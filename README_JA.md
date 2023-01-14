# kcg

![license](https://img.shields.io/github/license/kumak1/kcg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kumak1/kcg)](https://goreportcard.com/report/github.com/kumak1/kcg)
![build](https://img.shields.io/github/actions/workflow/status/kumak1/kcg/release.yml)
![Go version](https://img.shields.io/github/go-mod/go-version/kumak1/kcg)
![release](https://img.shields.io/github/v/release/kumak1/kcg)

[English Documents Available(英語ドキュメント)](README.md)

## 目次

- [概要](#概要)
- [インストール](#インストール)
- [セットアップ](#セットアップ)
  - [設定例](#設定例)
- [基本的な使い方](#基本的な使い方)

## 概要

アプリ開発で複数リポジトリ利用しており、それぞれ１つ１つ `git switch main` や `git pull` するのが煩わしく感じたことがあるかと思います。 `kcg` はこの手間を少しだけ減らしてくれるツールです。

#### 特徴

- 複数のリポジトリを簡単に管理する
- 対象を `filter` や `group` で絞りこめる

## インストール

#### homebrew

```shell
brew tap kumak1/homebrew-ktools 
brew install kcg
```

#### go

```shell
go get github.com/kumak1/kcg@latest
```

## セットアップ

デフォルトの設定ファイルは `~/.kcg` に配置します。

| command                                                   | description                                                     |
|:----------------------------------------------------------|:----------------------------------------------------------------|
| `kcg configure init`                                      | 設定ファイルを生成します（既に存在する場合はなにもしない）                                   |
| `kcg configure import --ghq`                              | [ghq](https://github.com/x-motemen/ghq) で管理しているリポジトリを元に設定を追加します |
| `kcg configure import --path="path/to/config"`            | 設定ファイル( `~/.kcg` ) に指定設定ファイルの設定を取り込みます                          |
| `kcg configure set <name> --repo="git@host:org/repo.git"` | 管理リポジトリを追加します（必須情報）                                             |
| `kcg configure set <name> --path="path/to/repo"`          | 管理リポジトリの保存場所を設定します（必須情報だが、ghqユーザは指定不要）                          |
| `kcg configure set <name> --group="group_a"`              | グループ設定します                                                       | 
| `kcg configure add <name> --group="group_a"`              | グループ設定を追加します                                                    | 
| `kcg configure set <name> --branch-alias="main:master"`   | branch 名のエイリアスを設定します<br>例) `main` を指定したら `master` を操作           |
| `kcg configure add <name> --branch-alias="main:master"`   | branch 名のエイリアスを追加します<br>例) `main` を指定したら `master` を操作           |
| `kcg configure delete <name>`                             | 管理リポジトリを削除します                                                   |

#### 設定例

```shell
kcg configure init
kcg configure set kumak1/kcg \
  --repo="git@github.com:kumak1/kcg.git" \
  --path="~/src/github.com/kumak1/kcg/"
```

## 基本的な使い方

| command                    | description                                               |
|:---------------------------|:----------------------------------------------------------|
| `kcg ls`                   | リポジトリの状態を一覧表示します。                                         |
| `kcg cleanup`              | リポジトリの local branch のうち、remote で merge 済みの branch を削除します。 |
| `kcg clone`                | リポジトリを `git clone` します。                                   |
| `kcg pull`                 | リポジトリを `git pull` します。                                    |
| `kcg switch <branch_name>` | リポジトリを `git switch` します。                                  |

上記のコマンド全てで `--filter="needle"` や `--group="group_name"` オプションによって対象リポジトリを絞ることが可能です。

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
