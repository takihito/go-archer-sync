package main

import (
        "io/ioutil"
//        "flag"
        "fmt"
        "log"
        "gopkg.in/yaml.v1"
)

type RsyncOption struct {
    ARCHIVE bool
    UPDATE bool
    COMPRESS bool
    DELETE bool
    EXCLUDE []string
    RSH string
    SOURCE string
    DEST string
    DRY_RUN bool
}

type ArcherSync struct {
    GLOBAL struct {
        WORK_DIR string
        DEST_DIR string
        ASSETS_PATH string
    }
    TASKS struct {
        INIT []struct {
            MODULE string
            NAME string
            CONFIG struct {
                MSG string
            }
        }
        PROCESS []struct {
            MODULE string
            NAME string
            CONFIG map[string] string
        }
    }
    PROJECTS map[string] map[string] []string
}

func ParseConf(yamlFile string) (*ArcherSync, error) {
    rsyncOption := RsyncOption{
        ARCHIVE:  true,
        UPDATE:   true,
        COMPRESS: true,
        DELETE:   true,
        RSH:      "ssh",
        DRY_RUN:  true,
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

   for _, value := range archerSync.TASKS.PROCESS {
        for option, ovalue := range value.CONFIG {
            fmt.Printf("option:%v value:%v\n", option, ovalue)
        }
    }

    for host, servers := range archerSync.PROJECTS {
        for _, server := range servers["servers"] {
            fmt.Printf("host:%s server:%s\n", host ,server)
        }
    }

    return &archerSync, err;
}

