# WUS-Go
[![License](https://img.shields.io/github/license/riiconnect24/wus-go.svg?style=flat-square)](http://www.gnu.org/licenses/agpl-3.0)
![Production List](https://img.shields.io/discord/206934458954153984.svg?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/RiiConnect24/WUS-Go?style=flat-square)](https://goreportcard.com/report/github.com/RiiConnect24/Mail-Go)

WUS is a server which games check to see if your other Wii Friends have the game that you have.

We have found that only 4 games used it:

- Animal Crossing Wii
- WarioWare: D.I.Y. Showcase
- Wii Music
- Wii no Ma (no longer active, and is technically not a game)

There are two scripts, `inquiry` and `notify`.

- `inquiry` is sent a list of friend codes from the Wii for the server to check. In the same order, it returns a list of 0s and 1s. 0 means the Wii doesn't have the game, and 1 means it does.
- `notify` tells the server the Wii has the game.