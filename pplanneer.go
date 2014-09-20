package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/zephyyrr/pplanneer/model"
	"github.com/zephyyrr/pplanneer/model/mock"

	"net/http"
)

func main() {
	m := martini.Classic()

	m.Group("/api", attachAPI,
		render.Renderer(),
		func(c martini.Context) {
			c.MapTo(mock.Instance(), (*model.Model)(nil))
		},
		func(c martini.Context, m model.Model) {
			c.Map(m.User("zephyyrr"))
		},
	)

	m.Run()
}

func attachAPI(api martini.Router) {
	api.Get("/user", getCurrentUser)
	api.Get("/users", getUsers)
	api.Post("/users", newUser)

	api.Get("/users/:name", getUser)
	//api.Put("/users/:name", updateUser)

	api.Get("/users/:name/projects", getProjects)
	//api.Post("/users/:name/projects", newProject)

	//apiGets.HandleFunc("/projects", getProjects)
}

func getUsers(r render.Render, m model.Model) {
	r.JSON(200, m.Users())
}

func getCurrentUser(r render.Render, u model.User) {
	r.JSON(200, u)
}

func getUser(params martini.Params, r render.Render, m model.Model) {
	r.JSON(200, m.User(params["name"]))
}

func newUser(w http.ResponseWriter, r *http.Request) {
	//err := m.NewUser(form[""])
	http.Error(w, "Not Implemented. Yet.", http.StatusNotImplemented)
}

func getProjects(params martini.Params, r render.Render, m model.Model) {
	if name := params["name"]; name != "" {
		r.JSON(200, m.User(name).Projects)
	} else {
		r.JSON(200, m.User("zephyyrr").Projects)

	}

}
