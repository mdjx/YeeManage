package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Reads all .json files in the current folder
// and encodes them as strings literals in data.go
func main() {
	fs, _ := ioutil.ReadDir("includes")
	out, _ := os.Create("data.go")
	out.Write([]byte("package main \n\nconst (\n"))
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".json") {
			out.Write([]byte(strings.TrimSuffix(f.Name(), ".json") + " = `"))
			f, err := os.Open("includes\\" + f.Name())

			if err != nil {
				fmt.Println(err)
			}

			io.Copy(out, f)
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}
