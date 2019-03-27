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

func gui() (err error) {
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

	cfg.F1056, err = target()
	if err != nil {
		return err
	}
	cfg.F2846, err = terminalGroup()
	if err != nil {
		return err
	}
	inf, err := iniFile()
	if err != nil {
		return err
	}

	section, s := iniSection(inf, cfg)
	// set file base, this includes extentions but not full path, set to upper
	// case
	fb := strings.ToUpper(path.Base(filename))
	// just the file name without extention
	f := strings.TrimSuffix(fb, path.Ext(fb))

	// set the batch # to the julian date, this "should" prevent collision
	s.Header.F902 = fmt.Sprintf("9%07v", sil.JulianNow())
	// #nosec
	s.Header.F913 = fmt.Sprintf("CONFIGURATOR UPDATE FOR %s %s", fb, section)

	// Write the SIL file useing the filename and section
	err = s.Write("CFG_" + f + "_" + section + ".sil")
	return err
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
