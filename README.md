# Simple Calculator (Go)

A tiny command-line calculator written in Go. This repository contains a minimal example of a calculator implemented across a small number of files.

## What’s in this project

- `calculator.go` — Core calculator functions (arithmetic operations and helpers).
- `main.go` — Program entry point / CLI that uses the calculator functions.
- `go.mod` — Go module file with project dependencies and module path.

The project is intentionally small and focused for learning purposes.

## How to run

From the project directory, run:

```powershell
go run .
```

This will execute the CLI in `main.go`. If the project exposes functions you can also build a binary with:

```powershell
go build -o simple-calculator
.
\simple-calculator.exe
```

## Example usage

Typical usage depends on how `main.go` accepts input as (first value, operation, second value):

- Calculation => 1
- Exit => 2

## Future enhancement: Persist history of calculations

Goal: add a feature to save past equations and their results to a file so users can view or export their calculation history.

Suggested design (low-risk, easy to implement):

- File location: store in the project directory as `history.jsonl` (JSON Lines) or `history.txt`. Another option is to put it in the user config directory (e.g., `%APPDATA%\\simple-calculator` on Windows).
- Format (JSON Lines example):

```
{"equation":"2 + 3","result":5,"timestamp":"2025-11-05T12:34:56Z"}
{"equation":"10 / 2","result":5,"timestamp":"2025-11-05T12:35:10Z"}
```

- Minimal API sketch (Go):

```go
type HistoryEntry struct {
    Equation  string    `json:"equation"`
    Result    float64   `json:"result"`
    Timestamp time.Time `json:"timestamp"`
}

func SaveHistory(path string, e HistoryEntry) error {
    // open file (create/append), marshal JSON, write newline
}

func ReadHistory(path string) ([]HistoryEntry, error) {
    // read lines, unmarshal each JSON line
}
```

Edge cases and considerations:

- Concurrency: use file locking or serialize writes if multiple goroutines can write history.
- Corruption: use JSON Lines to keep each entry independent so one malformed entry doesn't break the whole file.
- Rotation: consider limiting file size and rotating older history into a backup file.

## Next steps I can do for you

- Implement the history functions and wire them into `main.go` so the CLI automatically appends each calculation to the history file.
- Add a CLI flag like `--history` to print or export history.

If you'd like me to implement the history feature now, say so and I will add the necessary code and tests.

## License

This repository doesn't include a license file. If you want a license added (MIT, Apache-2.0, etc.), tell me which one and I'll add it.
