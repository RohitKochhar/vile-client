/*
Copyright © 2022 Rohit Singh

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
	"fmt"
	"rohitsingh/vile/vile/actions"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename {OLD_KEY} {NEW_KEY}",
	Short: "Rename changes the name of the key used to store a value",
	Long: `
Rename changes the name of the key used to store a value

Rename is equivalent to: 
	vile get {OLD_KEY}
	vile add {NEW_KEY} {GET_RESULT}
	vile delete {OLD_KEY}
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check that we have been given two arguments
		if len(args) != 2 {
			return fmt.Errorf("error: Rename takes two arguments (vile rename {OLD_KEY} {NEW_KEY})")
		}
		return actions.Rename(args[0], args[1], viper.GetViper())
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
