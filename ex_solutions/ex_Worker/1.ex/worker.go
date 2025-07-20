package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
  w := 5
  dataArr := getData()
  
  chanJob := make(chan string,len(dataArr))
  
  for i:=0;i<len(dataArr);i++{
    chanJob <- dataArr[i]  
  }
  close(chanJob)
  
  wg := &sync.WaitGroup{}
  
  wg.Add(w)
  for i:=0;i<w;i++{
    go func(){
      defer wg.Done()
      worker(chanJob)
    }()
  }
  wg.Wait()
  
}

func getData() []string {
  return []string{"1", "2","3","4","5","6","7","8","9","10","11", "12","13","14","15","16","17","18","19","20"}
}

func checkDomain(host string) bool {
  return true
}

var count int 
func worker(job chan string) {
  for v := range job {
    bo := checkDomain(v)
	count++
    fmt.Print(bo,count)
    time.Sleep(time.Second * 1)
  }  
}