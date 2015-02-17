package main

import (
  "time"
)

type SyncRequest struct {
  Group     string    `json:"group"`
  Code      string    `json:"code"`
  Status    string    `json:"status"`
  Initiator Client    `json:"initiator"`
  Target    Client    `json:"target"`
  CreatedAt time.Time `json:"createdAt"`
  ExpiresAt time.Time `json:"expiresAt"`
}

type SyncRequests []SyncRequest
