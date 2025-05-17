module pg_example

go 1.23.3

replace github.com/strowk/go-oauth2-hashed => ../..

require (
	github.com/go-oauth2/oauth2/v4 v4.5.3
	github.com/jackc/pgx/v4 v4.18.3
	github.com/strowk/go-oauth2-hashed v0.0.0-00010101000000-000000000000
	github.com/vgarvardt/go-oauth2-pg/v4 v4.4.4
	github.com/vgarvardt/go-pg-adapter v1.1.0
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgtype v1.14.4 // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/vgarvardt/pgx-helpers/v4 v4.2.0 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/text v0.25.0 // indirect
)
