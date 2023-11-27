# Auth0 + Golang Example

This sample demonstrates how to add authentication to a Go web app using Auth0.

## Running the App

To run the app, make sure you have **go** installed.

Rename the `.env.example` file to `.env` and provide your Auth0 credentials.

```bash
# .env

AUTH0_CLIENT_ID={CLIENT_ID}
AUTH0_DOMAIN={DOMAIN}
AUTH0_CLIENT_SECRET={CLIENT_SECRET}
AUTH0_CALLBACK_URL=http://localhost:3000/callback
```

Once you've set your Auth0 credentials in the `.env` file, run `go mod vendor` to download the Go dependencies.

Run `go run main.go` to start the app and navigate to [http://localhost:3000/](http://localhost:3000/).
