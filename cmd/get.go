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
	"rohitsingh/vile/vile/actions"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get {KEY}",
	Short: "Gets a value from remote vile server",
	Long: `
Gets a value from remote vile server

Retrieves the stored value associated with the provided key from the remote vile server
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := actions.Get(args, viper.GetViper())
		return err
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
