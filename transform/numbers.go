package transform

import "strconv"

// Convertit un mot représentant un nombre dans une base donnée en décimal //
// Exemple : convertit("1E", 16) renvoie "30" //
func convertit(mot string, base int) string {
	// Lit le mot comme un nombre écrit dans la base donnée (16 ou 2) //
	n, err := strconv.ParseInt(mot, base, 64)
	if err != nil {
		return mot // ce n'est pas un vrai nombre : on laisse le mot tel quel //
	}
	// Remet le nombre en texte, écrit en décimal (base 10) //
	return strconv.FormatInt(n, 10)
}
