# CLI Games

A collection of terminal games built with Go, [Bubble Tea](https://github.com/charmbracelet/bubbletea), and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

The launcher currently includes Sudoku, Snake, and Tetris.

The arcade fills the terminal with a fixed navigation bar, centered play area, and controls along the bottom. Scores, game status, and game-specific controls align beside each board. Layouts update when the terminal is resized; windows that cannot fit the board show the required dimensions and pause play until you resize and press `p`.

## Features

### Sudoku

- Randomly generated Sudoku boards with a matching solution.
- Six difficulty levels: Easy, Medium, Hard, Expert, Master, and Extreme.
- In-game difficulty selection.
- Three-mistake limit with a Game Over screen.
- Score calculated from correct entries, mistakes, and elapsed time.
- Pause, restart, and victory flows.
- Colored cells for fixed values, editable values, selected cells, matching values, and mistakes.
- Individually bordered cells with row, column, and matching-value highlighting.
- Visible-grid candidate suggestions and a filled-cell counter.
- Given clues and paused boards are protected from edits.

### Snake

- Real-time movement with keyboard arrows or `W` `A` `S` `D`.
- Food, growing body, score, wall and self-collision detection.
- Progressive speed and level increases every five foods.
- Best score preserved across restarts during the current session.
- Pause and restart flows.
- One turn per movement tick prevents accidental reversals.
- Fill the whole board to win; directional head and level progress show your next goal.

### Tetris

- Seven tetrominoes with rotation, automatic falling, and instant drop.
- Balanced seven-piece bags and horizontal wall kicks.
- Landing ghost, graphical next-piece preview, and a score sidebar.
- Line clearing, score, and progressive levels.
- Pause and restart flows.

## Requirements

- Go 1.26 or newer.
- A terminal that supports ANSI colors.

## Run the game

```bash
go run .
```

When the launcher starts, choose Sudoku, Snake, or Tetris and press `Enter`. Sudoku then asks you to choose a difficulty level. Press `m` in any game to return to the launcher; games pause and retain their state. Reopen a paused game and press `p` to resume. Snake and Tetris fit an 80×24 terminal. Use at least 80×31 for Sudoku’s original bordered-cell board; 100×40 gives it more breathing room.

## Install as a CLI

Install the latest version directly with Go:

```bash
go install github.com/misafidiniaina/soduku-cli/cmd/cligames@latest
```

Go installs the `cligames` executable in `GOBIN` when it is set, or otherwise in `GOPATH/bin`. Make sure that directory is in your `PATH`, then start the game from anywhere:

```bash
cligames
```

To install a specific version or commit, replace `@latest` with the desired version, for example `@v1.0.0`.

### Install from a checkout

Build and install the binary in `~/.local/bin`:

```bash
./install.sh
```

The destination can be customized without `sudo`:

```bash
BINDIR="$HOME/bin" ./install.sh
```

## Development

Run the test suite and static checks with:

```bash
go test ./...
go vet ./...
```

The project is organized into:

- `main.go` — root compatibility entry point for `go run .`.
- `cmd/cligames` — installable CLI entry point that produces the `cligames` binary.
- `internal/app` — shared game launcher and menu navigation.
- `internal/sudoku` — Sudoku Bubble Tea model, input handling, state, and rendering ([documentation](internal/sudoku/README.md)).
- `internal/snake` — Snake Bubble Tea model, game rules, state, and rendering ([documentation](internal/snake/README.md)).
- `internal/tetris` — Tetris Bubble Tea model, game rules, state, and rendering ([documentation](internal/tetris/README.md)).
- `internal/sudoku/logic` — Sudoku scoring, timing, cursor behavior, and shared helpers.
- `internal/sudoku/logic/gen` — Sudoku solved-board and puzzle generation.
- `internal/ui` — terminal rendering and styles.

Each game exposes the standard Bubble Tea `Init`, `Update`, and `View` methods. To add a game, create a package under `internal`, implement that interface, and register its model and display name in `newAppModel` in `internal/app/app.go`.

## License

This project does not currently specify a license.
