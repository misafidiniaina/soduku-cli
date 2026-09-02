# Sudoku TUI

A colorful terminal Sudoku game built with Go, [Bubble Tea](https://github.com/charmbracelet/bubbletea), and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

## Features

- Randomly generated Sudoku boards with a matching solution.
- Six difficulty levels: Easy, Medium, Hard, Expert, Master, and Extreme.
- In-game difficulty selection.
- Three-mistake limit with a Game Over screen.
- Score calculated from correct entries, mistakes, and elapsed time.
- Pause, restart, and victory flows.
- Colored cells for fixed values, editable values, selected cells, matching values, and mistakes.
- Board layout adapts to the terminal window.

## Requirements

- Go 1.26 or newer.
- A terminal that supports ANSI colors.

## Run the game

```bash
go run .
```

When the game starts, choose a level with the arrow keys or `1`–`6`, then press `Enter`.

## Controls

| Key | Action |
| --- | --- |
| `↑` `↓` `←` `→` | Move the cursor |
| `1`–`9` | Enter a number |
| `Backspace` / `Delete` | Clear an editable cell |
| `p` | Pause or resume |
| `r` | Open restart options |
| `q` / `Ctrl+C` | Quit |
| `Enter` | Confirm a menu selection |

After pressing `r`, choose one of these options with the arrows and press `Enter`:

1. Replay the existing board.
2. Generate a new board at the current level.
3. Choose another level.

The game ends after the third mistake. A completed board displays the victory screen. Both screens allow restarting with `r`.

## Difficulty levels

Difficulty controls how many cells are removed from the generated solution:

| Level | Empty cells |
| --- | ---: |
| Easy | 30 |
| Medium | 40 |
| Hard | 45 |
| Expert | 50 |
| Master | 55 |
| Extreme | 60 |

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
- `main.go` — Bubble Tea application state and input handling.

## License

This project does not currently specify a license.
