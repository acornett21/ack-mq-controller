// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package broker

import (
	"context"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/mq"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/mq-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.MQ{}
	_ = &svcapitypes.Broker{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeBrokerResponse
	resp, err = rm.sdkapi.DescribeBrokerWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeBroker", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AuthenticationStrategy != nil {
		ko.Spec.AuthenticationStrategy = resp.AuthenticationStrategy
	} else {
		ko.Spec.AuthenticationStrategy = nil
	}
	if resp.AutoMinorVersionUpgrade != nil {
		ko.Spec.AutoMinorVersionUpgrade = resp.AutoMinorVersionUpgrade
	} else {
		ko.Spec.AutoMinorVersionUpgrade = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.BrokerArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.BrokerArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.BrokerId != nil {
		ko.Status.BrokerID = resp.BrokerId
	} else {
		ko.Status.BrokerID = nil
	}
	if resp.BrokerInstances != nil {
		f4 := []*svcapitypes.BrokerInstance{}
		for _, f4iter := range resp.BrokerInstances {
			f4elem := &svcapitypes.BrokerInstance{}
			if f4iter.ConsoleURL != nil {
				f4elem.ConsoleURL = f4iter.ConsoleURL
			}
			if f4iter.Endpoints != nil {
				f4elemf1 := []*string{}
				for _, f4elemf1iter := range f4iter.Endpoints {
					var f4elemf1elem string
					f4elemf1elem = *f4elemf1iter
					f4elemf1 = append(f4elemf1, &f4elemf1elem)
				}
				f4elem.Endpoints = f4elemf1
			}
			if f4iter.IpAddress != nil {
				f4elem.IPAddress = f4iter.IpAddress
			}
			f4 = append(f4, f4elem)
		}
		ko.Status.BrokerInstances = f4
	} else {
		ko.Status.BrokerInstances = nil
	}
	if resp.BrokerState != nil {
		ko.Status.BrokerState = resp.BrokerState
	} else {
		ko.Status.BrokerState = nil
	}
	if resp.DeploymentMode != nil {
		ko.Spec.DeploymentMode = resp.DeploymentMode
	} else {
		ko.Spec.DeploymentMode = nil
	}
	if resp.EncryptionOptions != nil {
		f10 := &svcapitypes.EncryptionOptions{}
		if resp.EncryptionOptions.KmsKeyId != nil {
			f10.KMSKeyID = resp.EncryptionOptions.KmsKeyId
		}
		if resp.EncryptionOptions.UseAwsOwnedKey != nil {
			f10.UseAWSOwnedKey = resp.EncryptionOptions.UseAwsOwnedKey
		}
		ko.Spec.EncryptionOptions = f10
	} else {
		ko.Spec.EncryptionOptions = nil
	}
	if resp.EngineType != nil {
		ko.Spec.EngineType = resp.EngineType
	} else {
		ko.Spec.EngineType = nil
	}
	if resp.EngineVersion != nil {
		ko.Spec.EngineVersion = resp.EngineVersion
	} else {
		ko.Spec.EngineVersion = nil
	}
	if resp.HostInstanceType != nil {
		ko.Spec.HostInstanceType = resp.HostInstanceType
	} else {
		ko.Spec.HostInstanceType = nil
	}
	if resp.LdapServerMetadata != nil {
		f14 := &svcapitypes.LDAPServerMetadataInput{}
		if resp.LdapServerMetadata.Hosts != nil {
			f14f0 := []*string{}
			for _, f14f0iter := range resp.LdapServerMetadata.Hosts {
				var f14f0elem string
				f14f0elem = *f14f0iter
				f14f0 = append(f14f0, &f14f0elem)
			}
			f14.Hosts = f14f0
		}
		if resp.LdapServerMetadata.RoleBase != nil {
			f14.RoleBase = resp.LdapServerMetadata.RoleBase
		}
		if resp.LdapServerMetadata.RoleName != nil {
			f14.RoleName = resp.LdapServerMetadata.RoleName
		}
		if resp.LdapServerMetadata.RoleSearchMatching != nil {
			f14.RoleSearchMatching = resp.LdapServerMetadata.RoleSearchMatching
		}
		if resp.LdapServerMetadata.RoleSearchSubtree != nil {
			f14.RoleSearchSubtree = resp.LdapServerMetadata.RoleSearchSubtree
		}
		if resp.LdapServerMetadata.ServiceAccountUsername != nil {
			f14.ServiceAccountUsername = resp.LdapServerMetadata.ServiceAccountUsername
		}
		if resp.LdapServerMetadata.UserBase != nil {
			f14.UserBase = resp.LdapServerMetadata.UserBase
		}
		if resp.LdapServerMetadata.UserRoleName != nil {
			f14.UserRoleName = resp.LdapServerMetadata.UserRoleName
		}
		if resp.LdapServerMetadata.UserSearchMatching != nil {
			f14.UserSearchMatching = resp.LdapServerMetadata.UserSearchMatching
		}
		if resp.LdapServerMetadata.UserSearchSubtree != nil {
			f14.UserSearchSubtree = resp.LdapServerMetadata.UserSearchSubtree
		}
		ko.Spec.LDAPServerMetadata = f14
	} else {
		ko.Spec.LDAPServerMetadata = nil
	}
	if resp.Logs != nil {
		f15 := &svcapitypes.Logs{}
		if resp.Logs.Audit != nil {
			f15.Audit = resp.Logs.Audit
		}
		if resp.Logs.General != nil {
			f15.General = resp.Logs.General
		}
		ko.Spec.Logs = f15
	} else {
		ko.Spec.Logs = nil
	}
	if resp.MaintenanceWindowStartTime != nil {
		f16 := &svcapitypes.WeeklyStartTime{}
		if resp.MaintenanceWindowStartTime.DayOfWeek != nil {
			f16.DayOfWeek = resp.MaintenanceWindowStartTime.DayOfWeek
		}
		if resp.MaintenanceWindowStartTime.TimeOfDay != nil {
			f16.TimeOfDay = resp.MaintenanceWindowStartTime.TimeOfDay
		}
		if resp.MaintenanceWindowStartTime.TimeZone != nil {
			f16.TimeZone = resp.MaintenanceWindowStartTime.TimeZone
		}
		ko.Spec.MaintenanceWindowStartTime = f16
	} else {
		ko.Spec.MaintenanceWindowStartTime = nil
	}
	if resp.PubliclyAccessible != nil {
		ko.Spec.PubliclyAccessible = resp.PubliclyAccessible
	} else {
		ko.Spec.PubliclyAccessible = nil
	}
	if resp.SecurityGroups != nil {
		f23 := []*string{}
		for _, f23iter := range resp.SecurityGroups {
			var f23elem string
			f23elem = *f23iter
			f23 = append(f23, &f23elem)
		}
		ko.Spec.SecurityGroups = f23
	} else {
		ko.Spec.SecurityGroups = nil
	}
	if resp.StorageType != nil {
		ko.Spec.StorageType = resp.StorageType
	} else {
		ko.Spec.StorageType = nil
	}
	if resp.SubnetIds != nil {
		f25 := []*string{}
		for _, f25iter := range resp.SubnetIds {
			var f25elem string
			f25elem = *f25iter
			f25 = append(f25, &f25elem)
		}
		ko.Spec.SubnetIDs = f25
	} else {
		ko.Spec.SubnetIDs = nil
	}
	if resp.Tags != nil {
		f26 := map[string]*string{}
		for f26key, f26valiter := range resp.Tags {
			var f26val string
			f26val = *f26valiter
			f26[f26key] = &f26val
		}
		ko.Spec.Tags = f26
	} else {
		ko.Spec.Tags = nil
	}
	if resp.Users != nil {
		f27 := []*svcapitypes.User{}
		for _, f27iter := range resp.Users {
			f27elem := &svcapitypes.User{}
			if f27iter.Username != nil {
				f27elem.Username = f27iter.Username
			}
			f27 = append(f27, f27elem)
		}
		ko.Spec.Users = f27
	} else {
		ko.Spec.Users = nil
	}

	rm.setStatusDefaults(ko)
	if brokerCreateInProgress(&resource{ko}) {
		return &resource{ko}, requeueWaitWhileCreating
	}
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Status.BrokerID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeBrokerInput, error) {
	res := &svcsdk.DescribeBrokerInput{}

	if r.ko.Status.BrokerID != nil {
		res.SetBrokerId(*r.ko.Status.BrokerID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateBrokerResponse
	_ = resp
	resp, err = rm.sdkapi.CreateBrokerWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateBroker", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.BrokerArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.BrokerArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.BrokerId != nil {
		ko.Status.BrokerID = resp.BrokerId
	} else {
		ko.Status.BrokerID = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateBrokerRequest, error) {
	res := &svcsdk.CreateBrokerRequest{}

	if r.ko.Spec.AuthenticationStrategy != nil {
		res.SetAuthenticationStrategy(*r.ko.Spec.AuthenticationStrategy)
	}
	if r.ko.Spec.AutoMinorVersionUpgrade != nil {
		res.SetAutoMinorVersionUpgrade(*r.ko.Spec.AutoMinorVersionUpgrade)
	}
	if r.ko.Spec.Name != nil {
		res.SetBrokerName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.Configuration != nil {
		f3 := &svcsdk.ConfigurationId{}
		if r.ko.Spec.Configuration.ID != nil {
			f3.SetId(*r.ko.Spec.Configuration.ID)
		}
		if r.ko.Spec.Configuration.Revision != nil {
			f3.SetRevision(*r.ko.Spec.Configuration.Revision)
		}
		res.SetConfiguration(f3)
	}
	if r.ko.Spec.CreatorRequestID != nil {
		res.SetCreatorRequestId(*r.ko.Spec.CreatorRequestID)
	}
	if r.ko.Spec.DeploymentMode != nil {
		res.SetDeploymentMode(*r.ko.Spec.DeploymentMode)
	}
	if r.ko.Spec.EncryptionOptions != nil {
		f6 := &svcsdk.EncryptionOptions{}
		if r.ko.Spec.EncryptionOptions.KMSKeyID != nil {
			f6.SetKmsKeyId(*r.ko.Spec.EncryptionOptions.KMSKeyID)
		}
		if r.ko.Spec.EncryptionOptions.UseAWSOwnedKey != nil {
			f6.SetUseAwsOwnedKey(*r.ko.Spec.EncryptionOptions.UseAWSOwnedKey)
		}
		res.SetEncryptionOptions(f6)
	}
	if r.ko.Spec.EngineType != nil {
		res.SetEngineType(*r.ko.Spec.EngineType)
	}
	if r.ko.Spec.EngineVersion != nil {
		res.SetEngineVersion(*r.ko.Spec.EngineVersion)
	}
	if r.ko.Spec.HostInstanceType != nil {
		res.SetHostInstanceType(*r.ko.Spec.HostInstanceType)
	}
	if r.ko.Spec.LDAPServerMetadata != nil {
		f10 := &svcsdk.LdapServerMetadataInput{}
		if r.ko.Spec.LDAPServerMetadata.Hosts != nil {
			f10f0 := []*string{}
			for _, f10f0iter := range r.ko.Spec.LDAPServerMetadata.Hosts {
				var f10f0elem string
				f10f0elem = *f10f0iter
				f10f0 = append(f10f0, &f10f0elem)
			}
			f10.SetHosts(f10f0)
		}
		if r.ko.Spec.LDAPServerMetadata.RoleBase != nil {
			f10.SetRoleBase(*r.ko.Spec.LDAPServerMetadata.RoleBase)
		}
		if r.ko.Spec.LDAPServerMetadata.RoleName != nil {
			f10.SetRoleName(*r.ko.Spec.LDAPServerMetadata.RoleName)
		}
		if r.ko.Spec.LDAPServerMetadata.RoleSearchMatching != nil {
			f10.SetRoleSearchMatching(*r.ko.Spec.LDAPServerMetadata.RoleSearchMatching)
		}
		if r.ko.Spec.LDAPServerMetadata.RoleSearchSubtree != nil {
			f10.SetRoleSearchSubtree(*r.ko.Spec.LDAPServerMetadata.RoleSearchSubtree)
		}
		if r.ko.Spec.LDAPServerMetadata.ServiceAccountPassword != nil {
			f10.SetServiceAccountPassword(*r.ko.Spec.LDAPServerMetadata.ServiceAccountPassword)
		}
		if r.ko.Spec.LDAPServerMetadata.ServiceAccountUsername != nil {
			f10.SetServiceAccountUsername(*r.ko.Spec.LDAPServerMetadata.ServiceAccountUsername)
		}
		if r.ko.Spec.LDAPServerMetadata.UserBase != nil {
			f10.SetUserBase(*r.ko.Spec.LDAPServerMetadata.UserBase)
		}
		if r.ko.Spec.LDAPServerMetadata.UserRoleName != nil {
			f10.SetUserRoleName(*r.ko.Spec.LDAPServerMetadata.UserRoleName)
		}
		if r.ko.Spec.LDAPServerMetadata.UserSearchMatching != nil {
			f10.SetUserSearchMatching(*r.ko.Spec.LDAPServerMetadata.UserSearchMatching)
		}
		if r.ko.Spec.LDAPServerMetadata.UserSearchSubtree != nil {
			f10.SetUserSearchSubtree(*r.ko.Spec.LDAPServerMetadata.UserSearchSubtree)
		}
		res.SetLdapServerMetadata(f10)
	}
	if r.ko.Spec.Logs != nil {
		f11 := &svcsdk.Logs{}
		if r.ko.Spec.Logs.Audit != nil {
			f11.SetAudit(*r.ko.Spec.Logs.Audit)
		}
		if r.ko.Spec.Logs.General != nil {
			f11.SetGeneral(*r.ko.Spec.Logs.General)
		}
		res.SetLogs(f11)
	}
	if r.ko.Spec.MaintenanceWindowStartTime != nil {
		f12 := &svcsdk.WeeklyStartTime{}
		if r.ko.Spec.MaintenanceWindowStartTime.DayOfWeek != nil {
			f12.SetDayOfWeek(*r.ko.Spec.MaintenanceWindowStartTime.DayOfWeek)
		}
		if r.ko.Spec.MaintenanceWindowStartTime.TimeOfDay != nil {
			f12.SetTimeOfDay(*r.ko.Spec.MaintenanceWindowStartTime.TimeOfDay)
		}
		if r.ko.Spec.MaintenanceWindowStartTime.TimeZone != nil {
			f12.SetTimeZone(*r.ko.Spec.MaintenanceWindowStartTime.TimeZone)
		}
		res.SetMaintenanceWindowStartTime(f12)
	}
	if r.ko.Spec.PubliclyAccessible != nil {
		res.SetPubliclyAccessible(*r.ko.Spec.PubliclyAccessible)
	}
	if r.ko.Spec.SecurityGroups != nil {
		f14 := []*string{}
		for _, f14iter := range r.ko.Spec.SecurityGroups {
			var f14elem string
			f14elem = *f14iter
			f14 = append(f14, &f14elem)
		}
		res.SetSecurityGroups(f14)
	}
	if r.ko.Spec.StorageType != nil {
		res.SetStorageType(*r.ko.Spec.StorageType)
	}
	if r.ko.Spec.SubnetIDs != nil {
		f16 := []*string{}
		for _, f16iter := range r.ko.Spec.SubnetIDs {
			var f16elem string
			f16elem = *f16iter
			f16 = append(f16, &f16elem)
		}
		res.SetSubnetIds(f16)
	}
	if r.ko.Spec.Tags != nil {
		f17 := map[string]*string{}
		for f17key, f17valiter := range r.ko.Spec.Tags {
			var f17val string
			f17val = *f17valiter
			f17[f17key] = &f17val
		}
		res.SetTags(f17)
	}
	if r.ko.Spec.Users != nil {
		f18 := []*svcsdk.User{}
		for _, f18iter := range r.ko.Spec.Users {
			f18elem := &svcsdk.User{}
			if f18iter.ConsoleAccess != nil {
				f18elem.SetConsoleAccess(*f18iter.ConsoleAccess)
			}
			if f18iter.Groups != nil {
				f18elemf1 := []*string{}
				for _, f18elemf1iter := range f18iter.Groups {
					var f18elemf1elem string
					f18elemf1elem = *f18elemf1iter
					f18elemf1 = append(f18elemf1, &f18elemf1elem)
				}
				f18elem.SetGroups(f18elemf1)
			}
			if f18iter.Password != nil {
				tmpSecret, err := rm.rr.SecretValueFromReference(ctx, f18iter.Password)
				if err != nil {
					return nil, err
				}
				if tmpSecret != "" {
					f18elem.SetPassword(tmpSecret)
				}
			}
			if f18iter.Username != nil {
				f18elem.SetUsername(*f18iter.Username)
			}
			f18 = append(f18, f18elem)
		}
		res.SetUsers(f18)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer exit(err)
	if brokerCreateFailed(latest) {
		msg := "Broker state is CREATION_FAILED"
		setTerminalCondition(desired, corev1.ConditionTrue, &msg, nil)
		setSyncedCondition(desired, corev1.ConditionTrue, nil, nil)
		return desired, nil
	}
	if brokerCreateInProgress(latest) {
		msg := "Broker state is CREATION_IN_PROGRESS"
		setSyncedCondition(desired, corev1.ConditionFalse, &msg, nil)
		return desired, requeueWaitWhileCreating
	}
	if brokerDeleteInProgress(latest) {
		msg := "Broker state is DELETION_IN_PROGRESS"
		setSyncedCondition(desired, corev1.ConditionFalse, &msg, nil)
		return desired, requeueWaitWhileDeleting
	}

	input, err := rm.newUpdateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateBrokerResponse
	_ = resp
	resp, err = rm.sdkapi.UpdateBrokerWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateBroker", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()
	// Copy status from latest observed state
	latestKOStatus := latest.ko.DeepCopy().Status
	ko.Status = latestKOStatus

	if resp.AuthenticationStrategy != nil {
		ko.Spec.AuthenticationStrategy = resp.AuthenticationStrategy
	} else {
		ko.Spec.AuthenticationStrategy = nil
	}
	if resp.AutoMinorVersionUpgrade != nil {
		ko.Spec.AutoMinorVersionUpgrade = resp.AutoMinorVersionUpgrade
	} else {
		ko.Spec.AutoMinorVersionUpgrade = nil
	}
	if resp.BrokerId != nil {
		ko.Status.BrokerID = resp.BrokerId
	} else {
		ko.Status.BrokerID = nil
	}
	if resp.Configuration != nil {
		f3 := &svcapitypes.ConfigurationID{}
		if resp.Configuration.Id != nil {
			f3.ID = resp.Configuration.Id
		}
		if resp.Configuration.Revision != nil {
			f3.Revision = resp.Configuration.Revision
		}
		ko.Spec.Configuration = f3
	} else {
		ko.Spec.Configuration = nil
	}
	if resp.EngineVersion != nil {
		ko.Spec.EngineVersion = resp.EngineVersion
	} else {
		ko.Spec.EngineVersion = nil
	}
	if resp.HostInstanceType != nil {
		ko.Spec.HostInstanceType = resp.HostInstanceType
	} else {
		ko.Spec.HostInstanceType = nil
	}
	if resp.LdapServerMetadata != nil {
		f6 := &svcapitypes.LDAPServerMetadataInput{}
		if resp.LdapServerMetadata.Hosts != nil {
			f6f0 := []*string{}
			for _, f6f0iter := range resp.LdapServerMetadata.Hosts {
				var f6f0elem string
				f6f0elem = *f6f0iter
				f6f0 = append(f6f0, &f6f0elem)
			}
			f6.Hosts = f6f0
		}
		if resp.LdapServerMetadata.RoleBase != nil {
			f6.RoleBase = resp.LdapServerMetadata.RoleBase
		}
		if resp.LdapServerMetadata.RoleName != nil {
			f6.RoleName = resp.LdapServerMetadata.RoleName
		}
		if resp.LdapServerMetadata.RoleSearchMatching != nil {
			f6.RoleSearchMatching = resp.LdapServerMetadata.RoleSearchMatching
		}
		if resp.LdapServerMetadata.RoleSearchSubtree != nil {
			f6.RoleSearchSubtree = resp.LdapServerMetadata.RoleSearchSubtree
		}
		if resp.LdapServerMetadata.ServiceAccountUsername != nil {
			f6.ServiceAccountUsername = resp.LdapServerMetadata.ServiceAccountUsername
		}
		if resp.LdapServerMetadata.UserBase != nil {
			f6.UserBase = resp.LdapServerMetadata.UserBase
		}
		if resp.LdapServerMetadata.UserRoleName != nil {
			f6.UserRoleName = resp.LdapServerMetadata.UserRoleName
		}
		if resp.LdapServerMetadata.UserSearchMatching != nil {
			f6.UserSearchMatching = resp.LdapServerMetadata.UserSearchMatching
		}
		if resp.LdapServerMetadata.UserSearchSubtree != nil {
			f6.UserSearchSubtree = resp.LdapServerMetadata.UserSearchSubtree
		}
		ko.Spec.LDAPServerMetadata = f6
	} else {
		ko.Spec.LDAPServerMetadata = nil
	}
	if resp.Logs != nil {
		f7 := &svcapitypes.Logs{}
		if resp.Logs.Audit != nil {
			f7.Audit = resp.Logs.Audit
		}
		if resp.Logs.General != nil {
			f7.General = resp.Logs.General
		}
		ko.Spec.Logs = f7
	} else {
		ko.Spec.Logs = nil
	}
	if resp.SecurityGroups != nil {
		f8 := []*string{}
		for _, f8iter := range resp.SecurityGroups {
			var f8elem string
			f8elem = *f8iter
			f8 = append(f8, &f8elem)
		}
		ko.Spec.SecurityGroups = f8
	} else {
		ko.Spec.SecurityGroups = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.UpdateBrokerRequest, error) {
	res := &svcsdk.UpdateBrokerRequest{}

	if r.ko.Spec.AuthenticationStrategy != nil {
		res.SetAuthenticationStrategy(*r.ko.Spec.AuthenticationStrategy)
	}
	if r.ko.Spec.AutoMinorVersionUpgrade != nil {
		res.SetAutoMinorVersionUpgrade(*r.ko.Spec.AutoMinorVersionUpgrade)
	}
	if r.ko.Status.BrokerID != nil {
		res.SetBrokerId(*r.ko.Status.BrokerID)
	}
	if r.ko.Spec.Configuration != nil {
		f3 := &svcsdk.ConfigurationId{}
		if r.ko.Spec.Configuration.ID != nil {
			f3.SetId(*r.ko.Spec.Configuration.ID)
		}
		if r.ko.Spec.Configuration.Revision != nil {
			f3.SetRevision(*r.ko.Spec.Configuration.Revision)
		}
		res.SetConfiguration(f3)
	}
	if r.ko.Spec.EngineVersion != nil {
		res.SetEngineVersion(*r.ko.Spec.EngineVersion)
	}
	if r.ko.Spec.HostInstanceType != nil {
		res.SetHostInstanceType(*r.ko.Spec.HostInstanceType)
	}
	if r.ko.Spec.LDAPServerMetadata != nil {
		f6 := &svcsdk.LdapServerMetadataInput{}
		if r.ko.Spec.LDAPServerMetadata.Hosts != nil {
			f6f0 := []*string{}
			for _, f6f0iter := range r.ko.Spec.LDAPServerMetadata.Hosts {
				var f6f0elem string
				f6f0elem = *f6f0iter
				f6f0 = append(f6f0, &f6f0elem)
			}
			f6.SetHosts(f6f0)
		}
		if r.ko.Spec.LDAPServerMetadata.RoleBase != nil {
			f6.SetRoleBase(*r.ko.Spec.LDAPServerMetadata.RoleBase)
		}
		if r.ko.Spec.LDAPServerMetadata.RoleName != nil {
			f6.SetRoleName(*r.ko.Spec.LDAPServerMetadata.RoleName)
		}
		if r.ko.Spec.LDAPServerMetadata.RoleSearchMatching != nil {
			f6.SetRoleSearchMatching(*r.ko.Spec.LDAPServerMetadata.RoleSearchMatching)
		}
		if r.ko.Spec.LDAPServerMetadata.RoleSearchSubtree != nil {
			f6.SetRoleSearchSubtree(*r.ko.Spec.LDAPServerMetadata.RoleSearchSubtree)
		}
		if r.ko.Spec.LDAPServerMetadata.ServiceAccountPassword != nil {
			f6.SetServiceAccountPassword(*r.ko.Spec.LDAPServerMetadata.ServiceAccountPassword)
		}
		if r.ko.Spec.LDAPServerMetadata.ServiceAccountUsername != nil {
			f6.SetServiceAccountUsername(*r.ko.Spec.LDAPServerMetadata.ServiceAccountUsername)
		}
		if r.ko.Spec.LDAPServerMetadata.UserBase != nil {
			f6.SetUserBase(*r.ko.Spec.LDAPServerMetadata.UserBase)
		}
		if r.ko.Spec.LDAPServerMetadata.UserRoleName != nil {
			f6.SetUserRoleName(*r.ko.Spec.LDAPServerMetadata.UserRoleName)
		}
		if r.ko.Spec.LDAPServerMetadata.UserSearchMatching != nil {
			f6.SetUserSearchMatching(*r.ko.Spec.LDAPServerMetadata.UserSearchMatching)
		}
		if r.ko.Spec.LDAPServerMetadata.UserSearchSubtree != nil {
			f6.SetUserSearchSubtree(*r.ko.Spec.LDAPServerMetadata.UserSearchSubtree)
		}
		res.SetLdapServerMetadata(f6)
	}
	if r.ko.Spec.Logs != nil {
		f7 := &svcsdk.Logs{}
		if r.ko.Spec.Logs.Audit != nil {
			f7.SetAudit(*r.ko.Spec.Logs.Audit)
		}
		if r.ko.Spec.Logs.General != nil {
			f7.SetGeneral(*r.ko.Spec.Logs.General)
		}
		res.SetLogs(f7)
	}
	if r.ko.Spec.SecurityGroups != nil {
		f8 := []*string{}
		for _, f8iter := range r.ko.Spec.SecurityGroups {
			var f8elem string
			f8elem = *f8iter
			f8 = append(f8, &f8elem)
		}
		res.SetSecurityGroups(f8)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	if brokerDeleteInProgress(r) {
		return r, requeueWaitWhileDeleting
	}

	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteBrokerResponse
	_ = resp
	resp, err = rm.sdkapi.DeleteBrokerWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteBroker", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteBrokerInput, error) {
	res := &svcsdk.DeleteBrokerInput{}

	if r.ko.Status.BrokerID != nil {
		res.SetBrokerId(*r.ko.Status.BrokerID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Broker,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
