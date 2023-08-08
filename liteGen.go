//go:build afero_gen

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type symbolType string

const (
	constType   symbolType = "const"
	varType     symbolType = "var"
	funcType    symbolType = "func"
	typeType    symbolType = "type"
	packageType symbolType = "package"
)

func main() {
	f, err := os.Create("lite.go")
	check(err)
	defer f.Close()
	_, err = fmt.Fprint(f, `// Code generated with liteGen.go DO NOT EDIT.
package afero

import (
	"errors"

	lite "github.com/tbhartman/afero/lite"
)

// manually alias ErrOutOfRange
var ErrOutOfRange = errors.New("out of range")
`)
	check(err)

	var b strings.Builder
	cmd := exec.Command("go", "doc", "./lite")
	cmd.Stdout = &b
	err = cmd.Run()
	check(err)

	mapping := map[symbolType][]string{}

	s := bufio.NewScanner(strings.NewReader(b.String()))
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		sp := strings.SplitN(line, " ", 3)
		if len(sp) < 2 {
			panic("unexpected line: " + line)
		}
		var lineType = symbolType(sp[0])
		name := sp[1]
		switch lineType {
		case funcType:
			name = strings.Split(name, "(")[0]
		case typeType:
		case constType:
		case varType:
		case packageType:
			continue
		default:
			panic("unexpected line: " + line)
		}

		mapping[lineType] = append(mapping[lineType], name)
	}

	for k, names := range mapping {
		if len(names) == 0 {
			continue
		}
		var typeName = string(k)
		fmt.Fprintln(f, "")
		fmt.Fprint(f, "// aliasing ")
		switch k {
		case constType:
			fmt.Fprint(f, "constants")
		case varType:
			fmt.Fprint(f, "variables")
		case funcType:
			fmt.Fprint(f, "functions")
			typeName = "var"
		case typeType:
			fmt.Fprint(f, "types")
		}
		fmt.Fprintln(f, " from lite package")
		sort.Strings(names)
		for _, name := range names {
			fmt.Fprintf(f, "%s %s = lite.%[2]s\n", typeName, name)
		}
	}
	check(s.Err())
}
