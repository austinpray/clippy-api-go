package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/mholt/binding"
	"net/http"
)

type Client struct {
	Id        string  `json:"id",omitempty`
	Group     string  `json:"group,omitempty"`
	Name      string  `json:"name,omitempty"`
	Peers     Clients `json:"peers,omitempty"`
	PublicKey string  `json:"publicKey,omitempty"`
}

type Clients []Client

func (client *Client) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&client.Id: binding.Field{
			Form:     "id",
			Required: true,
		},
		&client.Group: binding.Field{
			Form:     "group",
			Required: true,
		},
		&client.Name: binding.Field{
			Form:     "name",
			Required: true,
		},
		&client.Peers: "peers",
		&client.PublicKey: binding.Field{
			Form:     "publicKey",
			Required: true,
		},
	}
}

func (client Client) Validate(req *http.Request, errs binding.Errors) binding.Errors {

	if !govalidator.IsUUIDv4(client.Id) {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"id"},
			Classification: "ContentTypeError",
			Message:        "`id` field must be a UUIDv4",
		})
	}

	if !govalidator.IsUUIDv4(client.Group) {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"group"},
			Classification: "ContentTypeError",
			Message:        "`group` field must be a UUIDv4",
		})
	}

	return errs
}
