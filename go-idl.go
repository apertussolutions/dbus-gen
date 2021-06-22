// SPDX-License-Identifier: BSD-3-Clause
package main

var goSingleSigMap = map[string]string{
	"y": "byte",
	"b": "bool",
	"n": "int16",
	"q": "uint16",
	"i": "int32",
	"u": "uint32",
	"x": "int64",
	"t": "uint64",
	"d": "float64",
	"s": "string",
	"o": "dbus.ObjectPath",
	"v": "dbus.Variant",
	"h": "unixFDIndex",
}

var goDoubleSigMap = map[string]string{
	"ay": "[]byte",
	"ai": "[]int32",
	"au": "[]uint32",
	"at": "[]uint64",
	"as": "[]string",
	"ao": "[]dbus.ObjectPath",
}

var goComplexSigMap = map[string]string{
	"a{ss}":     "map[string]string",
	"a{su}":     "map[string]uint32",
	"a{sv}":     "map[string]dbus.Variant",
	"aa{ss}":    "[]map[string]string",
	"aa{sv}":    "[]map[string]dbus.Variant",
	"a{sa{sv}}": "map[string]map[string]dbus.Variant",
}

func GoType(dt string) string {
	if t, ok := goSingleSigMap[dt]; ok {
		return t
	}
	if t, ok := goDoubleSigMap[dt]; ok {
		return t
	}
	if t, ok := goComplexSigMap[dt]; ok {
		return t
	}
	return "interface{}"
}

func (a Arg) GoType() string {
        return GoType(a.Type)
}

func (p Property) GoType() string {
        return GoType(p.Type)
}
