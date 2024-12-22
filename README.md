# Query Rewrite Utility

A Go utility that leverages the Anthropic Claude API to enhance and expand search queries into multiple comprehensive variations.

## Purpose

The `ReWriteQR` (Query Rewrite) function takes a single search query or question and generates 15 alternative, more robust versions while preserving the original intent. This helps in:
- Expanding search coverage
- Capturing different aspects of the same query
- Improving search result relevancy
- Generating comprehensive variations of the original question

## Prerequisites

- Go 1.x or higher
- Anthropic API key
- Environment variables configured in `.env` file
- Required dependencies:
  - github.com/joho/godotenv
  - github.com/rdhillbb/messagefile

## Configuration

1. Create a `.env` file in your project root
2. Add your Anthropic API key:
```
ANTHROPIC_API_KEY=your_api_key_here
```

## Usage

```go
results, err := ReWriteQR("what are the health benefits of garlic")
if err != nil {
    log.Fatal(err)
}

// Results will contain an array of 15 enhanced queries
for _, query := range results {
    fmt.Println(query)
}
```

## Response Format

The utility expects responses from Claude to be wrapped in XML tags:
```xml
<results>
["enhanced query 1", "enhanced query 2", ...]
</results>
```

## Error Handling

The utility includes comprehensive error handling for:
- Missing API keys
- API response errors
- JSON parsing issues
- Empty or invalid responses

## Dependencies

- Uses Claude 3 Sonnet model for query enhancement
- Requires properly formatted XML message templates
- Processes responses using regex pattern matching

## Keywords

- Query Expansion
- Search Enhancement
- Natural Language Processing
- Query Rewriting
- Claude API
- Go Utility
- Search Optimization
- Question Reformulation
- Semantic Search
- Query Preprocessing
