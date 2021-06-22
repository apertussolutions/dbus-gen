// SPDX-License-Identifier: BSD-3-Clause
package main

import (
	"regexp"
	"strings"
)

func Increment(i int) int {
	return i + 1
}

func Deduct(i int) int {
	return i - 1
}

func Initial(s string) string {
	return string([]byte(s)[0])
}

func addSplit(name, token, delim string) string {
	idx := strings.Index(name, token)
	if idx == -1 {
		return name
	}

	idx += len(token)
	if string(name[idx]) == delim {
		return name
	}

	return name[:idx] + delim + name[idx:]
}

func removeSplit(name, token, delim string) string {
	idx := strings.Index(name, token)
	if idx == -1 {
		return name
	}

	idx += len(token)
	if string(name[idx]) == delim {
		return name[:idx] + name[idx+1:]
	}

	return name
}

func SplitCamelCase(r rune) bool {
	return r == '.' || r == '-' || r == '_'
}

func CamelCase(name string, private bool) string {
	words := strings.FieldsFunc(name, SplitCamelCase)

	for i, w := range words {
		words[i] = strings.Title(w)
	}
	if private {
		words[0] = strings.ToLower(words[0])
	}

	return strings.Join(words, "")
}

func Underscore(s string) string {
	var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}

	combined := strings.ToLower(strings.Join(a, "_"))

	combined = removeSplit(combined, "xen", "_")

	return combined
}

func Hyphen(s string) string {
	var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}

	combined := strings.ToLower(strings.Join(a, "-"))

	combined = removeSplit(combined, "xen", "-")

	return combined
}
