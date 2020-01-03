package fetch

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/antonholmquist/jason"
)

func getBasicAuthEnv() (user, password string) {
	return os.Getenv("HTTP_USER"), os.Getenv("HTTP_PASSWORD")
}

type HTTPClientOption func(*http.Client)

var defaultClient = &http.Client{Timeout: time.Second}

// FetchExpvar fetches expvar by http for the given addr (host:port)
func FetchExpvar(u url.URL, opts ...HTTPClientOption) (*Expvar, error) {
	e := &Expvar{&jason.Object{}}
	client := defaultClient
	for _, opt := range opts {
		opt(client)
	}
	// it works but it seems really wierd.
	// MAYBE we can refactor it
	// error will be ignored here due to the bad format of the URL
	req, _ := http.NewRequest("GET", "localhost", nil)
	req.URL = &u
	req.Host = u.Host

	if user, pass := getBasicAuthEnv(); user != "" && pass != "" {
		req.SetBasicAuth(user, pass)
	}
	resp, err := client.Do(req)
	if err != nil {
		return e, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return e, errors.New("Vars not found. Did you import expvars?")
	}

	expvar, err := ParseExpvar(resp.Body)
	if err != nil {
		return e, err
	}
	e = expvar
	return e, nil
}
