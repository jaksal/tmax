package main

import (
	"strings"
)

func findStr(src string, fs [][]string) bool {
	//log.Println("find...", src, fs)

	upperSrc := strings.ToUpper(src)

	// or check
	for _, fo := range fs {
		// and check..
		exist := func(fas []string) bool {
			for _, s := range fas {
				if !strings.Contains(upperSrc, strings.ToUpper(s)) {
					return false
				}
			}
			return true
		}(fo)

		if exist {
			// log.Println("match", src, fo)
			return true
		}
	}
	// log.Println("mismatch", src, fs)
	return false
}
