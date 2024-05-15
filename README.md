# OONICTLogs
Is a project intend to collect all certificate chains from the webconnectivity OONI measurements to Certificate Transparency logs, using the Let’s Encrypt testing Twig logs as a target. 

## How to run
### Build
`GO111MODULE=auto go build -o service main.go` 
`./servcice`

### Development
`GO111MODULE=auto go run main.go`

## Project structure
Is composed by 2 microprocess: 
### Chainfinder
Queries and process all measurements to extract the certificate chains.

### Submitter
Submtis the certificate chains into Twig.

