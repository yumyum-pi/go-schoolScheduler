#!/bin/bash

#build directory
buildDir="./build"

#project name
projectName="schoolScheduler"


#project output filepath
output=$buildDir/$projectName


buildCMD=$(go build -o=$output)

if $buildCMD; then
    #make the executable
    chmod +x $output
    echo '> build:successful & executable'
else
    echo '> build:failed\n'
    exit 1
fi