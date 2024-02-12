package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

var (
	version string
	date    string
)

func main() {
	component := html()
	component.Render(context.Background(), os.Stdout)

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
