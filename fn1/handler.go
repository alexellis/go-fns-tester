package function

import (
	"fmt"
	"os"
)

// Handle a serverless request
func Handle(req []byte) string {
	ret := ""
	for _, v := range os.Environ() {
		ret = ret + v + "\n"
	}

	return fmt.Sprintf("%s", ret)
}
