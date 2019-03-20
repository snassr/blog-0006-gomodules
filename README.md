# TLDR Go Modules example
Go Modules: Go's versioning system is out in go1.12.1. This is an quick overview example of using Go modules with some sub packages and external dependencies.

This repository was created to accompany an explanation and walk-through in the following blog post: https://snassr.io/posts/0006_tldrgomodules.

The repository depends on a example external dependency at https://github.com/snassr/blog-0006-listener.

## Summary
1. Modules are designed to keep track of the evolving dependencies in your go project.
2. Modules use [SemVer](ttps://semver.org/) to keep track of the versions of dependencies in your project.
3. Dependencies which are not versioned get a default tag.
    - special pseudo version tag: `v0.0.0-{DateTime}-{ShortCommitHash}` or `v0.0.0-yyyymmddhhmmss-commit`.
4. Integrated into the go binary! 
    - All go commands work with modules `go get, go get -u, go build ./..., go test ./...`
    - When inside a module, commands are aware they may need to pull in existing dependencies, add dependencies or update dependencies (`go get -u`).
5. Use any directory in the filesystem! (~/i/put/my/packages/here/HelloWorldPkg)
6. `go.mod` file keeps track of the dependencies and their versions (commit to repo). 
7. `go.sum` file tracks checksums and version history (commit to repo).
8. The GOPATH will likely be deprecated in the future...

## TLDR of TLDR
1. Create/convert project to use modules
```bash
# to add go modules to a new or existing repository
go mod init github.com/username/projectname
```
2. Sub packages work (tree branch example below)
    - projectname -- this is what you used in the `go mod init command` --> `module github.com/username/projectname`
        - cmd
            - main.go -- import sub packages using the import below
        - pkg
            - mypackage -- use this import throughout the project --> `import github.com/username/projectname/mypackage`
3. Verify in-use/downloaded dependencies
```bash
# verify checksums from go.sum against the modules being imported
go mod verify
```
4. Download dependencies
```bash
# in the context of a module (inside a module directory) 
#   it will download all dependencies based on go.mod
go get
```
5. Update dependencies
```bash
# in the context of a module (inside a module directory) 
#   will update all dependencies in go.mod
#   can specify a certain package to update after the -u flag
go get -u $OptionalPackageNameToUpdate
```