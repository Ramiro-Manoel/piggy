# Piggy

Personal finance organizer built in Go. This is a personal project, used both as a real tool for managing my own finances and as a way to practice Go.

## What it does (planned)

- Fetches transaction data through Open Finance, using the Pluggy aggregator.
- Answers financial questions over WhatsApp.
- Automatically categorizes transactions using an LLM.

## Architecture

Clean Architecture. Domain and use cases have no external dependencies. Everything else (HTTP, CLI, WhatsApp, Postgres, Pluggy, AI) lives in separate adapters.

## Stack

- Go
- PostgreSQL
- Pluggy (Open Finance)

## Status

Early stage. Currently building the domain models and an in-memory repository, before adding any external infrastructure.

## Running tests

```bash
go test ./...
```
