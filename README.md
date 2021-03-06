# Log Server

This project is intended to show how a log stream into Elastic Search using Logstash can be set up. Then use kibana to visualise and query the logs on the Kibana console, with all services set up in a local environment using docker.

## Requirements

-   Docker
-   Docker Compose

Docker compose should come with the docker install for mac

To install with homebrew, use the following command:

`brew install docker`

## Setup

Once docker is installed, make sure the docker daemon is actually running and also make sure no other applications are listening on the ports specified in the docker-compose.yml, or the run will fail.

Then just run the make command:

`make run` (in the root of the logserver project)

This will spin up four containers, one for the Go server, logstash, elastic search and kibana.

Note: the first run will require downloads of the images you don't have locally, any runs after that will boot up much faster, if you want to play around with the code.

## Server Endpoint

### /getexample

to test this endpoint, you can simply go to localhost:8080/getexample or curl it:

`curl localhost:8080/getexample`

### /postexample

the easiest way to test this endpoint would be to curl it in your terminal:

`curl -X POST localhost:8080/postexample -d '{"data": "string message"}'`

The handler expects a JSON body with a data field. Invalid JSON will cause a 500 status code response

## Viewing Logs In Kibana

1. Go to localhost:5601
2. Go to the setting tab where you will see a "Configure an index pattern"
3. Keep the default settings and click "create index" at the bottom.
4. Go to the discover tab to see and query your logs.
5. Hitting the server endpoints will generate logs (you can start hitting the server before you set up the index and they will still show up after)

## Clean Up

after exiting the containers, run `docker system prune`, this will remove any stopped containers to free up resources on your machine.

## Acknowledgements

-   Nic Jackson for his book, "Building Microservices With Go"
