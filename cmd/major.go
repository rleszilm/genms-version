/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// majorCmd represents the major command
var majorCmd = &cobra.Command{
	Use:   "major",
	Short: "tool that returns a semver based off the repos history.",
	Long:  `major inspects the repository and uses the most recent tag to return a semantic version in the format "v<major>".`,
	Run:   major,
}

func init() {
	rootCmd.AddCommand(majorCmd)
}

func major(cmd *cobra.Command, args []string) {
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
		ver.IncMajor()
	}

	res, err := ver.Major()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(res)
}
