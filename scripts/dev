#!/bin/sh

# Install dependencies
go mod download

# Install templ and gow
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/mitranim/gow@latest

templ generate --watch &
gow -c run cmd/shorts/main.go
