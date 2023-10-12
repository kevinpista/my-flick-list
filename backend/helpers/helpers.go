package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/kevinpista/my-flick-list/backend/services"

)
// create write, read json methods and functions

type Envelope map[string] interface {} // help makes Json pretty

type Message struct {
	InfoLog *log.Logger
	ErrorLog *log.Logger
}

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Lshortfile) // tells us date and the file giving us the error

var MessageLogs = &Message{
    InfoLog: infoLog,
    ErrorLog: errorLog,
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
    maxBytes := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body) // Creates new JSON decoder to decode the request body
	err := dec.Decode(data) // Decodes the JSON into the data interface {} we established in services/response.go
	
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{}) // Attempts to decode again into an empty struct
	// Safe keeping code to check for any extra JSON
	if err != nil {
		return errors.New("body must have only a single JSON object")
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t") // serializes Go data structure into a JSON string
	if err != nil {
		return err
	}
	if len(headers) > 0 { // if greater than 0, means we have headers so we want to set those headers
		for key, value := range headers[0] {
			w.Header() [key] = value // set the headers of each in the writer if any
		}
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
	_, err = w.Write(out) // writes the JSON byte array to the body 

	if err != nil{
		return err
	}
	
	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload services.JsonResponse
	payload.Error = true
	payload.Message = err.Error()
	WriteJSON(w, statusCode, payload)
}
