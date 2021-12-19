package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/sjquant/nomadcoin/blockchain"
)

const (
	templateDir string = "templates"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

var templates *template.Template

func home(rw http.ResponseWriter, req *http.Request) {
	data := homeData{"Home", blockchain.GetBlockChain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		req.ParseForm()
		blockData := req.Form.Get("blockData")
		blockchain.GetBlockChain().AddBlock(blockData)
		http.Redirect(rw, req, "/", http.StatusPermanentRedirect)
	}
}

func Start(port int) {
	templates = template.Must(template.ParseGlob(templateDir + "/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "/partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
