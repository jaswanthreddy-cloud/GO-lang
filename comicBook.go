package main
import "fmt"
func main(){
  var writer, artist, title, publisher string = "Tracey Hatchet","Jewel Tampson","Mr. GoToSleep","DizzyBooks Publishing Inc."
  var year, pageNumber int = 1997,14
  var grade float32 = 6.5

  fmt.Println(title + " written by " +  writer + "and the name of artist is " + artist,  "publisher: " + publisher , "with a grade" , grade, " in the year " , year ,"total pages" , pageNumber)
  fmt.Println()

title="Epic Vol. 1"
writer= "Ryan N. Shawn"
artist= "Phoebe Paperclips"
year = 2013
pageNumber=160
grade = 9.0

  fmt.Println(title + " written by " +  writer + "and the name of artist is " + artist,  "publisher: " + publisher , "with a grade" , grade, " in the year " , year ,"total pages" , pageNumber)

}
