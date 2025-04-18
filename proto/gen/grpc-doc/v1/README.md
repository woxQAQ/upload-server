# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [v1/task_service.proto](#v1_task_service-proto)
    - [ApproveRequest](#woxqaq-v1-ApproveRequest)
    - [ApproveResponse](#woxqaq-v1-ApproveResponse)
    - [DataExportDetail](#woxqaq-v1-DataExportDetail)
    - [ExportAccordingToTable](#woxqaq-v1-ExportAccordingToTable)
    - [ExportAccordingToTable.TableOption](#woxqaq-v1-ExportAccordingToTable-TableOption)
    - [ExportBySQL](#woxqaq-v1-ExportBySQL)
    - [GetTaskDetailRequest](#woxqaq-v1-GetTaskDetailRequest)
    - [GetTaskDetailResponse](#woxqaq-v1-GetTaskDetailResponse)
    - [ListTaskRequest](#woxqaq-v1-ListTaskRequest)
    - [ListTaskResponse](#woxqaq-v1-ListTaskResponse)
    - [RejectRequest](#woxqaq-v1-RejectRequest)
    - [RejectResponse](#woxqaq-v1-RejectResponse)
    - [SubmitTaskRequest](#woxqaq-v1-SubmitTaskRequest)
    - [SubmitTaskResponse](#woxqaq-v1-SubmitTaskResponse)
    - [TaskListData](#woxqaq-v1-TaskListData)
  
    - [Encoding](#woxqaq-v1-Encoding)
    - [Exec](#woxqaq-v1-Exec)
    - [ExportContent](#woxqaq-v1-ExportContent)
    - [ExportMethod](#woxqaq-v1-ExportMethod)
    - [ExportType](#woxqaq-v1-ExportType)
    - [TaskKind](#woxqaq-v1-TaskKind)
    - [TaskRole](#woxqaq-v1-TaskRole)
    - [TaskStatus](#woxqaq-v1-TaskStatus)
    - [TriggerKind](#woxqaq-v1-TriggerKind)
  
    - [TaskService](#woxqaq-v1-TaskService)
  
- [v1/upload_service.proto](#v1_upload_service-proto)
    - [PreSignRequest](#woxqaq-v1-PreSignRequest)
    - [PreSignResponse](#woxqaq-v1-PreSignResponse)
  
    - [UploadService](#woxqaq-v1-UploadService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="v1_task_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/task_service.proto



<a name="woxqaq-v1-ApproveRequest"></a>

### ApproveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [uint64](#uint64) |  |  |
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
| ebs | [ExportBySQL](#woxqaq-v1-ExportBySQL) |  |  |
| eatt | [ExportAccordingToTable](#woxqaq-v1-ExportAccordingToTable) |  |  |
| type | [ExportType](#woxqaq-v1-ExportType) |  |  |
| encoding | [Encoding](#woxqaq-v1-Encoding) |  |  |
| export_reason | [string](#string) |  |  |
| exector | [Exec](#woxqaq-v1-Exec) |  |  |






<a name="woxqaq-v1-ExportAccordingToTable"></a>

### ExportAccordingToTable



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| export_all | [bool](#bool) |  |  |
| content | [ExportContent](#woxqaq-v1-ExportContent) |  |  |
| option | [ExportAccordingToTable.TableOption](#woxqaq-v1-ExportAccordingToTable-TableOption) | repeated |  |






<a name="woxqaq-v1-ExportAccordingToTable-TableOption"></a>

### ExportAccordingToTable.TableOption



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| table | [string](#string) |  |  |
| all_field | [bool](#bool) |  |  |
| fields | [string](#string) | repeated |  |
| filter | [string](#string) |  |  |






<a name="woxqaq-v1-ExportBySQL"></a>

### ExportBySQL



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| statement | [string](#string) |  |  |






<a name="woxqaq-v1-GetTaskDetailRequest"></a>

### GetTaskDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [string](#string) |  |  |






<a name="woxqaq-v1-GetTaskDetailResponse"></a>

### GetTaskDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |
| ded | [DataExportDetail](#woxqaq-v1-DataExportDetail) |  |  |






<a name="woxqaq-v1-ListTaskRequest"></a>

### ListTaskRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [TaskKind](#woxqaq-v1-TaskKind) |  |  |
| current | [int32](#int32) |  |  |
| page_size | [int32](#int32) |  |  |
| status | [TaskStatus](#woxqaq-v1-TaskStatus) |  |  |
| dsn | [string](#string) |  |  |
| role | [TaskRole](#woxqaq-v1-TaskRole) |  |  |
| task_id | [string](#string) |  |  |






<a name="woxqaq-v1-ListTaskResponse"></a>

### ListTaskResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |
| data | [TaskListData](#woxqaq-v1-TaskListData) | repeated |  |
| page_size | [int32](#int32) |  |  |
| total | [int32](#int32) |  |  |






<a name="woxqaq-v1-RejectRequest"></a>

### RejectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [uint64](#uint64) |  |  |
| approver | [string](#string) |  |  |
| comment | [string](#string) |  |  |






<a name="woxqaq-v1-RejectResponse"></a>

### RejectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="woxqaq-v1-SubmitTaskRequest"></a>

### SubmitTaskRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| kind | [TaskKind](#woxqaq-v1-TaskKind) |  |  |
| dsn | [string](#string) |  | TODO: datasource abstract |
| ded | [DataExportDetail](#woxqaq-v1-DataExportDetail) |  |  |






<a name="woxqaq-v1-SubmitTaskResponse"></a>

### SubmitTaskResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |
| task_id | [string](#string) |  |  |






<a name="woxqaq-v1-TaskListData"></a>

### TaskListData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| dsn | [string](#string) |  |  |
| kind | [TaskKind](#woxqaq-v1-TaskKind) |  |  |
| database | [string](#string) |  |  |
| create_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| update_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| submitter | [string](#string) |  |  |
| status | [TaskStatus](#woxqaq-v1-TaskStatus) |  |  |
| ded | [DataExportDetail](#woxqaq-v1-DataExportDetail) |  |  |





 


<a name="woxqaq-v1-Encoding"></a>

### Encoding


| Name | Number | Description |
| ---- | ------ | ----------- |
| ENCODING_UNSPECIFIED | 0 |  |
| ENCODING_UTF8 | 1 |  |
| ENCODING_GBK | 2 |  |



<a name="woxqaq-v1-Exec"></a>

### Exec


| Name | Number | Description |
| ---- | ------ | ----------- |
| EXEC_UNSPECIFIED | 0 |  |
| EXEC_APPROVER | 1 |  |
| EXEC_COMMITTER | 2 |  |
| EXEC_AUTO | 3 |  |



<a name="woxqaq-v1-ExportContent"></a>

### ExportContent


| Name | Number | Description |
| ---- | ------ | ----------- |
| EXPORT_CONTENT_UNSPECIFIED | 0 |  |
| EXPORT_CONTENT_DATA | 1 |  |
| EXPORT_CONTENT_STRUCT | 2 |  |
| EXPORT_CONTENT_ALL | 3 |  |



<a name="woxqaq-v1-ExportMethod"></a>

### ExportMethod


| Name | Number | Description |
| ---- | ------ | ----------- |
| EXPORT_METHOD_UNSPECIFIED | 0 |  |
| EXPORT_METHOD_SQL | 1 |  |
| EXPORT_METHOD_TABLES | 2 |  |



<a name="woxqaq-v1-ExportType"></a>

### ExportType


| Name | Number | Description |
| ---- | ------ | ----------- |
| EXPORT_TYPE_UNSPECIFIED | 0 |  |
| EXPORT_TYPE_CSV | 1 |  |
| EXPORT_TYPE_EXCEL | 2 |  |



<a name="woxqaq-v1-TaskKind"></a>

### TaskKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| TASK_KIND_UNSPECIFIED | 0 |  |
| TASK_KIND_DATA_IMPORT | 1 |  |
| TASK_KIND_DATA_EXPORT | 2 |  |
| TASK_KIND_DATA_GENERATE | 3 |  |



<a name="woxqaq-v1-TaskRole"></a>

### TaskRole


| Name | Number | Description |
| ---- | ------ | ----------- |
| TASK_ROLE_UNSPECIFIED | 0 |  |
| TASK_ROLE_COMMITTER | 1 |  |
| TASK_ROLE_TO_APPROVE | 2 |  |
| TASK_ROLE_TO_EXECUTE | 3 |  |
| TASK_ROLE_APPROVED | 4 |  |
| TASK_ROLE_EXECUTED | 5 |  |



<a name="woxqaq-v1-TaskStatus"></a>

### TaskStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| TASK_STATUS_UNSPECIFIED | 0 |  |
| TASK_STATUS_NOT_START | 1 |  |
| TASK_STATUS_PRE_CHECK | 2 |  |
| TASK_STATUS_PRE_CHECK_SUCCESS | 3 |  |
| TASK_STATUS_PRE_CHECK_FAILED | 4 |  |
| TASK_STATUS_APPROVING | 5 |  |
| TASK_STATUS_APPROVED | 6 |  |
| TASK_STATUS_NOT_APPROVED | 7 |  |



<a name="woxqaq-v1-TriggerKind"></a>

### TriggerKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| TRIGGER_KIND_UNSPECIFIED | 0 |  |
| TRIGGER_KIND_CURRENT | 1 |  |
| TRIGGER_KIND_CRON | 2 |  |


 

 


<a name="woxqaq-v1-TaskService"></a>

### TaskService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetTaskDetail | [GetTaskDetailRequest](#woxqaq-v1-GetTaskDetailRequest) | [GetTaskDetailResponse](#woxqaq-v1-GetTaskDetailResponse) |  |
| ListTaskDetail | [ListTaskRequest](#woxqaq-v1-ListTaskRequest) | [ListTaskResponse](#woxqaq-v1-ListTaskResponse) |  |
| SubmitTask | [SubmitTaskRequest](#woxqaq-v1-SubmitTaskRequest) | [SubmitTaskResponse](#woxqaq-v1-SubmitTaskResponse) |  |
| Approve | [ApproveRequest](#woxqaq-v1-ApproveRequest) | [ApproveResponse](#woxqaq-v1-ApproveResponse) |  |
| Reject | [RejectRequest](#woxqaq-v1-RejectRequest) | [RejectResponse](#woxqaq-v1-RejectResponse) |  |

 



<a name="v1_upload_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/upload_service.proto



<a name="woxqaq-v1-PreSignRequest"></a>

### PreSignRequest







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

