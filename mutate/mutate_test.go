package mutate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)




func TestMutate(t *testing.T) {
	//arrange
	//act
	_, err := Mutate([]byte(input))

	//assert
	assert.NoError(t,err)
}


const input= `
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
`

