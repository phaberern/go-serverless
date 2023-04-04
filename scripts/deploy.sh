#!/bin/bash
for dir in ../functions/*
do
    function_name="$(basename $dir)"  
    function_directory="$dir"  
    binary_file="../bin/${function_name}"
    zip_file="../bin/${function_name}.zip"

    echo "building $function_name..."
    GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o $binary_file $function_directory
    echo "zipping $function_name..."
    zip $zip_file $binary_file
done
