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

### MutatingWebhookConfiguration.webhooks.caBundle 

Fill caBundle property to make apiserver trust the TLS certificate of the webhook server.
Because we’ve signed our certificates with the Kubernetes API we can use the CA cert from our kubeconfig.
```bash
kubectl config view --raw --minify --flatten -o jsonpath='{.clusters[].cluster.certificate-authority-data}'
``` 
Returned value is base64 encoded certificate. 



## TEST DATA

Following json is prettify input for for `sleep` deployment 

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
               "kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{},\"name\":\"sleep\",\"namesp
ace\":\"mutant-test\"},\"spec\":{\"replicas\":1,\"selector\":{\"matchLabels\":{\"app\":\"sleep\"}},\"template\":{\"metadata\":{\"labels\":{\"app\":\"sleep\"}},\"spec\":{\"containers\":[{\"command\":[\"/bin/sleep\",\"infinity\"],\"im
age\":\"tutum/curl\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"sleep\"}]}}}}\n"
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


### todo:
 - check what happens when webhook request fails (restart policy always? )
 - what if deployment fails ? (crash-loop)
 - circular dependency ? 
 - 10k pods depending on our webhook ? 