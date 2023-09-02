package until

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomString(t *testing.T) {
	str := RandomString(64)
	str2 := RandomString(64)
	fmt.Println(str)
	fmt.Println(str2)
	assert.Equal(t, 64, len(str))
	assert.NotEqual(t, str, str2)
}
