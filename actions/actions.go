package actions

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"rohitsingh/vile/vile/encryption"
	"strings"

	"github.com/spf13/viper"
)

// getHttpUrl defines the URL to send requests to
func getHttpUrl(args []string, v *viper.Viper) (string, error) {
	// Use configuration specified in yaml
	host := v.Get("host")
	port := v.Get("port")
	secretKey := v.Get("secretKey")
	if secretKey == nil {
		return fmt.Sprintf("http://%s:%d/v1/key/%s", host, port, args[0]), nil
	}
	// Encrypt the key value
	encryptedKey, err := encryption.Encrypt(args[0], string(secretKey.(string)))
	if err != nil {
		return "", fmt.Errorf("error while encrypting key: %q", err)
	}
	// Return the formatted url
	return fmt.Sprintf("http://%s:%d/v1/key/%s", host, port, encryptedKey), nil
}

func Add(args []string, v *viper.Viper) error {
	// Check that we have been given a valid number of args
	if len(args) > 2 {
		return fmt.Errorf("cannot add multiple key value pairs")
	}
	if len(args) < 2 {
		return fmt.Errorf("no key value pair provided")
	}
	// Send a PUT request to the remote server
	path, err := getHttpUrl(args, v)
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

func Get(args []string, v *viper.Viper) error {
	if len(args) > 1 {
		// ToDo: Add support for multiple values
		return fmt.Errorf("cannot get multiple values")
	}
	if len(args) == 0 {
		return fmt.Errorf("no key provided")
	}
	// Send a GET request to the path
	path, err := getHttpUrl(args, v)
	if err != nil {
		return fmt.Errorf("error while creating HTTP url: %q", err)
	}
	r, err := http.Get(path)
	if err != nil {
		return err
	}
	switch {
	case strings.Contains(r.Header.Get("Content-Type"), "text/plain"):
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		if r.StatusCode == http.StatusOK {
			// Check if we need to decrypt the value
			secretKey := v.Get("secretKey")
			if secretKey == nil {
				print(string(body))
				return nil
			} else {
				decodedVal, err := encryption.Decrypt(string(body), secretKey.(string))
				if err != nil {
					return fmt.Errorf("error while decrypting value: %q", err)
				}
				fmt.Println(string(decodedVal))
			}
		}
		if r.StatusCode == http.StatusNotFound {
			fmt.Printf("Key \"%s\" was not found\n", args[0])
			return nil
		}

	default:
		return fmt.Errorf("unsupported Content-Type: %q", r.Header.Get("Content-Type"))
	}
	return nil
}

func Delete(args []string, v *viper.Viper) error {
	if len(args) > 1 {
		// ToDo: Add support for multiple values
		return fmt.Errorf("cannot get multiple values")
	}
	if len(args) == 0 {
		return fmt.Errorf("no key provided")
	}
	// Send a DELETE request to the path
	path, err := getHttpUrl(args, v)
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
	fmt.Printf("Successfully removed %s from remote vile store\n", args[0])
	return nil
}
