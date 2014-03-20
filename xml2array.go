package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	f    = flag.String("f", "", "filename")
	file = flag.String("file", "", "filename long option")
)

type Dataset struct {
	Table []Table
}

type Table struct {
	Column []string
	Row    []Row
	Name   string `xml:"name,attr"`
}

type Row struct {
	Value []string
}

func (t Table) String() string {
	return t.Name
}

func main() {
	flag.Parse()
	if *f == "" && *file == "" {
		fmt.Println("Missing filename (use --help for help)")
		return
	}
	fn := *f
	if *file != "" {
		fn = *file
	}

	b, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Println("file open error: ", err.Error())
		return
	}

	var data Dataset
	b = convertGoReadebleFormat(b)
	xml.Unmarshal(b, &data)

	// parse xml, create column, value strings
	for _, t := range data.Table {
		i := 0

		rCon := len(t.Column)
		tN := t.Name

		fmt.Printf("'%s' => array(\n", tN)

		for _, r := range t.Row {
			i = 0
			vs := make([]string, rCon)
			for _, v := range r.Value {
				vs[i] = v
				i++
			}
			showFormatedValue(t.Column, vs)
		}

		fmt.Println(");")
	}

}

// print php-array-format
func showFormatedValue(cs, vs []string) {
	i := 0
	fmt.Println("    array(")
	for _, c := range cs {

		fmt.Printf("        '%s' => '%s',\n", c, vs[i])
		i++
	}
	fmt.Println("    ),")
}

func convertGoReadebleFormat(b []byte) []byte {
	b2 := b
	// <: 60, /: 47, diff: 32, a: 97, z: 122
	i := 0
	for i < len(b) {
		if b[i] == 60 {
			i++
			if b[i] >= 97 && b[i] <= 122 {
				b2[i] = b[i] - 32
			}
			if b[i] == 47 {
				i++
				if b[i] >= 97 && b[i] <= 122 {
					b2[i] = b[i] - 32
				}
			}
		}
		i++
	}
	return b2
}
