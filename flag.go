package main

import (
	"flag"
	"os"
)

func bootFlag() {

	flag.Parse()

	if len(flag.Args()) == 0 {
		runFailed()
	}

	if flag.Arg(0) == "init" {

		dir := flag.Arg(1)

		if len(dir) == 0 {
			runFailed()
		} else if pathExists(dir) {
			pp("Dir `" + dir + "` already exists, use different name!")
		}

		handleFlagInit(dir)
		os.Exit(0)
	}

	iniFilePath = flag.Arg(0)

	if !fileExists(iniFilePath) {
		ps("The ini file `" + iniFilePath + "` doesn't exist!")
		pm("To auto-generate ini file, run  -")
		pe("$ blog-to-pdf init")
		os.Exit(0)
	}

}

func runFailed() {
	ps("\n+\tWrong instruction given!")
	pe(ConstHelpStr)
	pp("")
}

func handleFlagInit(dir string) {

	dir = AZ_AND_NUMBER_ONLY(dir)

	createDirsIfNotExists([]string{dir})

	p("Created directory: " + dir)

	var iniFN = dir + "/" + "config.ini"
	var strReplaceFN = dir + "/" + "string_replaces.json"
	var cssFN = dir + "/" + "custom.css"

	f, err := os.Create(iniFN)

	if err != nil {
		pp("Error creating `" + iniFN + "` file: " + err.Error())
	}

	p("Created ini file: " + iniFN)

	f2, err := os.Create(strReplaceFN)

	if err != nil {
		pp("Error creating `" + strReplaceFN + "` file: " + err.Error())
	}

	p("Created string replace file: " + strReplaceFN)

	f3, err := os.Create(cssFN)

	if err != nil {
		pp("Error creating `" + cssFN + "` file: " + err.Error())
	}

	p("Created custom css file: " + cssFN)

	defer f.Close()
	defer f2.Close()
	defer f3.Close()

	f.WriteString(ConstSampleINI)
	f2.WriteString(ConstReplaces)
	f3.WriteString(ConstCusotmCss)

	var instruction = `
+		Now -
+		$ cd ` + dir + `
+		Edit ` + iniFN + ` according to your needs, then -
+		$ blog-to-pdf config.ini
`
	p(instruction)
}

// func getAFileName(baseName string, ext string, i int) string {
// 	if !fileExists(baseName + ext) {
// 		return baseName + ext
// 	} else if !fileExists(baseName + "_" + strconv.Itoa(i) + ext) {
// 		return baseName + "_" + strconv.Itoa(i) + ext
// 	}
// 	return getAFileName(baseName, ext, i+1)
// }