package Users

import (
	"encoding/json"
	"net/http"
)

// w is what you're sending back to the client
// r  = what the client is sending to you (you, the server)
func getHandler(w http.ResponseWriter, r *http.Request) {
	//encoding the w into json
	json.NewEncoder(w).Encode("u got user1")
}

// curl localhost:9090/api/updateUser -X POST -d '{"ceva"}'
func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("you updated user1")
}
