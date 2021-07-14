package _task

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func NewManager() *Manager {
	return &Manager{
		cron: cron.New(cron.WithSeconds()),
	}
}

type Manager struct {
	address string
	cron    *cron.Cron
	lock    sync.Mutex
	tasks   map[string]Task
	entries map[string]cron.EntryID
}

func (p *Manager) SetAddress(address string) {
	p.address = address
}

func (p *Manager) SetCron(cron *cron.Cron) {
	p.cron = cron
}

func (p *Manager) Init(task ...Task) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	
	if p.tasks == nil {
		p.tasks = map[string]Task{}
	}
	
	for _, v := range task {
		_, ok := p.tasks[v.Slug()]
		if ok {
			return fmt.Errorf("duplicated task: %s", v.Slug())
		}
		p.tasks[v.Slug()] = v
	}
	
	return nil
}

func (p *Manager) Remove(slug string) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	
	if _, ok := p.tasks[slug]; !ok {
		return fmt.Errorf("task not exist")
	}
	
	entryId, ok := p.entries[slug]
	if !ok {
		return fmt.Errorf("task not add")
	}
	
	p.cron.Remove(entryId)
	
	return nil
}

func (p *Manager) Add(slug string) error {
	
	p.lock.Lock()
	defer p.lock.Unlock()
	
	task, ok := p.tasks[slug]
	if !ok {
		return fmt.Errorf("task not exist")
	}
	
	entryId, err := p.cron.AddJob(task.Spec(), task)
	if err != nil {
		return err
	}
	
	p.entries[slug] = entryId
	
	return nil
}

func (p *Manager) Run(slug string) error {
	
	p.lock.Lock()
	defer p.lock.Unlock()
	
	task, ok := p.tasks[slug]
	if !ok {
		return fmt.Errorf("task not exist")
	}
	
	if task.Running() {
		return fmt.Errorf("task is running")
	}
	
	task.Run()
	
	return nil
}

func (p *Manager) All() (tasks []map[string]interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for _, v := range p.tasks {
		tasks = append(tasks, map[string]interface{}{
			"slug":    v.Slug(),
			"spec":    v.Spec(),
			"name":    v.Name(),
			"running": v.Running(),
		})
	}
	return
}

func (p *Manager) Listen(serve bool, addr string) error {
	
	if p.entries == nil {
		p.entries = map[string]cron.EntryID{}
	}
	
	for _, v := range p.tasks {
		entryId, err := p.cron.AddFunc(v.Spec(), v.Run)
		if err != nil {
			return err
		}
		p.entries[v.Slug()] = entryId
	}
	p.cron.Start()
	
	if serve {
		go p.Serve(addr)
	}
	// TODO 等待任务执行完再退出
	// 平滑退出
	// 强制退出
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	p.cron.Stop()
	log.Println("cron manager force exit")
	return nil
}

func (p *Manager) Serve(addr string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("add", func(c echo.Context) error {
		err := p.Add(c.QueryParam("slug"))
		if err != nil {
			return c.JSON(400, map[string]interface{}{
				"error": err.Error(),
			})
		}
		return c.String(200, "ok")
	})
	e.GET("remove", func(c echo.Context) error {
		err := p.Remove(c.QueryParam("slug"))
		if err != nil {
			return c.JSON(400, map[string]interface{}{
				"error": err.Error(),
			})
		}
		return c.String(200, "ok")
	})
	e.GET("run", func(c echo.Context) error {
		err := p.Run(c.QueryParam("slug"))
		if err != nil {
			return c.JSON(400, map[string]interface{}{
				"error": err.Error(),
			})
		}
		return c.String(200, "ok")
	})
	e.GET("tasks", func(c echo.Context) error {
		return c.JSON(http.StatusOK, p.All())
	})
	if err := e.Start(addr); err != nil {
		log.Panicln(err)
	}
}
