package main

import (
	"fmt"
	"encoding/pem"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
)

type CertificateRequest struct {
	Certificate string `json:"certificate"`
}

func pemHandler(c *gin.Context) {
	// Parse the JSON request body
	var req CertificateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request: %s", err)})
		return
	}

	// Decode the PEM-encoded certificate
	block, _ := pem.Decode([]byte(req.Certificate))
	if block == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid pem certificate: %s", req.Certificate)})
		return
	}

	// Parse the DER-encoded certificate
	parsed, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to parse certificate: %s", err)})
		return
	}

	// Run zlint on the certificate
	zlintResultSet := zlint.LintCertificate(parsed)

	// Return the zlint results as the API response
	c.JSON(http.StatusOK, zlintResultSet)
}

func derHandler(c *gin.Context) {
	// Parse the JSON request body
	var req CertificateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request: %s", err)})
		return
	}

	// Parse the DER-encoded certificate
	parsed, err := x509.ParseCertificate([]byte(req.Certificate))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to parse certificate: %s", err)})
		return
	}

	// Run zlint on the certificate
	zlintResultSet := zlint.LintCertificate(parsed)

	// Return the zlint results as the API response
	c.JSON(http.StatusOK, zlintResultSet)
}


func main() {
	router := gin.Default()

	// Define the /pem route with the POST method
	router.POST("/pem", pemHandler)

	// Define the /der route with the POST method
	router.POST("/der", derHandler)

	// Start the HTTP server on port 8080
	router.Run(":8080")
}
