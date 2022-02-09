package xstring

import (
	"log"
	"strings"
)

// Between - extract from input string all substrings located between begin and end.
// Useful for extracting data between HTML tags.
func Between(input string, begin string, end string) []string {
	var result []string

	for {
		log.Printf("%v", input)
		var beginIndex = strings.Index(input, begin)
		if beginIndex == -1 {
			break
		}
		part := input[beginIndex+len(begin):]
		endIndex := strings.Index(part, end)
		if endIndex == -1 {
			break
		}
		finalPart := part[:endIndex]
		result = append(result, finalPart)

		input = input[beginIndex+len(begin)+len(finalPart):]
	}

	return result

}
