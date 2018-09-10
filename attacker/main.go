package main

import (
  "fmt"
  "os"
  "sync"
  "strconv"
  "net/http"
  "io/ioutil"
)

func main(){
	var wg sync.WaitGroup
  var totalCalls int = 1000000
  var serviseURI string = "http://127.0.0.1"

  if (len(os.Args) > 1) {
    totalCalls, _ = strconv.Atoi(os.Args[1])
  }

  if (len(os.Args) > 2) {
    serviseURI = os.Args[2]
  }
  for i := 0; i < totalCalls ; i++ {
    wg.Add(1)
    runAttack(i, serviseURI, &wg)
  }
  wg.Wait()
}


func runAttack(index int, serviseURI string, wg *sync.WaitGroup) bool {
  req, err := http.NewRequest("GET", serviseURI, nil)
  req.Header.Set("Content-Type", "text/xml")

  client := &http.Client{}
  resp, err := client.Do(req)

  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Printf("%d) %s \n", index, string(body))
  wg.Done()
  return true
}
