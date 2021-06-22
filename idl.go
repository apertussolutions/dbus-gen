// SPDX-License-Identifier: BSD-3-Clause
package main

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"
)

type DocString struct {
	XMLName xml.Name `xml:"docstring"`
	Body    string   `xml:",innerxml"`
}

type Arg struct {
	XMLName   xml.Name `xml:"arg"`
	Type      string   `xml:"type,attr"`
	Name      string   `xml:"name,attr"`
	Direction string   `xml:"direction,attr"`
}

type Method struct {
	XMLName xml.Name  `xml:"method"`
	Name    string    `xml:"name,attr"`
	Comment DocString `xml:"docstring"`
	Args    []Arg     `xml:"arg"`
}

func (m Method) Parameters() []Arg {
	params := make([]Arg, 0, 1)

	for _, p := range m.Args {
		if p.Direction == "in" {
			params = append(params, p)
		}
	}
	return params
}

func (m Method) Returns() []Arg {
	params := make([]Arg, 0, 1)

	for _, p := range m.Args {
		if p.Direction == "out" {
			params = append(params, p)
		}
	}
	return params
}

type Signal struct {
	XMLName xml.Name  `xml:"signal"`
	Name    string    `xml:"name,attr"`
	Comment DocString `xml:"docstring"`
	Args    []Arg     `xml:"arg"`
}

type Property struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Access  string   `xml:"access,attr"`
}

type Interface struct {
	XMLName    xml.Name   `xml:"interface"`
	Name       string     `xml:"name,attr"`
	Methods    []Method   `xml:"method"`
	Signals    []Signal   `xml:"signal"`
	Properties []Property `xml:"property"`
}

type Node struct {
	XMLName    xml.Name    `xml:"node"`
	Name       string      `xml:"name,attr"`
	Interfaces []Interface `xml:"interface"`
}

type Idl struct {
	Name string
	Path string
	Node *Node
}

func ParseIdlFile(path string) (*Node, error) {
	var node Node

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := xml.NewDecoder(f).Decode(&node); err != nil {
		return nil, err
	}

	return &node, nil
}

func NewIdl(path string) (*Idl, error) {
	var err error

	idl := Idl{
		Path: path,
	}

	fileName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	idl.Name = fileName
	idl.Node, err = ParseIdlFile(path)
	if err != nil {
		return nil, err
	}

	return &idl, nil
}
