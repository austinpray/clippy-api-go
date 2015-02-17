package main

import (
	"encoding/json"
	"github.com/dchest/uniuri"
	"log"
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

var SyncRequestTTL int = 3600

func NewSyncRequest(initiator Client) *SyncRequest {
	now := time.Now()

	return &SyncRequest{
		Group:     initiator.Group,
		Status:    "pending",
		Code:      uniuri.NewLen(6),
		Initiator: initiator,
		CreatedAt: now,
		ExpiresAt: now.Add(time.Second * 3600)}
}

func (s SyncRequest) Save() {
	conn := pool.Get()
	defer conn.Close()

	save, _ := json.Marshal(s)

	key := "SR-" + s.Code

	if _, err := conn.Do("SET", key, save); err != nil {
		log.Fatal(err)
	}
	if _, err := conn.Do("EXPIRE", key, 3600); err != nil {
		log.Fatal(err)
	}
}

type SyncRequests []SyncRequest
