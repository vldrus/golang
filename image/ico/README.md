This package implements PNG-based [ICO](https://en.wikipedia.org/wiki/ICO_(file_format)) image encoder.

```go
import "github.com/vldrus/golang/image/ico"

//...

imageFile, err := os.Open("image.jpg")
// if err...
img, fmt, err := image.Decode(imageFile)
// if err...

iconFile, err := os.Create("icon.ico")
// if err...
err = ico.Encode(iconFile, img)
// if err...
```
