package main

import (
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

func main() {
	fmt.Println("Hii")
	GetESClient()
}
func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}
