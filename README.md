# kcg

![license](https://img.shields.io/github/license/kumak1/kcg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kumak1/kcg)](https://goreportcard.com/report/github.com/kumak1/kcg)
![build](https://img.shields.io/github/actions/workflow/status/kumak1/kcg/release.yml)
![Go version](https://img.shields.io/github/go-mod/go-version/kumak1/kcg)
![release](https://img.shields.io/github/v/release/kumak1/kcg)

[日本語ドキュメント(Japanese Documents Available)](README_JA.md)

kumak1 Convenient Git tools.
inspired by [pdr](https://github.com/pyama86/pdr).

## Table of Contents

- [Overview](#overview)
- [Install](#install)
- [Getting Started](#getting-started)
    - [Configuration](#configuration)
        - [For ghq user](#for-ghq-user)
- [Usage](#usage)
    - [clone](#clone)
    - [ls](#ls)
    - [cleanup](#cleanup)
    - [switch](#switch)
    - [pull](#pull)
    - [Command Details](#command-details)

## Overview

If you use multiple repositories in your application development, you may have found it tedious to `git switch main` and `git pull` for each one. `kcg` is a tool that reduces this hassle a little.

#### Features

- Easily manage multiple repositories
- Narrow down the target with `filter` or `group` option.

## Install

#### homebrew

```shell
brew tap kumak1/homebrew-ktools 
brew install kcg
```

#### go

```shell
go get github.com/kumak1/kcg@latest
```

## Getting Started

### Configuration

default configuration file place is `~/.kcg` .

`Generate` setting file.

```shell
kcg configure init
```

<details>
<summary>more option</summary>

```shell
kcg configure init -h
Create an empty config file

Usage:
  kcg configure init [flags]

Flags:
  -h, --help                       help for init
      --import-from-ghq ghq list   create from ghq list
      --path string                write config file path

Global Flags:
      --config string   config file (default is $HOME/.kcg)
```
</details>

`Add` or `Update` repository setting.

```shell
 kcg configure set <name> --repo="git@github.com:kumak1/kcg.git" --path="~/src/github.com/kumak1/kcg/"
```

<details>
<summary>more option</summary>

```shell
kcg configure set -h 
Add repository config

Usage:
  kcg configure set <name> [flags]

Flags:
      --branch-alias stringArray   specify like "NAME:VALUE"
      --group stringArray          group
  -h, --help                       help for set
      --path string                local dir
      --repo string                remote repository
      --setup stringArray          setup command

Global Flags:
      --config string   config file (default is $HOME/.kcg)
```

</details>

`Delete` repository setting.

```shell
kcg configure delete <name>
```

#### For ghq user

If you are using [ghq](https://github.com/x-motemen/ghq), you can import repository settings.

```shell
kcg configure init --import-from-ghq
```

## Usage

### clone

Run `git clone` on each repository in configure file.

```shell
kcg clone
```

Can use narrow down repository option. `--filter="needle"` `--group="group_name"` 

### ls

Show repository data in configuration file.

```shell
kcg ls
```

Can use narrow down repository option. `--filter="needle"` `--group="group_name"`

### cleanup

Delete merged branch on each repository in configuration file.

```shell
kcg cleanup
```

Can use narrow down repository option. `--filter="needle"` `--group="group_name"`

### switch

Run `git switch` on each repository in configure file.

```shell
kcg switch <branch_name>
```

Can use narrow down repository option. `--filter="needle"` `--group="group_name"`

#### When main and master are mixed in the default branch

Setting branch name alias.
example: `main` to `master`

```shell
kcg configure set <name> --branch-alias="main:master"
```

### pull

Run `git pull` on each repository in configure file.

```shell
kcg pull
```

Can use narrow down repository option. `--filter="needle"` `--group="group_name"`

### Command Details

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
  -v, --version         version for kcg

Use "kcg [command] --help" for more information about a command.
```
