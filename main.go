package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

// JSON-RPC types
type JSONRPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      interface{}   `json:"id"`
}

type JSONRPCResponse struct {
	JSONRPC string        `json:"jsonrpc"`
	Result  interface{}   `json:"result,omitempty"`
	Error   *JSONRPCError `json:"error,omitempty"`
	ID      interface{}   `json:"id"`
}

type JSONRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ThrylosLightClient represents a lightweight node client for the Thrylos blockchain
type ThrylosLightClient struct {
	address    string
	seedNodes  []string
	httpClient *http.Client
	mu         sync.RWMutex
}

// NewThrylosLightClient creates a new Thrylos light client
func NewThrylosLightClient(address string, seedNodes []string) *ThrylosLightClient {
	return &ThrylosLightClient{
		address:    address,
		seedNodes:  seedNodes,
		httpClient: &http.Client{Timeout: time.Second * 10},
	}
}

// Start initializes and starts the Thrylos light client
func (c *ThrylosLightClient) Start() error {
	http.HandleFunc("/", c.handleJSONRPC)

	go c.discoverPeers()

	log.Printf("Thrylos Light Client starting on %s\n", c.address)
	log.Printf("Connected to Thrylos network via seed nodes: %v\n", c.seedNodes)
	return http.ListenAndServe(c.address, nil)
}

// handleJSONRPC handles incoming Thrylos JSON-RPC requests
func (c *ThrylosLightClient) handleJSONRPC(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req JSONRPCRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSONRPCError(w, &JSONRPCError{
			Code:    -32700,
			Message: "Parse error",
		}, nil)
		return
	}

	response, err := c.forwardRequest(req)
	if err != nil {
		sendJSONRPCError(w, &JSONRPCError{
			Code:    -32603,
			Message: err.Error(),
		}, req.ID)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// forwardRequest forwards JSON-RPC requests to a seed node
func (c *ThrylosLightClient) forwardRequest(req JSONRPCRequest) (*JSONRPCResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		log.Printf("Error marshaling request: %v", err)
		return nil, err
	}

	for _, node := range c.seedNodes {
		url := fmt.Sprintf("http://%s/", node) // Use root path
		log.Printf("Attempting to connect to node: %s", url)

		request, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
		if err != nil {
			log.Printf("Error creating request for %s: %v", url, err)
			continue
		}

		// Add required headers
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Origin", "http://localhost:8545") // Set appropriate origin

		resp, err := c.httpClient.Do(request)
		if err != nil {
			log.Printf("Error connecting to %s: %v", url, err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		log.Printf("Raw response from %s: Status=%d, Body=%s", url, resp.StatusCode, string(body))
		resp.Body.Close()

		// Create a new reader for the body since we've read it already
		var response JSONRPCResponse
		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&response); err != nil {
			log.Printf("Error decoding response from %s: %v\nResponse body: %s", url, err, string(body))
			continue
		}

		return &response, nil
	}

	return nil, fmt.Errorf("failed to forward request to any seed node")
}

// discoverPeers periodically updates the list of peers
func (c *ThrylosLightClient) discoverPeers() {
	ticker := time.NewTicker(time.Minute * 5)
	for range ticker.C {
		for _, node := range c.seedNodes {
			req := JSONRPCRequest{
				JSONRPC: "2.0",
				Method:  "getPeers",
				Params:  []interface{}{},
				ID:      1,
			}

			reqBytes, err := json.Marshal(req)
			if err != nil {
				continue
			}

			resp, err := c.httpClient.Post(
				fmt.Sprintf("http://%s/peers", node),
				"application/json",
				bytes.NewBuffer(reqBytes),
			)
			if err != nil {
				continue
			}
			defer resp.Body.Close()

			var discoveredPeers []string
			if err := json.NewDecoder(resp.Body).Decode(&discoveredPeers); err != nil {
				continue
			}

			c.mu.Lock()
			c.seedNodes = uniqueStrings(append(c.seedNodes, discoveredPeers...))
			c.mu.Unlock()
		}
	}
}

// sendJSONRPCError sends a JSON-RPC error response
func sendJSONRPCError(w http.ResponseWriter, jsonrpcErr *JSONRPCError, id interface{}) {
	response := JSONRPCResponse{
		JSONRPC: "2.0",
		Error:   jsonrpcErr,
		ID:      id,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}

// Utility function to get unique strings
func uniqueStrings(strings []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, str := range strings {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}
	return result
}

func main() {
	// Command line flags
	address := flag.String("address", ":8545", "Address to run the Thrylos light client on")
	seedNode := flag.String("seed", "thrylos-mainnet.example.com:8545", "Initial Thrylos seed node to connect to")
	networkFlag := flag.String("network", "mainnet", "Thrylos network to connect to (mainnet or testnet)")
	flag.Parse()

	// Set up seed nodes based on network
	var seedNodes []string
	if *networkFlag == "testnet" {
		seedNodes = []string{"thrylos-testnet.example.com:8545"}
	} else {
		seedNodes = []string{*seedNode}
	}

	// Create and start client
	client := NewThrylosLightClient(*address, seedNodes)

	log.Printf("Starting Thrylos Light Client")
	log.Printf("Network: %s", *networkFlag)
	log.Printf("Version: 0.1.0")
	log.Fatal(client.Start())
}
