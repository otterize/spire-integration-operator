// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package otterizegraphql

import (
	"context"
	"encoding/json"

	"github.com/Khan/genqlient/graphql"
)

type CertificateCustomization struct {
	DnsNames []*string `json:"dnsNames"`
	Ttl      *int      `json:"ttl"`
}

// GetDnsNames returns CertificateCustomization.DnsNames, and is useful for accessing the field via an interface.
func (v *CertificateCustomization) GetDnsNames() []*string { return v.DnsNames }

// GetTtl returns CertificateCustomization.Ttl, and is useful for accessing the field via an interface.
func (v *CertificateCustomization) GetTtl() *int { return v.Ttl }

// GetTLSKeyPairMockService includes the requested fields of the GraphQL type Service.
type GetTLSKeyPairMockService struct {
	TlsKeyPair *GetTLSKeyPairMockServiceTlsKeyPair `json:"tlsKeyPair"`
}

// GetTlsKeyPair returns GetTLSKeyPairMockService.TlsKeyPair, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairMockService) GetTlsKeyPair() *GetTLSKeyPairMockServiceTlsKeyPair {
	return v.TlsKeyPair
}

// GetTLSKeyPairMockServiceTlsKeyPair includes the requested fields of the GraphQL type KeyPair.
type GetTLSKeyPairMockServiceTlsKeyPair struct {
	TLSKeyPair `json:"-"`
}

// GetKeyPEM returns GetTLSKeyPairMockServiceTlsKeyPair.KeyPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairMockServiceTlsKeyPair) GetKeyPEM() string { return v.TLSKeyPair.KeyPEM }

// GetCertPEM returns GetTLSKeyPairMockServiceTlsKeyPair.CertPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairMockServiceTlsKeyPair) GetCertPEM() string { return v.TLSKeyPair.CertPEM }

// GetCaPEM returns GetTLSKeyPairMockServiceTlsKeyPair.CaPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairMockServiceTlsKeyPair) GetCaPEM() string { return v.TLSKeyPair.CaPEM }

// GetRootCAPEM returns GetTLSKeyPairMockServiceTlsKeyPair.RootCAPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairMockServiceTlsKeyPair) GetRootCAPEM() string { return v.TLSKeyPair.RootCAPEM }

// GetExpiresAt returns GetTLSKeyPairMockServiceTlsKeyPair.ExpiresAt, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairMockServiceTlsKeyPair) GetExpiresAt() int { return v.TLSKeyPair.ExpiresAt }

func (v *GetTLSKeyPairMockServiceTlsKeyPair) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*GetTLSKeyPairMockServiceTlsKeyPair
		graphql.NoUnmarshalJSON
	}
	firstPass.GetTLSKeyPairMockServiceTlsKeyPair = v

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

type __premarshalGetTLSKeyPairMockServiceTlsKeyPair struct {
	KeyPEM string `json:"keyPEM"`

	CertPEM string `json:"certPEM"`

	CaPEM string `json:"caPEM"`

	RootCAPEM string `json:"rootCAPEM"`

	ExpiresAt int `json:"expiresAt"`
}

func (v *GetTLSKeyPairMockServiceTlsKeyPair) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *GetTLSKeyPairMockServiceTlsKeyPair) __premarshalJSON() (*__premarshalGetTLSKeyPairMockServiceTlsKeyPair, error) {
	var retval __premarshalGetTLSKeyPairMockServiceTlsKeyPair

	retval.KeyPEM = v.TLSKeyPair.KeyPEM
	retval.CertPEM = v.TLSKeyPair.CertPEM
	retval.CaPEM = v.TLSKeyPair.CaPEM
	retval.RootCAPEM = v.TLSKeyPair.RootCAPEM
	retval.ExpiresAt = v.TLSKeyPair.ExpiresAt
	return &retval, nil
}

// GetTLSKeyPairResponse is returned by GetTLSKeyPair on success.
type GetTLSKeyPairResponse struct {
	MockService *GetTLSKeyPairMockService `json:"mockService"`
}

// GetMockService returns GetTLSKeyPairResponse.MockService, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairResponse) GetMockService() *GetTLSKeyPairMockService { return v.MockService }

// ReportKubernetesWorkloadReportKubernetesWorkloadService includes the requested fields of the GraphQL type Service.
type ReportKubernetesWorkloadReportKubernetesWorkloadService struct {
	Id string `json:"id"`
}

