package mutlabel


const podWithoutCustomLabel = `
{ 
   "kind":"AdmissionReview",
   "apiVersion":"admission.k8s.io/v1beta1",
   "request":{ 
      "uid":"60c76801-40ea-11ea-9a8c-66e02a952a06",
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
         "username":"aksService",
         "groups":[ 
            "system:masters",
            "system:authenticated"
         ]
      },
      "object":{ 
         "kind":"Pod",
         "apiVersion":"v1",
         "metadata":{ 
            "generateName":"sleep-7674d45776-",
            "creationTimestamp":null,
            "labels":{ 
               "app":"sleep",
               "pod-template-hash":"7674d45776"
            },
            "ownerReferences":[ 
               { 
                  "apiVersion":"apps/v1",
                  "kind":"ReplicaSet",
                  "name":"sleep-7674d45776",
                  "uid":"60c37a46-40ea-11ea-9a8c-66e02a952a06",
                  "controller":true,
                  "blockOwnerDeletion":true
               }
            ]
         },
         "spec":{ 
            "volumes":[ 
               { 
                  "name":"default-token-rlvqw",
                  "secret":{ 
                     "secretName":"default-token-rlvqw"
                  }
               }
            ],
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
                  "volumeMounts":[ 
                     { 
                        "name":"default-token-rlvqw",
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
      "dryRun":false,
      "options":null
   },
   "response":{ 
      "uid":"60c76801-40ea-11ea-9a8c-66e02a952a06",
      "allowed":true,
      "status":{ 
         "metadata":{ 

         },
         "status":"Success",
         "code":200
      },
      "patch":"W3sib3AiOiJyZW1vdmUiLCJwYXRoIjoiL21ldGFkYXRhL2xhYmVscyJ9LHsib3AiOiJhZGQiLCJwYXRoIjoiL21ldGFkYXRhL2xhYmVscyIsInZhbHVlIjp7ImFwcCI6InNsZWVwIiwiY29zdC1jZW50ZXIiOiI2MDAwMSIsImNyZWF0b3IiOiJnYTJrdWFsIiwiZW52aXJvbm1lbnQiOiJkZXYiLCJvd25lciI6ImdhMmt1YWwiLCJwb2QtdGVtcGxhdGUtaGFzaCI6Ijc2NzRkNDU3NzYiLCJwcm9kdWN0IjoiY2FzaC1zZXJ2aWNlcyJ9fV0=",
      "patchType":"JSONPatch"
   }
}
`


const podWithCustomLabel = `
{ 
   "kind":"AdmissionReview",
   "apiVersion":"admission.k8s.io/v1beta1",
   "request":{ 
      "uid":"fe379e6e-40e8-11ea-9a8c-66e02a952a06",
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
         "username":"aksService",
         "groups":[ 
            "system:masters",
            "system:authenticated"
         ]
      },
      "object":{ 
         "kind":"Pod",
         "apiVersion":"v1",
         "metadata":{ 
            "generateName":"beep-7964bf4878-",
            "creationTimestamp":null,
            "labels":{ 
               "app":"beep",
               "pod-template-hash":"7964bf4878"
            },
            "ownerReferences":[ 
               { 
                  "apiVersion":"apps/v1",
                  "kind":"ReplicaSet",
                  "name":"beep-7964bf4878",
                  "uid":"fe331cec-40e8-11ea-9a8c-66e02a952a06",
                  "controller":true,
                  "blockOwnerDeletion":true
               }
            ]
         },
         "spec":{ 
            "volumes":[ 
               { 
                  "name":"default-token-7zw6s",
                  "secret":{ 
                     "secretName":"default-token-7zw6s"
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
                        "name":"default-token-7zw6s",
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
      "dryRun":false,
      "options":null
   },
   "response":{ 
      "uid":"fe379e6e-40e8-11ea-9a8c-66e02a952a06",
      "allowed":true,
      "status":{ 
         "metadata":{ 

         },
         "status":"Success",
         "code":200
      },
      "patch":"W3sib3AiOiJyZW1vdmUiLCJwYXRoIjoiL21ldGFkYXRhL2xhYmVscyJ9LHsib3AiOiJhZGQiLCJwYXRoIjoiL21ldGFkYXRhL2xhYmVscyIsInZhbHVlIjp7ImFwcCI6ImJlZXAiLCJjb3N0LWNlbnRlciI6IjYwMDAxIiwiY3JlYXRvciI6ImdhMmt1YWwiLCJlbnZpcm9ubWVudCI6ImRldiIsIm93bmVyIjoiZ2Eya3VhbCIsInBvZC10ZW1wbGF0ZS1oYXNoIjoiNzk2NGJmNDg3OCIsInByb2R1Y3QiOiJjYXNoLXNlcnZpY2VzIn19XQ==",
      "patchType":"JSONPatch"
   }
}`

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

