package main

import (
	"context"
	"fmt"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"

	go_oauth2_hashed "github.com/strowk/go-oauth2-hashed"

	"github.com/go-oauth2/oauth2/v4/manage"
)

func main() {
	pgxConn, err := pgx.Connect(context.Background(), "postgres://postgres:mysecretpassword@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	defer pgxConn.Close(context.Background())

	adapter := pgx4adapter.NewConn(pgxConn)

	clientStore, err := pg.NewClientStore(adapter)
	if err != nil {
		panic(err)
	}

	clientStoreWithHash := go_oauth2_hashed.NewClientStoreWithBcrypt(clientStore)

	err = clientStoreWithHash.HashAndCreate(&models.Client{
		ID:     "client_id",
		Secret: "client_secret",
		Domain: "http://localhost:8080",
		Public: true,
		UserID: "user_id",
	}, clientStore.Create)
	if err != nil {
		panic(err)
	}

	client, err := clientStoreWithHash.GetByID(context.Background(), "client_id")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hashed secret: %s\n", client.GetSecret())

	if client.(oauth2.ClientPasswordVerifier).VerifyPassword("client_secret") {
		fmt.Println("Secret verified OK")
	} else {
		panic("Secret should be verified")
	}

	if client.(oauth2.ClientPasswordVerifier).VerifyPassword("wrong_password") {
		panic("Secret should not be verified")
	}

	manager := manage.NewDefaultManager()
	manager.MapClientStorage(clientStore)
}
