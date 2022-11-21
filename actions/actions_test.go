package actions

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*
Note that these tests will fail to run unless a vile client is running locally
*/

var ErrNoServer = errors.New("error: Couldn't connect to vile server at https://localhost:8080, is it running?")

// InitViper initializes the configuration for testing
func InitViper() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".vile" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".vile")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// TestCheckAction checks that the vile server is up and running using the check command
func TestCheckAction(t *testing.T) {
	InitViper()
	err := Check(viper.GetViper())
	if err != nil {
		t.Fatalf("%q: %q", ErrNoServer, err)
	}
}

// TestAddGetDeleteGetAction adds values to the server and attempts to get them
// then deletes them and tries to get them expecting the process to fail
func TestAddGetDeleteGetAction(t *testing.T) {
	// Using table-driven testing
	testCases := []struct {
		name          string            // name of test
		keyValuePairs map[string]string // Collection of key value pairs to add
		valType       string            // Type that the stored value should be able to be converted to
	}{
		{
			// TestSimpleAdd adds a single key value pair to the store and checks if it exists
			name:          "TestSimpleAddGet",
			keyValuePairs: map[string]string{"testKey": "testValue"},
		},
		{
			// TestIntegerAdd adds an integer value (string represented) to the store and checks that it can be gotten
			name:          "TestIntegerAddGet",
			keyValuePairs: map[string]string{"testKey": fmt.Sprint(41234134)},
		},
		{
			// TestJSONAdd adds an json value (string represented) to the store and checks that it can be gotten
			name:          "TestIntegerAddGet",
			keyValuePairs: map[string]string{"jsonTestKey": `"objectData":{"name":"objectName",type:"JSON"}`},
		},
		{
			// TestFloatAdd adds a float to the store and checks if it can be received
			name:          "TestFloatAdd",
			keyValuePairs: map[string]string{"floatKey": "413413.2545245234"},
		},
		{
			// TestMultiWordAdd adds a string of a few words and checks if it can be gotten
			name:          "TestMultiWordAdd",
			keyValuePairs: map[string]string{"multiWordKey": "there is a value here"},
		},
		{
			// TestStrangeInput adds a string with escpae characters be gotten
			name:          "TestStrangeInput",
			keyValuePairs: map[string]string{"strangeInputKey": "there\t\n\"\"\"is&(^(&^(^something````here"},
		},
	}
	// Check that the vile server is running
	TestCheckAction(t)
	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Iterate through the map of kv pairs
			for k, v := range tc.keyValuePairs {
				vip := viper.GetViper()
				// Add the pair
				if err := Add(k, v, vip); err != nil {
					t.Fatalf("error while adding key value pair to server: %q", err)
				}
				// Check that we can get the pair back
				val, err := Get(k, vip)
				if err != nil {
					t.Fatalf("error while getting key from server: %q", err)
				}
				// Check that the value is the one that we put up
				if val != v {
					t.Fatalf("incorrect value returned from server, expected %s, got %s", v, val)
				}
				// Delete the kv pair
				if err := Delete(k, vip); err != nil {
					t.Fatalf("error while deleting key from server: %q", err)
				}
				// Get the kv pair, expecting to fail
				emptyVal, err := Get(k, vip)
				if err == nil {
					t.Fatalf("expected to get an error, instead got nil")
				}
				if !errors.Is(err, ErrKeyNotFound) {
					t.Fatalf("expected to not find key, but instead got: %q", err)
				}
				if emptyVal != "" {
					t.Fatalf("expected returned get value to be empty, instead got %s", emptyVal)
				}
			}
		})
	}
}

// TestIntegration Tests all actions in one flow
func TestIntegration(t *testing.T) {
	// Define key value pair
	key := "integrationTest"
	value := "intergration Test Value"
	v := viper.GetViper()
	// Test the connection
	TestCheckAction(t)
	// Add a value to the store
	if err := Add(key, value, v); err != nil {
		t.Fatalf("error while adding value: %q", err)
	}
	// Get the value back
	gotVal, err := Get(key, v)
	if err != nil {
		t.Fatalf("error while getting value: %q", err)
	}
	if gotVal != value {
		t.Fatalf("incorrect value returned, expected %s, got %s", value, gotVal)
	}
	// Rename the value
	newKey := "newIntegrationTest"
	if err := Rename(key, newKey, v); err != nil {
		t.Fatalf("error while renaming value: %q", err)
	}
	// Get the newly renamed value
	renameGotVal, err := Get(newKey, v)
	if err != nil {
		t.Fatalf("error while getting renamed value: %q", err)
	}
	if renameGotVal != value {
		t.Fatalf("incorrect value returned, expected %s, got %s", value, renameGotVal)
	}
	// Check that we cannot get the original
	emptyVal, err := Get(key, v)
	if err == nil {
		t.Fatalf("expected to get an error, instead got nil")
	}
	if !errors.Is(err, ErrKeyNotFound) {
		t.Fatalf("expected to not find key, but instead got: %q", err)
	}
	if emptyVal != "" {
		t.Fatalf("expected returned get value to be empty, instead got %s", emptyVal)
	}
	// Delete the new key
	if err := Delete(newKey, v); err != nil {
		t.Fatalf("error while deleting renamed key: %q", err)
	}
	// Check that we can't get new key
	emptyNewVal, err := Get(newKey, v)
	if err == nil {
		t.Fatalf("expected to get an error, instead got nil")
	}
	if !errors.Is(err, ErrKeyNotFound) {
		t.Fatalf("expected to not find key, but instead got: %q", err)
	}
	if emptyVal != "" {
		t.Fatalf("expected returned get value to be empty, instead got %s", emptyNewVal)
	}
}
