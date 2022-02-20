local treesitter = require('nvim-treesitter.configs')

treesitter.setup {
    ensure_installed = { "python", "javascript", "css", "html", "tsx", "json", "yaml", "scss" },
    sync_install = false,
    ignore_install = { "" },
    highlight = {
        enable = true,
        disable = { "" },
        additional_vim_regex_highlighting = false,
    },
    indent = {
        enable = false,
        disable = {},
    },
    rainbow = {
        enable = true,
        extend_mode = true,
        max_file_lines = nil,
    },
}

local parser_config = require "nvim-treesitter.parsers".get_parser_configs()
parser_config.tsx.filetype_to_parsername = { "javascript", "typescript.tsx" }
