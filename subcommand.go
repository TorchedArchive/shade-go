package main

import (
  "os"
  "log"
  "fmt"
  "bufio"
  "strings"
)

func rev(s string) (string){
  return "\033[7m" + s + "\033[0m";
}

func help(){
}

func setup()  {
}

func install(s []string){
}

func uninstall(s []string){
}

func upgrade(s []string){
}

func update(){
}

func query(s string){
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
      if strings.Contains(line, s){
        fmt.Println(strings.Replace(line, s, fmt.Sprintf("\033[7m%s\033[0m", s), 1))
      }
    }
  }
}
