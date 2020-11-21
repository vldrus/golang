package ico

import (
	"bytes"
	"image/jpeg"
	"io/ioutil"
	"testing"
)

func TestEncode(t *testing.T) {
	files := []string{
		"test.jpg",
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatalf("cannot read test image: %v", err)
		}

		img, err := jpeg.Decode(bytes.NewBuffer(b))
		if err != nil {
			t.Fatalf("cannot decode test image: %v", err)
		}

		var buf bytes.Buffer

		err = Encode(&buf, img)
		if err != nil {
			t.Fatalf("cannot encode test image: %v", err)
		}

		if buf.Len() != 21123 {
			t.Fatalf("encoded image size is %d bytes", buf.Len())
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	bb, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		b.Fatalf("cannot read test image: %v", err)
	}

	img, err := jpeg.Decode(bytes.NewBuffer(bb))
	if err != nil {
		b.Fatalf("cannot decode test image: %v", err)
	}

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		if err := Encode(&buf, img); err != nil {
			b.Fatalf("cannot encode test image: %v", err)
		}
	}
}
