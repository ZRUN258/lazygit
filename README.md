## Windows LazyVim Fix

This fork carries a Windows-focused fix for a LazyVim/Neovim terminal input bug.
When lazygit is opened from LazyVim with `<leader>gg`, some Windows users see every key processed twice: `Tab` moves two rows, `j` moves two rows, and commit-message text appears as doubled characters.

The change is in the tcell-to-gocui input boundary:

- `pkg/gocui/tcell_driver.go` forces tcell to use the legacy keyboard protocol when lazygit starts.
- `pkg/gocui/tcell_driver.go` also ignores immediately repeated tcell key events whose full signature is identical.
- `pkg/gocui/gui.go` stores the last tcell key event signature used by that duplicate filter.
- `pkg/gocui/tcell_driver_test.go` covers the keyboard protocol selection, key-release handling, and duplicate-event filtering.

This is intended as a temporary Windows build for users affected by duplicate input in LazyVim. If the upstream lazygit release you are using no longer has this bug, prefer the official release.
