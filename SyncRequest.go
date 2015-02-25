package main

import (
	"encoding/json"
	"github.com/dchest/uniuri"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

type SyncRequest struct {
	Group     string    `json:"group,omitempty"`
	Code      string    `json:"code,omitempty"`
	Status    string    `json:"status,omitempty"`
	Initiator Client    `json:"initiator,omitempty"`
	Target    Client    `json:"target,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	ExpiresAt time.Time `json:"expiresAt,omitempty"`
}

var SyncRequestTTL int = 3600

func NewSyncRequestKey(code string) string {
	return "SR-" + code
}

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

func GetSyncRequest(code string) *SyncRequest {
	conn := pool.Get()
	defer conn.Close()

	key := NewSyncRequestKey(code)

	value, err := redis.String(conn.Do("GET", key))

	if err != nil {
		log.Panic(err)
	}

	syncRequest := &SyncRequest{}

	json.Unmarshal([]byte(value), &syncRequest)

	return syncRequest
}

func (s SyncRequest) Save() {
	conn := pool.Get()
	defer conn.Close()

	save, _ := json.Marshal(s)

	key := NewSyncRequestKey(s.Code)

	if _, err := conn.Do("SET", key, save); err != nil {
		log.Panic(err)
	}
	if _, err := conn.Do("EXPIRE", key, 3600); err != nil {
		log.Panic(err)
	}
}

type SyncRequests []SyncRequest
