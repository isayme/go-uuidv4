# uuidv4
generate a version 4 uuid.

> valid uuid with http://guid.one/parse

# Usage
```
package main

import (
	"fmt"

	"github.com/isayme/go-uuidv4"
)

func main() {
	id, _ := uuidv4.Generate()

	fmt.Println(id)
}
```
