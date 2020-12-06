package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetData(path string, splitBy string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not read data from file \"%s\".\n", path)
		os.Exit(1)
	}
	return strings.Split(string(b), splitBy)
}
