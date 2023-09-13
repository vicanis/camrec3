#!/bin/bash
set -eo pipefail
ARTIFACT_BUCKET=$(cat bucket-name.txt)
cd mailhandler
GOOS=linux CGO_ENABLED=0 go build main.go
cd ../
cd eventprocessor
GOOS=linux CGO_ENABLED=0 go build main.go
cd ../
aws cloudformation package --template-file template.yml --s3-bucket $ARTIFACT_BUCKET --output-template-file out.yml
aws cloudformation deploy --template-file out.yml --stack-name blank-go --capabilities CAPABILITY_NAMED_IAM
