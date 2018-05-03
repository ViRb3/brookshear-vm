package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"brookshear-vm/io"
)

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile assembly instructions to bytecode in a binary file",
	Example: `compile -f sample.asm.txt -o sample.bin`,
	Run:   compile,
}

func compile(cmd *cobra.Command, args []string) {
	if err := doCompile(); err != nil {
		fmt.Printf("Error during assembly compilation:\n%+v\n", err)
		return
	}
}

func doCompile() error {
	instrStrs, err := io.ReadAsmFile(Input)
	if err != nil {
		return err
	}

	//fmt.Println("Compiling to file:", Output)
	err = io.Compile(instrStrs, Output)
	//fmt.Println("Done!")
	return err
}

func init() {
	rootCmd.AddCommand(compileCmd)

	compileCmd.Flags().StringVarP(&Input, "file", "f", "", `Input assembly file path`)
	compileCmd.Flags().StringVarP(&Output, "out", "o", "", `Output binary file path`)

	compileCmd.MarkFlagRequired("file")
	compileCmd.MarkFlagRequired("out")
}