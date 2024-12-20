/*
 * Copyright contributors to the IBM Verify Identity Access Operator project
 */

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IBMSecurityVerifyAccessContainer defines the make-up of the container.  It
// is loosely based on the corev1.Container structure.

type IBMSecurityVerifyAccessContainer struct {
	// List of sources to populate environment variables in the container.
	// The keys defined within a source must be a C_IDENTIFIER. All invalid keys
	// will be reported as an event when the container is starting. When a key
	// exists in multiple sources, the value associated with the last source
	// will take precedence.  Values defined by an Env with a duplicate key
	// will take precedence.
	// Cannot be updated.
	// +optional
	EnvFrom []corev1.EnvFromSource `json:"envFrom,omitempty" protobuf:"bytes,19,rep,name=envFrom"`

	// List of environment variables to set in the container.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	Env []corev1.EnvVar `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"`

	// Compute Resources required by this container.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"`

	// Pod volumes to mount into the container's filesystem.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=mountPath
	// +patchStrategy=merge
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty" patchStrategy:"merge" patchMergeKey:"mountPath" protobuf:"bytes,9,rep,name=volumeMounts"`

	// volumeDevices is the list of block devices to be used by the container.
	// +patchMergeKey=devicePath
	// +patchStrategy=merge
	// +optional
	VolumeDevices []corev1.VolumeDevice `json:"volumeDevices,omitempty" patchStrategy:"merge" patchMergeKey:"devicePath" protobuf:"bytes,21,rep,name=volumeDevices"`

	// Periodic probe of container liveness.
	// Container will be restarted if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	LivenessProbe *corev1.Probe `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`

	// Periodic probe of container service readiness.
	// Container will be removed from service endpoints if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	ReadinessProbe *corev1.Probe `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`

	// StartupProbe indicates that the Pod has successfully initialized.
	// If specified, no other probes are executed until this completes
	// successfully.  If this probe fails, the Pod will be restarted, just as
	// if the livenessProbe failed.  This can be used to provide different
	// probe parameters at the beginning of a Pod's lifecycle, when it might
	// take a long time to load data or warm a cache, than during steady-state
	// operation.
	// This cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	StartupProbe *corev1.Probe `json:"startupProbe,omitempty" protobuf:"bytes,22,opt,name=startupProbe"`

	// Image pull policy.
	// One of Always, Never, IfNotPresent.
	// Defaults to Always if :latest tag is specified, or IfNotPresent
	// otherwise.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	// +optional
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty" protobuf:"bytes,14,opt,name=imagePullPolicy,casttype=PullPolicy"`

	// SecurityContext defines the security options the container should be run
	// with.  If set, the fields of SecurityContext override the equivalent
	// fields of PodSecurityContext.
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	// +optional
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty" protobuf:"bytes,15,opt,name=securityContext"`
}

// Language is the language in which messages will be displayed in the
// deployment.
type Language string

const (
	Chinese_Simplified  Language = "zh_CN.utf8"
	Chinese_Traditional Language = "zh_TW.utf8"
	Czech               Language = "cs_CZ.utf8"
	English             Language = "en_US.utf8"
	French              Language = "fr_FR.utf8"
	German              Language = "de_DE.utf8"
	Hungarian           Language = "hu_HU.utf8"
	Italian             Language = "it_IT.utf8"
	Japanese            Language = "ja_JP.utf8"
	Korean              Language = "ko_KR.utf8"
	Polish              Language = "pl_PL.utf8"
	Portuguese          Language = "pt_BR.utf8"
	Russian             Language = "ru_RU.utf8"
	Spanish             Language = "es_ES.utf8"
)

type LicenseModule string

const (
	Access_Control LicenseModule = "access_control"
	Federation     LicenseModule = "federation"
	ReverseProxy   LicenseModule = "webseal"
	Enterprise     LicenseModule = "enterprise"
)

// IBM License Metric Tool Annotations define the licence annotations added to runtime containers
// to track license usage by the IBM Licence Metric Tool.
type ILMTAnnotations struct {
	// +kubebuilder:default=webseal
	// +kubebuilder:validation:Enum=webseal;federation;access_control;enterprise
	// Licensed module to attach to container.
	Module LicenseModule `json:"module" protobuf:"bytes,32,rep,name=module,casttype=LicenseModule"`
	// +kubebuilder: default=true
	// Boolean flag to switch between production and development annotations.
	Production bool `json:"production"`
}

// Custom annotations to add to deployed Verify Access runtime container.
type CustomAnnotation struct {
	// Key of the annotation to create.
	Key string `json:"key" protobuf:"bytes,64,rep,name=key"`
	// Value of the annotation to create.
	Value string `json:"value" protobuf:"bytes,64,rep,name=value"`
}

