package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

func main() {
	var fileType string

	rootCmd := &cobra.Command{
		Use:   "gf [path]",
		Short: "A fast directory walker",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			root := "."
			if len(args) > 0 {
				root = args[0]
			}

			engine := NewEngine(root, runtime.NumCPU()*2, fileType)
			results := make(chan string, 100)

			go engine.Walk(root, results)

			for path := range results {
				fmt.Println(path)
			}
		},
	}

	rootCmd.Flags().StringVarP(&fileType, "type", "t", "", "Filter by type: file, dir")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
