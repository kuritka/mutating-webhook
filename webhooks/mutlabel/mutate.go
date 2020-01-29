package mutlabel

import (
	"encoding/json"
	"fmt"
	"mutating-webhook/common/extensions"
	"net/http"

	"k8s.io/api/admission/v1beta1"
	"mutating-webhook/common/log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var logger = log.Log

type patchOperation struct {
	Operation string      `json:"op"`
	Path      string      `json:"path"`
	Value     interface{} `json:"value,omitempty"`
}

type MutLabel struct {
	labels map[string]string
}

func NewMutLabel(labels map[string]string) *MutLabel {
	ml := new(MutLabel)
	ml.labels = labels
	return ml
}

func (ml *MutLabel) Mutate(body []byte) ([]byte, error) {
	review := new(v1beta1.AdmissionReview)
	if err := json.Unmarshal(body, review); err != nil {
		return nil, fmt.Errorf("unmarshaling request failed with %s", err)
	}

	var responseBody []byte
	status := &metav1.Status{Status: "Failure", Message: fmt.Sprintf("only pod can be mutated"), Reason: "Invalid", Code: http.StatusInternalServerError}

	if review.Request.Kind.Kind == "Pod" {
		var err error

		var pod corev1.Pod

		if err := json.Unmarshal(review.Request.Object.Raw, &pod); err != nil {
			return nil, fmt.Errorf("unmarshaling pod from request failed with %s", err)
		}

		labels := ml.updateLabels(pod.Labels, ml.labels)
		responseBody, err = json.Marshal(labels)
		if err != nil {
			return nil, fmt.Errorf("could not marshal %v;  %v", labels, err)
		}

		status = &metav1.Status{Status: "Success", Code: http.StatusOK}

		logger.Info().Msgf("applying %s: %s", review.Request.Kind.Kind, ml.getPodName(&pod))
	}

	review.Response = &v1beta1.AdmissionResponse{
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
	return reviewBody, nil
}

func (ml *MutLabel) getPodName(pod *corev1.Pod) string {
	if pod.Name != "" {
		return pod.Name
	}
	return pod.ObjectMeta.GenerateName
}

//remove all labels and append existing + new
func (ml *MutLabel) updateLabels(existingLabels map[string]string, added map[string]string) (patch []patchOperation) {
	labels := extensions.Union(existingLabels, added)
	//remove all
	if existingLabels != nil {
		patch = append(patch, patchOperation{
			Operation: "remove",
			Path:      "/metadata/labels",
		})
	}
	//add union
	patch = append(patch, patchOperation{
		Operation: "add",
		Path:      "/metadata/labels",
		Value:     labels,
	})
	return patch
}
