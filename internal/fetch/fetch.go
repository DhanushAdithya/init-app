package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const baseURL = "http://localhost:11434/api/generate"

const requestBody = `{"model": "phi", "prompt": "%s", "stream": false}`

var client = &http.Client{}

type Response struct {
	Model              string `json:"model"`
	CreatedAt          string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	Context            []int  `json:"context"`
	TotalDuration      int    `json:"total_duration"`
	LoadDuration       int    `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int    `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int    `json:"eval_duration"`
}

func Fetch(prompt string) (Response, error) {
	query := strings.NewReader(fmt.Sprintf(requestBody, prompt))
	req, err := http.NewRequest(http.MethodPost, baseURL, query)
	if err != nil {
		return Response{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}
	var response Response
	if err := json.Unmarshal(resp, &response); err != nil {
		return Response{}, err
	}

	return response, nil
}
