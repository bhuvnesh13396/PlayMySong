package kit

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type jsonLogger struct {
	enc *json.Encoder
}

func NewJSONLogger(w io.Writer) Logger {
	// Use synchronized writer provided by log.Logger
	sw := log.New(w, "", 0).Writer()
	enc := json.NewEncoder(sw)
	enc.SetEscapeHTML(false)
	return &jsonLogger{
		enc: enc,
	}
}

func (l *jsonLogger) Log(keyvals ...interface{}) error {
	n := len(keyvals)
	if n%2 != 0 {
		keyvals = append(keyvals, "(MISSING)")
	}

	m := make(map[string]interface{}, n/2)
	for i := 0; i < len(keyvals); i += 2 {
		k, ok := keyvals[i].(string)
		if !ok {
			k = "(UNKNOWN)"
		}
		m[k] = fmt.Sprint(keyvals[i+1])
	}

	return l.enc.Encode(m)
}
