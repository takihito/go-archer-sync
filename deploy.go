package main

import (
    "log"
    "fmt"
    "strings"
    "os/exec"
)

func Rsync(rsyncOption *RsyncOption, project string, server string) (msg string) {

    var rsyncCmd [] string
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

/* includeとexcludeが有効にならない? */
    for i := range rsyncOption.Include {
        rsyncCmd = append(rsyncCmd, fmt.Sprintf("--include='%s'", rsyncOption.Include[i]))
    }

    for i := range rsyncOption.Exclude {
        rsyncCmd = append(rsyncCmd, fmt.Sprintf("--exclude='%s'", rsyncOption.Exclude[i]))
    }

/*
    filterは原因不明のエラーとなる
    for i := range rsyncOption.Filter {
        rsyncCmd = append(rsyncCmd, fmt.Sprintf("--filter=\"%s\"", rsyncOption.Filter[i]))
    }
*/
    source := strings.Replace(rsyncOption.Source, "[% project %]", project, -1)
    rsyncCmd = append(rsyncCmd, fmt.Sprintf("%s", source))
    dest := strings.Replace(rsyncOption.Dest, "[% server %]", server, -1)
    rsyncCmd = append(rsyncCmd, fmt.Sprintf("%s@%s", rsyncOption.User, dest))

    log.Printf("project:%s, server:%s, command%v\n", rsyncCmd, project, server)

    cmd := exec.Command("rsync", rsyncCmd...)
    stdout, err := cmd.Output()
    if err != nil {
        log.Fatalf("error: deploy. server:%s, Output:%v", server, err)
    }
    cmd.Run()

    result := string(stdout)

    return result
}

