# Insights Translation Tool

This tool automatically translates Blue website insights from English to all supported languages using OpenAI's GPT-4o-mini model.

## Features

- **Automatic Language Detection**: Scans `/content/` directory to detect available languages
- **Smart Queue Building**: Compares each language against English insights to find missing translations
- **Parallel Processing**: Runs 100 concurrent translations for 100x faster processing
- **Real-time Progress Tracking**: Shows progress percentage, average time, and running statistics
- **Token & Cost Tracking**: Thread-safe tracking of token counts and cost calculation
- **Frontmatter Validation**: Ensures YAML frontmatter fields are correctly preserved
- **Professional Translations**: Uses system prompt from translation guidelines to maintain quality
- **Error Handling**: Continues processing even if individual translations fail

## Prerequisites

- Go 1.20 or higher
- OpenAI API key set in `.env` file as `OPENAI_API_KEY`

## Usage

From the project root directory:

```bash
go run cmd/translate-insights/main.go
```

## Output Example

```
Found 16 languages: [es fr de it pt ja ko zh zh-TW ru nl pl sv id km]
Found 64 English insights

=== Translation Summary ===
Spanish (es): 52 translations needed
French (fr): 50 translations needed
German (de): 52 translations needed
...
Total translations required: 743

=== Starting Translation Process ===
Running 100 parallel workers...

[125/743] 16.8% | Avg: 1.2s | Elapsed: 2m 30s | Tokens: 185230 in, 197845 out | Cost: $0.3903

=== Translation Complete ===
Total tasks: 743
Completed: 743
Total time: 15m 23s
Average time: 1.2s/translation
Total tokens: 1102340 input, 1175892 output
Total cost: $2.3225
```

## Cost Calculation

- Input tokens: $0.40 per 1M tokens
- Output tokens: $1.60 per 1M tokens
- Model: gpt-4o-mini-2024-07-18

## Translation Rules

The tool follows strict translation guidelines:
- Never translates "Blue" (product name)
- Preserves email addresses, URLs, and technical terms
- Only translates `title` and `description` in frontmatter
- Maintains professional B2B tone
- Preserves all markdown formatting

## Error Handling

- Skips files that fail to translate and continues with the queue
- Validates frontmatter after translation
- Automatically creates target directories if missing