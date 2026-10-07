package main

import (
	"os"
	"strings"
)

func Process(texte string) string {
	mots := strings.Fields(texte)
	resultat := []string{}
	for _, mot := range mots {
		if mot == "(up)" && len(resultat) > 0 {
			resultat[len(resultat)-1] = strings.ToUpper(resultat[len(resultat)-1])
		} else if mot == "(low)" && len(resultat) > 0 {
			resultat[len(resultat)-1] = strings.ToLower(resultat[len(resultat)-1])
		} else if mot == "(cap)" && len(resultat) > 0 {
			m := resultat[len(resultat)-1]
			resultat[len(resultat)-1] = strings.ToUpper(m[:1]) + strings.ToLower(m[1:])
		} else {
			resultat = append(resultat, mot)
		}
	}
	return strings.Join(resultat, " ")
}

func main() {
	contenu, _ := os.ReadFile("sample.txt")
	os.WriteFile("result.txt", []byte(Process(string(contenu))), 0644)
}
