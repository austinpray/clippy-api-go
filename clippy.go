package main

import (
  "net/http"
  "log"
  "github.com/gorilla/mux"
  "github.com/codegangsta/negroni"
  "github.com/rs/cors"
  "os"
)

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", CapabilitiesIndex)
  router.HandleFunc("/capabilities", CapabilitiesIndex)

  router.HandleFunc("/sync", SyncIndex)

  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
  })

  n := negroni.Classic()
  n.Use(c)
  n.Use(negroni.HandlerFunc(middlewareJSON))
  n.UseHandler(router)

  log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), n))
}
