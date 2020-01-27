package mutate

import (
	"encoding/json"
	"k8s.io/api/admission/v1beta1"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestEmptyLabelsOnPodWithoutLabels(t *testing.T){
	//arrange
	labels := map[string]string {}
	//act
	result, err := Mutate([]byte(podWithoutLabels), labels)
	review := new(v1beta1.AdmissionReview)
	_ = json.Unmarshal(result, review)
	patch := review.Response.Patch
	//assert
	assert.NoError(t,err)
	assert.Equal(t,`[{"op":"add","path":"/metadata/labels","value":{}}]`, string(patch))
}



func TestEmptyLabelsOnDeploymentWithCustomLabels(t *testing.T){
	//arrange
	labels := map[string]string {}
	//act
	result, err := Mutate([]byte(deploymentWithCustomLabel), labels)
	review := new(v1beta1.AdmissionReview)
	_ = json.Unmarshal(result, review)
	patch := review.Response.Patch
	//assert
	assert.NoError(t,err)
	assert.Equal(t,`[{"op":"remove","path":"/metadata/labels"},{"op":"add","path":"/metadata/labels","value":{"app":"beep","product":"EDTderivates"}}]`, string(patch))
}

func TestCorruptedJson(t *testing.T){
	//arrange
	labels := map[string]string {"environment":"dev", "product":"cash-services", "cost-center":"60001"}
	//act
	result, err := Mutate([]byte(corruptedPod), labels)
	//assert
	assert.Error(t,err)
	assert.Nil(t, result)
}

func TestDeploymentWithoutCustomLabel(t *testing.T) {
	//arrange
	labels := map[string]string {"environment":"dev", "product":"cash-services", "cost-center":"60001"}
	//act
	result, err := Mutate([]byte(deploymentWithoutCustomLabel), labels)
	review := new(v1beta1.AdmissionReview)
	_ = json.Unmarshal(result, review)
	patch := review.Response.Patch
	//assert
	assert.NoError(t,err)
	assert.Equal(t,`[{"op":"add","path":"/metadata/labels","value":{"cost-center":"60001","environment":"dev","product":"cash-services"}}]`, string(patch))

}


func TestDeploymentWithExistingLabel(t *testing.T) {
	//arrange
	labels := map[string]string {"environment":"dev", "product":"cash-services", "cost-center":"60001"}
	//act
	result, err := Mutate([]byte(deploymentWithCustomLabel), labels)
	review := new(v1beta1.AdmissionReview)
	_ = json.Unmarshal(result, review)
	patch := review.Response.Patch
	//assert
	assert.NoError(t,err)
	assert.Equal(t,`[{"op":"remove","path":"/metadata/labels"},{"op":"add","path":"/metadata/labels","value":{"app":"beep","cost-center":"60001","environment":"dev","product":"cash-services"}}]`, string(patch))
}

func TestPodWithoutLabels(t *testing.T) {
	//arrange
	labels := map[string]string {"environment":"dev", "product":"cash-services", "cost-center":"60001"}
	//act
	result, err := Mutate([]byte(deploymentWithCustomLabel), labels)
	review := new(v1beta1.AdmissionReview)
	_ = json.Unmarshal(result, review)
	patch := review.Response.Patch
	//assert
	assert.NoError(t,err)
	assert.Equal(t,`[{"op":"remove","path":"/metadata/labels"},{"op":"add","path":"/metadata/labels","value":{"app":"beep","cost-center":"60001","environment":"dev","product":"cash-services"}}]`, string(patch))
}





const deploymentWithCustomLabel = `
{ 
   "kind":"AdmissionReview",
   "apiVersion":"admission.k8s.io/v1beta1",
   "request":{ 
      "uid":"19bd36fb-3fc7-11ea-9a8c-66e02a952a06",
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
            "name":"beep",
            "namespace":"mutant-test",
            "creationTimestamp":null,
            "labels":{ 
               "app":"beep",
               "product":"EDTderivates"
            },
            "annotations":{ 
               "kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"beep\",\"product\":\"EDTderivates\"},\"name\":\"beep\",\"namespace\":\"mutant-test\"},\"spec\":{\"replicas\":1,\"selector\":{\"matchLabels\":{\"app\":\"beep\"}},\"template\":{\"metadata\":{\"labels\":{\"app\":\"beep\"}},\"spec\":{\"containers\":[{\"command\":[\"/bin/sleep\",\"infinity\"],\"image\":\"tutum/curl\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"beep\"}]}}}}\n"            }
         },
         "spec":{ 
            "replicas":1,
            "selector":{ 
               "matchLabels":{ 
                  "app":"beep"
               }
            },
            "template":{ 
               "metadata":{ 
                  "creationTimestamp":null,
                  "labels":{ 
                     "app":"beep"
                  }
               },
               "spec":{ 
                  "containers":[ 
                     { 
                        "name":"beep",
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
`

const deploymentWithoutCustomLabel = `
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
      "patchOperation":"CREATE",
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
`



const podWithoutLabels = `
{ 
   "kind":"AdmissionReview",
   "apiVersion":"admission.k8s.io/v1beta1",
   "request":{ 
      "uid":"d27ba71d-402f-11ea-9a8c-66e02a952a06",
      "kind":{ 
         "group":"",
         "version":"v1",
         "kind":"Pod"
      },
      "resource":{ 
         "group":"",
         "version":"v1",
         "resource":"pods"
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
         "kind":"Pod",
         "apiVersion":"v1",
         "metadata":{ 
            "name":"no-labels",
            "namespace":"mutant-test",
            "creationTimestamp":null,
            "annotations":{ 
               "kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"Pod\",\"metadata\":{\"annotations\":{},\"name\":\"no-labels\",\"namespace\":\"mutant-test\"},\"spec\":{\"containers\":[{\"command\":[\"/bin/sleep\",\"infinity\"],\"image\":\"tutum/curl\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"beep\"}]}}\n"
            }
         },
         "spec":{ 
            "volumes":[ 
               { 
                  "name":"default-token-88dzx",
                  "secret":{ 
                     "secretName":"default-token-88dzx"
                  }
               }
            ],
            "containers":[ 
               { 
                  "name":"beep",
                  "image":"tutum/curl",
                  "command":[ 
                     "/bin/sleep",
                     "infinity"
                  ],
                  "resources":{ 

                  },
                  "volumeMounts":[ 
                     { 
                        "name":"default-token-88dzx",
                        "readOnly":true,
                        "mountPath":"/var/run/secrets/kubernetes.io/serviceaccount"
                     }
                  ],
                  "terminationMessagePath":"/dev/termination-log",
                  "terminationMessagePolicy":"File",
                  "imagePullPolicy":"IfNotPresent"
               }
            ],
            "restartPolicy":"Always",
            "terminationGracePeriodSeconds":30,
            "dnsPolicy":"ClusterFirst",
            "serviceAccountName":"default",
            "serviceAccount":"default",
            "securityContext":{ 

            },
            "schedulerName":"default-scheduler",
            "tolerations":[ 
               { 
                  "key":"node.kubernetes.io/not-ready",
                  "operator":"Exists",
                  "effect":"NoExecute",
                  "tolerationSeconds":300
               },
               { 
                  "key":"node.kubernetes.io/unreachable",
                  "operator":"Exists",
                  "effect":"NoExecute",
                  "tolerationSeconds":300
               }
            ],
            "priority":0,
            "enableServiceLinks":true
         },
         "status":{ 

         }
      },
      "oldObject":null,
      "dryRun":false
   }
}`


const corruptedPod= `
{ 
   "kind:"AdmissionReview",
   "apiVersion":"admission.k8s.io/v1beta1",
   "request":{ 
      "uid":"d27ba71d-402f-11ea-9a8c-66e02a952a06",
      "kind":{ 
         "group":"",
         "version":"v1",
         "kind":"Pod"
      },
      "resource":{ 
         "group":"",
         "version":"v1",
         "resource":"pods"
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
         "kind":"Pod",
         "apiVersion":"v1",
         "metadata":{ 
            "name":"no-labels",
            "namespace":"mutant-test",
            "creationTimestamp":null,
            "annotations":{ 
               "kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"Pod\",\"metadata\":{\"annotations\":{},\"name\":\"no-labels\",\"namespace\":\"mutant-test\"},\"spec\":{\"containers\":[{\"command\":[\"/bin/sleep\",\"infinity\"],\"image\":\"tutum/curl\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"beep\"}]}}\n"
            }
         },
         "spec":{ 
            "volumes":[ 
               { 
                  "name":"default-token-88dzx",
                  "secret":{ 
                     "secretName":"default-token-88dzx"
                  }
               }
            ],
            "containers":[ 
               { 
                  "name":"beep",
                  "image":"tutum/curl",
                  "command":[ 
                     "/bin/sleep",
                     "infinity"
                  ],
                  "resources":{ 

                  },
                  "volumeMounts":[ 
                     { 
                        "name":"default-token-88dzx",
                        "readOnly":true,
                        "mountPath":"/var/run/secrets/kubernetes.io/serviceaccount"
                     }
                  ],
                  "terminationMessagePath":"/dev/termination-log",
                  "terminationMessagePolicy":"File",
                  "imagePullPolicy":"IfNotPresent"
               }
            ],
            "restartPolicy":"Always",
            "terminationGracePeriodSeconds":30,
            "dnsPolicy":"ClusterFirst",
            "serviceAccountName":"default",
            "serviceAccount":"default",
            "securityContext":{ 

            },
            "schedulerName":"default-scheduler",
            "tolerations":[ 
               { 
                  "key":"node.kubernetes.io/not-ready",
                  "operator":"Exists",
                  "effect":"NoExecute",
                  "tolerationSeconds":300
               },
               { 
                  "key":"node.kubernetes.io/unreachable",
                  "operator":"Exists",
                  "effect":"NoExecute",
                  "tolerationSeconds":300
               }
            ],
            "priority":0,
            "enableServiceLinks":true
         },
         "status":{ 

         }
      },
      "oldObject":null,
      "dryRun":false
   }
}`

