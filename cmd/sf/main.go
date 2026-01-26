package main

import (
	"fmt"
	"os"

	"github.com/sokinpui/sf"
	"github.com/spf13/cobra"
)

func main() {
	var fileType string
	var excludes []string
	var showHidden bool

	rootCmd := &cobra.Command{
		Use:   "sf [path]",
		Short: "A fast directory walker",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			results := sf.Run(args, fileType, excludes, showHidden)

			for _, path := range results {
				fmt.Println(path)
			}
		},
	}

	rootCmd.Flags().StringVarP(&fileType, "type", "t", "", "Filter by type: file, dir")
	rootCmd.Flags().StringSliceVarP(&excludes, "exclude", "e", []string{}, "Exclude entries that match the given glob pattern")
	rootCmd.Flags().BoolVarP(&showHidden, "hidden", "H", false, "Search hidden files and directories")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
