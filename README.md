# Number Classification API

A RESTful API that provides mathematical properties and fun facts about numbers.

## Features

- Get mathematical properties of a number (prime, perfect, Armstrong)
- Get the sum of digits
- Get interesting mathematical facts about the number
- CORS enabled
- JSON responses
- Input validation
- Error handling

## API Specification

### Endpoint

```
GET /api/classify-number?number={number}
```

### Success Response (200 OK)

```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```

### Error Response (400 Bad Request)

```json
{
    "number": "invalid_input",
    "error": true
}
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository
```bash
git clone https://github.com/yourusername/numbersfactsapi.git
```

2. Navigate to the project directory
```bash
cd numbersfactsapi
```

3. Install dependencies
```bash
go mod tidy
```

4. Run the application
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Usage Examples

Using curl:
```bash
curl "http://localhost:8080/api/classify-number?number=371"
```

Using browser:
Navigate to `http://localhost:8080/api/classify-number?number=371`

## Technical Details

- Built with Go and Gin web framework
- Uses the Numbers API (numbersapi.com) for fun facts
- Implements CORS for cross-origin requests
- Response time < 500ms
- Supports all valid integers

## Error Handling

The API implements proper error handling for:
- Invalid input types
- Missing parameters
- Network issues with external API

## License

This project is licensed under the MIT License - see the LICENSE file for details
