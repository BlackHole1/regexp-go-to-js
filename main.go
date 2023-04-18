package reg

import (
	"fmt"
	"regexp"
	"strings"
)

func GoToJs(goRegex string) string {
	// \xhh -> \u00HH
	re := regexp.MustCompile(`\\x([0-9a-fA-F]{2})`)
	goRegex = re.ReplaceAllStringFunc(goRegex, func(match string) string {
		hex := match[2:]
		return fmt.Sprintf("\\u%04s", strings.ToUpper(hex))
	})

	// \x{hhhh} -> \uHHHH
	re = regexp.MustCompile(`\\x\{([0-9a-fA-F]{4})}`)
	goRegex = re.ReplaceAllStringFunc(goRegex, func(match string) string {
		hex := match[3 : len(match)-1]
		return fmt.Sprintf("\\u%s", strings.ToUpper(hex))
	})

	return goRegex
}

