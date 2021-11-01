# meli-app
# Backend - Galaxy
This is a PoC for a meli-app API

### API endpoints


| Method | URL                                                  | Description                       |
|--------|------------------------------------------------------|-----------------------------------|
| POST   | /api/v1/topsecret                                    | Get all patients                  |
| POST   | /api/v1/topsecret_split/:satellite_name              | Get one patient                   |
| GET    | /api/v1/topsecret_split/:satellite_name              | Add one patient                   |

### Introduction
Just need to deploy the fuego instance locally or in a cloud provider and connect with a tcp client or as a web server.
Need different "modes" if you need a TCP plain connection or a web server, just add it in the json config file, located in the root.

### Installation
No further installation needed, just `go get github.com/jlopez34/meli-app`

### Build
`make build`

### Run
`make run`

### Docker
SID


### TEST
this is a test
------------------------------------------------------------
