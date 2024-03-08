package cmd

import (
	"bytes"
	"flag"
	"log"
	"math"
	"sort"
	"strings"
	"text/template"
)

func newBaseCommand(fs *flag.FlagSet, subCommands []Cmd, prevPath []string, description string) baseCommand {
	fs.SetOutput(bytes.NewBufferString(""))
	return baseCommand{
		fs:          fs,
		helpFlag:    fs.Bool("h", false, "help"),
		subCommands: subCommands,
		path:        append(prevPath, fs.Name()),
		description: description,
	}
}

type baseCommand struct {
	fs          *flag.FlagSet
	helpFlag    *bool
	subCommands []Cmd
	path        []string
	description string
}

func (bc *baseCommand) Name() string {
	return bc.fs.Name()
}

func (bc *baseCommand) Description() string {
	return bc.description
}

func (bc *baseCommand) Init(args []string) (bool, error) {
	if err := bc.fs.Parse(args); err != nil {
		return false, err
	}
	return true, nil
}

func align(rows []string) (res int) {
	for _, s := range rows {
		res = int(math.Max(float64(len(s)), float64(res)))
	}
	return
}

func names(mus []ManUnit) (res []string) {
	for _, s := range mus {
		res = append(res, s.Name)
	}
	return
}

func (bc *baseCommand) Usage() string {
	var commands, options []ManUnit
	for _, subCommand := range bc.subCommands {
		commands = append(commands, ManUnit{
			Name:        subCommand.Name(),
			Description: subCommand.Description(),
		})
	}
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Name < commands[j].Name
	})

	bc.fs.VisitAll(func(f *flag.Flag) {
		options = append(options, ManUnit{
			Name:        f.Name,
			Description: f.Usage,
		})
		sort.Slice(options, func(i, j int) bool {
			return options[i].Name < options[j].Name
		})
	})
	var s string
	res := bytes.NewBufferString(s)
	tmpl, err := template.New("man.tmpl").Funcs(template.FuncMap{
		"align": align,
		"names": names,
	}).ParseFiles("templates/man.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(res, struct {
		Description string
		SubCommands []ManUnit
		Options     []ManUnit
		PathToCMD   string
	}{
		Description: bc.Description(),
		SubCommands: commands,
		Options:     options,
		PathToCMD:   strings.Join(bc.path, " "),
	})
	if err != nil {
		log.Fatal(err)
	}
	return res.String()
}
