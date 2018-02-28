
# gotstufftodo

Sample golang lambda function that reads and writes entries from dynamodb

## Pre-reqs

serverless framework, google it and install it

## Deploy and Run

```
make
serverless deploy -v
serverless invoke -f dbread -l
```