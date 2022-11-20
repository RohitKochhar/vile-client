package actions

import (
	"fmt"
	"rohitsingh/vile/vile/encryption"

	"github.com/spf13/viper"
)

// getHttpUrl defines the URL to send requests to
func getHttpUrl(key string, v *viper.Viper) (string, error) {
	// Use configuration specified in yaml
	host := v.Get("host")
	port := v.Get("port")
	secretKey := v.Get("secretKey")
	if secretKey == nil {
		return fmt.Sprintf("https://%s:%d/v1/key/%s", host, port, key), nil
	}
	// Encrypt the key value
	encryptedKey, err := encryption.Encrypt(key, string(secretKey.(string)))
	if err != nil {
		return "", fmt.Errorf("error while encrypting key: %q", err)
	}
	// Return the formatted url
	return fmt.Sprintf("https://%s:%d/v1/key/%s", host, port, encryptedKey), nil
}
