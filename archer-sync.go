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
        parallel := make(chan bool, 4)
        log.Printf("\n\n\n\n msg:%v\n", servers["servers"])
        log.Printf("\nproject:%v\n", project)
go func() {
        for number, target  := range servers["servers"] {
            select {
            case para <- true:

log.Printf("\n >> serv:%s, i:%i, \n", target, number)
                go func(n int, server string) {
                    log.Printf("\n start serv:%s, msg:%s, n:%i, \n", server, msg, n)
    //            msg := Rsync(rsyncOption)
//    if ( n == 3 ) {
//        n = 7
//    }
    time.Sleep( time.Duration(n) * time.Second )
//    time.Sleep( time.Duration(1) * time.Second )
//                    log.Printf("\n end server:%s, msg:%s, n:%i, \n", servers["servers"][i], msg, n)
                    res := fmt.Sprintf("xyz:%s", server)
                    msg <- res
                    <- para
                }(number, target)
            }

            log.Printf("\nserver:%s\n", servers["servers"][number])
        }
}()
    for j := 0; j < len(servers["servers"]); j++ { 
        log.Printf("\nresult >>> msg:%s\n", <-msg)
    }




    }
    fmt.Printf(">>>>>>> WORK_DIR")

    log.Printf("main:%v\n", rsyncOption)
}

