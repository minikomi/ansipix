# ansipix

As seen on http://curlgur.tk

## image to ansi blocks in Go

![gopher](http://i.imgur.com/hYecglD.png)

## Can I try it? 

Sure, build ansipix.go in the /bin/ folder and you're good to go.

`$ ./ansipix octocat.jpg 50`

* First argument is filename
* 2nd is width in columns (default 80)

![octocat](http://i.imgur.com/t3cGshc.png)

## Can I use it?

Sure, just 

``` shell
$ go get "github.com/minikomi/ansipix"
```

Then pass an `io.Reader` & width as an int to Blockify:

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

`curl https://www.google.com/images/srpr/logo4w.png  | go run test.go`

![google](http://i.imgur.com/Vh3jngN.png)

Have fun! :octocat:
