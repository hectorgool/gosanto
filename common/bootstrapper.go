package common

import (
	"fmt"
)
// StartUp bootstrapps the application
func StartUp() {
	
	fmt.Printf("bootstrapper:0\n")
	// Initialize AppConfig variable
	initConfig()

	fmt.Printf("bootstrapper:1\n")
	// Initialize private/public keys for JWT authentication
	initKeys()

	fmt.Printf("bootstrapper:2\n")
	// Initialize Logger objects with Log Level
	setLogLevel(Level(AppConfig.LogLevel))

	fmt.Printf("bootstrapper:3\n")
	// Start a MongoDB session
	createDbSession()

	fmt.Printf("bootstrapper:4\n")
	// Add indexes into MongoDB
	addIndexes()
	
	fmt.Printf("bootstrapper:5\n")
}