/*
Copyright © 2022 Philipp Ucke <philippucke@googlemail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"qoubo.tech/path-diff/internal"
)

var (
	startPath       string
	destinationPath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "path-diff",
	Short: "Generate a relative path from start to destination path.",
	Long:  "Generate a relative path from start to destination path.",
	Example: `path-diff /home/user/Downloads /home/user/Documents
path-diff -s /home/user/Downloads -d /home/user/Documents`,
	Run: internal.Diff,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Flags().StringVarP(&startPath, "start", "s", "", "Start path")
	rootCmd.Flags().StringVarP(&destinationPath, "destination", "d", "", "Destination path")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
