package main

import (
        "io/ioutil"
//        "flag"
        "fmt"
        "log"
        "gopkg.in/yaml.v1"
)

type RsyncOption struct {
    Archive bool
    Update bool
    Compress bool
    Delete bool
    Exclude []string
    Rsh string
    Source string
    Dest string
    Dry_run bool
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
    Config struct {
        msg string
    }
}

type PROCESS []struct {
    Module string
    Name string
    Config map[string] string
}

func ParseConf(yamlFile string) (*ArcherSync, error) {
    rsyncOption := RsyncOption{
        Archive:  true,
        Update:   true,
        Compress: true,
        Delete:   true,
        Rsh:      "ssh",
        Dry_run:  true,
    }
    archerSync := ArcherSync{}

//    flag.StringVar(&yamlFile, "config", "sync.yaml", "yaml file")
    yamlString, _ := ioutil.ReadFile(yamlFile)
    log.Printf("start deploy project:%s", yamlString)
    log.Printf("start deploy project:%v", rsyncOption)

    err := yaml.Unmarshal([]byte(yamlString), &archerSync)
    if err != nil {
            log.Fatalf("error: %v", err)
    }
    fmt.Printf(">>>>>>> WORK_DIR:%v\n", archerSync.Global.Work_dir)

    for _, value := range archerSync.Tasks.Process {
        for option, ovalue := range value.Config {
            fmt.Printf("option:%v value:%v\n", option, ovalue)
        }
    }

    for host, servers := range archerSync.Projects {
        for _, server := range servers["servers"] {
            fmt.Printf("host:%s server:%s\n", host ,server)
        }
    }

    return &archerSync, err;
}

