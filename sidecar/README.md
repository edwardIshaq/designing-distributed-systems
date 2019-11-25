# Sidecar pattern
This pattern is usefull to add more than one container in the same pod.
The first example is going to be a logging sidecar.
So to begin with we will create two containers:

1- a web application that does some basic function like reversing a string
2- a logging container that exposes an API to log to disk
3- we will run these two containers on the same pod using kubernetes