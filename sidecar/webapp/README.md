# Webapp Demo
This will serve as the web application which will be the main container running in the pod

# Building the docker image
docker image build . -t edwardi/dds.sidecar.webapp

# Running using Docker
docker run -d -p 8080:8080  edwardi/dds.sidecar.webapp

# Running using k8s
kubectl run webapp --generator=run-pod/v1 --image=edwardi/dds.sidecar.webapp
kubectl port-forward webapp 8080:8080