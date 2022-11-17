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
	"bytes"
	"fmt"
	"net/http"

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
		// Use configuration specified in yaml
		host := viper.GetViper().Get("host")
		port := viper.GetViper().Get("port")
		url := fmt.Sprintf("http://%s:%d/v1/key/%s", host, port, args[0])
		req, err := http.NewRequest(
			http.MethodDelete,
			url,
			bytes.NewBuffer([]byte("")),
		)
		if err != nil {
			return fmt.Errorf("error while defining DELETE request: %q", err)
		}
		req.Header.Set("Content-Type", "text/plain")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("error while sending DELETE request: %q", err)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("error while making DELETE request: %s", http.StatusText(resp.StatusCode))
		}
		fmt.Printf("Successfully removed %s from remote vile store\n", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
