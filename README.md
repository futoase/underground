[![Build Status](https://travis-ci.org/futoase/underground.png?branch=support-test)](https://travis-ci.org/futoase/underground)

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

Documents
---------

http://godoc.org/github.com/futoase/underground

License
--------

MIT License


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/futoase/underground/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

