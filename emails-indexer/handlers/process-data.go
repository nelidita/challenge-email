package handlers

import (
	"emails-indexer/constants"
	"emails-indexer/models"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func EmailsDataDirectory() ([]models.Email, error) {
	var emails []models.Email
	var m sync.Mutex
	var wg sync.WaitGroup

	err := filepath.Walk(constants.EmailsRootFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("Error in walk func: ", err)
			return err
		}

		if !info.IsDir() {
			wg.Add(1)
			go func(p string) {
				defer wg.Done()
				emailData, err := processEmailFile(p)
				if err != nil {
					log.Fatal("Error in walk func: ", err)
					return
				}
				m.Lock()
				emails = append(emails, *emailData)
				m.Unlock()
			}(path)
		}

		return nil
	})

	if err != nil {
		log.Fatal("Error process emails files: ", err)
		return nil, err
	}

	wg.Wait()

	return emails, nil
}

func processEmailFile(path string) (*models.Email, error) {
	file, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		log.Println("Error read file: ", err)
		return nil, err
	}

	arr := strings.SplitN(string(file), "\r\n\r\n", 2)
	if len(arr) != 2 {
		log.Println("Error getting content from mail file: ", path)
		return nil, err
	}

	allDetails, content := arr[0], arr[1]
	details := strings.Split(allDetails, "\r\n")
	email := mapEmailDetails(details)
	email.Content = content
	email.Filepath = path

	return email, nil
}

func mapEmailDetails(details []string) *models.Email {
	email := &models.Email{}

	for i := 0; i < len(details); i++ {
		keyValue := strings.SplitN(details[i], ": ", 2)
		switch keyValue[0] {
		case "Message-ID":
			email.Id = keyValue[1]
		case "From":
			email.From = keyValue[1]
		case "To":
			email.To = keyValue[1]
		case "Subject":
			email.Subject = keyValue[1]
		case "Date":
			email.Date = keyValue[1]
		default:
			continue
		}
	}

	return email
}
