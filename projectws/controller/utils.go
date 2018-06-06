package controller

import (
	"html/template"
	"io/ioutil"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"database/sql"
)

// Create connection
func Conn() *sql.DB {
	db, err := sql.Open("mysql", "root:@/sekolah")
	if err != nil {
		panic(0)
	}
	return db
}

// File Reader (output dalam bentuk []byte)
func ReadHtmlFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	if err !=nil{
		panic(0)
	}
	return data
}

// menyiapkan dan mengeksekusi template html & data yg akan ditampilkan
func RenderTemplate(w http.ResponseWriter, name string, data map[string]interface{}){
	funcs := template.FuncMap{"Add": TemplateAdd, "ToHTML": TemplateHTML}
	t := template.New("t1")
	t, err := t.Funcs(funcs).Parse(string(ReadHtmlFile(name)))
	if err!= nil{
		panic(0)
	}
	if data == nil {
		data = make(map[string]interface{})
	}
	data["baseurl"] = BaseURL
	t.Execute(w, struct{ Data map[string]interface{} }{Data: data})
}

func TemplateAdd(x interface{}, y int) int {
	x1 := x.(int)
	return x1 + y
}

func TemplateHTML(x interface{}) template.HTML {
	x1 := x.(string)
	return template.HTML(x1)
}
