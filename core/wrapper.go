package core

import (
	"log"
	"os"
)

var awsToolPathEnvVar = "AWS_AUTH_TOOL_PATH"
var aAccountPassword = "AACCOUNT_PASSWORD"

type LmAWSAuthWrapper struct{}

func (wrapper LmAWSAuthWrapper) Authenticate() {
	pathToAwsTool := os.Getenv(awsToolPathEnvVar)
	if pathToAwsTool == "" {
		log.Printf("No path for AWS tool found, you must set it as %s env var", awsToolPathEnvVar)
		return
	}
	log.Printf("Path to AWS tool is %s", pathToAwsTool)
	password := os.Getenv(aAccountPassword)
	if password == "" {
		log.Printf("No AWS password found, you must set it as %s env var", aAccountPassword)
		return
	}
	log.Printf("AWS password found")
	authenticate()
	parseCredentials()
	setEnvVars()
	log.Printf("Trying to authenticate with password %s", password)

}

func authenticate() {
	log.Println("auth ...")
}

func parseCredentials() {
	log.Println("parsing credentials file...")
}

func setEnvVars(credentials AWSCredentials) {
	log.Println("setting credentials to env vars ...")

	os.Setenv("AWS_ACCESS_KEY_ID", credentials.accessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", credentials.secretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", credentials.sessionToken)
}
