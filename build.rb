#! /usr/bin/env ruby

Dir.chdir("#{ENV['GOPATH']}/src/github.com/camilo/fluid/parse")
`go tool yacc fluid.y`
abort "yacc failed" unless $?.success?
Dir.chdir("#{ENV['GOPATH']}/src/github.com/camilo/fluid/")
`go build -gcflags "-N -l"`
abort "build failed" unless $?.success?
