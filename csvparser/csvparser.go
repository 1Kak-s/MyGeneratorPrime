package csvparser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"MyGeneratorPrime/pirate"
)

func Parse(filePath string) ([]pirate.Pirate, error) {

	// Ouvre le fichier CSV
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier : %w", err)
	}
	defer file.Close()

	var pirates []pirate.Pirate

	// lit ligne par ligne
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		// ignore les lignes vides
		if line == "" {
			continue
		}

		// On découpe la ligne par ";"
		parts := strings.Split(line, ";")

		if len(parts) != 3 {
			fmt.Printf("⚠️  Ligne %d ignorée (format invalide) : %s\n", lineNumber, line)
			continue
		}

		name := strings.TrimSpace(parts[0])
		prime := strings.TrimSpace(parts[1])
		img := strings.TrimSpace(parts[2])

		// enlève les guillemets du chemin de l'image
		img = strings.Trim(img, "\"")

		// crée le pirate et valide les champs
		p, err := pirate.New(name, prime, img)
		if err != nil {
			fmt.Printf(" Ligne %d ignorée (erreur) : %s\n", lineNumber, err)
			continue
		}

		pirates = append(pirates, *p)
	}

	if len(pirates) == 0 {
		return nil, fmt.Errorf("aucun pirate trouvé dans le fichier")
	}

	return pirates, nil
}
