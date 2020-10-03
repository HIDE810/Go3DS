package main

import (
    "./ctr"
    "bufio"
    "fmt"
    "os"
    "os/exec"
)

func clear() {
    cmd := exec.Command("cmd", "/C", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func main() {
    fmt.Println("Go3DS - 3DS Analyzing Tool")
    fmt.Println("(c) 2020 HIDE. All rights reserved.\n")
    
    fmt.Println("1. Check a 3gx version")
    fmt.Println("2. Analyze a serial numver")
    fmt.Println("3. Analyze movable.sed")
    
    fmt.Print("\n")
    
    for {
        fmt.Print("(Type 1 or 2 or 3) >> ")
        
        sc:= bufio.NewScanner(os.Stdin)
        sc.Scan()
        
        switch sc.Text() {
            case "1":
                clear()
                ctr.Check3gxVer()
                goto Run
            case "2":
                clear()
                ctr.AnalyzeSerial()
                goto Run
            case "3":
                clear()
                ctr.AnalyzeMovable()
                goto Run
        }
    }
    
    Run:
        fmt.Print("\nPress Enter to exit.")
        fmt.Scanln()
}