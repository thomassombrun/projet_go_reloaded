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
		{"If I make you breakfast in bed just say thanks instead of: 'I am nothing without you'. Simply add 66 and 2 and you will see the result is 68. There is no greater agony than bearing a untold story inside you.", "If I make you breakfast in bed just say thanks instead of: 'I am nothing without you'. Simply add 66 and 2 and you will see the result is 68. There is no greater agony than bearing an untold story inside you."},
		{"", ""},
		{"go (up, 10)", "GO"},        // plus de mots demandés que disponibles
		{"ZZ (hex) done", "ZZ done"}, // pas un nombre hexa : le mot reste inchangé
		{"1010 (bin) and ff (hex)", "10 and 255"}, // hexa en minuscules accepté
		{"hello (cap) (up)", "HELLO"},             // deux marqueurs à la suite
		{"hello    world", "hello world"},         // espaces multiples
		{"' hello world '", "'hello world'"},      // apostrophes autour de plusieurs mots
		{"don't stop", "don't stop"},              // apostrophe dans un mot : on n'y touche pas
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"Hello ,world", "Hello, world"},
		{"a hour", "an hour"},
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
