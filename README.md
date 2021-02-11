# Golang Gio practice

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

(maybe) Current Gio have a redering bug on MS-Windows so, maximize then restore window when running.

![hello_screenshot](./screenshot.png)
