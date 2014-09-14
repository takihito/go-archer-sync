package main

import (
    "log"
    "fmt"
//    "os"
    "os/exec"
)

//    rsyncOption := RsyncOption{
//        Archive:  true,
//        Update:   true,
//        Compress: true,
//        Delete:   true,
//        Rsh:      "ssh",
//        Dry_run:  true,
//    }

//func Rsync(servers []string) (chan string, chan bool) {
func Rsync(archerSync *ArcherSync) (msg string) {
//    cmd := exec.Command("sleep", sleep)
    cmd := exec.Command("date")
    cmd.Args = []string {"-u"}
    stdout, err := cmd.Output()
    if err != nil {
        log.Fatalf("error: cmd.Output, %v", err)
    }


//    cmd.Stdout = os.Stdout
//    cmd.Stderr = os.Stderr
    cmd.Run()// この中で Start Waitを呼んでいる


    fmt.Printf(">>>>>>> WORK_DIR:%v\n", archerSync.Global.Work_dir)

    return string(stdout)
}

