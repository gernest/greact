package vdom

import (
	"bytes"
	"encoding/xml"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse([]byte("<ul><li>one</li><li>two</li><li>three</li></ul>"))
	}
}

func BenchmarkXMLDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer([]byte("<ul><li>one</li><li>two</li><li>three</li></ul>"))
		dec := xml.NewDecoder(buf)
		for _, err := dec.Token(); err == nil; _, err = dec.Token() {
		}
	}
}

func BenchmarkDiff(b *testing.B) {
	oldTree, _ := Parse([]byte("<ul><li>one</li><li>two</li><li>three</li></ul>"))
	newTree, _ := Parse([]byte("<ul><li>uno</li><li>two</li><li>three</li></ul>"))
	for i := 0; i < b.N; i++ {
		Diff(oldTree, newTree)
	}
}
