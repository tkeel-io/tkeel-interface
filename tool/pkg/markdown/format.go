/*
Copyright 2021 The tKeel Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package markdown

import (
	"github.com/russross/blackfriday"
	"github.com/shurcooL/markdownfmt/markdown"
	"io/ioutil"
	"os"
)

// Process formats Markdown.
// If opt is nil the defaults are used.
// Error can only occur when reading input from filename rather than src.
func FormatMarkdown(src []byte, opt *markdown.Options) ([]byte, error) {
	// extensions for GitHub Flavored Markdown-like parsing.
	const extensions = blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK

	output := blackfriday.Markdown(src, markdown.NewRenderer(opt), extensions)
	return output, nil
}



func FormatMarkdownFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	src, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	res, err := FormatMarkdown(src, &markdown.Options{})
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, res, 0)
	if err != nil {
		return err
	}

	return err
}

/*
// markdownfmt formats Markdown.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/scanner"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/shurcooL/markdownfmt/markdown"
	"golang.org/x/term"
)


func report(err error) {
	scanner.PrintError(os.Stderr, err)
	exitCode = 2
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: markdownfmt [flags] [path ...]\n")
	flag.PrintDefaults()
}

func isMarkdownFile(f os.FileInfo) bool {
	// Ignore non-Markdown files.
	name := f.Name()
	return !f.IsDir() && !strings.HasPrefix(name, ".") && (strings.HasSuffix(name, ".md") || strings.HasSuffix(name, ".markdown"))
}

func processFile(filename string, in io.Reader, out io.Writer, stdin bool) error {
	if in == nil {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		in = f
	}

	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	isTerminal := func() bool {
		return term.IsTerminal(int(os.Stdout.Fd())) && os.Getenv("TERM") != "dumb"
	}
	res, err := markdown.Process(filename, src, &markdown.Options{
		Terminal: !*list && !*write && !*doDiff && isTerminal(),
	})
	if err != nil {
		return err
	}

	if !bytes.Equal(src, res) {
		// formatting has changed
		if *list {
			fmt.Fprintln(out, filename)
		}
		if *write {
			err = ioutil.WriteFile(filename, res, 0)
			if err != nil {
				return err
			}
		}
		if *doDiff {
			data, err := diff(src, res)
			if err != nil {
				return fmt.Errorf("computing diff: %s", err)
			}
			fmt.Printf("diff %s markdownfmt/%s\n", filename, filename)
			out.Write(data)
		}
	}

	if !*list && !*write && !*doDiff {
		_, err = out.Write(res)
	}

	return err
}

func visitFile(path string, f os.FileInfo, err error) error {
	if err == nil && isMarkdownFile(f) {
		err = processFile(path, nil, os.Stdout, false)
	}
	if err != nil {
		report(err)
	}
	return nil
}

func walkDir(path string) {
	filepath.Walk(path, visitFile)
}

func main() {
	// call markdownfmtMain in a separate function
	// so that it can use defer and have them
	// run before the exit.
	markdownfmtMain()
	os.Exit(exitCode)
}

func markdownfmtMain() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		if err := processFile("<standard input>", os.Stdin, os.Stdout, true); err != nil {
			report(err)
		}
		return
	}

	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)
		switch dir, err := os.Stat(path); {
		case err != nil:
			report(err)
		case dir.IsDir():
			walkDir(path)
		default:
			if err := processFile(path, nil, os.Stdout, false); err != nil {
				report(err)
			}
		}
	}
}

func diff(b1, b2 []byte) (data []byte, err error) {
	f1, err := ioutil.TempFile("", "markdownfmt")
	if err != nil {
		return
	}
	defer os.Remove(f1.Name())
	defer f1.Close()

	f2, err := ioutil.TempFile("", "markdownfmt")
	if err != nil {
		return
	}
	defer os.Remove(f2.Name())
	defer f2.Close()

	f1.Write(b1)
	f2.Write(b2)

	data, err = exec.Command("diff", "-u", f1.Name(), f2.Name()).CombinedOutput()
	if len(data) > 0 {
		// diff exits with a non-zero status when the files don't match.
		// Ignore that failure as long as we get output.
		err = nil
	}
	return
}


*/