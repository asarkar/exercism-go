package grep

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type flags struct {
	filename        bool
	lineNumbers     bool
	nameOnly        bool
	invert          bool
	caseInsensitive bool
	wholeLine       bool
}

func newFlags(flagList []string, multipleFiles bool) flags {
	return flags{
		filename:        multipleFiles,
		lineNumbers:     slices.Contains(flagList, "-n"),
		nameOnly:        slices.Contains(flagList, "-l"),
		invert:          slices.Contains(flagList, "-v"),
		caseInsensitive: slices.Contains(flagList, "-i"),
		wholeLine:       slices.Contains(flagList, "-x"),
	}
}

// `Search` mimics grep function with flags.
func Search(pattern string, flagList []string, files []string) []string {
	fs := newFlags(flagList, len(files) > 1)

	regex := makePattern(fs, pattern)
	re := regexp.MustCompile(regex)

	var result []string
	for _, f := range files {
		//nolint:gosec // exercise expects reading files
		data, err := os.ReadFile(f)
		if err != nil {
			log.Printf("error reading file %s: %v", f, err)
			continue
		}

		result = append(result, grep(f, string(data), re, fs)...)
	}

	return result
}

// `makePattern` constructs the regex pattern based on flags.
func makePattern(fs flags, pattern string) string {
	var builder strings.Builder

	if fs.caseInsensitive {
		builder.WriteString("(?i)")
	}
	if fs.wholeLine {
		builder.WriteString("^")
	}

	builder.WriteString(fmt.Sprintf("(?:%s)", pattern))

	if fs.wholeLine {
		builder.WriteString("$")
	}

	return builder.String()
}

// grep processes the content of a single file and returns matching lines
func grep(filename, content string, re *regexp.Regexp, fs flags) []string {
	var result []string
	lines := strings.Split(content, "\n")

	for lineNum, line := range lines {
		if line == "" {
			continue
		}

		matched := re.MatchString(line)
		if fs.invert {
			matched = !matched
		}
		if !matched {
			continue
		}

		if fs.nameOnly {
			result = append(result, filename)
			break
		}

		var builder strings.Builder
		if fs.filename {
			builder.WriteString(filename + ":")
		}
		if fs.lineNumbers {
			builder.WriteString(strconv.Itoa(lineNum+1) + ":")
		}
		builder.WriteString(line)

		result = append(result, builder.String())
	}

	return result
}
