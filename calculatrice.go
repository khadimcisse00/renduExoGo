package main

import "fmt"

// fonction calcul
func calcul(a float64, b float64, op string) (float64, error) {

	// creation de l'operation
	operation := creerOperation(op)

	// si operation inconnue
	if operation == nil {
		return 0, fmt.Errorf("operation inconnue")
	}

	// division par 0
	if op == "/" && b == 0 {
		return 0, fmt.Errorf("division par zero")
	}

	// retourne le resultat
	return operation(a, b), nil
}

// fonction qui cree une operation
func creerOperation(op string) func(float64, float64) float64 {

	switch op {

	case "+":
		return func(a float64, b float64) float64 {
			return a + b
		}

	case "-":
		return func(a float64, b float64) float64 {
			return a - b
		}

	case "*":
		return func(a float64, b float64) float64 {
			return a * b
		}

	case "/":
		return func(a float64, b float64) float64 {
			return a / b
		}

	default:
		return nil
	}
}

func main() {

	// boucle jusqu'a quit
	for {

		var a, b float64
		var op string

		// saisie utilisateur
		fmt.Print("Entrez deux nombres et une operation (+, -, *, /) ou quit : ")
		fmt.Scan(&a, &b, &op)

		// quitter le programme
		if op == "quit" {
			break
		}

		// appel fonction calcul
		resultat, err := calcul(a, b, op)

		// affichage erreur
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			// affichage resultat
			fmt.Println("Resultat :", resultat)
		}
	}
}