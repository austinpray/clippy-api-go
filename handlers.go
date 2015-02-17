package main

import (
  "encoding/json"
  "net/http"
  "time"
)


func CapabilitiesIndex(w http.ResponseWriter, r *http.Request) {
  capabilities := &Capabilities{
    Name: "clippy.io Server: Golang Edition",
    Version: "1.0.0",
    Capabilities: []string{"REST"}}
  json.NewEncoder(w).Encode(capabilities);
}

func SyncIndex(w http.ResponseWriter, r *http.Request) {
  var syncRequest SyncRequest
  var initiator Client
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&initiator)

  if err != nil {
    panic(err)
  }

  now := time.Now()

  syncRequest = SyncRequest{
    Group: initiator.Group,
    Status:  "pending",
    Code:  "123456",
    Initiator:  initiator,
    CreatedAt:  now,
    ExpiresAt:  now.Add(time.Second*3600)}

  json.NewEncoder(w).Encode(syncRequest);
}
