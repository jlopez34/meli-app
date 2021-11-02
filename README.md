# meli-app
# Backend - Galaxy
This is a PoC for a meli-app API

### API endpoints


| Method | URL                                                  | Description                                                                                |
|--------|------------------------------------------------------|--------------------------------------------------------------------------------------------|
| POST   | /api/v1/topsecret                                    | Searching position and message sent by a spaceship lost                                    |
| POST   | /api/v1/topsecret_split/:satellite_name              | Find position and message sent by a spaceship lost using information from single satellite |
| GET    | /api/v1/topsecret_split/:satellite_name              | Searching position and message sent by a spaceship lost                                    |


### Installation
No further installation needed, just `go get github.com/jlopez34/meli-app`

### Build
`go build cmd/main.go`

### Run
`go run cmd/main.go`

### Docker
SID


### TEST
Since is build and run project using the code shown in previous section, it's possible to use the following postman collection. https://github.com/jlopez34/meli-app/blob/master/enviroment/GoMeli.postman_collection.json
