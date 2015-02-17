package main

type Client struct {
	Id        string  `json:"id"`
	Group     string  `json:"group"`
	Name      string  `json:"name"`
	Peers     Clients `json:"peers"`
	PublicKey string  `json:"publicKey"`
}

type Clients []Client
