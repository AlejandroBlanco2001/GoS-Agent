package parser

import (
	"strconv"
	"strings"
)

func ParseNetStatOutputLinux(output string) map[string]map[string]string {
	lines := strings.Split(output, "\n")
	results := make(map[string]map[string]string)

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) < 7 {
			continue
		}

		results[fields[4]] = map[string]string{
			"protocol":     fields[0],
			"state":        fields[1],
			"recv-q":       fields[2],
			"send-q":       fields[3],
			"local":        fields[4],
			"Peer Address": fields[5],
		}
	}

	return results
}

func ParseNetAdapterStatisticsLinux(output string, interfaceNames []string) (map[string]map[string]int64, error) {
	lines := strings.Split(output, "\n")

	results := make(map[string]map[string]int64)
	nameIndex := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		fields := strings.Fields(line)

		if fields[0] == "RX:" {
			interfaceName := interfaceNames[nameIndex]

			recievedBytes, err := strconv.ParseInt(strings.Fields(lines[i+1])[0], 10, 64)

			if err != nil {
				return nil, err
			}

			sentBytes, err := strconv.ParseInt(strings.Fields(lines[i+3])[0], 10, 64)

			if err != nil {
				return nil, err
			}

			results[interfaceName] = map[string]int64{
				"ReceivedBytes": recievedBytes,
				"SentBytes":     sentBytes,
			}

			nameIndex++
		}
	}
	return results, nil
}

func ParseInterfaceNamesLinux(output string) ([]string, error) {
	lines := strings.Split(output, "\n")

	results := make([]string, 0)

	// iterate two by two startring from zero
	for i := 0; i < len(lines); i += 2 {
		fields := strings.Fields(lines[i])

		interfaceName := strings.Replace(fields[1], ":", "", -1)

		results = append(results, interfaceName)
	}

	return results, nil
}
