package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(x); err != nil {
		return fmt.Errorf("ParseBody: %w", err)
	}
	return nil
}
