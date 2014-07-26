package main

import (
	"github.com/debber/debber-v0.3/deb"
	"github.com/debber/debber-v0.3/debgen"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func genChangelog(args []string) {
	pkg := deb.NewPackage("", "", "", "")
	build := debgen.NewBuildParams()
	debgen.ApplyGoDefaults(pkg)
	fs, params := InitFlags(cmdName, pkg, build)
	fs.StringVar(&params.Architecture, "arch", "all", "Architectures [any,386,armhf,amd64,all]")
	var entry string
	fs.StringVar(&entry, "entry", "", "Changelog entry data")

	err := ParseFlags(pkg, params, fs)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if entry == "" {
		log.Fatalf("Error: --entry is a required flag")

	}
	filename := filepath.Join(build.DebianDir, "changelog")
	templateVars := debgen.NewTemplateData(pkg)
	templateVars.ChangelogEntry = entry
	err = os.MkdirAll(filepath.Join(build.ResourcesDir, "debian"), 0777)
	if err != nil {
		log.Fatalf("Error making dirs: %v", err)
	}

	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		tpl, err := template.New("template").Parse(debgen.TemplateChangelogInitial)
		if err != nil {
			log.Fatalf("Error parsing template: %v", err)
		}
		//create ..
		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Error creating file: %v", err)
		}
		defer f.Close()
		err = tpl.Execute(f, templateVars)
		if err != nil {
			log.Fatalf("Error executing template: %v", err)
		}
		err = f.Close()
		if err != nil {
			log.Fatalf("Error closing written file: %v", err)
		}
	} else if err != nil {
		log.Fatalf("Error reading existing changelog: %v", err)
	} else {
		tpl, err := template.New("template").Parse(debgen.TemplateChangelogAdditionalEntry)
		if err != nil {
			log.Fatalf("Error parsing template: %v", err)
		}
		//append..
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer f.Close()
		err = tpl.Execute(f, templateVars)
		if err != nil {
			log.Fatalf("Error executing template: %v", err)
		}
		err = f.Close()
		if err != nil {
			log.Fatalf("Error closing written file: %v", err)
		}
	}

}

