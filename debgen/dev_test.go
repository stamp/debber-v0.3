package debgen_test

import (
	"github.com/laher/debgo-v0.2/deb"
	"github.com/laher/debgo-v0.2/debgen"
	"log"
)

func Example_genDevPackage() {
	pkg := deb.NewPackage("testpkg", "0.0.2", "me", "Dummy package for doing nothing\n")

	ddpkg := deb.NewDevPackage(pkg)
	build := deb.NewBuildParams()
	build.IsRmtemp = false
	var err error
	ddpkg.InitDev()
	ddpkg.Dev.MappedFiles, err = debgen.GlobForGoSources(".", []string{build.TmpDir, build.DestDir})
	if err != nil {
		log.Fatalf("Error building -dev: %v", err)
	}

	err = debgen.GenDevArtifact(ddpkg, build)
	if err != nil {
		log.Fatalf("Error building -dev: %v", err)
	}

	// Output:
	//
}
