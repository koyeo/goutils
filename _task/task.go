package _task

type Task interface {
	Slug() string
	Name() string
	Spec() string
	Running() bool
	Run()
}
