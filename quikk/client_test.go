package quikk_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"

	quikk2 "github.com/SirWaithaka/payments/quikk"
)

// TEST SUITES FOR REQUEST BUILDERS

func TestClient_ChargeRequest(t *testing.T) {
	endpoint := "http://foo.bar"

	requestID := xid.New().String()
	client := quikk2.New(quikk2.Config{Endpoint: endpoint})
	req, _ := client.ChargeRequest(quikk2.RequestCharge{}, requestID)

	// check endpoint
	assert.Equal(t, req.Request.URL.String(), endpoint+quikk2.EndpointCharge)
}

func TestClient_PayoutRequest(t *testing.T) {
	endpoint := "http://foo.bar"

	requestID := xid.New().String()
	client := quikk2.New(quikk2.Config{Endpoint: endpoint})
	req, _ := client.PayoutRequest(quikk2.RequestPayout{}, requestID)

	// check endpoint
	assert.Equal(t, req.Request.URL.String(), endpoint+quikk2.EndpointPayout)
}

func TestClient_TransferRequest(t *testing.T) {
	endpoint := "http://foo.bar"

	requestID := xid.New().String()
	client := quikk2.New(quikk2.Config{Endpoint: endpoint})
	req, _ := client.TransferRequest(quikk2.RequestTransfer{}, requestID)

	// check endpoint
	assert.Equal(t, req.Request.URL.String(), endpoint+quikk2.EndpointTransfer)
}

func TestClient_TransactionSearchRequest(t *testing.T) {
	endpoint := "http://foo.bar"

	requestID := xid.New().String()
	client := quikk2.New(quikk2.Config{Endpoint: endpoint})
	req, _ := client.TransactionSearchRequest(quikk2.RequestTransactionStatus{}, requestID)

	// check endpoint
	assert.Equal(t, req.Request.URL.String(), endpoint+quikk2.EndpointTransactionSearch)
}

func TestClient_BalanceRequest(t *testing.T) {
	endpoint := "http://foo.bar"

	requestID := xid.New().String()
	client := quikk2.New(quikk2.Config{Endpoint: endpoint})
	req, _ := client.BalanceRequest(quikk2.RequestAccountBalance{}, requestID)

	// check endpoint
	assert.Equal(t, req.Request.URL.String(), endpoint+quikk2.EndpointBalance)
}

// TEST SUITES FOR REQUEST EXECUTORS

func TestClient_TransactionSearch(t *testing.T) {
	resourceID := xid.New().String()

	// create a mock test server
	mux := http.NewServeMux()
	mux.HandleFunc(quikk2.EndpointTransactionSearch, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(fmt.Sprintf(`{"data":{"id":"12345","type":"search","attributes":{"resource_id":"%s"}}}`, resourceID)))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	client := quikk2.New(quikk2.Config{Endpoint: server.URL})
	res, err := client.TransactionSearch(t.Context(), quikk2.RequestTransactionStatus{}, xid.New().String())
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	assert.Equal(t, res.Data.Attributes.ResourceID, resourceID)
}
