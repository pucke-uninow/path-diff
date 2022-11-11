package internal

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

func toAbsolutePath(relativePath string) (string, error) {
	return filepath.Abs(relativePath)
}

func Diff(cmd *cobra.Command, args []string) {
	startPath := cmd.Flag("start").Value.String()
	destinationPath := cmd.Flag("destination").Value.String()

	if !(len(startPath) > 0 && len(destinationPath) > 0) {
		if len(args) < 2 {
			log.Println("Please provide start and destination path.")
			return
		}

		startPath = args[0]
		destinationPath = args[1]
	}

	absoluteStartPath, err := toAbsolutePath(startPath)
	if err != nil {
		log.Println(err.Error())
		return
	}

	absoluteDestinationPath, err := toAbsolutePath(destinationPath)
	if err != nil {
		log.Print(err.Error())
		return
	}

	relativePath, err := filepath.Rel(absoluteStartPath, absoluteDestinationPath)

	fmt.Printf("%s\n", relativePath)
}
