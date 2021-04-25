package core

import (
	"fmt"
	"log"
	"os"
)

type AWSCredentials struct {
	accessKeyID string
	secretAccessKey string
	sessionToken string
}

type CredentialsParser struct {
	credentialsFilePath string
}

func (parser CredentialsParser) parse() AWSCredentials {
	if parser.credentialsFilePath == "" {
		parser.credentialsFilePath = "~/.aws/credentials"
	}
	file, err := os.Open(parser.credentialsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
	return AWSCredentials{}
}
