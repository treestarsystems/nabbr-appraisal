#!/bin/bash

# Ran from the root of the project
# Frontend
cd ./frontend
npm i
npm run build:production
cd ..
# Remove all .js files in the frontend/src directory
find ./frontend/src -type f -name "*.js" -exec rm {} +
# Remove tsconfig.tsbuildinfo file
rm ./frontend/tsconfig.tsbuildinfo

# Backend
cd ./backend
make build