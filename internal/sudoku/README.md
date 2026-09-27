# Sudoku

Sudoku is selected from the root launcher and starts with a difficulty menu.

## Controls

| Key                    | Action                   |
| ---------------------- | ------------------------ |
| `↑` `↓` `←` `→`        | Move the cursor          |
| `1`-`9`                | Enter a number           |
| `Backspace` / `Delete` | Clear an editable cell   |
| `p`                    | Pause or resume          |
| `r`                    | Open restart options     |
| `m`                    | Return to the game menu  |
| `q` / `Ctrl+C`         | Quit                     |
| `Enter`                | Confirm a menu selection |

After pressing `r`, choose one of these options with the arrows and press `Enter`:

1. Replay the existing board.
2. Generate a new board at the current level.
3. Choose another level.

The game ends after the third mistake. A completed board displays the victory screen.

## Difficulty

| Level   | Empty cells |
| ------- | ----------: |
| Easy    |          30 |
| Medium  |          40 |
| Hard    |          45 |
| Expert  |          50 |
| Master  |          55 |
| Extreme |          60 |

## Scoring

- `+100` for each correctly filled editable cell.
- `-50` for each mistake.
- `-1` for each elapsed second.
- The score never drops below zero.

The compact grid highlights the selected row, column, box, and matching values. Empty or incorrect editable cells show candidate numbers computed from the visible board. The timer stops during pause, menus, and completed games. Press `Esc` to cancel the restart menu.

Returning to the launcher with `m` pauses play. Choose the game again and press `p` to resume.
