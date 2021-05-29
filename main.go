package main

import (
  "os"
  "runtime"
)


func main() {
  if !(runtime.GOOS == "darwin" || runtime.GOOS == "linux") {
    err("Shade supports only linux and macOSX")
    os.Exit(1)
  }
  if len(os.Args) < 2 {
    help()
  }

  subcommand := os.Args[1]

  switch subcommand {
  case "help":
    help()
  case "setup":
    setup()
  case "install":
    packages := os.Args[2:]
    install(packages)
  case "uninstall":
    packages := os.Args[2:]
    uninstall(packages)
  case "upgrade":
    packages := os.Args[2:]
    upgrade(packages)
  case "update":
    update()
  case "query":
    query(os.Args[2])
  default:
    err("Unknown command:", subcommand)
    os.Exit(1)
  }
}

