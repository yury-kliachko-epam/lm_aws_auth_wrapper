package core

import (
	"flag"
	"log"
)

type Enviroment string

const (
	Sandbox Enviroment = "sandbox"
	Development Enviroment = "development"
	NonProd Enviroment = "non-prod"
	Production Enviroment = "production"
)

func RunApp() {
	environmentFlag := flag.String("e", "development", "Enviroment to get creds for")
	awsAuthToolPathFlag := flag.String("t", "", "Path to AWS Auth tool")
	aAccountPasswordEnvVarFlag := flag.String("a", "", "A-account password env var name")
	flag.Parse()

	var environment Enviroment = Enviroment(*environmentFlag)
	switch environment {
	case Sandbox, Development, NonProd, Production:
		wrapper := LmAWSAuthWrapper{
			environment: environment, 
			awsToolPath: *awsAuthToolPathFlag, 
			aAccountPasswordEnvVar: *aAccountPasswordEnvVarFlag,
		}
		wrapper.Authenticate()
	default:
		log.Printf("Specified environment is invalid. Please, specify the correct value")
	}
}
