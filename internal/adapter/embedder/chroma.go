package embedder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func storeEmbedding(chromaURL string, embedding []float64, id string) error {
	embeddingData := map[string]any{
		"id":        id,
		"embedding": embedding,
	}

	jsonData, err := json.Marshal(embeddingData)
	if err != nil {
		return err
	}

	resp, err := http.Post(chromaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to store embedding: %s", body)
	}

	return nil
}

func querySimilarEmbeddings(chromaURL string, queryEmbedding []float64) ([]map[string]any, error) {
	queryData := map[string]any{
		"embedding": queryEmbedding,
	}

	jsonData, err := json.Marshal(queryData)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(chromaURL+"/search", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
