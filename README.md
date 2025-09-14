# Swini CLI

A Go CLI tool for interacting with Swini.

## Installation

```bash
curl -sSL https://nosebit.github.io/swini-cli/install.sh | bash
```

## Usage

```bash
# Create a new account
swini account create

# Setup a pay method (credit card)
swini pay method setup

# Setup a pay account (to receive funds)
swini pay account setup
```

## Useful Commands

```bash
# Generate grapqhl files
task build:gql

# Test the CLI
task run -- account create
# or: go main.go account create
```
