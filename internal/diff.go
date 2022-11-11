package internal

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

func toAbsolutePath(relativePath string) (string, error) {
	return filepath.Abs(relativePath)
}

func Diff(cmd *cobra.Command, args []string) {
	startPath := cmd.Flag("start").Value.String()
	destinationPath := cmd.Flag("destination").Value.String()

	if !(len(startPath) > 0 && len(destinationPath) > 0) {
		startPath = args[0]
		destinationPath = args[1]
	}

	absoluteStartPath, err := toAbsolutePath(startPath)
	if err != nil {
		panic(err)
	}

	absoluteDestinationPath, err := toAbsolutePath(destinationPath)
	if err != nil {
		panic(err)
	}

	relativePath, err := filepath.Rel(absoluteStartPath, absoluteDestinationPath)

	fmt.Printf("%s\n", relativePath)
}
