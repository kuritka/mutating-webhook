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
Initialize application by running following command: 
```bash
./webhook-init.sh
```
`webhook-init.sh` creates namespace `webhook-system` and secret containing key+cert pair. Certificate is signed by cluster authority.  


To redeploy run this:
```bash
./build.sh cicd v1.20
```
Script redeploy webhook(s). It is absolutely necessary to deploy by `./build.sh` as placeholders are filled dynamically.   

For testing run this:
```bash
kubectl apply -f ./deployment/_test.yaml
```
Command creates several pods in `mutant-test` namespace.

 ### terminal commands
 
 ```bash
watch kubectl logs `kubectl get pod -l app=mutating-webhook -n webhook-system -o jsonpath="{.items[0].metadata.name}"` -c mutating-webhook -n webhook-system  
```
 
 ---

### todo
 - [x] check what happens when webhook request fails - error is logged, no changes in pod, validator needs to take care of integrity until mute-webhook is fixed
 - [x] what if deployment fails ? (crash-loop) - restart policy always. If pods are created during initialization, than  pods without injections will occur. Validator should check label integrity
 - [x] circular dependency ? - no 
 - [x] 10k pods depending on our webhook ? - deployment + Replicasets
 - [X] multiple containers in pod ?  - np, it is about pods, not about containers. there is already injected istio 
 - [X] what about deployments ? 
 - [X] look at this: imagePullPolicy: Always
 - [X] check with multiple types of Kinds (Deploy, POD, Service); resources are given by MutatingWebhookConfiguration
 - [x] encapsulate mutlabel 
 - [x] put resources into dedicated namespace (check with webhook-create-signed-cert.sh )
 - [x] What if label already exists? or no labels?
 - [x] Configuration of labels would be part of Deployment, check json as envvar 
 - [x] build should inject caBundle, build version and annotation 
 - [x] return Result: &metav1.Status{ Message: err.Error(), }, refactor
 - [x] setup sa + rbac

 
 
