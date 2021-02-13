package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn:=true
  fmt.Println(time.Now())
  eludedGuards := rand.Intn(100)
  fmt.Println(eludedGuards)
  if(eludedGuards>=50){
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  }else{
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }
  openedVault:=rand.Intn(100)
  fmt.Println(isHeistOn)
  if (openedVault>=70) {
    fmt.Println("Grab and GO!")
  }else if(isHeistOn) {
    isHeistOn = false
    fmt.Println("Vault can't be opened")
  }
  leftSafely:= rand.Intn(5)
  if(isHeistOn){
    switch leftSafely{
      case 0:
      isHeistOn = false
      fmt.Println("failed heist message")
      case 1:
      isHeistOn = false
      fmt.Print("Turns out vault doors don't open from the inside...")
      case 2 :
      isHeistOn = false
      fmt.Print("Turns out vault doors .")
      default:
      fmt.Println("Start the getaway car!")
    }
  }
  time.Now()
  if(isHeistOn){
    amtStolen:=10000 + rand.Intn(1000000)
    fmt.Println(amtStolen)
    fmt.Println(time.Now())
  }
  }
