package main

import (
        "io/ioutil"
        "strings"
        "log"
        "gopkg.in/yaml.v1"
)

type RsyncOption struct {
    InitMessage string
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

type ArcherSync struct {
    Global GLOBAL
    Tasks TASKS
    Projects map[string] map[string] []string
}

type GLOBAL struct {
    Work_dir string
    Dest_dir string
    Assets_path string
}

type TASKS struct {
    Init INIT
    Process PROCESS
}

type INIT []struct {
    Module string
    Name string
    Config map[string] string
//    Config struct {
//        msg string
//    }
}

type PROCESS []struct {
    Module string
    Name string
    Config CONFIG
}

type CONFIG struct {
    User string
    Source string
    Dest string
    Dry_run string
    DryRun string
    Archive string
    Compress string
    Rsh string
    Update string
    Verbose string
    Delete string
    Progress string
    Include []string
    Exclude []string
    Filter []string
}

func ParseConf(yamlFile string) (*RsyncOption, map[string] map[string] []string, error) {
    archerSync := ArcherSync{}

    yamlString, err := ioutil.ReadFile(yamlFile)
    if err != nil {
        log.Fatalf("error: ioutil.ReadFile(yamlFile),  %v", err)
    }

    err = yaml.Unmarshal([]byte(yamlString), &archerSync)
    if err != nil {
        log.Fatalf("error: yaml.Unmarshal, %v", err)
    }

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
    for _, value := range archerSync.Tasks.Init {
        rsyncOption.InitMessage = value.Config["msg"]
    }
    for _, value := range archerSync.Tasks.Process {
        if (value.Config.Dry_run != "1") {
            rsyncOption.Dry_run = false
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

        if (len(value.Config.User) > 0) {
            rsyncOption.User = value.Config.User
        }

        if (len(value.Config.Source) > 0) {
            rsyncOption.Source = value.Config.Source
            rsyncOption.Source = strings.Replace(rsyncOption.Source, "[% work_dir %]", WorkDir, -1)
        }
        if (len(value.Config.Dest) > 0) {
            rsyncOption.Dest = value.Config.Dest
            rsyncOption.Dest = strings.Replace(rsyncOption.Dest, "[% dest_dir %]", DestDir, -1)
        }
    }

//    log.Printf("start parse config_yaml.")
    return &rsyncOption, archerSync.Projects, err;
}

