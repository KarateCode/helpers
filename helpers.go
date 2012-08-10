package helpers

import (
	"fmt"
	"regexp"
	"strings"
	"os"
	"runtime/debug"
)

const ( 
	Reset = "\x1b[0m"
	Bright = "\x1b[1m" 
	Dim = "\x1b[2m" 
	Underscore = "\x1b[4m" 
	Blink = "\x1b[5m" 
	Reverse = "\x1b[7m" 
	Hidden = "\x1b[8m" 
	FgBlack = "\x1b[30m" 
	FgRed = "\x1b[31m" 
	FgGreen = "\x1b[32m" 
	FgYellow = "\x1b[33m" 
	FgBlue = "\x1b[34m" 
	FgMagenta = "\x1b[35m" 
	FgCyan = "\x1b[36m" 
	FgWhite = "\x1b[37m" 
	BgBlack = "\x1b[40m" 
	BgRed = "\x1b[41m" 
	BgGreen = "\x1b[42m" 
	BgYellow = "\x1b[43m" 
	BgBlue = "\x1b[44m" 
	BgMagenta = "\x1b[45m" 
	BgCyan = "\x1b[46m" 
	BgWhite = "\x1b[47m" 
)

var filePath *regexp.Regexp

func init() {
	filePath = regexp.MustCompile("/[A-Za-z._]+:")
}

func ShouldEqual(comparer, comparee interface {}) {
	if comparer != comparee {
		fmt.Printf(FgRed + Bright + "\n    expected: " + FgWhite + "%v", comparer)
		fmt.Printf(FgRed +          "\n         got: " + FgWhite + "%v\n\n" + Reset, comparee)
		stack := string(debug.Stack())
		
		lines := strings.Split(stack, "\n")
		for _, line := range lines{
			if (strings.Contains(line, "helpers.go")) {
				continue
			}
			if (strings.Contains(line, "debug.Stack()")) {
				continue
			}
			if (strings.Contains(line, "testing.go:273")) {
				continue
			}
			if (strings.Contains(line, "test.F(t)")) {
				continue
			}
			if (strings.Contains(line, "proc.c:271")) {
				continue
			}
			if (strings.Contains(line, "goexit(void)")) {
				continue
			}
			indexes := filePath.FindStringIndex(line)
			if len(indexes) > 0 {
				match := filePath.FindString(line)
				fmt.Printf(line[:indexes[0]+1] + FgBlue + Bright + match[1:len(match)-1] + Reset + line[indexes[0]+len(match)-1:] + "\n")
			} else {
				println(line)
			}
		}
		os.Exit(-1)
	}
}

func IncludeExclude(excluded bool, origList []int, newList []int) []int {
	if excluded {
		return Union(origList, newList)
	}
	return Subtract(origList, newList)
}

func Union(a, b []int) []int {
	var c []int 
	for _, val := range a {
		if Index(b, val) == -1 {
			c = append(c, val)
		}
	}
	results := make([]int, len(c) + len(b))
	copy(results, c)
	copy(results[len(c):], b)
	return results
}

func Index(a []int, sep int) int {
	for i, val := range a {
		if val == sep {
			return i
		}
	}
	return -1
}

func Subtract(a, b []int) []int {
	var index int
	for _, num := range b {
		index = Index(a, num)
		if index != -1 {
			a = append(a[:index], a[index+1:]...)
		}
	}
	return a
}
