package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

var lines []string

func main() {
	content, err := ioutil.ReadFile("lines.txt")
	if err != nil {
		fmt.Println("Err")
	}
	lines = strings.Split(string(content), "\n")
	//fmt.Println(lines)
	handleRequests()
}

func rndLine(w http.ResponseWriter, r *http.Request) {
	i := rand.Intn(len(lines))
	line := lines[i]
	accept_header := r.Header.Get("Accept")
	content := line
	if strings.Contains(accept_header, "html") {
		content = fmt.Sprintf(`
      <div style='height: 100vh; display: grid; place-items: center; overflow: hidden'>
        <h1 style='text-align: center'>%s<h1/>
      </div>`, line)
	}
	fmt.Fprint(w, content)
}

func handleRequests() {
	http.HandleFunc("/", rndLine)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
