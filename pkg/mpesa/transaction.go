package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var transactionEndpoint = "mpesa/transactionstatus/v1/query"

// TransactionStatus Check the status of a transaction
func (sdk mSDK) TransactionStatus(tReq TransactionReq) (TransactionResp, error) {
	data, err := json.Marshal(tReq)
	if err != nil {
		return TransactionResp{}, err
	}
	url := fmt.Sprintf("%s/%s", sdk.baseURL, transactionEndpoint)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return TransactionResp{}, err
	}
	resp, err := sdk.sendRequest(req)
	if err != nil {
		return TransactionResp{}, err
	}

	var tr TransactionResp
	if err := json.Unmarshal(resp, &tr); err != nil {
		return TransactionResp{}, err
	}
	return tr, nil
}
