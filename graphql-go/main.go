package main

import (
  "log"
  "net/http"
  "fmt"
  "io/ioutil"
  "github.com/graph-gophers/graphql-go"
  "github.com/graph-gophers/graphql-go/relay"
)

var Schema = `
  schema {
    query: Query
  }

  type Query {
    hello: String!
  }
`

type Resolver struct{}

func (r *Resolver) Hello() string {
  return "Hello world!"
}

var schema *graphql.Schema

func init () {
  schema = graphql.MustParseSchema(Schema, &Resolver{})
}

var page = []byte(`
  hello
`)

func main() {
  http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  page, err := ioutil.ReadFile("graphiql.html")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(page)
  }))
  http.Handle("/graphql", &relay.Handler{Schema: schema})
  fmt.Println("Listening on localhost:4001")
  log.Fatal(http.ListenAndServe("127.0.0.1:4001", nil))
}
