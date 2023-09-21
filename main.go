package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	Version    string = "development"
	CommitHash string = "unknown"
	BuildTime  string = "unknown"
)

func main() {

	slog.Info("liapi", "version", Version, "commit", CommitHash, "buildTime", BuildTime)

	// Define our routes
	http.HandleFunc("/liatrio", LiatrioHandler)
	http.HandleFunc("/ping", PingHandler)

	// Look for bind address in environment
	address, ok := os.LookupEnv("LIAPI_ADDRESS")
	if !ok {
		address = ":8080"
	}

	// Start the server
	slog.Info("Starting the server", "address", address)
	if err := http.ListenAndServe(address, Log(http.DefaultServeMux)); err != nil {
		slog.Error("Listen and Serve within main", "error", err.Error())
	}
}

// Log is a middleware for the http.Handler
// This wraps the DefaultServeMux (other any other handler) to include structured logging of all requests
// to the http implementation
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request received",
			"address", r.RemoteAddr,
			"method", r.Method,
			"host", r.Host,
			"url", r.URL,
		)
		handler.ServeHTTP(w, r)
	})
}

// LiatrioHandler handles requests to the /liatrio endpoint
func LiatrioHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetLiatrio(w, r)
	case http.MethodPost:
		PostLiatrio(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetPing returns the current health status and version of the application
func GetPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Version, CommitHash, BuildTime})
}

// PingResponse struct for representing a response to the Ping endpoint
type PingResponse struct {
	Version    string `json:"version"`
	CommitHash string `json:"commitHash"`
	BuildTime  string `json:"buildTime"`
}

// PingHandler handles requests to the /liatrio endpoint
func PingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetPing(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// LiatrioRequest struct represents a request to the liapi Liatrio REST endpoints
type LiatrioRequest struct {
	Message string `json:"message"`
}

// LiatrioResponse struct for representing a response to the Liatrio REST endpoints
type LiatrioResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetLiatrio returns the static message and current timestamp
func GetLiatrio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LiatrioResponse{"Automate all the things!", strconv.FormatInt(time.Now().Unix(), 10)})
}

// PostLiatrio returns a dynamic message and current timestamp
// Currently, this function simply mirrors back the message posted along with the current timestamp
// This was added simply to show that dynamic messages and other http methods can be supported in this manner.
func PostLiatrio(w http.ResponseWriter, r *http.Request) {
	var newLiatrioRequest LiatrioRequest
	if err := json.NewDecoder(r.Body).Decode(&newLiatrioRequest); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		slog.Error("Parsing Liatrio POST request", "error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(LiatrioResponse{newLiatrioRequest.Message, strconv.FormatInt(time.Now().Unix(), 10)})
}
