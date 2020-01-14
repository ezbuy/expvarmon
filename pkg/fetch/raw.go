package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"expvar"
)

// FetchRaw fetches var object from expvar directly.
// Since expvar do not support the Range operation, the namesets here is required.
// e.g.: []string{"memstats","cmdline","YOUR_CUSTMIZED_FIELD"}
func FetchRaw(ctx context.Context, namesets []string) (*Expvar, error) {

	vs := make(map[string]json.RawMessage)
	for _, name := range namesets {
		v := expvar.Get(name)
		if v == nil {
			continue
		}
		vs[name] = json.RawMessage(v.String())
	}
	enc, err := json.Marshal(vs)
	if err != nil {
		return nil, err
	}
	return ParseExpvar(bytes.NewBuffer(enc))
}
