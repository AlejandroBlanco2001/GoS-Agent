package parser

import (
	"strings"
	"testing"
)

func TestParseNetStatOutput(t *testing.T) {
	const mockResultOfTheCommandWithConnections = `
		Active Connections

		Proto  Local Address          Foreign Address        State
		TCP    127.0.0.1:51718        127.0.0.1:51719        ESTABLISHED
		TCP    192.168.1.10:55432     93.184.216.34:443      CLOSE_WAIT
		UDP    0.0.0.0:123            *:*                   
	`

	expectedConnections := map[string]map[string]string{
		"127.0.0.1:51718": {
			"protocol":        "TCP",
			"foreign_address": "127.0.0.1:51719",
			"state":           "ESTABLISHED",
		},
		"192.168.1.10:55432": {
			"protocol":        "TCP",
			"foreign_address": "93.184.216.34:443",
			"state":           "CLOSE_WAIT",
		},
	}

	var tests = []struct {
		name  string
		input string
		want  map[string]map[string]string
	}{
		{"it should return the statistics of the connection", mockResultOfTheCommandWithConnections, expectedConnections},
		{"it should return an empty map if the input is empty", "", map[string]map[string]string{}},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			result := ParseNetStatOutput(strings.TrimSpace(it.input))

			for key, value := range it.want {
				if result[key]["protocol"] != value["protocol"] {
					t.Errorf("Expected %s, but got %s", value["protocol"], result[key]["protocol"])
				}
				if result[key]["foreign_address"] != value["foreign_address"] {
					t.Errorf("Expected %s, but got %s", value["foreign_address"], result[key]["foreign_address"])
				}
				if result[key]["state"] != value["state"] {
					t.Errorf("Expected %s, but got %s", value["state"], result[key]["state"])
				}
			}
		})
	}
}

func TestParseNetAdapterStatistics(t *testing.T) {
	mockResultOfTheCommandWithInterfaces := `
	Name           ReceivedBytes ReceivedUnicastPackets SentBytes   SentUnicastPackets
	----           ------------- ---------------------- ---------   ------------------
	Ethernet0      1048576000    1500000                2097152000 3000000
	Wi-Fi          524288000     800000                 1073741824 1600000
	`

	expectedInterfaces := map[string]map[string]int64{
		"Ethernet0": {
			"ReceivedBytes": 1048576000,
			"SentBytes":     2097152000,
		},
		"Wi-Fi": {
			"ReceivedBytes": 524288000,
			"SentBytes":     1073741824,
		},
	}

	interfacesAvailable := []string{"Ethernet0", "Wi-Fi"}

	var tests = []struct {
		name       string
		input      string
		interfaces []string
		want       map[string]map[string]int64
	}{
		{"it should return the statistics of the interfaces", mockResultOfTheCommandWithInterfaces, interfacesAvailable, expectedInterfaces},
		{"it should return an empty map if the input is empty", "", []string{}, map[string]map[string]int64{}},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			result, _ := ParseNetAdapterStatistics(strings.TrimSpace(it.input), it.interfaces)

			if len(result) != len(it.want) {
				t.Errorf("Expected %d, but got %d", len(it.want), len(result))
			}

			for networkInterface, stadistics := range it.want {
				for parameter, value := range stadistics {
					if result[networkInterface][parameter] != value {
						t.Errorf("Expected %d, but got %d", value, result[networkInterface][parameter])
					}
				}
			}
		})
	}
}

func TestParseInterfaceNames(t *testing.T) {
	const mockResultOfTheCommandWithInterfaces = `Name
		----
		Wi-Fi
		Ethernet 2
		Local Area Connection* 1
	`

	expectedCommandWithInterfaces := []string{"Wi-Fi", "Ethernet 2", "Local Area Connection* 1"}

	var tests = []struct {
		name  string
		input string
		want  []string
	}{
		{"it should return an array with the inferfaces if we have interfaces", mockResultOfTheCommandWithInterfaces, expectedCommandWithInterfaces},
		{"it should return an empty array if we don't have interfaces", "", []string{}},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			result, _ := ParseInterfaceNames(it.input)

			if len(result) != len(it.want) {
				t.Errorf("Expected %d, but got %d", len(it.want), len(result))
			}

			for i, v := range it.want {
				if result[i] != v {
					t.Errorf("Expected %s, but got %s", v, result[i])
				}
			}
		})
	}
}
