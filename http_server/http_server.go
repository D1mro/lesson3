package http_server

import (
	"errors"
	"net/http"
)

func StartHTTPServer() error{
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello server!!/n"))
	})
	err:= http.ListenAndServe(":5050", nil)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}