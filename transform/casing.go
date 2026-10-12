package transform

import "strings"

// Renvoie le mot avec la modification voulue selon le marqueur (up, low, cap) pour éviter les répétitions //
// Exemple : transforme("go", "up") renvoie "GO"//
func transforme(mot, mode string) string {
	switch mode {
	case "up":
		return strings.ToUpper(mot) // Modifie tout en majuscules //
	case "low":
		return strings.ToLower(mot) // Modifie tout en minuscules //
	case "cap":
		if mot == "" { // Sécurité : mot[:1] planterait sur un mot vide //
			return mot
		}
		// Met la 1re lettre en majuscule + tout le reste en minuscules //
		return strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:])
	}
	return mot // Marqueur inconnu : on renvoie le mot sans le modifier //
}
