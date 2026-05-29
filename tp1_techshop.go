package main

import (
	"fmt"
	"strings"
)

// structure produit
type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

// structure catalogue
type Catalogue struct {
	Produits []Produit
}

// ajouter un produit
func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, produit := range c.Produits {
		if produit.ID == p.ID {
			return fmt.Errorf("id deja utilise")
		}
	}

	c.Produits = append(c.Produits, p)
	return nil
}

// trouver produit par id
func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, produit := range c.Produits {
		if produit.ID == id {
			return produit, nil
		}
	}

	return Produit{}, fmt.Errorf("produit introuvable")
}

// trouver produit par categorie
func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultat []Produit

	for _, produit := range c.Produits {
		if strings.EqualFold(produit.Categorie, cat) {
			resultat = append(resultat, produit)
		}
	}

	return resultat
}

// appliquer une reduction
func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	nb := 0

	for i := range c.Produits {
		if strings.EqualFold(c.Produits[i].Categorie, categorie) {
			c.Produits[i].Prix = c.Produits[i].Prix - (c.Produits[i].Prix * pct / 100)
			nb++
		}
	}

	return nb
}

// vendre un produit
func (c *Catalogue) Vendre(id int, qte int) error {
	for i := range c.Produits {
		if c.Produits[i].ID == id {
			if c.Produits[i].Stock < qte {
				return fmt.Errorf("stock insuffisant")
			}

			c.Produits[i].Stock = c.Produits[i].Stock - qte
			return nil
		}
	}

	return fmt.Errorf("produit introuvable")
}

// rapport du catalogue
func (c Catalogue) Rapport() string {
	nbProduits := len(c.Produits)
	valeurTotal := 0.0

	for _, produit := range c.Produits {
		valeurTotal = valeurTotal + (produit.Prix * float64(produit.Stock))
	}

	return fmt.Sprintf("Nombre de produits : %d\nValeur totale du stock : %.2f euros", nbProduits, valeurTotal)
}

func main() {
	catalogue := Catalogue{}

	// produits de depart
	catalogue.AjouterProduit(Produit{1, "iPhone 15", "Apple", 899.99, 10, "Smartphone", true})
	catalogue.AjouterProduit(Produit{2, "MacBook Air", "Apple", 1299.99, 5, "Ordinateur", true})
	catalogue.AjouterProduit(Produit{3, "Galaxy S24", "Samsung", 799.99, 8, "Smartphone", true})
	catalogue.AjouterProduit(Produit{4, "ThinkPad X1", "Lenovo", 1499.99, 4, "Ordinateur", true})
	catalogue.AjouterProduit(Produit{5, "AirPods Pro", "Apple", 249.99, 15, "Accessoire", true})

	for {
		var choix int

		fmt.Println()
		fmt.Println("=== TECHSHOP ===")
		fmt.Println("1. Ajouter")
		fmt.Println("2. Chercher")
		fmt.Println("3. Soldes")
		fmt.Println("4. Vendre")
		fmt.Println("5. Rapport")
		fmt.Println("0. Quitter")
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		if choix == 0 {
			break
		}

		switch choix {

		case 1:
			var p Produit

			fmt.Print("ID : ")
			fmt.Scan(&p.ID)

			fmt.Print("Nom : ")
			fmt.Scan(&p.Nom)

			fmt.Print("Marque : ")
			fmt.Scan(&p.Marque)

			fmt.Print("Prix : ")
			fmt.Scan(&p.Prix)

			fmt.Print("Stock : ")
			fmt.Scan(&p.Stock)

			fmt.Print("Categorie : ")
			fmt.Scan(&p.Categorie)

			p.Actif = true

			err := catalogue.AjouterProduit(p)

			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Produit ajoute")
			}

		case 2:
			var id int

			fmt.Print("ID du produit : ")
			fmt.Scan(&id)

			produit, err := catalogue.TrouverParID(id)

			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Produit trouve :")
				fmt.Println(produit)
			}

		case 3:
			var categorie string
			var pct float64

			fmt.Print("Categorie : ")
			fmt.Scan(&categorie)

			fmt.Print("Reduction en % : ")
			fmt.Scan(&pct)

			nb := catalogue.AppliquerReduction(categorie, pct)

			fmt.Println(nb, "produit(s) modifie(s)")

		case 4:
			var id int
			var qte int

			fmt.Print("ID du produit : ")
			fmt.Scan(&id)

			fmt.Print("Quantite : ")
			fmt.Scan(&qte)

			err := catalogue.Vendre(id, qte)

			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Vente effectuee")
			}

		case 5:
			fmt.Println(catalogue.Rapport())

		default:
			fmt.Println("Choix incorrect")
		}
	}
}