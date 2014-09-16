package main

import (
    "log"
    "fmt"
)

func main() {

    rsyncOption, err := ParseConf("sync_test.yaml")

    if err != nil {
            log.Fatalf("error: %v", err)
    }

    msg := Rsync(rsyncOption)
    log.Printf("main:%v\n", rsyncOption)
    log.Printf("\nmsg:%s\n", msg)
    fmt.Printf(">>>>>>> WORK_DIR")

}

