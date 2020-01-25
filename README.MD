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



## TEST DATA

Following json is prettify input for for `sleep` deployment 

### AdmissionReview - request

```json
{ 
   "kind":"AdmissionReview",
   "apiVersion":"admission.k8s.io/v1beta1",
   "request":{ 
      "uid":"2de8bc03-3de2-11ea-9beb-b2bf8e0fdd36",
      "kind":{ 
         "group":"apps",
         "version":"v1",
         "kind":"Deployment"
      },
      "resource":{ 
         "group":"apps",
         "version":"v1",
         "resource":"deployments"
      },
      "namespace":"mutant-test",
      "operation":"CREATE",
      "userInfo":{ 
         "username":"masterclient",
         "groups":[ 
            "system:masters",
            "system:authenticated"
         ]
      },
      "object":{ 
         "kind":"Deployment",
         "apiVersion":"apps/v1",
         "metadata":{ 
            "name":"sleep",
            "namespace":"mutant-test",
            "creationTimestamp":null,
            "annotations":{ 
               "kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{},\"name\":\"sleep\",\"namespace\":\"mutant-test\"},\"spec\":{\"replicas\":1,\"selector\":{\"matchLabels\":{\"app\":\"sleep\"}},\"template\":{\"metadata\":{\"labels\":{\"app\":\"sleep\"}},\"spec\":{\"containers\":[{\"command\":[\"/bin/sleep\",\"infinity\"],\"image\":\"tutum/curl\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"sleep\"}]}}}}\n"
            }
         },
         "spec":{ 
            "replicas":1,
            "selector":{ 
               "matchLabels":{ 
                  "app":"sleep"
               }
            },
            "template":{ 
               "metadata":{ 
                  "creationTimestamp":null,
                  "labels":{ 
                     "app":"sleep"
                  }
               },
               "spec":{ 
                  "containers":[ 
                     { 
                        "name":"sleep",
                        "image":"tutum/curl",
                        "command":[ 
                           "/bin/sleep",
                           "infinity"
                        ],
                        "resources":{ 

                        },
                        "terminationMessagePath":"/dev/termination-log",
                        "terminationMessagePolicy":"File",
                        "imagePullPolicy":"IfNotPresent"
                     }
                  ],
                  "restartPolicy":"Always",
                  "terminationGracePeriodSeconds":30,
                  "dnsPolicy":"ClusterFirst",
                  "securityContext":{ 

                  },
                  "schedulerName":"default-scheduler"
               }
            },
            "strategy":{ 
               "type":"RollingUpdate",
               "rollingUpdate":{ 
                  "maxUnavailable":"25%",
                  "maxSurge":"25%"
               }
            },
            "revisionHistoryLimit":10,
            "progressDeadlineSeconds":600
         },
         "status":{ 

         }
      },
      "oldObject":null,
      "dryRun":false
   }
}
```


### AdmissionReview - response

```json
{ 
   "kind":"AdmissionReview",
   "apiVersion":"admission.k8s.io/v1beta1",
   "request":{ 
      "uid":"56a9f6c2-3f79-11ea-9a8c-66e02a952a06",
      "kind":{ 
         "group":"apps",
         "version":"v1",
         "kind":"Deployment"
      },
      "resource":{ 
         "group":"apps",
         "version":"v1",
         "resource":"deployments"
      },
      "namespace":"mutant-test",
      "operation":"CREATE",
      "userInfo":{ 
         "username":"masterclient",
         "groups":[ 
            "system:masters",
            "system:authenticated"
         ]
      },
      "object":{ 
         "kind":"Deployment",
         "apiVersion":"apps/v1",
         "metadata":{ 
            "name":"sleep",
            "namespace":"mutant-test",
            "creationTimestamp":null,
            "labels":{ 
               "app":"sleep"
            },
            "annotations":{ 
               "kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"sleep\"},\"name\":\"sleep\",\"namespace\":\"mutant-test\"},\"spec\":{\"replicas\":1,\"selector\":{\"matchLabels\":{\"app\":\"sleep\"}},\"template\":{\"metadata\":{\"labels\":{\"app\":\"sleep\"}},\"spec\":{\"containers\":[{\"command\":[\"/bin/sleep\",\"infinity\"],\"image\":\"tutum/curl\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"sleep\"}]}}}}\n"
            }
         },
         "spec":{ 
            "replicas":1,
            "selector":{ 
               "matchLabels":{ 
                  "app":"sleep"
               }
            },
            "template":{ 
               "metadata":{ 
                  "creationTimestamp":null,
                  "labels":{ 
                     "app":"sleep"
                  }
               },
               "spec":{ 
                  "containers":[ 
                     { 
                        "name":"sleep",
                        "image":"tutum/curl",
                        "command":[ 
                           "/bin/sleep",
                           "infinity"
                        ],
                        "resources":{ 

                        },
                        "terminationMessagePath":"/dev/termination-log",
                        "terminationMessagePolicy":"File",
                        "imagePullPolicy":"IfNotPresent"
                     }
                  ],
                  "restartPolicy":"Always",
                  "terminationGracePeriodSeconds":30,
                  "dnsPolicy":"ClusterFirst",
                  "securityContext":{ 

                  },
                  "schedulerName":"default-scheduler"
               }
            },
            "strategy":{ 
               "type":"RollingUpdate",
               "rollingUpdate":{ 
                  "maxUnavailable":"25%",
                  "maxSurge":"25%"
               }
            },
            "revisionHistoryLimit":10,
            "progressDeadlineSeconds":600
         },
         "status":{ 

         }
      },
      "oldObject":null,
      "dryRun":false,
      "options":null
   },
   "response":{ 
      "uid":"56a9f6c2-3f79-11ea-9a8c-66e02a952a06",
      "allowed":true,
      "status":{ 
         "metadata":{ 

         },
         "status":"Success"
      },
      "patch":"W3sib3AiOiJhZGQiLCJwYXRoIjoiL21ldGFkYXRhL2xhYmVscyIsInZhbHVlIjp7ImVudmlyb25tZW50IjoiZGV2In19XQ==",
      "patchType":"JSONPatch"
   }
}
```

### todo:
 - [ ] check what happens when webhook request fails (restart policy always? )
 - [ ] what if deployment fails ? (crash-loop)
 - [ ] circular dependency ? 
 - [ ] 10k pods depending on our webhook ? 
 - [ ] multiple containers in pod ?  - np, it is about pods, not about containers.. 
 - [X] what about deployments ? 
 - [X] look at this: imagePullPolicy: Always
 - [ ] check with multiple types of Kinds (Deploy, POD, Service)
 - [ ] mutate as strategy, refactor 
 - [ ] co novy namespace ?
 - [ ] What if label already exists? or no labels?
 - [ ] Configuration of labels would be part of Deployment
 - [x] build should inject caBundle, build version and annotation 
 - [ ] return Result: &metav1.Status{ Message: err.Error(), },

 
 
 ### terminal commands
 ```bash
watch kubectl logs `kubectl get pod -l app=mutating-webhook -o jsonpath="{.items[0].metadata.name}"` -c mutating-webhook 
```
 
