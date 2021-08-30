package main

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	marker := flag.String("M", "", "The `marker` to replace in the file.")
	filename := flag.String("F", "", "The file to replace snippet in.")
	flag.Parse()
	if *marker == "" || *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := snippet(*filename, *marker, os.Stdin); err != nil {
		log.Fatal(err)
	}
}

// snippet replace all regions in the file enclosed in `<!-- marker -->` with
// the content in `contentReader`.
func snippet(filename string, marker string, contentReader io.Reader) error {
	content, err := ioutil.ReadAll(contentReader)
	if err != nil {
		return err
	}
	before, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(before))
	lines := make([]string, 0, 10)
	markerLine := "<!-- " + marker + " -->"
	var open bool
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case !open && line == markerLine:
			open = true
			lines = append(lines, line)
		case open && line == markerLine:
			lines = append(lines, string(content))
			lines = append(lines, line)
		case open:
		// noop, this removes lines within the region
		default:
			lines = append(lines, line)
		}
	}
	next := strings.Join(lines, "\n")
	return os.WriteFile(filename, []byte(next), 0600)
}
