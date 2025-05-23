package main

import (
    "net/http"
    "log/slog"
    "encoding/json"
    "os"
)
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    switch r.URL.Path  {
    case "/add":
        s.AddHandler(w,r)
    case "/subtract":
        s.SubtractHandler(w,r)
    case "/divide":
        s.DivisionHandler(w,r)
    case "/multiply":
        s.MultiplicationHandler(w,r)
    case "/sum":
        s.SumHandler(w,r)
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
    w.WriteHeader(http.StatusOK)
    w.Write(jsonBytes)
    logger.Info(string(jsonBytes),
        "path",r.URL.Path,
        "ip", r.RemoteAddr,
    )
}
func (s server) SubtractHandler(w http.ResponseWriter,r *http.Request)  {
    w.Header().Set("Content Type", "application/json")
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
        json.NewEncoder(w).Encode(map[string]string  {
            "error": "Invalid JSON input",
        })
        return
    }

    solution := SingleValue{}
    solution.Value1 = calcReq.Value1 - calcReq.Value2

    jsonBytes, err := json.Marshal(solution)
    if err != nil  {
        logger.Error("Failed encoding",
            "path", r.URL.Path,
            "ip", r.RemoteAddr,
            "error", err.Error(),
        )
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string  {
            "error": "Failed encoding",
        })
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(jsonBytes)
    logger.Info(string(jsonBytes),
        "path",r.URL.Path,
        "ip", r.RemoteAddr,
    )
}
func (s server) DivisionHandler(w http.ResponseWriter,r *http.Request)  {
    w.Header().Set("Content Type", "application/json")
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
        json.NewEncoder(w).Encode(map[string]string  {
            "error": "Invalid JSON input",
        })
        return
    }
    if calcReq.Value2 == 0  {
        logger.Error("Invalid Divisor",
            "path", r.URL.Path,
            "ip", r.RemoteAddr,
        )
        w.WriteHeader(http.StatusBadRequest) 
        json.NewEncoder(w).Encode(map[string]string  {
            "error": "Cannot divide by 0",
        })
        return
    }
    solution := SingleValue{}
    solution.Value1 = calcReq.Value1/calcReq.Value2

    jsonBytes, err := json.Marshal(solution)
    if err != nil  {
        logger.Error("Failed encoding",
            "path", r.URL.Path,
            "ip", r.RemoteAddr,
            "error", err.Error(),
        )
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string  {
            "error": "Failed encoding",
        })
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(jsonBytes)
    logger.Info(string(jsonBytes),
        "path",r.URL.Path,
        "ip", r.RemoteAddr,
    )
}
func (s server) MultiplicationHandler(w http.ResponseWriter,r *http.Request)  {
    w.Header().Set("Content Type", "application/json")
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
        json.NewEncoder(w).Encode(map[string]string  {
            "error": "Invalid JSON input",
        })
        return
    }

    solution := SingleValue{}
    solution.Value1 = calcReq.Value1 * calcReq.Value2

    jsonBytes, err := json.Marshal(solution)
    if err != nil  {
        logger.Error("Failed encoding",
            "path", r.URL.Path,
            "ip", r.RemoteAddr,
            "error", err.Error(),
        )
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string  {
            "error": "Failed encoding",
        })
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(jsonBytes)
    logger.Info(string(jsonBytes),
        "path",r.URL.Path,
        "ip", r.RemoteAddr,
    )
}
func (s server) SumHandler(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Content-Type", "application/json")
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    defer r.Body.Close()
    
    jsonDecoder := json.NewDecoder(r.Body)

    calcReq := ArrayValue{}
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
    for _,v := range calcReq.Value1  {
        solution.Value1+= v
    }

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
    w.WriteHeader(http.StatusOK)
    w.Write(jsonBytes)
    logger.Info(string(jsonBytes),
        "path",r.URL.Path,
        "ip", r.RemoteAddr,
    )
}

