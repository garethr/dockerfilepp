//
// dockerfilepp provides a very simple library for building Dockerfile
// post-processors. These applications take a Dockerfile on stdin,
// replace (using Go templates) a set of passed in replacements, and then
// output the results to stdout.
//
package dockerfilepp

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

// A struct to provide a way of injecting values into Go templates
type Context struct {
	Value string
}

// render a template along with some context. Exit if we
// hit a problem with parsing or processing any of the templates
func render(temp string, args string) string {
	context := Context{args}
	tmpl, err := template.New("replacer").Parse(temp)
	if err != nil {
		fmt.Println("A problem occured parsing one of the processors:")
		fmt.Println(err)
		fmt.Print(temp)
		os.Exit(2)
	}
	buff := bytes.NewBufferString("")
	err = tmpl.Execute(buff, context)
	if err != nil {
		fmt.Println("A problem occured executing one of the processors:")
		fmt.Println(err)
		fmt.Print(temp)
		os.Exit(2)
	}
	return buff.String()
}

// Main entrypoint for building a post processor. Takes a map of replacements
// and a usage instructions
func Process(replacements map[string]string, docstring string) {
	stat, _ := os.Stdin.Stat()
	// We detect whether we have anything on stdin to process
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var buffer bytes.Buffer
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			buffer.WriteString(scanner.Text() + "\n")
		}
		value := buffer.String()
		for regex, tmpl := range replacements {
			re := regexp.MustCompile(regex + "(.*)")
			matches := re.FindStringSubmatch(value)
			if len(matches) == 2 {
				args := matches[1]
				args = strings.TrimLeft(args, " ")
				value = re.ReplaceAllString(value, render(tmpl, args))
			}
		}
		fmt.Print(value)
	} else {
		fmt.Println(docstring)
	}
}
