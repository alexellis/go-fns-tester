package function

import (
	"io/ioutil"
	"log"
	"os"
)

// Handle a serverless request
func Handle(req []byte) string {
	data, err := ioutil.ReadFile("/var/run/" + os.Getenv("secret_name"))

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
