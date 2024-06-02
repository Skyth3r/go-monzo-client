package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/99designs/keyring"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	client := NewClient()

	if client.clientID == "" || client.clientSecret == "" {
		return fmt.Errorf("the Client ID and Client secret were not found in env vars")
	}

	ring, err := keyring.Open(keyring.Config{
		ServiceName: "monzo-access-token",
	})
	if err != nil {
		return err
	}

	item, err := ring.Get("tokens")
	if err != nil {
		if errors.Is(err, keyring.ErrKeyNotFound) {
			if err := oauth(client); err != nil {
				return fmt.Errorf("failed to authenticate: %w", err)
			}
			tokens := client.accessToken + "::" + client.refreshToken
			if err := ring.Set(keyring.Item{
				Key:  "tokens",
				Data: []byte(tokens),
			}); err != nil {
				return fmt.Errorf("failed to set tokens in keychain: %w", err)
			}
		} else {
			return err
		}
	} else {
		tokens := string(item.Data)
		tokenSlice := strings.Split(tokens, "::")
		if len(tokenSlice) != 2 {
			return fmt.Errorf("unexpected token format: %s", tokens)
		}
		client.accessToken = tokenSlice[0]
		client.refreshToken = tokenSlice[1]
	}

	err = pingTest(client)
	if err != nil {
		return err
	}

	err = listAccounts(client)
	if err != nil {
		return err
	}

	// err = balance(client, "UK_RETAIL_ACCOUNT_ID")
	// if err != nil {
	// 	return err
	// }

	// err = listPots(client, "UK_RETAIL_ACCOUNT_ID")
	// if err != nil {
	// 	return err
	// }

	// err = depositToPot(client, "UK_RETAIL_ACCOUNT_ID", "POT_ID", 100)
	// if err != nil {
	// 	return err
	// }

	// err = withdrawFromPot(client, "UK_RETAIL_ACCOUNT_ID", "POT_ID", 100)
	// if err != nil {
	// 	return err
	// }

	return nil
}
