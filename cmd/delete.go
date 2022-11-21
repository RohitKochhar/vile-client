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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete {KEY}",
	Short: "Removes a value from remote vile server",
	Long: `
Removes a value from remote vile server

Removes the stored value assocaited with the provided key from the the remote vile server
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			// ToDo: Add support for multiple values
			return fmt.Errorf("cannot get multiple values")
		}
		if len(args) == 0 {
			return fmt.Errorf("no key provided")
		}
		return actions.Delete(args[0], viper.GetViper())
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
