# Terror

`Terror` is a Go library for creating and managing enhanced errors with stack trace details. It is designed to help developers easily trace the origins of errors and display them in a formatted tree structure for better debugging and readability.

## Features

- Create errors with contextual information (file, line number, and function name).
- Maintain an error chain with a wrapped cause.
- Format errors and their causes in a tree-like structure for better visualization.

## Installation

To use `terror`, run:

`go get -u github.com/iondodon/terror`


## Usage

### Creating a Terror

Create a `Terror` using the `New` function:

```go
package main

import (
	"fmt"

	"github.com/iondodon/terror"
)

func dbConnection() error {
	return terror.New("failed to connect to database", nil)
}

func fetchUserData() error {
	err := dbConnection()
	if err != nil {
		return terror.New("could not fetch user data", err)
	}
	return nil
}

func processUserRequest() error {
	err := fetchUserData()
	if err != nil {
		return terror.New("user request processing failed", err)
	}
	return nil
}

func main() {
	err := processUserRequest()
	if err != nil {
		fmt.Println(terror.FormatTree(err))
	}
}

```

## Output

The example above will print an error in the following format:

```
user request processing failed (/home/ion/terror-test/main.go:24 in main.processUserRequest)
└── could not fetch user data (/home/ion/terror-test/main.go:16 in main.fetchUserData)
    └── failed to connect to database (/home/ion/terror-test/main.go:10 in main.dbConnection)
```
