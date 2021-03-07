package main

import "github.com/yury-kliachko-epam/lm-aws-auth-helper/core"

func main() {
	wrapper := core.LmAWSAuthWrapper{}
	wrapper.Authenticate()
}
