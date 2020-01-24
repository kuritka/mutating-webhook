package mutate

import (
	"encoding/json"
	"fmt"
	"mutating-webhook/common/guards"

	v1beta1 "k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

func Mutate(body []byte) ([]byte, error){
	guards.FailOnNil(body, "body cn't be nil")

	review := new(v1beta1.AdmissionReview)
	if err := json.Unmarshal(body, review); err != nil {
		return body,  fmt.Errorf("unmarshaling request failed with %s", err)
	}

	var pod *corev1.Pod
	request := review.Request
	response := new(v1beta1.AdmissionResponse)

	//unmarshall pod
	if err := json.Unmarshal(request.Object.Raw, &pod); err != nil {
		return nil, fmt.Errorf("unable unmarshal pod json object %v", err)
	}

	// set response options
	response.Allowed = true
	response.UID = request.UID
	//http://jsonpatch.com/
	jsonPatch := v1beta1.PatchTypeJSONPatch
	response.PatchType = &jsonPatch

	response.AuditAnnotations = map[string]string{
		"mutateme": "injected by mutating-webhook",
	}

	var err error
	px := make([]patchOperation,0)
	patch :=  patchOperation{
		Op:    "add",
		Path:  "/metadata/labels",
		Value: map[string]string {"key1":"value1"},
	}
	px = append(px,patch )
	response.Patch, err = json.Marshal(px)
	if err != nil {
		return nil, fmt.Errorf("cannot append label %s", err)
	}
	fmt.Println(string(response.Patch))

	response.Result = &metav1.Status{ Status: "Success"}
	review.Response = response
	reviewBody, err := json.Marshal(review)
	if err != nil {
		return nil, fmt.Errorf("cannot create review body %s", err)
	}

	return reviewBody,nil
}