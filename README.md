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
repos:
  kcg:
    repo: git@github.com:kumak1/kcg.git
    path: ~/src/github.com/kumak1/kcg
    setup:
      - make setup
```

if you use [ghq](https://github.com/x-motemen/ghq)

```shell
ghq: true
repos:
  kcg:
    repo: git@github.com:kumak1/kcg.git
    setup:
      - make setup
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
  setup       run setup commands on each repository
  switch      run `git switch` on each repository dir

Flags:
      --config string   config file (default is $HOME/.kcg)
  -h, --help            help for kcg

Use "kcg [command] --help" for more information about a command.

```
