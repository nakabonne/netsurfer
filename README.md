# netsurfer

# Test

```
$ go run examples/main.go
```

# What can be done

- You can get the URL displayed on the first page when you google search
- You can get the HTML of the website displayed on the first page when you google search

# Installation

```
$ go get -u github.com/ryonakao/netsurfer
```

# Usage

```go
import (
	"fmt"
	"github.com/ryonakao/netsurfer"
)
urls, _ := netsurfer.SerpsURL("ruby")
for _, v := range urls {
	fmt.Println(v)
}
```

# License

`netsurfer` source code is available under the MIT [License](https://github.com/ryonakao/netsurfer/blob/master/LICENSE).