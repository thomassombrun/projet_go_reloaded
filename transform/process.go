package transform

import (
	"strconv"
	"strings"
)

// La fonction Process applique toutes les modifications au texte et renvoie le résultat
func Process(texte string) string {
	// Découpe le texte en mots (séparés par des espaces)
	mots := strings.Fields(texte)
	// Resultat contient les mots à garder, après modifications
	resultat := []string{}
	// On parcourt avec un index i pour pouvoir lire le mot suivant (mots[i+1])
	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		// Si il y a (up) => change les lettres en Majuscule
		// len(resultat) > 0 : on vérifie qu'il y a bien un mot avant le marqueur
		if mot == "(up)" && len(resultat) > 0 {
			resultat[len(resultat)-1] = strings.ToUpper(resultat[len(resultat)-1])
			// Si il y a (low) => change les lettres en minuscules
		} else if mot == "(low)" && len(resultat) > 0 {
			resultat[len(resultat)-1] = strings.ToLower(resultat[len(resultat)-1])
			// Si il y a (cap) => met la première lettre en Majuscule et le reste en Minuscule
		} else if mot == "(cap)" && len(resultat) > 0 {
			m := resultat[len(resultat)-1] // le dernier mot gardé
			resultat[len(resultat)-1] = strings.ToUpper(m[:1]) + strings.ToLower(m[1:])
			// Si il y a (hex) => convertit le nombre hexadécimal en décimal
		} else if mot == "(hex)" && len(resultat) > 0 {
			n, err := strconv.ParseInt(resultat[len(resultat)-1], 16, 64)
			if err == nil { // on remplace seulement si c'est un vrai nombre
				resultat[len(resultat)-1] = strconv.FormatInt(n, 10)
			}
			// Si il y a (bin) => convertit le nombre binaire en décimal
		} else if mot == "(bin)" && len(resultat) > 0 {
			n, err := strconv.ParseInt(resultat[len(resultat)-1], 2, 64)
			if err == nil {
				resultat[len(resultat)-1] = strconv.FormatInt(n, 10)
			}
			// Si il y a (up, n), (low, n) ou (cap, n) => change les n mots précédents
			// strings.Fields coupe "(up, 2)" en deux morceaux : "(up," et "2)"
		} else if (mot == "(up," || mot == "(low," || mot == "(cap,") && i+1 < len(mots) {
			mode := mot[1 : len(mot)-1] // "(up," devient "up"
			// Le mot suivant est "2)" : on enlève ")" pour garder "2", puis on le convertit en nombre
			nombre := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(nombre)
			if err == nil {
				// On part de n mots avant la fin, sans descendre sous 0
				debut := len(resultat) - n
				if debut < 0 {
					debut = 0
				}
				// On applique la transformation à chaque mot concerné
				for j := debut; j < len(resultat); j++ {
					resultat[j] = transforme(resultat[j], mode)
				}
			}
			// On saute le mot suivant ("2)") pour qu'il ne soit pas ajouté au texte
			i++
			// Cas normal : un mot ordinaire, on le garde tel quel
		} else {
			resultat = append(resultat, mot)
		}
	}
	// Recolle les mots avec un espace entre chacun et corrige la ponctuation
	return CorrigePonctuation(CorrigeApostrophes(CorrigeArticles(strings.Join(resultat, " "))))
}
