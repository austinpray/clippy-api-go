package main

import (
	"encoding/json"
	"net/http"
)

func CapabilitiesIndex(w http.ResponseWriter, r *http.Request) {
	capabilities := &Capabilities{
		Name:         "clippy.io Server: Golang Edition",
		Version:      "1.0.0",
		Capabilities: []string{"REST"}}
	json.NewEncoder(w).Encode(capabilities)
}

func SyncIndex(w http.ResponseWriter, r *http.Request) {
	var initiator Client
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&initiator)

	if err != nil {
		panic(err)
	}

	syncRequest := NewSyncRequest(initiator)

	syncRequest.Save()

	json.NewEncoder(w).Encode(syncRequest)
}
