/**
 * @Author: yy
 * @Description:
 * @File:  template_test
 * @Version: 1.0.0
 * @Date: 2023/06/04 13:36
 */

package main

import (
	"html/template"
	"os"
	"testing"
)

func TestParseFiles(t *testing.T) {
	type Todo struct {
		Title string
		Done  bool
	}
	type TodoPageData struct {
		PageTitle string
		Todos     []Todo
	}
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	files, err := template.ParseFiles("index.html")
	if err != nil {
		return
	}
	// 执行Execute方法，将数据写入模板
	if err := files.Execute(os.Stdout, data); err != nil {
		t.Fatal(err)
	}
}

func TestParse(t *testing.T) {
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

}
