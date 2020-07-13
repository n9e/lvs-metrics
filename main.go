package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/weizhenqian/lvs-metrics/cron"
	"github.com/weizhenqian/lvs-metrics/g"
	"github.com/weizhenqian/lvs-metrics/http"
)

var cfg = flag.String("c", "cfg.json", "configuration file")
var version = flag.Bool("version", false, "show version")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	defer glog.Flush()
	flag.Parse()

	g.HandleVersion(*version)
	if memfile, _ := g.HandleMemProfile(*memprofile); memfile != nil {
		defer memfile.Close()
	}

	// global config
	g.ParseConfig(*cfg)
	g.InitRpcClients()

	cron.Collect()

	// http
	go http.Start()

	select {}
}
