package models

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			ID        string                 `json:"_id"`
			Timestamp string                 `json:"@timestamp"`
			Source    map[string]interface{} `json:"_source"`
		} `json:"hits"`
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
	} `json:"hits"`
}

type BodyQuery struct {
	WordSearch string   `json:"word_search"`
	SortFields []string `json:"sort_fields"`
	Page       int      `json:"page"`
}

type SearchEmailsRequest struct {
	SearchType string                 `json:"search_type"`
	SortFields []string               `json:"sort_fields"`
	From       int                    `json:"from"`
	MaxResults int                    `json:"max_results"`
	Query      SearchEmailsQuery      `json:"query"`
	Source     map[string]interface{} `json:"_source"`
}

type SearchEmailsQuery struct {
	Term string `json:"term"`
}
