package parser

import "strings"

func ParseNetStatOutput(output string) []map[string]string {
	lines := strings.Split(output, "\n")
	var results []map[string]string

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}
		entry := map[string]string{
			"protocol":        fields[0],
			"local_address":   fields[1],
			"foreign_address": fields[2],
			"state":           fields[3],
		}
		results = append(results, entry)
	}

	return results
}
