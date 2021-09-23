package bot

import (
	"html"
	"regexp"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
)

func trimDescription(desc string, limit int) string {
	if limit == 0 {
		return ""
	}
	desc = strings.Trim(
		strip.StripTags(
			regexp.MustCompile("\n+").ReplaceAllLiteralString(
				strings.ReplaceAll(
					regexp.MustCompile(`<br(| /)>`).ReplaceAllString(
						html.UnescapeString(desc), "<br>"),
					"<br>", "\n"),
				"\n")),
		"\n")

	contentDescRune := []rune(desc)
	if len(contentDescRune) > limit {
		desc = string(contentDescRune[:limit])
	}

	return desc
}

// python
//    if int(version) == 1:
//        escape_chars = r'_*`['
//    elif int(version) == 2:
//        if entity_type in ['pre', 'code']:
//            escape_chars = r'\`'
//        elif entity_type == 'text_link':
//            escape_chars = r'\)'
//        else:
//            escape_chars = r'_*[]()~`>#+-=|{}.!'
//    else:
//        raise ValueError('Markdown version must be either 1 or 2!')
//
//    return re.sub(f'([{re.escape(escape_chars)}])', r'\\\1', text)
func EscapeMarkdown(text string) string {
	replacer := strings.NewReplacer("*", "\\*", "_", "\\_", "[", "\\[", "]", "\\]", "`", "\\`", "(", "\\(", "-", "\\-", ")", "\\)", ".", "\\.", "=", "\\=", "+", "\\+", "~", "\\~", "!", "\\!", "|", "\\|", "#", "\\#", ">", "\\>", "{", "\\{", "}", "\\}", "'", "\\'")
	return replacer.Replace(text)
}
