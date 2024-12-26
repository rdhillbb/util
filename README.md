# Query Rewrite System

This system provides functionality to enhance and rewrite search queries using the Anthropic Claude API. It generates multiple variations of an input query while maintaining the original intent and adding contextual depth.

## Overview

The Query Rewrite System consists of two main components:
- A message template file (`messagefile.xml.txt`) containing the prompt structure
- A Go implementation (`util.go`) that handles the query rewriting logic

## Message File Structure

The system uses an XML-based message file that contains templates for different operations. The primary template used is the `query_rewrite` template under `utilmessages`. This template provides instructions to Claude for generating query variations.

### Template Variables
The query rewrite template accepts two parameters:
- `%d`: Number of query rewrites to generate
- `%s`: Original query to be enhanced

## Environment Variables

The following environment variables must be configured:

| Variable | Description | Default |
|----------|-------------|---------|
| `ANTHROPIC_API_KEY` | API key for Anthropic Claude service | Required |
| `REWRITENUM` | Number of query rewrites to generate | 3 |

## Core Functions

### `ReWriteQR(query string) ([]string, error)`
Main function that processes a query and returns an array of rewritten queries.

### `buildPrompt(query string) string`
Constructs the prompt using the message template and query parameters.

### `makeAnthropicRequest(prompt, apiKey string) (*http.Response, error)`
Handles the API communication with Anthropic's Claude service.

### `processResponse(resp *http.Response) ([]string, error)`
Processes the API response and extracts the rewritten queries.

## Query Rewrite Guidelines

The system generates variations that include:
- Broader context queries
- More specific/detailed queries
- Alternative phrasings
- Related subtopics
- Different perspectives

Each rewritten query:
- Maintains the original intent
- Uses natural language
- Avoids redundancy
- Includes relevant context
- Varies in complexity
- Must not exceed 200 characters

## Example Usage

```go
func main() {
    results, err := ReWriteQR("what are the health benefits of garlic")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    for i, result := range results {
        fmt.Printf("%d: %s\n", i+1, result)
    }
}
```

## API Configuration

The system uses the Claude 3 Sonnet model (`claude-3-sonnet-20240229`) with the following settings:
- Max tokens: 1000
- API Version: 2023-06-01

## Dependencies

- github.com/rdhillbb/messagefile
- github.com/joho/godotenv

## Error Handling

The system includes comprehensive error handling for:
- Missing environment variables
- API communication issues
- Response parsing errors
- Invalid JSON formatting
- Empty results

## Installation

1. Clone the repository
2. Create a `.env` file with required environment variables
3. Install dependencies: `go get`
4. Build: `go build`

## Contributing

When contributing to this project, please ensure that:
- All new message templates are added to the XML file
- Environment variables are documented
- Error handling follows the established pattern
- Tests are included for new functionality
