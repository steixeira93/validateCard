# Credit Card Validation API

This API provides a POST route to validate credit card numbers using the Luhn algorithm.

## How to Run

1. Clone the repository to your local machine.
2. Navigate to the project folder.
3. Run the command `go run main.go` to start the server.

## Usage

To validate a credit card number, send a POST request to `http://localhost:9090/validate_card` with a JSON body containing the card number, as shown in the example below:

```json
{
	"number": "1234567890123456"
}
```

The response will be a JSON indicating whether the card is valid or not:

```json
{
  "valid": true
}
```

### Implementation
The API uses the ```gorilla/mux``` framework for routing and handling HTTP requests.
The ```/validate_card``` endpoint accepts POST requests and uses the ```validateCard``` function to check the validity of the provided card number.
The ```validateCard``` function implements the Luhn Algorithm to determine if the card number is valid.

### Security
It's importante to ensure that the API ir secure and that card nunmbers are handled carefully. Do not store card nunmbers and make sure that communications with the API are encrypted using HTTPS.

### Contribuitions
Contribuitions are welcome! To contribute, please fork the repository, create a branch for your feature or fix, and submit a pull request =P
