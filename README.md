# linkfire

#+++Instruction+++

a microservice application has been developed, certain endpoints can be called through this application,

At first, a Dockerfile was developed to turn this system into a docker image, then we ran the following command for this docker file to work,

ubuntu@kubernetes-worker:~/test2$ sudo docker build -t okcan/test2:1.1 .
Sending build context to Docker daemon 7.662MB
Step 1/6 : FROM golang:1.7.4-alpine
 ---> 00371bbb49d5
Step 2/6 : COPY . /home/ubuntu/test2
 ---> c20e4b2e2584
Step 3/6 : RUN export GOBIN=$GOPATH/bin && cd /home/ubuntu/test2 && CGO_ENABLED=0 go install
 ---> Running in a310ea7dc8ba
Removing intermediate container a310ea7dc8ba
 ---> ba6e26cb95b1
Step 4/6 : ENV PORT 8080
 ---> Running in 8b6f86599de9
Removing intermediate container 8b6f86599de9
 ---> 465fc4773557
Step 5/6 : EXPOSE 8080
 ---> Running in f0d88b847b4a
Removing intermediate container f0d88b847b4a
 ---> fa5a11833ffe
Step 6/6 : ENTRYPOINT test2
 ---> Running in 8576cde411d4
Removing intermediate container 8576cde411d4
 ---> e0e0b519e97c
 
 Code has become an image and will be deployed to our Kubernetes system thanks to the configs in the manifest file developed using this image.
 
 #ubuntu@kubernetes-worker:~/test2$ kubectl apply -f manifest.yaml 
 
 Finnally, code is deployed to kubernetes and check with curl command.
 
 #curl http://127.0.0.1:8080/
#[{"timestamp":"2022","hostname":"test","engine":"docker","visitor ":"192.168.1.1"}]

I prepared a pipeline for the establishment of this environment, but I could not write all the details because my time was limited, but this pipeline can be developed further.


P.S. 

If you are going to use it for yourself, you need to change the path and names in the code.

After the service is deployed, it needs to be port-forwarded as follows.

ubuntu@kubernetes-worker:~$ kubectl port-forward test2-65bfd6d685-zbrpm 8080:8080

P.S 2

I can share the resources I use if desired.
