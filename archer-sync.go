package main

import (
    "log"
    "flag"
    "fmt"
    "os"
)

func main() {

    var parallelFlag int
    var yamlConfigFlag string
    flag.IntVar(&parallelFlag, "parallel", 2, "parallel worker")
    flag.StringVar(&yamlConfigFlag, "config", "deploy_config.yaml", "yaml confige file")
    flag.Parse()

    rsyncOption, projects, err := ParseConf(yamlConfigFlag)
    if err != nil {
            log.Fatalf("error: %v", err)
    }

    var yn string
    fmt.Printf("> %s:", rsyncOption.InitMessage)
    fmt.Scan(&yn)
    if (yn != "y") {
        os.Exit(0)
    }

    for project, servers := range projects {
        msg := make(chan string)
        parallel := make(chan bool, parallelFlag)
        go func() {
                for number, target := range servers["servers"] {
                    select {
                    case parallel <- true:
                        go func(n int, server string) {
                            log.Printf("start deploy. project:%s, server:%s\n",project, server)
                            result := Rsync(rsyncOption, project, server)
                            msg <- result
                            <- parallel
                        }(number, target)
                    }
                }
        }()
        for j := 0; j < len(servers["servers"]); j++ {
            log.Printf("result\n%s\n", <-msg)
        }
    }
    log.Printf("...finished\n")
}

