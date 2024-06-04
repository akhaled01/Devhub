#!/bin/bash

if [[ $1 == "server" ]]; then
  go run .
elif [[ $1 == "frontend" ]]; then
  cd frontend
  npm run dev
else
  echo "Invalid argument. Please use 'server' or 'frontend'."
  exit 1
fi
