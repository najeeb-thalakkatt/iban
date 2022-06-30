# IBAN Validator

This solution is based on https://en.wikipedia.org/wiki/International_Bank_Account_Number#Validating_the_IBAN

An IBAN is validated by converting it into an integer and performing a basic mod-97 operation (as described in ISO 7064) on it. If the IBAN is valid, the remainder equals 1.[Note 1] The algorithm of IBAN validation is as follows:[8]

1. Check that the total IBAN length is correct as per the country. If not, the IBAN is invalid
2. Move the four initial characters to the end of the string
3. Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35
4. Interpret the string as a decimal integer and compute the remainder of that number on division by 97
5. If the remainder is 1, the check digit test is passed and the IBAN might be valid.

Example (fictitious United Kingdom bank, sort code 12-34-56, account number 98765432):

* IBAN:		GB82 WEST 1234 5698 7654 32	
* Rearrange:		W E S T12345698765432 G B82	
* Convert to integer:		3214282912345698765432161182	
* Compute remainder:		3214282912345698765432161182	mod 97 = 1

## How to run

For getting the dependencies from the root folder run:

`go mod vendor`

To start the service:

`export LOG_LEVEL=DEBUG;export SERVICE=iban;export VERSION=1; go run cmd/iban/main.go`

## How to run via Docker
`docker run --rm -d  -p 8080:8080/tcp iban:latest`
## How to run test

`go test ./... -v     `

## /iban/validate api 
### url: localhost:8080/iban/validate  POST
payload:

`{"IBAN":"BE71 0961 2345 6769"}`

response:

`{"result":"valid"}`

`{"Error":"invalid iban"}`

`{"Error":"IBAN length for DE is 22"}`
