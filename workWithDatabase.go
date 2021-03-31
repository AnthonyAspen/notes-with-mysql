package main

import (
  "fmt"
)

// this function will show every note of the user 
func (user User) getAllNotes()(dbNote [] Note,err error){
  db,err := getConnectionToMYSQL() 
  var str string = fmt.Sprintf("select IdUser,Title,Text from Notes where IdUser='%v'",user.Id)
  resp,err := db.Query(str)
  if err != nil {
    return dbNote,err
  }
  defer resp.Close()
  for resp.Next(){
    var pNote Note
    err = resp.Scan(&pNote.IdUser,&pNote.Title,&pNote.Text)
    if err != nil{
      return dbNote,err
    }
    dbNote = append(dbNote,pNote)
  }
    return dbNote,nil

}
func (user User) addANote(note Note)(err error){
  db,err := getConnectionToMYSQL() 
  str:= fmt.Sprintf("insert into Notes (IdUser,Title,Text) values (%v,'%s','%s')",user.Id,note.Title,note.Text)
  insert,err := db.Query(str)
  if err != nil{
    return err
  }
  defer insert.Close()
  return nil

}
func (user User) addToDataBase() (err error){
  db,err := getConnectionToMYSQL()
  str := fmt.Sprintf("insert into Users(Id,First,Second) values (%v,'%s','%s')",user.Id,user.First,user.Second)
  insert,err := db.Query(str)
  if err != nil {
    return err
  }
  defer insert.Close()
  return nil
}



