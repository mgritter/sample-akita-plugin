package main

import (
	"github.com/akitasoftware/akita-cli/plugin"
	"github.com/akitasoftware/akita-cli/printer"
	"github.com/akitasoftware/akita-libs/spec_util"
	pb "github.com/akitasoftware/akita-ir/go/api_spec"
	"strings"
)

type SillyAkitaPlugin struct{}

func (s SillyAkitaPlugin) Name() string {
	return "SillyAkitaPlugin"
}

func (s SillyAkitaPlugin) Transform(method *pb.Method) error {
	meta := spec_util.HTTPMetaFromMethod(method)
	if meta == nil {
		return nil
	}
	printer.Infof( "Bork! %v\n", meta.PathTemplate )
		
	path := strings.Split( meta.PathTemplate, "/" )
	for i := 1; i < len(path); i++ {
		path[i] = "bork"
	}
	meta.PathTemplate = strings.Join( path, "/" )
	return nil
}

func LoadAkitaPlugin() (plugin.AkitaPlugin, error) {
	return SillyAkitaPlugin{}, nil
}
