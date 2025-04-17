# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [enum.proto](#enum-proto)
    - [Exec](#woxQAQ-v1-Exec)
    - [TaskKind](#woxQAQ-v1-TaskKind)
    - [TriggerKind](#woxQAQ-v1-TriggerKind)
  
- [stub.proto](#stub-proto)
    - [ApproveRequest](#woxQAQ-v1-ApproveRequest)
    - [PreSignRequest](#woxQAQ-v1-PreSignRequest)
    - [PreSignResponse](#woxQAQ-v1-PreSignResponse)
    - [SubmitTaskRequest](#woxQAQ-v1-SubmitTaskRequest)
  
    - [UploadService](#woxQAQ-v1-UploadService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="enum-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## enum.proto


 


<a name="woxQAQ-v1-Exec"></a>

### Exec


| Name | Number | Description |
| ---- | ------ | ----------- |
| EXEC_APPROVER | 0 |  |
| EXEC_COMMITTER | 1 |  |
| EXEC_AUTO | 2 |  |



<a name="woxQAQ-v1-TaskKind"></a>

### TaskKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| DATA_IMPORT | 0 |  |
| DATA_EXPORT | 1 |  |
| DATA_GENERATE | 2 |  |



<a name="woxQAQ-v1-TriggerKind"></a>

### TriggerKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| current | 0 |  |
| cron | 1 |  |


 

 

 



<a name="stub-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub.proto



<a name="woxQAQ-v1-ApproveRequest"></a>

### ApproveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| progressId | [uint64](#uint64) |  |  |
| approver | [string](#string) |  |  |
| option | [string](#string) |  |  |
| exector | [Exec](#woxQAQ-v1-Exec) |  |  |






<a name="woxQAQ-v1-PreSignRequest"></a>

### PreSignRequest







<a name="woxQAQ-v1-PreSignResponse"></a>

### PreSignResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="woxQAQ-v1-SubmitTaskRequest"></a>

### SubmitTaskRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [TaskKind](#woxQAQ-v1-TaskKind) |  |  |





 

 

 


<a name="woxQAQ-v1-UploadService"></a>

### UploadService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetDataFilePreSignUrl | [PreSignRequest](#woxQAQ-v1-PreSignRequest) | [PreSignResponse](#woxQAQ-v1-PreSignResponse) |  |

 



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

