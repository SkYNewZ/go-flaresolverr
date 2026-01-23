# go-flaresolverr

Go client for https://github.com/FlareSolverr/FlareSolverr

## Features

- HTTP GET and POST requests through FlareSolverr proxy
- Session management (create, list, destroy) for persistent browser instances
- Cookie support for authenticated requests
- Proxy support for routing requests
- Context-aware with configurable timeouts
- Type-safe interface with proper error handling
- Full test coverage

## Installation

```bash
go get github.com/SkYNewZ/go-flaresolverr
```

## Usage

### Basic GET Request

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/SkYNewZ/go-flaresolverr"
    "github.com/google/uuid"
)

func main() {
    // Create a new FlareSolverr client
    client := flaresolverr.New("http://127.0.0.1:8191/v1", 60*time.Second, nil)

    // Make a simple GET request
    resp, err := client.Get(
        context.Background(),
        "https://example.com",
        uuid.Nil, // no session
        nil,      // no cookies
    )
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Status:", resp.Status)
    fmt.Println("Response:", resp.Solution.Response)
}
```

### Using Sessions

Sessions allow you to maintain browser state across multiple requests without re-solving challenges.

```go
// Create a new session
sessionID := uuid.New()
_, err := client.CreateSession(context.Background(), sessionID)
if err != nil {
    log.Fatal(err)
}
defer client.DestroySession(context.Background(), sessionID)

// Use the session for multiple requests
resp1, err := client.Get(context.Background(), "https://example.com", sessionID, nil)
if err != nil {
    log.Fatal(err)
}

resp2, err := client.Get(context.Background(), "https://example.com/page2", sessionID, nil)
if err != nil {
    log.Fatal(err)
}

// List all active sessions
sessions, err := client.ListSessions(context.Background())
if err != nil {
    log.Fatal(err)
}
fmt.Println("Active sessions:", sessions.Sessions)
```

### Using Cookies

Send custom cookies with your requests:

```go
cookies := []*flaresolverr.Cookie{
    {Name: "session_id", Value: "abc123"},
    {Name: "auth_token", Value: "xyz789"},
}

resp, err := client.Get(
    context.Background(),
    "https://example.com/protected",
    uuid.Nil, // no session
    cookies,  // with cookies
)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Response with cookies:", resp.Status)
```

### POST Requests

```go
// POST request with form data
postData := "username=john&password=secret"

resp, err := client.Post(
    context.Background(),
    "https://example.com/login",
    uuid.Nil,  // no session
    postData,  // application/x-www-form-urlencoded
    nil,       // no cookies
)
if err != nil {
    log.Fatal(err)
}

fmt.Println("POST response:", resp.Status)
```

### Using a Proxy

```go
// GET request through a proxy
resp, err := client.Get(
    context.Background(),
    "https://example.com",
    uuid.Nil,                          // no session
    nil,                               // no cookies
    "http://proxy.example.com:8080",   // proxy URL
)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Response via proxy:", resp.Status)
```

### Complete Example with Session, Cookies, and Proxy

```go
// Create session
sessionID := uuid.New()
_, err := client.CreateSession(context.Background(), sessionID)
if err != nil {
    log.Fatal(err)
}
defer client.DestroySession(context.Background(), sessionID)

// Prepare cookies
cookies := []*flaresolverr.Cookie{
    {Name: "preferences", Value: "dark_mode"},
}

// POST with session, cookies, and proxy
resp, err := client.Post(
    context.Background(),
    "https://example.com/api",
    sessionID,                         // with session
    "action=submit&data=value",        // POST data
    cookies,                           // with cookies
    "http://proxy.example.com:8080",   // with proxy
)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Complete example:", resp.Status)
fmt.Println("User Agent:", resp.Solution.UserAgent)
fmt.Println("Response cookies:", resp.Solution.Cookies)
```

## Error Handling

The library provides typed errors for common scenarios:

```go
resp, err := client.Get(context.Background(), "https://example.com", uuid.Nil, nil)
if err != nil {
    if errors.Is(err, flaresolverr.ErrRequestTimeout) {
        log.Println("Request timed out")
    } else if errors.Is(err, flaresolverr.ErrUnexpectedError) {
        log.Println("Unexpected error from FlareSolverr")
    } else {
        log.Fatal(err)
    }
}
```

## License

GPL-3.0
