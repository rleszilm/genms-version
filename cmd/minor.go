/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// minorCmd represents the minor command
var minorCmd = &cobra.Command{
	Use:   "minor",
	Short: "tool that returns a semver based off the repos history.",
	Long:  `minor inspects the repository and uses the most recent tag to return a semantic version in the format "v<major>.<minor>".`,
	Run:   minor,
}

func init() {
	rootCmd.AddCommand(minorCmd)
}

func minor(cmd *cobra.Command, args []string) {
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
		ver.IncMinor()
	}

	res, err := ver.Minor()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(res)
}
