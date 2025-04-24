# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [v1/common.proto](#v1_common-proto)
    - [Engine](#woxqaq-v1-Engine)
  
- [v1/generate_service.proto](#v1_generate_service-proto)
    - [AddRuleCatagoryRequest](#woxqaq-v1-AddRuleCatagoryRequest)
    - [AddRuleCatagoryRequest.RuleCatagoryMeta](#woxqaq-v1-AddRuleCatagoryRequest-RuleCatagoryMeta)
    - [AddRuleCatagoryResponse](#woxqaq-v1-AddRuleCatagoryResponse)
    - [AlgoValue](#woxqaq-v1-AlgoValue)
    - [CreateRuleRequest](#woxqaq-v1-CreateRuleRequest)
    - [CreateRuleResponse](#woxqaq-v1-CreateRuleResponse)
    - [DeleteRuleRequest](#woxqaq-v1-DeleteRuleRequest)
    - [DeleteRuleResponse](#woxqaq-v1-DeleteRuleResponse)
    - [GetRuleRequest](#woxqaq-v1-GetRuleRequest)
    - [GetRuleResponse](#woxqaq-v1-GetRuleResponse)
    - [ListRuleRequest](#woxqaq-v1-ListRuleRequest)
    - [ListRuleResponse](#woxqaq-v1-ListRuleResponse)
    - [PreviewRuleGenerateRequest](#woxqaq-v1-PreviewRuleGenerateRequest)
    - [PreviewRuleGenerateRequest.algoMeta](#woxqaq-v1-PreviewRuleGenerateRequest-algoMeta)
    - [PreviewRuleGenerateRequest.algoMeta.AlgoMetaEntry](#woxqaq-v1-PreviewRuleGenerateRequest-algoMeta-AlgoMetaEntry)
    - [PreviewRuleGenerateResponse](#woxqaq-v1-PreviewRuleGenerateResponse)
    - [Rule](#woxqaq-v1-Rule)
    - [Rule.AlgoMetaEntry](#woxqaq-v1-Rule-AlgoMetaEntry)
    - [StringList](#woxqaq-v1-StringList)
    - [UpdateRuleRequest](#woxqaq-v1-UpdateRuleRequest)
    - [UpdateRuleResponse](#woxqaq-v1-UpdateRuleResponse)
  
    - [GenerateAlgorithm](#woxqaq-v1-GenerateAlgorithm)
    - [RuleCategory](#woxqaq-v1-RuleCategory)
  
    - [RuleService](#woxqaq-v1-RuleService)
  
- [v1/issue_service.proto](#v1_issue_service-proto)
    - [ApprovalFlow](#woxqaq-v1-ApprovalFlow)
    - [ApprovalNode](#woxqaq-v1-ApprovalNode)
    - [ApprovalNodeNotifyConfig](#woxqaq-v1-ApprovalNodeNotifyConfig)
    - [ApprovalNodeNotifyConfig.NotifyTypesEntry](#woxqaq-v1-ApprovalNodeNotifyConfig-NotifyTypesEntry)
    - [ApprovalStep](#woxqaq-v1-ApprovalStep)
    - [ApprovalTemplate](#woxqaq-v1-ApprovalTemplate)
    - [ApproveRequest](#woxqaq-v1-ApproveRequest)
    - [ApproveRequest.approveMeta](#woxqaq-v1-ApproveRequest-approveMeta)
    - [ApproveResponse](#woxqaq-v1-ApproveResponse)
    - [DataExportDetail](#woxqaq-v1-DataExportDetail)
    - [DataExportDetail.ExportAccordingToTable](#woxqaq-v1-DataExportDetail-ExportAccordingToTable)
    - [DataExportDetail.ExportAccordingToTable.TableOption](#woxqaq-v1-DataExportDetail-ExportAccordingToTable-TableOption)
    - [DataExportDetail.ExportBySQL](#woxqaq-v1-DataExportDetail-ExportBySQL)
    - [GetIssueDetailRequest](#woxqaq-v1-GetIssueDetailRequest)
    - [GetIssueDetailResponse](#woxqaq-v1-GetIssueDetailResponse)
    - [Issue](#woxqaq-v1-Issue)
    - [Issue.Approver](#woxqaq-v1-Issue-Approver)
    - [IssueDetail](#woxqaq-v1-IssueDetail)
    - [ListApproveConfigRequest](#woxqaq-v1-ListApproveConfigRequest)
    - [ListApproveConfigResponse](#woxqaq-v1-ListApproveConfigResponse)
    - [ListIssueRequest](#woxqaq-v1-ListIssueRequest)
    - [ListIssueResponse](#woxqaq-v1-ListIssueResponse)
    - [RejectRequest](#woxqaq-v1-RejectRequest)
    - [RejectRequest.approveMeta](#woxqaq-v1-RejectRequest-approveMeta)
    - [RejectResponse](#woxqaq-v1-RejectResponse)
    - [SubmitIssueRequest](#woxqaq-v1-SubmitIssueRequest)
    - [SubmitIssueResponse](#woxqaq-v1-SubmitIssueResponse)
    - [UpdateApproveNotifyRequest](#woxqaq-v1-UpdateApproveNotifyRequest)
  
    - [ApprovalNode.Type](#woxqaq-v1-ApprovalNode-Type)
    - [ApprovalNodeNotifyConfig.NotifyType](#woxqaq-v1-ApprovalNodeNotifyConfig-NotifyType)
    - [ApprovalNodeNotifyConfig.Receiver](#woxqaq-v1-ApprovalNodeNotifyConfig-Receiver)
    - [ApprovalStep.Type](#woxqaq-v1-ApprovalStep-Type)
    - [Encoding](#woxqaq-v1-Encoding)
    - [Exec](#woxqaq-v1-Exec)
    - [ExportContent](#woxqaq-v1-ExportContent)
    - [ExportMethod](#woxqaq-v1-ExportMethod)
    - [ExportType](#woxqaq-v1-ExportType)
    - [Issue.Approver.Status](#woxqaq-v1-Issue-Approver-Status)
    - [IssueKind](#woxqaq-v1-IssueKind)
    - [IssueStatus](#woxqaq-v1-IssueStatus)
    - [ListIssueRequest.IssueRole](#woxqaq-v1-ListIssueRequest-IssueRole)
    - [TriggerKind](#woxqaq-v1-TriggerKind)
  
    - [IssueService](#woxqaq-v1-IssueService)
  
- [v1/mask_service.proto](#v1_mask_service-proto)
- [v1/notify_service.proto](#v1_notify_service-proto)
    - [NotifyRequest](#woxqaq-v1-NotifyRequest)
    - [NotifyResponse](#woxqaq-v1-NotifyResponse)
  
    - [NotifyServices](#woxqaq-v1-NotifyServices)
  
- [v1/review_service.proto](#v1_review_service-proto)
    - [DeleteReviewGroupRequest](#woxqaq-v1-DeleteReviewGroupRequest)
    - [DeleteReviewResponse](#woxqaq-v1-DeleteReviewResponse)
    - [ListReviewGroupsRequest](#woxqaq-v1-ListReviewGroupsRequest)
    - [ListReviewGroupsResponse](#woxqaq-v1-ListReviewGroupsResponse)
    - [ReviewGroup](#woxqaq-v1-ReviewGroup)
    - [ReviewRule](#woxqaq-v1-ReviewRule)
  
    - [ReviewRuleLevel](#woxqaq-v1-ReviewRuleLevel)
  
    - [ReviewServive](#woxqaq-v1-ReviewServive)
  
- [v1/upload_service.proto](#v1_upload_service-proto)
    - [PreSignRequest](#woxqaq-v1-PreSignRequest)
    - [PreSignResponse](#woxqaq-v1-PreSignResponse)
  
    - [UploadService](#woxqaq-v1-UploadService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="v1_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/common.proto


 


<a name="woxqaq-v1-Engine"></a>

### Engine


| Name | Number | Description |
| ---- | ------ | ----------- |
| ENGINE_UNSPECIFIED | 0 |  |
| CLICKHOUSE | 1 |  |
| MYSQL | 2 |  |
| POSTGRES | 3 |  |
| SNOWFLAKE | 4 |  |
| SQLITE | 5 |  |
| TIDB | 6 |  |
| MONGODB | 7 |  |
| REDIS | 8 |  |
| ORACLE | 9 |  |
| SPANNER | 10 |  |
| MSSQL | 11 |  |
| REDSHIFT | 12 |  |
| MARIADB | 13 |  |
| OCEANBASE | 14 |  |
| DM | 15 |  |
| RISINGWAVE | 16 |  |
| OCEANBASE_ORACLE | 17 |  |
| STARROCKS | 18 |  |
| DORIS | 19 |  |
| HIVE | 20 |  |
| ELASTICSEARCH | 21 |  |
| BIGQUERY | 22 |  |
| DYNAMODB | 23 |  |
| DATABRICKS | 24 |  |
| COCKROACHDB | 25 |  |
| COSMOSDB | 26 |  |
| TRINO | 27 |  |
| CASSANDRA | 28 |  |


 

 

 



<a name="v1_generate_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/generate_service.proto



<a name="woxqaq-v1-AddRuleCatagoryRequest"></a>

### AddRuleCatagoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| meta | [AddRuleCatagoryRequest.RuleCatagoryMeta](#woxqaq-v1-AddRuleCatagoryRequest-RuleCatagoryMeta) |  |  |






<a name="woxqaq-v1-AddRuleCatagoryRequest-RuleCatagoryMeta"></a>

### AddRuleCatagoryRequest.RuleCatagoryMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |






<a name="woxqaq-v1-AddRuleCatagoryResponse"></a>

### AddRuleCatagoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="woxqaq-v1-AlgoValue"></a>

### AlgoValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| single_string | [string](#string) |  |  |
| string_list | [StringList](#woxqaq-v1-StringList) |  |  |






<a name="woxqaq-v1-CreateRuleRequest"></a>

### CreateRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| rule | [Rule](#woxqaq-v1-Rule) |  |  |






<a name="woxqaq-v1-CreateRuleResponse"></a>

### CreateRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="woxqaq-v1-DeleteRuleRequest"></a>

### DeleteRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="woxqaq-v1-DeleteRuleResponse"></a>

### DeleteRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="woxqaq-v1-GetRuleRequest"></a>

### GetRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="woxqaq-v1-GetRuleResponse"></a>

### GetRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rule | [Rule](#woxqaq-v1-Rule) |  |  |






<a name="woxqaq-v1-ListRuleRequest"></a>

### ListRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="woxqaq-v1-ListRuleResponse"></a>

### ListRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rules | [Rule](#woxqaq-v1-Rule) | repeated |  |






<a name="woxqaq-v1-PreviewRuleGenerateRequest"></a>

### PreviewRuleGenerateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| meta | [PreviewRuleGenerateRequest.algoMeta](#woxqaq-v1-PreviewRuleGenerateRequest-algoMeta) |  |  |






<a name="woxqaq-v1-PreviewRuleGenerateRequest-algoMeta"></a>

### PreviewRuleGenerateRequest.algoMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| algo | [GenerateAlgorithm](#woxqaq-v1-GenerateAlgorithm) |  |  |
| algo_meta | [PreviewRuleGenerateRequest.algoMeta.AlgoMetaEntry](#woxqaq-v1-PreviewRuleGenerateRequest-algoMeta-AlgoMetaEntry) | repeated |  |






<a name="woxqaq-v1-PreviewRuleGenerateRequest-algoMeta-AlgoMetaEntry"></a>

### PreviewRuleGenerateRequest.algoMeta.AlgoMetaEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [AlgoValue](#woxqaq-v1-AlgoValue) |  |  |






<a name="woxqaq-v1-PreviewRuleGenerateResponse"></a>

### PreviewRuleGenerateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |
| value | [string](#string) |  |  |






<a name="woxqaq-v1-Rule"></a>

### Rule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| category | [RuleCategory](#woxqaq-v1-RuleCategory) |  |  |
| enable | [bool](#bool) |  |  |
| column_condition | [string](#string) |  |  |
| empty_ratio | [float](#float) |  |  |
| algorithm | [GenerateAlgorithm](#woxqaq-v1-GenerateAlgorithm) |  |  |
| algo_meta | [Rule.AlgoMetaEntry](#woxqaq-v1-Rule-AlgoMetaEntry) | repeated |  |
| internal | [bool](#bool) |  |  |






<a name="woxqaq-v1-Rule-AlgoMetaEntry"></a>

### Rule.AlgoMetaEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [AlgoValue](#woxqaq-v1-AlgoValue) |  |  |






<a name="woxqaq-v1-StringList"></a>

### StringList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [string](#string) | repeated |  |






<a name="woxqaq-v1-UpdateRuleRequest"></a>

### UpdateRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| rule | [Rule](#woxqaq-v1-Rule) |  |  |






<a name="woxqaq-v1-UpdateRuleResponse"></a>

### UpdateRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |





 


<a name="woxqaq-v1-GenerateAlgorithm"></a>

### GenerateAlgorithm


| Name | Number | Description |
| ---- | ------ | ----------- |
| GENERATE_ALGORITHM_UNSPECIFIED | 0 |  |
| GENERATE_ALGORITHM_TIME | 1 |  |
| GENERATE_ALGORITHM_TIME_DATE | 2 |  |
| GENERATE_ALGORITHM_DEPARTMENT | 3 |  |
| GENERATE_ALGORITHM_COMPANY | 4 |  |
| GENERATE_ALGORITHM_SCRIPT | 5 |  |



<a name="woxqaq-v1-RuleCategory"></a>

### RuleCategory


| Name | Number | Description |
| ---- | ------ | ----------- |
| RULE_CATEGORY_UNSPECIFIED | 0 |  |
| RULE_CATEGORY_PERSONAL | 1 |  |
| RULE_CATEGORY_PRODUCT | 2 |  |
| RULE_CATEGORY_COMMERCE | 3 |  |
| RULE_CATEGORY_COMMON | 4 |  |
| RULE_CATEGORY_POSITION | 5 |  |
| RULE_CATEGORY_CUSTOM | 6 |  |


 

 


<a name="woxqaq-v1-RuleService"></a>

### RuleService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetRule | [GetRuleRequest](#woxqaq-v1-GetRuleRequest) | [GetRuleResponse](#woxqaq-v1-GetRuleResponse) |  |
| ListRule | [ListRuleRequest](#woxqaq-v1-ListRuleRequest) | [ListRuleResponse](#woxqaq-v1-ListRuleResponse) |  |
| CreateRule | [CreateRuleRequest](#woxqaq-v1-CreateRuleRequest) | [CreateRuleResponse](#woxqaq-v1-CreateRuleResponse) |  |
| UpdateRule | [UpdateRuleRequest](#woxqaq-v1-UpdateRuleRequest) | [UpdateRuleResponse](#woxqaq-v1-UpdateRuleResponse) |  |
| DeleteRule | [DeleteRuleRequest](#woxqaq-v1-DeleteRuleRequest) | [DeleteRuleResponse](#woxqaq-v1-DeleteRuleResponse) |  |
| AddRuleCatagory | [AddRuleCatagoryRequest](#woxqaq-v1-AddRuleCatagoryRequest) | [AddRuleCatagoryResponse](#woxqaq-v1-AddRuleCatagoryResponse) |  |
| PreviewRuleGenerate | [PreviewRuleGenerateRequest](#woxqaq-v1-PreviewRuleGenerateRequest) | [PreviewRuleGenerateResponse](#woxqaq-v1-PreviewRuleGenerateResponse) |  |

 



<a name="v1_issue_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/issue_service.proto



<a name="woxqaq-v1-ApprovalFlow"></a>

### ApprovalFlow



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| steps | [ApprovalStep](#woxqaq-v1-ApprovalStep) | repeated |  |






<a name="woxqaq-v1-ApprovalNode"></a>

### ApprovalNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ApprovalNode.Type](#woxqaq-v1-ApprovalNode-Type) |  |  |
| role | [string](#string) |  |  |






<a name="woxqaq-v1-ApprovalNodeNotifyConfig"></a>

### ApprovalNodeNotifyConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notify_types | [ApprovalNodeNotifyConfig.NotifyTypesEntry](#woxqaq-v1-ApprovalNodeNotifyConfig-NotifyTypesEntry) | repeated |  |
| webhooks | [string](#string) | repeated |  |
| available_receiver | [ApprovalNodeNotifyConfig.Receiver](#woxqaq-v1-ApprovalNodeNotifyConfig-Receiver) | repeated |  |
| selected_receiver | [ApprovalNodeNotifyConfig.Receiver](#woxqaq-v1-ApprovalNodeNotifyConfig-Receiver) | repeated |  |
| custom_receivers | [string](#string) | repeated |  |






<a name="woxqaq-v1-ApprovalNodeNotifyConfig-NotifyTypesEntry"></a>

### ApprovalNodeNotifyConfig.NotifyTypesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [bool](#bool) |  |  |
| value | [ApprovalNodeNotifyConfig.NotifyType](#woxqaq-v1-ApprovalNodeNotifyConfig-NotifyType) |  |  |






<a name="woxqaq-v1-ApprovalStep"></a>

### ApprovalStep



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ApprovalStep.Type](#woxqaq-v1-ApprovalStep-Type) |  |  |
| nodes | [ApprovalNode](#woxqaq-v1-ApprovalNode) | repeated |  |






<a name="woxqaq-v1-ApprovalTemplate"></a>

### ApprovalTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flow | [ApprovalFlow](#woxqaq-v1-ApprovalFlow) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| required_approver | [string](#string) | repeated |  |






<a name="woxqaq-v1-ApproveRequest"></a>

### ApproveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| meta | [ApproveRequest.approveMeta](#woxqaq-v1-ApproveRequest-approveMeta) |  |  |






<a name="woxqaq-v1-ApproveRequest-approveMeta"></a>

### ApproveRequest.approveMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| approver | [string](#string) |  |  |
| comment | [string](#string) |  |  |






<a name="woxqaq-v1-ApproveResponse"></a>

### ApproveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="woxqaq-v1-DataExportDetail"></a>

### DataExportDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [string](#string) |  |  |
| ebs | [DataExportDetail.ExportBySQL](#woxqaq-v1-DataExportDetail-ExportBySQL) |  |  |
| eatt | [DataExportDetail.ExportAccordingToTable](#woxqaq-v1-DataExportDetail-ExportAccordingToTable) |  |  |
| type | [ExportType](#woxqaq-v1-ExportType) |  |  |
| encoding | [Encoding](#woxqaq-v1-Encoding) |  |  |
| export_reason | [string](#string) |  |  |
| exector | [Exec](#woxqaq-v1-Exec) |  |  |






<a name="woxqaq-v1-DataExportDetail-ExportAccordingToTable"></a>

### DataExportDetail.ExportAccordingToTable



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| export_all | [bool](#bool) |  |  |
| content | [ExportContent](#woxqaq-v1-ExportContent) |  |  |
| option | [DataExportDetail.ExportAccordingToTable.TableOption](#woxqaq-v1-DataExportDetail-ExportAccordingToTable-TableOption) | repeated |  |






<a name="woxqaq-v1-DataExportDetail-ExportAccordingToTable-TableOption"></a>

### DataExportDetail.ExportAccordingToTable.TableOption



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| table | [string](#string) |  |  |
| all_field | [bool](#bool) |  |  |
| fields | [string](#string) | repeated |  |
| filter | [string](#string) |  |  |






<a name="woxqaq-v1-DataExportDetail-ExportBySQL"></a>

### DataExportDetail.ExportBySQL



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| statement | [string](#string) |  |  |






<a name="woxqaq-v1-GetIssueDetailRequest"></a>

### GetIssueDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="woxqaq-v1-GetIssueDetailResponse"></a>

### GetIssueDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |
| ded | [DataExportDetail](#woxqaq-v1-DataExportDetail) |  |  |






<a name="woxqaq-v1-Issue"></a>

### Issue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| kind | [IssueKind](#woxqaq-v1-IssueKind) |  |  |
| database | [string](#string) |  |  |
| create_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| update_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| submitter | [string](#string) |  |  |
| status | [IssueStatus](#woxqaq-v1-IssueStatus) |  |  |
| detail | [IssueDetail](#woxqaq-v1-IssueDetail) |  |  |
| approvers | [Issue.Approver](#woxqaq-v1-Issue-Approver) | repeated |  |
| approval_templates | [ApprovalTemplate](#woxqaq-v1-ApprovalTemplate) | repeated |  |






<a name="woxqaq-v1-Issue-Approver"></a>

### Issue.Approver



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Issue.Approver.Status](#woxqaq-v1-Issue-Approver-Status) |  | The new status. |
| principal | [string](#string) |  |  |






<a name="woxqaq-v1-IssueDetail"></a>

### IssueDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ded | [DataExportDetail](#woxqaq-v1-DataExportDetail) |  |  |






<a name="woxqaq-v1-ListApproveConfigRequest"></a>

### ListApproveConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="woxqaq-v1-ListApproveConfigResponse"></a>

### ListApproveConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |
| data | [ApprovalNodeNotifyConfig](#woxqaq-v1-ApprovalNodeNotifyConfig) | repeated |  |






<a name="woxqaq-v1-ListIssueRequest"></a>

### ListIssueRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| kind | [IssueKind](#woxqaq-v1-IssueKind) |  |  |
| current | [int32](#int32) |  |  |
| page_size | [int32](#int32) |  |  |
| status | [IssueStatus](#woxqaq-v1-IssueStatus) |  |  |
| role | [ListIssueRequest.IssueRole](#woxqaq-v1-ListIssueRequest-IssueRole) |  |  |






<a name="woxqaq-v1-ListIssueResponse"></a>

### ListIssueResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |
| data | [Issue](#woxqaq-v1-Issue) | repeated |  |
| page_size | [int32](#int32) |  |  |
| total | [int32](#int32) |  |  |






<a name="woxqaq-v1-RejectRequest"></a>

### RejectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| meta | [RejectRequest.approveMeta](#woxqaq-v1-RejectRequest-approveMeta) |  |  |






<a name="woxqaq-v1-RejectRequest-approveMeta"></a>

### RejectRequest.approveMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| approver | [string](#string) |  |  |
| comment | [string](#string) |  |  |






<a name="woxqaq-v1-RejectResponse"></a>

### RejectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="woxqaq-v1-SubmitIssueRequest"></a>

### SubmitIssueRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| kind | [IssueKind](#woxqaq-v1-IssueKind) |  |  |
| detail | [IssueDetail](#woxqaq-v1-IssueDetail) |  |  |






<a name="woxqaq-v1-SubmitIssueResponse"></a>

### SubmitIssueResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |
| issue_id | [string](#string) |  |  |






<a name="woxqaq-v1-UpdateApproveNotifyRequest"></a>

### UpdateApproveNotifyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| data | [ApprovalNodeNotifyConfig](#woxqaq-v1-ApprovalNodeNotifyConfig) |  |  |





 


<a name="woxqaq-v1-ApprovalNode-Type"></a>

### ApprovalNode.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| ANY_IN_GROUP | 1 |  |



<a name="woxqaq-v1-ApprovalNodeNotifyConfig-NotifyType"></a>

### ApprovalNodeNotifyConfig.NotifyType


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| SITE_NOTIFY | 1 |  |
| SMS | 2 |  |
| EMAIL | 3 |  |
| PHONE | 4 |  |
| WEBHOOK | 5 |  |



<a name="woxqaq-v1-ApprovalNodeNotifyConfig-Receiver"></a>

### ApprovalNodeNotifyConfig.Receiver


| Name | Number | Description |
| ---- | ------ | ----------- |
| RECEIVER_UNSPECIFIED | 0 |  |
| APPROVER | 1 |  |
| COMMITTER | 2 |  |
| EXECUTOR | 3 |  |
| CUSTOM | 4 |  |



<a name="woxqaq-v1-ApprovalStep-Type"></a>

### ApprovalStep.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| ALL | 1 |  |
| ANY | 2 |  |



<a name="woxqaq-v1-Encoding"></a>

### Encoding


| Name | Number | Description |
| ---- | ------ | ----------- |
| ENCODING_UNSPECIFIED | 0 |  |
| UTF8 | 1 |  |
| GBK | 2 |  |



<a name="woxqaq-v1-Exec"></a>

### Exec


| Name | Number | Description |
| ---- | ------ | ----------- |
| EXEC_UNSPECIFIED | 0 |  |
| APPROVER | 1 |  |
| COMMITTER | 2 |  |
| AUTO | 3 |  |



<a name="woxqaq-v1-ExportContent"></a>

### ExportContent


| Name | Number | Description |
| ---- | ------ | ----------- |
| EXPORT_CONTENT_UNSPECIFIED | 0 |  |
| DATA | 1 |  |
| STRUCT | 2 |  |
| ALL | 3 |  |



<a name="woxqaq-v1-ExportMethod"></a>

### ExportMethod


| Name | Number | Description |
| ---- | ------ | ----------- |
| EXPORT_METHOD_UNSPECIFIED | 0 |  |
| SQL | 1 |  |
| TABLES | 2 |  |



<a name="woxqaq-v1-ExportType"></a>

### ExportType


| Name | Number | Description |
| ---- | ------ | ----------- |
| EXPORT_TYPE_UNSPECIFIED | 0 |  |
| CSV | 1 |  |
| EXCEL | 2 |  |



<a name="woxqaq-v1-Issue-Approver-Status"></a>

### Issue.Approver.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 |  |
| PENDING | 1 |  |
| APPROVED | 2 |  |
| REJECTED | 3 |  |



<a name="woxqaq-v1-IssueKind"></a>

### IssueKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| ISSUE_KIND_UNSPECIFIED | 0 |  |
| DATA_IMPORT | 1 |  |
| DATA_EXPORT | 2 |  |
| DATA_GENERATE | 3 |  |



<a name="woxqaq-v1-IssueStatus"></a>

### IssueStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| ISSUE_STATUS_UNSPECIFIED | 0 |  |
| NOT_START | 1 |  |
| PRE_CHECK | 2 |  |
| PRE_CHECK_SUCCESS | 3 |  |
| PRE_CHECK_FAILED | 4 |  |
| APPROVING | 5 |  |
| APPROVED | 6 |  |
| NOT_APPROVED | 7 |  |



<a name="woxqaq-v1-ListIssueRequest-IssueRole"></a>

### ListIssueRequest.IssueRole


| Name | Number | Description |
| ---- | ------ | ----------- |
| ISSUE_ROLE_UNSPECIFIED | 0 |  |
| COMMITTER | 1 |  |
| TO_APPROVE | 2 |  |
| TO_EXECUTE | 3 |  |
| APPROVED | 4 |  |
| EXECUTED | 5 |  |



<a name="woxqaq-v1-TriggerKind"></a>

### TriggerKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| TRIGGER_KIND_UNSPECIFIED | 0 |  |
| CURRENT | 1 |  |
| CRON | 2 |  |


 

 


<a name="woxqaq-v1-IssueService"></a>

### IssueService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetIssueDetail | [GetIssueDetailRequest](#woxqaq-v1-GetIssueDetailRequest) | [GetIssueDetailResponse](#woxqaq-v1-GetIssueDetailResponse) |  |
| ListIssue | [ListIssueRequest](#woxqaq-v1-ListIssueRequest) | [ListIssueResponse](#woxqaq-v1-ListIssueResponse) |  |
| SubmitIssue | [SubmitIssueRequest](#woxqaq-v1-SubmitIssueRequest) | [SubmitIssueResponse](#woxqaq-v1-SubmitIssueResponse) |  |
| Approve | [ApproveRequest](#woxqaq-v1-ApproveRequest) | [ApproveResponse](#woxqaq-v1-ApproveResponse) |  |
| Reject | [RejectRequest](#woxqaq-v1-RejectRequest) | [RejectResponse](#woxqaq-v1-RejectResponse) |  |
| ListApproveConfig | [ListApproveConfigRequest](#woxqaq-v1-ListApproveConfigRequest) | [ListApproveConfigResponse](#woxqaq-v1-ListApproveConfigResponse) |  |
| UpdateApproveRequestConfig | [UpdateApproveNotifyRequest](#woxqaq-v1-UpdateApproveNotifyRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |

 



<a name="v1_mask_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/mask_service.proto


 

 

 

 



<a name="v1_notify_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/notify_service.proto



<a name="woxqaq-v1-NotifyRequest"></a>

### NotifyRequest







<a name="woxqaq-v1-NotifyResponse"></a>

### NotifyResponse






 

 

 


<a name="woxqaq-v1-NotifyServices"></a>

### NotifyServices


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Notify | [NotifyRequest](#woxqaq-v1-NotifyRequest) | [NotifyResponse](#woxqaq-v1-NotifyResponse) |  |

 



<a name="v1_review_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/review_service.proto



<a name="woxqaq-v1-DeleteReviewGroupRequest"></a>

### DeleteReviewGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| org | [string](#string) |  |  |






<a name="woxqaq-v1-DeleteReviewResponse"></a>

### DeleteReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="woxqaq-v1-ListReviewGroupsRequest"></a>

### ListReviewGroupsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| org | [string](#string) |  |  |






<a name="woxqaq-v1-ListReviewGroupsResponse"></a>

### ListReviewGroupsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| review_groups | [ReviewGroup](#woxqaq-v1-ReviewGroup) | repeated |  |






<a name="woxqaq-v1-ReviewGroup"></a>

### ReviewGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| org | [string](#string) |  |  |
| title | [string](#string) |  |  |
| enabled | [bool](#bool) |  |  |
| rules | [ReviewRule](#woxqaq-v1-ReviewRule) | repeated |  |






<a name="woxqaq-v1-ReviewRule"></a>

### ReviewRule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| level | [ReviewRuleLevel](#woxqaq-v1-ReviewRuleLevel) |  |  |
| payload | [string](#string) |  |  |
| engine | [Engine](#woxqaq-v1-Engine) |  |  |
| comment | [string](#string) |  |  |





 


<a name="woxqaq-v1-ReviewRuleLevel"></a>

### ReviewRuleLevel


| Name | Number | Description |
| ---- | ------ | ----------- |
| LEVEL_UNSPECIFIED | 0 |  |
| ERROR | 1 |  |
| WARNING | 2 |  |
| DISABLED | 3 |  |


 

 


<a name="woxqaq-v1-ReviewServive"></a>

### ReviewServive


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListReviewGroups | [ListReviewGroupsRequest](#woxqaq-v1-ListReviewGroupsRequest) | [ListReviewGroupsResponse](#woxqaq-v1-ListReviewGroupsResponse) |  |
| CreateReviewGroup | [ReviewGroup](#woxqaq-v1-ReviewGroup) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| UpdateReviewGroup | [ReviewGroup](#woxqaq-v1-ReviewGroup) | [ReviewGroup](#woxqaq-v1-ReviewGroup) |  |
| DeleteReviewGroup | [DeleteReviewGroupRequest](#woxqaq-v1-DeleteReviewGroupRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |

 



<a name="v1_upload_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/upload_service.proto



<a name="woxqaq-v1-PreSignRequest"></a>

### PreSignRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_name | [string](#string) |  |  |
| instance_name | [string](#string) |  |  |






<a name="woxqaq-v1-PreSignResponse"></a>

### PreSignResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |





 

 

 


<a name="woxqaq-v1-UploadService"></a>

### UploadService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PreSign | [PreSignRequest](#woxqaq-v1-PreSignRequest) | [PreSignResponse](#woxqaq-v1-PreSignResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

