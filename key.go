package main

import (
	"fmt"
	"path"
	"strings"
)

// Key is the data making F2845
type Key struct {
	filename string
	section  string
	key      string
}

func (k *Key) String() string {
	justFile := strings.ToUpper(path.Base(k.filename))
	return fmt.Sprintf("%s[%s]%s", justFile, k.section, k.key)
}
