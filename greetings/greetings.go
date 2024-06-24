package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string, age int) string {
   // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!, I'm %d years old.", name,age)
return message
}




    
 
  