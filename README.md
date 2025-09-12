# swini-cli

A Go CLI tool for profile creation and management.

## Features
- `swini profile create`: Generates an EC key pair, stores the private key locally, calls a GraphQL mutation to create a profile, and stores the returned user ID.

## Usage
```
swini profile create
```


## Useful Commands

```bash
# Generate grapqhl files
task build:gql

# Test the CLI
task run -- profile create
go main.go profile create
```