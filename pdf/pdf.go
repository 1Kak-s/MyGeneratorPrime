package pdf

import (
	"MyGeneratorPrime/pirate"
	"fmt"
	"os"

	"github.com/signintech/gopdf"
)

// ou enregistrer les PDF
type PdfSaver struct {
	OutputDir string
}

// crée un PdfSaver (vérifie que le dossier output existe)
func New(outputDir string) (*PdfSaver, error) {
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("impossible de créer le dossier de sortie : %w", err)
	}

	return &PdfSaver{OutputDir: outputDir}, nil
}

// génère le PDF pour un pirate

func (p *PdfSaver) Save(pi pirate.Pirate) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4, // 595 x 842 points
	})
	pdf.AddPage()

	// Image de fond
	err := pdf.Image("assets/wantedVierge.jpg", 0, 0, &gopdf.Rect{W: 595, H: 842})
	if err != nil {
		return fmt.Errorf("impossible de charger le template : %w", err)
	}

	// Photo du pirate
	if pi.Img != "" {
		err = pdf.Image(pi.Img, 83, 200, &gopdf.Rect{W: 430, H: 325})
		if err != nil {
			fmt.Println("Image du pirate non trouvée :", pi.Img)
		}
	}

	//  Police pour le nom et la prime
	err = pdf.AddTTFFont("main", "assets/fonts/LiberationSans-Bold.ttf")
	if err != nil {
		return fmt.Errorf("police introuvable : %w", err)
	}

	// Nom du pirate
	pdf.SetFont("main", "", 39)
	pdf.SetTextColor(40, 25, 10)

	nameWidth, _ := pdf.MeasureTextWidth(pi.Name)
	pdf.SetX((595 - nameWidth) / 2) // Centré
	pdf.SetY(630)
	pdf.Cell(nil, pi.Name)

	// prime
	pdf.SetFont("main", "", 50)

	primeText := pi.Prime
	primeWidth, _ := pdf.MeasureTextWidth(primeText)
	pdf.SetX((620 - primeWidth) / 2)
	pdf.SetY(704)
	pdf.Cell(nil, primeText)

	//  Enregistrer le PDF
	outputPath := fmt.Sprintf("%s/%s.pdf", p.OutputDir, pi.Name)
	err = pdf.WritePdf(outputPath)
	if err != nil {
		return fmt.Errorf("erreur sauvegarde PDF : %w", err)
	}

	fmt.Printf("Avis de recherche créé : %s\n", outputPath)
	return nil
}
