package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/go-ini/ini"
	"github.com/kjbreil/sil"
	"github.com/kjbreil/sil/loc"
	"github.com/manifoldco/promptui"
)

func gui() {
	// declare the CFG table type
	cfg := loc.CFG{
		F1000: "999901",
		F1056: "999",
		F2846: "GROC_LANE",
		F253:  sil.JulianNow(),
		F940:  999,
		F941:  999,
		F1001: 1,
		F1264: sil.JulianNow(),
	}

	cfg.F1056, _ = target()
	cfg.F2846, _ = terminalGroup()

	inf, _ := iniFile()

	section, s := iniSection(inf, cfg)
	fmt.Println("CFG_" + filename + section + ".sil")

	f := path.Base(filename)
	f = strings.ToUpper(strings.TrimSuffix(f, path.Ext(f)))

	s.Write("CFG_" + f + "_" + section + ".sil")
}

// make a SIL file from a section of the file
func iniSection(inf *ini.File, cfg loc.CFG) (string, sil.SIL) {

	sections := inf.SectionStrings()
	sections = sections[1:]
	sections = append([]string{"ALL"}, sections...)
	prompt := promptui.Select{
		Label: "Select Section: ",
		Items: sections,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

	}

	s := sil.Make("CFG", loc.CFG{})

	if result == "ALL" {
		for _, sec := range sections[1:] {
			s.View.Data = append(s.View.Data, singleSection(sec, cfg, inf)...)
		}
	} else {
		s.View.Data = append(s.View.Data, singleSection(result, cfg, inf)...)
	}

	return result, s
}

// singleSection returns the CFG array of a single section, since sil.View.Data
// is an interface{} the return value also needs to be an interface :-(
func singleSection(name string, cfg loc.CFG, inf *ini.File) (cfgs []interface{}) {

	sec := inf.Section(name)

	// make the key to fill
	var k Key

	k.filename = filename
	k.section = sec.Name()
	for _, ele := range sec.Keys() {
		k.key = ele.Name()
		cfg.F2845 = k.String()
		cfg.F2847 = ele.Value()

		cfgs = append(cfgs, cfg)

	}
	return
}

// validate that a correct ini file was entered
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
		Default:  "samples/System.ini",
	}
	// run the prompt
	filename, err = prompt.Run()
	if err != nil {
		return inf, fmt.Errorf("prompt failed %v", err)
	}

	return inf, nil
}

func terminalGroup() (string, error) {
	validate := func(input string) error {

		switch {
		case strings.ToUpper(input) != input:
			return fmt.Errorf("group needs to be uppercase")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Terminal Group",
		Validate: validate,
		Default:  "GROC_LANE",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

func target() (string, error) {
	validate := func(input string) error {

		switch {
		case strings.ToUpper(input) != input:
			return fmt.Errorf("target needs to be uppercase")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Target?",
		Validate: validate,
		Default:  "PAL",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

// func directory() string {

// }
