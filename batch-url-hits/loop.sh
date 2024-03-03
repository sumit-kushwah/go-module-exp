# write a code to run command each 20 seconds for 1 hour
# command: echo "Hello World"

#!/bin/bash

export BEARER_TOKEN=$(gcloud auth print-access-token)
echo $BEARER_TOKEN

for i in $(seq 1 600)
do
    go run main.go
    sleep 10
    # if i is divisible by 20
    if [ $((i % 20)) -eq 0 ]
    then
        export BEARER_TOKEN=$(gcloud auth print-access-token)
        echo "new auth token $BEARER_TOKEN"
    fi
done