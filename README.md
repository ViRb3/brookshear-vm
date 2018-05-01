# Brookshear VM

```
 _                     _        _                            _   ____  ___
| |                   | |      | |                          | | | |  \/  |
| |__  _ __ ___   ___ | | _____| |__   ___  __ _ _ __ ______| | | | .  . |
| '_ \| '__/ _ \ / _ \| |/ / __| '_ \ / _ \/ _` | '__|______| | | | |\/| |
| |_) | | | (_) | (_) |   <\__ \ | | |  __/ (_| | |         \ \_/ / |  | |
|_.__/|_|  \___/ \___/|_|\_\___/_| |_|\___|\__,_|_|          \___/\_|  |_/
```
A brookshear virtual machine implementation capable of executing and compiling brookshear assembly instructions.

## Features

* Executes brookshear assembly instructions
* Three levels of verbosity, allowing for easy step-by-step debugging of code
* Ability to compile to binary (brookshear bytecode)

## Help screen (run with `-help`)

```
  -a string
    	(REQUIRED) Action to perform, one of:
    	compile - compiles an input asm file to binary
    	execute - executes an input asm file
  -f string
    	(REQUIRED) Input file. Ignores blank lines or lines starting with '//' or '#'
  -v int
    	Set level of verbosity, one of:
    	0 - print memory dump on end
    	1 - print each executed instruction
    	2 - print data changes on each executed instruction
```