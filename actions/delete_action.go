package actions

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

// Delete takes an array containing a single key

func Delete(key string, v *viper.Viper) error {
	// Send a DELETE request to the path
	// ToDo: Remove this once server certificate signing is automated
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	path, err := getHttpUrl(key, v)
	if err != nil {
		return fmt.Errorf("error while creating HTTP url: %q", err)
	}
	req, err := http.NewRequest(
		http.MethodDelete,
		path,
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
	fmt.Printf("Successfully removed %s from remote vile store\n", key)
	return nil
}
