//Comic Mischief(Variables & Types)
//Use variables and types in the Go programming language to maintain the catalog of items in your comic book store Comic Mischief.

package main
import "fmt"

func main(){
  var publisher,writer,artist,title string
  title="Mr.GoToSleep"
  writer="Tracey Hatchet"
  artist="Jewel Tampson"
  publisher="DizzyBooks Publishing Inc"
  var year,pageNumber uint
  year=1997
  pageNumber=14
  var grade float32
  grade=6.5
  fmt.Println(title,"written by",writer,"drawn by",artist,"published by",publisher,"in the year of",year,"totl pages are",pageNumber,"and it is graded",grade)
  title="Epic Vol. 1"
  writer="Ryan N.Shawn"
  artist="Phoebe Paperclips"
  year=2013
  pageNumber=160
  grade=9.0
  fmt.Println(title,"written by",writer,"drawn by",artist,"published by",publisher,"in the year of",year,"totl pages are",pageNumber,"and it is graded",grade)
}
