package mutlabel

import (
	"encoding/json"
	"k8s.io/api/admission/v1beta1"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyLabelsOnPodWithoutLabels(t *testing.T) {
	//arrange
	//empty configuration
	labels := map[string]string{}
	ml := NewMutLabel(labels)

	//act
	result, err := ml.Mutate([]byte(podWithoutLabels))
	review := new(v1beta1.AdmissionReview)
	_ = json.Unmarshal(result, review)
	patch := review.Response.Patch

	//assert
	assert.NoError(t, err)
	assert.Equal(t, `[{"op":"add","path":"/metadata/labels","value":{}}]`, string(patch))
	assert.Equal(t, review.Response.UID, review.Request.UID)
	assert.Equal(t, int32(http.StatusOK), review.Response.Result.Code)
	assert.Equal(t, "Success", review.Response.Result.Status)
}

func TestPodWithCustomLabels(t *testing.T) {
	//arrange
	labels := map[string]string{"environment": "dev", "product": "cash-services", "cost-center": "60001"}
	ml := NewMutLabel(labels)

	//act
	result, err := ml.Mutate([]byte(podWithCustomLabel))
	review := new(v1beta1.AdmissionReview)
	_ = json.Unmarshal(result, review)
	patch := review.Response.Patch

	//assert
	assert.NoError(t, err)
	assert.Equal(t, `[{"op":"remove","path":"/metadata/labels"},{"op":"add","path":"/metadata/labels","value":{"app":"beep","cost-center":"60001","environment":"dev","pod-template-hash":"7964bf4878","product":"cash-services"}}]`, string(patch))
	assert.Equal(t, review.Response.UID, review.Request.UID)
	assert.Equal(t, int32(http.StatusOK), review.Response.Result.Code)
	assert.Equal(t, "Success", review.Response.Result.Status)
}

func TestPodWithoutCustomLabel(t *testing.T) {
	//arrange
	labels := map[string]string{"environment": "dev", "product": "cash-services", "cost-center": "60001"}
	ml := NewMutLabel(labels)

	//act
	result, err := ml.Mutate([]byte(podWithoutCustomLabel))
	review := new(v1beta1.AdmissionReview)
	_ = json.Unmarshal(result, review)
	patch := review.Response.Patch

	//assert
	assert.NoError(t, err)
	assert.Equal(t, `[{"op":"remove","path":"/metadata/labels"},{"op":"add","path":"/metadata/labels","value":{"app":"sleep","cost-center":"60001","environment":"dev","pod-template-hash":"7674d45776","product":"cash-services"}}]`, string(patch))
	assert.Equal(t, review.Response.UID, review.Request.UID)
	assert.Equal(t, int32(http.StatusOK), review.Response.Result.Code)
	assert.Equal(t, "Success", review.Response.Result.Status)
}

func TestCorruptedJson(t *testing.T) {
	//arrange
	labels := map[string]string{"environment": "dev", "product": "cash-services", "cost-center": "60001"}
	ml := NewMutLabel(labels)

	//act
	result, err := ml.Mutate([]byte(corruptedPod))

	//assert
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestEmptyLabelsOnDeploymentWithCustomLabels(t *testing.T) {
	//arrange
	labels := map[string]string{}
	ml := NewMutLabel(labels)

	//act
	result, err := ml.Mutate([]byte(deploymentWithCustomLabel))
	review := new(v1beta1.AdmissionReview)
	_ = json.Unmarshal(result, review)
	patch := review.Response.Patch

	//assert
	assert.NoError(t, err)
	assert.Nil(t, patch)
	assert.Equal(t, review.Response.UID, review.Request.UID)
	assert.Equal(t, int32(http.StatusInternalServerError), review.Response.Result.Code)
	assert.Equal(t, "Failure", review.Response.Result.Status)
}

func TestDeploymentWithoutCustomLabel(t *testing.T) {
	//arrange
	labels := map[string]string{"environment": "dev", "product": "cash-services", "cost-center": "60001"}
	ml := NewMutLabel(labels)

	//act
	result, err := ml.Mutate([]byte(deploymentWithoutCustomLabel))
	review := new(v1beta1.AdmissionReview)
	_ = json.Unmarshal(result, review)
	patch := review.Response.Patch

	//assert
	assert.NoError(t, err)
	assert.Nil(t, patch)
	assert.Equal(t, review.Response.UID, review.Request.UID)
	assert.Equal(t, int32(http.StatusInternalServerError), review.Response.Result.Code)
	assert.Equal(t, "Failure", review.Response.Result.Status)
}

func TestPodWithoutLabels(t *testing.T) {
	//arrange
	labels := map[string]string{"environment": "dev", "product": "cash-services", "cost-center": "60001"}
	ml := NewMutLabel(labels)

	//act
	result, err := ml.Mutate([]byte(podWithoutLabels))
	review := new(v1beta1.AdmissionReview)
	_ = json.Unmarshal(result, review)
	patch := review.Response.Patch

	//assert
	assert.NoError(t, err)
	assert.Equal(t, `[{"op":"add","path":"/metadata/labels","value":{"cost-center":"60001","environment":"dev","product":"cash-services"}}]`, string(patch))
	assert.Equal(t, review.Response.UID, review.Request.UID)
	assert.Equal(t, int32(http.StatusOK), review.Response.Result.Code)
	assert.Equal(t, "Success", review.Response.Result.Status)
}
