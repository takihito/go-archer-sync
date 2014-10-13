package main

import (
	. "."
	"testing"
)

func TestParse(t *testing.T) {
	option, projects, err := ParseConf("deploy_config.yaml")
	if err != nil {
		t.Error(err)
	}

	if option.InitMessage != "really deploy app? [y/n]" {
		t.Error(option.InitMessage)
	}

	if option.Dest != "[% server %]:/home/deploy/" {
		t.Error(option.Dest)
	}

	if option.Source != "/home/deploy_trunk/[% project %]" {
		t.Error(option.Source)
	}

	for project, _ := range projects {
		if project != "example.com" && project != "xyz.localhost" {
			t.Error(project)
		}
	}
}
