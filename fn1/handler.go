package function

import (
	"fmt"
	"os"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("%v", os.Environ())
}
