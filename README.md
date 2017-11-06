# netsurfer

# Test

```
$ go run examples/main.go
```

# What can be done

- You can get the URL displayed on the first page when you google search
- You can get the HTML of the website displayed on the first page when you google search

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