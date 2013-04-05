# ansipix

As seen on http://curlgur.tk

## Golang image to ansi blocks

![gopher](http://i.imgur.com/hYecglD.png)

## Can I try it? 

Sure, build ansipix.go in the /bin/ folder and you're good to go.

`$ ./ansipix octocat.jpg 50`

* First argument is filename
* 2nd is width in columns (default 80)

![octocat](http://i.imgur.com/t3cGshc.png)

## Can I use it?

Sure, just 

``` go
include "github.com/minikomi/ansipix"
```

Then pass an `io.Reader` to Blockify:

```go
package main

// AnsipixOnPipe.go
import (
	"bufio"
	"fmt"
	"github.com/minikomi/ansipix"
	"os"
)

func main() {
	Stdin := bufio.NewReader(os.Stdin)

	str, err := ansipix.Blockify(Stdin, 80)
	if err != nil {
		panic(err)
	}

	fmt.Println(str)
}
```

and

`cat awesomeimage.gif | go run AnsipixOnPipe.go `

Have fun! :octocat:
