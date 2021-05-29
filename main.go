package main

import "runtime"
import "os"
import "github.com/gookit/color"
import "os/user"

func isRoot() bool {
  currentUser, err := user.Current()
  if err != nil {
    color.Error.Println("Unable to get current user: %s", err)
  }
  return currentUser.Username == "root"
}

func success(s string)  {
  color.Style{color.Green, color.Bold}.Println(s)
}
func error(s string)  {
  color.Style{color.Red, color.Bold}.Println(s)
}
func main() {

  if !(runtime.GOOS == "darwin" || runtime.GOOS == "linux"){
    color.Danger.Println("Shade supports only linux and macOSX")
    os.Exit(0)
  }

  if !(isRoot()){
    error("Shade must be run as a root user")
    os.Exit(0)
  }

  packages := os.Args[2:]
  subcommand := os.Args[1]

  switch subcommand {
  case "help":
    help()
  case "setup":
    setup()
  case "install":
    install(packages)
  case "uninstall":
    uninstall(packages)
  case "upgrade":
    upgrade(packages)
  case "update":
    update()
  case "query":
    query(os.Args[2])
  default: color.Info.Println("No subcommand called ", subcommand)
  }
}
