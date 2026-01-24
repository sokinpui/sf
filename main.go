package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

func main() {
	var fileType string
	var excludes []string

	rootCmd := &cobra.Command{
		Use:   "sf [path]",
		Short: "A fast directory walker",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			roots := args
			if len(roots) == 0 {
				roots = []string{"."}
			}

			engine := NewEngine(runtime.NumCPU()*2, fileType, excludes)
			results := make(chan string, 100)

			go engine.Walk(roots, results)

			for path := range results {
				fmt.Println(path)
			}
		},
	}

	rootCmd.Flags().StringVarP(&fileType, "type", "t", "", "Filter by type: file, dir")
	rootCmd.Flags().StringSliceVarP(&excludes, "exclude", "e", []string{}, "Exclude entries that match the given glob pattern")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
