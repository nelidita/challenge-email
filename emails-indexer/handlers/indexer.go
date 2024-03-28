package handlers

import (
	"log"
)

func ProcessDataIndexer() {
	log.Println("Process data directory")
	emails, err := EmailsDataDirectory()

	if err != nil {
		log.Fatalf("Failed process data: %s", err)
	}

	log.Println("Starting data indexing")
	errCreateIndex := CreateIndex()
	if errCreateIndex != nil {
		log.Fatalf("Failed to create index: %s", errCreateIndex)
		return
	}

	errIndexDocuments := CreateDocuments(emails)
	if errIndexDocuments != nil {
		log.Fatalf("Failed to create documents in index: %s", errIndexDocuments)
	}
}
