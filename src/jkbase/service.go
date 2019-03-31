package jkbase

import (
	"github.com/kardianos/service"
	"jk/jklog"
)

type Program struct {
	Name        string
	DisplayName string
	Desc        string

	s      service.Service
	Runner func()
}

func NewProgram(name, displayname, desc string) *Program {
	return &Program{
		Name:        name,
		DisplayName: displayname,
		Desc:        desc,
	}
}

func (p *Program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *Program) run() {
	// Do work here
	p.Runner()
	p.s.Stop()
}
func (p *Program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func (p *Program) CreateService() error {
	svcConfig := &service.Config{
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Description: p.Desc,
	}

	s, err := service.New(p, svcConfig)
	if err != nil {
		return err
	}
	p.s = s
	return nil
}

func (p *Program) Ctrl(cmd string) error {
	switch cmd {
	case "run":
		err := p.Run()
		if err != nil {
			jklog.L().Errorln("Run service failed: ", err)
			return err
		}
	case "stop":
		err := p.s.Stop()
		if err != nil {
			jklog.L().Errorln("Stop service failed: ", err)
			return err
		}
	case "install":
		err := p.Install()
		if err != nil {
			jklog.L().Errorln("install service failed ", err)
			return err
		}
	case "remove":
		err := p.Uninstall()
		if err != nil {
			jklog.L().Errorln("remove service failed ", err)
			return err
		}
	}
	return nil
}

func (p *Program) Run() error {
	return p.s.Run()
}

func (p *Program) Install() error {
	return p.s.Install()
}

func (p *Program) Uninstall() error {
	return p.s.Uninstall()
}
