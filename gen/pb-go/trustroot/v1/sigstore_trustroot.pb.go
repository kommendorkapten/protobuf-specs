// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: sigstore_trustroot.proto

package v1

import (
	v1 "github.com/sigstore/protobuf-specs/gen/pb-go/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TransparencyLogInstance describes the immutable parameters from a
// transparency log.
// See https://www.rfc-editor.org/rfc/rfc9162.html#name-log-parameters
// for more details.
// The incluced parameters are the minimal set required to identify a log,
// and verify an inclusion promise.
type TransparencyLogInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The base URL at which can be used to URLs for the client.
	BaseUrl string `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	// The hash algorithm used for the Merkle Tree.
	HashAlgorithm v1.HashAlgorithm `protobuf:"varint,2,opt,name=hash_algorithm,json=hashAlgorithm,proto3,enum=dev.sigstore.common.v1.HashAlgorithm" json:"hash_algorithm,omitempty"`
	// The public key used to verify signatures generated by the log.
	// This attribute contains the signature algorithm used by the log.
	PublicKey *v1.PublicKey `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The unique identifier for this transparency log.
	LogId *v1.LogId `protobuf:"bytes,4,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
}

func (x *TransparencyLogInstance) Reset() {
	*x = TransparencyLogInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sigstore_trustroot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransparencyLogInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransparencyLogInstance) ProtoMessage() {}

func (x *TransparencyLogInstance) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_trustroot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransparencyLogInstance.ProtoReflect.Descriptor instead.
func (*TransparencyLogInstance) Descriptor() ([]byte, []int) {
	return file_sigstore_trustroot_proto_rawDescGZIP(), []int{0}
}

func (x *TransparencyLogInstance) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *TransparencyLogInstance) GetHashAlgorithm() v1.HashAlgorithm {
	if x != nil {
		return x.HashAlgorithm
	}
	return v1.HashAlgorithm(0)
}

func (x *TransparencyLogInstance) GetPublicKey() *v1.PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *TransparencyLogInstance) GetLogId() *v1.LogId {
	if x != nil {
		return x.LogId
	}
	return nil
}

// CertificateAuthority enlists the information required to identify which
// CA to use and perform signature verification.
type CertificateAuthority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The root certificate MUST be self-signed, and so the subject and
	// issuer are the same.
	Subject *v1.DistinguishedName `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// The URI at which the CA can be accessed.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// The certificate chain for this CA.
	CertChain *v1.X509CertificateChain `protobuf:"bytes,3,opt,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	// The time the *entire* chain was valid. This is at max the
	// longest interval when *all* certificates in the chain were valid,
	// but it MAY be shorter.
	ValidFor *v1.TimeRange `protobuf:"bytes,4,opt,name=valid_for,json=validFor,proto3" json:"valid_for,omitempty"`
}

func (x *CertificateAuthority) Reset() {
	*x = CertificateAuthority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sigstore_trustroot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateAuthority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateAuthority) ProtoMessage() {}

func (x *CertificateAuthority) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_trustroot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateAuthority.ProtoReflect.Descriptor instead.
func (*CertificateAuthority) Descriptor() ([]byte, []int) {
	return file_sigstore_trustroot_proto_rawDescGZIP(), []int{1}
}

func (x *CertificateAuthority) GetSubject() *v1.DistinguishedName {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *CertificateAuthority) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *CertificateAuthority) GetCertChain() *v1.X509CertificateChain {
	if x != nil {
		return x.CertChain
	}
	return nil
}

func (x *CertificateAuthority) GetValidFor() *v1.TimeRange {
	if x != nil {
		return x.ValidFor
	}
	return nil
}

// TrustedRoot describes the client's complete set of trusted entities.
// How the TrustedRoot is populated is not specified, but can be a
// combination of many sources such as TUF repositories, files on disk etc.
//
// The TrustedRoot is not meant to be used for any artifact verification, only
// to capture the complete/global set of trusted verification materials.
// When verifying an artifact, based on the artifact and policies, a selection
// of keys/authorities are expected to be extracted and provided to the
// verification function. This way the set of keys/authorities kan be kept to
// a minimal set by the policy to gain better control over what signatures
// that are allowed.
type TrustedRoot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A set of trusted Rekor servers.
	Tlogs []*TransparencyLogInstance `protobuf:"bytes,1,rep,name=tlogs,proto3" json:"tlogs,omitempty"`
	// A set of trusted certificate authorites (e.g Fulcio), and any
	// intermediate certificates they provide.
	// If a CA is issuing multiple intermediate certificate, each
	// combination shall be represented as separate chain. I.e, a single
	// root cert may appear in multiple chains but with different
	// intermediate and/or leaf certificates.
	// The certificates are intended to be used for verifying artifact
	// signatures.
	CertificateAuthorities []*CertificateAuthority `protobuf:"bytes,2,rep,name=certificate_authorities,json=certificateAuthorities,proto3" json:"certificate_authorities,omitempty"`
	// A set of trusted certificate transparency logs.
	Ctlogs []*TransparencyLogInstance `protobuf:"bytes,3,rep,name=ctlogs,proto3" json:"ctlogs,omitempty"`
	// A set of trusted timestamping authorities.
	TimestampAuthorities []*CertificateAuthority `protobuf:"bytes,4,rep,name=timestamp_authorities,json=timestampAuthorities,proto3" json:"timestamp_authorities,omitempty"`
}

func (x *TrustedRoot) Reset() {
	*x = TrustedRoot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sigstore_trustroot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrustedRoot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustedRoot) ProtoMessage() {}

func (x *TrustedRoot) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_trustroot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustedRoot.ProtoReflect.Descriptor instead.
func (*TrustedRoot) Descriptor() ([]byte, []int) {
	return file_sigstore_trustroot_proto_rawDescGZIP(), []int{2}
}

func (x *TrustedRoot) GetTlogs() []*TransparencyLogInstance {
	if x != nil {
		return x.Tlogs
	}
	return nil
}

func (x *TrustedRoot) GetCertificateAuthorities() []*CertificateAuthority {
	if x != nil {
		return x.CertificateAuthorities
	}
	return nil
}

func (x *TrustedRoot) GetCtlogs() []*TransparencyLogInstance {
	if x != nil {
		return x.Ctlogs
	}
	return nil
}

func (x *TrustedRoot) GetTimestampAuthorities() []*CertificateAuthority {
	if x != nil {
		return x.TimestampAuthorities
	}
	return nil
}

// TransparencyLogIdentifier contains information that can be used to
// identify a `TransparencyLogInstance`, either by its log id or its base_uri.
type TransparencyLogIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Id:
	//	*TransparencyLogIdentifier_LogId
	//	*TransparencyLogIdentifier_Uri
	Id isTransparencyLogIdentifier_Id `protobuf_oneof:"id"`
}

func (x *TransparencyLogIdentifier) Reset() {
	*x = TransparencyLogIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sigstore_trustroot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransparencyLogIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransparencyLogIdentifier) ProtoMessage() {}

func (x *TransparencyLogIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_trustroot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransparencyLogIdentifier.ProtoReflect.Descriptor instead.
func (*TransparencyLogIdentifier) Descriptor() ([]byte, []int) {
	return file_sigstore_trustroot_proto_rawDescGZIP(), []int{3}
}

func (m *TransparencyLogIdentifier) GetId() isTransparencyLogIdentifier_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *TransparencyLogIdentifier) GetLogId() *v1.LogId {
	if x, ok := x.GetId().(*TransparencyLogIdentifier_LogId); ok {
		return x.LogId
	}
	return nil
}

func (x *TransparencyLogIdentifier) GetUri() string {
	if x, ok := x.GetId().(*TransparencyLogIdentifier_Uri); ok {
		return x.Uri
	}
	return ""
}

type isTransparencyLogIdentifier_Id interface {
	isTransparencyLogIdentifier_Id()
}

type TransparencyLogIdentifier_LogId struct {
	// The log id of the transparency log.
	LogId *v1.LogId `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3,oneof"`
}

