<p align="center">
  A Simple Go library that validates AWS cognito JWT token
</p>

<hr>

### :building_construction: Installation

```bash
  # Add to your go project using:
  go get github.com/eryk-vieira/go-cognito-jwt-validator
```

### :technologist: Usage

```go
package main

import (
  validator "github.com/eryk-vieira/go-cognito-jwt-validator"
)

func main(){
  // Change the config with your valid credentials
  validator := validator.New(&validator.Config{
		Region:          "AWS Cognito Region",
		CognitoPoolId:   "Cognito Pool Id",
		CognitoClientId: "Cognito Client Id",
	})
  
  err := validator.Validate("JWT Token")

  if err != nil {
    fmt.Println(err)
  }
}
```

## :closed_book: License

<p align="center">Released in 2022. This project is under the<a href="https://github.com/eryk-vieira/go-cognito-jwt-validator/blob/master/LICENSE"> MIT license</a> 🚀</p>

<p align="center"> Made with love by <a href="https://github.com/Eryk-Luiz">Eryk Luiz</a> 🚀</p>

