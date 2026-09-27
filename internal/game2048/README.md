# 2048

Select **2048** from the launcher. The game starts with two tiles on a 4×4 board and fits an 80×24 terminal.

## Controls

| Key | Action |
| --- | --- |
| `↑` `↓` `←` `→` or `W` `A` `S` `D` | Slide all tiles |
| `p` / `Space` | Pause or resume |
| `Enter` | Keep playing after reaching 2048 |
| `r` | Start a fresh board |
| `m` | Pause and return to the launcher |
| `q` / `Ctrl+C` | Quit |

## Rules

Matching tiles merge when they meet. Each tile can merge once per move, so `2 2 4` becomes `4 4`, while `2 2 2 2` becomes `4 4`. Each merge adds the new tile's value to your score.

After a move changes the board, a new tile appears in an empty cell: a 2 with 90% probability, or a 4 with 10% probability. Blocked moves do not spawn tiles or increase the move count. Merged tiles briefly flash, and a new tile briefly shows dots around its number.

Reach 2048 to win, then press `Enter` to continue toward larger tiles. The game ends when no empty cells or adjacent matching tiles remain. The best score is retained across restarts during the current session.

Returning to the launcher preserves your board. Choose 2048 again and press `p` to resume. If the terminal is too small, resize it as instructed and press `p` to continue.
