package greetings

import (
  "fmt"
  "testing"
  "regexp"
)

/*
  Two test function 
  1. TestHelloName : send the input as name and match with expected behavior 
  2. TestHelloEmpty : Send empty string as name and match with expected behavior``
*/

func TestHelloName(t *testing.T) {
  name := "Jimmi"
  want := regexp.MustCompile(`\b`+name+`\b`)
  message, err := Hello("Jimmi")
  if !want.MatchString(message) || err != nil {
    t.Fatalf(`Hello("Jimmi") = %q, %v, want match for %#q, nil`, message, err, want)
  }

}

func TestHelloEmpty(t *testing.T) {
  message, err := Hello("")
   if message != "" {
    for i := 0;i < len(message);i++ {
      fmt.Printf("The ASCII value of %c is %d \n", message[i], message[i]) 
    }
  }
  
  if message != "" || err == nil {
    t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
  }
}

