# Blackjack Deck CLI (Go)

[![CI](https://img.shields.io/github/actions/workflow/status/mahmoudnasser1561/blackjack/ci.yml?branch=main&label=CI)](https://github.com/mahmoudnasser1561/blackjack/actions/workflows/ci.yml?query=branch%3Amain)
[![License](https://img.shields.io/github/license/mahmoudnasser1561/blackjack)](./LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/mahmoudnasser1561/blackjack?label=go)](./go.mod)

A focused Go learning project that implements deck-of-cards utilities behind a small CLI.
It is structured with `cmd/` and `internal/`, includes unit tests, and ships with CI/release workflows.

## Features

- Create a deck and save it to a file.
- Load a deck from disk.
- Shuffle an existing deck file.
- Deal a hand from a deck file and print remaining deck size.
- Validate invalid input with clear CLI errors and non-zero exit codes.
- Unit tests for core deck behavior.
- GitHub Actions CI and tag-based release packaging.

## Quickstart

```bash
go test ./...
go run ./cmd/deck --help
go run ./cmd/deck new --out assets/my_cards.txt
go run ./cmd/deck shuffle --in assets/my_cards.txt
go run ./cmd/deck deal --in assets/my_cards.txt --hand 5
```

Example output (card order is not deterministic after shuffle):

```text
$ go run ./cmd/deck deal --in assets/my_cards.txt --hand 5
Hand (5 cards):
0 Ace of Spades
1 Two of Diamonds
2 Four of Clubs
3 Ace of Hearts
4 Three of Spades
Remaining deck size: 11
```

## Usage

```text
deck new [--out <path>]
deck shuffle --in <path> [--out <path>]
deck deal --in <path> [--hand <n>]
```

| Subcommand | Description | Flags |
|---|---|---|
| `new` | Create a new deck and save it | `--out` (default: `assets/my_cards.txt`) |
| `shuffle` | Load from file, shuffle, write result | `--in` (required), `--out` (default: same as `--in`) |
| `deal` | Deal hand from file and print result | `--in` (required), `--hand` (default: `5`) |

Help examples:

```bash
go run ./cmd/deck --help
go run ./cmd/deck new --help
go run ./cmd/deck shuffle --help
go run ./cmd/deck deal --help
```

## Project Structure

```text
.
├── assets
│   └── my_cards.txt
├── cmd
│   └── deck
│       └── main.go
├── internal
│   └── deck
│       ├── deck.go
│       └── deck_test.go
├── go.mod
├── LICENSE
└── README.md
```

Why this layout:

- `cmd/deck`: CLI entrypoint (`main` package), keeping executable wiring separate from domain logic.
- `internal/deck`: Core deck logic and tests, scoped for internal module use.
- `assets`: Sample/default deck file path used by CLI commands.

## Development

Local quality commands (Go standard tools):

```bash
go fmt ./...
go vet ./...
go test ./...
go test ./... -race
```

CI runs on pushes and pull requests and must pass before merge.

Convenience commands via Makefile:

```bash
make test
make ci
make build
make run
```

## Testing

Unit tests in `internal/deck/deck_test.go` cover:

- deck creation and expected size/card ordering assumptions
- save/load behavior for deck files
- hand-size validation for deal operations

Run all tests:

```bash
go test ./...
```

## Releasing

Push a semantic version tag to trigger `.github/workflows/release.yml`:

```bash
git tag vX.Y.Z
git push origin vX.Y.Z
```

Release artifacts are attached to the repository’s GitHub Releases page.

## License

Licensed under the terms in [`LICENSE`](./LICENSE).
