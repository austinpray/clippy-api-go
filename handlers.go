package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
	"github.com/unrolled/render"
	"log"
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
	initiator := new(Client)

	errs := binding.Bind(r, initiator)
	if errs.Handle(w) {
		return
	}

	syncRequest := NewSyncRequest(*initiator)

	syncRequest.Save()

	resp := render.New()
	resp.JSON(w, http.StatusOK, syncRequest)
}

func SyncHandler(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	syncRequest := GetSyncRequest(code)

	if r.Method == "POST" {
		var target Client
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&target)

		if err != nil {
			log.Panic(err)
		}

		syncRequest.Status = "accepted"
		syncRequest.Target = target

		syncRequest.Save()
	}

	json.NewEncoder(w).Encode(syncRequest)
}
