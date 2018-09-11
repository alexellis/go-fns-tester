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
		if strings.Contains(v, "PORT") == false {
			ret = ret + v + "\n"
		}
	}

	return fmt.Sprintf("%s", ret)
}
