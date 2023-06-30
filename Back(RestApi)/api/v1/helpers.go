package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// WriteJSON write data to response
func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	// marshal data
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	// add headers to response
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func ReadJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		if errors.As(err, &syntaxError) {
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		} else if errors.As(err, &unmarshalTypeError) {
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", unmarshalTypeError.Offset)

		} else if errors.As(err, &invalidUnmarshalError) {
			panic(err) //If our program reaches app point where it cannot be recovered due to some major errors

		} else if errors.Is(err, io.ErrUnexpectedEOF) {
			return errors.New("body contains badly-formed JSON")

		} else if errors.Is(err, io.EOF) {
			return errors.New("body must not be empty")

		} else {
			return err
		}
	}

	return nil
}
