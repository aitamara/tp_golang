package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func heure(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, time.Now().Format("15:04"))
}

func add(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if err := req.ParseForm(); err != nil {
		fmt.Println("Something went bad")
		fmt.Fprintln(w, "Something went bad")
		return
	}
	author := req.PostForm.Get("author")
	entry := req.PostForm.Get("entry")
	fmt.Println(author, ": ", entry)
	if len(author) > 0 && len(entry) > 0 {
		saveOnFile(author, entry)
		fmt.Fprintf(w, author+": "+entry)
	} else {
		fmt.Fprintf(w, "Missing parameters")
	}

}

func saveOnFile(author, entry string) {
	f, err := os.OpenFile("save.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(author + ": " + entry + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Enregistré dans le fichier save.txt")
}

func allEntries() []string {
	raw, err := os.ReadFile("save.txt")

	if err != nil {
		panic(err)
	}

	entries := strings.Split(string(raw), "\n")

	return entries
}

func entries(w http.ResponseWriter, req *http.Request) {
	entries := allEntries()

	for _, entry := range entries {
		entry := strings.Split(entry, ":")

		fmt.Fprintf(w, entry[1]+"\n")
	}
}

func main() {
	http.HandleFunc("/", heure)
	http.HandleFunc("/add", add)
	http.HandleFunc("/entries", entries)

	fmt.Println("Serveur démarré")
	http.ListenAndServe(":4567", nil)
}
