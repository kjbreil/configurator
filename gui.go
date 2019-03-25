package main

import (
	"fmt"
	"log"
	"path"

	"github.com/go-ini/ini"
	"github.com/manifoldco/promptui"
)

func gui() {
	inf, _ := iniFile()
	log.Println(inf)
}

// validate that a correct ini file was
func iniFile() (inf *ini.File, err error) {
	// ini file validation
	iniFile := func(input string) error {
		if path.Ext(input) != ".ini" {
			return fmt.Errorf("%s does not have an INI extension", input)
		}

		inf, err = ini.Load(input)
		if err != nil {
			return fmt.Errorf("%s cannot be opened: %v", input, err)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Name of INI file to load",
		Validate: iniFile,
	}
	// run the prompt
	_, err = prompt.Run()
	if err != nil {
		return inf, fmt.Errorf("prompt failed %v", err)
	}

	return inf, nil
}
