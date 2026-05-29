package main
import (
	"fmt"
	"time"
)

func main() {
	var naissance string

	fmt.Print("Entrez votre date de naissance (JJ/MM/AAAA) : ")
	fmt.Scanln(&naissance)

	dateNaissance, _ := time.Parse("02/01/2006", naissance)

	aujourdhui := time.Now()

	age := aujourdhui.Year() - dateNaissance.Year()

	// Vérifie si l'anniversaire est déjà passé cette année
	if aujourdhui.YearDay() < dateNaissance.YearDay() {
		age--
	}

	fmt.Println("Vous avez", age, "ans")
}