package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/taoshihan1991/imaptool/common"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "example:go-fly version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-fly " + common.Version)
	},
}
