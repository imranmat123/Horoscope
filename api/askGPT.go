package askGPT

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ChatGPTRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func AskGPT(userInput string) (string, error) {
	apiKey := os.Getenv("API_KEY")
	reqBody := ChatGPTRequest{
		Model: "gpt-3.5-turbo", // Correct model
		Messages: []Message{
			{Role: "user", Content: userInput},
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// CHECK if response code is not 200
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("OpenAI API error: %s", respBody)
	}

	var gptResp ChatGPTResponse
	err = json.Unmarshal(respBody, &gptResp)
	if err != nil {
		return "", err
	}

	if len(gptResp.Choices) > 0 {
		return gptResp.Choices[0].Message.Content, nil
	}

	return "No response from GPT", nil
}
