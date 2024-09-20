# Air_Qual - GoRestDemo

Air_Qual is a simple application that allows you to push data to a server and retrieve it later.

## Technologies used

-   [x] Go
-   [x] Docker
-   [x] Echo

## Usage Details

### Run the application

```bash
docker compose up --build
```

### Run the client

```bash
go build -o air_qual client/main.go
./air_qual
```

```
Send air quality data to the server, which then stores it in a basic database.

Usage:
  air_qual [flags]

Flags:
  -k, --apiKey string    API key
  -c, --city string      City (default "Warsaw")
  -t, --country string   Country (default "Poland")
  -h, --help             help for air_qual
  -s, --state string     State (default "Mazovia")
```
