package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "time"
)


type Tache struct {
    Nom   string
    Faite bool
    Date  string
}

var fichier = "taches.json"

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage : go run main.go [add/list/done/delete]")
        return
    }

    commande := os.Args[1]
    taches := chargerTaches()

    switch commande {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Usage : go run main.go add \"Ma tâche\"")
            return
        }
        nom := os.Args[2]
        date := time.Now().Format("2006-01-02 15:04:05")
        taches = append(taches, Tache{Nom: nom, Faite: false, Date: date})
        sauvegarderTaches(taches)
        fmt.Println("Tâche ajoutée :", nom)

    case "list":
        if len(taches) == 0 {
            fmt.Println("Aucune tâche.")
            return
        }
        for i, t := range taches {
            etat := "[ ]"
            if t.Faite {
                etat = "[x]"
            }
            fmt.Printf("%d. %s %s (ajoutée le %s)\n", i+1, etat, t.Nom, t.Date)
        }

    case "done":
        if len(os.Args) < 3 {
            fmt.Println("Usage : go run main.go done <numero>")
            return
        }
        var index int
        fmt.Sscanf(os.Args[2], "%d", &index)
        if index < 1 || index > len(taches) {
            fmt.Println("Numéro invalide")
            return
        }
        taches[index-1].Faite = true
        sauvegarderTaches(taches)
        fmt.Println("Tâche marquée comme faite :", taches[index-1].Nom)

    case "delete":
        if len(os.Args) < 3 {
            fmt.Println("Usage : go run main.go delete <numero>")
            return
        }
        var index int
        fmt.Sscanf(os.Args[2], "%d", &index)
        if index < 1 || index > len(taches) {
            fmt.Println("Numéro invalide")
            return
        }
        supprimée := taches[index-1].Nom
        taches = append(taches[:index-1], taches[index:]...)
        sauvegarderTaches(taches)
        fmt.Println("Tâche supprimée :", supprimée)

    default:
        fmt.Println("Commande inconnue :", commande)
    }
}

func chargerTaches() []Tache {
    contenu, err := ioutil.ReadFile(fichier)
    if err != nil {
        return []Tache{}
    }

    var taches []Tache
    json.Unmarshal(contenu, &taches)
    return taches
}

func sauvegarderTaches(taches []Tache) {
    data, _ := json.MarshalIndent(taches, "", "  ")
    ioutil.WriteFile(fichier, data, 0644)
}
