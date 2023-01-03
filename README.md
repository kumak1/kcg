# kcg

kumak1 Convenient Git tools.
inspired by [pdr](https://github.com/pyama86/pdr).

## install

```shell
brew tap kumak1/homebrew-ktools 
brew install kcg
```

## configure

`~/.kcg`

```shell
ghq: true
repos:
  kumaoche:
    repo: git@github.com:kumak1/kcg.git
    path: ~/src/github.com/kumak1/kcg
```

## usage

```shell
% kcg -h                        
This is git command wrapper CLI.

Usage:
  kcg [command]

Available Commands:
  clone       run `git clone` each repository
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ls          Show repository list.
  pull        run `git pull` on each repository dir
  setup       run configfile setupCommands each repository
  switch      run `git switch` on each repository dir

Flags:
      --config string   config file (default is $HOME/.kcg)
  -h, --help            help for kcg

Use "kcg [command] --help" for more information about a command.
```
