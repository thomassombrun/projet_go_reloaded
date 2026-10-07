package main

import (
	"os"
	"strconv"
	"strings"
)

func Process(texte string) string {
	mots := strings.Fields(texte)
	resultat := []string{}
	for _, mot := range mots {
		//si il y a (up) => change les lettres en Majuscule//
		if mot == "(up)" && len(resultat) > 0 {
			resultat[len(resultat)-1] = strings.ToUpper(resultat[len(resultat)-1])
			//si il y a (low) => change les lettres en Minuscule//
		} else if mot == "(low)" && len(resultat) > 0 {
			resultat[len(resultat)-1] = strings.ToLower(resultat[len(resultat)-1])
			//si il y a (cap) => met la première lettre en Majuscule et le reste en Minuscule//
		} else if mot == "(cap)" && len(resultat) > 0 {
			m := resultat[len(resultat)-1]
			resultat[len(resultat)-1] = strings.ToUpper(m[:1]) + strings.ToLower(m[1:])
			//si il y a (hex) => convertit le nombre hexadécimal en décimal//
		} else if mot == "(hex)" && len(resultat) > 0 {
			n, err := strconv.ParseInt(resultat[len(resultat)-1], 16, 64)
			if err == nil {
				resultat[len(resultat)-1] = strconv.FormatInt(n, 10)
			}
			//si il y a (bin) => convertit le nombre binaire en décimal//
		} else if mot == "(bin)" && len(resultat) > 0 {
			n, err := strconv.ParseInt(resultat[len(resultat)-1], 2, 64)
			if err == nil {
				resultat[len(resultat)-1] = strconv.FormatInt(n, 10)
			}
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
