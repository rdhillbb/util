package util 
//package main 

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rdhillbb/messagefile"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// AnthropicResponse represents the structure of the response from Anthropic API
type AnthropicResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

func ReWriteQR(query string) ([]string, error) {
	var stringArray []string

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("ANTHROPIC_API_KEY environment variable not set")
	}

	prompt := buildPrompt(query)

	resp, err := makeAnthropicRequest(prompt, apiKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	stringArray, err = processResponse(resp)
	if err != nil {
		return nil, err
	}

	return stringArray, nil
}

func getRewriteNum() int {
    numStr := os.Getenv("REWRITENUM")
    if numStr == "" {
        return 3 // Default value
    }
    
    num, err := strconv.Atoi(numStr)
    if err != nil {
        log.Printf("Invalid REWRITENUM value: %s, using default of 3", numStr)
        return 3
    }
    if num > 8 {
	    num=8
     }
    return num
}

func buildPrompt(query string) string {
	// Retrieve a message
	msg, err := messagefile.GetMSG("utilmessages:query_rewrite")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ""
	}

	// Get number of rewrites
	numRewrites := getRewriteNum()
	
	// Now using two format specifiers - one for number of rewrites, one for query
	return fmt.Sprintf(msg, numRewrites, query)
}

func makeAnthropicRequest(prompt, apiKey string) (*http.Response, error) {
	payload := map[string]interface{}{
		"model":      "claude-3-sonnet-20240229",
		"max_tokens": 1000,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling payload: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	return client.Do(req)
}

func processResponse(resp *http.Response) ([]string, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status: %d, body: %s", resp.StatusCode, string(body))
	}

	var anthropicResp AnthropicResponse
	if err := json.Unmarshal(body, &anthropicResp); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	if len(anthropicResp.Content) == 0 {
		return nil, fmt.Errorf("no content in response")
	}

	re := regexp.MustCompile(`(?s)<results>\s*(.*?)\s*</results>`)
	matches := re.FindStringSubmatch(anthropicResp.Content[0].Text)
	if len(matches) < 2 {
		return nil, fmt.Errorf("could not find results in content")
	}

	jsonStr := matches[1]
	jsonStr = strings.TrimSpace(jsonStr)

	var stringArray []string
	err = json.Unmarshal([]byte(jsonStr), &stringArray)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v\n raw JSON: %s", err, jsonStr)
	}

	if len(stringArray) == 0 {
		return nil, fmt.Errorf("parsed JSON resulted in empty array")
	}

	return stringArray, nil
}

func testmain() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	results, err := ReWriteQR("I am doing resesarch on Jefferson's view of democracy. for this I need to examine all aspects of his life, business, family and poitical career.")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Final Results:")
	for i, result := range results {
		fmt.Printf("%d: %s\n", i+1, result)
	}
}
