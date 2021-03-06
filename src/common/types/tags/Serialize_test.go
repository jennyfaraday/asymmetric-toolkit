package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bytes"
	"fmt"
	"testing"
)

func TestTag_Serialize(t *testing.T) {
	key := "name"
	o := NewTag()
	o.Add(key)
	tag := o.Serialize()
	expectedTag := []byte{1, 0, 4, 110, 97, 109, 101}
	fmt.Println("    :0.......8......F.")
	fmt.Println("tag1:", tag)
	fmt.Println("tag2:", expectedTag)
	errors.Assert(bytes.Equal(tag, expectedTag), fmt.Sprintf("Key-value mismatch. v:%s", string(tag)))
}
