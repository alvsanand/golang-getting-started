package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func readTextFromConsole(message string) (string, error) {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print(message)
    text, err := reader.ReadString('\n')
    
    return text, err
}

var (
    floatingNumberREGEX = regexp.MustCompile(`^(\d+)\.?(\d*)$`)
)

// GetIntegerFromFloatingNumber Parse a FloatingNumber from a string and obtain the integer part
func GetIntegerFromFloatingNumber(number string) (int64, error) {
    matched := floatingNumberREGEX.FindStringSubmatch(number)

    if len(matched) == 0 {
        return 0, fmt.Errorf("Error parsing: %s is not a Floating Number", number)
    }
    
    return strconv.ParseInt(matched[1], 10, 64)
}

func main() {
     //number, err := readTextFromConsole("floating point number: ")
    number := "123.45"    
    var err error

    if err != nil {
        log.Fatal("Error reading text from console: ", err)
    } else {
        var integer, err = GetIntegerFromFloatingNumber(strings.Trim(number, "\n"))

        if err != nil {
            log.Fatal("Error parsing FloatingNumber: ", err)
        } else {
            fmt.Println(fmt.Sprintf("Integer parsed: %d", integer))
        }        
    }
}