package core

import (
	"log"
	"os"
	"os/exec"
	"io"
	"fmt"
)

type LmAWSAuthWrapper struct {
	environment Enviroment
	awsToolPath string
	aAccountPasswordEnvVar string
}

func (w LmAWSAuthWrapper) Authenticate() {
	log.Printf("Authenticate to %s environment", w.environment)
	password := os.Getenv(w.aAccountPasswordEnvVar)
	if password == "" {
		log.Printf("No AWS password found, you must set it as %s env var", w.aAccountPasswordEnvVar)
		return
	}
	log.Printf("AWS password found")
	authenticate()
	parseCredentials()
	//setEnvVars()
	log.Printf("Trying to authenticate with password %s", password)
}

func authenticate() {
	cmd := exec.Command("echo")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
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
