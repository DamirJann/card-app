package main

import (
	"app/cmd"
	"app/repository"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version string `yaml:"version"`
	DB      struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"db"`
}

func (cfg *Config) Init() error {
	file, err := os.OpenFile("config/configuration.yaml", 0, 0)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		fmt.Println(fmt.Errorf("failed to parse configuration.yaml: %v", err))

	}

	file, err = os.OpenFile("config/secret.yaml", 0, 0)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return fmt.Errorf("failed to parse secret.yaml: %v", err)

	}

	if cfg.Version == "" {
		return fmt.Errorf("app version is not set")
	}
	return nil
}

func main() {
	var cfg Config
	if err := cfg.Init(); err != nil {
		fmt.Println(fmt.Errorf("failed to init config: %v", err))
		os.Exit(1)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=default port=5432 sslmode=disable", cfg.DB.Host, cfg.DB.User, cfg.DB.Password)
	repository, err := repository.NewPostgres(dsn)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		os.Exit(1)
	}

	rootCMD := cmd.NewRoot(cfg.Version, repository)
	_, err = rootCMD.Init(os.Args[1:])
	if err != nil {
		fmt.Println(fmt.Errorf("failed to parse command.\n%v", err))
		os.Exit(1)
	}

	err = rootCMD.Run()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to run command: %v", err))
		os.Exit(1)
	}

}
