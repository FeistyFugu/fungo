package main

import (
	"flag"
	"fmt"
	tpl "github.com/FeistyFugu/fungo/templates"
	"os"
	"path"
	"strings"
	"text/template"
	"time"
)

const version = "0.0.1"

const (
	_ = iota
	exitCodeArgRequired
	exitCodeInvalidArg
	exitCodeUnknownTemplate
	exitCodeFailedToGenerate
	exitCodeOther
)

type templateArgs struct {
	Version string
	Date string
	PackageName string
	FunctionName string
	T1 string
	T2 string
}

func crash(code int, data ...interface{}) {
	var errMsg = map[int]string {
		exitCodeArgRequired: "Argument required: %s",
		exitCodeInvalidArg: "Invalid %s: '%s'",
		exitCodeUnknownTemplate: "Unknown template: %s",
		exitCodeFailedToGenerate: "Could not generate source file: %s",
		exitCodeOther: "Failed! The following error occurred: %s",
	}

	fmt.Printf(errMsg[code], data...)
	os.Exit(code)
}

func isValidName(name string) bool {
	alpha := "abcdefghijklmnopqrstuvwxyz_"
	numbers := "0123456789"
	if name == "" {
		return false
	}
	if !strings.Contains(alpha, strings.ToLower(string(name[0]))) {
		return false
	}
	valid := alpha + numbers
	for i := 1; i < len(name); i++ {
		if !strings.Contains(valid, strings.ToLower(string(name[i]))) {
			return false
		}
	}
	return true
}

func main() {
	var templates = map[string]string {
		"All": tpl.All,
		"Apply": tpl.Apply,
		"Contains": tpl.Contains,
		"FanOut": tpl.FanOut,
		"Filter": tpl.Filter,
		"FindFirst": tpl.FindFirst,
		"FindLast": tpl.FindLast,
		"GroupBy": tpl.GroupBy,
		"Match": tpl.Match,
		"Max": tpl.Max,
		"Min": tpl.Min,
		"Reduce": tpl.Reduce,
	}

	tplPtr := flag.String("template", "", "Name of the function template to invoke")
	pkgNamePtr := flag.String("packageName", "", "Name of the package to which the function must be added")
	funcNamePtr := flag.String("functionName", "", "Name of the function to create from the template")
	fileNamePtr := flag.String("fileName", "", "Name of the source file to create")
	t1Ptr := flag.String("t1", "string", "Name of the first type")
	t2Ptr := flag.String("t2", "string", "Name of the second type")
	flag.Parse()

	if *tplPtr == "" {
		crash(exitCodeArgRequired, "template")
	}
	tplName := *tplPtr
	if _, exist := templates[tplName]; !exist {
		crash(exitCodeUnknownTemplate, tplName)
	}

	var pkgName string
	if *pkgNamePtr == "" {
		dir, err := os.Getwd()
		if err != nil {
			crash(exitCodeOther, err.Error())
		}
		pkgName = path.Base(dir)
	} else {
		pkgName = *pkgNamePtr
	}
	if !isValidName(pkgName) {
		crash(exitCodeInvalidArg, "packageName", pkgName)
	}

	var funcName string
	if *funcNamePtr == "" {
		funcName = tplName
	} else {
		funcName = *funcNamePtr
	}
	if !isValidName(funcName) {
		crash(exitCodeInvalidArg, "functionName", funcName)
	}

	var fileName string
	if *fileNamePtr == "" {
		fileName = funcName
	} else {
		fileName = *fileNamePtr
	}
	if !isValidName(fileName) {
		crash(exitCodeInvalidArg, "fileName", fileName)
	}
	fileName += ".go"

	t1 := *t1Ptr
	if !isValidName(t1) {
		crash(exitCodeInvalidArg, "t1", t1)
	}

	t2 := *t2Ptr
	if !isValidName(t2) {
		crash(exitCodeInvalidArg, "t2", t2)
	}

	f, err := os.Create(fileName)
	if err != nil {
		crash(exitCodeFailedToGenerate, err.Error())
	}

	t := template.Must(template.New(tplName).Parse(tpl.Header + templates[tplName]))
	ta := templateArgs{version, time.Now().Format("2006-01-02 15:04:05"), pkgName, funcName, t1, t2}
	err = t.Execute(f, ta)
	if err != nil {
		crash(exitCodeFailedToGenerate, err)
	}
}
