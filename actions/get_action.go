package actions

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"rohitsingh/vile/vile/encryption"
	"strings"

	"github.com/spf13/viper"
)

var ErrKeyNotFound = errors.New("error: could not find provided key in vile store")

// Get fetches the value associated with the provided key
// from the store described by the provided viper configuration
// and returns an error if the process could not be completed.
func Get(key string, v *viper.Viper) (string, error) {
	// Send a GET request to the path
	// ToDo: Remove this once server certificate signing is automated
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	path, err := getHttpUrl(key, v)
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
				return string(body), nil
			} else {
				decodedVal, err := encryption.Decrypt(string(body), secretKey.(string))
				if err != nil {
					return "", fmt.Errorf("error while decrypting value: %q", err)
				}
				return string(decodedVal), nil
			}
		}
		if r.StatusCode == http.StatusNotFound {
			fmt.Printf("Key \"%s\" was not found\n", key)
			return "", ErrKeyNotFound
		}
		return "", fmt.Errorf("unexpected error occured while getting value")
	default:
		return "", fmt.Errorf("unsupported Content-Type: %q", r.Header.Get("Content-Type"))
	}
}
