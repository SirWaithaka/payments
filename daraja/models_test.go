package daraja

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestResponseCode_MarshalJSON(t *testing.T) {

	type ts struct {
		ResCode ResponseCode `json:"ResponseCode"`
	}

	tcs := []struct {
		input    ts
		expected string
	}{
		{ts{ResCode: SuccessSubmission}, fmt.Sprintf(`{"ResponseCode":"%s"}`, SuccessSubmission.String())},
		{ts{ResCode: InvalidAccessToken}, fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAccessToken.String())},
		{ts{ResCode: InvalidAuthHeader}, fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAuthHeader.String())},
		{ts{ResCode: InvalidAuthType}, fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAuthType.String())},
		{ts{ResCode: InvalidGrantType}, fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidGrantType.String())},
		{ts{ResCode: InternalServerError}, fmt.Sprintf(`{"ResponseCode":"%s"}`, InternalServerError.String())},
	}

	for _, tc := range tcs {
		var buf bytes.Buffer
		err := jsoniter.NewEncoder(&buf).Encode(tc.input)
		assert.NoError(t, err)

		result := bytes.TrimRight(buf.Bytes(), "\n")

		assert.Equal(t, tc.expected, string(result))
	}

}

func TestResponseCode_MarshalText(t *testing.T) {

	type ts struct {
		ResCode ResponseCode `json:"ResponseCode"`
	}

	tcs := []struct {
		input    ts
		expected string
	}{
		{ts{ResCode: SuccessSubmission}, fmt.Sprintf(`{"ResponseCode":"%s"}`, SuccessSubmission.String())},
		{ts{ResCode: InvalidAccessToken}, fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAccessToken.String())},
		{ts{ResCode: InvalidAuthHeader}, fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAuthHeader.String())},
		{ts{ResCode: InvalidAuthType}, fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAuthType.String())},
		{ts{ResCode: InvalidGrantType}, fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidGrantType.String())},
		{ts{ResCode: InternalServerError}, fmt.Sprintf(`{"ResponseCode":"%s"}`, InternalServerError.String())},
	}

	for _, tc := range tcs {
		result, err := jsoniter.Marshal(&tc.input)

		assert.NoError(t, err)
		assert.Equal(t, tc.expected, string(result))
	}

}

func TestResponseCode_UnmarshalText(t *testing.T) {
	type ts struct {
		ResCode ResponseCode `json:"ResponseCode"`
	}

	tcs := []struct {
		input    string
		expected ResponseCode
	}{
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, SuccessSubmission.String()), SuccessSubmission},
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAccessToken.String()), InvalidAccessToken},
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAuthHeader.String()), InvalidAuthHeader},
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAuthType.String()), InvalidAuthType},
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidGrantType.String()), InvalidGrantType},
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, InternalServerError.String()), InternalServerError},
	}

	for _, tc := range tcs {
		var result ts
		err := jsoniter.Unmarshal([]byte(tc.input), &result)

		assert.NoError(t, err)
		assert.Equal(t, tc.expected, result.ResCode)
	}

}

func TestResponseCode_UnmarshalJSON(t *testing.T) {

	type ts struct {
		ResCode ResponseCode `json:"ResponseCode"`
	}

	tcs := []struct {
		input    string
		expected ResponseCode
	}{
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, SuccessSubmission.String()), SuccessSubmission},
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAccessToken.String()), InvalidAccessToken},
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAuthHeader.String()), InvalidAuthHeader},
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidAuthType.String()), InvalidAuthType},
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, InvalidGrantType.String()), InvalidGrantType},
		{fmt.Sprintf(`{"ResponseCode":"%s"}`, InternalServerError.String()), InternalServerError},
	}

	for _, tc := range tcs {
		var result ts
		err := jsoniter.NewDecoder(strings.NewReader(tc.input)).Decode(&result)

		assert.NoError(t, err)
		assert.Equal(t, tc.expected, result.ResCode)
	}

}

func TestToResponseCode(t *testing.T) {
	result := ToResponseCode(InvalidGrantType.String())
	assert.Equal(t, InvalidGrantType, result)
}
