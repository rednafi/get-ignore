<div align="center">

<h2>GetIgnore</h2>

<h4>A Pointless CLI to Download Gitignore Files with a Single Command</h4>

![title](https://github.com/rednafi/get-ignore/blob/master/art/Screenshot%20from%202020-04-29%2016-28-23.png)

</div>

## Description

This downloads `.gitignore` of the language(s) provided by the user.

```bash
getignore --languages Python
```

or,

```bash
getignore -languages Python Go Nim
```

In case of multiple languages, the contents of their corresponding `.gitignore` files will be appended to a single `.gitignore` file.

To see a list of all the supported languages, run:

```bash
getignore --list
```

Invoke the help with:

```bash
getignore -h
```

```
NAME:
   getignore - A Pointless CLI to Download Gitignore Files 📥

USAGE:
   getignore [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --languages value, --lg value  Provide the desired languages 🔥
   --list, --ls                   Show a list of available languages 📝 (default: false)
   --help, -h                     show help (default: false)
```

## Installation

If you have `go` installed in your system, then run:

```bash
go get github.com/rednafi/get-ignore/getignore
```
