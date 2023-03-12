package gomux

import (
    "encoding/json"
    "net/http"
)

func SimpleHandler[T any](handler func(T, http.ResponseWriter, *http.Request) any) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var t T
        if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
            http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
            return
        }

        result := handler(t, w, r)

        if result != nil {
            b, err := json.Marshal(result)
            if err != nil {
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
                return
            }

            w.Header().Set("Content-Type", "application/json")
            _, _ = w.Write(b)
        }
    }
}
