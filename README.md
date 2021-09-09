# ah-scanner-api

This API connects to the AH Mobile App API to easily allow a Raspberry Pi / ESP8266 to connect. This project belongs to the ah-scanner project which allows a barcode scanner to scan a GTIN to automatically add products to your shopping list.

## Usage

### Login

Get an auth code by logging in to:
```
https://login.ah.nl/secure/oauth/authorize?client_id=appie&redirect_uri=appie://login-exit&response_type=code
```

This will then redirect you to: `appie://login-exit?code=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` where `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` is your auth code. Put this into the root of the project in a file named `code.txt`
This will be used the first time to get the access token. The access token will then be stored in `token.json`.

### Start the server

```bash
go run main.go
```

## Example Requests

### Add product to shopping list using FIR (ID)
`POST("/shoppinglist/add")`

Body:
```json
{
  "type": "FIR",
  "id": 366138
}
```

### Add product to shopping list using GTIN
`POST("/shoppinglist/add")`

Body:
```json
{
  "type": "GTIN",
  "id": 8710739491087
}
```
