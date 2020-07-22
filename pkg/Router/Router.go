package Router

import (
	"fmt"
	"github/ivalexandru/golang_chi_002/pkg/Users"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func StartServer() *chi.Mux {
	router := chi.NewRouter()

	//we can stack routes on top of each other
	//practic, u mount more endpoints, that are stacked in a function, ON TOP
	//of this one
	//  /api/users/cevaDinAlteHandlers
	// curl localhost:9090/api/users/getUser
	router.Mount("/api/users", Users.UsersRoutes())

	//a handler is a kind of constructor
	fmt.Println("starting server on :9090")
	log.Fatal(http.ListenAndServe(":9090", router))
	return router
}
