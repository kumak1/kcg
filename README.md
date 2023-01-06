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

### minimum

```shell
repos:
  kcg:
    repo: git@github.com:kumak1/kcg.git
    path: ~/src/github.com/kumak1/kcg
```

if you use [ghq](https://github.com/x-motemen/ghq)

```shell
ghq: true
repos:
  kcg:
    repo: git@github.com:kumak1/kcg.git
```

### more options

```shell
repos:
  with_setup_commands:
    repo: git@github.com:kumak1/with_setup_commands.git
    path: ~/src/github.com/kumak1/with_setup_commands
    setup:
      - "make setup"
  with_groups:
    repo: git@github.com:kumak1/with_groups.git
    path: ~/src/github.com/kumak1/with_groups
    groups:
      - ktools
  homebrew-ktools:
    repo: git@github.com:kumak1/homebrew-ktools.git
    path: ~/src/github.com/kumak1/homebrew-ktools
    groups:
      - ktools
```

## usage

```shell
% kcg -h
This is git command wrapper CLI.

Usage:
  kcg [command]

Available Commands:
  cleanup     delete merged branch on each repository dir
  clone       run `git clone` each repository
  completion  Generate the autocompletion script for the specified shell
  configure   Operate config file
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
