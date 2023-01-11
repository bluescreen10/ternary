# ternary

Ternary operator for Go

# Usage

A ternary operator is a short version of an if-else statement that can make
the code more readable for some use cases such is variable initialization:

```go
import "github.com/bluescreen10/ternary"

function createServer(desiredPort int) {
    port := ternary.T( desiredPort > 0, desiredPort, 8080)
    // some code
}
```
