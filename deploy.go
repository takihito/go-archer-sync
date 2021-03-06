package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Rsync(rsyncOption *RsyncOption, project string, server string) (msg string) {

	var rsyncCmd []string
	if rsyncOption.Archive {
		rsyncCmd = append(rsyncCmd, "--archive")
	}
	if rsyncOption.Update {
		rsyncCmd = append(rsyncCmd, "--update")
	}
	if rsyncOption.Compress {
		rsyncCmd = append(rsyncCmd, "--compress")
	}
	if rsyncOption.Delete {
		rsyncCmd = append(rsyncCmd, "--delete")
	}
	if rsyncOption.DryRun {
		rsyncCmd = append(rsyncCmd, "--dry-run")
	}
	if rsyncOption.Verbose {
		rsyncCmd = append(rsyncCmd, "--verbose")
	}
	if rsyncOption.Progress {
		rsyncCmd = append(rsyncCmd, "--progress")
	}
	if len(rsyncOption.Rsh) > 0 {
		rsyncCmd = append(rsyncCmd, fmt.Sprintf("--rsh=%s", rsyncOption.Rsh))
	}

	for i := range rsyncOption.Include {
		rsyncCmd = append(rsyncCmd, fmt.Sprintf("--include"))
		rsyncCmd = append(rsyncCmd, fmt.Sprintf("%s", rsyncOption.Include[i]))
	}

	for i := range rsyncOption.Exclude {
		rsyncCmd = append(rsyncCmd, fmt.Sprintf("--exclude"))
		rsyncCmd = append(rsyncCmd, fmt.Sprintf("%s", rsyncOption.Exclude[i]))
	}

	for i := range rsyncOption.Filter {
		rsyncCmd = append(rsyncCmd, fmt.Sprintf("--filter"))
		rsyncCmd = append(rsyncCmd, fmt.Sprintf("%s", rsyncOption.Filter[i]))
	}

	rsyncCmd = append(rsyncCmd, "-v")

	rsyncCmd = append(rsyncCmd, rsyncOption.ProjectSource(project))
	rsyncCmd = append(rsyncCmd, rsyncOption.ServerDest(server))

	rsync_cli := strings.Join(rsyncCmd, " ")
	log.Printf("cmd:%s, server:%s, command%v\n", rsync_cli, project, server)

	cmd := exec.Command("rsync", rsyncCmd...)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatalf("error: deploy. server:%s, Output:%v", server, err)
	}

	result := string(stdout)
	return result
}
