package cmd

import (
	"app/repository"
	"errors"
	"flag"
	"fmt"
)

type Remove struct {
	baseCommand

	idFlag   NullableInt
	allFlag  *bool
	nameFlag *string

	repository repository.Repository
}

func NewRemove(repository repository.Repository, path []string) Cmd {
	fs := flag.NewFlagSet("remove", flag.ContinueOnError)
	var idFlag NullableInt
	fs.Var(&idFlag, "id", "remove by card id")
	nameFlag := fs.String("name", "", "remove by name")
	allFlag := fs.Bool("all", false, "remove all cards")

	return &Remove{
		baseCommand: newBaseCommand(
			fs,
			[]Cmd{},
			path,
			"Remove cards from collection",
		),
		nameFlag: nameFlag,
		idFlag:   idFlag,
		allFlag:  allFlag,

		repository: repository,
	}
}

func (r *Remove) Run() error {
	if *r.helpFlag {
		fmt.Printf(r.Usage())
		return nil
	}

	var cards []repository.Card
	var err error
	if *r.allFlag {
		cards, err = r.repository.Delete(repository.CardFilter{})
	} else if *r.nameFlag != "" {
		cards, err = r.repository.Delete(repository.CardFilter{Name: r.nameFlag})
	} else {
		cards, err = r.repository.Delete(repository.CardFilter{ID: &r.idFlag.Val})
	}
	if err != nil {
		return err
	}

	fmt.Printf("successfully deleted %d cards\n", len(cards))
	return nil
}

func (r *Remove) Init(args []string) (bool, error) {
	if len(args) == 0 || args[0] != r.fs.Name() {
		return false, nil
	}
	if ok, err := r.baseCommand.Init(args[1:]); !ok {
		return false, err
	}

	if !r.idFlag.IsSet && !*r.allFlag && *r.nameFlag == "" && !*r.helpFlag {
		return false, errors.New(r.Usage())
	}
	return true, nil
}
