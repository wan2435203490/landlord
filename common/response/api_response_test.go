package r

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSuccess(t *testing.T) {
	success := Success0("a")
	fmt.Printf("%#v\n", success)
	success2 := Success0(1)
	fmt.Printf("%#v\n", success2)
	success3 := Success0(true)
	fmt.Printf("%#v\n", success3)
}

func TestError(t *testing.T) {
	err1 := Error0(http.StatusPermanentRedirect, "err1")
	fmt.Printf("%#v\n", err1)
	err2 := Error0(http.StatusUnsupportedMediaType, "err2")
	fmt.Printf("%#v\n", err2)
	err3 := Error0(http.StatusServiceUnavailable, "")
	fmt.Printf("%#v\n", err3)
}
