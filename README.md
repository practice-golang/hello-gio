# [Gio](https://gioui.org/) practice

![hello_screenshot](./screenshot.png)

Copied [gio-example/hello](https://git.sr.ht/~eliasnaur/gio-example/tree/main/item/hello)

## Run
```sh
$ go run ./main.go
```

## Font
```go
....

once.Do(func() {
		// register(text.Font{}, goregular.TTF)

		ttfBytes, err := ioutil.ReadFile("./fonts/dongdong.ttf")

		if err != nil {
			log.Fatal(err)
		}
		register(text.Font{}, ttfBytes)

....
```

For CJK showing, Changed font from `gioui.org/font/gofont` to `giofont`
