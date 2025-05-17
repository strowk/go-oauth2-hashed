## Postgres Example

This example demostrates how to wrap a Postgres store into `ClientStoreWithHash` in order to store clients with hashed secrets.

To run this example, you will need to have a Postgres database running. `run.sh` script will create a database in Docker, so with docker installed you can run:

```bash
bash /run.sh
```

### Understanding the example

Essentially all we are doing here is after creating store, wrap it like this:

```golang
clientStoreWithHash := go_oauth2_hashed.NewClientStoreWithBcrypt(clientStore)
```

Then to store we need to use `clientStoreWithHash.HashAndCreate` instead of `clientStore.Create`, but we pass a reference to original `Create` method like this:

```golang
clientStoreWithHash.HashAndCreate(&models.Client{
    ID:     "client_id",
    Secret: "client_secret",
    Domain: "http://localhost:8080",
    Public: true,
    UserID: "user_id",
}, clientStore.Create)
```

This way client store would hash secret before calling original `Create` method.

`ClientStoreWithHash` implements `oauth2.ClientStore` in a way that returned client supports `oauth2.ClientPasswordVerifier` using hash, so you can just give `ClientStoreWithHash` directly to `oauth2` library without any changes.

