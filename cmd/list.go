package cmd

import (
	"app/repository"
	"flag"
	"fmt"
)

type List struct {
	baseCommand
	repository repository.Repository
}

func (l List) Description() string {
	return "list all cards"
}

func NewList(repository repository.Repository, path []string) Cmd {
	return &List{
		baseCommand: newBaseCommand(
			flag.NewFlagSet("list", flag.ContinueOnError),
			[]Cmd{},
			path,
			"List all cards",
		),
		repository: repository,
	}
}

func (l List) Run() error {
	if *l.helpFlag {
		fmt.Printf(l.Usage())
		return nil
	}

	cards, err := l.repository.Find()
	if err != nil {
		return err
	}
	fmt.Printf("total: %d\n", len(cards))
	for _, card := range cards {
		fmt.Printf("#%d. %v\n", card.ID, card.Name)
	}
	return nil
}

func (l List) Init(args []string) (bool, error) {
	if len(args) == 0 || args[0] != l.fs.Name() {
		return false, nil
	}
	return l.baseCommand.Init(args[1:])
}