// GetId returns ReportKubernetesWorkloadReportKubernetesWorkloadService.Id, and is useful for accessing the field via an interface.
func (v *ReportKubernetesWorkloadReportKubernetesWorkloadService) GetId() string { return v.Id }

// ReportKubernetesWorkloadResponse is returned by ReportKubernetesWorkload on success.
type ReportKubernetesWorkloadResponse struct {
	ReportKubernetesWorkload ReportKubernetesWorkloadReportKubernetesWorkloadService `json:"reportKubernetesWorkload"`
}

// GetReportKubernetesWorkload returns ReportKubernetesWorkloadResponse.ReportKubernetesWorkload, and is useful for accessing the field via an interface.
func (v *ReportKubernetesWorkloadResponse) GetReportKubernetesWorkload() ReportKubernetesWorkloadReportKubernetesWorkloadService {
	return v.ReportKubernetesWorkload
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

// __GetTLSKeyPairInput is used internally by genqlient
type __GetTLSKeyPairInput struct {
	Id                        *string                   `json:"id"`
	Name                      *string                   `json:"name"`
	Namespace                 *string                   `json:"namespace"`
	CertificateCustomizations *CertificateCustomization `json:"certificateCustomizations"`
}

// GetId returns __GetTLSKeyPairInput.Id, and is useful for accessing the field via an interface.
func (v *__GetTLSKeyPairInput) GetId() *string { return v.Id }

// GetName returns __GetTLSKeyPairInput.Name, and is useful for accessing the field via an interface.
func (v *__GetTLSKeyPairInput) GetName() *string { return v.Name }

// GetNamespace returns __GetTLSKeyPairInput.Namespace, and is useful for accessing the field via an interface.
func (v *__GetTLSKeyPairInput) GetNamespace() *string { return v.Namespace }

// GetCertificateCustomizations returns __GetTLSKeyPairInput.CertificateCustomizations, and is useful for accessing the field via an interface.
func (v *__GetTLSKeyPairInput) GetCertificateCustomizations() *CertificateCustomization {
	return v.CertificateCustomizations
}

// __ReportKubernetesWorkloadInput is used internally by genqlient
type __ReportKubernetesWorkloadInput struct {
	Namespace    string `json:"namespace"`
	PodOwnerName string `json:"podOwnerName"`
}

// GetNamespace returns __ReportKubernetesWorkloadInput.Namespace, and is useful for accessing the field via an interface.
func (v *__ReportKubernetesWorkloadInput) GetNamespace() string { return v.Namespace }

// GetPodOwnerName returns __ReportKubernetesWorkloadInput.PodOwnerName, and is useful for accessing the field via an interface.
func (v *__ReportKubernetesWorkloadInput) GetPodOwnerName() string { return v.PodOwnerName }

func GetTLSKeyPair(
	ctx context.Context,
	client graphql.Client,
	id *string,
	name *string,
	namespace *string,
	certificateCustomizations *CertificateCustomization,
) (*GetTLSKeyPairResponse, error) {
	req := &graphql.Request{
		OpName: "GetTLSKeyPair",
		Query: `
query GetTLSKeyPair ($id: ID!, $name: String!, $namespace: String!, $certificateCustomizations: CertificateCustomization) {
	mockService(id: $id, name: $name, namespace: $namespace) {
		tlsKeyPair(certificateCustomization: $certificateCustomizations) {
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
			Id:                        id,
			Name:                      name,
			Namespace:                 namespace,
			CertificateCustomizations: certificateCustomizations,
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

func ReportKubernetesWorkload(
	ctx context.Context,
	client graphql.Client,
	namespace string,
	podOwnerName string,
) (*ReportKubernetesWorkloadResponse, error) {
	req := &graphql.Request{
		OpName: "ReportKubernetesWorkload",
		Query: `
mutation ReportKubernetesWorkload ($namespace: String!, $podOwnerName: String!) {
	reportKubernetesWorkload(namespace: $namespace, podOwnerName: $podOwnerName) {
		id
	}
}
`,
		Variables: &__ReportKubernetesWorkloadInput{
			Namespace:    namespace,
			PodOwnerName: podOwnerName,
		},
	}
	var err error

	var data ReportKubernetesWorkloadResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
