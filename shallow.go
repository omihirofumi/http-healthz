package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func healthShallowHandler(w http.ResponseWriter, r *http.Request) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "shallow-")
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer os.Remove(tmpFile.Name())

	text := []byte("Check.")
	if _, err := tmpFile.Write(text); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	if err := tmpFile.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
}
