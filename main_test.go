package main

import "testing"

func TestTransforme(t *testing.T) {
	tests := []struct {
		mot, mode, want string
	}{
		{"go", "up", "GO"},
		{"SHOUTING", "low", "shouting"},
		{"bridge", "cap", "Bridge"},
		{"BRIDGE", "cap", "Bridge"},
		{"", "cap", ""},
	}
	for _, tc := range tests {
		got := transforme(tc.mot, tc.mode)
		if got != tc.want {
			t.Errorf("transforme(%q, %q) = %q, want %q", tc.mot, tc.mode, got, tc.want)
		}
	}
}

func TestConvertit(t *testing.T) {
	tests := []struct {
		mot  string
		base int
		want string
	}{
		{"1E", 16, "30"},
		{"10", 2, "2"},
		{"salut", 16, "salut"},
	}
	for _, tc := range tests {
		got := convertit(tc.mot, tc.base)
		if got != tc.want {
			t.Errorf("convertit(%q, %d) = %q, want %q", tc.mot, tc.base, got, tc.want)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct{ in, want string }{
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"Welcome to the Brooklyn bridge is nice (cap, 3)", "Welcome to the Brooklyn Bridge Is Nice"},
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
	}
	for _, tc := range tests {
		got := Process(tc.in)
		if got != tc.want {
			t.Errorf("Process(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestPonctuation(t *testing.T) {
	tests := []struct{ in, want string }{
		{"Hello , world", "Hello, world"},
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"What ?! Really", "What?! Really"},
	}
	for _, tc := range tests {
		got := CorrigePonctuation(tc.in)
		if got != tc.want {
			t.Errorf("CorrigePonctuation(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestApostrophes(t *testing.T) {
	tests := []struct{ in, want string }{
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"hello ' world ' and ' bye '", "hello 'world' and 'bye'"},
	}
	for _, tc := range tests {
		got := CorrigeApostrophes(tc.in)
		if got != tc.want {
			t.Errorf("CorrigeApostrophes(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestArticles(t *testing.T) {
	tests := []struct{ in, want string }{
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"a apple and a banana", "an apple and a banana"},
		{"a hour", "an hour"},
	}
	for _, tc := range tests {
		got := CorrigeArticles(tc.in)
		if got != tc.want {
			t.Errorf("CorrigeArticles(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
