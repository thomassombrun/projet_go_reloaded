package transform

import "strings"

// CorrigeArticles transforme "a" en "an" quand le mot suivant commence par une voyelle ou un h //
func CorrigeArticles(texte string) string {
	mots := strings.Fields(texte)
	// len(mots)-1 : on s'arrête avant le dernier mot, car il n'a pas de mot suivant //
	for i := 0; i < len(mots)-1; i++ {
		if mots[i] == "a" || mots[i] == "A" {
			// Première lettre du mot suivant //
			premiere := []rune(mots[i+1])[0]
			if strings.ContainsRune("aeiouhAEIOUH", premiere) {
				if mots[i] == "a" {
					mots[i] = "an"
				} else {
					mots[i] = "An" // On garde la majuscule //
				}
			}
		}
	}
	return strings.Join(mots, " ")
}
