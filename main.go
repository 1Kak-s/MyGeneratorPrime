package main

import (
	"flag"
	"fmt"
	"os"

	"MyGeneratorPrime/csvparser"
	"MyGeneratorPrime/generateprime"
	"MyGeneratorPrime/pdf"
)

func main() {
	// Déclaration des modes
	cliMode := flag.Bool("cli", false, "mode ligne de commande")
	fileMode := flag.Bool("file", false, "mode fichier CSV")

	// Arguments du mode -cli
	name := flag.String("name", "", "nom du pirate")
	prime := flag.String("prime", "", "prime du pirate")
	img := flag.String("img", "", "chemin de l'image")

	// Argument du mode -file
	path := flag.String("path", "", "chemin du fichier CSV")

	flag.Parse()

	// MODE -cli : un seul pirate
	if *cliMode {
		if *name == "" || *prime == "" || *img == "" {
			fmt.Println(" Erreur : il manque des arguments")
			fmt.Println("Usage : ./MyGeneratorPrime -cli -name \"LUFFY\" -prime \"1,500,000,000\" -img \"assets/luffy.png\"")
			os.Exit(1)
		}

		err := generateprime.Generate(*name, *prime, *img, "./output")
		if err != nil {
			fmt.Println(" Erreur :", err)
			os.Exit(1)
		}

		// MODE -file : plusieurs pirates via CSV
	} else if *fileMode {
		if *path == "" {
			fmt.Println(" Erreur : il manque le chemin du fichier CSV")
			fmt.Println("Usage : ./MyGeneratorPrime -file -path \"pirates.csv\"")
			os.Exit(1)
		}

		//  parse le CSV
		pirates, err := csvparser.Parse(*path)
		if err != nil {
			fmt.Println(" Erreur CSV :", err)
			os.Exit(1)
		}

		//  crée le PdfSaver une seule fois
		saver, err := pdf.New("./output")
		if err != nil {
			fmt.Println(" Erreur :", err)
			os.Exit(1)
		}

		//  génère un PDF pour chaque pirate
		for _, p := range pirates {
			err = saver.Save(p)
			if err != nil {
				fmt.Printf("  Erreur pour %s : %s\n", p.Name, err)
			}
		}

	} else {
		fmt.Println(" Erreur : choisis un mode")
		fmt.Println("  -cli  : ./MyGeneratorPrime -cli -name \"LUFFY\" -prime \"1,500,000,000\" -img \"assets/luffy.png\"")
		fmt.Println("  -file : ./MyGeneratorPrime -file -path \"pirates.csv\"")
		os.Exit(1)
	}
}
