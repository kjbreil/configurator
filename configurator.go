package main

import (
	"log"

	"github.com/go-ini/ini"
	"github.com/kjbreil/sil"
	"github.com/kjbreil/sil/loc"
)

var (
	filename = "./sample.ini"
)

// main function - keep less than 10 lines of code
func main() {
	gui()
	// cfg()
}

func cfg() {
	// make SIL file
	s := sil.Make("CFG", loc.CFG{})

	// import the ini file with go-ini
	ini, err := ini.Load(filename)

	if err != nil {
		log.Panicf("could not load ini: %v", err)
	}

	for _, section := range ini.Sections() {
		// declare the CFG table type
		var c loc.CFG

		// make the key to fill
		var k Key

		k.filename = filename
		k.section = section.Name()
		for _, ele := range section.Keys() {
			k.key = ele.Name()
			c.F2845 = k.String()
			c.F2847 = ele.Value()

			s.View.Data = append(s.View.Data, c)

		}
	}
	// should read error return
	s.Write("out.sil")
}
