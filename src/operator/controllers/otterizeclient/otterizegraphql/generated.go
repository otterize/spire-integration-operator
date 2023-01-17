// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package otterizegraphql

import (
	"context"
	"encoding/json"

	"github.com/Khan/genqlient/graphql"
)

type CertificateCustomization struct {
	DnsNames []string `json:"dnsNames"`
	Ttl      int      `json:"ttl"`
}

// GetDnsNames returns CertificateCustomization.DnsNames, and is useful for accessing the field via an interface.
func (v *CertificateCustomization) GetDnsNames() []string { return v.DnsNames }

// GetTtl returns CertificateCustomization.Ttl, and is useful for accessing the field via an interface.
func (v *CertificateCustomization) GetTtl() int { return v.Ttl }

// CleanupOrphanK8SPodEntriesResponse is returned by CleanupOrphanK8SPodEntries on success.
type CleanupOrphanK8SPodEntriesResponse struct {
	// Removes certificateRequests of pod owners from the context's cluster which does not appear in the "activePodOwners" list
	RemoveOrphanedCertificateRequests bool `json:"removeOrphanedCertificateRequests"`
}

// GetRemoveOrphanedCertificateRequests returns CleanupOrphanK8SPodEntriesResponse.RemoveOrphanedCertificateRequests, and is useful for accessing the field via an interface.
func (v *CleanupOrphanK8SPodEntriesResponse) GetRemoveOrphanedCertificateRequests() bool {
	return v.RemoveOrphanedCertificateRequests
}

type ComponentType string

const (
	ComponentTypeIntentsOperator     ComponentType = "INTENTS_OPERATOR"
	ComponentTypeCredentialsOperator ComponentType = "CREDENTIALS_OPERATOR"
	ComponentTypeNetworkMapper       ComponentType = "NETWORK_MAPPER"
)

// GetTLSKeyPairResponse is returned by GetTLSKeyPair on success.
type GetTLSKeyPairResponse struct {
	// Get service
	Service *GetTLSKeyPairService `json:"service"`
}

// GetService returns GetTLSKeyPairResponse.Service, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairResponse) GetService() *GetTLSKeyPairService { return v.Service }

// GetTLSKeyPairService includes the requested fields of the GraphQL type Service.
type GetTLSKeyPairService struct {
	TlsKeyPair *GetTLSKeyPairServiceTlsKeyPair `json:"tlsKeyPair"`
}

// GetTlsKeyPair returns GetTLSKeyPairService.TlsKeyPair, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairService) GetTlsKeyPair() *GetTLSKeyPairServiceTlsKeyPair { return v.TlsKeyPair }

// GetTLSKeyPairServiceTlsKeyPair includes the requested fields of the GraphQL type KeyPair.
type GetTLSKeyPairServiceTlsKeyPair struct {
	TLSKeyPair `json:"-"`
}

// GetKeyPEM returns GetTLSKeyPairServiceTlsKeyPair.KeyPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairServiceTlsKeyPair) GetKeyPEM() string { return v.TLSKeyPair.KeyPEM }

// GetCertPEM returns GetTLSKeyPairServiceTlsKeyPair.CertPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairServiceTlsKeyPair) GetCertPEM() string { return v.TLSKeyPair.CertPEM }

// GetCaPEM returns GetTLSKeyPairServiceTlsKeyPair.CaPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairServiceTlsKeyPair) GetCaPEM() string { return v.TLSKeyPair.CaPEM }

// GetRootCAPEM returns GetTLSKeyPairServiceTlsKeyPair.RootCAPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairServiceTlsKeyPair) GetRootCAPEM() string { return v.TLSKeyPair.RootCAPEM }

// GetExpiresAt returns GetTLSKeyPairServiceTlsKeyPair.ExpiresAt, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairServiceTlsKeyPair) GetExpiresAt() int { return v.TLSKeyPair.ExpiresAt }

func (v *GetTLSKeyPairServiceTlsKeyPair) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*GetTLSKeyPairServiceTlsKeyPair
		graphql.NoUnmarshalJSON
	}
	firstPass.GetTLSKeyPairServiceTlsKeyPair = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.TLSKeyPair)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalGetTLSKeyPairServiceTlsKeyPair struct {
	KeyPEM string `json:"keyPEM"`

	CertPEM string `json:"certPEM"`

	CaPEM string `json:"caPEM"`

	RootCAPEM string `json:"rootCAPEM"`

	ExpiresAt int `json:"expiresAt"`
}

func (v *GetTLSKeyPairServiceTlsKeyPair) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *GetTLSKeyPairServiceTlsKeyPair) __premarshalJSON() (*__premarshalGetTLSKeyPairServiceTlsKeyPair, error) {
	var retval __premarshalGetTLSKeyPairServiceTlsKeyPair

	retval.KeyPEM = v.TLSKeyPair.KeyPEM
	retval.CertPEM = v.TLSKeyPair.CertPEM
	retval.CaPEM = v.TLSKeyPair.CaPEM
	retval.RootCAPEM = v.TLSKeyPair.RootCAPEM
	retval.ExpiresAt = v.TLSKeyPair.ExpiresAt
	return &retval, nil
}

