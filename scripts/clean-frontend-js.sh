#!/bin/bash 

# Ran from the root of the project
# Remove all .js files in the frontend/src directory
find ./frontend/src -type f -name "*.js" -exec rm {} +

# Remove tsconfig.tsbuildinfo file
rm ./frontend/tsconfig.tsbuildinfo