# swini-cli

A Go CLI tool for account creation and management.

## Features
- `swini account create`: Generates an EC key pair, stores the private key locally, calls a GraphQL mutation to create an account, and stores the returned user ID.

## Usage
```
swini account create
```


## Useful Commands

```bash
# Generate grapqhl files
task build:gql

# Test the CLI
task run -- account create
go main.go account create
```