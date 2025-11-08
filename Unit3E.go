// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-06
// Fileoverview: This program asks the user for their first name and then
// displays the lyrics to the song "Happy Birthday," inserting the user's
// name at the appropriate place in the song.

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // variables
    var userName string

    reader := bufio.NewReader(os.Stdin)

    // input
    fmt.Print("What is your first name? ")
    userName, _ = reader.ReadString('\n')
    userName = strings.TrimSpace(userName)

    // output
    fmt.Println()
    fmt.Println("Here is your birthday song:")
    fmt.Println()
    fmt.Println("Happy Birthday to you,")
    fmt.Println("Happy Birthday to you,")
    fmt.Printf("Happy Birthday dear %s,\n", userName)
    fmt.Println("Happy Birthday to you!")
    fmt.Println()
    fmt.Println("Done.")
}
