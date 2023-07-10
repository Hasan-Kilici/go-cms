package Utils

import(
	"time"
	"math/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type File struct {
	Path    string
	Name    string
	Content string
}

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

func GetFiles(folderPath string) ([]File, error) {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	var fileStructs []File
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(folderPath, file.Name())
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Dosya okunamadı: %v\n", err)
			continue
		}

		fileStruct := File{
			Path:    filePath,
			Name:    file.Name(),
			Content: string(content),
		}
		fileStructs = append(fileStructs, fileStruct)
	}

	return fileStructs, nil
}