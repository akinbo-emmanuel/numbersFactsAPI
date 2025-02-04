# numbersFactsAPI

A powerful and efficient RESTful API service that provides comprehensive mathematical properties and fascinating facts about numbers. Perfect for educational applications, math enthusiasts, or any project requiring number analysis.

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
git clone https://github.com/akinbo-emmanuel/numbersFactsAPI.git
```

2. Navigate to the project directory
```bash
cd numbersFactsAPI
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
# Get properties of a number
curl "http://localhost:8080/api/classify-number?number=371"

# Get properties of a prime number
curl "http://localhost:8080/api/classify-number?number=17"

# Get properties of a perfect number
curl "http://localhost:8080/api/classify-number?number=28"
```

Using JavaScript:
```javascript
fetch('http://localhost:8080/api/classify-number?number=371')
  .then(response => response.json())
  .then(data => console.log(data));
```

Using Python:
```python
import requests
response = requests.get('http://localhost:8080/api/classify-number?number=371')
data = response.json()
print(data)
```

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

MIT License

Copyright (c) 2025 numbersFactsAPI

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
