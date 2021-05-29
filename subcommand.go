package main

import (
  "os"
  "log"
  "fmt"
  "bufio"
  "strings"
)

func help() {
	// TODO
	os.Exit(0)
}

func setup()  {
  checkRoot()

  info("Starting setup")
  setupDirs := []string{"log", "cache", "bin", "include", "lib", "share"}

  for _, curdir := range setupDirs {
    os.MkdirAll(shadeDir + curdir, 0755)
  }

  run("git clone " + shadeRepo + " " + shadeDir + "main")
  success("Setup done!")
}

func install(s []string) {
	// TODO
	checkRoot()
}

func uninstall(s []string) {
	// TODO
	checkRoot()
}

func upgrade(s []string) {
	// TODO
	checkRoot()
}

func update() {
	// TODO
	checkRoot()
}

func query(s string) {
  buildscripts := "/usr/local/shade/main/packages"
  files, err := os.ReadDir(buildscripts)
  if err != nil {
    log.Fatal(err)
  }

  for _, file := range files {
    path := fmt.Sprintf("%s/%s/meta", buildscripts, file.Name())
    inFile, err := os.Open(path)
    if err != nil {
      fmt.Println(err.Error() + `: ` + path)
      return
    }
    defer inFile.Close()

    scanner := bufio.NewScanner(inFile)
    for scanner.Scan() {
      line := scanner.Text()
      if strings.Contains(line, s) {
        fmt.Println(strings.Replace(line, s, fmt.Sprintf("\033[7m%s\033[0m", s), 1))
      }
    }
  }
}
