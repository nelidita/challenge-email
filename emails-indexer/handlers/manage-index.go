package handlers

import (
	"bytes"
	"emails-indexer/models"
	"emails-indexer/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func CkeckIndexExists() *http.Response {
	url := fmt.Sprintf("%sapi/%s/_mapping", os.Getenv("ZS_BASE_API_URL"), os.Getenv("ZS_INDEX"))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Error in searching for the existence of the index: %v", err)
		return nil
	}

	utils.SetBasicHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("error sending api/_mapping HTTP request: %v", err)
		return nil
	}

	return res
}

func CreateIndex(emails []models.Email) error {
	url := fmt.Sprintf("%sapi/_bulkv2", os.Getenv("ZS_BASE_API_URL"))
	indexName := os.Getenv("ZS_INDEX")

	body := models.CreateIndexRequest{
		Index:   indexName,
		Records: emails,
	}

	// Returns a chunk of bytes containing the JSON data
	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Panicf("Error to marshall data for create index request: %v", err)
		return err
	}

	// POST request to create index with your documents
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		log.Fatalf(
			"Error to create new index (%s): %v", os.Getenv("ZS_INDEX"), err)
		return err
	}

	utils.SetBasicHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error to create index: %v", err)
		return err
	}

	defer res.Body.Close()

	// Get the json of the response to display in the terminal
	var responseMap map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&responseMap)
	if err != nil {
		log.Fatalf("Error to decode response JSON: %v", err)
		return err
	}

	message, _ := responseMap["message"].(string)
	recordCount, _ := responseMap["record_count"].(float64)
	fmt.Println("Message:", message)
	fmt.Println("Record count:", recordCount)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error creating index: %v", res.Status)
	}

	return nil
}

func DeleteIndex(index string) error {
	url := fmt.Sprintf("%sapi/index/%s", os.Getenv("ZS_BASE_API_URL"), index)

	req, err := http.NewRequest("DELETE", url, nil)
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
