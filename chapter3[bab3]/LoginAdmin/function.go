package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/Rumahkopi/rkbackend"
)

func init() {
	functions.HTTP("Login", petaPediaPost)
}

func petaPediaPost(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://rumahkopi.yumerp.me")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://rumahkopi.yumerp.me")
	fmt.Fprintf(w, rkbackend.Login("PASETOPRIVATEKEY", "MONGOSTRING", "proyek3", "admin", r))

}