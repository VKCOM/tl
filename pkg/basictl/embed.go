package basictl

import (
	_ "embed"
	"strings"
)

//go:embed  basictl.go
var basicTLContent string

//go:embed  basictl2.go
var basicTL2Content string

func BasicTLContent(headerComment string, packageName string) string {
	const packageBasictl = "package basictl"
	ind := strings.Index(basicTLContent, packageBasictl)
	if ind < 0 {
		panic("basictl.go must contain line 'package basictl' after optional comment")
	}
	return headerComment + "\npackage " + packageName + basicTLContent[ind+len(packageBasictl):]
}

func BasicTL2Content(headerComment string, packageName string) string {
	const packageBasictl = "package basictl"
	ind := strings.Index(basicTL2Content, packageBasictl)
	if ind < 0 {
		panic("basictl2.go must contain line 'package basictl' after optional comment")
	}
	return headerComment + "\npackage " + packageName + basicTL2Content[ind+len(packageBasictl):]
}
