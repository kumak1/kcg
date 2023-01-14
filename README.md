# kcg

![license](https://img.shields.io/github/license/kumak1/kcg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kumak1/kcg)](https://goreportcard.com/report/github.com/kumak1/kcg)
![build](https://img.shields.io/github/actions/workflow/status/kumak1/kcg/release.yml)
![Go version](https://img.shields.io/github/go-mod/go-version/kumak1/kcg)
![release](https://img.shields.io/github/v/release/kumak1/kcg)

[日本語ドキュメント (Japanese Documents Available)](README_JA.md)

## Table of Contents

- [Overview](#overview)
- [Install](#install)
- [Getting Started](#getting-started)
    - [Configuration](#configuration)
    - [For ghq user](#for-ghq-user)
- [Usage](#usage)

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

default configuration file place is `~/.kcg` .

### Configuration

#### `Generate` setting file.

```shell
kcg configure init
```

#### `Add` or `Update` repository setting.

```shell
 kcg configure set <name> --repo="git@github.com:kumak1/kcg.git" --path="~/src/github.com/kumak1/kcg/"
```

##### Tips

`group` option set, useful for narrow down operation.

```shell
kcg configure set <name> --group="group_a" --group="group_b"
```
or
```shell
kcg configure add group <name> "group_c"
```

When `main` and `master` are mixed in the default branch, setting branch name alias.

example: `main` to `master`

```shell
kcg configure set <name> --branch-alias="main:master"
```
or
```shell
kcg configure add branch-alias <name> "main:master"
```

#### `Delete` repository setting.

```shell
kcg configure delete <name>
```

### For ghq user

#### `Import` setting file.

If you are using [ghq](https://github.com/x-motemen/ghq), you can import repository settings.
This command is non-destructive except for `--path` option configuration. recommend rerun when you have more repositories to manage with `ghq`.

```shell
kcg configure import --ghq
```

#### `Add` or `Update` repository setting.

`--path` option is not required.

```shell
 kcg configure set <name> --repo="git@github.com:kumak1/kcg.git"
```

## Usage

| command                    | description                                                    |
|:---------------------------|:---------------------------------------------------------------|
| `kcg ls`                   | Show repository data in configuration file.                    |
| `kcg cleanup`              | Delete merged branch on each repository in configuration file. |
| `kcg clone`                | Run `git clone` on each repository in configure file.          |
| `kcg pull`                 | Run `git pull` on each repository in configure file.           |
| `kcg switch <branch_name>` | Run `git switch` on each repository in configure file.         |

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
  exec        Run commands on each repository
  help        Help about any command
  ls          Show repository list.
  pull        run `git pull` on each repository dir
  switch      run `git switch` on each repository dir

Flags:
      --config string   config file (default is $HOME/.kcg)
  -h, --help            help for kcg
  -v, --version         version for kcg

Use "kcg [command] --help" for more information about a command.
```
