package main

import "github.com/itchio/butler/comm"
import "github.com/itchio/wharf/archiver"

func untar(file string, dir string) {
	res, err := archiver.ExtractTar(file, dir, comm.NewStateConsumer())
	must(err)
	comm.Logf("Extracted %d dirs, %d files, %d symlinks", res.Dirs, res.Files, res.Symlinks)
}
