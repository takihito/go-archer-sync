package main

import (
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"strings"
)

type RsyncOption struct {
	InitMessage string
	Archive     bool
	Update      bool
	Compress    bool
	Verbose     bool
	Delete      bool
	Progress    bool
	DryRun      bool
	Include     []string
	Exclude     []string
	Filter      []string
	Rsh         string
	User        string
	Source      string
	Dest        string
}

func (r *RsyncOption) ProjectSource(project string) string {
	source := strings.Replace(r.Source, "[% project %]", project, -1)
	return source
}

func (r *RsyncOption) ServerDest(server string) string {
	dest := strings.Replace(r.Dest, "[% server %]", server, -1)
	return dest
}

type ArcherSync struct {
	Global   GLOBAL
	Tasks    TASKS
	Projects map[string]map[string][]string
}

type GLOBAL struct {
	WorkDir string `yaml:"work_dir"`
	DestDir string `yaml:"dest_dir"`
}

type TASKS struct {
	Init    INIT
	Process PROCESS
}

type INIT []struct {
	Module string
	Name   string
	Config map[string]string
}

type PROCESS []struct {
	Module string
	Name   string
	Config CONFIG
}

type CONFIG struct {
	User   string
	Source string
	Dest   string
	DryRun string `yaml:dry_run`
	//    DryRun string
	Archive  string
	Compress string
	Rsh      string
	Update   string
	Verbose  string
	Delete   string
	Progress string
	Include  []string
	Exclude  []string
	Filter   []string
}

func ParseConf(yamlFile string) (*RsyncOption, map[string]map[string][]string, error) {
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
		DryRun:   true,
	}

	var WorkDir = archerSync.Global.WorkDir
	var DestDir = archerSync.Global.DestDir
	for _, value := range archerSync.Tasks.Init {
		rsyncOption.InitMessage = value.Config["msg"]
	}
	for _, value := range archerSync.Tasks.Process {
		if value.Config.DryRun != "1" {
			rsyncOption.DryRun = false
		}
		if value.Config.Archive == "1" {
			rsyncOption.Archive = true
		}
		if value.Config.Compress == "1" {
			rsyncOption.Compress = true
		}
		if value.Config.Verbose == "1" {
			rsyncOption.Verbose = true
		}
		if value.Config.Update == "1" {
			rsyncOption.Update = true
		}
		if value.Config.Delete == "1" {
			rsyncOption.Delete = true
		}
		if value.Config.Progress == "1" {
			rsyncOption.Progress = true
		}
		if len(value.Config.Include) > 0 {
			rsyncOption.Include = value.Config.Include
		}
		if len(value.Config.Exclude) > 0 {
			rsyncOption.Exclude = value.Config.Exclude
		}
		if len(value.Config.Filter) > 0 {
			rsyncOption.Filter = value.Config.Filter
		}

		if len(value.Config.User) > 0 {
			rsyncOption.User = value.Config.User
		}

		if len(value.Config.Source) > 0 {
			rsyncOption.Source = value.Config.Source
			rsyncOption.Source = strings.Replace(rsyncOption.Source, "[% work_dir %]", WorkDir, -1)
		}
		if len(value.Config.Dest) > 0 {
			rsyncOption.Dest = value.Config.Dest
			rsyncOption.Dest = strings.Replace(rsyncOption.Dest, "[% dest_dir %]", DestDir, -1)
		}
	}

	return &rsyncOption, archerSync.Projects, err
}
