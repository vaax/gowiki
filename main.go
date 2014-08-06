package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "html/template"
)

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func main() {
    //p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
    //p1.save()
    //p2, _ := loadPage("TestPage")
    //fmt.Println(string(p2.Body))
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
//    http.HandleFunc("/save/", saveHandler)
    http.ListenAndServe(":8081", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w,p)
}

//func saveHandler(w http.ResponseWriter, r *http.Request) error {
//    title := r.URL.Path[len("/save/"):]
//    filename := title + ".txt"
//    err := ioutil.WriteFile(filename, byte("p.Body"), 0600)
//    if err != nil {
//        p = &Page{Title: title}
//    }
//}
