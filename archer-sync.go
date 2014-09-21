package main

import (
    "log"
    "fmt"
//    "sync"
    "time"
)

func main() {

    rsyncOption, projects, err := ParseConf("sync_test.yaml")

    if err != nil {
            log.Fatalf("error: %v", err)
    }

    for project, servers := range projects {
        msg := make(chan string)
        log.Printf("\nmsg:%v\n", servers["servers"])
        log.Printf("\nmsg:%v\n", project)
        for i := range servers["servers"] {
            go func(n int) {
//            msg := Rsync(rsyncOption)
                log.Printf("\n 0xxxxxxxxxx nserver:%s, msg:%i, \n", msg, n)
//if ( n == 3 ) {
//    n = 7
//}
time.Sleep( time.Duration(n) * time.Second )
                log.Printf("\n 1xxxxxxxxxx nserver:%s, msg:%s, \n", msg, n)
                res := "xyz"
                msg <- res
            }(i)
            log.Printf("\n+++++ 1 server:%s\n", servers["servers"][i])
        }

    for j := 0; j < len(servers["servers"]); j++ { 
        log.Printf("\n+++++ 2 msg:%s\n", <-msg)
    }




    }
    fmt.Printf(">>>>>>> WORK_DIR")

    log.Printf("main:%v\n", rsyncOption)
}

