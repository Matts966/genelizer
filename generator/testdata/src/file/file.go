package a

import (
	"os"
)

func test10() {
	f, _ := os.Open("") // OK
	f.Close()
}

func test11() {
	_, _ = os.Open("") // want `should call os.Close when using os.File`
}
