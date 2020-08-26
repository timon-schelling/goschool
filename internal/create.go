package internal

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/otiai10/copy"
	"github.com/timon-schelling/goschool/config"
)

func Create(parent string, template string, name string, prependDate bool) {
	directoryName := ""
	if prependDate {
		time := time.Now()
		directoryName += time.Format("2006-01-02")
	}
	if len(strings.TrimSpace(name)) != 0 {
		directoryName = spacing(directoryName, name)
	}

	directory := parent + "/" + directoryName
	templateDir := config.GoschoolHome + "/templates/" + template
	stat, err := os.Stat(templateDir)
	if os.IsNotExist(err) || !stat.IsDir() {
		return
	}
	stat, err = os.Stat(parent)
	if os.IsNotExist(err) || !stat.IsDir() {
		return
	}
	stat, err = os.Stat(directory)
	if !os.IsNotExist(err) {
		return
	}
	err = copy.Copy(templateDir, directory)
	if err != nil {
		log.Fatal(err)
	}
}

func spacing(s string, a string) string {
	if len(strings.TrimSpace(s)) != 0 {
		s += " "
	}
	s += a
	return s
}
