package actions

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func Check(v *viper.Viper) (string, error) {
	// Send a GET request to the path
	// ToDo: Remove this once server certificate signing is automated
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	path := fmt.Sprintf("https://%s:%d/", v.Get("host"), v.Get("port"))
	// Send a get request to the root path
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
			fmt.Printf("Vile server up and running at https://%s:%d\n", v.Get("host"), v.Get("port"))
		}
		if r.StatusCode == http.StatusNotFound {
			fmt.Printf("Could not connect to vile at https://%s:%d\n", v.Get("host"), v.Get("port"))
			return "", fmt.Errorf("could not connect to vile server")
		}
		return string(body), nil
	default:
		return "", fmt.Errorf("unsupported Content-Type: %q", r.Header.Get("Content-Type"))
	}
}
