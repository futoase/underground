underground
===========

go package (testing)

How to Install
---------------

```shell
> go get github.com/futoase/underground
```

How to Use
-----------

**main.go**

```go
package main

import (
  "fmt"
  . "github.com/futoase/underground"
)

func main() {
  fmt.Printf(WelcomeToUnderground())
}
```

**build**

```shell
> go build main.go
> ./main
> Welcome to underground...
```

License
--------

MIT License
