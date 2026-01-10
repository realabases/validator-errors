# validator-errors

A simple lightweight Go package to format error messages from [`validator/v10`](https://github.com/go-playground/validator).

## Installation

```bash
go get github.com/realabases/validator-errors
```

## Usage

### Create an instance

```go
ve := validatorerrors.New()
```

### Add/Remove rules

Below are pre-made messages for some common tags (required, min, max, email):

```go
ve.AddDefaultRule("required") // "Username is required"
ve.AddDefaultRule("email")    // "Email must be a valid email"
```

You can also add your custom rules:

```go
ve.AddRule("max", func(e validator.FieldError) string {
    return fmt.Sprintf("%s cantttt beee moreee thannn %s chaaaarsss", e.Field(), e.Param())
})
```

And remove rules for whatever reason:

```go
ve.RemoveRule("min")
```

### Format errors

```go
errors := ve.FormatValidationErrors(err)
```

It returns a `map[string]string` with your custom messages

## Example

```go
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
    ve "github.com/realabases/validator-errors"
)

type LoginRequest struct {
    Username string `json:"username" validate:"required,min=4"`
    Password string `json:"password" validate:"required,min=8"`
}

func main() {
    validate := validator.New()
    loginErrors := ve.New()

    loginErrors.AddDefaultRule("required")
    loginErrors.AddRule("min", func(e validator.FieldError) string {
		return fmt.Sprintf("%s must be at least %s chars!!!!!", e.Field(), e.Param())
	})

    loginERrors.RemoveRule("required")

    req := LoginRequest{
        Username: "",
        Password: "123",
    }

    // Validate struct
    err := validate.Struct(req)
    if err != nil {
        customErrs := loginErrors.FormatValidationErrors(err)
        fmt.Println(customErrs)
    }
}
```

**Output:**

```go
map[Password:Password must be at least 8 chars!!!!!]
```
