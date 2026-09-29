package main

import (
        "errors"
        "fmt"
        "os"
        "time"
)

func main() {
        if len(os.Args) == 1 {
                f, e := os.ReadFile(".wip")
                if e != nil {
                        if errors.Is(e, os.ErrNotExist) {
                                _, e := os.Create(".wip")
                                if e != nil {
                                        println("error creating wip file: ", e.Error())
                                        return
                                }
                        }
                }
                println(string(f))
                return
        }

        if len(os.Args) == 2 && os.Args[1] != "ok" {
                now := time.Now().Local()
                res := fmt.Sprintf("[%s] %s", now.Format("2006/01/02 15:04"), os.Args[1])
                e := os.WriteFile(".wip", []byte(res), os.FileMode(0644))
                if e != nil {
                        println("error writing wip file: ", e.Error())
                        return
                }
                println(res)
                return
        }
}
