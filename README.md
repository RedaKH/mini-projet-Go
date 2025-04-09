
```markdown
# 📝 Gestionnaire de Tâches en Go (CLI)

Une application en ligne de commande simple pour gérer des tâches.

---

## ⚙️ Commandes disponibles

### ➕ Ajouter une tâche
Ajoute une nouvelle tâche avec la date du jour.

```bash
go run main.go add "Nom de la tâche"
```

---

### 📋 Lister les tâches
Affiche toutes les tâches, leur état (faite ou non) et la date d’ajout.

```bash
go run main.go list
```

---

### ✅ Marquer une tâche comme faite
Coche une tâche à partir de son numéro dans la liste.

```bash
go run main.go done <numero>
```

Exemple :
```bash
go run main.go done 2
```

---

### ❌ Supprimer une tâche
Supprime une tâche à partir de son numéro dans la liste.

```bash
go run main.go delete <numero>
```

Exemple :
```bash
go run main.go delete 1
```

---

## 💾 Fichier de sauvegarde

Les tâches sont automatiquement sauvegardées dans le fichier :
```
taches.json
```

---

## 🧪 Exemple d'exécution
```bash
go run main.go add "Apprendre Go"
go run main.go list
go run main.go done 1
go run main.go delete 1
```



## ✅ Technologies
- Go (Golang)
- JSON (sérialisation)
- Terminal (ligne de commande)