type NamespacedPodOwner struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// GetName returns NamespacedPodOwner.Name, and is useful for accessing the field via an interface.
func (v *NamespacedPodOwner) GetName() string { return v.Name }

// GetNamespace returns NamespacedPodOwner.Namespace, and is useful for accessing the field via an interface.
func (v *NamespacedPodOwner) GetNamespace() string { return v.Namespace }

// RegisterKubernetesPodOwnerCertificateRequestRegisterKubernetesPodOwnerCertificateRequestService includes the requested fields of the GraphQL type Service.
type RegisterKubernetesPodOwnerCertificateRequestRegisterKubernetesPodOwnerCertificateRequestService struct {
	Id string `json:"id"`
}

// GetId returns RegisterKubernetesPodOwnerCertificateRequestRegisterKubernetesPodOwnerCertificateRequestService.Id, and is useful for accessing the field via an interface.
func (v *RegisterKubernetesPodOwnerCertificateRequestRegisterKubernetesPodOwnerCertificateRequestService) GetId() string {
	return v.Id
}

// RegisterKubernetesPodOwnerCertificateRequestResponse is returned by RegisterKubernetesPodOwnerCertificateRequest on success.
type RegisterKubernetesPodOwnerCertificateRequestResponse struct {
	// Register certificate-request details for kubernetes pod owner, returns the service associated with this pod owner
	RegisterKubernetesPodOwnerCertificateRequest RegisterKubernetesPodOwnerCertificateRequestRegisterKubernetesPodOwnerCertificateRequestService `json:"registerKubernetesPodOwnerCertificateRequest"`
}

// GetRegisterKubernetesPodOwnerCertificateRequest returns RegisterKubernetesPodOwnerCertificateRequestResponse.RegisterKubernetesPodOwnerCertificateRequest, and is useful for accessing the field via an interface.
func (v *RegisterKubernetesPodOwnerCertificateRequestResponse) GetRegisterKubernetesPodOwnerCertificateRequest() RegisterKubernetesPodOwnerCertificateRequestRegisterKubernetesPodOwnerCertificateRequestService {
	return v.RegisterKubernetesPodOwnerCertificateRequest
}

// ReportComponentStatusResponse is returned by ReportComponentStatus on success.
type ReportComponentStatusResponse struct {
	ReportIntegrationComponentStatus bool `json:"reportIntegrationComponentStatus"`
}

// GetReportIntegrationComponentStatus returns ReportComponentStatusResponse.ReportIntegrationComponentStatus, and is useful for accessing the field via an interface.
func (v *ReportComponentStatusResponse) GetReportIntegrationComponentStatus() bool {
	return v.ReportIntegrationComponentStatus
}

// TLSKeyPair includes the GraphQL fields of KeyPair requested by the fragment TLSKeyPair.
type TLSKeyPair struct {
	KeyPEM    string `json:"keyPEM"`
	CertPEM   string `json:"certPEM"`
	CaPEM     string `json:"caPEM"`
	RootCAPEM string `json:"rootCAPEM"`
	ExpiresAt int    `json:"expiresAt"`
}

// GetKeyPEM returns TLSKeyPair.KeyPEM, and is useful for accessing the field via an interface.
func (v *TLSKeyPair) GetKeyPEM() string { return v.KeyPEM }

// GetCertPEM returns TLSKeyPair.CertPEM, and is useful for accessing the field via an interface.
func (v *TLSKeyPair) GetCertPEM() string { return v.CertPEM }

// GetCaPEM returns TLSKeyPair.CaPEM, and is useful for accessing the field via an interface.
func (v *TLSKeyPair) GetCaPEM() string { return v.CaPEM }

// GetRootCAPEM returns TLSKeyPair.RootCAPEM, and is useful for accessing the field via an interface.
func (v *TLSKeyPair) GetRootCAPEM() string { return v.RootCAPEM }

// GetExpiresAt returns TLSKeyPair.ExpiresAt, and is useful for accessing the field via an interface.
func (v *TLSKeyPair) GetExpiresAt() int { return v.ExpiresAt }

// __CleanupOrphanK8SPodEntriesInput is used internally by genqlient
type __CleanupOrphanK8SPodEntriesInput struct {
	ExistingPodOwners []NamespacedPodOwner `json:"existingPodOwners"`
}

// GetExistingPodOwners returns __CleanupOrphanK8SPodEntriesInput.ExistingPodOwners, and is useful for accessing the field via an interface.
func (v *__CleanupOrphanK8SPodEntriesInput) GetExistingPodOwners() []NamespacedPodOwner {
	return v.ExistingPodOwners
}

