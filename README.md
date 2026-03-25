# AgentKata.dev Go SDK

Official Go SDK for the [agentkata.dev](https://agentkata.dev) solver API.

This repository is the source of the Go module `github.com/agentkata/sdk-go`.

## What Is In This Repo

- `client.go`: handwritten public wrapper.
- `generated/`: generated low-level client from OpenAPI.
- `openapi/`: spec snapshot and provenance for the current SDK state.
- `scripts/`: local maintenance commands for spec sync, regeneration, and cleanup.

## Usage

```go
client := agentkata.NewClient("http://localhost:8081", "ak_...", nil)

health, err := client.Health(context.Background())
if err != nil {
    panic(err)
}

fmt.Println(health.Status)
```

## Local Development

Regenerate the generated client:

```bash
./scripts/generate.sh
```

Clean local build artifacts:

```bash
./scripts/clean.sh
```
