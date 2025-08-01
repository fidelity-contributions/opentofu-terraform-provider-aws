// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksight

// Exports for use in tests only.
var (
	ResourceAccountSettings     = newAccountSettingsResource
	ResourceAccountSubscription = resourceAccountSubscription
	ResourceAnalysis            = resourceAnalysis
	ResourceDashboard           = resourceDashboard
	ResourceDataSet             = resourceDataSet
	ResourceDataSource          = resourceDataSource
	ResourceFolder              = resourceFolder
	ResourceFolderMembership    = newFolderMembershipResource
	ResourceGroup               = resourceGroup
	ResourceGroupMembership     = resourceGroupMembership
	ResourceIAMPolicyAssignment = newIAMPolicyAssignmentResource
	ResourceIngestion           = newIngestionResource
	ResourceIPRestriction       = newIPRestrictionResource
	ResourceKeyRegistration     = newKeyRegistrationResource
	ResourceNamespace           = newNamespaceResource
	ResourceRefreshSchedule     = newRefreshScheduleResource
	ResourceRoleMembership      = newRoleMembershipResource
	ResourceTemplate            = resourceTemplate
	ResourceTemplateAlias       = newTemplateAliasResource
	ResourceTheme               = resourceTheme
	ResourceUser                = resourceUser
	ResourceVPCConnection       = newVPCConnectionResource

	DashboardLatestVersion                = dashboardLatestVersion
	DefaultGroupNamespace                 = defaultGroupNamespace
	DefaultUserNamespace                  = defaultUserNamespace
	FindAccountSettingsByID               = findAccountSettingsByID
	FindAccountSubscriptionByID           = findAccountSubscriptionByID
	FindAnalysisByTwoPartKey              = findAnalysisByTwoPartKey
	FindDashboardByThreePartKey           = findDashboardByThreePartKey
	FindDataSetByTwoPartKey               = findDataSetByTwoPartKey
	FindDataSourceByTwoPartKey            = findDataSourceByTwoPartKey
	FindFolderByTwoPartKey                = findFolderByTwoPartKey
	FindFolderMembershipByFourPartKey     = findFolderMembershipByFourPartKey
	FindGroupByThreePartKey               = findGroupByThreePartKey
	FindGroupMembershipByFourPartKey      = findGroupMembershipByFourPartKey
	FindIAMPolicyAssignmentByThreePartKey = findIAMPolicyAssignmentByThreePartKey
	FindIngestionByThreePartKey           = findIngestionByThreePartKey
	FindIPRestrictionByID                 = findIPRestrictionByID
	FindKeyRegistrationByID               = findKeyRegistrationByID
	FindNamespaceByTwoPartKey             = findNamespaceByTwoPartKey
	FindRefreshScheduleByThreePartKey     = findRefreshScheduleByThreePartKey
	FindRoleMembershipByMultiPartKey      = findRoleMembershipByMultiPartKey
	FindTemplateAliasByThreePartKey       = findTemplateAliasByThreePartKey
	FindTemplateByTwoPartKey              = findTemplateByTwoPartKey
	FindThemeByTwoPartKey                 = findThemeByTwoPartKey
	FindUserByThreePartKey                = findUserByThreePartKey
	FindVPCConnectionByTwoPartKey         = findVPCConnectionByTwoPartKey

	StartAfterDateTimeLayout = startAfterDateTimeLayout
)
