package main


/*
Exercise: Readers
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
 */
import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (myR MyReader) Read (b []byte) (int, error) {
	b[0] = 'A' // 65
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}