# Piggy

Personal finance organizer built in Go. Personal project for both real use and Go practice.

## What it does (planned)

- Fetches transaction data through Open Finance (Pluggy).
- Automatically categorizes transactions using an LLM.
- Answers financial questions over WhatsApp.

## Stack

- Go
- PostgreSQL
- Pluggy (Open Finance)

## Status

Phase 1 in progress. Transaction context implemented with in-memory repository. No external infrastructure yet.

## Running tests

```bash
go test ./...
```
