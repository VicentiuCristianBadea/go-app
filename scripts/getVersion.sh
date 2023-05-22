#!/bin/sh

echo 'Testing 2 things. I want to see the current tag increments, and if the new tag is exported as an environment variable.'

current_tag="1.0.0"
echo 'Current tag: ' $current_tag

# Get the version of the last tag from DockerHub
TAG=$(curl -s https://registry.hub.docker.com/v1/repositories/$1/tags | jq -r '.[].name' | grep -v latest | sort -V | tail -n 1)
echo 'Received tag: ' $TAG

# Split the current tag into an array using '.' as delimiter
IFS='.' read -ra ADDR <<< "$current_tag"
echo 'IFS: ' $IFS

# Increment the patch version
new_patch_version=$((ADDR[2] + 1))
echo 'New patch version: ' $new_patch_version

# Construct the new tag
new_tag="${ADDR[0]}.${ADDR[1]}.${new_patch_version}"
echo 'new_tag: ' $new_tag

# Export the new tag as an environment variable
export TAG=$new_tag
echo $TAG
