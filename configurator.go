package main

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
	"github.com/kjbreil/sil"
	"github.com/kjbreil/sil/loc"
)

var (
	filename = "samples/system.ini"
)

// main function - keep less than 10 lines of code
func main() {
	// gui()
	cfg()
}

func cfg() {
	// make SIL file
	s := sil.Make("CFG", loc.CFG{})

	// import the ini file with go-ini
	inf, err := ini.Load(filename)

	if err != nil {
		log.Panicf("could not load ini: %v", err)
	}

	sec := inf.Section("SMS")
	var c loc.CFG
	c.F1000 = "001901"
	c.F1056 = "001"
	c.F2846 = "GROC_LANE"
	c.F253 = "2123131"
	c.F940 = 999
	c.F941 = 999
	c.F1001 = 1
	// c.F1264 = "0132156"

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

	// for _, section := range ini.Sections() {
	// 	// declare the CFG table type
	// 	var c loc.CFG

	// 	// make the key to fill
	// 	var k Key

	// 	k.filename = filename
	// 	k.section = section.Name()
	// 	for _, ele := range section.Keys() {
	// 		k.key = ele.Name()
	// 		c.F2845 = k.String()
	// 		c.F2847 = ele.Value()

	// 		s.View.Data = append(s.View.Data, c)

	// 	}
	// }
	// should read error return
	fmt.Println(s)

	s.Write("out.sil")
}
