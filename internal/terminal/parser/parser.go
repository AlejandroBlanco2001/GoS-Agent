package parser

import (
	"strconv"
	"strings"
)

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

func ParseNetAdapterStatistics(output string) (map[string]map[string]int64, error) {
	lines := strings.Split(output, "\n")

	results := make(map[string]map[string]int64)

	for index, line := range lines {
		// Windows PowerShell command output has a header and a line
		if index == 0 || index == 1 {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 5 {
			continue
		}

		interfaceName := fields[0]

		// TODO: This is a dirty workaround, I need to use the command GetInterfaceNames to get the interface name properly and remove based on this output
		// Basically find if the strings is inside of the fields, remove it to that point and join the rest, in that way, i can avoid the if
		if len(fields) > 5 {
			interfaceName = strings.Join(fields[:len(fields)-4], " ")
			fields = append(fields[:1], fields[2:]...)
		}

		recievedBytes, err := strconv.ParseInt(fields[1], 10, 64)

		if err != nil {
			return nil, err
		}

		recievedUnicastPackets, err := strconv.ParseInt(fields[2], 10, 64)

		if err != nil {
			return nil, err
		}

		sentBytes, err := strconv.ParseInt(fields[3], 10, 64)

		if err != nil {
			return nil, err
		}

		sentUnicastPackets, err := strconv.ParseInt(fields[4], 10, 64)

		if err != nil {
			return nil, err
		}

		results[interfaceName] = map[string]int64{
			"ReceivedBytes":          recievedBytes,
			"ReceivedUnicastPackets": recievedUnicastPackets,
			"SentBytes":              sentBytes,
			"SentUnicastPackets":     sentUnicastPackets,
		}
	}

	return results, nil
}
