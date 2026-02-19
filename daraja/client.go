package daraja

import (
	"context"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/SirWaithaka/gorequest"
	"github.com/SirWaithaka/gorequest/corehooks"
)

type AuthenticationRequestFunc func() (*gorequest.Request, *ResponseAuthorization)

type Config struct {
	Endpoint string
	Hooks    gorequest.Hooks
	LogLevel gorequest.LogLevel
}

func DefaultHooks() gorequest.Hooks {
	// create default hooks
	hooks := corehooks.Default()

	// create client with default timeout of 5 seconds
	client := &http.Client{Timeout: 5 * time.Second}
	hooks.Build.PushBackHook(HTTPClient(client))
	hooks.Build.PushBackHook(corehooks.EncodeRequestBody)
	hooks.Unmarshal.PushBackHook(ResponseDecoder)
	return hooks

}

func PasswordEncode(shortcode, passphrase, timestamp string) string {
	return base64.StdEncoding.EncodeToString([]byte(shortcode + passphrase + timestamp))
}

// Client provides the API operation methods for making requests
// to MPESA daraja service.
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

	return Client{endpoint: cfg.Endpoint, Hooks: cfg.Hooks}
}

func (client Client) AuthenticationRequest(key, secret string) AuthenticationRequestFunc {
	return func() (*gorequest.Request, *ResponseAuthorization) {
		op := gorequest.Operation{
			Name:   "Authenticate",
			Method: http.MethodGet,
			Path:   EndpointAuthentication + "?grant_type=client_credentials",
		}

		// create a client with a 40-second timeout
		cl := &http.Client{Timeout: time.Second * 40}
		cfg := gorequest.Config{HTTPClient: cl, Endpoint: client.endpoint}

		// default hooks
		hooks := corehooks.Default()
		hooks.Build.PushBackHook(corehooks.SetBasicAuth(key, secret))
		hooks.Send.PushFrontHook(corehooks.LogHTTPRequest)
		hooks.Unmarshal.PushBackHook(ResponseDecoder)

		output := &ResponseAuthorization{}
		req := gorequest.New(cfg, op, hooks, nil, nil, output)

		return req, output
	}
}

