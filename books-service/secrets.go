package main

var (
	AvaibleClients = []string{
		"gateway",
	}
	Secret = "books"
)

func ContainsKey(key string) bool {
	for _, k := range AvaibleClients {
		if k == key {
			return true
		}
	}
	return false
}
