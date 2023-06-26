package Utils
import(
	"time"
	"math/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)
func GenerateToken(length int) string {
	rand.Seed(time.Now().UnixNano())

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func CheckPassword(password, hashedPassword string) bool {
	decodedHash, err := hex.DecodeString(hashedPassword)
	if err != nil {
		fmt.Println("Hata:", err)
		return false
	}

	decodedText := string(decodedHash)
	return password == decodedText
}