package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

var versionRxp, protoRxp *regexp.Regexp
var version, destDir string

func init() {
	var err error
	versionRxp, err = regexp.Compile(`^[0-9]+\.[0-9]+\.[0-9]+`)
	if err != nil {
		panic("cannot compile versionRxp")
	}
	protoRxp, err = regexp.Compile(`\.proto$`)
	if err != nil {
		panic("cannot compile protoRxp")
	}
	flag.StringVar(&version, "version", "3.10.0", "google protobuf version to download from github")
	flag.StringVar(&destDir, "destination", "proto", "parent path to the 'google' directory")
	flag.Parse()
}

func main() {
	sane := saneVersion(version)
	if !sane {
		log.Fatalf("%q is not a sane version number", version)
	}
	url := assembleURL(version)
	log.Printf("downloading: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error downloading %s: %s", url, err)
	}
	log.Print(resp.StatusCode)
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Fatalf("error reading gzip: %s", err)
	}
	tb := tar.NewReader(gz)

outer:
	for {
		header, err := tb.Next()
		switch {
		case err == io.EOF:
			break outer
		case err != nil:
			log.Fatalf("error reading tarball: %s", err)
		}
		if protoFilter(header.Name) {
			stripped, err := stripPath(header.Name)
			if err != nil {
				log.Fatal(err)
			}

			fqp := path.Join(destDir, stripped)
			dn, _ := path.Split(fqp)
			err = os.MkdirAll(dn, 0755)
			if err != nil {
				log.Fatalf("error creating directory %s: %s", dn, err)
			}
			f, err := os.Create(fqp)
			if err != nil {
				log.Fatalf("error creating file %s: %s", fqp, err)
			}
			defer f.Close()

			io.Copy(f, tb)
		}
	}
}

func stripPath(p string) (string, error) {
	splitPath := strings.Split(p, "/")
	if len(splitPath) < 3 {
		return "", fmt.Errorf("path %s too short to strip", p)
	}
	keep := splitPath[2:]
	return path.Join(keep...), nil
}

func protoFilter(path string) bool {
	split := strings.Split(path, "/")
	if len(split) < 3 {
		return false
	}
	match := protoRxp.Match([]byte(split[len(split)-1]))
	if split[1] == "src" &&
		split[2] == "google" &&
		match {
		return true
	}
	return false
}

func assembleURL(v string) string {
	url := strings.Join([]string{
		"https://github.com",
		"protocolbuffers",
		"protobuf",
		"releases",
		"download",
		"v" + v,
		assembleFilename(v)}, "/")
	return url
}

func assembleFilename(v string) string {
	return fmt.Sprintf("protobuf-all-%s.tar.gz", v)
}

func saneVersion(v string) bool {
	return versionRxp.Match([]byte(v))
}
