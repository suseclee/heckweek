package main

import (
    "log"
    "net/http"
    "encoding/json"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        names := r.URL.Query()["name"]
        name := "anonymous"
        if len(names) == 1 {
            name = names[0]
        }

        m := map[string]string{"name": name}
        w.Write([]byte("Hello GO\n"))
        enc := json.NewEncoder(w)
        enc.Encode(m)
    })

    err := http.ListenAndServe(":3000", nil)

    if err != nil {
        log.Fatal(err)
    }
}