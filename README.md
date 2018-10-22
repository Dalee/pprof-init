# Pprof init

Simple library to automate pprof serving.

Example:

```go
package main

import (
	_ "github.com/Dalee/pprof_init"
	"time"
)

func main() {
	time.Sleep(time.Second * 100)
}
```

Now you have a running pprof on `:6060`.
