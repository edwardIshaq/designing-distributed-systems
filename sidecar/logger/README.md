## Building the Docker image
docker build . -t edwardi/dds.sidecar.logger

## Running Docker image
we're also mapping port since our container needs to listen on that port
docker run -d -p 8090:8090  edwardi/dds.sidecar.logger
