#!/bin/sh

# Read each module path from the go.work file and run go mod download for each.
grep '^./' go.work | while read -r module; do
    (cd "$module" && go mod download)
done