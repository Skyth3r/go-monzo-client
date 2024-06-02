# Go Monzo Client

A third party OAuth client for accessing the [Monzo API](https://docs.monzo.com/) written in Go. Access and refresh tokens are stored within the operating system's keystore system via the [keyring package](https://github.com/99designs/keyring)

## Prerequisites

Before running the program, you'll need to register a new API client on the [Monzo Developer Portal](https://developers.monzo.com). 

To register a new API client, log inot the Monzo Devloper Portal (remember to approve the login via your Monzo app) and click "New OAuth Client".

Then provide the following details for your OAuth client (Logo URL can remain blank):

```
Name: Go Monzo Client 
Redirect URL: http://127.0.0.1:21234/callback
Description: Go Monzo Client Application
Confidentiality: True
```

Once the client is registered you will recieve a Client ID and a Client Secert. Make a note of these!

## How to use
Clone the repository (I like using the GitHub CLI for this)
```bash
gh repo clone skyth3r/go-monzo-client
```

Install dependencies
```bash
go mod tidy
```

Set Client ID and Client Secert in environment variables
```bash
export MONZO_CLIENT_ID=YOUR_CLIENT_ID_HERE

export MONZO_CLIENT_SECRET=YOUR_CLIENT_SECRET_HERE
```

Run the program
```bash
go run ./
```

## Expected results

The first time this code is run, the client will start the OAuth flow, and attempt to open a browser with the login URL. On the login page, type in your email address linked to your personal Monzo account and then click the link sent to your email address and go back to the app.

You will then be prompted to open the Monzo app and grant access to the app by clicking "Allow access to your data". This process is related to Strong Customer Authentication. Once access has been granted via the Monzo app, go back to the app and press the [Enter] key to continue.

The app will attempt to make an API request to the accounts endpoint and list all the accounts you own that are not closed.

For future uses of the app, the access and refresh token will be retrieved from the system's keystore (e.g. on MacOS it woud be retrieved from Keychain). You might be prompted to allow the app to access to keystore. For now it is adviced you grant access with the 'Allow' option rather than the 'Always Allow' option.

The app also has support for the balance and pots endpoints (with more coming soon!). You can try these out by uncommenting the code in main for each of these endpoints and providing the function with the required arguments to make the API call.
