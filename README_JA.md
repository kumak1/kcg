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
  - [設定ファイル](#設定ファイル)
    - [ghq利用者の場合](#ghq利用者の場合)
- [基本的な使い方](#基本的な使い方)
  - [clone](#clone)
  - [ls](#ls)
  - [cleanup](#cleanup)
  - [switch](#switch)
  - [pull](#pull)
  - [Command Details](#command-details)

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

### 設定ファイル

デフォルトの設定ファイルは `~/.kcg` に配置します。

#### 初期化

設定ファイルが存在しない場合は、以下のコマンドで生成します。

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

#### 追加・更新

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

#### 削除

```shell
kcg configure delete <name>
```

#### ghq利用者の場合

[ghq](https://github.com/x-motemen/ghq) コマンドを利用している場合、以下のコマンドで ghq で管理しているリポジトリを元に設定ファイルを生成できます。

```shell
kcg configure init --import-from-ghq
```

このコマンドは `--path` オプションの設定値以外は非破壊的に動作します。`ghq` で管理するリポジトリが増えたらまた実行することをおすすめします。

## 基本的な使い方

### clone

設定ファイルに記載されたリポジトリを `git clone` します。

```shell
kcg clone
```

`--filter="needle"` や `--group="group_name"` で対象をリポジトリを絞ることが可能です。

### ls

設定ファイルに記載されたリポジトリの状態を一覧します。

```shell
kcg ls
```

`--filter="needle"` や `--group="group_name"` で対象をリポジトリを絞ることが可能です。

### cleanup

設定ファイルに記載されたリポジトリの local branch のうち、remote で merge 済みの branch を削除します。

```shell
kcg cleanup
```

`--filter="needle"` や `--group="group_name"` で対象をリポジトリを絞ることが可能です。

### switch

設定ファイルに記載されたリポジトリを `git switch` します。

```shell
kcg switch <branch_name>
```

`--filter="needle"` や `--group="group_name"` で対象をリポジトリを絞ることが可能です。

#### default branch に main と master が混在する場合

以下のコマンドで branch 名のエイリアス（例: `main` を指定したら `master` を操作 ）を設定できます。

```shell
kcg configure set <name> --branch-alias="main:master"
```

### pull

設定ファイルに記載されたリポジトリを `git pull` します。

```shell
kcg pull
```

`--filter="needle"` や `--group="group_name"` で対象をリポジトリを絞ることが可能です。

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
