package controllers

import (
	"bytes"
	"email-challenge/constants"
	"email-challenge/models"
	utilsIndexer "emails-indexer/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (a *App) SearchEmails(w http.ResponseWriter, r *http.Request) {
	var query models.BodyQuery

	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		log.Println("Error query invalid: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	searchResult, err := searchZS(query)
	if err != nil {
		log.Println("Error search in zincSearch: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(searchResult)
}

func searchZS(query models.BodyQuery) (*models.SearchResponse, error) {
	url := fmt.Sprintf("%sapi/%s/_search", os.Getenv("ZS_BASE_API_URL"), os.Getenv("ZS_INDEX"))

	var searchType string
	if query.WordSearch == "" {
		searchType = constants.AllEmails
	} else {
		searchType = constants.EmailSearchType
	}

	from := (query.Page - 1) * constants.EmailMaxResults

	body := models.SearchEmailsRequest{
		SearchType: searchType,
		Query: models.SearchEmailsQuery{
			Term: query.WordSearch,
		},
		SortFields: query.SortFields,
		From:       from,
		MaxResults: constants.EmailMaxResults,
	}

	jsonQuery, err := json.Marshal(body)
	if err != nil {
		log.Println("Error process body: ", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonQuery))
	if err != nil {
		log.Println("error in method POST /_search: ", err)
		return nil, err
	}

	utilsIndexer.SetBasicHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending api/index/_search HTTP request: ", err)
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("bad status code from server sincSearch: %d", res.StatusCode)
		return nil, err
	}

	bodyRes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error read response body: %s", err)
	}

	var response *models.SearchResponse
	json.Unmarshal(bodyRes, &response)

	return response, nil
}
