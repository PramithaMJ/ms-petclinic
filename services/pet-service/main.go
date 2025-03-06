package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hudl/fargo"  // Eureka client library
	"github.com/satori/go.uuid"
	"net/http"
	"log"
)
func registerWithEureka() {
    // Create a Eureka connection
    conn := fargo.NewConn("http://localhost:8761/eureka")
    
    // Create an instance to register
    instance := fargo.Instance{
        HostName:         "localhost",
        Port:             8081,
        App:              "pet-service",
        IPAddr:           "127.0.0.1",
        VipAddress:       "pet-service",
        SecureVipAddress: "pet-service",
        Status:           fargo.UP,
        DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.MyOwn},
        LeaseInfo:        fargo.LeaseInfo{RenewalIntervalInSecs: 30},
    }
    
    // Register the instance
    err := conn.RegisterInstance(&instance)
    if err != nil {
        log.Fatalf("Error registering with Eureka: %v", err)
    }
    log.Println("Successfully registered with Eureka")
}

func main() {
	// Register the service with Eureka
	registerWithEureka()

	// Initialize the Gin framework
	r := gin.Default()

	// Endpoint to retrieve a pet
	r.GET("/pets", func(c *gin.Context) {
		// Generate a unique pet ID
		petID := uuid.NewV4()

		// Return a sample pet
		c.JSON(http.StatusOK, gin.H{
			"pet_id": petID.String(),
			"name":   "Fido",
			"type":   "Dog",
		})
	})

	// Start the service on port 8081
	r.Run(":8081")
}
