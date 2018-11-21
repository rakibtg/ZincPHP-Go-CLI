package main
import "fmt"
import "log"
import "os/exec"
import "os"

func main() {
  system("php", "./test.php", "nice", os.Args[1])
}

func system(cmd string, arg ...string) {
  out, err := exec.Command(cmd, arg...).Output()
  if err != nil {
    log.Fatal(err)
  }
  // fmt.Println(os.Args)
  fmt.Println("🐳 🚀")
  fmt.Println(string(out))
}

func shiftArgs() {
  // Not completed yet...
}