package writer

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
)


func (o *Writer) Close() {
	if o.file != nil {
		err := o.file.Close()
		errors.Assert(err == nil, fmt.Sprintf("Failed to close file.  Error:%v", err))
	}
}