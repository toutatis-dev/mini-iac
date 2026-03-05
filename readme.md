# mini-iac

A declarative infrastructure tool built from scratch in Go.
Parses a custom configuration language, diffs desired state
against current state, and applies changes idempotently.

## How It Works

Define resources in a manifest file. The tool lexes and parses
the configuration into an AST, reads current state from
state.json, computes a diff, and applies only the necessary
changes — creating, updating, or deleting resources to match
the declared state. After each run, state is written back to
disk.

## Config Syntax
```
resource "file" "readme.md" {
    content = "# My Project";
}
```

Resources are declared with a provider type, a name, and
properties. Removing a resource from the manifest deletes it
on the next apply.

## Supported Providers

- **file** — creates files with specified content

## Run

go build -o mini-iac .
./mini-iac manifest.iac


## Planned

- Tests
- Directory provider
- Additional file properties (permissions, owner)
- Plan command (preview changes without applying)
- Multiple providers