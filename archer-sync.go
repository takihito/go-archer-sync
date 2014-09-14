package main

import (
    "log"
    "fmt"
)

//    rsyncOption := RsyncOption{
//        Archive:  true,
//        Update:   true,
//        Compress: true,
//        Delete:   true,
//        Rsh:      "ssh",
//        Dry_run:  true,
//    }

func main() {

    archerSync, err := ParseConf("sync_test.yaml")

    if err != nil {
            log.Fatalf("error: %v", err)
    }

    msg := Rsync(archerSync)

    log.Printf("main:%v\n", archerSync)
    log.Printf("\nmsg:%s\n", msg)
    fmt.Printf(">>>>>>> WORK_DIR:%v\n", archerSync.Global.Work_dir)

}

