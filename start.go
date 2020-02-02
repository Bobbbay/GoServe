package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {

  reader := bufio.NewReader(os.Stdin)
  fmt.Println(`
GoServe                                            
`)
  fmt.Println("---------------------")

  steps := []string{"Port", "Source"}

  for i:=0; i<len(steps);i++{
    fmt.Print(steps[i] + " >---> ")
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	num, err := strconv.Atoi(text)
	if err == nil && num <= 1024 {
		fmt.Println("We don't support low ports yet")
		steps = append([]string{"Prepend Item"}, steps...)
	}

	// At this point, we consider all is good
	settings, err := os.OpenFile("goserve.conf", os.O_APPEND|os.O_WRONLY, 0644)
	
	if err != nil {
        fmt.Println(err)
        return
    }
	
	write, err := settings.WriteString(text)
	
	if err != nil {
        fmt.Println(err)
        settings.Close()
        return
    }
	
	fmt.Println(write, "Settings saved")
	
	err = settings.Close()
	
	if err != nil {
        fmt.Println(err)
        return
    }

  }

}