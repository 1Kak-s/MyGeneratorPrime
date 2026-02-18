package generateprime

import (
	"MyGeneratorPrime/pdf"
	"MyGeneratorPrime/pirate"
	"fmt"
)

// Generate crée un seul avis de recherche
func Generate(name, prime, img, outputDir string) error {

	// 1. Créer le pirate
	p, err := pirate.New(name, prime, img)
	if err != nil {
		return fmt.Errorf("erreur création pirate : %w", err)
	}

	// 2. Créer le PdfSaver
	saver, err := pdf.New(outputDir)
	if err != nil {
		return fmt.Errorf("erreur création saver : %w", err)
	}

	// 3. Générer le PDF
	return saver.Save(*p)
}
