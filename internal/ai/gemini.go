package ai

import (
    "bytes"
    "encoding/json"
    "net/http"
    "os"
)

type GeminiClient struct {
    APIKey string
}

func NewGeminiClient() *GeminiClient {
    apiKey := os.Getenv("GEMINI_API_KEY")
    if apiKey == "" {
        panic("GEMINI_API_KEY not set in .env")
    }
    return &GeminiClient{APIKey: apiKey}
}

func (c *GeminiClient) AnalyzeCommand(command string) (string, error) {
    url := "https://api.google.com/gemini/v1/analyze" // Thay bằng URL thật
    payload := map[string]string{"command": command}
    jsonData, _ := json.Marshal(payload)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", err
    }
    req.Header.Set("Authorization", "Bearer "+c.APIKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var result struct {
        Response string `json:"response"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }
    return result.Response, nil
}