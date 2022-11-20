package actions

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"rohitsingh/vile/vile/encryption"
	"strings"

	"github.com/spf13/viper"
)

func Get(args []string, v *viper.Viper) (string, error) {
	if len(args) > 1 {
		// ToDo: Add support for multiple values
		return "", fmt.Errorf("cannot get multiple values")
	}
	if len(args) == 0 {
		return "", fmt.Errorf("no key provided")
	}
	// Send a GET request to the path
	// ToDo: Remove this once server certificate signing is automated
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	path, err := getHttpUrl(args[0], v)
	if err != nil {
		return "", fmt.Errorf("error while creating HTTP url: %q", err)
	}
	r, err := http.Get(path)
	if err != nil {
		return "", err
	}
	switch {
	case strings.Contains(r.Header.Get("Content-Type"), "text/plain"):
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return "", err
		}
		if r.StatusCode == http.StatusOK {
			// Check if we need to decrypt the value
			secretKey := v.Get("secretKey")
			if secretKey == nil {
				print(string(body))
				return "", nil
			} else {
				decodedVal, err := encryption.Decrypt(string(body), secretKey.(string))
				if err != nil {
					return "", fmt.Errorf("error while decrypting value: %q", err)
				}
				fmt.Println(string(decodedVal))
			}
		}
		if r.StatusCode == http.StatusNotFound {
			fmt.Printf("Key \"%s\" was not found\n", args[0])
			return "", nil
		}
		return string(body), nil
	default:
		return "", fmt.Errorf("unsupported Content-Type: %q", r.Header.Get("Content-Type"))
	}
	return "", nil
}