// IBMSecurityVerifyAccessSpec defines the desired state of an
// IBMSecurityVerifyAccess resource.
type IBMSecurityVerifyAccessSpec struct {
	// The name of the image which will be used in the deployment.
	// Cannot be updated.
	Image string `json:"image"`

	//+kubebuilder:validation:Minimum=0
	//+kubebuilder:default=1
	// Replicas is the number of pods which will be started for the deployment.
	// +optional
	Replicas int32 `json:"replicas"`

	//+kubebuilder:default=true
	// AutoRestart is a boolean which indicates whether the deployment should
	// be restarted if a new snapshot is published
	// +optional
	AutoRestart bool `json:"autoRestart"`

	//+kubebuilder:default=published
	// SnapshotId is a string which is used to indicate the identifier of the
	// snapshot which should be used.  If no identifier is specified a default
	// snapshot of 'published' will be used.
	// Cannot be updated.
	// +optional
	SnapshotId string `json:"snapshotId"`

	// List of secrets to decrypt configuration snapshot files. Secrets are separated by '||'. This option is the
	// equivalent of setting the CONFIG_SNAPSHOT_SECRETS environment property.
	// +optional
	SnapshotSecrets string `json:"snapshotSecrets"`

	//+kubebuilder:default=operator
	// SnapshotTLSCacert is a string which defines how the Verify Identity Access runtime containers
	// verify connections to the snapshot management service. This option is the equivalent
	// of setting the CONFIG_SERVICE_TLS_CACERT environment property. The default option for this
	// property is to read the X509 certificate for the Operator's snapshot management service
	// from the verify-access-operator secret.
	// Note: Administrators must ensure that the service account for the runtime containers has
	// permission to read Secrets in the namespace that the Pod is deployed to in order for this to work.
	// +optional
	SnapshotTLSCacert string `json:"snapshotTLSCacert"`

	// Fixpacks is an array of strings which indicate the name of fixpacks
	// which should be installed in the deployment.  This corresponds to
	// setting the FIXPACKS environment variable in the deployment itself.
	// Cannot be updated.
	// +optional
	Fixpacks []string `json:"fixpacks,omitempty"`

	// Instance is the name of the Verify Identity Access instance which is being
	// started.  This value is only used for WRP and DSC deployments and is
	// ignored for Runtime deployments.
	// Defaults to 'default'.
	// Cannot be updated.
	// +optional
	Instance string `json:"instance"`

	// +kubebuilder:validation:Enum=zh_CN.utf8;zh_TW.utf8;cs_CZ.utf8;en_US.utf8;fr_FR.utf8;de_DE.utf8;hu_HU.utf8;it_IT.utf8;ja_JP.utf8;ko_KR.utf8;pl_PL.utf8;pt_BR.utf8;ru_RU.utf8;es_ES.utf8
	// Language is the language which will be used for messages which are logged
	// by the deployment.
	// Cannot be updated.
	// +optional
	Language Language `json:"language,omitempty" protobuf:"bytes,14,opt,name=language,casttype=Language"`

	// List of volumes that can be mounted by containers belonging to the pod.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge,retainKeys
	Volumes []corev1.Volume `json:"volumes,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name" protobuf:"bytes,1,rep,name=volumes"`

	// ImagePullSecrets is an optional list of references to secrets in the same
	// namespace to use for pulling any of the images used by this PodSpec.
	// If specified, these secrets will be passed to individual puller
	// implementations for them to use. For example,
	// in the case of docker, only DockerConfig type secrets are honored.
	// More info:
	// https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"`

	// ServiceAccountName is the name of the ServiceAccount to use to run this pod.
	// More info:
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty" protobuf:"bytes,8,opt,name=serviceAccountName"`

	// The set of custom annotations to add to the container being created.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=key
	// +patchStrategy=merge,
	CustomAnnotations []CustomAnnotation `json:"customAnnotations,omitempty" patchStrategy:"merge" patchMergeKey:"key" protobuf:"bytes,5,opt,name=customAnnotations"`

	// The IBM License Metric Tool annotations to add to runtime containers. Annotations are used by IBM to
	// track license usage for containerised environments.
	// +optional
	LicenseAnnotations *ILMTAnnotations `json:"ilmtAnnotations,omitempty" protobuf:"bytes,64,opt,name=ilmt_annotations,casttype=ILMTAnnotations"`

	// The definition for the container which is being created.
	// Cannot be updated.
	// +optional
	Container IBMSecurityVerifyAccessContainer `json:"container,omitempty"`
}

// IBMSecurityVerifyAccessStatus defines the observed state of an
// IBMSecurityVerifyAccess resource.
type IBMSecurityVerifyAccessStatus struct {
	// Conditions is the list of status conditions for this resource
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// IBMSecurityVerifyAccess is the Schema for the ibmsecurityverifyaccesses API.
// +kubebuilder:subresource:status
type IBMSecurityVerifyAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IBMSecurityVerifyAccessSpec   `json:"spec,omitempty"`
	Status IBMSecurityVerifyAccessStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IBMSecurityVerifyAccessList contains a list of IBMSecurityVerifyAccess
// resources.
type IBMSecurityVerifyAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IBMSecurityVerifyAccess `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&IBMSecurityVerifyAccess{}, &IBMSecurityVerifyAccessList{})
}
