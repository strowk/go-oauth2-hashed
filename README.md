# go-oauth2-hashed

Small wrapper for go oauth2 client stores to keep client secret hashed.

## Usage

```go
package main

import (
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	go_oauth2_hashed "github.com/strowk/go-oauth2-hashed"
)

func main() {
	// ... create your store here

	// wrap the store using Bcrypt hashing
	clientStoreWithHash := go_oauth2_hashed.NewClientStoreWithBcrypt(clientStore)

	// store a client with hashed secret
	err := clientStoreWithHash.HashAndCreate(&models.Client{
		ID:     "client_id",
		Secret: "client_secret",
		Domain: "http://localhost:8080",
		Public: true,
		UserID: "user_id",
	}, clientStore.Create) // refer to original Create method
	if err != nil {
		panic(err)
	}

	// Now you can pass clientStoreWithHash to oauth2 library as usual store
	manager := manage.NewDefaultManager()
	manager.MapClientStorage(clientStore)
}

```
