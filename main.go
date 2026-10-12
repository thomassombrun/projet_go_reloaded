package main

import (
	"fmt"
	"go-reloaded/render"
	"os"
)

func main() {
	// L'énoncé veut 2 arguments : le fichier d'entrée (sample.txt) et le fichier de sortie (result.txt) //
	// os.Args[0] est le nom du programme, donc il y en a 3 en tout //
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <sample.txt> <result.txt>")
		return
	}

	// Lit le fichier d'entrée et affiche l'erreur s'il n'existe pas //
	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

	// Applique les modifications puis écrit le résultat dans le fichier de sortie //
	//0644 = permissions standard d'un fichier texte //
	err = os.WriteFile(os.Args[2], []byte(render.Lines(string(contenu))), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
	}
}
