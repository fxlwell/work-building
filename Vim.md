# work-building
```
mkdir -p ~/.vim/bundle
git clone https://github.com/gmarik/Vundle.vim.git ~/.vim/bundle/Vundle.vim
vim ~/.vimrc

syntax on
" tab宽度和缩进同样设置为4
set tabstop=4
set softtabstop=4
set shiftwidth=4
 
set nocompatible
 
" 你在此设置运行时路径
set rtp+=~/.vim/bundle/Vundle.vim
 
call vundle#begin()
 
" Vundle 本身就是一个插件
Plugin 'gmarik/Vundle.vim'
 
 
"所有插件都应该在这一行之前 
call vundle#end()
 
" filetype off
filetype plugin indent on

:PluginInstall
Plugin 'fatih/vim-go'

:GoInstallBinaries 安装必要工具
```
