local treesitter = require('nvim-treesitter.configs')

treesitter.setup {
    ensure_installed = { "python", "javascript", "css", "html" },
    sync_install = false,
    ignore_install = { "" },
    highlight = {
        enable = true,
        disable = { "" },
        additional_vim_regex_highlighting = false,
    },
    rainbow = {
        enable = true,
        extend_mode = true,
        max_file_lines = nil,
    }
}
