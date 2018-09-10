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
  var totalCalls int = 100000
  var serviseURI string = "http://127.0.0.1"
  var statusCode int
  var total200 int
  var totalNo200 int

  if (len(os.Args) > 1) {
    totalCalls, _ = strconv.Atoi(os.Args[1])
  }

  if (len(os.Args) > 2) {
    serviseURI = os.Args[2]
  }

  for i := 0; i < totalCalls ; i++ {
    wg.Add(1)
    statusCode = runAttack(i, serviseURI, &wg)
    if (statusCode == 200){
      total200 += 1
    }else{
      totalNo200 += 1
    }
  }
  wg.Wait()


  fmt.Printf("Total calls: %d \n", totalCalls)
  fmt.Printf("Total 200 response: %d \n", total200)
  fmt.Printf("Total diferent from 200: %d \n", totalNo200)



}


func runAttack(index int, serviseURI string, wg *sync.WaitGroup) int {
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
  //fmt.Printf("code %d \n", resp.StatusCode)
  wg.Done()
  return resp.StatusCode
}
