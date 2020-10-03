package ctr

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var serial []string

func AnalyzeConsole() {
    
    var model, region string
    
    for {
        fmt.Print("Serial >> ")
        
        sc:= bufio.NewScanner(os.Stdin)
        sc.Scan()
        
        if len(sc.Text()) == 11 {
            serial = strings.Split(strings.ToUpper(sc.Text()), "")
            break
        } else if len(sc.Text()) != 0 {
            fmt.Println("\nError: Not a serial number\n")
        }
    }
    
    switch serial[0] {
        case "C": model = "3DS"
        case "S": model = "3DS XL/LL"
        case "A": model = "2DS"
        case "Y": model = "New 3DS"
        case "Q": model = "New 3DS XL/LL"
        case "N": model = "New 2DS XL/LL"
        default: model = "Unknown"
    }
    
    switch serial[1] {
        case "J": region = "Japan"
        case "W": region = "North America"
        case "S": region = "Middle East/Southeast Asia"
        case "E": region = "Europe"
        case "A": region = "Australia"
        case "K": region = "South Korea"
        case "C": region = "China"
        default: region = "Unknown"
    }
    
    fmt.Print("\n")
    
    fmt.Println("Model :", model)
    fmt.Println("Region:", region)
}

func AnalyzeGameCard() {
    
    var month string
    
    for {
        fmt.Print("Serial >> ")
        
        sc:= bufio.NewScanner(os.Stdin)
        sc.Scan()
        
        if len(sc.Text()) == 10 {
            serial = strings.Split(strings.ToUpper(sc.Text()), "")
            break
        } else if len(sc.Text()) != 0 {
            fmt.Println("\nError: Not a serial number\n")
        }
    }
    
    fmt.Print("\n")
    
    fmt.Printf("Product code : CTR-P-%s\n", strings.Join(serial[0:4], ""))
    
    year, _ := strconv.Atoi(serial[5])
    
    if year >= 1 && year <= 9 {
        fmt.Printf("Product year : %d\n", 2010 + year)
    } else {
        fmt.Println("Product year : Unknown")
    }
    
    switch serial[4] {
        case "1": month = "Jan."
        case "2": month = "Feb."
        case "3": month = "Mar."
        case "4": month = "Apr."
        case "5": month = "May"
        case "6": month = "Jun."
        case "7": month = "Jul."
        case "8": month = "Aug."
        case "9": month = "Sep."
        case "X": month = "Oct."
        case "Y": month = "Nov."
        case "Z": month = "Dec."
        default: month = "Unknown"
    }
    
    fmt.Println("Product month:", month)
}

func AnalyzeSerial() {
    
    fmt.Println("2. Analyze a serial numver\n")
    fmt.Println("Which type you want to analyze? (1.Console, 2.Gamecard)\n")
    
    for {
        fmt.Print("(Type 1 or 2) >> ")
        
        sc:= bufio.NewScanner(os.Stdin)
        sc.Scan()
        
        switch sc.Text() {
            case "1":
                fmt.Print("\n")
                AnalyzeConsole()
                goto Run
            case "2":
                fmt.Print("\n")
                AnalyzeGameCard()
                goto Run
        }
    }
    Run:
}