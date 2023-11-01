# erratum

simple golang stdlib error wrapper

## How To Use

```go
package main

import (
	"errors"
	"fmt"

	"github.com/Hidayathamir/erratum"
)

func main() {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err3 := errors.New("err3")

	err := erratum.Wrap(err1, err2)
	err = erratum.Wrap(err, err3)

	fmt.Println("err\t\t=", err)
	fmt.Println("err is err1\t=", errors.Is(err, err1))
	fmt.Println("err is err2\t=", errors.Is(err, err2))
	fmt.Println("err is err3\t=", errors.Is(err, err3))
	fmt.Println("err cause\t=", erratum.Cause(err))

	err4 := errors.New("err4")
	err = erratum.AddCause(err, err4)
	fmt.Println("err\t\t=", err)
	fmt.Println("err is err4\t=", errors.Is(err, err4))
	fmt.Println("err cause\t=", erratum.Cause(err))
}
```

output:

```
err             = err3: err2: err1
err is err1     = true
err is err2     = true
err is err3     = true
err cause       = err1
err             = err3: err2: err1: err4
err is err4     = true
err cause       = err4
```
