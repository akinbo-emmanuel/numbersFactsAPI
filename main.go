// Package main implements a Number Classification API that provides mathematical properties
// and fun facts about numbers through a RESTful HTTP interface.
package main

import (
	// "encoding/json"
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
type ErrorResponse struct {
	Number string `json:"number"` // The invalid input that caused the error
	Error  bool   `json:"error"`  // Always true for error responses
}

// isArmstrong checks if a number is an Armstrong number.
// An Armstrong number is a number that is the sum of its own digits each raised to
// the power of the number of digits.
// For example, 371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371
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

// isPrime determines if a number is prime.
// A prime number is a natural number greater than 1 that is only divisible by 1 and itself.
// The function uses trial division up to the square root of n for efficiency.
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
// A perfect number is a positive integer that is equal to the sum of its proper positive divisors.
// For example, 6 is perfect because 1 + 2 + 3 = 6
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
// For example, digitSum(123) returns 6 (1 + 2 + 3)
func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// getFunFact retrieves an interesting mathematical fact about a number from the Numbers API.
// If the API call fails, it returns a generic message about the number.
// The function includes a timeout to prevent long waiting times.
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
// It accepts a number parameter via query string and returns a JSON response with various
// mathematical properties of the number.
// If the input is invalid, it returns a 400 Bad Request status with an error response.
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
// The server listens on port 8080 and provides the /api/classify-number endpoint.
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
