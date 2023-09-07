package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"html/template"
	"github.com/gorilla/mux"
)


var tmpl *template.Template

type Todo struct{
  ID string
  Todo string  
  IsDone bool 

}


func todosHundler(w http.ResponseWriter, r *http.Request){


  tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


    err = tmpl.Execute(w, todos)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
   
  
}


func createTodoHundler(w http.ResponseWriter, r *http.Request){

  
  err := r.ParseForm()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

  
    newTodo := r.FormValue("todo")

    
    todos = append(todos, Todo{ID:strconv.Itoa(rand.Intn(100000)), Todo: newTodo, IsDone: false})

    
    http.Redirect(w, r, "/todos", http.StatusSeeOther)

}



func deleteTodoHundler(w http.ResponseWriter, r *http.Request){


  
  params := mux.Vars(r)

  for index, item := range todos{
    if item.ID == params["id"]{

      todos = append(todos[:index],todos[index+1:]...)
    }
  }
  



}

func doneTodoHundler( w  http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)


  for index, item := range todos{
    if item.ID == params["id"]{
      item.IsDone = !item.IsDone
      todos = append(todos[:index], todos[index+1:]...)
      todos = append(todos, Todo{ID:item.ID, Todo:item.Todo, IsDone:item.IsDone })
    }
  }
}


var todos[]Todo

func main(){
  
  

  fmt.Println(todos)
  r:= mux.NewRouter() 
  r.HandleFunc("/todos",todosHundler).Methods("GET")
  r.HandleFunc("/todos",createTodoHundler).Methods("POST")
  r.HandleFunc("/todos/{id}", deleteTodoHundler).Methods("DELETE")
  r.HandleFunc("/todos/{id}", doneTodoHundler).Methods("GET")
  fmt.Println("Listening and Serving at port 8080!")
  log.Fatal(http.ListenAndServe(":8080",r))
}

  
