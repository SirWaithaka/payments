package tanda

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/SirWaithaka/gorequest"
	"github.com/SirWaithaka/gorequest/corehooks"
)

func DefaultHooks() gorequest.Hooks {
	// create default hooks
	hooks := corehooks.Default()

	// create client with default timeout of 5 seconds
	client := &http.Client{Timeout: 5 * time.Second}
	hooks.Build.PushFront(gorequest.WithHTTPClient(client))
	hooks.Build.PushBackHook(corehooks.EncodeRequestBody)
	hooks.Unmarshal.PushBackHook(ResponseDecoder)
	return hooks

}

type Config struct {
	Endpoint string
	Hooks    gorequest.Hooks
	LogLevel gorequest.LogLevel
}

func New(cfg Config) Client {
	if cfg.Hooks.IsEmpty() {
		cfg.Hooks = DefaultHooks()
	}

	// add log level to request config
	cfg.Hooks.Build.PushFront(gorequest.WithLogLevel(cfg.LogLevel))

	return Client{endpoint: cfg.Endpoint, Hooks: cfg.Hooks}
}

type Client struct {
	endpoint string
	Hooks    gorequest.Hooks
}

func (client Client) AuthenticationRequest(clientID, secret string) (*gorequest.Request, *ResponseAuthentication) {
	op := gorequest.Operation{
		Name:   OperationAuthenticate,
		Method: http.MethodPost,
		Path:   EndpointAuthentication,
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", secret)

	// create a client with a 30-second timeout
	cl := &http.Client{Timeout: time.Second * 30}
	cfg := gorequest.Config{HTTPClient: cl, Endpoint: client.endpoint}

	// default hooks
	hooks := corehooks.Default()
	hooks.Build.PushFront(gorequest.WithRequestHeader("Content-Type", "application/x-www-form-urlencoded"))
	hooks.Send.PushFrontHook(corehooks.LogHTTPRequest)
	hooks.Unmarshal.PushBackHook(ResponseDecoder)

	input := strings.NewReader(data.Encode())
	output := &ResponseAuthentication{}
	req := gorequest.New(cfg, op, hooks, nil, input, output)

	return req, output
}

func (client Client) PaymentRequest(orgID string, payload RequestPayment, opts ...gorequest.Option) (*gorequest.Request, *ResponsePayment) {
	op := gorequest.Operation{
		Name:   OperationPayment,
		Method: http.MethodPost,
		Path:   fmt.Sprintf(EndpointPayments, orgID),
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	output := &ResponsePayment{}
	req := gorequest.New(cfg, op, client.Hooks, nil, payload, output)
	req.ApplyOptions(opts...)

	return req, output

}

func (client Client) TransactionStatusRequest(orgID, trackingID, shortCode string, opts ...gorequest.Option) (*gorequest.Request, *ResponseTransactionStatus) {
	op := gorequest.Operation{
		Name:   OperationTransactionStatus,
		Method: http.MethodGet,
		Path:   fmt.Sprintf(EndpointTransactionStatus, orgID, trackingID),
	}

	// add short code as a query param to the request path
	uParams := url.Values{}
	uParams.Set("shortCode", shortCode)
	op.Path = op.Path + "?" + uParams.Encode()

	cfg := gorequest.Config{Endpoint: client.endpoint}

	output := &ResponseTransactionStatus{}
	req := gorequest.New(cfg, op, client.Hooks, nil, nil, output)
	req.ApplyOptions(opts...)

	return req, output
}
