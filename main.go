package main

import (
	"os"
	"strconv"
	"strings"
)

// Renvoie le mot avec la modification voulue selon le mode (up, low, cap) pour éviter les répétitions)//
func transforme(mot, mode string) string {
	switch mode {
	case "up":
		return strings.ToUpper(mot)
	case "low":
		return strings.ToLower(mot)
	case "cap":
		if mot == "" {
			return mot
		}
		return strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:])
	}
	return mot
}

// Convertit un mot représentant un nombre dans une base donnée en décimal//
func convertit(mot string, base int) string {
	n, err := strconv.ParseInt(mot, base, 64)
	if err != nil {
		return mot
	}
	return strconv.FormatInt(n, 10)
}

func Process(texte string) string {
	mots := strings.Fields(texte)
	resultat := []string{}
	for i := 0; i < len(mots); i++ {
		mot := mots[i]
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
			//si il y a (up, n), (low, n) ou (cap, n) => change les n mots précédents//
		} else if (mot == "(up," || mot == "(low," || mot == "(cap,") && i+1 < len(mots) {
			mode := mot[1 : len(mot)-1] // "(up," devient "up"//
			nombre := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(nombre)
			// on part de n mots avant la fin, sans descendre sous 0//
			if err == nil {
				debut := len(resultat) - n
				if debut < 0 {
					debut = 0
				}
				// on applique la transformation à chaque mot concerné//
				for j := debut; j < len(resultat); j++ {
					resultat[j] = transforme(resultat[j], mode)
				}
			}
			i++
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
