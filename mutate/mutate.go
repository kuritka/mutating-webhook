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


func Mutate(body []byte) ([]byte, error){

	review := new(v1beta1.AdmissionReview)
	if err := json.Unmarshal(body, review); err != nil {
		return body,  fmt.Errorf("unmarshaling request failed with %s", err)
	}

	request := review.Request


	var responseBytes []byte
	var err error
	var deployment appsv1.Deployment
	if err := json.Unmarshal(request.Object.Raw, &deployment); err != nil {
		return nil, fmt.Errorf("could not unmarshal raw object: %v", err)
	}
	availableLabels := deployment.Labels
	addedLabels := updateLabels(availableLabels, map[string]string {"environment":"dev", "product":"cash-services", "cost-center":"60001"})
	if responseBytes, err = json.Marshal(addedLabels); err != nil {
		return nil, fmt.Errorf("could not marshall response: %v", err)
	}

	response :=&v1beta1.AdmissionResponse{
		Allowed: true,
		Patch:   responseBytes,
		UID: request.UID,
		PatchType: func() *v1beta1.PatchType {
			pt := v1beta1.PatchTypeJSONPatch
			return &pt
		}(),
	}



	response.Result = &metav1.Status{ Status: "Success"}
	review.Response = response
	reviewBody, err := json.Marshal(review)
	if err != nil {
		return nil, fmt.Errorf("cannot create review body %s", err)
	}
	return reviewBody,nil
}