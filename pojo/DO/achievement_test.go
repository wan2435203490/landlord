package DO

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestUUID(t *testing.T) {
	newString := strings.ReplaceAll(uuid.NewString(), "-", "")
	fmt.Println(newString)
}
