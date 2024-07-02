package main

import (
    "log"
    "net/http"
    "github.com/r3labs/sse/v2"
)

func main() {
    // Create an SSE server
    stream := sse.New()
    stream.CreateStream("chat")

    // Handle message sending
    http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
        message := r.URL.Query().Get("message")
        stream.Publish("chat", &sse.Event{
            Data: []byte(message),
        })
        w.WriteHeader(http.StatusOK)
    })

    // Serve the SSE stream
    http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
        stream.ServeHTTP(w, r)
    })


    log.Println("Server started at :1395")
    log.Fatal(http.ListenAndServe(":1395", nil))
}