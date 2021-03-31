package main

import(
  "net/http"
  "html/template"
  "strconv"
  "log"
)

type User struct{
  Id int
  First string
  Second string
}
type Note struct{
  IdUser int
  Title string
  Text string
}
var tmpl = template.Must(template.ParseGlob("./templates/*.html"))

func Handler(w http.ResponseWriter, r *http.Request){
  tmpl.ExecuteTemplate(w,"index.html",nil)
}

func procHandler(w http.ResponseWriter, r *http.Request){
  if r.Method != "POST" {
    http.Redirect (w,r,"/", http.StatusSeeOther)
    return
  }
  id_user,err := strconv.Atoi(r.FormValue("ider"))
  if err != nil{
    log.Fatal(err)
  }
  first_name := r.FormValue("firster")
  second_name := r.FormValue("laster")
  user = User {
    Id: id_user,
    First: first_name,
    Second: second_name,
  }
  err = user.addToDataBase()
  // here I just continue even tho the user exists ( you won't see any notify )
  tmpl.ExecuteTemplate(w,"processor.html",user)
}
var user User 
func NoteHandler(w http.ResponseWriter, r *http.Request){
  if r.Method != "POST" {
    http.Redirect (w,r,"/", http.StatusSeeOther)
    return
  }
  ///taking the user
  


  //// taking a note

  title := r.FormValue("title")
  text := r.FormValue("text")
  note := Note {
    Title: title,
    Text: text,
  }

  err := user.addANote(note)
  if err != nil {
    log.Fatal()
  }

  
  userSlice,err := user.getAllNotes()
  if err != nil {
    log.Fatal()

  }
  //user get all notes should be here 
  tmpl.ExecuteTemplate(w,"showNotes.html",userSlice)
  
}

func main(){
  http.HandleFunc("/",Handler)
  http.HandleFunc("/process",procHandler)
  http.HandleFunc("/yourNotes",NoteHandler)
  http.ListenAndServe(":8080",nil)
}
