package Users

import (
	"User_Management/Database"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

/* Permet de supprimer les users en double dans la liste */
func removeDuplicateValues(users []Database.User) []Database.User {
	keys := make(map[string]bool)
	list := []Database.User{}

	for _, entry := range users {
		if _, value := keys[entry.ID]; !value {
			keys[entry.ID] = true
			list = append(list, entry)
		}
	}
	return list
}

/* Permet de vérfier si un user existe déjà dans la BD, si c'est pas le cas, il l'ajoute.
 * Le choix de l'UpdateOne a été fait pour éviter l'ajout en double d'un user
 */
func addUserInDatabase(ch chan string, ctx context.Context, user Database.User) string {
	var user2 Database.User
	trouve := Database.GetUsersCollection("users").FindOne(ctx, bson.M{"id": user.ID}).Decode(&user2)
	if trouve == nil {
		return "User already exist"
	}
	options := options.Update().SetUpsert(true)
	_, err := Database.GetUsersCollection("users").UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": user}, options)
	if err != nil {
		ch <- err.Error()
	}
	return ""
}

/* Permet de :
 *   - Hasher le mot de passe
 *   - Insérer dans la BD l'utilisateur avec son mot de passe hasher si il n'existe pas
 *   - Créer un fichier local avec comme nom l'id du user et comme contenu le champ Data
 */
func hashAndInsertUser(ch chan string, ctx context.Context, user Database.User, wg *sync.WaitGroup) {
	defer wg.Done()
	var mu sync.Mutex

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		ch <- err.Error()
	}

	user.Password = string(passwordHash)

	/* Gère le cas où deux clients créent le même user */
	mu.Lock()
	res := addUserInDatabase(ch, ctx, user)
	mu.Unlock()
	if res == "User already exist" {
		ch <- res
		return
	} else {
		ch <- res
	}

	file, err := os.OpenFile(user.ID+".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()

	if err != nil {
		ch <- err.Error()
	}

	_, err = file.WriteString(user.Data)
	if err != nil {
		ch <- err.Error()
	}
}

/* Permet la création des users dans la Database */
func CreateUser(c *fiber.Ctx) error {
	var users []Database.User
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := json.Unmarshal(c.Body(), &users); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	users = removeDuplicateValues(users) // Gère le cas où un user n'existe pas dans la DB, mais il est en double dans la liste.
	ch := make(chan string, len(users))  // pour reçevoir les erreurs des goroutines

	for _, user := range users {
		wg.Add(1)
		go hashAndInsertUser(ch, ctx, user, &wg)

	}

	elem := <-ch
	if elem != "" && elem != "User already exist" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": elem})
	}
	wg.Wait()
	close(ch)
	return c.Status(http.StatusCreated).JSON(fiber.Map{"Message": "Created successfully"})
}
