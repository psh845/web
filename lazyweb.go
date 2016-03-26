package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

type Web struct {
	Title string
}

func www_root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("template/index.html")
	p := Web{Title: "Lazypic"}
	t.Execute(w, p)
}

func www_about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("template/about.html")
	p := Web{Title: "Lazypic:about"}
	t.Execute(w, p)
}

func www_opensource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("template/opensource.html")
	p := Web{Title: "Lazypic:Opensource"}
	t.Execute(w, p)
}

func www_fun(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("template/fun.html")
	p := Web{Title: "Lazypic:Fun"}
	t.Execute(w, p)
}

func www_coffeecat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	imagestr := "<center>"
	for i := 1; i < 6; i++ {
		imagestr += fmt.Sprintf(`<img src="/images/coffeecat/%02d.png"><br>`, i)
	}
	io.WriteString(w, head("Coffeecat")+menu("coffeecat")+imagestr+tail())
}

func main() {
	portPtr := flag.String("http", "", "service port ex):80")
	flag.Parse()
	if *portPtr == "" {
		fmt.Println("lazyweb service")
		flag.PrintDefaults()
		os.Exit(1)
	}
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.HandleFunc("/", www_root)
	http.HandleFunc("/fun", www_fun)
	http.HandleFunc("/opensource", www_opensource)
	http.HandleFunc("/coffeecat", www_coffeecat)
	http.HandleFunc("/about", www_about)
	http.ListenAndServe(*portPtr, nil)
}
