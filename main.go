// Package main implements a Number Classification API that provides mathematical properties
// and fun facts about numbers through a RESTful HTTP interface.
//
// This API is part of the HNG DevOps Stage 1 task and demonstrates:
// - RESTful API design principles
// - Mathematical computations
// - External API integration
// - Error handling
// - CORS configuration
//
// The API exposes a single endpoint:
// GET /api/classify-number?number={number}
//
// Example usage:
// curl "http://localhost:8080/api/classify-number?number=371"
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NumberResponse represents the JSON response structure for successful number classification requests.
// It includes various mathematical properties and interesting facts about the number.
//
// Example response:
//
//	{
//	    "number": 371,
//	    "is_prime": false,
//	    "is_perfect": false,
//	    "properties": ["armstrong", "odd"],
//	    "digit_sum": 11,
//	    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
//	}
type NumberResponse struct {
	Number     int      `json:"number"`     // The input number being classified
	IsPrime    bool     `json:"is_prime"`   // Whether the number is prime
	IsPerfect  bool     `json:"is_perfect"` // Whether the number is perfect
	Properties []string `json:"properties"` // List of properties (e.g., "armstrong", "odd", "even")
	DigitSum   int      `json:"digit_sum"`  // Sum of all digits in the number
	FunFact    string   `json:"fun_fact"`   // Interesting mathematical fact about the number
}

// ErrorResponse represents the JSON response structure for failed requests.
// It is returned when the input is invalid or cannot be processed.
//
// Example response:
//
//	{
//	    "number": "abc",
//	    "error": true
//	}
type ErrorResponse struct {
	Number string `json:"number"` // The invalid input that caused the error
	Error  bool   `json:"error"`  // Always true for error responses
}

// isArmstrong checks if a number is an Armstrong number.
// An Armstrong number (also known as a narcissistic number) is a number that is equal to
// the sum of its own digits each raised to the power of the number of digits.
//
// Mathematical definition:
// For a number with n digits d1, d2, ..., dn:
// number = d1^n + d2^n + ... + dn^n
//
// Examples:
// - 371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 27 + 343 + 1 = 371
// - 153 is an Armstrong number because 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153
//
// Parameters:
//   - n: The number to check
//
// Returns:
//   - bool: true if n is an Armstrong number, false otherwise
func isArmstrong(n int) bool {
	original := n
	sum := 0
	numDigits := int(math.Log10(float64(n))) + 1

	for n > 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		n /= 10
	}

	return sum == original
}

// isPrime determines if a number is prime using trial division algorithm.
// A prime number is a natural number greater than 1 that is only divisible by 1 and itself.
//
// The function uses trial division up to the square root of n for efficiency.
// This optimization is based on the fact that if n is divisible by a number greater
// than its square root, it must also be divisible by a smaller number.
//
// Examples:
// - 2 is prime (smallest prime number)
// - 17 is prime
// - 4 is not prime (divisible by 2)
//
// Parameters:
//   - n: The number to check for primality
//
// Returns:
//   - bool: true if n is prime, false otherwise
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// isPerfect checks if a number is perfect.
// A perfect number is a positive integer that is equal to the sum of its proper positive divisors
// (all its positive divisors excluding the number itself).
//
// Mathematical definition:
// n is perfect if n = sum of all proper divisors of n
//
// Examples:
// - 6 is perfect because 1 + 2 + 3 = 6
// - 28 is perfect because 1 + 2 + 4 + 7 + 14 = 28
// - 496 is perfect because 1 + 2 + 4 + 8 + 16 + 31 + 62 + 124 + 248 = 496
//
// The function optimizes divisor calculation by:
// 1. Only checking up to square root of n
// 2. Adding both divisors when finding one (if i divides n, n/i is also a divisor)
//
// Parameters:
//   - n: The number to check
//
// Returns:
//   - bool: true if n is a perfect number, false otherwise
func isPerfect(n int) bool {
	if n <= 1 {
		return false
	}
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n
}

// digitSum calculates the sum of all digits in a number.
// This function breaks down a number into its individual digits and adds them together.
//
// Examples:
// - digitSum(123) = 1 + 2 + 3 = 6
// - digitSum(999) = 9 + 9 + 9 = 27
// - digitSum(1000) = 1 + 0 + 0 + 0 = 1
//
// Parameters:
//   - n: The number whose digits are to be summed
//
// Returns:
//   - int: The sum of all digits in the number
func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// getFunFact retrieves an interesting mathematical fact about a number from the Numbers API.
// The function makes an HTTP GET request to numbersapi.com, which provides mathematical
// trivia about numbers.
//
// Features:
// - 10-second timeout to prevent hanging
// - Graceful fallback if API is unreachable
// - Automatic response body closing
//
// The function will return a generic message if:
// - The HTTP request fails
// - The API response cannot be read
// - The timeout is exceeded
//
// Parameters:
//   - n: The number to get a fact about
//
// Returns:
//   - string: A mathematical fact about the number, or a generic message if the API call fails
func getFunFact(n int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", n)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Sprintf("%d is an interesting number.", n)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("%d is an interesting number.", n)
	}

	return string(body)
}

// classifyNumber is the main HTTP handler function that processes number classification requests.
// This function serves as the core of the API, combining all number properties into a single response.
//
// Request Processing:
// 1. Extracts the 'number' query parameter
// 2. Validates and converts the input to an integer
// 3. Determines various mathematical properties
// 4. Fetches an interesting fact about the number
// 5. Returns a JSON response
//
// Error Handling:
// - Returns 400 Bad Request for invalid number inputs
// - Properly formats error responses
//
// Example Usage:
// GET /api/classify-number?number=371
//
// Parameters:
//   - c: Gin context containing the HTTP request and response utilities
func classifyNumber(c *gin.Context) {
	numberStr := c.Query("number")
	number, err := strconv.Atoi(numberStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Number: numberStr,
			Error:  true,
		})
		return
	}

	var properties []string
	if isArmstrong(number) {
		if number%2 == 0 {
			properties = []string{"armstrong", "even"}
		} else {
			properties = []string{"armstrong", "odd"}
		}
	} else {
		if number%2 == 0 {
			properties = []string{"even"}
		} else {
			properties = []string{"odd"}
		}
	}

	response := NumberResponse{
		Number:     number,
		IsPrime:    isPrime(number),
		IsPerfect:  isPerfect(number),
		Properties: properties,
		DigitSum:   digitSum(number),
		FunFact:    getFunFact(number),
	}

	c.JSON(http.StatusOK, response)
}

// main initializes and starts the HTTP server with proper CORS configuration.
// This function sets up the Gin router, configures CORS, defines routes,
// and starts the server on port 8080.
//
// Server Configuration:
// - Uses Gin's default middleware (Logger and Recovery)
// - Enables CORS for all origins
// - Allows only GET methods
// - Serves on port 8080
//
// Endpoints:
// - GET /api/classify-number: Classifies a number and returns its properties
//
// The server can be stopped with Ctrl+C (SIGINT)
func main() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET"}
	r.Use(cors.New(config))

	// Routes
	r.GET("/api/classify-number", classifyNumber)

	// Start server
	r.Run(":8080")
}
