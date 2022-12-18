//LEARN GO
//Bank Heist

package main
import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn:=true
  eludedGuards:=rand.Intn(100)
  if eludedGuards>=50{
    fmt.Println("Good Job!")
  }else{
    isHeistOn=false
    fmt.Println("Need a better a plan.")
  }
  openedVault:=rand.Intn(100)
  if openedVault>=70{
    fmt.Println("Grab and Go!")
  }else if openedVault<70{
    isHeistOn=false
    fmt.Println("Vault can't be Opened")
  }
  leftSafely:=rand.Intn(3)
  if isHeistOn{
    switch leftSafely{
      case 0:
      isHeistOn=false
      fmt.Println("Alarmed")
      case 1:
      isHeistOn=false
      fmt.Println("Fignerprint not accepting...")
      case 2:
      isHeistOn=false
      fmt.Println("They have Dogs in the Vaults")
      default:
      fmt.Println("Nothing to worry!")
    }
    if isHeistOn{
      amtStolen:=10000+rand.Intn(1000000)
      fmt.Println("We've taken:",amtStolen)
    }
  }
  fmt.Println("isHeistOn is currently:",isHeistOn)
}
