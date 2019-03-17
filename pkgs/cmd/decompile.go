package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"brookshear-vm/pkgs/io"
)

// decompileCmd represents the decompile command
var decompileCmd = &cobra.Command{
	Use:   "decompile",
	Short: "Decompile bytecode from a binary file to assembly instructions",
	Example: `decompile -f sample.bin`,
	Run: decompile,
}

func decompile(cmd *cobra.Command, args []string) {
	if err := doDecompile(); err != nil {
		fmt.Printf("Error during binary decompilation:\n%+v\n", err)
		return
	}
}

func doDecompile() error {
	instrStr, err := io.Decompile(Input)
	if err != nil {
		return err
	}

	for _, instr := range instrStr {
		fmt.Println(instr.ToString())
	}
	return nil
}

func init() {
	rootCmd.AddCommand(decompileCmd)

	decompileCmd.Flags().StringVarP(&Input, "file", "f", "", `Input binary file path`)

	decompileCmd.MarkFlagRequired("file")
}