package _cli

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"reflect"
)

const (
	FLAG_CONFIG = "config"
	FLAG_C      = "c"
)

type App struct {
	Cmd cobra.Command
}

func (p *App) InitCmd(use, usage string) {
	p.Cmd.Use = ""
	p.Cmd.Flags().StringP(FLAG_CONFIG, FLAG_C, "", "config file path")
}

func (p *App) InitConfig(pointer interface{}) error {
	rv := reflect.TypeOf(pointer)
	if rv.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("expected pointer,go %s", rv))
	}
	if err := viper.BindPFlags(p.Cmd.Flags()); err != nil {
		return errors.Wrap(err, "failed to bind p.Cmd flags")
	}
	
	filePath, err := p.Cmd.Flags().GetString(FLAG_CONFIG)
	if err != nil {
		return errors.Wrap(err, "p.Cmd get config flag error")
	}
	
	viper.SetConfigFile(filePath)
	if err = viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "read config file error")
	}
	
	if err = viper.Unmarshal(pointer); err != nil {
		return errors.Wrap(err, "config unmarshal failed")
	}
	return nil
}
