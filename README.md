# MyMac .dotfiles Cheatsheets

## ZSH-Cheatsheet

| Shortcut | Function                                          |
| -------- | ------------------------------------------------- |
| ll       | List dir                                          |
| lt       | List dir w/ tree                                  |
| so       | Reset config.                                     |
| zconfig  | Open zsh config file                              |
| kconfig  | Open kitty config file                            |
| vconfig  | Open nvim .config                                 |
| tconfig  | Open tmux.conf file                               |
| cda      | Activate data analysis env.                       |
| base     | Activate base env.                                |
| qt       | Activate Qt env.                                  |
| v.       | Open nvim project on current directory            |
| gitback  | Reset git repo in current directory               |
| vqt      | Open nvim project and activate Qt env.            |
| vda      | Open nvim project and activate data analysis env. |
| proj     | Change my project directory                       |
| fof      | Search and open file using fzf (external)         |
| fod      | Search and open director (external)               |
| fcd      | Search and change directory                       |
| ff       | Search and get path the file.                     |

## NVIM–Cheatsheet

| Function                        | Pluggin                | shortcuts                                 |
| ------------------------------- | ---------------------- | ----------------------------------------- |
| Change bracket                  | tpope/vim-surround     | cs (ex.: cs"' to change "this" to 'this') |
| make comment                    | tpope/vim-commentary   | gcc or gc                                 |
| Show tagbar panel               | preservim/tagbar       | F8                                        |
| Multi select                    | mg979/vim-visual-ulti | ==???==                                   |
| Find files                      | Telescope              | `<space>ff`                               |
| Live grep                       | Telescope              | `<space>fg`                               |
| Buffers                         | Telescope              | `<space>fb`                               |
| Help tags                       | Telescope              | `<space>fh`                               |
| lsp buf declaration             | lsp                    | `gD`                                      |
| lsp buf definition              | lsp                    | `gd`                                      |
| lsp buf hover                   | lsp                    | 'K'                                       |
| lsp buf implementation          | lsp                    | `gi`                                      |
| lsp buf signature               | lsp                    | `C-k`                                     |
| lsp buf add_workspace_folder    | lsp                    | `<space>wa`                               |
| lsp buf remove_workspace_folder | lsp                    | `<space>wr`                               |
| lsp print inspect               | lsp                    | `<space>wl`                               |
| lsp type definition             | lsp                    | `<space>D`                                |
| lsp rename                      | lsp                    | `<space>rn`                               |
| lsp code-action                 | lsp                    | `<space>ca`                               |
| lsp references                  | lsp                    | `gr`                                      |
| lsp formating                   | lsp                    | `<space>f`                                |
| File explorer toggle            | NvimTree               | `<space>p`                                |
| File explorer refresh           | NvimTree               | `<space>r`                                |
| File explorer find file         | NvimTree               | `<space>n`                                |
| Manual Formating                | Formatter              | F9                                        |
| Navigation key-h                | keymap                 | n: `<space>hh`                            |
| Navigation key-j                | keymap                 | n: `<space>jj`                            |
| Navigation key-k                | keymap                 | n: `<space>kk`                            |
| Navigation key-l                | keymap                 | n: `<space>ll`                            |
| Navigation buffer-next          | keymap                 | n: Shift+l                                |
| Navigation buffer-previous      | keymap                 | n: Shift+h                                |
| indent+                         | keymap                 | v: "<"                                    |
| indent-                         | keymap                 | v: ">"                                    |
| Move up                         | keymap                 | v: Shift+j                                |
| Move down                       | keymap                 | v: Shift+k                                |
| Paste+                          | keymap                 | v: "p"                                    |
| Open terminal                   | toggleterm             | Ctrl+\                                    |
| Open LazyGit                    | toggleterm             | n: `<space>gg`                            |
| Lsp Diagnostics All             | lspfuzzy               | n: `gf`                                   |

## Requirements:

```php
brew install ctags
brew install fd
brew install ripgrep
brew install fzf
brew install bat
brew install node
brew install tree
brew install miniconda
brew install exa
brew install git
brew install lazygit
brew install neovim
brew install peco
brew install stow
brew install tmux

```

### Python-LSP

- pyright: `sudo npm i -g pyright`
- Treesitter config: `:TSInstall python` | to update `:TSUpdate python`
- Formating:
	- black
	- isort

#### pyrightconfig.json

Example ignore error message:

```json

{
        "reportOptionalMemberAccess": false,
}

```
---
# References

- [🌐 Getting Started With Vim - A Practical Guide](https://www.jakewiesler.com/blog/getting-started-with-vim#neovim)
- [🌐 Switching to NeoVim Native LSP | Mark Story](https://mark-story.com/posts/view/switching-to-neovim-native-lsp)
- [📽️ Up And Running With Neovim + LSP](https://youtu.be/36o52-2_83M)
- [📽️ Neovim - NvimTree File Explorer Written In Lua](https://youtu.be/SpexCBrZ1pQ)
- [craftzdog/dotfiles-public: My personal dotfiles](https://github.com/craftzdog/dotfiles-public)
- [jakewies/.dotfiles: My portable development environment](https://github.com/jakewies/.dotfiles)
- [ChristianChiarulli/nvim: My neovim config](https://github.com/ChristianChiarulli/nvim)
- [How to set up Neovim 0.5](https://blog.inkdrop.app/how-to-set-up-neovim-0-5-modern-plugins-lsp-treesitter-etc-542c3d9c9887)
- [pyright/docs/configuration](pyright-doc.md)
- [How to set up Neovim 0.5 + Modern plugins (LSP, Treesitter, Fuzzy finder, etc) | by Takuya Matsuyama | Dev as Life](https://blog.inkdrop.app/how-to-set-up-neovim-0-5-modern-plugins-lsp-treesitter-etc-542c3d9c9887)

