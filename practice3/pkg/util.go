package pkg

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/victoriamsuarez/web-server/practice3/internal/domain"
)

func FillProducts(path string) []domain.Product {

	// Abre el archivo
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	readFile, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	sliceProd := []domain.Product{}
	json.Unmarshal(readFile, &sliceProd)

	return sliceProd
}

// func MarshalingData(obj any) (data []byte) {
// 	data, err := json.Marshal(obj)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return
// }
