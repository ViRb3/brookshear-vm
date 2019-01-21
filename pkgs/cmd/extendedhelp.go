package cmd

import (
	"github.com/spf13/cobra"
)

// extendedHelp represents the extendedhelp command
var extendedHelp = &cobra.Command{
	Use:   "extendedHelp",
	Short: "Extra instructions help",
	Long: `This release contains an extra instruction (opcode 0xD) for relative (offset) branch.
It accepts a single nibble as offset, and can be either positive (+4) or negative (-2).
When compiling, the instruction is encoded as follows:

D - opcode                  D
R - destination register  [0-F]
O - branch direction      [0-1] : 0 for +/forward, 1 for -/backward
X - offset                [0-F]

Examples:

Bytecode | Instruction  | Comment
-----------------------------------------------------------------------
D104     | jmpeq +4, r1 | ; branch forward by 4 bytes (2 instructions)
D312     | jmpeq -2, r3 | ; branch backward by 2 bytes (1 instruction)`,
}

func init() {
	rootCmd.AddCommand(extendedHelp)
}