package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("What I like most about '%s': '%s'", string(req), string(req))
}
