" Basic settings
set nocompatible
set number
set relativenumber
set expandtab
set tabstop=4
set shiftwidth=4
set softtabstop=4
set autoindent
set smartindent
set showmatch
set hlsearch
set incsearch
set ignorecase
set smartcase
set ruler
set laststatus=2
set encoding=utf-8
set fileencoding=utf-8
set termencoding=utf-8

" Go specific settings
syntax on
filetype plugin indent on
set foldmethod=syntax
set nofoldenable

" Enable mouse support
set mouse=a

" Enable clipboard support
set clipboard=unnamed

" Set leader key to space
let mapleader = " "

" Basic key mappings
nnoremap <leader>w :w<CR>
nnoremap <leader>q :q<CR>
nnoremap <leader>e :e<CR>
nnoremap <leader>h :nohlsearch<CR>

" Go to last position in file
if has("autocmd")
    au BufReadPost * if line("'\"") > 1 && line("'\"") <= line("$") | exe "normal! g`\"" | endif
endif

" Highlight trailing whitespace
highlight ExtraWhitespace ctermbg=red guibg=red
match ExtraWhitespace /\s\+$/
autocmd BufWinEnter * match ExtraWhitespace /\s\+$/
autocmd InsertEnter * match ExtraWhitespace /\s\+\%#\@<!$/
autocmd InsertLeave * match ExtraWhitespace /\s\+$/
autocmd BufWinLeave * call clearmatches()

" Auto remove trailing whitespace on save
autocmd BufWritePre * :%s/\s\+$//e 