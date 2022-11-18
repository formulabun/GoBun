package strings

import "strings"

func Escape(inString string, toEscape, escapeChar rune) string {
  var sb strings.Builder
  for _, c := range inString {
    if c == toEscape {
      sb.WriteRune(escapeChar)
    }
    sb.WriteRune(c)
  }
  return sb.String()
}
