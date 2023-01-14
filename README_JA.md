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
  - [設定](#設定)
  - [ghq利用者の場合](#ghq利用者の場合)
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

### 設定

デフォルトの設定ファイルは `~/.kcg` に配置します。

#### 初期化

設定ファイルが存在しない場合は、以下のコマンドで生成します。

```shell
kcg configure init
```

#### 追加・更新

```shell
 kcg configure set <name> --repo="git@github.com:kumak1/kcg.git" --path="~/src/github.com/kumak1/kcg/"
```

##### Tips

`group` の設定をしておくと、後述の `pull` などの操作をする際に対象を絞れるので便利です。

```shell
kcg configure set <name> --group="group_a" --group="group_b"
```
または
```shell
kcg configure add group <name> "group_c"
```

管理対象リポジトリの default branch に `main` と `master` が混在する場合、
以下のコマンドで branch 名のエイリアスを設定できます。

例) `main` を指定したら `master` を操作

```shell
kcg configure set <name> --branch-alias="main:master"
```
または
```shell
kcg configure add branch-alias <name> "main:master"
```

#### 削除

```shell
kcg configure delete <name>
```

### ghq利用者の場合

#### 取り込み

[ghq](https://github.com/x-motemen/ghq) コマンドを利用している場合、以下のコマンドで ghq で管理しているリポジトリを元に設定ファイルを生成できます。
このコマンドは `--path` オプションの設定値以外は非破壊的に動作します。`ghq` で管理するリポジトリが増えたらまた実行することをおすすめします。

```shell
kcg configure import --ghq
```

#### 追加・更新

`--path` オプションは不要です。

```shell
 kcg configure set <name> --repo="git@github.com:kumak1/kcg.git"
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
