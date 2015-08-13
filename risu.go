package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"code.google.com/p/go-uuid/uuid"
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"

	"github.com/wantedly/risu/registry"
	"github.com/wantedly/risu/schema"
)

func create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()
	var opts schema.BuildCreateOpts
	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		log.Fatal(err)
	}

	if opts.Dockerfile == "" {
		opts.Dockerfile = "Dockerfile"
	}

	build := schema.Build{
		ID:             uuid.NewUUID(),
		SourceRepo:     opts.SourceRepo,
		SourceRevision: opts.SourceRevision,
		Name:           opts.Name,
		Dockerfile:     opts.Dockerfile,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	reg := registry.NewRegistry("localfs", "")
	reg.Set(build)

	// debug code
	builddata, err := reg.Get(build.ID)
	fmt.Fprintln(w, builddata)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reg := registry.NewRegistry("localfs", "")
	builds, err := reg.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(builds)
}

func show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	image := ps.ByName("image")
	fmt.Fprintf(w, "Build %s!\n", image)
}

func main() {
	router := httprouter.New()
	router.GET("/builds", index)
	router.GET("/builds/:image", show)
	router.POST("/builds", create)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}
