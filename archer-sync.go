package main

import (
    "log"
)

func main() {

    archerSync, err := ParseConf("sync_test.yaml")

    if err != nil {
            log.Fatalf("error: %v", err)
    }

    log.Printf("main:%v", archerSync)

}

