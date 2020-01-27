# mutating-webhook

## instructions

- Setup a local Kubernetes cluster using Kind

- Write an admission webhook server that will add a label of your choosing to all pods created (https://github.com/banzaicloud/admission-webhook-example is a good start)

![](https://d33wubrfki0l68.cloudfront.net/af21ecd38ec67b3d81c1b762221b4ac777fcf02d/7c60e/images/blog/2019-03-21-a-guide-to-kubernetes-admission-controllers/admission-controller-phases.png)

- Add the necessary mutating admission webhook configuration to your Kubernetes cluster to get it working

- Record a screencast of your admission webhook in action

- Send us a link to the video as well as a public GitHub repo with the code to implement the above




## How to start

### install local cluster

install and start [kind](https://kind.sigs.k8s.io/) by `go get sigs.k8s.io/kind@v0.7.0 && kind create cluster`
and verify that `kind-kind` context is properly set `kubectl config get-contexts` 


check if `admissionregistration` is installed by running `kubectl api-versions | grep admissionregistration.k8s.io/v1beta1`
 

### specify private registry

Run following command to enable [private registry](https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod)
```bash
kubectl create secret docker-registry <name> --docker-server=DOCKER_REGISTRY_SERVER --docker-username=DOCKER_USER --docker-password=DOCKER_PASSWORD
```

All secrets are within [azure registry](https://portal.azure.com/#@deutscheboerse.onmicrosoft.com/resource/subscriptions/28ed73f5-4bb4-4064-bf48-b520cc638475/resourceGroups/rg-onho-sbx/providers/Microsoft.ContainerRegistry/registries/acronhosbx/accessKey)

### build and deploy 

run this:
```bash
./build.sh cicd v1.18
```



### todo:
 - [ ] check what happens when webhook request fails (restart policy always? )
 - [ ] what if deployment fails ? (crash-loop)
 - [ ] circular dependency ? 
 - [ ] 10k pods depending on our webhook ? 
 - [X] multiple containers in pod ?  - np, it is about pods, not about containers. there is already injected istio 
 - [X] what about deployments ? 
 - [X] look at this: imagePullPolicy: Always
 - [X] check with multiple types of Kinds (Deploy, POD, Service); resources are given by MutatingWebhookConfiguration
 - [ ] object model, refactor 
 - [ ] put app into new namespace (check with webhook-create-signed-cert.sh )
 - [x] What if label already exists? or no labels?
 - [x] Configuration of labels would be part of Deployment, check json as envvar 
 - [x] build should inject caBundle, build version and annotation 
 - [x] return Result: &metav1.Status{ Message: err.Error(), }, refactor

 
 
 ### terminal commands
 ```bash
watch kubectl logs `kubectl get pod -l app=mutating-webhook -o jsonpath="{.items[0].metadata.name}"` -c mutating-webhook 
```
 
