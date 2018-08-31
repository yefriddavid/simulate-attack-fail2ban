package main

import (
  "fmt"
  "sync"
  "net/http"
  "io/ioutil"
)

func main(){
	var wg sync.WaitGroup
  for i := 0; i < 100000 ; i++ {
  //for i := 0; i < 100 ; i++ {
    wg.Add(1)
    runAttack(i, &wg)
  }
  wg.Wait()
}


func runAttack(index int, wg *sync.WaitGroup) bool {
	serviseURI := `http://127.0.0.1`

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
  //fmt.Println(string(body))
  //fmt.Println(index)
  wg.Done()
  return true
}
