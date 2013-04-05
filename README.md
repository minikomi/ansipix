ansipix
=======

## Golang image to ansi blocks

![gopher](http://i.imgur.com/uJJBQiO.png)

## Can I try it? 

Sure, build ansipix.go in the /bin/ folder and you're good to go.

`$ ./ansipix octocat.jpg 50`

* First argument is filename
* 2nd is width in columns (default 80)

![octocat](http://i.imgur.com/SC5yP7J.png)

## Can I use it?

Sure, just 

``` go
include "github.com/minikomi/ansipix"
```

Then pass an `io.Reader` to Blockify:

```go
  f, err := os.Open("octocat.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

  // str - ansified image as a string
  // err - error decoding image
	str, err := ansipix.Blockify(f, columns)
	if err != nil {
		panic(err)
	}
  fmt.Println(str)
```

Have fun! :octocat:
