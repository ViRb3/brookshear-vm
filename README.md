# Brookshear VM

```
 _                     _        _                            _   ____  ___
| |                   | |      | |                          | | | |  \/  |
| |__  _ __ ___   ___ | | _____| |__   ___  __ _ _ __ ______| | | | .  . |
| '_ \| '__/ _ \ / _ \| |/ / __| '_ \ / _ \/ _` | '__|______| | | | |\/| |
| |_) | | | (_) | (_) |   <\__ \ | | |  __/ (_| | |         \ \_/ / |  | |
|_.__/|_|  \___/ \___/|_|\_\___/_| |_|\___|\__,_|_|          \___/\_|  |_/
```

A [brookshear](https://uk.mathworks.com/matlabcentral/fileexchange/22593-extended-brookshear-machine-emulator-and-assembler?focused=5204034&tab=example) virtual machine with support for step-by-step emulation, compilation, and decompilation.

## Features

* Run assembly code or binary bytecode
* Various verbosity levels allow detail-rich, step-by-step debugging
* Compile assembly instructions to bytecode in a binary file
* Decompile bytecode from a binary file to assembly instructions
* Extra instruction *(opcode `0xD`)* for relative (offset) branch

## Relative branch if equal

```
Examples:

Bytecode | Instruction  | Comment
-----------------------------------------------------------------------
D104     | jmpeq +4, r1 | ; branch forward by 4 bytes (2 instructions)
D312     | jmpeq -2, r3 | ; branch backward by 2 bytes (1 instruction)
```

It accepts a single nibble as offset, and can be either positive (`+4`) or negative (`-2`). When compiling, the instruction is encoded as follows:

```
D - opcode                  D
R - destination register  [0-F]
O - branch direction      [0-1] : 0 for +/forward, 1 for -/backward
X - offset                [0-F]
```

## [Releases](./releases)

*For general help run with argument `-h` or `--help`*

*For help on the extra instructions run the command `extendedHelp`*