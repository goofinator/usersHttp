package binders

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type key int

const (
	//ID is a key to access id in the context
	ID key = iota
)

// IDBinder puts id into the context
func IDBinder(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sid, ok := vars["id"]
		if !ok {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(sid)
		if err != nil {
			http.Error(w, fmt.Sprintf("bad id: %s", err), http.StatusBadRequest)
			return
		}

		context.Set(r, ID, id)
		h.ServeHTTP(w, r)
	})
}
