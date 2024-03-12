package cmd

import (
	"app/repository"
	"errors"
	"flag"
	"fmt"
)

const (
	rootName = "app"
)

type RunFunc func() error

type Cmd interface {
	Run() error
	Init(args []string) (bool, error)

	Usage() string
	Name() string
	Description() string
}

type Root struct {
	baseCommand

	version     string
	versionFlag *bool
	run         RunFunc

	repository repository.Repository
}

func NewRoot(version string, repository repository.Repository) Cmd {
	fs := flag.NewFlagSet(rootName, flag.ContinueOnError)
	vf := fs.Bool("v", false, "app version")
	return &Root{
		version:     version,
		repository:  repository,
		versionFlag: vf,
		baseCommand: newBaseCommand(
			fs,
			[]Cmd{
				NewAdd(repository, []string{rootName}),
				NewList(repository, []string{rootName}),
				NewRemove(repository, []string{rootName}),
			},
			[]string{},
			"System for keeping cards with answers",
		),
	}
}

func (v *Root) Init(args []string) (bool, error) {
	found, err := v.InitFlag(args)
	if err != nil {
		return false, err
	}
	if found {
		return true, nil
	}

	found, err = v.InitSubcommand(args)
	if err != nil {
		return false, err
	}
	if found {
		return true, nil
	}
	return false, errors.New(v.Usage())
}

func (v *Root) Run() error {
	return v.run()
}

func (v *Root) InitSubcommand(args []string) (bool, error) {
	for _, cmd := range v.subCommands {
		found, err := cmd.Init(args)
		if err != nil {
			return false, err
		}
		if found {
			v.run = cmd.Run
			return true, nil
		}
	}
	return false, nil
}

func (v *Root) InitFlag(args []string) (bool, error) {
	if err := v.fs.Parse(args); err != nil {
		return false, err
	}
	if !*v.versionFlag && !*v.helpFlag {
		return false, nil
	}
	v.run = v.flagRun
	return true, nil
}

func (v *Root) flagRun() error {
	if *v.versionFlag {
		fmt.Printf("app version %s\n", v.version)
	} else {
		fmt.Printf(v.Usage())
	}
	return nil
}
