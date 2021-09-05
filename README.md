# gotp

Command-line TOTP Tool

This tool enable you to manage your TOTP (such as google authenticator) in command-line.

[![GitHub release](https://img.shields.io/github/tag/kzk-maeda/gotp.svg?label=latest)](https://github.com/kzk-maeda/gotp/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/kzk-maeda/gotp)](https://goreportcard.com/report/github.com/kzk-maeda/gotp)
[![License](https://img.shields.io/badge/license-BSD-blue.svg)](./LICENSE.md)

## Overview

gotp is a command-line tool that enable you to manage your TOTP keys simply and enable you select one time password in command line


## Basic Usage

### Install and configure

under construncting...

### Add your secret to config file

You can add your secret to your config file just running the command `gotp add`

```bash
$ ./gotp add
Initializing your config og gotp
Input key name: <Name>
Name
Input secret key: <Your Secret Key>
Your Secret Key
```

the config file is created in `$HOME/.gotp/config.yml` as multi-document yaml file.

you can also manage your key by editing that config file.

### Select key and get One-time password

You can get one-time password just running the command `gotp`

```bash
$ gotp
GOTP
Use the arrow keys to navigate: ↓ ↑ → ←
? Select Key:
  ▸ google
    docbase
    ...

✔ google
346068
```

and then, the one-time password is saved in your clipboard, so you can just paste it to use.