package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
	b, err := ioutil.ReadFile("target.xml")
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
		rs := make([]string, rCon)
		vs := make([]string, rCon)
		for _, c := range t.Column {
			rs[i] = c
			i++
		}

		i = 0
		for _, r := range t.Row {
			for _, v := range r.Value {
				vs[i] = v
				i++
			}
		}
		showFormatedValue(tN, rs, vs)
	}

}

// print php-array-format
func showFormatedValue(tN string, rs, vs []string) {
	fmt.Printf("'%s' => array (\n", tN)
	i := 0
	for _, r := range rs {
		fmt.Printf("    '%s' => '%s',\n", r, vs[i])
	}
	fmt.Println(");")
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
