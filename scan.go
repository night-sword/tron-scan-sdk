package tronscan

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type Scan struct {
	*apiKeyWRR
	client *resty.Client
}

func NewScan(client *resty.Client) *Scan {
	return &Scan{
		apiKeyWRR: newApiKeyWRR([]string{}),
		client:    client,
	}
}

// Get account detail information
//
//	see : https://docs.tronscan.org/api-endpoints/account#get-account-detail-information
//	demo: https://apilist.tronscanapi.com/api/accountv2?address=TSTVYwFDp7SBfZk7Hrz3tucwQVASyJdwC7
func (inst *Scan) GetAccount(ctx context.Context, address string) (response *GetAccountResponse, err error) {
	api := "/api/accountv2"
	params := map[string]string{"address": address}

	rsp, err := inst.get(ctx, api, params)
	if err != nil {
		return
	}

	return jsonDecode[GetAccountResponse](rsp)
}

// see: https://docs.tronscan.org/security-service/security-service-api
//
//	https://apilist.tronscanapi.com/api/security/auth/data?address=THSiB9MT2sCnAUgnYs9euMCY9aiZCD4HB5
func (inst *Scan) GetAccountAuthSecurity(ctx context.Context, address string) (response *GetAccountAuthSecurityResponse, err error) {
	api := "/api/security/auth/data"
	params := map[string]string{"address": address}

	rsp, err := inst.get(ctx, api, params)
	if err != nil {
		return
	}

	return jsonDecode[GetAccountAuthSecurityResponse](rsp)
}

// Get account's token list
//
//	see : https://docs.tronscan.org/api-endpoints/account#get-accounts-token-list
//	demo: https://apilist.tronscanapi.com/api/account/tokens?address=TSTVYwFDp7SBfZk7Hrz3tucwQVASyJdwC7&start=0&limit=20&hidden=0&show=0&sortType=0&sortBy=0&token=
func (inst *Scan) ListAccountTokens(ctx context.Context, request *ListAccountTokensRequest) (response *ListAccountTokensResponse, err error) {
	api := "/api/account/tokens"

	req := toStringMap(request)
	rsp, err := inst.get(ctx, api, req)
	if err != nil {
		return
	}

	return jsonDecode[ListAccountTokensResponse](rsp)
}

// Get account's transaction data
//
//	see : https://docs.tronscan.org/api-endpoints/transactions-and-transfers#get-accounts-transaction-datas
//	demo: https://apilist.tronscanapi.com/api/token_trc20/transfers-with-status?limit=1&start=0&trc20Id=TUpMhErZL2fhh4sVNULAbNKLokS4GjC1F4&address=TV6MuMXfmLbBqPZvBHdwFsDnQeVfnmiuSi
func (inst *Scan) ListTrc20TransfersWithStatus(ctx context.Context, request *ListTrc20TransfersWithStatusRequest) (response *ListTrc20TransfersWithStatusResponse, err error) {
	api := "/api/token_trc20/transfers-with-status"

	req := toStringMap(request)
	rsp, err := inst.get(ctx, api, req)
	if err != nil {
		return
	}

	return jsonDecode[ListTrc20TransfersWithStatusResponse](rsp)
}

func (inst *Scan) get(ctx context.Context, api string, params map[string]string) (body []byte, err error) {
	client := inst.client.R().SetContext(ctx)

	key := inst.getKey()
	if key != "" {
		client = client.SetHeader(API_KEY_HEADER_NAME, key)
	}

	rsp, err := client.SetQueryParams(params).Get(api)
	if err != nil {
		return
	}

	body = rsp.Body()
	return
}

func (inst *Scan) post(ctx context.Context, api string, input any) (body []byte, err error) {
	client := inst.client.R().SetContext(ctx)

	key := inst.getKey()
	if key != "" {
		client = client.SetHeader(API_KEY_HEADER_NAME, key)
	}

	client.SetBody(input)

	rsp, err := client.Post(api)
	if err != nil {
		return
	}

	body = rsp.Body()
	return
}
