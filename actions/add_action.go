package actions

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"rohitsingh/vile/vile/encryption"

	"github.com/spf13/viper"
)

func Add(args []string, v *viper.Viper) error {
	// Check that we have been given a valid number of args
	if len(args) > 2 {
		return fmt.Errorf("cannot add multiple key value pairs")
	}
	if len(args) < 2 {
		return fmt.Errorf("no key value pair provided")
	}
	// Send a PUT request to the remote server
	// ToDo: Remove this once server certificate signing is automated
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	path, err := getHttpUrl(args[0], v)
	if err != nil {
		return fmt.Errorf("error while creating HTTP url: %q", err)
	}
	// Check if we want to encrypt the value
	secretKey := v.Get("secretKey")
	var val *bytes.Buffer
	if secretKey == nil {
		val = bytes.NewBuffer([]byte(args[1]))
	} else {
		encryptedValue, err := encryption.Encrypt(args[1], secretKey.(string))
		if err != nil {
			return fmt.Errorf("error while encrypyting value: %q", err)
		}
		val = bytes.NewBuffer([]byte(encryptedValue))
	}
	req, err := http.NewRequest(
		http.MethodPut,
		path,
		val,
	)
	if err != nil {
		return fmt.Errorf("error while defining PUT request: %q", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error while sending PUT request: %q", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error while making PUT request: %s", http.StatusText(resp.StatusCode))
	}
	fmt.Printf("Successfully added %s:\"%s\" to remote vile store\n", args[0], args[1])
	return nil
}
