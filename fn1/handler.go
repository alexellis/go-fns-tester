package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("What I like most about rootless builds is: '%s'", string(req))
}
