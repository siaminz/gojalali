# github.com/ItsAminZamani/gojalali
Jalali Wrapper around [time.Time](https://golang.org/pkg/time), you can use all method from `time` package with this module.

## Usage


### Get the gojalali module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/ItsAminZamani/gojalali@v0.1.0
```

```go
package main

import (
    "fmt"

    "github.com/ItsAminZamani/gojalali"
)

func main() {
    // Get Current time in jalali format
    now := gojalali.Now()

    // Import Time from Go Standard Time
    gojalali.From(time.Now())

    // return time in string ex: 1401/08/24 15:34:65
    now.String()
}
```