package parser

import "testing"

func TestParseNetStatOutput(t *testing.T) {
	const mockResultOfTheCommand = `
		Active Connections

		Proto  Local Address          Foreign Address        State
		TCP    127.0.0.1:51718        127.0.0.1:51719        ESTABLISHED
		TCP    192.168.1.10:55432     93.184.216.34:443      CLOSE_WAIT
		UDP    0.0.0.0:123            *:*                   
	`

	expected := map[string]map[string]string{
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

	result := ParseNetStatOutput(mockResultOfTheCommand)

	for key, value := range expected {
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
}
