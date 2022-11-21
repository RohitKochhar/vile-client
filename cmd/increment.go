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
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// incrementCmd represents the increment command
var incrementCmd = &cobra.Command{
	Use:   "increment {KEY} {AMOUNT}",
	Short: "Increment adds the provided amount to a stored value",
	Long: `
Increment adds the provided amount to a stored value

Increment is used to increase a value (found by key) by a specified amount.

Increment will return an error if the stored value is not an integer or float
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check that we have been given two arguments
		if len(args) != 2 {
			return fmt.Errorf("error: increment takes two arguments (vile rename {KEY} {AMOUNT})")
		}
		amount, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			return fmt.Errorf("error: invalid increment value specified: %s", args[1])
		}
		return actions.Increment(args[0], amount, viper.GetViper())
	},
}

func init() {
	rootCmd.AddCommand(incrementCmd)
	incrementCmd.DisableFlagParsing = true
}
