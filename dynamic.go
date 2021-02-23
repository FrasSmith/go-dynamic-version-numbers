package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	// Version : current version
	Version string = strings.TrimSpace(version)
	//go:embed version.txt
	version string
)

func main() {
	fmt.Println("Dynamic Version Numbers in Go Example")
	vData := strings.Split(Version, "\n")
	if len(vData) != 3 {
		fmt.Println("Embedded version data is badly formed or missing")
	} else {
		VersionNumber := vData[0]
		BuildTime := vData[1]
		BuildNo := vData[2]

		fmt.Printf("Version       : %s\n", VersionNumber)
		fmt.Printf("Build Time    : %s\n", BuildTime)
		fmt.Printf("Build Number  : %s\n", BuildNo)
	}
}
