package main

import (
	"fmt"     // Bibliothèque qui sert a gérer les opérations d'entrée et de sortie formatées //
	"os"      // Bibliothèque pour lire et écrire des fichiers //
	"strconv" // Bibliothèque pour convertir des nombres en texte et inversement //
	"strings" // Bibliothèque pour manipuler du texte (majuscules, minuscules, etc...) //
)

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

// La fonction Process applique toutes les modifications au texte et renvoie le résultat //
func Process(texte string) string {
	// Découpe le texte en mots (séparés par des espaces) //
	mots := strings.Fields(texte)
	// Resultat contient les mots à garder, après modifications //
	resultat := []string{}
	// On parcourt avec un index i pour pouvoir lire le mot suivant (mots[i+1]) //
	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		// Si il y a (up) => change les lettres en Majuscule //
		// len(resultat) > 0 : on vérifie qu'il y a bien un mot avant le marqueur //
		if mot == "(up)" && len(resultat) > 0 {
			resultat[len(resultat)-1] = strings.ToUpper(resultat[len(resultat)-1])
			// Si il y a (low) => change les lettres en minuscules //
		} else if mot == "(low)" && len(resultat) > 0 {
			resultat[len(resultat)-1] = strings.ToLower(resultat[len(resultat)-1])
			// Si il y a (cap) => met la première lettre en Majuscule et le reste en Minuscule //
		} else if mot == "(cap)" && len(resultat) > 0 {
			m := resultat[len(resultat)-1] // le dernier mot gardé //
			resultat[len(resultat)-1] = strings.ToUpper(m[:1]) + strings.ToLower(m[1:])
			// Si il y a (hex) => convertit le nombre hexadécimal en décimal //
		} else if mot == "(hex)" && len(resultat) > 0 {
			n, err := strconv.ParseInt(resultat[len(resultat)-1], 16, 64)
			if err == nil { // on remplace seulement si c'est un vrai nombre //
				resultat[len(resultat)-1] = strconv.FormatInt(n, 10)
			}
			// Si il y a (bin) => convertit le nombre binaire en décimal //
		} else if mot == "(bin)" && len(resultat) > 0 {
			n, err := strconv.ParseInt(resultat[len(resultat)-1], 2, 64)
			if err == nil {
				resultat[len(resultat)-1] = strconv.FormatInt(n, 10)
			}
			// Si il y a (up, n), (low, n) ou (cap, n) => change les n mots précédents //
			// strings.Fields coupe "(up, 2)" en deux morceaux : "(up," et "2)" //
		} else if (mot == "(up," || mot == "(low," || mot == "(cap,") && i+1 < len(mots) {
			mode := mot[1 : len(mot)-1] // "(up," devient "up"
			// Le mot suivant est "2)" : on enlève ")" pour garder "2", puis on le convertit en nombre //
			nombre := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(nombre)
			if err == nil {
				// On part de n mots avant la fin, sans descendre sous 0 //
				debut := len(resultat) - n
				if debut < 0 {
					debut = 0
				}
				// On applique la transformation à chaque mot concerné //
				for j := debut; j < len(resultat); j++ {
					resultat[j] = transforme(resultat[j], mode)
				}
			}
			// On saute le mot suivant ("2)") pour qu'il ne soit pas ajouté au texte //
			i++
			// Cas normal : un mot ordinaire, on le garde tel quel //
		} else {
			resultat = append(resultat, mot)
		}
	}
	// Recolle les mots avec un espace entre chacun et corrige la ponctuation //
	return CorrigePonctuation(CorrigeApostrophes(CorrigeArticles(strings.Join(resultat, " "))))
}

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

// La fonction ProcessLines traite le texte ligne par ligne pour conserver les retours a la lignes //
func ProcessLines(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, l := range lignes {
		lignes[i] = Process(l)
	}
	return strings.Join(lignes, "\n")
}

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
	err = os.WriteFile(os.Args[2], []byte(ProcessLines(string(contenu))), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
	}
}
