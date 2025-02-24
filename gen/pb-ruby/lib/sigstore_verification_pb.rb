# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: sigstore_verification.proto

require 'google/protobuf'

require 'sigstore_common_pb'
require 'sigstore_trustroot_pb'
require 'sigstore_bundle_pb'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("sigstore_verification.proto", :syntax => :proto3) do
    add_message "dev.sigstore.verification.v1.CertificateIdentity" do
      optional :issuer, :string, 1
      optional :san, :message, 2, "dev.sigstore.common.v1.SubjectAlternativeName"
      repeated :oids, :message, 3, "dev.sigstore.common.v1.ObjectIdentifierValuePair"
    end
    add_message "dev.sigstore.verification.v1.CertificateIdentities" do
      repeated :identities, :message, 1, "dev.sigstore.verification.v1.CertificateIdentity"
    end
    add_message "dev.sigstore.verification.v1.PublicKeyIdentities" do
      repeated :public_keys, :message, 1, "dev.sigstore.common.v1.PublicKey"
    end
    add_message "dev.sigstore.verification.v1.ArtifactVerificationOptions" do
      proto3_optional :tlog_options, :message, 3, "dev.sigstore.verification.v1.ArtifactVerificationOptions.TlogOptions"
      proto3_optional :ctlog_options, :message, 4, "dev.sigstore.verification.v1.ArtifactVerificationOptions.CtlogOptions"
      proto3_optional :tsa_options, :message, 5, "dev.sigstore.verification.v1.ArtifactVerificationOptions.TimestampAuthorityOptions"
      oneof :signers do
        optional :certificate_identities, :message, 1, "dev.sigstore.verification.v1.CertificateIdentities"
        optional :public_keys, :message, 2, "dev.sigstore.verification.v1.PublicKeyIdentities"
      end
    end
    add_message "dev.sigstore.verification.v1.ArtifactVerificationOptions.TlogOptions" do
      optional :threshold, :int32, 1
      optional :perform_online_verification, :bool, 2
      optional :disable, :bool, 3
    end
    add_message "dev.sigstore.verification.v1.ArtifactVerificationOptions.CtlogOptions" do
      optional :threshold, :int32, 1
      optional :disable, :bool, 3
    end
    add_message "dev.sigstore.verification.v1.ArtifactVerificationOptions.TimestampAuthorityOptions" do
      optional :threshold, :int32, 1
      optional :disable, :bool, 2
    end
    add_message "dev.sigstore.verification.v1.Artifact" do
      oneof :data do
        optional :artifact_uri, :string, 1
        optional :artifact, :bytes, 2
      end
    end
    add_message "dev.sigstore.verification.v1.Input" do
      optional :artifact_trust_root, :message, 1, "dev.sigstore.trustroot.v1.TrustedRoot"
      optional :artifact_verification_options, :message, 2, "dev.sigstore.verification.v1.ArtifactVerificationOptions"
      optional :bundle, :message, 3, "dev.sigstore.bundle.v1.Bundle"
      proto3_optional :artifact, :message, 4, "dev.sigstore.verification.v1.Artifact"
    end
  end
end

module Sigstore
  module Verification
    module V1
      CertificateIdentity = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("dev.sigstore.verification.v1.CertificateIdentity").msgclass
      CertificateIdentities = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("dev.sigstore.verification.v1.CertificateIdentities").msgclass
      PublicKeyIdentities = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("dev.sigstore.verification.v1.PublicKeyIdentities").msgclass
      ArtifactVerificationOptions = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("dev.sigstore.verification.v1.ArtifactVerificationOptions").msgclass
      ArtifactVerificationOptions::TlogOptions = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("dev.sigstore.verification.v1.ArtifactVerificationOptions.TlogOptions").msgclass
      ArtifactVerificationOptions::CtlogOptions = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("dev.sigstore.verification.v1.ArtifactVerificationOptions.CtlogOptions").msgclass
      ArtifactVerificationOptions::TimestampAuthorityOptions = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("dev.sigstore.verification.v1.ArtifactVerificationOptions.TimestampAuthorityOptions").msgclass
      Artifact = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("dev.sigstore.verification.v1.Artifact").msgclass
      Input = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("dev.sigstore.verification.v1.Input").msgclass
    end
  end
end
