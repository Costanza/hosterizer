# Go Coding Standards

## Code Style

- Follow official Go style guide and `gofmt` formatting
- Use `goimports` for automatic import management
- Run `go vet` and `golangci-lint` before committing
- Use tabs for indentation (Go standard)
- No line length limit, but keep lines readable

## Project Structure

```
project/
├── cmd/
│   └── appname/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   ├── repository/
│   └── service/
├── pkg/
│   └── utils/
├── api/
├── configs/
├── go.mod
├── go.sum
└── README.md
```

- `cmd/`: Application entry points
- `internal/`: Private application code (cannot be imported by other projects)
- `pkg/`: Public library code (can be imported)
- `api/`: API definitions (OpenAPI, protobuf)

## Naming Conventions

- Packages: short, lowercase, single word (no underscores)
- Interfaces: `Reader`, `Writer`, `Handler` (noun or agent noun)
- Variables: `camelCase` for local, `PascalCase` for exported
- Constants: `PascalCase` for exported, `camelCase` for private
- Acronyms: all caps (`HTTPServer`, `URLPath`, `userID`)

## Error Handling

- Always check errors, never ignore them
- Return errors as the last return value
- Use `errors.New()` or `fmt.Errorf()` for simple errors
- Wrap errors with context using `fmt.Errorf()` with `%w` verb
- Create custom error types for domain-specific errors

```go
func ProcessData(id string) (*Data, error) {
    data, err := fetchData(id)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch data for id %s: %w", id, err)
    }
    return data, nil
}

// Custom error type
type ValidationError struct {
    Field string
    Msg   string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Msg)
}
```

## Interfaces

- Keep interfaces small (1-3 methods ideal)
- Define interfaces where they're used, not where they're implemented
- Accept interfaces, return structs
- Use standard library interfaces when possible (`io.Reader`, `io.Writer`)

```go
// Good: small, focused interface
type UserStore interface {
    GetUser(id string) (*User, error)
    SaveUser(user *User) error
}

// Usage
func ProcessUser(store UserStore, id string) error {
    user, err := store.GetUser(id)
    // ...
}
```

## Concurrency

- Use channels for communication between goroutines
- Use `sync.WaitGroup` for waiting on multiple goroutines
- Use `context.Context` for cancellation and timeouts
- Protect shared state with `sync.Mutex` or `sync.RWMutex`
- Avoid goroutine leaks by ensuring all goroutines can exit

```go
func ProcessItems(ctx context.Context, items []Item) error {
    var wg sync.WaitGroup
    errCh := make(chan error, len(items))
    
    for _, item := range items {
        wg.Add(1)
        go func(i Item) {
            defer wg.Done()
            if err := processItem(ctx, i); err != nil {
                errCh <- err
            }
        }(item)
    }
    
    wg.Wait()
    close(errCh)
    
    for err := range errCh {
        if err != nil {
            return err
        }
    }
    return nil
}
```

## Testing

- Test files: `*_test.go` in the same package
- Use table-driven tests for multiple test cases
- Use `t.Helper()` for test helper functions
- Use subtests with `t.Run()` for organization
- Mock interfaces, not concrete types

```go
func TestCalculateTotal(t *testing.T) {
    tests := []struct {
        name     string
        prices   []float64
        taxRate  float64
        expected float64
    }{
        {"no tax", []float64{10.0, 20.0}, 0.0, 30.0},
        {"with tax", []float64{10.0, 20.0}, 0.1, 33.0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := CalculateTotal(tt.prices, tt.taxRate)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}
```

## Dependencies

- Use Go modules (`go.mod`, `go.sum`)
- Keep dependencies minimal
- Use semantic versioning
- Run `go mod tidy` regularly
- Vendor dependencies for production builds if needed

## Documentation

- Use godoc comments for all exported types, functions, and constants
- Start comments with the name of the thing being documented
- Use complete sentences
- Add examples in `*_example_test.go` files

```go
// User represents a system user with authentication credentials.
type User struct {
    ID       string
    Username string
    Email    string
}

// NewUser creates a new User with the given username and email.
// It returns an error if the email format is invalid.
func NewUser(username, email string) (*User, error) {
    // ...
}
```

## Best Practices

- Prefer explicit over implicit
- Use `defer` for cleanup operations
- Initialize structs with field names: `User{ID: "123", Name: "John"}`
- Use zero values effectively
- Avoid package-level state (globals)
- Use `context.Context` as the first parameter for functions that need it
- Return early to reduce nesting
- Use meaningful variable names (avoid single letters except for short scopes)
