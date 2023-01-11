# kcg

![license](https://img.shields.io/github/license/kumak1/kcg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kumak1/kcg)](https://goreportcard.com/report/github.com/kumak1/kcg)
![build](https://img.shields.io/github/actions/workflow/status/kumak1/kcg/release.yml)
![Go version](https://img.shields.io/github/go-mod/go-version/kumak1/kcg)
![release](https://img.shields.io/github/v/release/kumak1/kcg)

[English Documents Available(�Ѹ�ɥ������)](README.md)

kumak1 Convenient Git tools.
inspired by [pdr](https://github.com/pyama86/pdr).

## �ܼ�

- [����](#����)
- [���󥹥ȡ���](#���󥹥ȡ���)
- [���åȥ��å�](#���åȥ��å�)
  - [����ե�����](#����ե�����)
    - [ghq���ѼԤξ��](#ghq���ѼԤξ��)
- [����Ū�ʻȤ���](#����Ū�ʻȤ���)
  - [clone](#clone)
  - [ls](#ls)
  - [cleanup](#cleanup)
  - [switch](#switch)
  - [pull](#pull)
  - [Command Details](#command-details)

## ����

���ץ곫ȯ��ʣ����ݥ��ȥ����Ѥ��Ƥ��ꡢ���줾�죱�ģ��� `git switch main` �� `git pull` ����Τ��Ѥ路�����������Ȥ����뤫�Ȼפ��ޤ��� `kcg` �Ϥ��μ�֤򾯤��������餷�Ƥ����ġ���Ǥ���

#### ��ħ

- ʣ���Υ�ݥ��ȥ���ñ�˴�������
- �оݤ� `filter` �� `group` �ǹʤꤳ���

## ���󥹥ȡ���

#### homebrew

```shell
brew tap kumak1/homebrew-ktools 
brew install kcg
```

#### go

```shell
go get github.com/kumak1/kcg@latest
```

## ���åȥ��å�

### ����ե�����

�ǥե���Ȥ�����ե������ `~/.kcg` �����֤��ޤ��� 
����ե����뤬¸�ߤ��ʤ����ϡ��ʲ��Υ��ޥ�ɤ��������ޤ���

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

������ɲá������򤹤���ϡ��ʲ��Υ��ޥ�ɤ�¹Ԥ��ޤ���

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

����������������ϡ��ʲ��Υ��ޥ�ɤ�¹Ԥ��ޤ���

```shell
kcg configure delete <name>
```

#### ghq���ѼԤξ��

[ghq](https://github.com/x-motemen/ghq) ���ޥ�ɤ����Ѥ��Ƥ����硢�ʲ��Υ��ޥ�ɤ� ghq �Ǵ������Ƥ����ݥ��ȥ�򸵤�����ե�����������Ǥ��ޤ���

```shell
kcg configure init --import-from-ghq
```


## ����Ū�ʻȤ���

### clone

����ե�����˵��ܤ��줿��ݥ��ȥ�� `git clone` ���ޤ���

```shell
kcg clone
```

`--filter="needle"` �� `--group="group_name"` ���оݤ��ݥ��ȥ��ʤ뤳�Ȥ���ǽ�Ǥ���

### ls

����ե�����˵��ܤ��줿��ݥ��ȥ�ξ��֤�������ޤ���

```shell
kcg ls
```

`--filter="needle"` �� `--group="group_name"` ���оݤ��ݥ��ȥ��ʤ뤳�Ȥ���ǽ�Ǥ���

### cleanup

����ե�����˵��ܤ��줿��ݥ��ȥ�� local branch �Τ�����remote �� merge �Ѥߤ� branch �������ޤ���

```shell
kcg cleanup
```

`--filter="needle"` �� `--group="group_name"` ���оݤ��ݥ��ȥ��ʤ뤳�Ȥ���ǽ�Ǥ���

### switch

����ե�����˵��ܤ��줿��ݥ��ȥ�� `git switch` ���ޤ���

```shell
kcg switch <branch_name>
```

`--filter="needle"` �� `--group="group_name"` ���оݤ��ݥ��ȥ��ʤ뤳�Ȥ���ǽ�Ǥ���

#### default branch �� main �� master �����ߤ�����

�ʲ��Υ��ޥ�ɤ� branch ̾�Υ����ꥢ������: `main` ����ꤷ���� `master` ����� �ˤ�����Ǥ��ޤ���

```shell
kcg configure set <name> --branch-alias="main:master"
```

### pull

����ե�����˵��ܤ��줿��ݥ��ȥ�� `git pull` ���ޤ���

```shell
kcg pull
```

`--filter="needle"` �� `--group="group_name"` ���оݤ��ݥ��ȥ��ʤ뤳�Ȥ���ǽ�Ǥ���

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
