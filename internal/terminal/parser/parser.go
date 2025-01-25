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

func ParseNetAdapterStatistics(output string, interfaceNames []string) (map[string]map[string]int64, error) {
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

		// Hack: The command that brings the interface name and the one that hold the statistics are the same
		// so the output is the same, but the fields are different, so I can substract the lenght of the name splitted to get the statistics
		// Therefore, the index-2 is the interface row
		interfaceName := interfaceNames[index-2]
		nameOffset := len(strings.Fields(interfaceName))
		statisticsArray := fields[nameOffset:]

		recievedBytes, err := strconv.ParseInt(statisticsArray[0], 10, 64)

		if err != nil {
			return nil, err
		}

		recievedUnicastPackets, err := strconv.ParseInt(statisticsArray[1], 10, 64)

		if err != nil {
			return nil, err
		}

		sentBytes, err := strconv.ParseInt(statisticsArray[2], 10, 64)

		if err != nil {
			return nil, err
		}

		sentUnicastPackets, err := strconv.ParseInt(statisticsArray[3], 10, 64)

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

func ParseInterfaceNames(output string) ([]string, error) {
	lines := strings.Split(output, "\n")

	results := make([]string, 0)

	for index, line := range lines {
		line = strings.TrimSpace(line)

		// This is just to help the test cases, and avoid calling functions from the terminal package (cleaning)
		if line == "" {
			continue
		}

		if index == 0 || index == 1 {
			continue
		}

		results = append(results, line)
	}

	return results, nil
}
