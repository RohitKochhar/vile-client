package actions

import (
	"fmt"
	"rohitsingh/vile/vile/encryption"

	"github.com/spf13/viper"
)

// Rename gets the original key value, stores it in a
// new key and deletes the original
func Rename(args []string, v *viper.Viper) error {
	// Check that we have been given two arguments
	if len(args) != 2 {
		return fmt.Errorf("error: Rename takes two arguments (vile rename {OLD_KEY} {NEW_KEY})")
	}
	// Get the value associated with the provided key
	encryptedVal, err := Get([]string{args[0]}, viper.GetViper())
	if err != nil {
		return err
	}
	// Decrypt the value to store again
	secretKey := v.Get("secretKey")
	decryptedVal, err := encryption.Decrypt(encryptedVal, secretKey.(string))
	if err != nil {
		return err
	}
	// Store the value under a new name
	if err := Add([]string{args[1], string(decryptedVal)}, v); err != nil {
		return err
	}
	// Delete the old key
	if err := Delete([]string{args[0]}, v); err != nil {
		return err
	}
	return nil
}
