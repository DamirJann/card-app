package cmd

import (
	"app/repository"
	"errors"
	"flag"
	"fmt"
)

type Add struct {
	baseCommand

	nameFlag   *string
	answerFlag *string

	repository repository.Repository
}

func NewAdd(repository repository.Repository, path []string) Cmd {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	nameFlag := fs.String("name", "", "name or title of a card")
	answerFlag := fs.String("answer", "", "card back side or answer")

	return &Add{
		nameFlag:   nameFlag,
		answerFlag: answerFlag,
		baseCommand: newBaseCommand(
			fs,
			[]Cmd{},
			path,
			"Add new cards to collection",
		),
		repository: repository,
	}
}

func (a *Add) Run() error {
	if *a.helpFlag {
		fmt.Printf(a.Usage())
		return nil
	}

	err := a.repository.Create(*a.nameFlag, *a.answerFlag)
	if err != nil {
		return err
	}
	fmt.Printf("successfully added \n")
	return nil
}

func (a *Add) Init(args []string) (bool, error) {
	if len(args) == 0 || args[0] != a.fs.Name() {
		return false, nil
	}
	if ok, err := a.baseCommand.Init(args[1:]); !ok {
		return false, err
	}

	if *a.nameFlag == "" && *a.answerFlag == "" && !*a.helpFlag {
		return false, errors.New(a.Usage())
	}
	if (*a.nameFlag == "" || *a.answerFlag == "") && !*a.helpFlag {
		return false, errors.New(a.Usage())
	}

	return true, nil
}
