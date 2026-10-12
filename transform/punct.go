package transform

func estPonct(r rune) bool {
	return r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';'
}

func CorrigePonctuation(texte string) string {
	runes := []rune(texte) // On travaille caractère par caractère //
	out := []rune{}        // Le texte corrigé, construit petit à petit //

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if estPonct(r) {
			// On retire les espaces juste avant la ponctuation //
			for len(out) > 0 && out[len(out)-1] == ' ' {
				out = out[:len(out)-1]
			}
			// On ajoute la ponctuation, collée au mot précédent //
			out = append(out, r)
			// Espace après, sauf si le suivant est un espace, une ponctuation, ou la fin du texte //
			if i+1 < len(runes) && runes[i+1] != ' ' && !estPonct(runes[i+1]) {
				out = append(out, ' ')
			}
		} else {
			out = append(out, r)
		}
	}
	return string(out)
}
