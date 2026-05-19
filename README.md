# pak

Git decorator for AI-assisted development. Captures Claude Code sessions, tokens, and code attribution alongside your commits.

## Install

```bash
go install github.com/janedzumerko/pak/cmd/pak@latest
```

## Usage

Use pak exactly like git. All commands pass through transparently.

```bash
pak status
pak commit -m "fix bug"
pak push origin main
```

License

MIT
