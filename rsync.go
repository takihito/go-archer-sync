package main

import (
    "log"
    "fmt"
    "time"
    "strings"
    "os/exec"
)

func Rsync(rsyncOption *RsyncOption, project string, server string) (msg string) {

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

    source := strings.Replace(rsyncOption.Source, "[% project %]", project, -1)
    rsyncCmd = append(rsyncCmd, fmt.Sprintf("%s", source))
    dest := strings.Replace(rsyncOption.Dest, "[% server %]", server, -1)
    rsyncCmd = append(rsyncCmd, fmt.Sprintf("%s@%s", rsyncOption.User, dest))

fmt.Printf("command%v, project:%s, server:%s\n", rsyncCmd, project, server)
time.Sleep( time.Duration(2) * time.Second )

    cmd := exec.Command("date")
    cmd.Args = []string {"-u"}
    stdout, err := cmd.Output()
    if err != nil {
        log.Fatalf("error: cmd.Output, %v", err)
    }
    cmd.Run()
    result := string(stdout)

    return fmt.Sprintf("%s", result)
}

