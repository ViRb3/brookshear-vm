# Brookshear VM

```
 _                     _        _                            _   ____  ___
| |                   | |      | |                          | | | |  \/  |
| |__  _ __ ___   ___ | | _____| |__   ___  __ _ _ __ ______| | | | .  . |
| '_ \| '__/ _ \ / _ \| |/ / __| '_ \ / _ \/ _` | '__|______| | | | |\/| |
| |_) | | | (_) | (_) |   <\__ \ | | |  __/ (_| | |         \ \_/ / |  | |
|_.__/|_|  \___/ \___/|_|\_\___/_| |_|\___|\__,_|_|          \___/\_|  |_/
```

A brookshear virtual machine with support for step-by-step emulation, compilation, and decompilation.

## Features

* Run an assembly or binary file
* Various verbosity levels allow detail-rich, step-by-step debugging
* Compile assembly instructions to bytecode in a binary file
* Decompile bytecode from a binary file

## Help screen (run with `-h` or `--help`)

```
Usage:
  brookshear-vm [command]

Available Commands:
  compile     Compile assembly instructions to bytecode in a binary file
  decompile   Decompile bytecode from a binary file
  help        Help about any command
  run         Run an assembly or binary file
```