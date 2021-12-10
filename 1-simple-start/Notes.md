### Questions:

#### How to run the code?
  - We use the Go CLI to build, run, ... our project's code.
  
  - `go build` -> Compile a bunch of go source code files.
  - `go run` -> Compiles and executes one or two files.
  - `go fmt` -> Formats all the code in each file in the current directory.
  - `go install` -> Compiles and installs a package.
  - `go get` -> Doenloads the raw source code of someone else's package.
  - `go test` -> Runs any tests associated with the current project.

  - When we build a project, we just compile the source into an executable file (exe in case of windows).

#### What `package main` mean?
  - A package is just a project or a workspace that is comprised of multiple files.
  - These files, constituting a package, must have  `package {package name}` as the first line, decalring what package this file belongs to.
  - In Go, there are two types of packages, executable and reusable.
  1. An executable package generates a file we can run.
  2. A reusable package is meant to be used as helper, so it is a good place to put reusable logic.
  - The name of the package determines what kind of package we are building, so the name `main` is a unique package name. And the executable package file **must** have a function named `main`. 

#### What `import fmt` mean?
  - fmt is a standard library or package in the Go language, and importing it will add some functionality to our main package.
  - Check this link for docs about Go's standard libraries -> https://golang.org/pkg

#### What `func main () {}` mean?
  - This is the entry point of our package, so when the package is executed this function's code will be run.

#### How the `main.go` file is organized?
  - Package name, then the imports, and finally the code (functions, interfaces,....)

