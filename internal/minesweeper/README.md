# Minesweeper

Select Minesweeper from the launcher, choose a difficulty, and press `Enter`. All difficulties fit an 80×24 terminal.

## Difficulties

| Level | Board | Mines |
| --- | ---: | ---: |
| Beginner | 9×9 | 10 |
| Intermediate | 16×16 | 40 |
| Expert | 30×16 | 99 |

## Controls

| Key | Action |
| --- | --- |
| `↑` `↓` `←` `→` or `W` `A` `S` `D` | Move the cursor |
| `Enter` / `Space` | Reveal a cell |
| `f` | Flag or unflag a covered cell |
| `p` | Pause or resume |
| `r` | Start a new board at the same difficulty |
| `m` | Pause and return to the launcher |
| `q` / `Ctrl+C` | Quit |

## Rules

The first cell you reveal is always safe. A number shows how many mines touch that cell; revealing an empty area opens its neighboring safe cells automatically. Flags are limited to the board's mine count.

Reveal every safe cell to win. Revealing a mine ends the round and shows the mine locations. The timer starts with your first reveal.

Returning to the launcher with `m` pauses play. Choose Minesweeper again and press `p` to resume.
