package main

import (
	"fmt"

	"github.com/cz-theng/ethspider"
	"github.com/spf13/cobra"
)

var (
	version bool
)

var rootCMD = &cobra.Command{
	Use:   "ethspider",
	Short: "start an ethspider",
	Long:  `start an ethspider`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			dumpVersion()
			return
		}
		cmd.Usage()
	},
}

func dumpVersion() {
	fmt.Printf("%s\n", ethspider.Version())
}

// Execute executes the root command.
func Execute() error {
	return rootCMD.Execute()
}

func init() {
	rootCMD.PersistentFlags().BoolVarP(&version, "version", "v", false, "Print version of ethspider")
}
