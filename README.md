# CLI Games

A collection of terminal games built with Go, [Bubble Tea](https://github.com/charmbracelet/bubbletea), and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

The launcher currently includes Sudoku and Snake.

## Features

### Sudoku

- Randomly generated Sudoku boards with a matching solution.
- Six difficulty levels: Easy, Medium, Hard, Expert, Master, and Extreme.
- In-game difficulty selection.
- Three-mistake limit with a Game Over screen.
- Score calculated from correct entries, mistakes, and elapsed time.
- Pause, restart, and victory flows.
- Colored cells for fixed values, editable values, selected cells, matching values, and mistakes.
- Board layout adapts to the terminal window.

### Snake

- Real-time movement with keyboard arrows or `W` `A` `S` `D`.
- Food, growing body, score, wall and self-collision detection.
- Pause and restart flows.

## Requirements

- Go 1.26 or newer.
- A terminal that supports ANSI colors.

## Run the game

```bash
go run .
```

When the launcher starts, choose Sudoku or Snake and press `Enter`. Sudoku then asks you to choose a difficulty level.

## Install as a CLI

Install the latest version directly with Go:

```bash
go install github.com/misafidiniaina/sudoku@latest
```

Go installs the `sudoku` executable in `GOBIN` when it is set, or otherwise in `GOPATH/bin`. Make sure that directory is in your `PATH`, then start the game from anywhere:

```bash
sudoku
```

To install a specific version or commit, replace `@latest` with the desired version, for example `@v1.0.0`.

## Controls

| Key                    | Action                   |
| ---------------------- | ------------------------ |
| `↑` `↓` `←` `→`        | Move the cursor          |
| `1`–`9`                | Enter a number           |
| `Backspace` / `Delete` | Clear an editable cell   |
| `p`                    | Pause or resume          |
| `r`                    | Open restart options     |
| `q` / `Ctrl+C`         | Quit                     |
| `Enter`                | Confirm a menu selection |

### Snake controls

| Key                                | Action                  |
| ---------------------------------- | ----------------------- |
| `↑` `↓` `←` `→` or `W` `A` `S` `D` | Move the snake          |
| `p` / `Space`                      | Pause or resume         |
| `r`                                | Restart                 |
| `m`                                | Return to the game menu |
| `q` / `Ctrl+C`                     | Quit                    |

After pressing `r`, choose one of these options with the arrows and press `Enter`:

1. Replay the existing board.
2. Generate a new board at the current level.
3. Choose another level.

The game ends after the third mistake. A completed board displays the victory screen. Both screens allow restarting with `r`.

## Difficulty levels

Difficulty controls how many cells are removed from the generated solution:

| Level   | Empty cells |
| ------- | ----------: |
| Easy    |          30 |
| Medium  |          40 |
| Hard    |          45 |
| Expert  |          50 |
| Master  |          55 |
| Extreme |          60 |

## Scoring

The score starts at zero and is recalculated during play:

- `+100` for each correctly filled editable cell.
- `-50` for each mistake.
- `-1` for each elapsed second.
- The score never drops below zero.

## Development

Run the test suite and static checks with:

```bash
go test ./...
go vet ./...
```

The project is organized into:

- `internal/logic` — game rules, scoring, timing, and cursor behavior.
- `internal/logic/gen` — solved-board and puzzle generation.
- `internal/ui` — terminal rendering and styles.
- `main.go` — game launcher, Bubble Tea application state, and input handling.
- `internal/snake` — Snake game rules, state, rendering, and tests.

## License

This project does not currently specify a license.
