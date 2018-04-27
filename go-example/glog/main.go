package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	glog.Info("This is Info log")
	glog.Warning("This is a Warning log")
	glog.Error("this is a Error log")


	glog.V(0).Infoln("level 0")
	glog.V(1).Infoln("level 1")
	glog.V(2).Infoln("level 2")

	glog.Flush()
}
