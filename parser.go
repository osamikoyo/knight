package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var FILEPATHS []string = []string{
	"Knight.yaml",
	"Knight.yml",
	"KnightManifect.yml",
	"KnightManifect.yaml",
	"KnightMan.yml",
}

func ParseFile() (*Manifest, error) {
	var (
		man  Manifest
		file *os.File = nil
		err  error
	)
	for _, files := range FILEPATHS {
		file, err = os.Open(files)
		if err == nil {
			break
		}
	}

	if file == nil {
		return nil, errors.New("knight manifest not found")
	}

	defer file.Close()

	if err = yaml.NewDecoder(file).Decode(&man); err != nil {
		return nil, fmt.Errorf("error decode manifest: %v", err)
	}

	return &man, nil
}
