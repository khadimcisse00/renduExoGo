package main

import "fmt"

func main() {

	// Variables
	var poids float64 = 70.5
	var taille float64 = 1.75

	// Constantes
	const IMCMaigreur = 18.5
	const IMCNormal = 25.0
	const IMCSurpoids = 30.0

	// Calcul IMC
	imc := poids / (taille * taille)

	// Affichage IMC
	fmt.Printf("IMC : %.2f\n", imc)

	// Catégorie
	if imc < IMCMaigreur {
		fmt.Println("Categorie : Maigreur")
	} else if imc < IMCNormal {
		fmt.Println("Categorie : Normal")
	} else if imc < IMCSurpoids {
		fmt.Println("Categorie : Surpoids")
	} else {
		fmt.Println("Categorie : Obesite")
	}

	// Bonus
	fmt.Println("Khadim Cisse")
}