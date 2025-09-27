//go:build ignore

package main

import (
	"log"
	"net/http"

	"github.com/SirWaithaka/gorequest"
	"github.com/SirWaithaka/gorequest/corehooks"

	daraja2 "github.com/SirWaithaka/payments/daraja"
)

// this example demonstrates an example of calling the authentication api on daraja

func main() {
	key := "fake_key"
	secret := "fake_secret"

	hooks := gorequest.Hooks{}
	// needs a hook to set basic auth headers
	hooks.Build.PushFrontHook(corehooks.SetBasicAuth(key, secret))
	// needs a hook to unmarshal the response
	// the daraja auth endpoint returns a json response
	hooks.Unmarshal.PushBackHook(daraja2.ResponseDecoder)

	// configure the request method and path
	op := gorequest.Operation{Method: http.MethodGet, Path: "/oauth/v1/generate?grant_type=client_credentials"}
	// model where the response will be unmarshalled to
	res := &daraja2.ResponseAuthorization{}
	// create an instance of request
	req := gorequest.New(gorequest.Config{Endpoint: daraja2.SandboxUrl}, op, hooks, nil, nil, res)
	// make request
	if err := req.Send(); err != nil {
		log.Println(err)
		return
	}
	// print response
	log.Println(res)
}
