/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// patchCmd represents the patch command
var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "tool that returns a semver based off the repos history.",
	Long:  `patch inspects the repository and uses the most recent tag to return a semantic version in the format "v<major>.<minor>.<patch>".`,
	Run:   patch,
}

func init() {
	rootCmd.AddCommand(patchCmd)
}

func patch(cmd *cobra.Command, args []string) {
	ver, err := buildVersion(cmd, args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inc, err := cmd.Flags().GetBool("increment")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if inc {
		ver.IncPatch()
	}

	res, err := ver.Patch()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(res)
}
