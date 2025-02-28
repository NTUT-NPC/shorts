#!/bin/sh

# Install dependencies
go mod download

# Install templ and gow
go install github.com/mitranim/gow@latest

gow -c run cmd/shorts/main.go
