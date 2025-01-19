package parser

import "strings"

func ParseNetStatOutput(output string) map[string]map[string]string {
	lines := strings.Split(output, "\n")
	results := make(map[string]map[string]string)

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) < 4 {
			continue
		}

		results[fields[1]] = map[string]string{
			"protocol":        fields[0],
			"foreign_address": fields[2],
			"state":           fields[3],
		}
	}

	return results
}
