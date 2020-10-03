package ctr

import (
    "bufio"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
	"reflect"
)

func AnalyzeMovable() {

	var path string
    var id0 uint32
    var buf = make([]byte, 288)
    
    fmt.Println("3. Analyze movable.sed\n")
    
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
    
	fp.Read(buf)
    
	seed := []byte{0x53, 0x45, 0x45, 0x44} /* SEED */
    
    fmt.Print("\n")
    
	if reflect.DeepEqual(buf[0:4], seed) {
		keyY := fmt.Sprintf("%X", buf[272:288])
        
        fmt.Println("Friendcode seed:", keyY[0:16])
        fmt.Println("keyY           :", keyY)
        
        hash := sha256.Sum256([]byte(buf[272:288]))
        
        fmt.Print("ID0            : ")
        
        for i := 0; i <= 12; i += 4 {
            binary.Read(bytes.NewReader(hash[i:i+4]), binary.LittleEndian, &id0)
            fmt.Printf("%X", id0)
        }
        fmt.Print("\n")
    
    } else {
		fmt.Println("Error: Invalid file.")
	}
}