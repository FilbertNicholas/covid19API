package main

import (
	"fmt"
	"os"
)

var path = "covidActiveCases.json"

func createFile() {
    // check if file exists
    var _, err = os.Stat(path)

    // create file if not exists
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        if isError(err) {
            return
        }
        defer file.Close()
    }

    fmt.Println("File Created Successfully", path)
}

func writeFile(input string) {
	var fileInput  = input
    // Open file using READ & WRITE permission.
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if isError(err) {
        return
    }
    defer file.Close()

    // Write some text line-by-line to file.
    _, err = file.WriteString(fileInput)
    if isError(err) {
        return
    }

    // Save file changes.
    err = file.Sync()
    if isError(err) {
        return
    }

    fmt.Println("File Updated Successfully.")
}

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}