// __GetTLSKeyPairInput is used internally by genqlient
type __GetTLSKeyPairInput struct {
	Id *string `json:"id"`
}

// GetId returns __GetTLSKeyPairInput.Id, and is useful for accessing the field via an interface.
func (v *__GetTLSKeyPairInput) GetId() *string { return v.Id }

// __RegisterKubernetesPodOwnerCertificateRequestInput is used internally by genqlient
type __RegisterKubernetesPodOwnerCertificateRequestInput struct {
	Namespace                 string                   `json:"namespace"`
	PodOwnerName              string                   `json:"podOwnerName"`
	CertificateCustomizations CertificateCustomization `json:"certificateCustomizations"`
}

// GetNamespace returns __RegisterKubernetesPodOwnerCertificateRequestInput.Namespace, and is useful for accessing the field via an interface.
func (v *__RegisterKubernetesPodOwnerCertificateRequestInput) GetNamespace() string {
	return v.Namespace
}

// GetPodOwnerName returns __RegisterKubernetesPodOwnerCertificateRequestInput.PodOwnerName, and is useful for accessing the field via an interface.
func (v *__RegisterKubernetesPodOwnerCertificateRequestInput) GetPodOwnerName() string {
	return v.PodOwnerName
}

// GetCertificateCustomizations returns __RegisterKubernetesPodOwnerCertificateRequestInput.CertificateCustomizations, and is useful for accessing the field via an interface.
func (v *__RegisterKubernetesPodOwnerCertificateRequestInput) GetCertificateCustomizations() CertificateCustomization {
	return v.CertificateCustomizations
}

// __ReportComponentStatusInput is used internally by genqlient
type __ReportComponentStatusInput struct {
	Component ComponentType `json:"component"`
}

// GetComponent returns __ReportComponentStatusInput.Component, and is useful for accessing the field via an interface.
func (v *__ReportComponentStatusInput) GetComponent() ComponentType { return v.Component }

func CleanupOrphanK8SPodEntries(
	ctx context.Context,
	client graphql.Client,
	existingPodOwners []NamespacedPodOwner,
) (*CleanupOrphanK8SPodEntriesResponse, error) {
	req := &graphql.Request{
		OpName: "CleanupOrphanK8SPodEntries",
		Query: `
mutation CleanupOrphanK8SPodEntries ($existingPodOwners: [NamespacedPodOwner!]!) {
	removeOrphanedCertificateRequests(activePodOwners: $existingPodOwners)
}
`,
		Variables: &__CleanupOrphanK8SPodEntriesInput{
			ExistingPodOwners: existingPodOwners,
		},
	}
	var err error

	var data CleanupOrphanK8SPodEntriesResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetTLSKeyPair(
	ctx context.Context,
	client graphql.Client,
	id *string,
) (*GetTLSKeyPairResponse, error) {
	req := &graphql.Request{
		OpName: "GetTLSKeyPair",
		Query: `
query GetTLSKeyPair ($id: ID!) {
	service(id: $id) {
		tlsKeyPair {
			... TLSKeyPair
		}
	}
}
fragment TLSKeyPair on KeyPair {
	keyPEM
	certPEM
	caPEM
	rootCAPEM
	expiresAt
}
`,
		Variables: &__GetTLSKeyPairInput{
			Id: id,
		},
	}
	var err error

	var data GetTLSKeyPairResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func RegisterKubernetesPodOwnerCertificateRequest(
	ctx context.Context,
	client graphql.Client,
	namespace string,
	podOwnerName string,
	certificateCustomizations CertificateCustomization,
) (*RegisterKubernetesPodOwnerCertificateRequestResponse, error) {
	req := &graphql.Request{
		OpName: "RegisterKubernetesPodOwnerCertificateRequest",
		Query: `
mutation RegisterKubernetesPodOwnerCertificateRequest ($namespace: String!, $podOwnerName: String!, $certificateCustomizations: CertificateCustomization) {
	registerKubernetesPodOwnerCertificateRequest(namespace: $namespace, podOwnerName: $podOwnerName, certificateCustomization: $certificateCustomizations) {
		id
	}
}
`,
		Variables: &__RegisterKubernetesPodOwnerCertificateRequestInput{
			Namespace:                 namespace,
			PodOwnerName:              podOwnerName,
			CertificateCustomizations: certificateCustomizations,
		},
	}
	var err error

	var data RegisterKubernetesPodOwnerCertificateRequestResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func ReportComponentStatus(
	ctx context.Context,
	client graphql.Client,
	component ComponentType,
) (*ReportComponentStatusResponse, error) {
	req := &graphql.Request{
		OpName: "ReportComponentStatus",
		Query: `
mutation ReportComponentStatus ($component: ComponentType!) {
	reportIntegrationComponentStatus(component: $component)
}
`,
		Variables: &__ReportComponentStatusInput{
			Component: component,
		},
	}
	var err error

	var data ReportComponentStatusResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
