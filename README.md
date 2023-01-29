# kcg

![license](https://img.shields.io/github/license/kumak1/kcg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kumak1/kcg)](https://goreportcard.com/report/github.com/kumak1/kcg)
![build](https://img.shields.io/github/actions/workflow/status/kumak1/kcg/release.yml)
![Go version](https://img.shields.io/github/go-mod/go-version/kumak1/kcg)
![release](https://img.shields.io/github/v/release/kumak1/kcg)
[![Coverage Status](https://coveralls.io/repos/github/kumak1/kcg/badge.svg)](https://coveralls.io/github/kumak1/kcg)

[日本語ドキュメント (Japanese Documents Available)](README_JA.md)

## Table of Contents

- [Overview](#overview)
- [Install](#install)
- [Getting Started](#getting-started)
    - [Quick Start Example](#quick-start-example)
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

| command                                                   | description                                                                                              |
|:----------------------------------------------------------|:---------------------------------------------------------------------------------------------------------|
| `kcg configure init`                                      | Generate configuration file (if not file exists)                                                         |
| `kcg configure import --ghq`                              | Add setting managed by [ghq](https://github.com/x-motemen/ghq)                                           |
| `kcg configure import --path="path/to/config"`            | Import settings from specified file into configure file ( `~/.kcg` )                                     |
| `kcg configure import --url="url/to/config"`              | Import settings from specified url file into configure file ( `~/.kcg` )                                 |
| `kcg configure set <name> --repo="git@host:org/repo.git"` | Set repository setting (required)                                                                        |
| `kcg configure set <name> --path="path/to/repo"`          | Set repository save path setting（required. [ghq](https://github.com/x-motemen/ghq) user is not required） |
| `kcg configure set <name> --group="group_a"`              | Set group setting                                                                                        | 
| `kcg configure add <name> --group="group_a"`              | Add group setting                                                                                        | 
| `kcg configure set <name> --branch-alias="main:master"`   | Set branch alias setting. <br> example: `main` to `master`                                               |
| `kcg configure add <name> --branch-alias="main:master"`   | Add branch alias setting. <br> example: `main` to `master`                                               |
| `kcg configure delete <name>`                             | Delete repository setting                                                                                |

#### Quick Start Example

```shell
kcg configure init
kcg configure set kumak1/kcg \
  --repo="git@github.com:kumak1/kcg.git" \
  --path="~/src/github.com/kumak1/kcg/"
```

if you using [ghq](https://github.com/x-motemen/ghq)

```shell
kcg configure import --ghq
```

#### Share config file Example

##### sender

```shell
kcg configure export --filter="kcg" | gh gist create --public
```

##### receiver

```shell
kcg configure import --url="gist_raw_file_url"
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

#### Command Details

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
