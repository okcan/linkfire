# Linkfire

# +++Instruction+++

a microservice application has been developed, certain endpoints can be called through this application,

At first, a Dockerfile was developed to turn this system into a docker image, then we ran the following command for this docker file to work,

```
ubuntu@kubernetes-worker:~/test2$ sudo docker build -t okcan/test2:1.1 .
Sending build context to Docker daemon 7.662MB
```
 
 Code has become an image and will be deployed to our Kubernetes system thanks to the configs in the manifest file developed using this image.

``` 
 #ubuntu@kubernetes-worker:~/test2$ kubectl apply -f manifest.yaml 
 ```

Finnally, code is deployed to kubernetes and check with curl command.

```
  curl http://127.0.0.1:8080/
  [{"timestamp":"2022","hostname":"test","engine":"docker","visitor ":"192.168.1.1"}]
```

I prepared a pipeline for the establishment of this environment, but I could not write all the details because my time was limited, but this pipeline can be developed further.


P.S. 

If you are going to use it for yourself, you need to change the path and names in the code.

After the service is deployed, it needs to be port-forwarded as follows.
```
ubuntu@kubernetes-worker:~$ kubectl port-forward test2-65bfd6d685-zbrpm 8080:8080
```
P.S 2

I can share the resources I use if desired.
