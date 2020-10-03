package ctr

import (
    "bufio"
    "fmt"
    "os"
    "reflect"
)

func Check3gxVer() {
    
    var path string
    
    fmt.Println("1. Check a 3gx version\n")
    
    Input:
    
    for {
        fmt.Print("Enter the file path >> ")
        
        sc:= bufio.NewScanner(os.Stdin)
        sc.Scan()
        
        if len(sc.Text()) != 0 {
            path = sc.Text()
            break
        }
    }
    
    fp, err := os.Open(path)
    if err != nil {
        fmt.Println("\nError: The file path not found.\n")
        goto Input
    }
    defer fp.Close()
    
    buf := make([]byte, 8)
    
    fp.Read(buf)
    
    ver1 := []byte {0x33, 0x47, 0x58, 0x24, 0x30, 0x30, 0x30, 0x31} /* "3GX$0001" */
    ver2 := []byte {0x33, 0x47, 0x58, 0x24, 0x30, 0x30, 0x30, 0x32} /* "3GX$0002" */
    
    fmt.Print("\n")
    
    switch {
        case reflect.DeepEqual(buf, ver1): fmt.Println("CTRPF v0.5.1")
        case reflect.DeepEqual(buf, ver2): fmt.Println("CTRPF v0.6.0")
        default: fmt.Println("Error: Invalid file.")
    }
}