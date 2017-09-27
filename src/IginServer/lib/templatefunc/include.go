package templatefunc

import (
	"io/ioutil"
	"log"
	// "os"
)

func include(file string) string {
	b, err := ioutil.ReadFile("templates/" + file)
	if err != nil {
		log.Println(err)
	}
	s := string(b)
	return s
}
