package render

import (
	"go-reloaded/transform"
	"strings"
)

func Lines(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, l := range lignes {
		lignes[i] = transform.Process(l)
	}
	return strings.Join(lignes, "\n")
}
