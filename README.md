# memcli
memcached cli

## Installation

Download from https://github.com/tkuchiki/memcli/releases

## Usage

```console
$ memcli --help
usage: memcli [<flags>] <command> [<args> ...]

memcached cli tool

Flags:
  --help                         Show context-sensitive help (also try --help-long and --help-man).
  --servers=127.0.0.1:11211 ...  servers
  --version                      Show application version.

Commands:
  help [<command>...]
    Show help.

  get --key=KEY
    get

  set --key=KEY [<flags>]
    set

  delete --key=KEY
    delete

  delete-all
    delete-all

  flush-all
    flush-all
```
