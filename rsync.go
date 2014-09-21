package main

import (
    "log"
    "fmt"
//    "strings"
//    "os"
    "os/exec"
)


func Rsync(rsyncOption *RsyncOption) (msg string) {

    var rsyncCmd [] string
    rsyncCmd = append(rsyncCmd, "rsync" )
    if (rsyncOption.Archive) {
        rsyncCmd = append(rsyncCmd, "--archive" )
    }
    if (rsyncOption.Update) {
        rsyncCmd = append(rsyncCmd, "--update" )
    }
    if (rsyncOption.Compress) {
        rsyncCmd = append(rsyncCmd, "--compress" )
    }
    if (rsyncOption.Delete) {
        rsyncCmd = append(rsyncCmd, "--delete" )
    }
    if (rsyncOption.Dry_run) {
        rsyncCmd = append(rsyncCmd, "--dry-run" )
    }
    if (rsyncOption.Verbose) {
        rsyncCmd = append(rsyncCmd, "--verbose" )
    }
    if (rsyncOption.Progress) {
        rsyncCmd = append(rsyncCmd, "--progress" )
    }
    if (len(rsyncOption.Rsh) > 0) {
        rsyncCmd = append(rsyncCmd, fmt.Sprintf("--rsh=%s", rsyncOption.Rsh))
    }

    for i := range rsyncOption.Include {
        rsyncCmd = append(rsyncCmd, fmt.Sprintf("--include=\"%s\"", rsyncOption.Include[i]))
    }

    for i := range rsyncOption.Exclude {
        rsyncCmd = append(rsyncCmd, fmt.Sprintf("--exclude=\"%s\"", rsyncOption.Exclude[i]))
    }

    for i := range rsyncOption.Filter {
        rsyncCmd = append(rsyncCmd, fmt.Sprintf("--filter=\"%s\"", rsyncOption.Filter[i]))
    }

    rsyncCmd = append(rsyncCmd, fmt.Sprintf("%s", rsyncOption.Source))
    rsyncCmd = append(rsyncCmd, fmt.Sprintf("%s@%s", rsyncOption.User, rsyncOption.Dest))
    fmt.Printf("XXXXXXXXX %v\n", rsyncCmd )

    cmd := exec.Command("date")
    cmd.Args = []string {"-u"}
    stdout, err := cmd.Output()
    if err != nil {
        log.Fatalf("error: cmd.Output, %v", err)
    }
    cmd.Run()

    return string(stdout)
}

