package main

import (
    "log"
    "fmt"
    "strings"
//    "os"
    "os/exec"
)
type RsyncOption struct {
    Archive bool
    Update bool
    Compress bool
    Verbose bool
    Delete bool
    Progress bool
    Dry_run bool
    Include []string
    Exclude []string
    Filter []string
    Rsh string
    User string
    Source string
    Dest string
}

func Rsync(archerSync *ArcherSync) (msg string) {

    rsyncOption := RsyncOption{
        Archive:  true,
        Update:   true,
        Compress: true,
        Verbose:  true,
        Delete:   true,
        Progress: true,
        Rsh:      "ssh",
        Dry_run:  true,
    }

    var WorkDir = archerSync.Global.Work_dir
    var DestDir = archerSync.Global.Dest_dir
    for _, value := range archerSync.Tasks.Process {
        if (value.Config.Dry_run == "1") {
            rsyncOption.Dry_run = true
        }
        if (value.Config.Archive == "1") {
            rsyncOption.Archive = true
        }
        if (value.Config.Compress == "1") {
            rsyncOption.Compress = true
        }
        if (value.Config.Verbose == "1") {
            rsyncOption.Verbose = true
        }
        if (value.Config.Update == "1") {
            rsyncOption.Update = true
        }
        if (value.Config.Delete == "1") {
            rsyncOption.Delete = true
        }
        if (value.Config.Progress == "1") {
            rsyncOption.Progress = true
        }
        if (len(value.Config.Include) > 0) {
            rsyncOption.Include = value.Config.Include
        }
        if (len(value.Config.Exclude)  > 0) {
            rsyncOption.Exclude = value.Config.Exclude
        }
        if (len(value.Config.Filter) > 0) {
            rsyncOption.Filter = value.Config.Filter
        }

        if (len(value.Config.Source) > 0) {
            rsyncOption.Source = value.Config.Source
        }
        if (len(value.Config.Dest) > 0) {
            rsyncOption.Dest = value.Config.Dest
        }
    }

//        if value, ok := value.Config["filter"]; ok {
//            rsyncOption.Filter = value
//        }

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
        rsyncCmd = append(rsyncCmd, "--dry_run" )
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

    rsyncOption.Source = strings.Replace(rsyncOption.Source, "[% work_dir %]", WorkDir, -1)
    rsyncCmd = append(rsyncCmd, fmt.Sprintf("%s", rsyncOption.Source))
    rsyncOption.Dest = strings.Replace(rsyncOption.Dest, "[% dest_dir %]", DestDir, -1)
    rsyncCmd = append(rsyncCmd, fmt.Sprintf("%s", rsyncOption.Dest))
    fmt.Printf("XXXXXXXXX %v\n", rsyncCmd )

    cmd := exec.Command("date")
    cmd.Args = []string {"-u"}
    stdout, err := cmd.Output()
    if err != nil {
        log.Fatalf("error: cmd.Output, %v", err)
    }
    cmd.Run()

    //fmt.Printf(">>>>>>> WORK_DIR:%v\n", archerSync.Global.Work_dir)
    return string(stdout)
}

