/* LowLevelServer
 */

package main

import (
	"net/http"
)

func main() {

	myHandler := http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		// Just return no content - arbitrary headers can be set, arbitrary body
		rw.WriteHeader(http.StatusNoContent)
	})

	http.ListenAndServe(":8080", myHandler)
}
