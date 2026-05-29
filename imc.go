
package main

import "fmt"
//date 29/05/2026
func main() {

	// Variables
	var poids float64 = 70.5
	var taille float64 = 1.75

	// Constantes
	const IMCMaigreur = 18.5
	const IMCNormal = 25.0
	const IMCSurpoids = 30.0

	// Calcul 
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