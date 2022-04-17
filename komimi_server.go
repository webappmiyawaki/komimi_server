package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Page struct {
	Title                      string
	Body                       []byte
	Key1                       string
	Key2                       string
	CustomerBalance            string
	CustomerPaymentInformation bool
	Information                string
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

var templates = template.Must(template.ParseFiles("view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	//t, _ := template.ParseFiles(tmpl + ".html")
	//t.Execute(w, p)
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "view"+title, http.StatusFound)
		return
	}
	key1 := r.URL.Query().Get("key1")
	key2 := r.URL.Query().Get("key2")
	p.Key1 = key1
	p.Key2 = key2
	account := AccountInfo{key1, key2, "0", false}
	ac := account.SelectAnyId()
	p.CustomerBalance = ac[0].CustomerBalance
	balance, _ := strconv.Atoi(p.CustomerBalance)
	pay, _ := strconv.Atoi(key2)
	if balance > pay {
		p.CustomerPaymentInformation = true

		account.CustomerBalance = strconv.Itoa(balance - pay)
		p.Information = key2 + "の決済が完了できました。残高は「" + account.CustomerBalance + "」です。"
		account.Update()
	} else {
		p.CustomerPaymentInformation = false
		p.Information = "決済が出来ませんでした。"
	}
	renderTemplate(w, "view", p)
}

var validPath = regexp.MustCompile("^/(view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func komimiServer() {
	http.HandleFunc("/view/", makeHandler(viewHandler))

	//:5050の前に何も書かなければローカルホストになる
	log.Fatal(http.ListenAndServe(":5050", nil))
}
