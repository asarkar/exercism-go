package markdown

import (
	"fmt"
	"strings"
)

// Render translates a simplified Markdown subset to HTML.
//
// Supported:
//   - Headings: lines starting with '#' up to ######
//   - Bold: text enclosed in __double underscores__
//   - Italic: text enclosed in _single underscores_
//   - Unordered lists: lines starting with '*'
//   - Paragraphs: plain text lines
func Render(markdown string) string {
	var html strings.Builder
	lines := strings.Split(markdown, "\n")

	inList := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		switch {
		// --- Headings ---
		case strings.HasPrefix(line, "#"):
			level := headingLevel(line)
			if level > 6 {
				html.WriteString(fmt.Sprintf("<p>%s</p>", escape(line)))
				continue
			}
			content := strings.TrimSpace(line[level:])
			content = renderInline(content)
			fmt.Fprintf(&html, "<h%d>%s</h%d>", level, content, level)

		// --- List items ---
		case strings.HasPrefix(line, "* "):
			if !inList {
				html.WriteString("<ul>")
				inList = true
			}
			item := strings.TrimPrefix(line, "* ")
			item = renderInline(item)
			fmt.Fprintf(&html, "<li>%s</li>", item)

		// --- Paragraphs ---
		default:
			if inList {
				html.WriteString("</ul>")
				inList = false
			}
			html.WriteString("<p>")
			html.WriteString(renderInline(line))
			html.WriteString("</p>")
		}
	}

	if inList {
		html.WriteString("</ul>")
	}

	return html.String()
}

// headingLevel counts consecutive '#' characters at the start of a line.
func headingLevel(line string) int {
	level := 0
	for _, r := range line {
		if r == '#' {
			level++
		} else {
			break
		}
	}
	return level
}

// renderInline converts inline markdown (_emphasis_, __bold__) to HTML.
func renderInline(text string) string {
	text = renderInlineTag(text, "__", "strong")
	text = renderInlineTag(text, "_", "em")
	return text
}

func renderInlineTag(text, delim, tag string) string {
	const marker = "\x00" // temporary non-printable placeholder

	text = strings.ReplaceAll(text, delim, marker)
	parts := strings.Split(text, marker)

	for i := range parts {
		if i%2 == 1 {
			parts[i] = fmt.Sprintf("<%s>%s</%s>", tag, parts[i], tag)
		}
	}

	return strings.Join(parts, "")
}

// escape escapes any stray '#' symbols that exceed level 6 headings.
func escape(s string) string {
	return strings.ReplaceAll(s, "<", "&lt;")
}
