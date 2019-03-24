package main

import "fmt"

// Key is the data making F2845
type Key struct {
	filename string
	section  string
	key      string
}

func (k *Key) String() string {
	return fmt.Sprintf("%s[%s}%s", k.filename, k.section, k.key)
}