func (client Client) C2BExpressRequest(input RequestC2BExpress, opts ...gorequest.Option) (*gorequest.Request, *ResponseC2BExpress) {
	op := gorequest.Operation{
		Name:   OperationC2BExpress,
		Method: http.MethodPost,
		Path:   EndpointC2bExpress,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// append to request options
	opts = append(opts, gorequest.WithRequestHeader("Content-Type", "application/json"))

	output := ResponseC2BExpress{}
	req := gorequest.New(cfg, op, client.Hooks, nil, input, &output)
	req.ApplyOptions(opts...)

	return req, &output
}

func (client Client) C2BExpress(ctx context.Context, payload RequestC2BExpress) (ResponseC2BExpress, error) {
	req, out := client.C2BExpressRequest(payload)
	req.WithContext(ctx)

	if err := req.Send(); err != nil {
		return ResponseC2BExpress{}, err
	}

	return *out, nil
}

func (client Client) C2BQueryRequest(input RequestC2BExpressQuery, opts ...gorequest.Option) (*gorequest.Request, *ResponseC2BExpressQuery) {
	op := gorequest.Operation{
		Name:   OperationC2BQuery,
		Method: http.MethodPost,
		Path:   EndpointC2bExpressQuery,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// append to request options
	opts = append(opts, gorequest.WithRequestHeader("Content-Type", "application/json"))

	output := &ResponseC2BExpressQuery{}
	req := gorequest.New(cfg, op, client.Hooks, nil, input, output)
	req.ApplyOptions(opts...)

	return req, output
}

func (client Client) C2BQuery(ctx context.Context, payload RequestC2BExpressQuery) (ResponseC2BExpressQuery, error) {
	req, out := client.C2BQueryRequest(payload)
	req.WithContext(ctx)

	if err := req.Send(); err != nil {
		return ResponseC2BExpressQuery{}, err
	}

	return *out, nil
}

func (client Client) ReversalRequest(input RequestReversal, opts ...gorequest.Option) (*gorequest.Request, *ResponseReversal) {
	op := gorequest.Operation{
		Name:   OperationReversal,
		Method: http.MethodPost,
		Path:   EndpointReversal,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// append to request options
	opts = append(opts, gorequest.WithRequestHeader("Content-Type", "application/json"))

	output := &ResponseReversal{}
	req := gorequest.New(cfg, op, client.Hooks, nil, input, output)
	req.ApplyOptions(opts...)

	return req, output
}

func (client Client) Reverse(ctx context.Context, payload RequestReversal) (ResponseReversal, error) {
	req, out := client.ReversalRequest(payload)
	req.WithContext(ctx)

	if err := req.Send(); err != nil {
		return ResponseReversal{}, err
	}

	return *out, nil
}

func (client Client) B2CRequest(input RequestB2C, opts ...gorequest.Option) (*gorequest.Request, *ResponseB2C) {
	op := gorequest.Operation{
		Name:   OperationB2C,
		Method: http.MethodPost,
		Path:   EndpointB2cPayment,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// append to request options
	opts = append(opts, gorequest.WithRequestHeader("Content-Type", "application/json"))

	output := &ResponseB2C{}
	req := gorequest.New(cfg, op, client.Hooks, nil, input, output)
	req.ApplyOptions(opts...)

	return req, output
}

func (client Client) B2C(ctx context.Context, payload RequestB2C) (ResponseB2C, error) {
	req, out := client.B2CRequest(payload)
	req.WithContext(ctx)

	if err := req.Send(); err != nil {
		return ResponseB2C{}, err
	}

	return *out, nil
}

func (client Client) B2BRequest(input RequestB2B, opts ...gorequest.Option) (*gorequest.Request, *ResponseB2B) {
	op := gorequest.Operation{
		Name:   OperationB2B,
		Method: http.MethodPost,
		Path:   EndpointB2bPayment,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// append to request options
	opts = append(opts, gorequest.WithRequestHeader("Content-Type", "application/json"))

	output := &ResponseB2B{}
	req := gorequest.New(cfg, op, client.Hooks, nil, input, output)
	req.ApplyOptions(opts...)

	return req, output
}

func (client Client) B2B(ctx context.Context, payload RequestB2B) (ResponseB2B, error) {
	req, out := client.B2BRequest(payload)
	req.WithContext(ctx)

	if err := req.Send(); err != nil {
		return ResponseB2B{}, err
	}

	return *out, nil
}

func (client Client) TransactionStatusRequest(input RequestTransactionStatus, opts ...gorequest.Option) (*gorequest.Request, *ResponseTransactionStatus) {
	op := gorequest.Operation{
		Name:   OperationTransactionStatus,
		Method: http.MethodPost,
		Path:   EndpointTransactionStatus,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// append to request options
	opts = append(opts, gorequest.WithRequestHeader("Content-Type", "application/json"))

	output := &ResponseTransactionStatus{}
	req := gorequest.New(cfg, op, client.Hooks, nil, input, output)
	req.ApplyOptions(opts...)

	return req, output
}

func (client Client) TransactionStatus(ctx context.Context, payload RequestTransactionStatus) (ResponseTransactionStatus, error) {
	req, out := client.TransactionStatusRequest(payload)
	req.WithContext(ctx)

	if err := req.Send(); err != nil {
		return ResponseTransactionStatus{}, err
	}

	return *out, nil
}

func (client Client) BalanceRequest(input RequestBalance, opts ...gorequest.Option) (*gorequest.Request, *ResponseBalance) {
	op := gorequest.Operation{
		Name:   OperationBalance,
		Method: http.MethodPost,
		Path:   EndpointAccountBalance,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// append to request options
	opts = append(opts, gorequest.WithRequestHeader("Content-Type", "application/json"))

	output := &ResponseBalance{}
	req := gorequest.New(cfg, op, client.Hooks, nil, input, output)
	req.ApplyOptions(opts...)

	return req, output
}

func (client Client) Balance(ctx context.Context, payload RequestBalance) (ResponseBalance, error) {
	req, out := client.BalanceRequest(payload)
	req.WithContext(ctx)

	if err := req.Send(); err != nil {
		return ResponseBalance{}, err
	}

	return *out, nil
}

func (client Client) QueryOrgInfoRequest(input RequestOrgInfoQuery, opts ...gorequest.Option) (*gorequest.Request, *ResponseOrgInfoQuery) {
	op := gorequest.Operation{
		Name:   OperationQueryOrgInfo,
		Method: http.MethodPost,
		Path:   EndpointQueryOrgInfo,
	}

	cfg := gorequest.Config{Endpoint: client.endpoint}

	// append to request options
	opts = append(opts, gorequest.WithRequestHeader("Content-Type", "application/json"))

	output := &ResponseOrgInfoQuery{}
	req := gorequest.New(cfg, op, client.Hooks, nil, input, output)
	req.ApplyOptions(opts...)

	return req, output
}

func (client Client) QueryOrgInfo(ctx context.Context, payload RequestOrgInfoQuery) (ResponseOrgInfoQuery, error) {
	req, out := client.QueryOrgInfoRequest(payload)
	req.WithContext(ctx)

	if err := req.Send(); err != nil {
		return ResponseOrgInfoQuery{}, err
	}

	return *out, nil
}
