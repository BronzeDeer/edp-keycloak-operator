package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// KeycloakComponentSpec defines the desired state of KeycloakRealmComponent.
type KeycloakComponentSpec struct {
	// Name of keycloak component.
	Name string `json:"name"`

	// Realm is name of KeycloakRealm custom resource.
	Realm string `json:"realm"`

	// ProviderID is a provider ID of component.
	ProviderID string `json:"providerId"`

	// ProviderType is a provider type of component.
	ProviderType string `json:"providerType"`

	// Config is a map of component configuration.
	// +nullable
	// +optional
	Config map[string][]string `json:"config,omitempty"`
}

// KeycloakComponentStatus defines the observed state of KeycloakRealmComponent.
type KeycloakComponentStatus struct {
	// +optional
	Value string `json:"value,omitempty"`

	// +optional
	FailureCount int64 `json:"failureCount,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// KeycloakRealmComponent is the Schema for the keycloak component API.
type KeycloakRealmComponent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeycloakComponentSpec   `json:"spec,omitempty"`
	Status KeycloakComponentStatus `json:"status,omitempty"`
}

func (in *KeycloakRealmComponent) GetFailureCount() int64 {
	return in.Status.FailureCount
}

func (in *KeycloakRealmComponent) SetFailureCount(count int64) {
	in.Status.FailureCount = count
}

func (in *KeycloakRealmComponent) GetStatus() string {
	return in.Status.Value
}

func (in *KeycloakRealmComponent) SetStatus(value string) {
	in.Status.Value = value
}

func (in *KeycloakRealmComponent) K8SParentRealmName() (string, error) {
	return in.Spec.Realm, nil
}

// +kubebuilder:object:root=true

// KeycloakRealmComponentList contains a list of KeycloakRealmComponent.
type KeycloakRealmComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KeycloakRealmComponent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeycloakRealmComponent{}, &KeycloakRealmComponentList{})
}
