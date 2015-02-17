package main

type Capabilities struct {
  Name         string    `json:"name"`
  Version      string    `json:"version"`
  Capabilities []string  `json:"capabilities"`
}