type TransparencyLogIdentifier_Uri struct {
	// The base_uri for the transparency log.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof"`
}

func (*TransparencyLogIdentifier_LogId) isTransparencyLogIdentifier_Id() {}

func (*TransparencyLogIdentifier_Uri) isTransparencyLogIdentifier_Id() {}

// CertificateAuthorityIdentifier contains information that can be used to
// identify a `CertificateAuthority`, either by its subject or by its uri.
type CertificateAuthorityIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Id:
	//	*CertificateAuthorityIdentifier_Subject
	//	*CertificateAuthorityIdentifier_Uri
	Id isCertificateAuthorityIdentifier_Id `protobuf_oneof:"id"`
}

func (x *CertificateAuthorityIdentifier) Reset() {
	*x = CertificateAuthorityIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sigstore_trustroot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateAuthorityIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateAuthorityIdentifier) ProtoMessage() {}

func (x *CertificateAuthorityIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_trustroot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateAuthorityIdentifier.ProtoReflect.Descriptor instead.
func (*CertificateAuthorityIdentifier) Descriptor() ([]byte, []int) {
	return file_sigstore_trustroot_proto_rawDescGZIP(), []int{4}
}

func (m *CertificateAuthorityIdentifier) GetId() isCertificateAuthorityIdentifier_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *CertificateAuthorityIdentifier) GetSubject() *v1.DistinguishedName {
	if x, ok := x.GetId().(*CertificateAuthorityIdentifier_Subject); ok {
		return x.Subject
	}
	return nil
}

