package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func listAccounts(c *MonzoClient) error {
	path := "accounts"
	requestURL := fmt.Sprintf("%s/%s", c.endpoints["APIURL"], path)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	rsp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}

	if rsp.Body == nil {
		return fmt.Errorf("response body is empty")
	}

	var accountsResp AccountsResp

	err = json.NewDecoder(rsp.Body).Decode(&accountsResp)
	if err != nil {
		return err
	}

	for _, account := range accountsResp.Accounts {
		if !account.Closed {
			fmt.Printf("Account: %s - %s\n", account.Type, account.ID)
		}
	}

	return nil
}
