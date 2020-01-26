package mutate

import (
	"encoding/json"
	"fmt"

	"k8s.io/api/admission/v1beta1"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type patchOperation struct {
	Operation string      `json:"op"`
	Path      string      `json:"path"`
	Value     interface{} `json:"value,omitempty"`
}


func Mutate(body []byte, customLabels map[string]string) ([]byte, error){

	var err error
	var responseBody []byte
	status := &metav1.Status{ Status: "Success"}

	review := new(v1beta1.AdmissionReview)
	if err := json.Unmarshal(body, review); err != nil {
		return nil,  fmt.Errorf("unmarshaling request failed with %s", err)
	}


	var deployment appsv1.Deployment
	err = json.Unmarshal(review.Request.Object.Raw, &deployment)
	setStatusOnFail(status,err, "could not unmarshal raw deployment %v",err)

	labels := updateLabels(deployment.Labels, customLabels )
	responseBody, err = json.Marshal(labels)
	setStatusOnFail(status,err, "could not marshal %v;  %v",labels,err)

	review.Response =&v1beta1.AdmissionResponse{
		Allowed: true,
		Patch:   responseBody,
		UID:     review.Request.UID,
		PatchType: func() *v1beta1.PatchType {
			pt := v1beta1.PatchTypeJSONPatch
			return &pt
		}(),
		Result: status,
	}
	reviewBody, err := json.Marshal(review)
	if err != nil {
		return nil, fmt.Errorf("cannot create review body %s", err)
	}
	return reviewBody,nil
}


func setStatusOnFail(response *metav1.Status, err error, message string, v ...interface{}){
	if err != nil {
		response = &metav1.Status{Status: "Failure", Message: fmt.Sprintf("%s %v", message, v), Reason: "Invalid"}
	}
}


func getLabels(existingLabels map[string]string, added map[string]string) (labels map[string]string) {
	newMap := make(map[string]string, len(existingLabels) + len(added))
	for key, value := range existingLabels {
		newMap[key] = value
	}
	for key, value := range added {
		newMap[key] = value
	}
	return newMap
}


func updateLabels(existingLabels map[string]string, added map[string]string) (patch []patchOperation) {
	labels := getLabels(existingLabels,added)
	if existingLabels != nil {
		patch = append(patch, patchOperation{
			Operation: "remove",
			Path:      "/metadata/labels",
		})
	}
	patch = append(patch, patchOperation{
		Operation: "add",
		Path:      "/metadata/labels",
		Value:     labels,
	})
	return patch
}
