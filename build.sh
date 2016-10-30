#!/bin/bash

set -euf -o pipefail

echo "Building the reload application..."
go build -o bin/reload -i . 
echo "Build done!"
