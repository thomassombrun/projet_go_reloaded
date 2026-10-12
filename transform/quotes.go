package transform

import "strings"

func CorrigeApostrophes(texte string) string {
	mots := strings.Fields(texte)
	resultat := []string{}
	ouvert := false  // true quand on est entre deux apostrophes //
	aColler := false // true si l'apostrophe ouvrante doit être collée au mot suivant //

	for _, mot := range mots {
		if mot == "'" {
			if !ouvert {
				// Première apostrophe : elle ira devant le mot suivant //
				ouvert = true
				aColler = true
			} else {
				// Deuxième apostrophe : on la colle derrière le dernier mot de resultat //
				if len(resultat) > 0 {
					resultat[len(resultat)-1] = resultat[len(resultat)-1] + "'"
				}
				ouvert = false
			}
		} else if ouvert && strings.HasPrefix(mot, "'") && strings.Trim(mot[1:], ".,!?:;") == "" {
			// Apostrophe fermante suivie de ponctuation (ex : "'.") : on colle le tout au dernier mot //
			if len(resultat) > 0 {
				resultat[len(resultat)-1] += mot
			}
			ouvert = false
		} else {
			// Mot normal : s'il vient juste après une apostrophe ouvrante, on ajoute "'" devant //
			if aColler {
				mot = "'" + mot
				aColler = false
			}
			resultat = append(resultat, mot)
		}
	}
	return strings.Join(resultat, " ")
}
