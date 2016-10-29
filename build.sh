#!/bin/bash

set -euf -o pipefail

echo "Building library"
go build -o bin/reload -i . 
