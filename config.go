package main

import (
        "io/ioutil"
//        "flag"
//        "fmt"
        "log"
        "gopkg.in/yaml.v1"
)


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

func ParseConf(yamlFile string) (*ArcherSync, error) {
    archerSync := ArcherSync{}

    yamlString, err := ioutil.ReadFile(yamlFile)
    if err != nil {
        log.Fatalf("error: ioutil.ReadFile(yamlFile),  %v", err)
    }

    err = yaml.Unmarshal([]byte(yamlString), &archerSync)
    if err != nil {
        log.Fatalf("error: yaml.Unmarshal, %v", err)
    }

    log.Printf("start parse config_yaml.")
    return &archerSync, err;
}

