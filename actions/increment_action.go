package actions

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

// Rename gets the original key value, stores it in a
// new key and deletes the original
func Increment(key string, incVal float64, v *viper.Viper) error {
	// Get the value associated with the provided key
	val, err := Get(key, viper.GetViper())
	if err != nil {
		return err
	}
	// Try to convert the value to a float
	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fmt.Errorf("stored value is not a number: %s", val)
	}
	// Add the new value to the store
	if err := Add(key, fmt.Sprintf("%f", valFloat+incVal), v); err != nil {
		return err
	}
	return nil
}
