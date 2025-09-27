package quikk

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/SirWaithaka/gorequest"
	"github.com/SirWaithaka/gorequest/corehooks"
)

// given a data and secret, signer generates a base64 encoded hmac signature
func signer(date, secret []byte) string {
	data := []byte(fmt.Sprintf("date: %s", date))
	h := hmac.New(sha256.New, secret)
	h.Write(data)
	b64 := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return url.QueryEscape(b64)
}

func DefaultHooks() gorequest.Hooks {
	hooks := corehooks.Default()

	hooks.Build.PushBackHook(corehooks.EncodeRequestBody)
	hooks.Unmarshal.PushBackHook(ResponseDecoder)
	return hooks
}

type Config struct {
	Endpoint string
	Hooks    gorequest.Hooks
	LogLevel gorequest.LogLevel
}

type Client struct {
	endpoint string
	Hooks    gorequest.Hooks
}

func New(cfg Config) Client {
	if cfg.Hooks.IsEmpty() {
		cfg.Hooks = DefaultHooks()
	}

	// add log level to request config
	cfg.Hooks.Build.PushFront(gorequest.WithLogLevel(cfg.LogLevel))
	cfg.Hooks.Build.PushFront(gorequest.WithRequestHeader("accept", "application/vnd.api+json"))
	cfg.Hooks.Build.PushFront(gorequest.WithRequestHeader("content-type", "application/vnd.api+json"))

	return Client{endpoint: cfg.Endpoint, Hooks: cfg.Hooks}
}

func (client Client) VerifyAuth(opts ...gorequest.Option) (*gorequest.Request, []byte) {
	op := gorequest.Operation{
		Name:   OperationAuthCheck,
		Method: http.MethodGet,
		Path:   EndpointAuthCheck,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	var output []byte
	req := gorequest.New(cfg, op, client.Hooks, nil, nil, &output)
	req.ApplyOptions(opts...)

	return req, output
}

func (client Client) ChargeRequest(input RequestCharge, ref string, opts ...gorequest.Option) (*gorequest.Request, *ResponseDefault) {
	op := gorequest.Operation{
		Name:   OperationCharge,
		Method: http.MethodPost,
		Path:   EndpointCharge,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// append to request options
	//opts = append(opts, gorequest.WithRequestHeader("Content-Type", "application/json"))

	// build actual payload
	payload := RequestDefault[RequestCharge]{
		Data: Data[RequestCharge]{
			ID:         ref,
			Type:       "charge",
			Attributes: input,
		},
	}

	output := ResponseDefault{}
	req := gorequest.New(cfg, op, client.Hooks, nil, payload, &output)
	req.ApplyOptions(opts...)

	return req, &output
}

func (client Client) PayoutRequest(input RequestPayout, ref string, opts ...gorequest.Option) (*gorequest.Request, *ResponseDefault) {
	op := gorequest.Operation{
		Name:   OperationPayout,
		Method: http.MethodPost,
		Path:   EndpointPayout,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// build actual payload
	payload := RequestDefault[RequestPayout]{
		Data: Data[RequestPayout]{
			ID:         ref,
			Type:       "payout",
			Attributes: input,
		},
	}

	output := ResponseDefault{}
	req := gorequest.New(cfg, op, client.Hooks, nil, payload, &output)
	req.ApplyOptions(opts...)

	return req, &output
}

func (client Client) TransferRequest(input RequestTransfer, ref string, opts ...gorequest.Option) (*gorequest.Request, *ResponseDefault) {
	op := gorequest.Operation{
		Name:   OperationTransfer,
		Method: http.MethodPost,
		Path:   EndpointTransfer,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// build actual payload
	payload := RequestDefault[RequestTransfer]{
		Data: Data[RequestTransfer]{
			ID:         ref,
			Type:       "transfer",
			Attributes: input,
		},
	}

	output := ResponseDefault{}
	req := gorequest.New(cfg, op, client.Hooks, nil, payload, &output)
	req.ApplyOptions(opts...)

	return req, &output
}

func (client Client) BalanceRequest(input RequestAccountBalance, ref string, opts ...gorequest.Option) (*gorequest.Request, *ResponseDefault) {
	op := gorequest.Operation{
		Name:   OperationBalance,
		Method: http.MethodPost,
		Path:   EndpointBalance,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// build actual payload
	payload := RequestDefault[RequestAccountBalance]{
		Data: Data[RequestAccountBalance]{
			ID:         ref,
			Type:       "search",
			Attributes: input,
		},
	}

	output := ResponseDefault{}
	req := gorequest.New(cfg, op, client.Hooks, nil, payload, &output)
	req.ApplyOptions(opts...)

	return req, &output
}

func (client Client) TransactionSearchRequest(input RequestTransactionStatus, ref string, opts ...gorequest.Option) (*gorequest.Request, *ResponseDefault) {
	op := gorequest.Operation{
		Name:   OperationTransactionSearch,
		Method: http.MethodPost,
		Path:   EndpointTransactionSearch,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// build actual payload
	payload := RequestDefault[RequestTransactionStatus]{
		Data: Data[RequestTransactionStatus]{
			ID:         ref,
			Type:       "search",
			Attributes: input,
		},
	}

	output := ResponseDefault{}
	req := gorequest.New(cfg, op, client.Hooks, nil, payload, &output)
	req.ApplyOptions(opts...)

	return req, &output
}

func (client Client) TransactionSearch(ctx context.Context, input RequestTransactionStatus, ref string) (ResponseDefault, error) {
	req, out := client.TransactionSearchRequest(input, ref)
	req.WithContext(ctx)

	if err := req.Send(); err != nil {
		return ResponseDefault{}, err
	}

	return *out, nil
}
