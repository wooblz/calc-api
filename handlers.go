package main

import (
    "net/http"
    "log"
    "encoding/json"
)
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    switch r.URL.Path  {
    case "/add":
        s.AddHandler(w,r)
    }
}
func (s server) AddHandler(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Content-Type", "application/json")
    defer r.Body.Close()
    
    jsonDecoder := json.NewDecoder(r.Body)

    calcReq := TwoValues{}
    err := jsonDecoder.Decode(&calcReq)
    if err != nil  {
        log.Fatal(err)
    }

    solution := SingleValue{}
    solution.Value1 = calcReq.Value1 + calcReq.Value2

    jsonBytes, err := json.Marshal(solution)
    if err != nil  {
        log.Fatal(err)
    }
    log.Println(string(jsonBytes))
    w.Write(jsonBytes)
}

func main()  {
    log.Fatal(http.ListenAndServe(":8080",server{}))
}

