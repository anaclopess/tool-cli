package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "uso: taillog <arquivo>")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	sc := bufio.NewScanner(f)
	lines := []string{}
	for sc.Scan() {
		lines = append(lines, sc.Text())
		if len(lines) > 80 {
			lines = lines[1:]
		}
	}
	for _, ln := range lines {
		if strings.TrimSpace(ln) != "" {
			fmt.Println(ln)
		}
	}
}
