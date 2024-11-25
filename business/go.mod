module github.com/portierglobal/vision-online-companion/business

go 1.23.2

replace github.com/portierglobal/vision-online-companion/db => ../db

require (
	github.com/go-resty/resty/v2 v2.16.1
	github.com/portierglobal/vision-online-companion/db v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.33.0
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)
