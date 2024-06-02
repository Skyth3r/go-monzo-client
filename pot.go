package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func listPots(c *MonzoClient, id string) error {
	path := "pots"
	requestURL := fmt.Sprintf("%s/%s", c.endpoints["APIURL"], path)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	q := req.URL.Query()
	q.Add("current_account_id", id)
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

	var pots Pots

	err = json.NewDecoder(rsp.Body).Decode(&pots)
	if err != nil {
		return err
	}

	for _, pot := range pots.Pots {
		if !pot.Deleted {
			fmt.Printf("Pot: %s - %s\n", pot.Name, pot.ID)
		}
	}

	return nil
}

func depositToPot(c *MonzoClient, accountID, potID string, amount int64) error {
	path := "pots/" + potID + "/deposit"
	requestURL := fmt.Sprintf("%s/%s", c.endpoints["APIURL"], path)

	dedupeID := fmt.Sprintf("dedupe_id_%d", time.Now().UnixNano())

	form := url.Values{}
	form.Set("source_account_id", accountID)
	form.Set("amount", fmt.Sprintf("%d", amount))
	form.Set("dedupe_id", dedupeID)
	formData := form.Encode()

	req, err := http.NewRequest("PUT", requestURL, bytes.NewBufferString(formData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

	fmt.Println("Deposited to pot successfully!")

	return nil
}

func withdrawFromPot(c *MonzoClient, accountID, potID string, amount int64) error {
	path := "pots/" + potID + "/withdraw"
	requestURL := fmt.Sprintf("%s/%s", c.endpoints["APIURL"], path)

	dedupeID := fmt.Sprintf("dedupe_id_%d", time.Now().UnixNano())

	form := url.Values{}
	form.Set("destination_account_id", accountID)
	form.Set("amount", fmt.Sprintf("%d", amount))
	form.Set("dedupe_id", dedupeID)
	formData := form.Encode()

	req, err := http.NewRequest("PUT", requestURL, bytes.NewBufferString(formData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

	fmt.Println("Withdrawn from pot successfully!")

	return nil
}