func (x *CertificateAuthorityIdentifier) GetUri() string {
	if x, ok := x.GetId().(*CertificateAuthorityIdentifier_Uri); ok {
		return x.Uri
	}
	return ""
}

type isCertificateAuthorityIdentifier_Id interface {
	isCertificateAuthorityIdentifier_Id()
}

type CertificateAuthorityIdentifier_Subject struct {
	// The subject of the certificate authority.
	Subject *v1.DistinguishedName `protobuf:"bytes,1,opt,name=subject,proto3,oneof"`
}

type CertificateAuthorityIdentifier_Uri struct {
	// The uri of the certificate authority.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof"`
}

func (*CertificateAuthorityIdentifier_Subject) isCertificateAuthorityIdentifier_Id() {}

func (*CertificateAuthorityIdentifier_Uri) isCertificateAuthorityIdentifier_Id() {}

// Environment acts a selector to filter down a trust root to a smaller one.
// A policy should reference an environment, or embed one that can be used
// during artifact verification to filter the global trust root into a one
// only containing the relevant instances. An environment does not take
// temporality into account when selecting the instances from a trust root.
// Prior to performing verification the temporal aspect should be further
// examined to reduce the set of instances.
type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the environment.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Ids for artifact signature transparency log to include.
	Tlog []*TransparencyLogIdentifier `protobuf:"bytes,2,rep,name=tlog,proto3" json:"tlog,omitempty"`
	// Ids for certificate authorities to include.
	Cas []*CertificateAuthorityIdentifier `protobuf:"bytes,3,rep,name=cas,proto3" json:"cas,omitempty"`
	// Ids for certificate transparency logs to include.
	Ctlogs []*TransparencyLogIdentifier `protobuf:"bytes,4,rep,name=ctlogs,proto3" json:"ctlogs,omitempty"`
	// Ids for timestamp authorities to include.
	Tsas []*CertificateAuthorityIdentifier `protobuf:"bytes,5,rep,name=tsas,proto3" json:"tsas,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sigstore_trustroot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_trustroot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_sigstore_trustroot_proto_rawDescGZIP(), []int{5}
}

func (x *Environment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Environment) GetTlog() []*TransparencyLogIdentifier {
	if x != nil {
		return x.Tlog
	}
	return nil
}

func (x *Environment) GetCas() []*CertificateAuthorityIdentifier {
	if x != nil {
		return x.Cas
	}
	return nil
}

func (x *Environment) GetCtlogs() []*TransparencyLogIdentifier {
	if x != nil {
		return x.Ctlogs
	}
	return nil
}

func (x *Environment) GetTsas() []*CertificateAuthorityIdentifier {
	if x != nil {
		return x.Tsas
	}
	return nil
}

// Policy to apply to a trust root prior to creating the verification Input
// when calling a verifier.
// Not that the Policy does NOT contain any reference to a bundle. How this
// mapping happens is specific for each client.
type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Environment *Environment `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
	// Types that are assignable to Signers:
	//	*Policy_CertificateIdentities
	//	*Policy_PublicKeys
	Signers isPolicy_Signers `protobuf_oneof:"signers"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sigstore_trustroot_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_trustroot_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_sigstore_trustroot_proto_rawDescGZIP(), []int{6}
}

func (x *Policy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policy) GetEnvironment() *Environment {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (m *Policy) GetSigners() isPolicy_Signers {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (x *Policy) GetCertificateIdentities() *v1.CertificateIdentities {
	if x, ok := x.GetSigners().(*Policy_CertificateIdentities); ok {
		return x.CertificateIdentities
	}
	return nil
}

func (x *Policy) GetPublicKeys() *v1.PublicKeyIdentities {
	if x, ok := x.GetSigners().(*Policy_PublicKeys); ok {
		return x.PublicKeys
	}
	return nil
}

type isPolicy_Signers interface {
	isPolicy_Signers()
}

type Policy_CertificateIdentities struct {
	CertificateIdentities *v1.CertificateIdentities `protobuf:"bytes,3,opt,name=certificate_identities,json=certificateIdentities,proto3,oneof"`
}

type Policy_PublicKeys struct {
	PublicKeys *v1.PublicKeyIdentities `protobuf:"bytes,4,opt,name=public_keys,json=publicKeys,proto3,oneof"`
}

func (*Policy_CertificateIdentities) isPolicy_Signers() {}

func (*Policy_PublicKeys) isPolicy_Signers() {}

var File_sigstore_trustroot_proto protoreflect.FileDescriptor

var file_sigstore_trustroot_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x64, 0x65, 0x76, 0x2e,
	0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x72, 0x6f,
	0x6f, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a,
	0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x4c, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x49, 0x64, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0xfa, 0x01, 0x0a, 0x14, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x43, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x75, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x4b, 0x0a, 0x0a, 0x63, 0x65, 0x72,
	0x74, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x35, 0x30, 0x39, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x09, 0x63, 0x65, 0x72,
	0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x3e, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f,
	0x66, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x08, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x46, 0x6f, 0x72, 0x22, 0xf3, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x75, 0x73, 0x74,
	0x65, 0x64, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x48, 0x0a, 0x05, 0x74, 0x6c, 0x6f, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x05, 0x74, 0x6c, 0x6f, 0x67, 0x73,
	0x12, 0x68, 0x0a, 0x17, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x16, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x06, 0x63, 0x74,
	0x6c, 0x6f, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x72,
	0x6f, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x06,
	0x63, 0x74, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x64, 0x0a, 0x15, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x14, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x6d, 0x0a, 0x19,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x6f, 0x67, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x06, 0x6c, 0x6f, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x03, 0x75, 0x72, 0x69, 0x42, 0x04, 0x0a, 0x02, 0x69, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x1e,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x45,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x75, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x69, 0x42, 0x04, 0x0a, 0x02, 0x69, 0x64, 0x22,
	0xd5, 0x02, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x04, 0x74, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x04, 0x74, 0x6c, 0x6f, 0x67, 0x12, 0x4b, 0x0a,
	0x03, 0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x72,
	0x6f, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x03, 0x63, 0x61, 0x73, 0x12, 0x4c, 0x0a, 0x06, 0x63, 0x74,
	0x6c, 0x6f, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x72,
	0x6f, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x52, 0x06, 0x63, 0x74, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x4d, 0x0a, 0x04, 0x74, 0x73, 0x61, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x72, 0x6f, 0x6f, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x52, 0x04, 0x74, 0x73, 0x61, 0x73, 0x22, 0xa9, 0x02, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x66, 0x0a, 0x16, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x48,
	0x00, 0x52, 0x15, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x73, 0x42, 0x6e, 0x0a, 0x1f, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x72,
	0x6f, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x6f, 0x6f,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x72, 0x6f, 0x6f, 0x74,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sigstore_trustroot_proto_rawDescOnce sync.Once
	file_sigstore_trustroot_proto_rawDescData = file_sigstore_trustroot_proto_rawDesc
)

func file_sigstore_trustroot_proto_rawDescGZIP() []byte {
	file_sigstore_trustroot_proto_rawDescOnce.Do(func() {
		file_sigstore_trustroot_proto_rawDescData = protoimpl.X.CompressGZIP(file_sigstore_trustroot_proto_rawDescData)
	})
	return file_sigstore_trustroot_proto_rawDescData
}

var file_sigstore_trustroot_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sigstore_trustroot_proto_goTypes = []interface{}{
	(*TransparencyLogInstance)(nil),        // 0: dev.sigstore.trustroot.v1.TransparencyLogInstance
	(*CertificateAuthority)(nil),           // 1: dev.sigstore.trustroot.v1.CertificateAuthority
	(*TrustedRoot)(nil),                    // 2: dev.sigstore.trustroot.v1.TrustedRoot
	(*TransparencyLogIdentifier)(nil),      // 3: dev.sigstore.trustroot.v1.TransparencyLogIdentifier
	(*CertificateAuthorityIdentifier)(nil), // 4: dev.sigstore.trustroot.v1.CertificateAuthorityIdentifier
	(*Environment)(nil),                    // 5: dev.sigstore.trustroot.v1.Environment
	(*Policy)(nil),                         // 6: dev.sigstore.trustroot.v1.Policy
	(v1.HashAlgorithm)(0),                  // 7: dev.sigstore.common.v1.HashAlgorithm
	(*v1.PublicKey)(nil),                   // 8: dev.sigstore.common.v1.PublicKey
	(*v1.LogId)(nil),                       // 9: dev.sigstore.common.v1.LogId
	(*v1.DistinguishedName)(nil),           // 10: dev.sigstore.common.v1.DistinguishedName
	(*v1.X509CertificateChain)(nil),        // 11: dev.sigstore.common.v1.X509CertificateChain
	(*v1.TimeRange)(nil),                   // 12: dev.sigstore.common.v1.TimeRange
	(*v1.CertificateIdentities)(nil),       // 13: dev.sigstore.common.v1.CertificateIdentities
	(*v1.PublicKeyIdentities)(nil),         // 14: dev.sigstore.common.v1.PublicKeyIdentities
}
var file_sigstore_trustroot_proto_depIdxs = []int32{
	7,  // 0: dev.sigstore.trustroot.v1.TransparencyLogInstance.hash_algorithm:type_name -> dev.sigstore.common.v1.HashAlgorithm
	8,  // 1: dev.sigstore.trustroot.v1.TransparencyLogInstance.public_key:type_name -> dev.sigstore.common.v1.PublicKey
	9,  // 2: dev.sigstore.trustroot.v1.TransparencyLogInstance.log_id:type_name -> dev.sigstore.common.v1.LogId
	10, // 3: dev.sigstore.trustroot.v1.CertificateAuthority.subject:type_name -> dev.sigstore.common.v1.DistinguishedName
	11, // 4: dev.sigstore.trustroot.v1.CertificateAuthority.cert_chain:type_name -> dev.sigstore.common.v1.X509CertificateChain
	12, // 5: dev.sigstore.trustroot.v1.CertificateAuthority.valid_for:type_name -> dev.sigstore.common.v1.TimeRange
	0,  // 6: dev.sigstore.trustroot.v1.TrustedRoot.tlogs:type_name -> dev.sigstore.trustroot.v1.TransparencyLogInstance
	1,  // 7: dev.sigstore.trustroot.v1.TrustedRoot.certificate_authorities:type_name -> dev.sigstore.trustroot.v1.CertificateAuthority
	0,  // 8: dev.sigstore.trustroot.v1.TrustedRoot.ctlogs:type_name -> dev.sigstore.trustroot.v1.TransparencyLogInstance
	1,  // 9: dev.sigstore.trustroot.v1.TrustedRoot.timestamp_authorities:type_name -> dev.sigstore.trustroot.v1.CertificateAuthority
	9,  // 10: dev.sigstore.trustroot.v1.TransparencyLogIdentifier.log_id:type_name -> dev.sigstore.common.v1.LogId
	10, // 11: dev.sigstore.trustroot.v1.CertificateAuthorityIdentifier.subject:type_name -> dev.sigstore.common.v1.DistinguishedName
	3,  // 12: dev.sigstore.trustroot.v1.Environment.tlog:type_name -> dev.sigstore.trustroot.v1.TransparencyLogIdentifier
	4,  // 13: dev.sigstore.trustroot.v1.Environment.cas:type_name -> dev.sigstore.trustroot.v1.CertificateAuthorityIdentifier
	3,  // 14: dev.sigstore.trustroot.v1.Environment.ctlogs:type_name -> dev.sigstore.trustroot.v1.TransparencyLogIdentifier
	4,  // 15: dev.sigstore.trustroot.v1.Environment.tsas:type_name -> dev.sigstore.trustroot.v1.CertificateAuthorityIdentifier
	5,  // 16: dev.sigstore.trustroot.v1.Policy.environment:type_name -> dev.sigstore.trustroot.v1.Environment
	13, // 17: dev.sigstore.trustroot.v1.Policy.certificate_identities:type_name -> dev.sigstore.common.v1.CertificateIdentities
	14, // 18: dev.sigstore.trustroot.v1.Policy.public_keys:type_name -> dev.sigstore.common.v1.PublicKeyIdentities
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_sigstore_trustroot_proto_init() }
func file_sigstore_trustroot_proto_init() {
	if File_sigstore_trustroot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sigstore_trustroot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransparencyLogInstance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sigstore_trustroot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateAuthority); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sigstore_trustroot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrustedRoot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sigstore_trustroot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransparencyLogIdentifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sigstore_trustroot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateAuthorityIdentifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sigstore_trustroot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sigstore_trustroot_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_sigstore_trustroot_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*TransparencyLogIdentifier_LogId)(nil),
		(*TransparencyLogIdentifier_Uri)(nil),
	}
	file_sigstore_trustroot_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*CertificateAuthorityIdentifier_Subject)(nil),
		(*CertificateAuthorityIdentifier_Uri)(nil),
	}
	file_sigstore_trustroot_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Policy_CertificateIdentities)(nil),
		(*Policy_PublicKeys)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sigstore_trustroot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sigstore_trustroot_proto_goTypes,
		DependencyIndexes: file_sigstore_trustroot_proto_depIdxs,
		MessageInfos:      file_sigstore_trustroot_proto_msgTypes,
	}.Build()
	File_sigstore_trustroot_proto = out.File
	file_sigstore_trustroot_proto_rawDesc = nil
	file_sigstore_trustroot_proto_goTypes = nil
	file_sigstore_trustroot_proto_depIdxs = nil
}
