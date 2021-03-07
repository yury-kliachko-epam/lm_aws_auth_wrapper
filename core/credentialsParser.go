package core

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
	return AWSCredentials{}
}
