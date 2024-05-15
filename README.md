# OONICTLogs
Is a project intended to collect all certificate chains from the webconnectivity OONI measurements to create Certificate Transparency logs, using the Let’s Encrypt testing Twig logs as a target.


## How to run
### Build
`GO111MODULE=auto go build -o service main.go` 
`./servcice`

## Development
`GO111MODULE=auto go run main.go`

### Tests
Unit tests within the package by running the command: ``GO111MODULE=auto go test`



## Project structure
Is composed by 2 sub-processes: 

### Finder
Queries and processes all measurements to extract the certificate chains.

### Submitter
Submits the certificate chains into Twig.

