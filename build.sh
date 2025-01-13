#!/bin/bash

targets=("linux/amd64" "darwin/amd64" "windows/amd64")

for target in "${targets[@]}"; do
  GOOS=$(echo $target | cut -d '/' -f 1)
  GOARCH=$(echo $target | cut -d '/' -f 2)
  output="myapp-$GOOS-$GOARCH"
  echo "Building for $GOOS/$GOARCH..."
  GOOS=$GOOS GOARCH=$GOARCH go build -o $output
done
