package data_api

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetData(api *string) ([]CustomerParam, error) {
	var res *http.Response
	var dataCollection CustomersParam
	var err error
	var body []byte
	res, err = http.Get(*api)
	body, err = io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &dataCollection)
	if err != nil {
		return nil, err
	}
	return dataCollection.Data, nil
}
