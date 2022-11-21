package actions

import (
	"github.com/spf13/viper"
)

// Rename gets the original key value, stores it in a
// new key and deletes the original
func Rename(oldKey, newKey string, v *viper.Viper) error {
	// Get the value associated with the provided key
	val, err := Get(oldKey, viper.GetViper())
	if err != nil {
		return err
	}

	if err := Add(newKey, string(val), v); err != nil {
		return err
	}
	// Delete the old key
	if err := Delete(oldKey, v); err != nil {
		return err
	}
	return nil
}
