package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

func InitKnight() error {
	file, err := os.Create("Knight.yaml")
	if err != nil {
		return err
	}

	man := &Manifest{
		ProjectName: "knight",
		Pipeline: []Pipe{
			{
				Cmds: []string{"echo hello", "echo kapi"},
			},
			{
				Cmds:     []string{"knight is cool", "very cool"},
				TimeoutS: 10,
			},
		},
	}

	if err = yaml.NewEncoder(file).Encode(man); err != nil {
		return err
	}

	return err
}
