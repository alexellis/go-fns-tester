package function

import (
	"fmt"
	"os"
	"strings"
)

// Handle a serverless request
func Handle(req []byte) string {
	ret := ""
	for _, v := range os.Environ() {
		if strings.Contains(v, "_PORT") == false && strings.Contains(v, "_SERVICE_HOST") == false {
			ret = ret + v + "\n"
		}
	}

	return fmt.Sprintf("Filtered env-vars:\n%s", ret)
}
