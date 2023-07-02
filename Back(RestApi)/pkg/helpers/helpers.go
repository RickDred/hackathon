package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"hackathon/pkg/validator"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func ReadString(qs url.Values, key string, defaultValue string) string {
	if qs.Has(key) {
		return qs.Get(key)
	}
	return defaultValue
}

func ReadInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	if !qs.Has(key) {
		return defaultValue
	}
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		v.AddError(key, "must be an integer value")
		return defaultValue
	}
	return i
}

func ReadIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	fmt.Println(params)
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

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
