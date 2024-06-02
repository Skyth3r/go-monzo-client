package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func balance(c *MonzoClient, id string) error {
	path := "balance"
	requestURL := fmt.Sprintf("%s/%s", c.endpoints["APIURL"], path)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	q := req.URL.Query()
	q.Add("account_id", id)
	req.URL.RawQuery = q.Encode()

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

	var balance Balance

	err = json.NewDecoder(rsp.Body).Decode(&balance)
	if err != nil {
		return err
	}

	fmt.Printf("Balance: %d\n", balance.Balance)

	return nil
}
