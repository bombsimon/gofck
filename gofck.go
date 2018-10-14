package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	out := flag.String("out", "", "name to compile to")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: gofuck <braincode.bf>")
		fmt.Println("       or")
		fmt.Println("       gofuck --out bf-app <braincode.bf>")
		os.Exit(0)
	}

	content, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	r := compile(content)

	if *out != "" {
		buildApp(r, *out)
		return
	}

	fmt.Println(string(r))
}

func buildApp(source []byte, appName string) {
	boilerPlate := `package main
import "fmt"
func main() { fmt.Println(string([]byte("%s"))) }`

	ss := string(source)
	ss = strings.Replace(ss, "\n", "\\n", -1)

	fc := fmt.Sprintf(boilerPlate, ss)

	fn := "tmp.go"
	if _, err := os.Stat(fn); !os.IsNotExist(err) {
		fmt.Printf("a file with the temporary name %s already exists, cannot continue\n", fn)
		return
	}

	f, _ := os.Create(fn)
	defer os.Remove(fn)

	f.Write([]byte(fc))
	f.Close()

	cmd := exec.Command("go", "build", "-o", appName, "tmp.go")
	cmd.Run()
}

func compile(c []byte) []byte {
	a := []int{0}
	b := []byte{}
	i := 0
	bm := map[int]int{}

	if strings.Contains(string(c), "[") {
		p := []int{}

		for i, v := range c {
			switch v {
			case byte('['):
				p = append(p, i)
			case byte(']'):
				bm[i] = p[len(p)-1]
				p = p[:len(p)-1]
			}
		}

		for k, v := range bm {
			bm[v] = k
		}
	}

	for ii := 0; ii < len(c); ii++ {
		switch c[ii] {
		case byte('>'):
			if i == len(a)-1 {
				a = append(a, 0)
			}

			i++
		case byte('<'):
			if i == 0 {
				a = append([]int{0}, a...)
			}

			i--
		case byte('+'):
			a[i]++
			if a[i] > 255 {
				a[i] = 0
			}
		case byte('-'):
			a[i]--
			if a[i] < 0 {
				a[i] = 255
			}
		case byte('.'):
			b = append(b, byte(a[i]))
		case byte(','):
			// TODO
		case byte('['):
			if a[i] == 0 {
				ii = bm[ii]
			}
		case byte(']'):
			if a[i] != 0 {
				ii = bm[ii]
			}
		}
	}

	return b
}
