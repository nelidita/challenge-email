package handlers

import (
	"bytes"
	"emails-indexer/models"
	"emails-indexer/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func CkeckIndexExists() *http.Response {
	url := fmt.Sprintf("%sapi/%s/_mapping", os.Getenv("ZS_BASE_API_URL"), os.Getenv("ZS_INDEX"))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Error new request _mapping searching for the existence of the index: %v", err)
		return nil
	}

	utils.SetBasicHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error in searching for the existence of the index: %v", err)
		return nil
	}

	return res
}

func CreateIndex() error {
	url := fmt.Sprintf("%sapi/index", os.Getenv("ZS_BASE_API_URL"))
	indexName := os.Getenv("ZS_INDEX")

	body := models.Index{
		Name:         indexName,
		Storage_type: "disk",
		Shard_num:    6,
		Mappings: models.Mapping{
			Properties: models.Properties{
				Id: models.Property{
					TypeProperty: "text",
					Index:        true,
					Store:        true,
					Sortable:     false,
				},
				From: models.Property{
					TypeProperty: "text",
					Index:        true,
					Store:        true,
					Sortable:     true,
				},
				To: models.Property{
					TypeProperty: "text",
					Index:        true,
					Store:        true,
					Sortable:     true,
				},
				Content: models.Property{
					TypeProperty: "text",
					Index:        true,
					Store:        true,
					Sortable:     true,
				},
				Subject: models.Property{
					TypeProperty: "text",
					Index:        true,
					Store:        true,
					Sortable:     true,
				},
				Date: models.Property{
					TypeProperty: "date",
					Index:        true,
					Store:        true,
					Sortable:     true,
				},
				Filepath: models.Property{
					TypeProperty: "text",
					Index:        true,
					Store:        true,
					Sortable:     false,
				},
			},
		},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Panicf("Error convert data to JSON slice bytes in createIndex: %v", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf(
			"Error to create index NewRequest in (%s): %v", indexName, err)
		return err
	}

	utils.SetBasicHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error to create index Do: %v", err)
		return err
	}

	defer res.Body.Close()

	var resCreateIndex models.ResCreateIndex
	err = json.NewDecoder(res.Body).Decode(&resCreateIndex)
	if err != nil {
		log.Fatalf("Error to decode response JSON createIndex: %v", err)
		return err
	}

	fmt.Println("Index:", resCreateIndex.Index)
	fmt.Println("Message create Index:", resCreateIndex.Message)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error creating index: %v", res.Status)
	}

	return nil
}

func CreateDocuments(emails []models.Email) error {
	url := fmt.Sprintf("%sapi/_bulkv2", os.Getenv("ZS_BASE_API_URL"))
	indexName := os.Getenv("ZS_INDEX")
	fmt.Printf("Bulk indexName: %s\n", indexName)

	body := models.CreateIndexRequest{
		Index:   indexName,
		Records: emails,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Error convert data to JSON slice bytes: %v", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf(
			"Error to create documents NewRequest (%s): %v", indexName, err)
		return err
	}

	utils.SetBasicHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error to create documents Do: %v", err)
		return err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		return err
	}
	fmt.Printf("Zinc server response body: %s\n", string(resBody))

	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("error creating documents statusCode: %v", res.Status)
		return err
	}

	return nil
}

func DeleteIndex(index string) error {
	url := fmt.Sprintf("%sapi/index/%s", os.Getenv("ZS_BASE_API_URL"), index)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Println("The HTTP request DELETE index failed with error ", err)
		return err
	}

	utils.SetBasicHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("The HTTP request DELETE index returned an error: ", err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error deleting index: %v", res.Status)
	}

	log.Println("index deleted successfully")
	return nil
}
