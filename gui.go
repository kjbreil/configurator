package main

import (
	"fmt"
	"log"
	"path"

	"github.com/go-ini/ini"
	"github.com/kjbreil/sil"
	"github.com/kjbreil/sil/loc"
	"github.com/manifoldco/promptui"
)

func gui() {
	inf, _ := iniFile()

	s := iniSection(inf)
	s.Write("out.sil")
}

// make a SIL file from a section of the file
func iniSection(inf *ini.File) sil.SIL {
	fmt.Println()
	prompt := promptui.Select{
		Label: "Select Section: ",
		Items: inf.SectionStrings(),
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

	}

	sec := inf.Section(result)

	s := sil.Make("CFG", loc.CFG{})

	// declare the CFG table type
	var c loc.CFG

	// make the key to fill
	var k Key

	k.filename = filename
	k.section = sec.Name()
	for _, ele := range sec.Keys() {
		k.key = ele.Name()
		c.F2845 = k.String()
		c.F2847 = ele.Value()

		s.View.Data = append(s.View.Data, c)

		log.Println(k.String(), ele.Value())

	}

	return s
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
		Default:  "samples/system.ini",
	}
	// run the prompt
	_, err = prompt.Run()
	if err != nil {
		return inf, fmt.Errorf("prompt failed %v", err)
	}

	return inf, nil
}
