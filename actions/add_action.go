package actions

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"rohitsingh/vile/vile/encryption"

	"github.com/spf13/viper"
)

// Add adds the provided key value pair to a vile server specified from the config
// it takes two strings (key, value), and a viper configuration instance
// and returns an error if the PUT process could not be completed
func Add(key, value string, v *viper.Viper) error {
	// Check the input is value
	// if err := sanitizeInput(key); err != nil {
	// 	return err
	// }
	// Send a PUT request to the remote server
	// ToDo: Remove this once server certificate signing is automated
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	path, err := getHttpUrl(key, v)
	if err != nil {
		return fmt.Errorf("error while creating HTTP url: %q", err)
	}
	// Check if we want to encrypt the value
	secretKey := v.Get("secretKey")
	var val *bytes.Buffer
	if secretKey == nil {
		val = bytes.NewBuffer([]byte(value))
	} else {
		encryptedValue, err := encryption.Encrypt(value, secretKey.(string))
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
	fmt.Printf("Successfully added %s:\"%s\" to remote vile store\n", key, value)
	return nil
}
