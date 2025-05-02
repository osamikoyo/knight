package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
	Purpur = "\033[35m"
)

type Pipe struct {
	Name     string   `yaml:"name"`
	Cmds     []string `yaml:"cmds"`
	TimeoutS int      `yaml:"timeout"`
}

type Manifest struct {
	ProjectName string `yaml:"project"`
	Pipeline    []Pipe `yaml:"pipeline"`
}

func (p *Pipe) Run(_ context.Context) error {
	for _, cmd := range p.Cmds {
		words := strings.Split(cmd, " ")
		now := time.Now()

		timeStr := now.Format("15:04:05")

		fmt.Printf(Purpur+"[%s]: %s"+Reset, timeStr, cmd)

		ecmd := exec.Command(words[0], words[0:]...)
		ecmd.Stderr = os.Stderr
		ecmd.Stdout = os.Stdout
		if err := ecmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manifest) Run(pipename string) error {
	var ctx context.Context

	for _, pipe := range m.Pipeline {
		if pipe.Name != pipename {
			continue
		}

		if pipe.TimeoutS != 0 {
			subCtx, cancel := context.WithTimeout(context.Background(), time.Duration(pipe.TimeoutS)*time.Second)
			defer cancel()

			ctx = subCtx
		} else {
			ctx = context.Background()
		}

		if err := pipe.Run(ctx); err != nil {
			return err
		}

		return nil
	}

	return errors.New("pipe not found")
}
