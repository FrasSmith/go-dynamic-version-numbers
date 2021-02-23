package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Read the existing version.txt file from the project root
// Update build time and increment build number before
// writing back to version.txt for embedding in the
// main binary file
//
func main() {
	vFileData, _ := os.ReadFile("version.txt")

	vLines := strings.Split(string(vFileData), "\n")
	vno := vLines[0]
	bTime := time.Now().Format("20060102-1504")
	bNum, _ := strconv.Atoi(vLines[2])
	bNum++

	outStr := vno + "\n" + bTime + "\n" + fmt.Sprint(bNum)

	_ = os.WriteFile("version.txt", []byte(outStr), 0777)
}
