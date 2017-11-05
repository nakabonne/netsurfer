# netsurfer

# Test

```
$ go run examples/main.go
```

# Usage

```go
urls, err := netsurfer.SerpsURL("ruby")
if err != nil {
	fmt.Println("erorr!", err)
} else {
	fmt.Println("Success!")
	for _, v := range urls {
		fmt.Println(v)
	}
}
```

# License

`netsurfer` source code is available under the MIT [License](https://github.com/ryonakao/netsurfer/blob/master/LICENSE).