package main

import (
    "net/http"
    "log/slog"
    "encoding/json"
    "os"
    "log"
)
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    switch r.URL.Path  {
    case "/add":
        s.AddHandler(w,r)
    }
}
func (s server) AddHandler(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Content-Type", "application/json")
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    defer r.Body.Close()
    
    jsonDecoder := json.NewDecoder(r.Body)

    calcReq := TwoValues{}
    err := jsonDecoder.Decode(&calcReq)
    if err != nil  {
        logger.Error("Failed to decode JSON",
            "path", r.URL.Path,
            "ip", r.RemoteAddr,
            "error", err.Error(),
        )
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string {
            "error":"Invalid JSONInput",
        })
        return
    }

    solution := SingleValue{}
    solution.Value1 = calcReq.Value1 + calcReq.Value2

    jsonBytes, err := json.Marshal(solution)
    if err != nil  {
        logger.Error("Failed to marshal data",
            "path",r.URL.Path,
            "ip", r.RemoteAddr,
            "error", err.Error(),
        )
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string  {
            "error": "Failed to Encode",
        })
        return
    }
    w.Write(jsonBytes)
    logger.Info(string(jsonBytes),
        "path",r.URL.Path,
        "ip", r.RemoteAddr,
    )
}

func main()  {
    log.Fatal(http.ListenAndServe(":8080",server{}))
}

