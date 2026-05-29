package main

import "fmt"

// structure personne
type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

// retourne prenom + nom
func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

// presentation de la personne
func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans, email : %s", p.NomComplet(), p.Age, p.Email)
}

// structure adresse
type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

// format de l'adresse
func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

// structure employe
type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

// fiche employe
func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"Nom : %s\nAge : %d\nEmail : %s\nAdresse : %s\nPoste : %s\nSalaire : %.2f €",
		e.NomComplet(),
		e.Age,
		e.Email,
		e.Format(),
		e.Poste,
		e.Salaire,
	)
}

// augmentation salaire
func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire + (e.Salaire * pct / 100)
}

// structure etudiant
type Etudiant struct {
	Personne
	Promo    string
	Moyenne  float64
}

// mention selon la moyenne
func (e Etudiant) MentionObtenue() string {

	switch {
	case e.Moyenne >= 16:
		return "TB"
	case e.Moyenne >= 14:
		return "B"
	case e.Moyenne >= 12:
		return "AB"
	default:
		return "P"
	}
}

func main() {

	// employe 1
	emp1 := Employe{
		Personne: Personne{
			Prenom: "Khadim",
			Nom:    "Cisse",
			Age:    26,
			Email:  "khadim@test.com",
		},
		Adresse: Adresse{
			Rue:        "10 rue de Lyon",
			Ville:      "Lyon",
			CodePostal: "69000",
		},
		Poste:   "Developpeur",
		Salaire: 2500,
	}

	emp2 := Employe{
		Personne: Personne{
			Prenom: "Khadim2",
			Nom:    "Cisse2",
			Age:    30,
			Email:  "khadim@mail.com",
		},
		Adresse: Adresse{
			Rue:        "20 rue Victor Hugo",
			Ville:      "Paris",
			CodePostal: "75000",
		},
		Poste:   "Chef de projet",
		Salaire: 3500,
	}
	//augmentation salaire de
	emp1.AugmenterSalaire(10)

	// etudiant 1
	etu1 := Etudiant{
		Personne: Personne{
			Prenom: "khadim",
			Nom:    "etudaint",
			Age:    20,
			Email:  "khadimetud@mail.com",
		},
		Promo:   "M1",
		Moyenne: 17,
	}

	// etudiant 2
	etu2 := Etudiant{
		Personne: Personne{
			Prenom: "khadim",
			Nom:    "etudiant2",
			Age:    21,
			Email:  "khadimet@mail.com",
		},
		Promo:   "M1",
		Moyenne: 13,
	}

	// affichage employes
	fmt.Println("=== EMPLOYES ===")
	fmt.Println(emp1.FicheEmploye())
	fmt.Println()
	fmt.Println(emp2.FicheEmploye())

	// affichage etudiants
	fmt.Println("\n=== ETUDIANTS ===")
	fmt.Println(etu1.NomComplet(), "- Mention :", etu1.MentionObtenue())
	fmt.Println(etu2.NomComplet(), "- Mention :", etu2.MentionObtenue())
}