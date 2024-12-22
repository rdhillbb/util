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

## Message Template Format

The utility requires a properly structured XML file containing message templates. Here's the expected format:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<messages>
    <utilmessages>
        <query_rewrite>
            You are tasked with enhancing a user's query by creating multiple rewrites. 
            This process aims to generate 15 more comprehensive and effective search 
            queries while maintaining the original intent of the user's question or request. 
            You are only to provide the results. No additional information is to be added.
            Here is the user's original query:
            <user_query>
            %s
            </user_query>
            Place the results in a json string in the tag <results></results>
        </query_rewrite>
    </utilmessages>
    <supportmessages>
        <help>
            Please provide assistance for the following request:
            <request>
            %s
            </request>
            Format your response with priority and action items.
        </help>
    </supportmessages>
</messages>
```

The XML structure contains:
- `<messages>`: Root element
- `<utilmessages>`: Contains utility-related message templates
  - `<query_rewrite>`: Template for query rewriting instructions
    - Uses `%s` placeholder for the user's query
- `<supportmessages>`: Contains support-related message templates
  - `<help>`: Template for help requests

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
