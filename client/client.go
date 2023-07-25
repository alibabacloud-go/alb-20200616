// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddEntriesToAclRequest struct {
	// The IP entries that you want to add. You can add up to 20 IP entries in each call.
	AclEntries []*AddEntriesToAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The ACL ID.
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s AddEntriesToAclRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclRequest) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclRequest) SetAclEntries(v []*AddEntriesToAclRequestAclEntries) *AddEntriesToAclRequest {
	s.AclEntries = v
	return s
}

func (s *AddEntriesToAclRequest) SetAclId(v string) *AddEntriesToAclRequest {
	s.AclId = &v
	return s
}

func (s *AddEntriesToAclRequest) SetClientToken(v string) *AddEntriesToAclRequest {
	s.ClientToken = &v
	return s
}

func (s *AddEntriesToAclRequest) SetDryRun(v bool) *AddEntriesToAclRequest {
	s.DryRun = &v
	return s
}

type AddEntriesToAclRequestAclEntries struct {
	// The description of the IP entry. The description must be 2 to 256 characters in length, and can contain letters, digits, and the following special characters: , . ; / @ \_ -.
	//
	// You can add up to 20 IP entries in each call.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The CIDR block of the IP entry.
	//
	// You can add up to 20 IP entries in each call.
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s AddEntriesToAclRequestAclEntries) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclRequestAclEntries) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclRequestAclEntries) SetDescription(v string) *AddEntriesToAclRequestAclEntries {
	s.Description = &v
	return s
}

func (s *AddEntriesToAclRequestAclEntries) SetEntry(v string) *AddEntriesToAclRequestAclEntries {
	s.Entry = &v
	return s
}

type AddEntriesToAclResponseBody struct {
	// The asynchronous task ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddEntriesToAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclResponseBody) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclResponseBody) SetJobId(v string) *AddEntriesToAclResponseBody {
	s.JobId = &v
	return s
}

func (s *AddEntriesToAclResponseBody) SetRequestId(v string) *AddEntriesToAclResponseBody {
	s.RequestId = &v
	return s
}

type AddEntriesToAclResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddEntriesToAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddEntriesToAclResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclResponse) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclResponse) SetHeaders(v map[string]*string) *AddEntriesToAclResponse {
	s.Headers = v
	return s
}

func (s *AddEntriesToAclResponse) SetStatusCode(v int32) *AddEntriesToAclResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEntriesToAclResponse) SetBody(v *AddEntriesToAclResponseBody) *AddEntriesToAclResponse {
	s.Body = v
	return s
}

type AddServersToServerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: prechecks the request, but does not add a backend server to a server group. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The server group ID.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The backend servers that you want to add to the server group. You can specify up to 40 backend servers in each call.
	Servers []*AddServersToServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s AddServersToServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupRequest) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupRequest) SetClientToken(v string) *AddServersToServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetDryRun(v bool) *AddServersToServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetServerGroupId(v string) *AddServersToServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetServers(v []*AddServersToServerGroupRequestServers) *AddServersToServerGroupRequest {
	s.Servers = v
	return s
}

type AddServersToServerGroupRequestServers struct {
	// The description of the backend server. The description must be 2 to 256 characters in length and can contain letters, digits, periods (.), underscores (\_), hyphens (-), commas (,), semicolons (;), forward slashes (/), and at signs (@). You can specify up to 40 servers in each call.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port used by the backend server. Valid values: **1** to **65535**. You can specify up to 40 server IDs in each call.
	//
	// > This parameter is required if the **ServerType** parameter is set to **Ecs**, **Eni**, **Eci**, or **Ip**. You do not need to set this parameter if **ServerType** is set to **Fc**.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Specifies whether to enable the remote IP address feature. You can specify up to 40 server IDs in each call. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// Regions that support the remote IP address feature: China (Hangzhou), China (Shenzhen), China (Qingdao), China (Beijing), China (Zhangjiakou), China (Ulanqab), China (Shanghai), China (Chengdu), China (Guangzhou), China (Hong Kong), US (Virginia), Japan (Tokyo), UK (London), US (Silicon Valley), Germany (Frankfurt), Indonesia (Jakarta), Singapore, Malaysia (Kuala Lumpur), Australia (Sydney), and India (Mumbai).
	//
	// > If **ServerType** is set to **Ip**, this parameter is available.
	RemoteIpEnabled *bool `json:"RemoteIpEnabled,omitempty" xml:"RemoteIpEnabled,omitempty"`
	// The backend server ID. You can specify up to 40 server IDs in each call.
	//
	// *   If ServerType is set to **Instance**, set the ServerId parameter to the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. These backend servers are specified by **Ecs**, **Eni**, or **Eci**.
	// *   If ServerType is set to **Ip**, set the ServerId parameter to an IP address specified in the server group.
	// *   If the backend server group is of the **Fc** type, set this parameter to the Alibaba Cloud Resource Name (ARN) of a function.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address in inclusive ENI mode. You can specify up to 40 server IDs in each call.
	//
	// > You do not need to set this parameter if **ServerType** is set to **Fc**.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server that you want to add to the server group. You can specify up to 40 server IDs in each call. Valid values:
	//
	// *   **Ecs**
	// *   **Eni**
	// *   **Eci**
	// *   **Ip**
	// *   **fc**
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The weight of the backend server. Valid values: **0** to **100**. Default value: **100**. If the weight of a backend server is set to **0**, no requests are forwarded to the backend server. You can specify up to 40 server IDs in each call.
	//
	// > You do not need to set this parameter if **ServerType** is set to **Fc**.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AddServersToServerGroupRequestServers) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupRequestServers) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupRequestServers) SetDescription(v string) *AddServersToServerGroupRequestServers {
	s.Description = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetPort(v int32) *AddServersToServerGroupRequestServers {
	s.Port = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetRemoteIpEnabled(v bool) *AddServersToServerGroupRequestServers {
	s.RemoteIpEnabled = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerId(v string) *AddServersToServerGroupRequestServers {
	s.ServerId = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerIp(v string) *AddServersToServerGroupRequestServers {
	s.ServerIp = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerType(v string) *AddServersToServerGroupRequestServers {
	s.ServerType = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetWeight(v int32) *AddServersToServerGroupRequestServers {
	s.Weight = &v
	return s
}

type AddServersToServerGroupResponseBody struct {
	// The ID of the asynchronous job.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddServersToServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupResponseBody) SetJobId(v string) *AddServersToServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetRequestId(v string) *AddServersToServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddServersToServerGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddServersToServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddServersToServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupResponse) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupResponse) SetHeaders(v map[string]*string) *AddServersToServerGroupResponse {
	s.Headers = v
	return s
}

func (s *AddServersToServerGroupResponse) SetStatusCode(v int32) *AddServersToServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddServersToServerGroupResponse) SetBody(v *AddServersToServerGroupResponseBody) *AddServersToServerGroupResponse {
	s.Body = v
	return s
}

type ApplyHealthCheckTemplateToServerGroupRequest struct {
	// The ID of the asynchronous task.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not set this parameter, the system automatically uses **RequestId** as **ClientToken**. **RequestId** may be different for each API request.
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// Specifies whether to check the request without applying the template. Valid values:
	//
	// *   **true**: checks the request without applying the template. The system checks the required parameters, request syntax, and limits. If the request fails to pass the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the API request. If the request passes the check, a 2xx HTTP status code is returned and the template is applied.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s ApplyHealthCheckTemplateToServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyHealthCheckTemplateToServerGroupRequest) GoString() string {
	return s.String()
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) SetClientToken(v string) *ApplyHealthCheckTemplateToServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) SetDryRun(v bool) *ApplyHealthCheckTemplateToServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) SetHealthCheckTemplateId(v string) *ApplyHealthCheckTemplateToServerGroupRequest {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupRequest) SetServerGroupId(v string) *ApplyHealthCheckTemplateToServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

type ApplyHealthCheckTemplateToServerGroupResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyHealthCheckTemplateToServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyHealthCheckTemplateToServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyHealthCheckTemplateToServerGroupResponseBody) SetJobId(v string) *ApplyHealthCheckTemplateToServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupResponseBody) SetRequestId(v string) *ApplyHealthCheckTemplateToServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type ApplyHealthCheckTemplateToServerGroupResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyHealthCheckTemplateToServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyHealthCheckTemplateToServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyHealthCheckTemplateToServerGroupResponse) GoString() string {
	return s.String()
}

func (s *ApplyHealthCheckTemplateToServerGroupResponse) SetHeaders(v map[string]*string) *ApplyHealthCheckTemplateToServerGroupResponse {
	s.Headers = v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupResponse) SetStatusCode(v int32) *ApplyHealthCheckTemplateToServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupResponse) SetBody(v *ApplyHealthCheckTemplateToServerGroupResponseBody) *ApplyHealthCheckTemplateToServerGroupResponse {
	s.Body = v
	return s
}

type AssociateAclsWithListenerRequest struct {
	AclIds      []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	AclType     *string   `json:"AclType,omitempty" xml:"AclType,omitempty"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the request.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s AssociateAclsWithListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateAclsWithListenerRequest) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerRequest) SetAclIds(v []*string) *AssociateAclsWithListenerRequest {
	s.AclIds = v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetAclType(v string) *AssociateAclsWithListenerRequest {
	s.AclType = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetClientToken(v string) *AssociateAclsWithListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetDryRun(v bool) *AssociateAclsWithListenerRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetListenerId(v string) *AssociateAclsWithListenerRequest {
	s.ListenerId = &v
	return s
}

type AssociateAclsWithListenerResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAclsWithListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateAclsWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerResponseBody) SetJobId(v string) *AssociateAclsWithListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *AssociateAclsWithListenerResponseBody) SetRequestId(v string) *AssociateAclsWithListenerResponseBody {
	s.RequestId = &v
	return s
}

type AssociateAclsWithListenerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssociateAclsWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateAclsWithListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateAclsWithListenerResponse) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerResponse) SetHeaders(v map[string]*string) *AssociateAclsWithListenerResponse {
	s.Headers = v
	return s
}

func (s *AssociateAclsWithListenerResponse) SetStatusCode(v int32) *AssociateAclsWithListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateAclsWithListenerResponse) SetBody(v *AssociateAclsWithListenerResponseBody) *AssociateAclsWithListenerResponse {
	s.Body = v
	return s
}

type AssociateAdditionalCertificatesWithListenerRequest struct {
	Certificates []*AssociateAdditionalCertificatesWithListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId** as **ClientToken**. The value of **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, an HTTP `2xx` status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener. You must specify the ID of an HTTPS listener or a QUIC listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerRequest) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetCertificates(v []*AssociateAdditionalCertificatesWithListenerRequestCertificates) *AssociateAdditionalCertificatesWithListenerRequest {
	s.Certificates = v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetClientToken(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetDryRun(v bool) *AssociateAdditionalCertificatesWithListenerRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetListenerId(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.ListenerId = &v
	return s
}

type AssociateAdditionalCertificatesWithListenerRequestCertificates struct {
	// The ID of the certificate. Only server certificates are supported.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerRequestCertificates) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) SetCertificateId(v string) *AssociateAdditionalCertificatesWithListenerRequestCertificates {
	s.CertificateId = &v
	return s
}

type AssociateAdditionalCertificatesWithListenerResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) SetJobId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) SetRequestId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody {
	s.RequestId = &v
	return s
}

type AssociateAdditionalCertificatesWithListenerResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssociateAdditionalCertificatesWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateAdditionalCertificatesWithListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerResponse) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerResponse) SetHeaders(v map[string]*string) *AssociateAdditionalCertificatesWithListenerResponse {
	s.Headers = v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponse) SetStatusCode(v int32) *AssociateAdditionalCertificatesWithListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponse) SetBody(v *AssociateAdditionalCertificatesWithListenerResponseBody) *AssociateAdditionalCertificatesWithListenerResponse {
	s.Body = v
	return s
}

type AttachCommonBandwidthPackageToLoadBalancerRequest struct {
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The ID of the asynchronous task.
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun         *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachCommonBandwidthPackageToLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachCommonBandwidthPackageToLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetBandwidthPackageId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetClientToken(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetDryRun(v bool) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetLoadBalancerId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetRegionId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.RegionId = &v
	return s
}

type AttachCommonBandwidthPackageToLoadBalancerResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachCommonBandwidthPackageToLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachCommonBandwidthPackageToLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *AttachCommonBandwidthPackageToLoadBalancerResponseBody) SetJobId(v string) *AttachCommonBandwidthPackageToLoadBalancerResponseBody {
	s.JobId = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerResponseBody) SetRequestId(v string) *AttachCommonBandwidthPackageToLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type AttachCommonBandwidthPackageToLoadBalancerResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachCommonBandwidthPackageToLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachCommonBandwidthPackageToLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachCommonBandwidthPackageToLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *AttachCommonBandwidthPackageToLoadBalancerResponse) SetHeaders(v map[string]*string) *AttachCommonBandwidthPackageToLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerResponse) SetStatusCode(v int32) *AttachCommonBandwidthPackageToLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerResponse) SetBody(v *AttachCommonBandwidthPackageToLoadBalancerResponseBody) *AttachCommonBandwidthPackageToLoadBalancerResponse {
	s.Body = v
	return s
}

type CreateAScriptsRequest struct {
	AScripts []*CreateAScriptsRequestAScripts `json:"AScripts,omitempty" xml:"AScripts,omitempty" type:"Repeated"`
	// The ID of the asynchronous task.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IDs of AScript rules.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s CreateAScriptsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAScriptsRequest) GoString() string {
	return s.String()
}

func (s *CreateAScriptsRequest) SetAScripts(v []*CreateAScriptsRequestAScripts) *CreateAScriptsRequest {
	s.AScripts = v
	return s
}

func (s *CreateAScriptsRequest) SetClientToken(v string) *CreateAScriptsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAScriptsRequest) SetDryRun(v bool) *CreateAScriptsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAScriptsRequest) SetListenerId(v string) *CreateAScriptsRequest {
	s.ListenerId = &v
	return s
}

type CreateAScriptsRequestAScripts struct {
	// The ID of the AScript rule.
	AScriptName   *string `json:"AScriptName,omitempty" xml:"AScriptName,omitempty"`
	Enabled       *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
}

func (s CreateAScriptsRequestAScripts) String() string {
	return tea.Prettify(s)
}

func (s CreateAScriptsRequestAScripts) GoString() string {
	return s.String()
}

func (s *CreateAScriptsRequestAScripts) SetAScriptName(v string) *CreateAScriptsRequestAScripts {
	s.AScriptName = &v
	return s
}

func (s *CreateAScriptsRequestAScripts) SetEnabled(v bool) *CreateAScriptsRequestAScripts {
	s.Enabled = &v
	return s
}

func (s *CreateAScriptsRequestAScripts) SetScriptContent(v string) *CreateAScriptsRequestAScripts {
	s.ScriptContent = &v
	return s
}

type CreateAScriptsResponseBody struct {
	AScriptIds []*CreateAScriptsResponseBodyAScriptIds `json:"AScriptIds,omitempty" xml:"AScriptIds,omitempty" type:"Repeated"`
	JobId      *string                                 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAScriptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAScriptsResponseBody) SetAScriptIds(v []*CreateAScriptsResponseBodyAScriptIds) *CreateAScriptsResponseBody {
	s.AScriptIds = v
	return s
}

func (s *CreateAScriptsResponseBody) SetJobId(v string) *CreateAScriptsResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateAScriptsResponseBody) SetRequestId(v string) *CreateAScriptsResponseBody {
	s.RequestId = &v
	return s
}

type CreateAScriptsResponseBodyAScriptIds struct {
	AScriptId *string `json:"AScriptId,omitempty" xml:"AScriptId,omitempty"`
}

func (s CreateAScriptsResponseBodyAScriptIds) String() string {
	return tea.Prettify(s)
}

func (s CreateAScriptsResponseBodyAScriptIds) GoString() string {
	return s.String()
}

func (s *CreateAScriptsResponseBodyAScriptIds) SetAScriptId(v string) *CreateAScriptsResponseBodyAScriptIds {
	s.AScriptId = &v
	return s
}

type CreateAScriptsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAScriptsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAScriptsResponse) GoString() string {
	return s.String()
}

func (s *CreateAScriptsResponse) SetHeaders(v map[string]*string) *CreateAScriptsResponse {
	s.Headers = v
	return s
}

func (s *CreateAScriptsResponse) SetStatusCode(v int32) *CreateAScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAScriptsResponse) SetBody(v *CreateAScriptsResponseBody) *CreateAScriptsResponse {
	s.Body = v
	return s
}

type CreateAclRequest struct {
	// The ID of the ACL.
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The ID of the asynchronous task.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request.
	DryRun          *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateAclRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAclRequest) GoString() string {
	return s.String()
}

func (s *CreateAclRequest) SetAclName(v string) *CreateAclRequest {
	s.AclName = &v
	return s
}

func (s *CreateAclRequest) SetClientToken(v string) *CreateAclRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAclRequest) SetDryRun(v bool) *CreateAclRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAclRequest) SetResourceGroupId(v string) *CreateAclRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateAclResponseBody struct {
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAclResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclResponseBody) SetAclId(v string) *CreateAclResponseBody {
	s.AclId = &v
	return s
}

func (s *CreateAclResponseBody) SetJobId(v string) *CreateAclResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateAclResponseBody) SetRequestId(v string) *CreateAclResponseBody {
	s.RequestId = &v
	return s
}

type CreateAclResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAclResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAclResponse) GoString() string {
	return s.String()
}

func (s *CreateAclResponse) SetHeaders(v map[string]*string) *CreateAclResponse {
	s.Headers = v
	return s
}

func (s *CreateAclResponse) SetStatusCode(v int32) *CreateAclResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAclResponse) SetBody(v *CreateAclResponseBody) *CreateAclResponse {
	s.Body = v
	return s
}

type CreateHealthCheckTemplateRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a **2xx** HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The interval at which health checks are performed.
	//
	// Valid values: **1 to 50**.
	//
	// Default value: **2**.
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The port that is used for health checks.
	//
	// Valid values: **0 to 65535**.
	//
	// Default value: **0**. If you set the value to 0, the port of a backend server is used for health checks.
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that you want to use for the health check.
	//
	// Default value: **$SERVER_IP**. The domain name must be 1 to 80 characters in length. The domain name must meet the following requirements:
	//
	// *   The domain name can contain lowercase letters, digits, hyphens (-), and periods (.).
	// *   The domain name must contain at least one period (.) but cannot start or end with a period (.).
	// *   The rightmost domain label can contain only letters but cannot contain digits or hyphens (-).
	// *   Other fields cannot start or end with a hyphen (-).
	//
	// This parameter is required only if the **HealthCheckProtocol** parameter is set to **HTTP**.
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The HTTP version that is used for health checks.
	//
	// Valid values: **HTTP 1.0** and **HTTP 1.1**.
	//
	// Default value: **HTTP 1.1**.
	//
	// > This parameter is valid only if the `HealthCheckProtocol` parameter is set to **HTTP**.
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval at which health checks are performed.
	//
	// Valid values: **1 to 50**.
	//
	// Default value: **2**.
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The method that you want to use for the health check. Valid values:
	//
	// *   **HEAD**: By default, the ALB instance sends HEAD requests to a backend server to perform HTTP health checks.
	// *   **POST**: gRPC health checks automatically use the POST method.
	// *   **GET**: If the length of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	//
	// > This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The URL that is used for health checks.
	//
	// It must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&). It can also contain the following extended characters: `_ ; ~ ! ( ) * [ ] @ $ ^ : \" , +`. The URL must start with a forward slash (/).
	//
	// > This parameter is valid only if the `HealthCheckProtocol` parameter is set to **HTTP**.
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that you want to use for health checks. Valid values:
	//
	// *   **HTTP** (default): To perform HTTP health checks, ALB sends HEAD or GET requests to a backend server to check whether the backend server is healthy.
	// *   **TCP**: To perform TCP health checks, ALB sends SYN packets to a backend server to check whether the port of the backend server is available to receive requests.
	// *   **gRPC**: To perform gRPC health checks, ALB sends POST or GET requests to a backend server to check whether the backend server is healthy.
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The name of the health check template.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// The timeout period of a health check. If a backend server does not respond within the specified timeout period, the backend server fails the health check.
	//
	// Valid values: **1 to 300**.
	//
	// Default value: **5**.
	//
	// > If the value of the `HealthCheckTimeout` parameter is smaller than that of the `HealthCheckInterval` parameter, the timeout period specified by the `HealthCheckTimeout` parameter is ignored and the value of the `HealthCheckInterval` parameter is used as the timeout period.
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail** to **success**.
	//
	// Valid values: **2 to 10**.
	//
	// Default value: **3**.
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success** to **fail**.
	//
	// Valid values: **2 to 10**.
	//
	// Default value: **3**.
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s CreateHealthCheckTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHealthCheckTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckTemplateRequest) SetClientToken(v string) *CreateHealthCheckTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetDryRun(v bool) *CreateHealthCheckTemplateRequest {
	s.DryRun = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckCodes(v []*string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckCodes = v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckConnectPort(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckHost(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckHost = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckHttpVersion(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckInterval(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckMethod(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckPath(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckProtocol(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckTemplateName(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckTimeout(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthyThreshold(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetUnhealthyThreshold(v int32) *CreateHealthCheckTemplateRequest {
	s.UnhealthyThreshold = &v
	return s
}

type CreateHealthCheckTemplateResponseBody struct {
	// The ID of the health check template.
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHealthCheckTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHealthCheckTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckTemplateResponseBody) SetHealthCheckTemplateId(v string) *CreateHealthCheckTemplateResponseBody {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *CreateHealthCheckTemplateResponseBody) SetRequestId(v string) *CreateHealthCheckTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateHealthCheckTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHealthCheckTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHealthCheckTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHealthCheckTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckTemplateResponse) SetHeaders(v map[string]*string) *CreateHealthCheckTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateHealthCheckTemplateResponse) SetStatusCode(v int32) *CreateHealthCheckTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHealthCheckTemplateResponse) SetBody(v *CreateHealthCheckTemplateResponseBody) *CreateHealthCheckTemplateResponse {
	s.Body = v
	return s
}

type CreateListenerRequest struct {
	// A list of certificates.
	CaCertificates []*CreateListenerRequestCaCertificates `json:"CaCertificates,omitempty" xml:"CaCertificates,omitempty" type:"Repeated"`
	// Specifies whether to enable mutual authentication. Valid values:
	//
	// *   **true**
	// *   **false** (default):
	CaEnabled *bool `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// A list of certificates.
	Certificates []*CreateListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The actions of the forwarding rule.
	DefaultActions []*CreateListenerRequestDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable `GZIP` compression to compress specific types of files. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// Specifies whether to enable `HTTP/2`. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The timeout period of an idle connection. Unit: seconds.
	//
	// Valid values: **1 to 60**.
	//
	// Default value: **15**.
	//
	// If no requests are received within the specified timeout period, ALB closes the current connection. When a new request is received, ALB establishes a new connection.
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The name of the listener.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (\_). Regular expressions are supported.
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The frontend port that is used by the ALB instance.
	//
	// Valid values: **1 to 65535**.
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listener protocol.
	//
	// Valid values: **HTTP**, **HTTPS**, and **QUIC**.
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the Application Load Balancer (ALB) instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// Selects a QUIC listener and associates it with the HTTPS listener of the ALB instance.
	QuicConfig *CreateListenerRequestQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// The timeout period of a request. Unit: seconds.
	//
	// Valid values: **1 to 180**.
	//
	// Default value: **60**.
	//
	// If no response is received from the backend server during the request timeout period, ALB sends an `HTTP 504` error code to the client.
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The ID of the security policy. System security policies and custom security policies are supported.
	//
	// Default value: **tls_cipher_policy\_1\_0** (system security policy).
	//
	// > This parameter is available only when you create an HTTPS listener.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The configuration of the XForward header.
	XForwardedForConfig *CreateListenerRequestXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s CreateListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateListenerRequest) SetCaCertificates(v []*CreateListenerRequestCaCertificates) *CreateListenerRequest {
	s.CaCertificates = v
	return s
}

func (s *CreateListenerRequest) SetCaEnabled(v bool) *CreateListenerRequest {
	s.CaEnabled = &v
	return s
}

func (s *CreateListenerRequest) SetCertificates(v []*CreateListenerRequestCertificates) *CreateListenerRequest {
	s.Certificates = v
	return s
}

func (s *CreateListenerRequest) SetClientToken(v string) *CreateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateListenerRequest) SetDefaultActions(v []*CreateListenerRequestDefaultActions) *CreateListenerRequest {
	s.DefaultActions = v
	return s
}

func (s *CreateListenerRequest) SetDryRun(v bool) *CreateListenerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateListenerRequest) SetGzipEnabled(v bool) *CreateListenerRequest {
	s.GzipEnabled = &v
	return s
}

func (s *CreateListenerRequest) SetHttp2Enabled(v bool) *CreateListenerRequest {
	s.Http2Enabled = &v
	return s
}

func (s *CreateListenerRequest) SetIdleTimeout(v int32) *CreateListenerRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateListenerRequest) SetListenerDescription(v string) *CreateListenerRequest {
	s.ListenerDescription = &v
	return s
}

func (s *CreateListenerRequest) SetListenerPort(v int32) *CreateListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateListenerRequest) SetListenerProtocol(v string) *CreateListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *CreateListenerRequest) SetLoadBalancerId(v string) *CreateListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateListenerRequest) SetQuicConfig(v *CreateListenerRequestQuicConfig) *CreateListenerRequest {
	s.QuicConfig = v
	return s
}

func (s *CreateListenerRequest) SetRequestTimeout(v int32) *CreateListenerRequest {
	s.RequestTimeout = &v
	return s
}

func (s *CreateListenerRequest) SetSecurityPolicyId(v string) *CreateListenerRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *CreateListenerRequest) SetXForwardedForConfig(v *CreateListenerRequestXForwardedForConfig) *CreateListenerRequest {
	s.XForwardedForConfig = v
	return s
}

type CreateListenerRequestCaCertificates struct {
}

func (s CreateListenerRequestCaCertificates) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestCaCertificates) GoString() string {
	return s.String()
}

type CreateListenerRequestCertificates struct {
	// The ID of the certificate. Only server certificates are supported. You can specify a maximum of 20 certificate IDs.
	//
	// > This parameter is required if **ListenerProtocol** is set to **HTTPS** or **QUIC**.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s CreateListenerRequestCertificates) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCertificates) SetCertificateId(v string) *CreateListenerRequestCertificates {
	s.CertificateId = &v
	return s
}

type CreateListenerRequestDefaultActions struct {
	// Specifies the configurations of the forwarding action. You can specify a maximum of 20 configurations.
	ForwardGroupConfig *CreateListenerRequestDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The type of the action. You can specify only one action type.
	//
	// Set the value to **ForwardGroup** to forward requests to multiple vServer groups.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateListenerRequestDefaultActions) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestDefaultActions) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestDefaultActions) SetForwardGroupConfig(v *CreateListenerRequestDefaultActionsForwardGroupConfig) *CreateListenerRequestDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateListenerRequestDefaultActions) SetType(v string) *CreateListenerRequestDefaultActions {
	s.Type = &v
	return s
}

type CreateListenerRequestDefaultActionsForwardGroupConfig struct {
	// The server group to which requests are forwarded.
	ServerGroupTuples []*CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) *CreateListenerRequestDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// The ID of the server group to which requests are forwarded.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type CreateListenerRequestQuicConfig struct {
	// The ID of the QUIC listener that you want to associate with the HTTPS listener. Only HTTPS listeners support this parameter. This parameter is required when **QuicUpgradeEnabled** is set to **true**.
	//
	// > You must add the HTTPS listener and the QUIC listener to the same ALB instance. In addition, make sure that the QUIC listener has never been associated with another listener.
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// Specifies whether to enable QUIC upgrade. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	//
	// > This parameter is available only when you create an HTTPS listener.
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s CreateListenerRequestQuicConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestQuicConfig) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestQuicConfig) SetQuicListenerId(v string) *CreateListenerRequestQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *CreateListenerRequestQuicConfig) SetQuicUpgradeEnabled(v bool) *CreateListenerRequestQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

type CreateListenerRequestXForwardedForConfig struct {
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertClientVerifyEnabled** is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (\_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertClientVerifyAlias *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-clientverify` header to retrieve the verification result of the client certificate. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertClientVerifyEnabled *bool `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertFingerprintEnabled** is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (\_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertFingerprintAlias *string `json:"XForwardedForClientCertFingerprintAlias,omitempty" xml:"XForwardedForClientCertFingerprintAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-fingerprint` header to retrieve the fingerprint of the client certificate. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertFingerprintEnabled *bool `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertIssuerDNEnabled** is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (\_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertIssuerDNAlias *string `json:"XForwardedForClientCertIssuerDNAlias,omitempty" xml:"XForwardedForClientCertIssuerDNAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-issuerdn` header to retrieve information about the authority that issues the client certificate. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertIssuerDNEnabled *bool `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertSubjectDNEnabled** is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (\_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertSubjectDNAlias *string `json:"XForwardedForClientCertSubjectDNAlias,omitempty" xml:"XForwardedForClientCertSubjectDNAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-subjectdn` header to retrieve information about the owner of the client certificate. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertSubjectDNEnabled *bool `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-Client-Ip` header to obtain the source IP address of the ALB instance. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	//
	// > This parameter is available only when you create an HTTP, HTTPS, or QUIC listener. The feature specified by this parameter is unavailable by default. To use the feature, contact your account manager.
	XForwardedForClientSourceIpsEnabled *bool `json:"XForwardedForClientSourceIpsEnabled,omitempty" xml:"XForwardedForClientSourceIpsEnabled,omitempty"`
	// The trusted proxy IP address.
	//
	// ALB traverses `X-Forwarded-For` backward and selects the first IP address that is not in the trusted IP address list as the real IP address of the client. The IP address is used in source IP address throttling.
	XForwardedForClientSourceIpsTrusted *string `json:"XForwardedForClientSourceIpsTrusted,omitempty" xml:"XForwardedForClientSourceIpsTrusted,omitempty"`
	// Specifies whether to use the `X-Forwarded-Client-Port` header to retrieve the client port. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	//
	// > This parameter is available only when you create an HTTP or HTTPS listener.
	XForwardedForClientSrcPortEnabled *bool `json:"XForwardedForClientSrcPortEnabled,omitempty" xml:"XForwardedForClientSrcPortEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-For` header to retrieve client IP addresses. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	//
	// > This parameter is available only when you create an HTTP or HTTPS listener.
	XForwardedForEnabled *bool `json:"XForwardedForEnabled,omitempty" xml:"XForwardedForEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-Proto` header to retrieve the listener protocol. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	//
	// > This parameter is available only when you create an HTTP, HTTPS, or QUIC listener.
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// Specifies whether to use the `SLB-ID` header to retrieve the ID of the CLB instance. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	//
	// > This parameter is available only when you create an HTTP, HTTPS, or QUIC listener.
	XForwardedForSLBIdEnabled *bool `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-Port` header to retrieve the listener port of the ALB instance. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	//
	// > This parameter is available only when you create an HTTP, HTTPS, or QUIC listener.
	XForwardedForSLBPortEnabled *bool `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s CreateListenerRequestXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientSourceIpsEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientSourceIpsEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientSourceIpsTrusted(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientSourceIpsTrusted = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

type CreateListenerResponseBody struct {
	// The asynchronous task ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The listener ID.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateListenerResponseBody) SetJobId(v string) *CreateListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateListenerResponseBody) SetListenerId(v string) *CreateListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *CreateListenerResponseBody) SetRequestId(v string) *CreateListenerResponseBody {
	s.RequestId = &v
	return s
}

type CreateListenerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateListenerResponse) SetHeaders(v map[string]*string) *CreateListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateListenerResponse) SetStatusCode(v int32) *CreateListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateListenerResponse) SetBody(v *CreateListenerResponseBody) *CreateListenerResponse {
	s.Body = v
	return s
}

type CreateLoadBalancerRequest struct {
	// The mode used to assign IP addresses to zones of the ALB instance. Default value: Dynamic. Valid values:
	//
	// *   **Fixed:** assigns a static IP address to the ALB instance.
	// *   **Dynamic:** dynamically assigns an IP address to each zone of the ALB instance.
	AddressAllocatedMode *string `json:"AddressAllocatedMode,omitempty" xml:"AddressAllocatedMode,omitempty"`
	// The protocol version. Valid values:
	//
	// *   **IPv4:** IPv4.
	// *   **DualStack:** dual stack.
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The type of the address of the ALB instance. Valid values:
	//
	// *   **Internet:** The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. In this case, the ALB instance can be accessed over the Internet.
	// *   **Intranet:** The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. In this case, the ALB instance can be accessed over the VPC in which the ALB instance is deployed.
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters.
	//
	// >  If you do not specify this parameter, the system uses the value of **RequestId** as the value of **ClientToken**. The value of the **RequestId** parameter may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable deletion protection. Default value: false. Valid values:
	//
	// *   **true:** enables deletion protection.
	// *   **false:** disables deletion protection.
	DeletionProtectionEnabled *bool `json:"DeletionProtectionEnabled,omitempty" xml:"DeletionProtectionEnabled,omitempty"`
	// Specifies whether to perform a dry run. Default value: false. Valid values:
	//
	// *   **true:** performs a dry run. The system checks the required parameters, request format, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false:** performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configuration of the billing method of the ALB instance.
	LoadBalancerBillingConfig *CreateLoadBalancerRequestLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	// The edition of the ALB instance. The features and billing rules vary based on the edition of the ALB instance. Valid values:
	//
	// *   **Basic:** basic.
	// *   **Standard:** standard.
	// *   **StandardWithWaf:** WAF-enabled.
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// The name of the ALB instance.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The configuration of the configuration read-only mode.
	ModificationProtectionConfig *CreateLoadBalancerRequestModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the virtual private cloud (VPC) in which you want to create the ALB instance.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zones and the vSwitches. You must specify at least two zones.
	ZoneMappings []*CreateLoadBalancerRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s CreateLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequest) SetAddressAllocatedMode(v string) *CreateLoadBalancerRequest {
	s.AddressAllocatedMode = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressIpVersion(v string) *CreateLoadBalancerRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressType(v string) *CreateLoadBalancerRequest {
	s.AddressType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetClientToken(v string) *CreateLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDeletionProtectionEnabled(v bool) *CreateLoadBalancerRequest {
	s.DeletionProtectionEnabled = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDryRun(v bool) *CreateLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerBillingConfig(v *CreateLoadBalancerRequestLoadBalancerBillingConfig) *CreateLoadBalancerRequest {
	s.LoadBalancerBillingConfig = v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerEdition(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerEdition = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerName(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetModificationProtectionConfig(v *CreateLoadBalancerRequestModificationProtectionConfig) *CreateLoadBalancerRequest {
	s.ModificationProtectionConfig = v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceGroupId(v string) *CreateLoadBalancerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetVpcId(v string) *CreateLoadBalancerRequest {
	s.VpcId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetZoneMappings(v []*CreateLoadBalancerRequestZoneMappings) *CreateLoadBalancerRequest {
	s.ZoneMappings = v
	return s
}

type CreateLoadBalancerRequestLoadBalancerBillingConfig struct {
	// The ID of the Elastic IP Address (EIP) bandwidth plan that is associated with the ALB instance if the ALB instance uses a public IP address.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The billing method of the ALB instance.
	//
	// Set the value to **PostPay**, which specifies the pay-as-you-go billing method.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s CreateLoadBalancerRequestLoadBalancerBillingConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) SetBandwidthPackageId(v string) *CreateLoadBalancerRequestLoadBalancerBillingConfig {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) SetPayType(v string) *CreateLoadBalancerRequestLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

type CreateLoadBalancerRequestModificationProtectionConfig struct {
	// The reason why the configuration read-only mode is enabled. The reason must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The reason must start with a letter.
	//
	// >  This parameter takes effect only when you set the `Status` parameter to **ConsoleProtection**.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Specifies whether to enable the configuration read-only mode for the ALB instance. Valid values:
	//
	// *   **NonProtection:** disables the configuration read-only mode. In this case, you cannot specify the ModificationProtectionReason parameter. If you specify the ModificationProtectionReason parameter, the value is cleared.
	// *   **ConsoleProtection:** enables the configuration read-only mode. In this case, you can specify the ModificationProtectionReason parameter.
	//
	// >  If you set this parameter to **ConsoleProtection**, you cannot modify the configurations of the ALB instance in the ALB console. However, you can call API operations to modify the configurations of the ALB instance.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateLoadBalancerRequestModificationProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestModificationProtectionConfig) SetReason(v string) *CreateLoadBalancerRequestModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *CreateLoadBalancerRequestModificationProtectionConfig) SetStatus(v string) *CreateLoadBalancerRequestModificationProtectionConfig {
	s.Status = &v
	return s
}

type CreateLoadBalancerRequestZoneMappings struct {
	// The ID of the vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of an ALB instance. You can specify up to 10 vSwitch IDs.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone where the ALB instance is deployed. You can specify up to 10 zone IDs.
	//
	// You can call the [DescribeZones](~~36064~~) operation to query the zones of the ALB instance.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLoadBalancerRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestZoneMappings) SetVSwitchId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetZoneId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type CreateLoadBalancerResponseBody struct {
	// The ID of the ALB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponseBody) SetLoadBalancerId(v string) *CreateLoadBalancerResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetRequestId(v string) *CreateLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type CreateLoadBalancerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerResponse) SetStatusCode(v int32) *CreateLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerResponse) SetBody(v *CreateLoadBalancerResponseBody) *CreateLoadBalancerResponse {
	s.Body = v
	return s
}

type CreateRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The direction to which the forwarding rule is applied. Valid values:
	//
	// *   **Request** (default): The forwarding rule is applied to the client requests received by ALB.
	// *   **Response**: The forwarding rule is applied to the responses returned by backend servers.
	//
	// > Basic ALB instances do not support the **Response** value.
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener of the ALB instance.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The priority of the forwarding rule. Valid values: **1 to 10000**. A smaller value indicates a higher priority.
	//
	// > The priorities of the forwarding rules created for the same listener must be unique.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The actions of the forwarding rule.
	//
	// The maximum number of actions that can be specified for a forwarding rule.
	//
	// *   Basic ALB instances: You can specify at most three actions for a forwarding rule.
	// *   Standard ALB instances: You can specify at most five actions for a forwarding rule.
	RuleActions []*CreateRuleRequestRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// The match conditions of the forwarding rule.
	//
	// The maximum number of conditions supported by a forwarding rule.
	//
	// *   Basic ALB instances: You can specify at most five conditions for a forwarding rule.
	// *   Standard ALB instances: You can specify at most 10 conditions for a forwarding rule.
	RuleConditions []*CreateRuleRequestRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The name of the forwarding rule.
	//
	// *   It must be 2 to 128 characters in length.
	// *   It can contain letters, digits, periods (.), underscores (\_), and hyphens (-). It must start with a letter.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) SetClientToken(v string) *CreateRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRuleRequest) SetDirection(v string) *CreateRuleRequest {
	s.Direction = &v
	return s
}

func (s *CreateRuleRequest) SetDryRun(v bool) *CreateRuleRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRuleRequest) SetListenerId(v string) *CreateRuleRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateRuleRequest) SetPriority(v int32) *CreateRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateRuleRequest) SetRuleActions(v []*CreateRuleRequestRuleActions) *CreateRuleRequest {
	s.RuleActions = v
	return s
}

func (s *CreateRuleRequest) SetRuleConditions(v []*CreateRuleRequestRuleConditions) *CreateRuleRequest {
	s.RuleConditions = v
	return s
}

func (s *CreateRuleRequest) SetRuleName(v string) *CreateRuleRequest {
	s.RuleName = &v
	return s
}

type CreateRuleRequestRuleActions struct {
	// The origins allowed.
	CorsConfig *CreateRuleRequestRuleActionsCorsConfig `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	// The configuration of the custom response.
	FixedResponseConfig *CreateRuleRequestRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	// The server groups to which requests are distributed. Each forwarding rule supports at most five server groups.
	ForwardGroupConfig *CreateRuleRequestRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The configuration of the header to be inserted.
	InsertHeaderConfig *CreateRuleRequestRuleActionsInsertHeaderConfig `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// The priority of the action. Valid values: **1 to 50000**. A smaller value indicates a higher priority. The actions of a forwarding rule are applied in descending order of priority. This parameter is required. The priority of each action within a forwarding rule must be unique.
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The redirect configuration.
	//
	// > When you configure the **RedirectConfig** action, you can use the default value only for the **httpCode** parameter. Do not use the default values for the other parameters.
	RedirectConfig *CreateRuleRequestRuleActionsRedirectConfig `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	// The rewrite configuration.
	//
	// > If multiple actions are configured within a forwarding rule, you must set **RewriteConfig** to the value of **ForwardGroup**.
	RewriteConfig *CreateRuleRequestRuleActionsRewriteConfig `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	// The action to throttle traffic.
	TrafficLimitConfig *CreateRuleRequestRuleActionsTrafficLimitConfig `json:"TrafficLimitConfig,omitempty" xml:"TrafficLimitConfig,omitempty" type:"Struct"`
	// The action to mirror traffic.
	TrafficMirrorConfig *CreateRuleRequestRuleActionsTrafficMirrorConfig `json:"TrafficMirrorConfig,omitempty" xml:"TrafficMirrorConfig,omitempty" type:"Struct"`
	// The action type. Valid values:
	//
	// *   **ForwardGroup**: forwards a request to multiple vServer groups.
	// *   **Redirect**: redirects a request.
	// *   **FixedResponse**: returns a custom response.
	// *   **Rewrite**: rewrites a request.
	// *   **InsertHeader**: inserts a header.
	// *   **RemoveHeaderConfig**: deletes a header.
	// *   **TrafficLimitConfig**: throttles network traffic.
	// *   **TrafficMirrorConfig**: mirrors traffic.
	// *   **CorsConfig**: forwards requests based on CORS.
	//
	// The following action types are supported:
	//
	// *   **FinalType**: the last action to be performed in a forwarding rule. Each forwarding rule can contain only one FinalType action. You can specify the **ForwardGroup**, **Redirect**, or **FixedResponse** action as the FinalType action.
	// *   **ExtType**: the action or the actions to be performed before the **FinalType** action. A forwarding rule can contain one or more **ExtType** actions. To specify this parameter, you must also specify **FinalType**. You can specify multiple **InsertHeader** actions or one **Rewrite** action.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRuleRequestRuleActions) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActions) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActions) SetCorsConfig(v *CreateRuleRequestRuleActionsCorsConfig) *CreateRuleRequestRuleActions {
	s.CorsConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetFixedResponseConfig(v *CreateRuleRequestRuleActionsFixedResponseConfig) *CreateRuleRequestRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetForwardGroupConfig(v *CreateRuleRequestRuleActionsForwardGroupConfig) *CreateRuleRequestRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetInsertHeaderConfig(v *CreateRuleRequestRuleActionsInsertHeaderConfig) *CreateRuleRequestRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetOrder(v int32) *CreateRuleRequestRuleActions {
	s.Order = &v
	return s
}

func (s *CreateRuleRequestRuleActions) SetRedirectConfig(v *CreateRuleRequestRuleActionsRedirectConfig) *CreateRuleRequestRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetRewriteConfig(v *CreateRuleRequestRuleActionsRewriteConfig) *CreateRuleRequestRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetTrafficLimitConfig(v *CreateRuleRequestRuleActionsTrafficLimitConfig) *CreateRuleRequestRuleActions {
	s.TrafficLimitConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetTrafficMirrorConfig(v *CreateRuleRequestRuleActionsTrafficMirrorConfig) *CreateRuleRequestRuleActions {
	s.TrafficMirrorConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetType(v string) *CreateRuleRequestRuleActions {
	s.Type = &v
	return s
}

type CreateRuleRequestRuleActionsCorsConfig struct {
	// Specifies whether credentials can be carried in CORS requests. Valid values:
	//
	// *   **on**: yes
	// *   **off**: no
	AllowCredentials *string `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	// The allowed headers for CORS requests.
	AllowHeaders []*string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty" type:"Repeated"`
	// The allowed HTTP methods for CORS requests.
	AllowMethods []*string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty" type:"Repeated"`
	// The allowed origins of CORS requests.
	AllowOrigin []*string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty" type:"Repeated"`
	// The headers that can be exposed.
	ExposeHeaders []*string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty" type:"Repeated"`
	// The maximum cache time of dry run requests in the browser. Unit: seconds.
	//
	// Valid values: **-1** to **172800**.
	MaxAge *int64 `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
}

func (s CreateRuleRequestRuleActionsCorsConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsCorsConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetAllowCredentials(v string) *CreateRuleRequestRuleActionsCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetAllowHeaders(v []*string) *CreateRuleRequestRuleActionsCorsConfig {
	s.AllowHeaders = v
	return s
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetAllowMethods(v []*string) *CreateRuleRequestRuleActionsCorsConfig {
	s.AllowMethods = v
	return s
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetAllowOrigin(v []*string) *CreateRuleRequestRuleActionsCorsConfig {
	s.AllowOrigin = v
	return s
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetExposeHeaders(v []*string) *CreateRuleRequestRuleActionsCorsConfig {
	s.ExposeHeaders = v
	return s
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetMaxAge(v int64) *CreateRuleRequestRuleActionsCorsConfig {
	s.MaxAge = &v
	return s
}

type CreateRuleRequestRuleActionsFixedResponseConfig struct {
	// The content of the custom response. The content can be up to 1 KB in size and can contain only ASCII characters.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The format of the response.
	//
	// Valid values: **text/plain**, **text/css**, **text/html**, **application/javascript**, and **application/json**.
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The HTTP status code in the response. Valid values: **HTTP\_2xx**, **HTTP\_4xx**, and **HTTP\_5xx**. **x** must be a digit.
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s CreateRuleRequestRuleActionsFixedResponseConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) SetContent(v string) *CreateRuleRequestRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) SetContentType(v string) *CreateRuleRequestRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) SetHttpCode(v string) *CreateRuleRequestRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

type CreateRuleRequestRuleActionsForwardGroupConfig struct {
	// The configuration of session persistence for server groups.
	ServerGroupStickySession *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession `json:"ServerGroupStickySession,omitempty" xml:"ServerGroupStickySession,omitempty" type:"Struct"`
	// The server groups to which requests are distributed. Each forwarding rule supports at most five server groups.
	ServerGroupTuples []*CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfig) SetServerGroupStickySession(v *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) *CreateRuleRequestRuleActionsForwardGroupConfig {
	s.ServerGroupStickySession = v
	return s
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) *CreateRuleRequestRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession struct {
	// Specifies whether to enable session persistence. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The timeout period of sessions. Unit: seconds. Valid values: **1** to **86400**. Default value: **1000**.
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) SetEnabled(v bool) *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Enabled = &v
	return s
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) SetTimeout(v int32) *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Timeout = &v
	return s
}

type CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples struct {
	// The server group to which requests are distributed.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The weight of the server group. A larger value specifies a higher weight. A server group with a higher weight receives more requests. Valid values: **0** to **100**.
	//
	// *   If only one destination server group exists and you do not specify a weight, the default value **100** is used.
	// *   If more than one destination server group exists, you must specify weights.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) SetWeight(v int32) *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

type CreateRuleRequestRuleActionsInsertHeaderConfig struct {
	// The key of the header. The key must be 1 to 40 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The header key specified by **InsertHeaderConfig** must be unique.
	//
	// > You cannot specify the following header keys (case-insensitive): `slb-id`, `slb-ip`, `x-forwarded-for`, `x-forwarded-proto`, `x-forwarded-eip`, `x-forwarded-port`, `x-forwarded-client-srcport`, `connection`, `upgrade`, `content-length`, `transfer-encoding`, `keep-alive`, `te`, `host`, `cookie`, `remoteip`, and `authority`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the header to be inserted.
	//
	// *   If **ValueType** is set to **SystemDefined**, you can specify one of the following header values:
	//
	//     *   **ClientSrcPort**: the client port.
	//     *   **ClientSrcIp**: the client IP address.
	//     *   **Protocol**: the request protocol (HTTP or HTTPS).
	//     *   **SLBId**: the ID of the ALB instance.
	//     *   **SLBPort**: the listening port.
	//
	// *   If **ValueType** is set to **UserDefined**, you can specify a custom header value. The header value must be 1 to 128 characters in length, and can contain printable characters whose ASCII values are `greater than or equal to 32 and lower than 127`. You can use asterisks (\*) and question marks (?) as wildcards. The value cannot start or end with a space character.
	//
	// *   If **ValueType** is set to **ReferenceHeader**, you can reference one of the request headers. The header value must be 1 to 128 characters in length, and can contain lowercase letters, digits, underscores (\_), and hyphens (-).
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The type of header. Valid values:
	//
	// *   **UserDefined**: a custom header
	// *   **ReferenceHeader**: a header that references one of the request headers
	// *   **SystemDefined**: a header predefined by the system
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s CreateRuleRequestRuleActionsInsertHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) SetKey(v string) *CreateRuleRequestRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) SetValue(v string) *CreateRuleRequestRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) SetValueType(v string) *CreateRuleRequestRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

type CreateRuleRequestRuleActionsRedirectConfig struct {
	// The hostname to which requests are redirected. Valid values:
	//
	// *   **${host}** (default): If you set the value to ${host}, you cannot append other characters.
	//
	// *   A custom value. Make sure that the custom value meets the following requirements:
	//
	//     *   The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, and the following special characters: - . \* = ~ \_ + \ ^ ! $ & | ( ) \[ ] ?.
	//     *   The hostname must contain at least one period (.) but cannot start or end with a period (.).
	//     *   The rightmost domain label can contain only letters and wildcards, and cannot contain digits or hyphens (-). The leftmost `domain label` can be an asterisk (\*).
	//     *   The domain labels cannot start or end with a hyphen (-).
	//     *   You can use asterisks (\*) and question marks (?) anywhere in a domain label as wildcard characters.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The redirect type. Valid values: **301**, **302**, **303**, **307**, and **308**.
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The path to which requests are redirected. Valid values:
	//
	// *   Default value: **${path}**. \*\*${host}**, **${protocol}**, and **${port}\*\* are also supported. Each variable can be specified only once. You can specify one or more of the preceding variables in each request. You can also combine them with a custom value.
	//
	// *   A custom value. You must make sure that the custom value meets the following requirements:
	//
	//     *   The value must be 1 to 128 characters in length, and can contain asterisks (\*) and question marks (?) as wildcards. The value is case-sensitive.
	//     *   It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ .+ / & ~ @ :`. It cannot contain the following special characters: `" % # ; ! ( ) [ ] ^ , "`. You can use asterisks (\*) and question marks (?) as wildcard characters.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port to which requests are redirected.
	//
	// *   **${port}** (default): If you set the value to ${port}, you cannot add other characters to the value.
	// *   You can also enter a port number. Valid values: **1 to 63335**.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The redirect protocol. Valid values:
	//
	// *   **${protocol}** (default): If you set the value to ${protocol}, you cannot add other characters to the value.
	// *   **HTTP** or **HTTPS**.
	//
	// > HTTPS listeners support only HTTPS to HTTPS redirects.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The query string of the URL to which requests are redirected.
	//
	// *   Default value: **${query}**. \*\*${host}**, **${protocol}**, and **${port}\*\* are also supported. Each variable can be specified only once. You can specify one or more of the preceding variables in each request. You can also combine them with a custom value.
	//
	// *   A custom value. You must make sure that the custom value meets the following requirements:
	//
	//     *   The value must be 1 to 128 characters in length.
	//     *   It can contain printable characters, except space characters, the special characters `# [ ] { } \ | < > &`, and uppercase letters.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRuleRequestRuleActionsRedirectConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetHost(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetHttpCode(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetPath(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetPort(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetProtocol(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetQuery(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

type CreateRuleRequestRuleActionsRewriteConfig struct {
	// The hostname to which requests are redirected. Valid values:
	//
	// *   **${host}** (default): If you set the value to ${host}, you cannot append other characters.
	//
	// *   If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     *   The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, and the following special characters: - . \* = ~ \_ + \ ^ ! $ & | ( ) \[ ] ?.
	//     *   The hostname must contain at least one period (.) but cannot start or end with a period (.).
	//     *   The rightmost domain label can contain only letters and wildcards, and cannot contain digits or hyphens (-). The leftmost `domain label` can be an asterisk (\*).
	//     *   The domain labels cannot start or end with a hyphen (-). You can use asterisks (\*) and question marks (?) anywhere in a domain label as wildcard characters.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The path to which requests are redirected. Valid values:
	//
	// *   Default value: **${path}**. \*\*${host}**, **${protocol}**, and **${port}\*\* are also supported. Each variable can be specified only once. You can specify one or more of the preceding variables in each request. You can also combine them with a custom value.
	//
	// *   A custom value. You must make sure that the custom value meets the following requirements:
	//
	//     *   The value must be 1 to 128 characters in length, and can contain asterisks (\*) and question marks (?) as wildcards. The value is case-sensitive.
	//     *   It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ .+ / & ~ @ :`. It cannot contain the following special characters: `" % # ; ! ( ) [ ] ^ , "`. You can use asterisks (\*) and question marks (?) as wildcard characters.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The query string of the URL to which requests are redirected.
	//
	// *   Default value: **${query}**. \*\*${host}**, **${protocol}**, and **${port}\*\* are also supported. Each variable can be specified only once. You can specify one or more of the preceding variables in each request. You can also combine them with a custom value.
	//
	// *   A custom value. You must make sure that the custom value meets the following requirements:
	//
	//     *   The value must be 1 to 128 characters in length.
	//     *   It can contain printable characters, except space characters, the special characters `# [ ] { } \ | < > &`, and uppercase letters.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRuleRequestRuleActionsRewriteConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) SetHost(v string) *CreateRuleRequestRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) SetPath(v string) *CreateRuleRequestRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) SetQuery(v string) *CreateRuleRequestRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

type CreateRuleRequestRuleActionsTrafficLimitConfig struct {
	// The QPS of each IP address. Valid values: **1 to 100000**.
	//
	// > If both the **QPS** and **PerIpQps** properties are specified, make sure that the value of the **QPS** property is smaller than the value of the PerIpQps property.
	PerIpQps *int32 `json:"PerIpQps,omitempty" xml:"PerIpQps,omitempty"`
	// The queries per second (QPS). Valid values: **1 to 100000**.
	QPS *int32 `json:"QPS,omitempty" xml:"QPS,omitempty"`
}

func (s CreateRuleRequestRuleActionsTrafficLimitConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsTrafficLimitConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsTrafficLimitConfig) SetPerIpQps(v int32) *CreateRuleRequestRuleActionsTrafficLimitConfig {
	s.PerIpQps = &v
	return s
}

func (s *CreateRuleRequestRuleActionsTrafficLimitConfig) SetQPS(v int32) *CreateRuleRequestRuleActionsTrafficLimitConfig {
	s.QPS = &v
	return s
}

type CreateRuleRequestRuleActionsTrafficMirrorConfig struct {
	// The configuration of the server group to which traffic is mirrored.
	MirrorGroupConfig *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig `json:"MirrorGroupConfig,omitempty" xml:"MirrorGroupConfig,omitempty" type:"Struct"`
	// The type of destination to which network traffic is mirrored. Valid values:
	//
	// *   **ForwardGroupMirror**: a server group
	// *   **SlsMirror**: Log Service
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfig) SetMirrorGroupConfig(v *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) *CreateRuleRequestRuleActionsTrafficMirrorConfig {
	s.MirrorGroupConfig = v
	return s
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfig) SetTargetType(v string) *CreateRuleRequestRuleActionsTrafficMirrorConfig {
	s.TargetType = &v
	return s
}

type CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig struct {
	// The configuration of the server group to which traffic is mirrored.
	ServerGroupTuples []*CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) SetServerGroupTuples(v []*CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples struct {
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type CreateRuleRequestRuleConditions struct {
	// The configurations of the cookies.
	CookieConfig *CreateRuleRequestRuleConditionsCookieConfig `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	// The configuration of the header.
	HeaderConfig *CreateRuleRequestRuleConditionsHeaderConfig `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	// The configurations of the host.
	HostConfig *CreateRuleRequestRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// The configurations of the request methods.
	MethodConfig *CreateRuleRequestRuleConditionsMethodConfig `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	// The configurations of the URLs.
	PathConfig *CreateRuleRequestRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// The configurations of the query strings.
	QueryStringConfig *CreateRuleRequestRuleConditionsQueryStringConfig `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	// The configuration of the source IP-based forwarding rule. This parameter is required and takes effect only when **Type** is set to **SourceIP**.
	SourceIpConfig *CreateRuleRequestRuleConditionsSourceIpConfig `json:"SourceIpConfig,omitempty" xml:"SourceIpConfig,omitempty" type:"Struct"`
	// The type of forwarding rule. Valid values:
	//
	// *   **Host**: Requests are distributed based on hosts.
	// *   **Path**: Requests are distributed based on paths.
	// *   **Header**: Requests are distributed based on HTTP headers.
	// *   **QueryString**: Requests are distributed based on query strings.
	// *   **Method**: Requests are distributed based on request methods.
	// *   **Cookie**: Requests are distributed based on cookies.
	// *   **SourceIp**: Requests are distributed based on source IP addresses.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRuleRequestRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditions) SetCookieConfig(v *CreateRuleRequestRuleConditionsCookieConfig) *CreateRuleRequestRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetHeaderConfig(v *CreateRuleRequestRuleConditionsHeaderConfig) *CreateRuleRequestRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetHostConfig(v *CreateRuleRequestRuleConditionsHostConfig) *CreateRuleRequestRuleConditions {
	s.HostConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetMethodConfig(v *CreateRuleRequestRuleConditionsMethodConfig) *CreateRuleRequestRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetPathConfig(v *CreateRuleRequestRuleConditionsPathConfig) *CreateRuleRequestRuleConditions {
	s.PathConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetQueryStringConfig(v *CreateRuleRequestRuleConditionsQueryStringConfig) *CreateRuleRequestRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetSourceIpConfig(v *CreateRuleRequestRuleConditionsSourceIpConfig) *CreateRuleRequestRuleConditions {
	s.SourceIpConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetType(v string) *CreateRuleRequestRuleConditions {
	s.Type = &v
	return s
}

type CreateRuleRequestRuleConditionsCookieConfig struct {
	// The cookie values.
	Values []*CreateRuleRequestRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsCookieConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsCookieConfig) SetValues(v []*CreateRuleRequestRuleConditionsCookieConfigValues) *CreateRuleRequestRuleConditionsCookieConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsCookieConfigValues struct {
	// The key of the cookie.
	//
	// *   The key must be 1 to 100 characters in length.
	// *   You can use asterisks (\*) and question marks (?) as wildcard characters.
	// *   The key can contain printable characters, except uppercase letters, space characters, and the following special characters: `; # [ ] { } \ | < > &`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the cookie.
	//
	// *   The value must be 1 to 100 characters in length.
	// *   You can use asterisks (\*) and question marks (?) as wildcard characters.
	// *   The value can contain printable characters, except uppercase letters, space characters, and the following special characters: `; # [ ] { } \ | < > &`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRuleRequestRuleConditionsCookieConfigValues) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsCookieConfigValues) SetKey(v string) *CreateRuleRequestRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsCookieConfigValues) SetValue(v string) *CreateRuleRequestRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

type CreateRuleRequestRuleConditionsHeaderConfig struct {
	// The key of the header.
	//
	// *   The key must be 1 to 40 characters in length.
	// *   It can contain lowercase letters, digits, hyphens (-), and underscores (\_).
	// *   Cookie and Host are not supported.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the header.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsHeaderConfig) SetKey(v string) *CreateRuleRequestRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsHeaderConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsHostConfig struct {
	// The hostname.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsHostConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsHostConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsMethodConfig struct {
	// The request methods.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsMethodConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsMethodConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsMethodConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsPathConfig struct {
	// The path.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsPathConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsPathConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsQueryStringConfig struct {
	// The query strings.
	Values []*CreateRuleRequestRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsQueryStringConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfig) SetValues(v []*CreateRuleRequestRuleConditionsQueryStringConfigValues) *CreateRuleRequestRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsQueryStringConfigValues struct {
	// The key of the query string.
	//
	// *   The key must be 1 to 100 characters in length.
	// *   You can use asterisks (\*) and question marks (?) as wildcards. The key can contain printable characters, except uppercase letters, space characters, and the following special characters: `# [ ] { } \ | < > &`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the query string.
	//
	// *   The value must be 1 to 128 characters in length.
	// *   It can contain printable characters, except uppercase letters, space characters, and the following special characters: `# [ ] { } \ | < > &`. You can use asterisks (\*) and question marks (?) as wildcard characters.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRuleRequestRuleConditionsQueryStringConfigValues) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfigValues) SetKey(v string) *CreateRuleRequestRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfigValues) SetValue(v string) *CreateRuleRequestRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

type CreateRuleRequestRuleConditionsSourceIpConfig struct {
	// The configuration of the source IP-based forwarding rule.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsSourceIpConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsSourceIpConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsSourceIpConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsSourceIpConfig {
	s.Values = v
	return s
}

type CreateRuleResponseBody struct {
	// The asynchronous task ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the forwarding rule.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) SetJobId(v string) *CreateRuleResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRuleResponseBody) SetRuleId(v string) *CreateRuleResponseBody {
	s.RuleId = &v
	return s
}

type CreateRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRuleResponse) SetHeaders(v map[string]*string) *CreateRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRuleResponse) SetStatusCode(v int32) *CreateRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRuleResponse) SetBody(v *CreateRuleResponseBody) *CreateRuleResponse {
	s.Body = v
	return s
}

type CreateRulesRequest struct {
	// The format of the response. Valid values:
	//
	// *   **text/plain**
	// *   **text/css**
	// *   **text/html**
	// *   **application/javascript**
	// *   **application/json**
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The HTTP status code in the response. Valid values: **HTTP\_2xx**, **HTTP\_4xx**, and **HTTP\_5xx**. **x** must be a digit.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The content of the custom response. The content can be up to 1 KB in size and can contain only ASCII characters.
	ListenerId *string                    `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Rules      []*CreateRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreateRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateRulesRequest) SetClientToken(v string) *CreateRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRulesRequest) SetDryRun(v bool) *CreateRulesRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRulesRequest) SetListenerId(v string) *CreateRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateRulesRequest) SetRules(v []*CreateRulesRequestRules) *CreateRulesRequest {
	s.Rules = v
	return s
}

type CreateRulesRequestRules struct {
	// The ID of the forwarding rule.
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The server group to which requests are distributed.
	Priority       *int32                                   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RuleActions    []*CreateRulesRequestRulesRuleActions    `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	RuleConditions []*CreateRulesRequestRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The list of forwarding rules.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateRulesRequestRules) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRules) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRules) SetDirection(v string) *CreateRulesRequestRules {
	s.Direction = &v
	return s
}

func (s *CreateRulesRequestRules) SetPriority(v int32) *CreateRulesRequestRules {
	s.Priority = &v
	return s
}

func (s *CreateRulesRequestRules) SetRuleActions(v []*CreateRulesRequestRulesRuleActions) *CreateRulesRequestRules {
	s.RuleActions = v
	return s
}

func (s *CreateRulesRequestRules) SetRuleConditions(v []*CreateRulesRequestRulesRuleConditions) *CreateRulesRequestRules {
	s.RuleConditions = v
	return s
}

func (s *CreateRulesRequestRules) SetRuleName(v string) *CreateRulesRequestRules {
	s.RuleName = &v
	return s
}

type CreateRulesRequestRulesRuleActions struct {
	CorsConfig          *CreateRulesRequestRulesRuleActionsCorsConfig          `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	FixedResponseConfig *CreateRulesRequestRulesRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	ForwardGroupConfig  *CreateRulesRequestRulesRuleActionsForwardGroupConfig  `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	InsertHeaderConfig  *CreateRulesRequestRulesRuleActionsInsertHeaderConfig  `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// The port to which requests are redirected.
	//
	// *   **${port}** (default): If you set the value to ${port}, you cannot append other characters.
	// *   You can also enter a port number. Valid values: **1 to 63335**.
	Order               *int32                                                 `json:"Order,omitempty" xml:"Order,omitempty"`
	RedirectConfig      *CreateRulesRequestRulesRuleActionsRedirectConfig      `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	RewriteConfig       *CreateRulesRequestRulesRuleActionsRewriteConfig       `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	TrafficLimitConfig  *CreateRulesRequestRulesRuleActionsTrafficLimitConfig  `json:"TrafficLimitConfig,omitempty" xml:"TrafficLimitConfig,omitempty" type:"Struct"`
	TrafficMirrorConfig *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig `json:"TrafficMirrorConfig,omitempty" xml:"TrafficMirrorConfig,omitempty" type:"Struct"`
	// The ID of the vServer group.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRulesRequestRulesRuleActions) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActions) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActions) SetCorsConfig(v *CreateRulesRequestRulesRuleActionsCorsConfig) *CreateRulesRequestRulesRuleActions {
	s.CorsConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetFixedResponseConfig(v *CreateRulesRequestRulesRuleActionsFixedResponseConfig) *CreateRulesRequestRulesRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetForwardGroupConfig(v *CreateRulesRequestRulesRuleActionsForwardGroupConfig) *CreateRulesRequestRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetInsertHeaderConfig(v *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) *CreateRulesRequestRulesRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetOrder(v int32) *CreateRulesRequestRulesRuleActions {
	s.Order = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetRedirectConfig(v *CreateRulesRequestRulesRuleActionsRedirectConfig) *CreateRulesRequestRulesRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetRewriteConfig(v *CreateRulesRequestRulesRuleActionsRewriteConfig) *CreateRulesRequestRulesRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetTrafficLimitConfig(v *CreateRulesRequestRulesRuleActionsTrafficLimitConfig) *CreateRulesRequestRulesRuleActions {
	s.TrafficLimitConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetTrafficMirrorConfig(v *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) *CreateRulesRequestRulesRuleActions {
	s.TrafficMirrorConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetType(v string) *CreateRulesRequestRulesRuleActions {
	s.Type = &v
	return s
}

type CreateRulesRequestRulesRuleActionsCorsConfig struct {
	// The key of the header.
	//
	// *   The key must be 1 to 40 characters in length.
	// *   It can contain letters, digits, hyphens (-), and underscores (\_).
	// *   You cannot set Cookie or Host.
	AllowCredentials *string   `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	AllowHeaders     []*string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty" type:"Repeated"`
	AllowMethods     []*string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty" type:"Repeated"`
	AllowOrigin      []*string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty" type:"Repeated"`
	ExposeHeaders    []*string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty" type:"Repeated"`
	// The value of the header. The header values within a forwarding rule must be unique.
	//
	// *   The value must be 1 to 128 characters in length.
	// *   It can contain printable characters whose ASCII values are `greater than or equal to 32 and lower than 127`. You can use asterisks (\*) and question marks (?) as wildcard characters.
	// *   The value cannot start or end with a space character.
	MaxAge *int64 `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsCorsConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsCorsConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetAllowCredentials(v string) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetAllowHeaders(v []*string) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.AllowHeaders = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetAllowMethods(v []*string) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.AllowMethods = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetAllowOrigin(v []*string) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.AllowOrigin = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetExposeHeaders(v []*string) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.ExposeHeaders = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetMaxAge(v int64) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.MaxAge = &v
	return s
}

type CreateRulesRequestRulesRuleActionsFixedResponseConfig struct {
	// The weight of the server group. A larger value indicates a higher weight. A server group with a higher weight receives more requests. Valid values: **1 to 100**. Default value: **100**.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Specifies whether to enable session persistence. Valid values:
	//
	// *   **true**: enables session persistence.
	// *   **false** (default): disables session persistence.
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The timeout period of sessions. Unit: seconds. Valid values: **1 to 86400**.
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsFixedResponseConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) SetContent(v string) *CreateRulesRequestRulesRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) SetContentType(v string) *CreateRulesRequestRulesRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) SetHttpCode(v string) *CreateRulesRequestRulesRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

type CreateRulesRequestRulesRuleActionsForwardGroupConfig struct {
	ServerGroupStickySession *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession `json:"ServerGroupStickySession,omitempty" xml:"ServerGroupStickySession,omitempty" type:"Struct"`
	ServerGroupTuples        []*CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples      `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfig) SetServerGroupStickySession(v *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) *CreateRulesRequestRulesRuleActionsForwardGroupConfig {
	s.ServerGroupStickySession = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) *CreateRulesRequestRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession struct {
	// The type of header. Valid values:
	//
	// *   **UserDefined**: a custom header.
	// *   **ReferenceHeader**: a header that is referenced from one of the request headers.
	// *   **SystemDefined**: a header predefined by the system.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The priority of the action within the forwarding rule. Valid values: **1 to 50000**. A lower value indicates a higher priority. The actions of a forwarding rule are applied in descending order of priority. This parameter is required. The priority of each action within a forwarding rule must be unique. You can specify priorities for at most 20 actions.
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) SetEnabled(v bool) *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Enabled = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) SetTimeout(v int32) *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Timeout = &v
	return s
}

type CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	// The name of the header to insert. The name must be 1 to 40 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The header names specified by **InsertHeaderConfig** must be unique.
	//
	// >  You cannot set the name of the header to any of the following values (case-insensitive): `slb-id`, `slb-ip`, `x-forwarded-for`, `x-forwarded-proto`, `x-forwarded-eip`, `x-forwarded-port`, `x-forwarded-client-srcport`, `connection`, `upgrade`, `content-length`, `transfer-encoding`, `keep-alive`, `te`, `host`, `cookie`, `remoteip`, and `authority`.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The value of the header to insert.
	//
	// *   If **ValueType** is set to **SystemDefined**, you can set one of the following header values:
	//
	//     *   **ClientSrcPort**: the client port.
	//     *   **ClientSrcIp**: the client IP address.
	//     *   **Protocol**: the request protocol (HTTP or HTTPS).
	//     *   **SLBId**: the ID of the ALB instance.
	//     *   **SLBPort**: the listening port.
	//
	// *   If **ValueType** is set to **UserDefined**, you can specify a custom header value. The header value must be 1 to 128 characters in length and can contain printable characters whose ASCII character values are `greater than or equal to 32 and lower than 127`. You can use asterisks (\*) and question marks (?) as wildcards. The value cannot start or end with a space character.
	//
	// *   If **ValueType** is set to **ReferenceHeader**, you can reference one of the request headers. The header value must be 1 to 128 characters in length, and can contain lowercase letters, digits, underscores (\_), and hyphens (-).
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) SetWeight(v int32) *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

type CreateRulesRequestRulesRuleActionsInsertHeaderConfig struct {
	// The hostname to which requests are distributed. Valid values:
	//
	// *   **${host}** (default): If you set the value to ${host}, you cannot append other characters.
	//
	// *   If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     *   The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), and periods (.). You can use asterisks (\*) and question marks (?) as wildcard characters.
	//     *   The hostname must contain at least one period (.) but cannot start or end with a period (.).
	//     *   The rightmost domain label can contain only letters and wildcard characters. It cannot contain digits or hyphens (-).
	//     *   The domain labels cannot start or end with a hyphen (-).
	//     *   You can use asterisks (\*) and question marks (?) as wildcards anywhere in a domain label.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The HTTP status code that indicates the redirect type. Valid values: **301**, **302**, **303**, **307**, and **308**.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The path to which requests are redirected. Valid values:
	//
	// *   Default value: **${path}**. **${host}**, **${protocol}**, and **${port}** are also supported. Each variable cannot be specified more than once. You can specify one or more of the preceding variables in each request. You can also combine them with the following characters.
	//
	// *   If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     *   The value is 1 to 128 characters in length.
	//     *   It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ .+ / & ~ @ :`. It cannot contain the following special characters: `" % # ; ! ( ) [ ]^ , "`. You can use asterisks (\*) and question marks (?) as wildcards.
	//     *   The value is case-sensitive.
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsInsertHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) SetKey(v string) *CreateRulesRequestRulesRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) SetValue(v string) *CreateRulesRequestRulesRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) SetValueType(v string) *CreateRulesRequestRulesRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

type CreateRulesRequestRulesRuleActionsRedirectConfig struct {
	// The redirect protocol. Valid values:
	//
	// *   **${protocol}** (default): If you set the value to ${protocol}, you cannot append other characters.
	// *   You can set the protocol to **HTTP** or **HTTPS**.
	//
	// >  HTTPS listeners do not support HTTPS-to-HTTP redirects.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The query string of the URL to which requests are redirected.
	//
	// *   Default value: **${query}**. **${host}**, **${protocol}**, and **${port}** are also supported. Each variable cannot be specified more than once. You can specify one or more of the preceding variables in each request. You can also combine them with the following characters.
	//
	// *   If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     *   The value is 1 to 128 characters in length.
	//     *   It can contain printable characters, except space characters, the special characters `# [ ] { } \ | < > &`, and uppercase letters.
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The hostname to which requests are redirected. Valid values:
	//
	// *   **${host}** (default): If you set the value to ${host}, you cannot append other characters.
	//
	// *   If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     *   The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), and periods (.). You can use asterisks (\*) and question marks (?) as wildcard characters.
	//     *   The hostname must contain at least one period (.) but cannot start or end with a period (.).
	//     *   The rightmost domain label can contain only letters and wildcard characters. It cannot contain digits or hyphens (-).
	//     *   The domain labels cannot start or end with a hyphen (-). You can use an asterisk (\*) and question mark (?) as wildcards anywhere in a domain label.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The path to which requests are redirected. Valid values:
	//
	// *   Default value: **${path}**. **${host}**, **${protocol}**, and **${port}** are also supported. Each variable cannot be specified more than once. You can specify one or more of the preceding variables in each request. You can also combine them with the following characters.
	//
	// *   If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     *   The value is 1 to 128 characters in length.
	//     *   It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ .+ / & ~ @ :`. It cannot contain the following special characters: `" % # ; ! ( ) [ ]^ , "`. You can use asterisks (\*) and question marks (?) as wildcards.
	//     *   The value is case-sensitive.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The query string of the URL to which requests are redirected.
	//
	// *   Default value: **${query}**. **${host}**, **${protocol}**, and **${port}** are also supported. Each variable cannot be specified more than once. You can specify one or more of the preceding variables in each request. You can also combine them with the following characters.
	//
	// *   If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     *   The value is 1 to 128 characters in length.
	//     *   It can contain printable characters, except space characters, the special characters `# [ ] { } \ | < > &`, and uppercase letters.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The action type. You can specify at most 11 types of action. Valid values:
	//
	// *   **ForwardGroup**: forwards a request to multiple vServer groups.
	// *   **Redirect**: redirects a request.
	// *   **FixedResponse**: returns a custom response.
	// *   **Rewrite**: rewrites a request.
	// *   **InsertHeader**: inserts a header.
	// *   **RemoveHeaderConfig**: deletes a header.
	// *   **TrafficLimitConfig**: throttles network traffic.
	// *   **TrafficMirrorConfig**: mirrors network traffic.
	// *   **CORS**: enables cross-origin resource sharing (CORS).
	//
	// You can specify the last action and the actions that you want to perform before the last action:
	//
	// *   **FinalType**: the last action to be performed in a forwarding rule. Each forwarding rule can contain only one FinalType action. You can specify a **ForwardGroup**, **Redirect**, or **FixedResponse** action as the FinalType action.
	// *   **ExtType**: the action to be performed before the FinalType action. A forwarding rule can contain one or more ExtType actions. To specify this parameter, you must also specify FinalType. You can specify multiple **InsertHeader** actions or one **Rewrite** action.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsRedirectConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetHost(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetHttpCode(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetPath(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetPort(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetProtocol(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetQuery(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

type CreateRulesRequestRulesRuleActionsRewriteConfig struct {
	// Queries per second (QPS). Valid values: **1 to 100000**.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The QPS of each IP address. Valid values: **1 to 100000**.
	//
	// >  If **QPS** and PerIpQps are configured at the same time, the value of the **PerIpQps** parameter must be smaller than that of the **QPS** parameter.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The type of destination to which network traffic is mirrored. Valid values:
	//
	// *   **ForwardGroupMirror**: a server group.
	// *   **SlsMirror**: Log Service.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsRewriteConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) SetHost(v string) *CreateRulesRequestRulesRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) SetPath(v string) *CreateRulesRequestRulesRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) SetQuery(v string) *CreateRulesRequestRulesRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

type CreateRulesRequestRulesRuleActionsTrafficLimitConfig struct {
	// The allowed HTTP methods for CORS requests. Valid values:
	//
	// *   **GET**
	// *   **POST**
	// *   **PUT**
	// *   **DELETE**
	// *   **HEAD**
	// *   **OPTIONS**
	// *   **PATCH**
	PerIpQps *int32 `json:"PerIpQps,omitempty" xml:"PerIpQps,omitempty"`
	// The origin site that is allowed to access. You can specify an asterisk (`*`) or one or more values. The value cannot be an asterisk (`*`).
	//
	// *   The value must start with `http://` or `https://` and include a valid domain name or top-level wildcard domain name, such as `*.test.abc.example.com`.
	// *   You can choose to include a port number from **1** to **65535** in each value based on your business requirement.
	QPS *int32 `json:"QPS,omitempty" xml:"QPS,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsTrafficLimitConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsTrafficLimitConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsTrafficLimitConfig) SetPerIpQps(v int32) *CreateRulesRequestRulesRuleActionsTrafficLimitConfig {
	s.PerIpQps = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsTrafficLimitConfig) SetQPS(v int32) *CreateRulesRequestRulesRuleActionsTrafficLimitConfig {
	s.QPS = &v
	return s
}

type CreateRulesRequestRulesRuleActionsTrafficMirrorConfig struct {
	MirrorGroupConfig *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig `json:"MirrorGroupConfig,omitempty" xml:"MirrorGroupConfig,omitempty" type:"Struct"`
	// The allowed headers for CORS requests. You can specify an asterisk (`*`) or one or more values. Separate multiple values with commas (,). The value must be 1 to 32 characters in length, and can contain letters and digits. The value cannot start or end with an underscore (\_) or hyphen (-).
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) SetMirrorGroupConfig(v *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig {
	s.MirrorGroupConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) SetTargetType(v string) *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig {
	s.TargetType = &v
	return s
}

type CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig struct {
	ServerGroupTuples []*CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) SetServerGroupTuples(v []*CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples struct {
	// The headers that are allowed to expose. You can specify an asterisk (`*`) or one or more values. Separate multiple values with commas (,). The value must be 1 to 32 characters in length, and can contain letters and digits. The value cannot start or end with an underscore (\_) or hyphen (-).
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type CreateRulesRequestRulesRuleConditions struct {
	CookieConfig         *CreateRulesRequestRulesRuleConditionsCookieConfig         `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	HeaderConfig         *CreateRulesRequestRulesRuleConditionsHeaderConfig         `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	HostConfig           *CreateRulesRequestRulesRuleConditionsHostConfig           `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	MethodConfig         *CreateRulesRequestRulesRuleConditionsMethodConfig         `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	PathConfig           *CreateRulesRequestRulesRuleConditionsPathConfig           `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	QueryStringConfig    *CreateRulesRequestRulesRuleConditionsQueryStringConfig    `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	ResponseHeaderConfig *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig `json:"ResponseHeaderConfig,omitempty" xml:"ResponseHeaderConfig,omitempty" type:"Struct"`
	SourceIpConfig       *CreateRulesRequestRulesRuleConditionsSourceIpConfig       `json:"SourceIpConfig,omitempty" xml:"SourceIpConfig,omitempty" type:"Struct"`
	// The ID of the asynchronous task.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRulesRequestRulesRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditions) SetCookieConfig(v *CreateRulesRequestRulesRuleConditionsCookieConfig) *CreateRulesRequestRulesRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetHeaderConfig(v *CreateRulesRequestRulesRuleConditionsHeaderConfig) *CreateRulesRequestRulesRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetHostConfig(v *CreateRulesRequestRulesRuleConditionsHostConfig) *CreateRulesRequestRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetMethodConfig(v *CreateRulesRequestRulesRuleConditionsMethodConfig) *CreateRulesRequestRulesRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetPathConfig(v *CreateRulesRequestRulesRuleConditionsPathConfig) *CreateRulesRequestRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetQueryStringConfig(v *CreateRulesRequestRulesRuleConditionsQueryStringConfig) *CreateRulesRequestRulesRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetResponseHeaderConfig(v *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) *CreateRulesRequestRulesRuleConditions {
	s.ResponseHeaderConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetSourceIpConfig(v *CreateRulesRequestRulesRuleConditionsSourceIpConfig) *CreateRulesRequestRulesRuleConditions {
	s.SourceIpConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetType(v string) *CreateRulesRequestRulesRuleConditions {
	s.Type = &v
	return s
}

type CreateRulesRequestRulesRuleConditionsCookieConfig struct {
	Values []*CreateRulesRequestRulesRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfig) SetValues(v []*CreateRulesRequestRulesRuleConditionsCookieConfigValues) *CreateRulesRequestRulesRuleConditionsCookieConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsCookieConfigValues struct {
	// The hostname. A forwarding rule can contain only one unique hostname.
	//
	// *   The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), periods (.), asterisks (\*), and question marks (?).
	// *   The hostname must contain at least one period (.) but cannot start or end with a period (.).
	// *   The rightmost domain label can contain only letters and wildcard characters. It cannot contain digits or hyphens (-).
	// *   The domain labels do not start or end with hyphens (-). You can use an asterisk (\*) and question mark (?) as wildcards anywhere in a domain label.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The request methods. Valid values: **HEAD**, **GET**, **POST**, **OPTIONS**, **PUT**, **PATCH**, and **DELETE**.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfigValues) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfigValues) SetKey(v string) *CreateRulesRequestRulesRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfigValues) SetValue(v string) *CreateRulesRequestRulesRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

type CreateRulesRequestRulesRuleConditionsHeaderConfig struct {
	// The path to which requests are forwarded. Limits:
	//
	// *   The path must be 1 to 128 characters in length.
	// *   It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ .+ / & ~ @ :`. It cannot contain the following special characters: `" % # ; ! ( ) [ ]^ , "`. You can use asterisks (\*) and question marks (?) as wildcards.
	// *   The value is case-sensitive.
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsHeaderConfig) SetKey(v string) *CreateRulesRequestRulesRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsHeaderConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsHostConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsHostConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsMethodConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsMethodConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsMethodConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsMethodConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsPathConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsPathConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsQueryStringConfig struct {
	Values []*CreateRulesRequestRulesRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfig) SetValues(v []*CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) *CreateRulesRequestRulesRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsQueryStringConfigValues struct {
	// The type of forwarding rule. You can specify at most seven types. Valid values:
	//
	// *   **Host**: Requests are forwarded based on hosts.
	// *   **Path**: Requests are forwarded based on paths.
	// *   **Header**: Requests are forwarded based on HTTP headers.
	// *   **QueryString**: Requests are forwarded based on query strings.
	// *   **Method**: Requests are forwarded based on request methods.
	// *   **Cookie**: Requests are forwarded based on cookies.
	// *   **SourceIp**: Requests are forwarded based on source IP addresses.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The IP addresses or CIDR blocks.
	//
	// You can specify at most five values for **SourceIp**.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) SetKey(v string) *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) SetValue(v string) *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

type CreateRulesRequestRulesRuleConditionsResponseHeaderConfig struct {
	// The name of the forwarding rule. You can name at most 20 forwarding rules.
	//
	// *   The name must be 2 to 128 characters in length.
	// *   It can contain letters, digits, periods (.), underscores (\_), and hyphens (-). It must start with a letter.
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) SetKey(v string) *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsSourceIpConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsSourceIpConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsSourceIpConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsSourceIpConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsSourceIpConfig {
	s.Values = v
	return s
}

type CreateRulesResponseBody struct {
	// The priority of the forwarding rule. Valid values: **1 to 10000**. A smaller value indicates a higher priority.
	//
	// >  The priority of each forwarding rule added to a listener is unique.
	JobId     *string                           `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleIds   []*CreateRulesResponseBodyRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s CreateRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBody) SetJobId(v string) *CreateRulesResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRulesResponseBody) SetRequestId(v string) *CreateRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRulesResponseBody) SetRuleIds(v []*CreateRulesResponseBodyRuleIds) *CreateRulesResponseBody {
	s.RuleIds = v
	return s
}

type CreateRulesResponseBodyRuleIds struct {
	Priority *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RuleId   *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateRulesResponseBodyRuleIds) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesResponseBodyRuleIds) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBodyRuleIds) SetPriority(v int32) *CreateRulesResponseBodyRuleIds {
	s.Priority = &v
	return s
}

func (s *CreateRulesResponseBodyRuleIds) SetRuleId(v string) *CreateRulesResponseBodyRuleIds {
	s.RuleId = &v
	return s
}

type CreateRulesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateRulesResponse) SetHeaders(v map[string]*string) *CreateRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateRulesResponse) SetStatusCode(v int32) *CreateRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRulesResponse) SetBody(v *CreateRulesResponseBody) *CreateRulesResponse {
	s.Body = v
	return s
}

type CreateSecurityPolicyRequest struct {
	// The supported cipher suites, which are determined by the TLS protocol version.****
	//
	// The specified cipher suites must be supported by at least one TLS protocol version that you specify.****
	//
	// >  For example, if you set the TLSVersions parameter to TLSv1.3, you must specify cipher suites that are supported by TLS 1.3.
	//
	// *   TLS 1.0 and TLS 1.1 support the following cipher suites:
	//
	//     *   ECDHE-ECDSA-AES128-SHA
	//     *   ECDHE-ECDSA-AES256-SHA
	//     *   ECDHE-RSA-AES128-SHA
	//     *   ECDHE-RSA-AES256-SHA
	//     *   AES128-SHA
	//     *   AES256-SHA
	//     *   DES-CBC3-SHA
	//
	// *   TLS 1.2 supports the following cipher suites:
	//
	//     *   ECDHE-ECDSA-AES128-SHA
	//     *   ECDHE-ECDSA-AES256-SHA
	//     *   ECDHE-RSA-AES128-SHA
	//     *   ECDHE-RSA-AES256-SHA
	//     *   AES128-SHA
	//     *   AES256-SHA
	//     *   DES-CBC3-SHA
	//     *   ECDHE-ECDSA-AES128-GCM-SHA256
	//     *   ECDHE-ECDSA-AES256-GCM-SHA384
	//     *   ECDHE-ECDSA-AES128-SHA256
	//     *   ECDHE-ECDSA-AES256-SHA384
	//     *   ECDHE-RSA-AES128-GCM-SHA256
	//     *   ECDHE-RSA-AES256-GCM-SHA384
	//     *   ECDHE-RSA-AES128-SHA256
	//     *   ECDHE-RSA-AES256-SHA384
	//     *   AES128-GCM-SHA256
	//     *   AES256-GCM-SHA384
	//     *   AES128-SHA256
	//     *   AES256-SHA256
	//
	// *   TLS 1.3 supports the following cipher suites:
	//
	//     *   TLS_AES\_128\_GCM_SHA256
	//     *   TLS_AES\_256\_GCM_SHA384
	//     *   TLS_CHACHA20\_POLY1305\_SHA256
	//     *   TLS_AES\_128\_CCM_SHA256
	//     *   TLS_AES\_128\_CCM\_8\_SHA256
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The ID of each request is unique.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: only prechecks the request and does not perform the requested operation. The system checks the required parameters, request format, and service limits. If the request fails the precheck, an error code is returned based on the cause of the failure. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false**: prechecks the request and performs the requested operation. After the request passes the precheck, an HTTP 2xx status code is returned and the system performs the operation. This is the default value.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the security policy.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	// The supported versions of the Transport Layer Security (TLS) protocol. Valid values: **TLSv1.0**, **TLSv1.1**, **TLSv1.2**, and **TLSv1.3**.
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s CreateSecurityPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyRequest) SetCiphers(v []*string) *CreateSecurityPolicyRequest {
	s.Ciphers = v
	return s
}

func (s *CreateSecurityPolicyRequest) SetClientToken(v string) *CreateSecurityPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetDryRun(v bool) *CreateSecurityPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetResourceGroupId(v string) *CreateSecurityPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetSecurityPolicyName(v string) *CreateSecurityPolicyRequest {
	s.SecurityPolicyName = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetTLSVersions(v []*string) *CreateSecurityPolicyRequest {
	s.TLSVersions = v
	return s
}

type CreateSecurityPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the security policy.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s CreateSecurityPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyResponseBody) SetRequestId(v string) *CreateSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetSecurityPolicyId(v string) *CreateSecurityPolicyResponseBody {
	s.SecurityPolicyId = &v
	return s
}

type CreateSecurityPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSecurityPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyResponse) SetHeaders(v map[string]*string) *CreateSecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityPolicyResponse) SetStatusCode(v int32) *CreateSecurityPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityPolicyResponse) SetBody(v *CreateSecurityPolicyResponseBody) *CreateSecurityPolicyResponse {
	s.Body = v
	return s
}

type CreateServerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configuration of health checks.
	HealthCheckConfig *CreateServerGroupRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// The backend protocol. Valid values:
	//
	// *   **HTTP** (default): The server group can be associated with HTTPS, HTTP, and QUIC listeners.
	// *   **HTTPS**: The server group can be associated with HTTPS listeners.
	// *   **gRPC**: The server group can be associated with HTTPS and QUIC listeners.
	//
	// > If the **ServerGroupType** parameter is set to **Fc**, you can set Protocol only to **HTTP**.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// *   **Wrr** (default): The weighted round-robin algorithm is used. Backend servers that have higher weights receive more requests than those that have lower weights.
	// *   **Wlc**: The weighted least connections algorithm is used. Requests are distributed based on the weights and the number of connections to backend servers. If two backend servers have the same weight, the backend server that has fewer connections is expected to receive more requests.
	// *   **Sch**: The consistent hashing algorithm is used. Requests from the same source IP address are distributed to the same backend server.
	//
	// > This parameter takes effect when the **ServerGroupType** parameter is set to **Instance** or **Ip**.
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The type of server group. Valid values:
	//
	// *   **Instance** (default): allows you to add servers by specifying **Ecs**, **Ens**, or **Eci**.
	// *   **Ip**: allows you to add servers by specifying IP addresses.
	// *   **Fc**: allows you to add servers by specifying functions of Function Compute.
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// This parameter is available only if the ALB Ingress controller is used. In this case, set this parameter to the name of the `Kubernetes Service` that is associated with the server group.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The configuration of session persistence.
	//
	// > This parameter takes effect when the **ServerGroupType** parameter is set to **Instance** or **Ip**.
	StickySessionConfig *CreateServerGroupRequestStickySessionConfig `json:"StickySessionConfig,omitempty" xml:"StickySessionConfig,omitempty" type:"Struct"`
	// The setting of consistent hashing based on URLs.
	UchConfig *CreateServerGroupRequestUchConfig `json:"UchConfig,omitempty" xml:"UchConfig,omitempty" type:"Struct"`
	// The ID of the virtual private cloud (VPC). You can add only backend servers that are deployed in the specified VPC to the server group.
	//
	// > This parameter takes effect when the **ServerGroupType** parameter is set to **Instance** or **Ip**.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequest) SetClientToken(v string) *CreateServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServerGroupRequest) SetDryRun(v bool) *CreateServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServerGroupRequest) SetHealthCheckConfig(v *CreateServerGroupRequestHealthCheckConfig) *CreateServerGroupRequest {
	s.HealthCheckConfig = v
	return s
}

func (s *CreateServerGroupRequest) SetProtocol(v string) *CreateServerGroupRequest {
	s.Protocol = &v
	return s
}

func (s *CreateServerGroupRequest) SetResourceGroupId(v string) *CreateServerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServerGroupRequest) SetScheduler(v string) *CreateServerGroupRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateServerGroupRequest) SetServerGroupName(v string) *CreateServerGroupRequest {
	s.ServerGroupName = &v
	return s
}

func (s *CreateServerGroupRequest) SetServerGroupType(v string) *CreateServerGroupRequest {
	s.ServerGroupType = &v
	return s
}

func (s *CreateServerGroupRequest) SetServiceName(v string) *CreateServerGroupRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateServerGroupRequest) SetStickySessionConfig(v *CreateServerGroupRequestStickySessionConfig) *CreateServerGroupRequest {
	s.StickySessionConfig = v
	return s
}

func (s *CreateServerGroupRequest) SetUchConfig(v *CreateServerGroupRequestUchConfig) *CreateServerGroupRequest {
	s.UchConfig = v
	return s
}

func (s *CreateServerGroupRequest) SetVpcId(v string) *CreateServerGroupRequest {
	s.VpcId = &v
	return s
}

type CreateServerGroupRequestHealthCheckConfig struct {
	// The HTTP status codes that are used to determine whether the backend server passes the health check.
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The port that you want to use for health checks on backend servers.
	//
	// Valid values: **0** to **65535**.
	//
	// Default value: **0**. If you set the value to 0, the port of a backend server is used for health checks.
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > If the **ServerGroupType** parameter is set to **Instance** or **Ip**, the health check feature is enabled by default. If the **ServerGroupType** parameter is set to **Fc**, the health check feature is disabled by default.
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The domain name that is used for health checks. The domain name must meet the following requirements:
	//
	// *   The domain name must be 1 to 80 characters in length.
	// *   The domain name can contain lowercase letters, digits, hyphens (-), and periods (.).
	// *   It must contain at least one period (.) but cannot start or end with a period (.).
	// *   The rightmost domain label of the domain name can contain only letters, and cannot contain digits or hyphens (-).
	// *   The domain name cannot start or end with a hyphen (-).
	//
	// > This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP**.
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The HTTP version. Valid values: **HTTP1.0** and **HTTP1.1**. Default value: HTTP1.1.
	//
	// > This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP**.
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval at which health checks are performed. Unit: seconds.
	//
	// Valid values: **1** to **50**.
	//
	// Default value: **2**.
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP method that is used for health checks. Valid values:
	//
	// *   **GET**: If the length of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	// *   **POST**: gRPC health checks automatically use the POST method.
	// *   **HEAD**: By default, HTTP health checks use the HEAD method.
	//
	// > This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The path that is used for health checks.
	//
	// The path must be 1 to 80 characters in length and can contain only letters, digits, and the following special characters: `- / . % ? # & =`. It can also contain the following extended characters: `_ ; ~ ! ( ) * [ ] @ $ ^ : \" , +`. The URL must start with a forward slash (/).
	//
	// > This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP**.
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// *   **HTTP**: To perform HTTP health checks, Application Load Balancer (ALB) sends HEAD or GET requests to a backend server to check whether the backend server is healthy.
	// *   **HTTPS**: To perform HTTPS health checks, ALB sends SYN packets to a backend server to check whether the port of the backend server is available to receive requests.
	// *   **gRPC**: To perform gRPC health checks, ALB sends POST or GET requests to a backend server to check whether the backend server is healthy.
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// Specify the timeout period of a health check response. If a backend server, such as an Elastic Compute Service (ECS) instance, does not return a health check response within the specified timeout period, the server fails the health check. Unit: seconds.
	//
	// Valid values: **1** to **300**.
	//
	// Default value: **5**.
	//
	// > If the value of the **HealthCheckTimeout** parameter is smaller than that of the **HealthCheckInterval** parameter, the timeout period specified by the **HealthCheckTimeout** parameter is ignored and the value of the **HealthCheckInterval** parameter is used as the timeout period.
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail** to **success**.
	//
	// Valid values: **2** to **10**.
	//
	// Default value: **3**.
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success** to **fail**.
	//
	// Valid values: **2** to **10**.
	//
	// Default value: **3**.
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s CreateServerGroupRequestHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupRequestHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckCodes(v []*string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckCodes = v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckConnectPort(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckEnabled(v bool) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckHost(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckHost = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckHttpVersion(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckMethod(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckPath(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckProtocol(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckTimeout(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthyThreshold(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

type CreateServerGroupRequestStickySessionConfig struct {
	// The cookie to be configured on the server.
	//
	// The cookie must be 1 to 200 characters in length and can contain only ASCII characters and digits. It cannot contain commas (,), semicolons (;), or space characters. It cannot start with a dollar sign ($).
	//
	// > This parameter takes effect when the **StickySessionEnabled** parameter is set to **true** and the **StickySessionType** parameter is set to **Server**.
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of a cookie. Unit: seconds.
	//
	// Valid values: **1** to **86400**.
	//
	// Default value: **1000**.
	//
	// > This parameter takes effect only when the **StickySessionEnabled** parameter is set to **true** and the **StickySessionType** parameter is set to **Insert**.
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// Specifies whether to enable session persistence. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	//
	// > This parameter takes effect when the **ServerGroupType** parameter is set to **Instance** or **Ip**.
	StickySessionEnabled *bool `json:"StickySessionEnabled,omitempty" xml:"StickySessionEnabled,omitempty"`
	// The method that is used to handle a cookie. Valid values:
	//
	// *   **Insert** (default): inserts a cookie.
	//
	// ALB inserts a session cookie (SERVERID) into the first HTTP or HTTPS response that is sent to a client. Subsequent requests to ALB carry this cookie, and ALB determines the destination servers of the requests based on the cookies.
	//
	// *   **Server**: rewrites a cookie.
	//
	// When ALB detects a user-defined cookie, it overwrites the original cookie with the user-defined cookie. Subsequent requests to ALB carry this user-defined cookie, and ALB determines the destination servers of the requests based on the cookies.
	//
	// > This parameter takes effect when the **StickySessionEnabled** parameter is set to **true**.
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
}

func (s CreateServerGroupRequestStickySessionConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupRequestStickySessionConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestStickySessionConfig) SetCookie(v string) *CreateServerGroupRequestStickySessionConfig {
	s.Cookie = &v
	return s
}

func (s *CreateServerGroupRequestStickySessionConfig) SetCookieTimeout(v int32) *CreateServerGroupRequestStickySessionConfig {
	s.CookieTimeout = &v
	return s
}

func (s *CreateServerGroupRequestStickySessionConfig) SetStickySessionEnabled(v bool) *CreateServerGroupRequestStickySessionConfig {
	s.StickySessionEnabled = &v
	return s
}

func (s *CreateServerGroupRequestStickySessionConfig) SetStickySessionType(v string) *CreateServerGroupRequestStickySessionConfig {
	s.StickySessionType = &v
	return s
}

type CreateServerGroupRequestUchConfig struct {
	// Type
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The setting of consistent hashing.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServerGroupRequestUchConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupRequestUchConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestUchConfig) SetType(v string) *CreateServerGroupRequestUchConfig {
	s.Type = &v
	return s
}

func (s *CreateServerGroupRequestUchConfig) SetValue(v string) *CreateServerGroupRequestUchConfig {
	s.Value = &v
	return s
}

type CreateServerGroupResponseBody struct {
	// The ID of the asynchronous job.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerGroupResponseBody) SetJobId(v string) *CreateServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetRequestId(v string) *CreateServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetServerGroupId(v string) *CreateServerGroupResponseBody {
	s.ServerGroupId = &v
	return s
}

type CreateServerGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateServerGroupResponse) SetHeaders(v map[string]*string) *CreateServerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateServerGroupResponse) SetStatusCode(v int32) *CreateServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServerGroupResponse) SetBody(v *CreateServerGroupResponseBody) *CreateServerGroupResponse {
	s.Body = v
	return s
}

type DeleteAScriptsRequest struct {
	AScriptIds  []*string `json:"AScriptIds,omitempty" xml:"AScriptIds,omitempty" type:"Repeated"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DeleteAScriptsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAScriptsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAScriptsRequest) SetAScriptIds(v []*string) *DeleteAScriptsRequest {
	s.AScriptIds = v
	return s
}

func (s *DeleteAScriptsRequest) SetClientToken(v string) *DeleteAScriptsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAScriptsRequest) SetDryRun(v bool) *DeleteAScriptsRequest {
	s.DryRun = &v
	return s
}

type DeleteAScriptsResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAScriptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAScriptsResponseBody) SetJobId(v string) *DeleteAScriptsResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteAScriptsResponseBody) SetRequestId(v string) *DeleteAScriptsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAScriptsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAScriptsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAScriptsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAScriptsResponse) SetHeaders(v map[string]*string) *DeleteAScriptsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAScriptsResponse) SetStatusCode(v int32) *DeleteAScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAScriptsResponse) SetBody(v *DeleteAScriptsResponseBody) *DeleteAScriptsResponse {
	s.Body = v
	return s
}

type DeleteAclRequest struct {
	AclId       *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DeleteAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclRequest) GoString() string {
	return s.String()
}

func (s *DeleteAclRequest) SetAclId(v string) *DeleteAclRequest {
	s.AclId = &v
	return s
}

func (s *DeleteAclRequest) SetClientToken(v string) *DeleteAclRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAclRequest) SetDryRun(v bool) *DeleteAclRequest {
	s.DryRun = &v
	return s
}

type DeleteAclResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclResponseBody) SetJobId(v string) *DeleteAclResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteAclResponseBody) SetRequestId(v string) *DeleteAclResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAclResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclResponse) GoString() string {
	return s.String()
}

func (s *DeleteAclResponse) SetHeaders(v map[string]*string) *DeleteAclResponse {
	s.Headers = v
	return s
}

func (s *DeleteAclResponse) SetStatusCode(v int32) *DeleteAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAclResponse) SetBody(v *DeleteAclResponseBody) *DeleteAclResponse {
	s.Body = v
	return s
}

type DeleteHealthCheckTemplatesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to precheck the API request. Valid values:
	//
	// *   **true**: only prechecks the API request. If you select this option, the specified endpoint service is not created after the request passes the precheck. The system prechecks the required parameters, request format, and service limits. If the request fails the precheck, the corresponding error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false**: checks the request. After the request passes the check, an **http\_2xx** status code is returned and the operation is performed. This is the default value.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IDs of the health check templates that you want to delete. You can specify up to 10 IDs.
	HealthCheckTemplateIds []*string `json:"HealthCheckTemplateIds,omitempty" xml:"HealthCheckTemplateIds,omitempty" type:"Repeated"`
}

func (s DeleteHealthCheckTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHealthCheckTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckTemplatesRequest) SetClientToken(v string) *DeleteHealthCheckTemplatesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteHealthCheckTemplatesRequest) SetDryRun(v bool) *DeleteHealthCheckTemplatesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteHealthCheckTemplatesRequest) SetHealthCheckTemplateIds(v []*string) *DeleteHealthCheckTemplatesRequest {
	s.HealthCheckTemplateIds = v
	return s
}

type DeleteHealthCheckTemplatesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHealthCheckTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHealthCheckTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckTemplatesResponseBody) SetRequestId(v string) *DeleteHealthCheckTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHealthCheckTemplatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteHealthCheckTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHealthCheckTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHealthCheckTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckTemplatesResponse) SetHeaders(v map[string]*string) *DeleteHealthCheckTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteHealthCheckTemplatesResponse) SetStatusCode(v int32) *DeleteHealthCheckTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHealthCheckTemplatesResponse) SetBody(v *DeleteHealthCheckTemplatesResponseBody) *DeleteHealthCheckTemplatesResponse {
	s.Body = v
	return s
}

type DeleteListenerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck this request. Valid values:
	//
	// *   **true**: prechecks the request but does not delete the listener. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false**: sends the API request. This is the default value. If the request passes the precheck, an `HTTP_2xx` status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the ALB instance is deployed.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s DeleteListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteListenerRequest) SetClientToken(v string) *DeleteListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteListenerRequest) SetDryRun(v bool) *DeleteListenerRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteListenerRequest) SetListenerId(v string) *DeleteListenerRequest {
	s.ListenerId = &v
	return s
}

type DeleteListenerResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponseBody) SetJobId(v string) *DeleteListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteListenerResponseBody) SetRequestId(v string) *DeleteListenerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteListenerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponse) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponse) SetHeaders(v map[string]*string) *DeleteListenerResponse {
	s.Headers = v
	return s
}

func (s *DeleteListenerResponse) SetStatusCode(v int32) *DeleteListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteListenerResponse) SetBody(v *DeleteListenerResponseBody) *DeleteListenerResponse {
	s.Body = v
	return s
}

type DeleteLoadBalancerRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun         *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DeleteLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerRequest) SetClientToken(v string) *DeleteLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetDryRun(v bool) *DeleteLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

type DeleteLoadBalancerResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponseBody) SetJobId(v string) *DeleteLoadBalancerResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteLoadBalancerResponseBody) SetRequestId(v string) *DeleteLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLoadBalancerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponse) SetHeaders(v map[string]*string) *DeleteLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoadBalancerResponse) SetStatusCode(v int32) *DeleteLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLoadBalancerResponse) SetBody(v *DeleteLoadBalancerResponseBody) *DeleteLoadBalancerResponse {
	s.Body = v
	return s
}

type DeleteRuleRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) SetClientToken(v string) *DeleteRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRuleRequest) SetDryRun(v bool) *DeleteRuleRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleId(v string) *DeleteRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteRuleResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) SetJobId(v string) *DeleteRuleResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponse) SetHeaders(v map[string]*string) *DeleteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRuleResponse) SetStatusCode(v int32) *DeleteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRuleResponse) SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse {
	s.Body = v
	return s
}

type DeleteRulesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId** as **ClientToken**. The value of **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, the system returns an `HTTP 2xx` status code and performs the operation.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The list of forwarding rules that you want to delete.
	//
	// >  The RuleIds parameter is required. You can specify up to 10 forwarding rules in each request.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s DeleteRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRulesRequest) SetClientToken(v string) *DeleteRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRulesRequest) SetDryRun(v bool) *DeleteRulesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRulesRequest) SetRuleIds(v []*string) *DeleteRulesRequest {
	s.RuleIds = v
	return s
}

type DeleteRulesResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRulesResponseBody) SetJobId(v string) *DeleteRulesResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteRulesResponseBody) SetRequestId(v string) *DeleteRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRulesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRulesResponse) SetHeaders(v map[string]*string) *DeleteRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRulesResponse) SetStatusCode(v int32) *DeleteRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRulesResponse) SetBody(v *DeleteRulesResponseBody) *DeleteRulesResponse {
	s.Body = v
	return s
}

type DeleteSecurityPolicyRequest struct {
	// The ID of the security policy.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request.
	DryRun           *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s DeleteSecurityPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyRequest) SetClientToken(v string) *DeleteSecurityPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) SetDryRun(v bool) *DeleteSecurityPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) SetSecurityPolicyId(v string) *DeleteSecurityPolicyRequest {
	s.SecurityPolicyId = &v
	return s
}

type DeleteSecurityPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyResponseBody) SetRequestId(v string) *DeleteSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecurityPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecurityPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyResponse) SetHeaders(v map[string]*string) *DeleteSecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityPolicyResponse) SetStatusCode(v int32) *DeleteSecurityPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityPolicyResponse) SetBody(v *DeleteSecurityPolicyResponseBody) *DeleteSecurityPolicyResponse {
	s.Body = v
	return s
}

type DeleteServerGroupRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s DeleteServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupRequest) SetClientToken(v string) *DeleteServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServerGroupRequest) SetDryRun(v bool) *DeleteServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteServerGroupRequest) SetServerGroupId(v string) *DeleteServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

type DeleteServerGroupResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupResponseBody) SetJobId(v string) *DeleteServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteServerGroupResponseBody) SetRequestId(v string) *DeleteServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServerGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupResponse) SetHeaders(v map[string]*string) *DeleteServerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerGroupResponse) SetStatusCode(v int32) *DeleteServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServerGroupResponse) SetBody(v *DeleteServerGroupResponseBody) *DeleteServerGroupResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The supported natural language. Valid values:
	//
	// *   zh-CN: **Chinese**
	// *   en-US: **English**
	// *   ja: **Japanese**
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The name of the region.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of region service.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	// The supported natural language. Valid values:
	//
	// *   **zh-CN**: Chinese
	// *   **en-US**: English
	// *   **ja**: Japanese
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetAcceptLanguage(v string) *DescribeZonesRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeZonesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of zones.
	Zones []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	// The name of the zone.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetLocalName(v string) *DescribeZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DetachCommonBandwidthPackageFromLoadBalancerRequest struct {
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The ID of the asynchronous task.
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun         *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachCommonBandwidthPackageFromLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachCommonBandwidthPackageFromLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetBandwidthPackageId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetClientToken(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetDryRun(v bool) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetLoadBalancerId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetRegionId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.RegionId = &v
	return s
}

type DetachCommonBandwidthPackageFromLoadBalancerResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachCommonBandwidthPackageFromLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachCommonBandwidthPackageFromLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerResponseBody) SetJobId(v string) *DetachCommonBandwidthPackageFromLoadBalancerResponseBody {
	s.JobId = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerResponseBody) SetRequestId(v string) *DetachCommonBandwidthPackageFromLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type DetachCommonBandwidthPackageFromLoadBalancerResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachCommonBandwidthPackageFromLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachCommonBandwidthPackageFromLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachCommonBandwidthPackageFromLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerResponse) SetHeaders(v map[string]*string) *DetachCommonBandwidthPackageFromLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerResponse) SetStatusCode(v int32) *DetachCommonBandwidthPackageFromLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerResponse) SetBody(v *DetachCommonBandwidthPackageFromLoadBalancerResponseBody) *DetachCommonBandwidthPackageFromLoadBalancerResponse {
	s.Body = v
	return s
}

type DisableDeletionProtectionRequest struct {
	// The ID of the Application Load Balancer (ALB) instance.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request.
	DryRun     *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DisableDeletionProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *DisableDeletionProtectionRequest) SetClientToken(v string) *DisableDeletionProtectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableDeletionProtectionRequest) SetDryRun(v bool) *DisableDeletionProtectionRequest {
	s.DryRun = &v
	return s
}

func (s *DisableDeletionProtectionRequest) SetResourceId(v string) *DisableDeletionProtectionRequest {
	s.ResourceId = &v
	return s
}

type DisableDeletionProtectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDeletionProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDeletionProtectionResponseBody) SetRequestId(v string) *DisableDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

type DisableDeletionProtectionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableDeletionProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *DisableDeletionProtectionResponse) SetHeaders(v map[string]*string) *DisableDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *DisableDeletionProtectionResponse) SetStatusCode(v int32) *DisableDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDeletionProtectionResponse) SetBody(v *DisableDeletionProtectionResponseBody) *DisableDeletionProtectionResponse {
	s.Body = v
	return s
}

type DisableLoadBalancerAccessLogRequest struct {
	// The ID of the ALB instance.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request.
	DryRun         *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DisableLoadBalancerAccessLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerAccessLogRequest) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerAccessLogRequest) SetClientToken(v string) *DisableLoadBalancerAccessLogRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableLoadBalancerAccessLogRequest) SetDryRun(v bool) *DisableLoadBalancerAccessLogRequest {
	s.DryRun = &v
	return s
}

func (s *DisableLoadBalancerAccessLogRequest) SetLoadBalancerId(v string) *DisableLoadBalancerAccessLogRequest {
	s.LoadBalancerId = &v
	return s
}

type DisableLoadBalancerAccessLogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLoadBalancerAccessLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerAccessLogResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerAccessLogResponseBody) SetRequestId(v string) *DisableLoadBalancerAccessLogResponseBody {
	s.RequestId = &v
	return s
}

type DisableLoadBalancerAccessLogResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableLoadBalancerAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableLoadBalancerAccessLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerAccessLogResponse) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerAccessLogResponse) SetHeaders(v map[string]*string) *DisableLoadBalancerAccessLogResponse {
	s.Headers = v
	return s
}

func (s *DisableLoadBalancerAccessLogResponse) SetStatusCode(v int32) *DisableLoadBalancerAccessLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableLoadBalancerAccessLogResponse) SetBody(v *DisableLoadBalancerAccessLogResponseBody) *DisableLoadBalancerAccessLogResponse {
	s.Body = v
	return s
}

type DisableLoadBalancerIpv6InternetRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among all requests. The token can only contain ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically sets **ClientToken** to the value of **RequestId**. The value of **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// *   **true:** performs a dry run. The system prechecks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false:** performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed. This is the default value.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the ALB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DisableLoadBalancerIpv6InternetRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerIpv6InternetRequest) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerIpv6InternetRequest) SetClientToken(v string) *DisableLoadBalancerIpv6InternetRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetRequest) SetDryRun(v bool) *DisableLoadBalancerIpv6InternetRequest {
	s.DryRun = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetRequest) SetLoadBalancerId(v string) *DisableLoadBalancerIpv6InternetRequest {
	s.LoadBalancerId = &v
	return s
}

type DisableLoadBalancerIpv6InternetResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLoadBalancerIpv6InternetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerIpv6InternetResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerIpv6InternetResponseBody) SetJobId(v string) *DisableLoadBalancerIpv6InternetResponseBody {
	s.JobId = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetResponseBody) SetRequestId(v string) *DisableLoadBalancerIpv6InternetResponseBody {
	s.RequestId = &v
	return s
}

type DisableLoadBalancerIpv6InternetResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableLoadBalancerIpv6InternetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableLoadBalancerIpv6InternetResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerIpv6InternetResponse) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerIpv6InternetResponse) SetHeaders(v map[string]*string) *DisableLoadBalancerIpv6InternetResponse {
	s.Headers = v
	return s
}

func (s *DisableLoadBalancerIpv6InternetResponse) SetStatusCode(v int32) *DisableLoadBalancerIpv6InternetResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetResponse) SetBody(v *DisableLoadBalancerIpv6InternetResponseBody) *DisableLoadBalancerIpv6InternetResponse {
	s.Body = v
	return s
}

type DissociateAclsFromListenerRequest struct {
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	// The ID of the asynchronous task.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** for each API request may be different.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether only to precheck this request. Valid values:
	//
	// *   **true**: prechecks the request but does not disassociate the network ACL from the listener. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the NAT CIDR block is created.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s DissociateAclsFromListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateAclsFromListenerRequest) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerRequest) SetAclIds(v []*string) *DissociateAclsFromListenerRequest {
	s.AclIds = v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetClientToken(v string) *DissociateAclsFromListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetDryRun(v bool) *DissociateAclsFromListenerRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetListenerId(v string) *DissociateAclsFromListenerRequest {
	s.ListenerId = &v
	return s
}

type DissociateAclsFromListenerResponseBody struct {
	// The ID of the request.
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateAclsFromListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateAclsFromListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerResponseBody) SetJobId(v string) *DissociateAclsFromListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *DissociateAclsFromListenerResponseBody) SetRequestId(v string) *DissociateAclsFromListenerResponseBody {
	s.RequestId = &v
	return s
}

type DissociateAclsFromListenerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DissociateAclsFromListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissociateAclsFromListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateAclsFromListenerResponse) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerResponse) SetHeaders(v map[string]*string) *DissociateAclsFromListenerResponse {
	s.Headers = v
	return s
}

func (s *DissociateAclsFromListenerResponse) SetStatusCode(v int32) *DissociateAclsFromListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateAclsFromListenerResponse) SetBody(v *DissociateAclsFromListenerResponseBody) *DissociateAclsFromListenerResponse {
	s.Body = v
	return s
}

type DissociateAdditionalCertificatesFromListenerRequest struct {
	// The ID of the asynchronous task.
	Certificates []*DissociateAdditionalCertificatesFromListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	ClientToken  *string                                                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun       *bool                                                              `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, the system returns an **HTTP 2xx** status code and performs the operation.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerRequest) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetCertificates(v []*DissociateAdditionalCertificatesFromListenerRequestCertificates) *DissociateAdditionalCertificatesFromListenerRequest {
	s.Certificates = v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetClientToken(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetDryRun(v bool) *DissociateAdditionalCertificatesFromListenerRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetListenerId(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.ListenerId = &v
	return s
}

type DissociateAdditionalCertificatesFromListenerRequestCertificates struct {
	// The ID of the request.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerRequestCertificates) String() string {
	return tea.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerRequestCertificates) SetCertificateId(v string) *DissociateAdditionalCertificatesFromListenerRequestCertificates {
	s.CertificateId = &v
	return s
}

type DissociateAdditionalCertificatesFromListenerResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) SetJobId(v string) *DissociateAdditionalCertificatesFromListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) SetRequestId(v string) *DissociateAdditionalCertificatesFromListenerResponseBody {
	s.RequestId = &v
	return s
}

type DissociateAdditionalCertificatesFromListenerResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DissociateAdditionalCertificatesFromListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissociateAdditionalCertificatesFromListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerResponse) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) SetHeaders(v map[string]*string) *DissociateAdditionalCertificatesFromListenerResponse {
	s.Headers = v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) SetStatusCode(v int32) *DissociateAdditionalCertificatesFromListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) SetBody(v *DissociateAdditionalCertificatesFromListenerResponseBody) *DissociateAdditionalCertificatesFromListenerResponse {
	s.Body = v
	return s
}

type EnableDeletionProtectionRequest struct {
	// The ID of the Application Load Balancer (ALB) instance.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request.
	DryRun     *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s EnableDeletionProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *EnableDeletionProtectionRequest) SetClientToken(v string) *EnableDeletionProtectionRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableDeletionProtectionRequest) SetDryRun(v bool) *EnableDeletionProtectionRequest {
	s.DryRun = &v
	return s
}

func (s *EnableDeletionProtectionRequest) SetResourceId(v string) *EnableDeletionProtectionRequest {
	s.ResourceId = &v
	return s
}

type EnableDeletionProtectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDeletionProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *EnableDeletionProtectionResponseBody) SetRequestId(v string) *EnableDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

type EnableDeletionProtectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableDeletionProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *EnableDeletionProtectionResponse) SetHeaders(v map[string]*string) *EnableDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *EnableDeletionProtectionResponse) SetStatusCode(v int32) *EnableDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableDeletionProtectionResponse) SetBody(v *EnableDeletionProtectionResponseBody) *EnableDeletionProtectionResponse {
	s.Body = v
	return s
}

type EnableLoadBalancerAccessLogRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun         *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LogProject     *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	LogStore       *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
}

func (s EnableLoadBalancerAccessLogRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerAccessLogRequest) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerAccessLogRequest) SetClientToken(v string) *EnableLoadBalancerAccessLogRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetDryRun(v bool) *EnableLoadBalancerAccessLogRequest {
	s.DryRun = &v
	return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetLoadBalancerId(v string) *EnableLoadBalancerAccessLogRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetLogProject(v string) *EnableLoadBalancerAccessLogRequest {
	s.LogProject = &v
	return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetLogStore(v string) *EnableLoadBalancerAccessLogRequest {
	s.LogStore = &v
	return s
}

type EnableLoadBalancerAccessLogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableLoadBalancerAccessLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerAccessLogResponseBody) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerAccessLogResponseBody) SetRequestId(v string) *EnableLoadBalancerAccessLogResponseBody {
	s.RequestId = &v
	return s
}

type EnableLoadBalancerAccessLogResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableLoadBalancerAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableLoadBalancerAccessLogResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerAccessLogResponse) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerAccessLogResponse) SetHeaders(v map[string]*string) *EnableLoadBalancerAccessLogResponse {
	s.Headers = v
	return s
}

func (s *EnableLoadBalancerAccessLogResponse) SetStatusCode(v int32) *EnableLoadBalancerAccessLogResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableLoadBalancerAccessLogResponse) SetBody(v *EnableLoadBalancerAccessLogResponseBody) *EnableLoadBalancerAccessLogResponse {
	s.Body = v
	return s
}

type EnableLoadBalancerIpv6InternetRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The ClientToken value contains only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically sets **ClientToken** to the value of **RequestId**. The value of **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// *   **true:** performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false:** performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed. This is the default value.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the ALB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s EnableLoadBalancerIpv6InternetRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerIpv6InternetRequest) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerIpv6InternetRequest) SetClientToken(v string) *EnableLoadBalancerIpv6InternetRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableLoadBalancerIpv6InternetRequest) SetDryRun(v bool) *EnableLoadBalancerIpv6InternetRequest {
	s.DryRun = &v
	return s
}

func (s *EnableLoadBalancerIpv6InternetRequest) SetLoadBalancerId(v string) *EnableLoadBalancerIpv6InternetRequest {
	s.LoadBalancerId = &v
	return s
}

type EnableLoadBalancerIpv6InternetResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableLoadBalancerIpv6InternetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerIpv6InternetResponseBody) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerIpv6InternetResponseBody) SetJobId(v string) *EnableLoadBalancerIpv6InternetResponseBody {
	s.JobId = &v
	return s
}

func (s *EnableLoadBalancerIpv6InternetResponseBody) SetRequestId(v string) *EnableLoadBalancerIpv6InternetResponseBody {
	s.RequestId = &v
	return s
}

type EnableLoadBalancerIpv6InternetResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableLoadBalancerIpv6InternetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableLoadBalancerIpv6InternetResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerIpv6InternetResponse) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerIpv6InternetResponse) SetHeaders(v map[string]*string) *EnableLoadBalancerIpv6InternetResponse {
	s.Headers = v
	return s
}

func (s *EnableLoadBalancerIpv6InternetResponse) SetStatusCode(v int32) *EnableLoadBalancerIpv6InternetResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableLoadBalancerIpv6InternetResponse) SetBody(v *EnableLoadBalancerIpv6InternetResponseBody) *EnableLoadBalancerIpv6InternetResponse {
	s.Body = v
	return s
}

type GetHealthCheckTemplateAttributeRequest struct {
	// The domain name that is used for health checks. Valid values:
	//
	// *   **$SERVER_IP**: the private IP addresses of backend servers. If you do not set the HealthCheckHost parameter or set the parameter to $SERVER_IP, the Application Load Balancer (ALB) instance uses the private IP addresses of backend servers for health checks.
	// *   **domain**: The domain name must be 1 to 80 characters in length and can contain letters, digits, periods (.), and hyphens (-).
	//
	// >  This parameter takes effect only when the `HealthCheckProtocol` parameter is set to **HTTP**.
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
}

func (s GetHealthCheckTemplateAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHealthCheckTemplateAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetHealthCheckTemplateAttributeRequest) SetHealthCheckTemplateId(v string) *GetHealthCheckTemplateAttributeRequest {
	s.HealthCheckTemplateId = &v
	return s
}

type GetHealthCheckTemplateAttributeResponseBody struct {
	// The ID of the health check template.
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The HTTP status codes that are used to determine whether the backend server passes the health check.
	//
	// *   If **HealthCheckProtocol** is set to **HTTP**, **HealthCheckCodes** can be set to **http\_2xx** (default), **http\_3xx**, **http\_4xx**, and **http\_5xx**. Separate multiple HTTP status codes with a comma (,).
	// *   If **HealthCheckProtocol** is set to **gRPC**, **HealthCheckCodes** can be set to **0 to 99**. Default value: **0**. Value ranges are supported. You can enter 20 value ranges at most and must separate each range with a comma (,).
	//
	// >  This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The HTTP version that is used for health checks.
	//
	// Valid values: **HTTP1.0** and **HTTP1.1**.
	//
	// >  This parameter takes effect only when the `HealthCheckProtocol` parameter is set to **HTTP**.
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds. Valid values: **1 to 50**.
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length and can contain only letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&). The URL can also contain the following extended characters: `_ ; ~ ! ( ) * [ ] @ $ ^ : \" , +`.
	//
	// The URL path must start with a forward slash (/).
	//
	// >  This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// *   **HTTP** (default): To perform HTTP health checks, ALB sends HEAD or GET requests to a backend server to check whether the backend server is healthy.
	// *   **TCP**: To perform TCP health checks, ALB sends SYN packets to a backend server to check whether the port of the backend server is available to receive requests.
	// *   **gRPC**: To perform gRPC health checks, ALB sends POST or GET requests to a backend server to check whether the backend server is healthy.
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The name of the health check template.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The timeout period of a health check. If a backend server does not return a health check response within the specified timeout period, the server fails the health check. Unit: seconds.
	//
	// Valid values: **1** to **300**.
	//
	// >  If the value of the `HealthCheckTimeout` parameter is smaller than that of the `HealthCheckInterval` parameter, the timeout period specified by the `HealthCheckTimeout` parameter is ignored and the value of the `HealthCheckInterval` parameter is used as the timeout period.
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The HTTP method that is used for health checks. Valid values:
	//
	// *   **HEAD**: By default, HTTP health checks use the HEAD method.
	// *   **GET**: If the size of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	// *   **POST**: By default, gRPC health checks use the POST method.
	//
	// >  This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail** to **success**.
	//
	// Valid values: **2** to **10**.
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// The ID of the request.
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success** to **fail**.
	//
	// Valid values: **2** to **10**.
	HealthyThreshold   *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UnhealthyThreshold *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s GetHealthCheckTemplateAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHealthCheckTemplateAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckCodes(v []*string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckCodes = v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckConnectPort(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckHost(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckHost = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckHttpVersion(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckInterval(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckMethod(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckMethod = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckPath(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckPath = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckProtocol(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckProtocol = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckTemplateId(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckTemplateName(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckTimeout(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckTimeout = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthyThreshold(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetRequestId(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetUnhealthyThreshold(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

type GetHealthCheckTemplateAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHealthCheckTemplateAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHealthCheckTemplateAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHealthCheckTemplateAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetHealthCheckTemplateAttributeResponse) SetHeaders(v map[string]*string) *GetHealthCheckTemplateAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponse) SetStatusCode(v int32) *GetHealthCheckTemplateAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponse) SetBody(v *GetHealthCheckTemplateAttributeResponseBody) *GetHealthCheckTemplateAttributeResponse {
	s.Body = v
	return s
}

type GetListenerAttributeRequest struct {
	// The configuration of the action. This parameter is returned and takes effect when the Type parameter is set to **FowardGroup**.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s GetListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeRequest) SetListenerId(v string) *GetListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

type GetListenerAttributeResponseBody struct {
	// The server groups to which requests are forwarded.
	AclConfig *GetListenerAttributeResponseBodyAclConfig `json:"AclConfig,omitempty" xml:"AclConfig,omitempty" type:"Struct"`
	// The ID of the listener.
	CaCertificates []*GetListenerAttributeResponseBodyCaCertificates `json:"CaCertificates,omitempty" xml:"CaCertificates,omitempty" type:"Repeated"`
	// The ID of the ALB instance.
	CaEnabled *bool `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// The timeout period of an idle connection. Unit: seconds.
	//
	// If no requests are received within the specified timeout period, Application Load Balancer (ALB) closes the current connection. When another request is received, ALB establishes a new connection.
	Certificates []*GetListenerAttributeResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The configuration of the access log.
	DefaultActions []*GetListenerAttributeResponseBodyDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
	// The Xtrace type. Supported Xtrace type: **Zipkin**.
	//
	// >  This parameter takes effect when **TracingEnabled** is set to **true**.
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// The configuration information when the listener is associated with a QUIC listener.
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The ID of the QUIC listener. This parameter is returned when **QuicUpgradeEnabled** is set to **true**. Only HTTPS listeners support this parameter.
	//
	// >  The HTTPS listener and the QUIC listener must be added to the same ALB instance. In addition, make sure that the QUIC listener is not associated with another listener.
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Indicates whether QUIC upgrade is enabled. Valid values:
	//
	// - **true**: enabled
	// - **false**: disabled
	//
	// >  Only HTTPS listeners support this parameter.
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The timeout period of a request. Unit: seconds.
	//
	// If no responses are received from the backend server within the specified timeout period, ALB returns an `HTTP 504` error code to the client.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The security policy.
	//
	// >  Only HTTPS listeners support this parameter.
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The configuration of the XForward headers.
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertClientVerifyEnabled** is set to **true**.
	//
	// The name is 1 to 40 characters in length, and can contain lowercase letters, hyphens (-), underscores (\_), and digits.
	//
	// >  Only HTTPS listeners support this parameter.
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-clientverify` header is used to retrieve the verification result of the client certificate. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	//
	// >  Only HTTPS listeners support this parameter.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-fingerprint` header is used to retrieve the fingerprint of the client certificate. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	//
	// >  Only HTTPS listeners support this parameter.
	LogConfig *GetListenerAttributeResponseBodyLogConfig `json:"LogConfig,omitempty" xml:"LogConfig,omitempty" type:"Struct"`
	// Indicates whether the `X-Forwarded-For` header is used to retrieve the client IP address. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	//
	// >  HTTP and HTTPS listeners support this parameter.
	QuicConfig *GetListenerAttributeResponseBodyQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// Indicates whether the `X-Forwarded-Port` header is used to retrieve the listening port of the ALB instance. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	//
	// >  HTTP, HTTPS, and QUIC listeners support this parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the `X-Forwarded-Client-Ip` header is used to retrieve the source port of the ALB instance. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	//
	// >  HTTP, HTTPS, and QUIC listeners support this parameter.
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The trusted proxy IP address.
	//
	// ALB traverses `X-Forwarded-For` backwards and selects the first IP address that is not in the trusted IP list as the originating IP address of the client, which will be throttled if source IP address throttling is enabled.
	SecurityPolicyId    *string                                              `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	XForwardedForConfig *GetListenerAttributeResponseBodyXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s GetListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBody) SetAclConfig(v *GetListenerAttributeResponseBodyAclConfig) *GetListenerAttributeResponseBody {
	s.AclConfig = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCaCertificates(v []*GetListenerAttributeResponseBodyCaCertificates) *GetListenerAttributeResponseBody {
	s.CaCertificates = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCaEnabled(v bool) *GetListenerAttributeResponseBody {
	s.CaEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCertificates(v []*GetListenerAttributeResponseBodyCertificates) *GetListenerAttributeResponseBody {
	s.Certificates = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetDefaultActions(v []*GetListenerAttributeResponseBodyDefaultActions) *GetListenerAttributeResponseBody {
	s.DefaultActions = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetGzipEnabled(v bool) *GetListenerAttributeResponseBody {
	s.GzipEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetHttp2Enabled(v bool) *GetListenerAttributeResponseBody {
	s.Http2Enabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetIdleTimeout(v int32) *GetListenerAttributeResponseBody {
	s.IdleTimeout = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerDescription(v string) *GetListenerAttributeResponseBody {
	s.ListenerDescription = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerId(v string) *GetListenerAttributeResponseBody {
	s.ListenerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerPort(v int32) *GetListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerProtocol(v string) *GetListenerAttributeResponseBody {
	s.ListenerProtocol = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerStatus(v string) *GetListenerAttributeResponseBody {
	s.ListenerStatus = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetLoadBalancerId(v string) *GetListenerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetLogConfig(v *GetListenerAttributeResponseBodyLogConfig) *GetListenerAttributeResponseBody {
	s.LogConfig = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetQuicConfig(v *GetListenerAttributeResponseBodyQuicConfig) *GetListenerAttributeResponseBody {
	s.QuicConfig = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRequestId(v string) *GetListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRequestTimeout(v int32) *GetListenerAttributeResponseBody {
	s.RequestTimeout = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetSecurityPolicyId(v string) *GetListenerAttributeResponseBody {
	s.SecurityPolicyId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetXForwardedForConfig(v *GetListenerAttributeResponseBodyXForwardedForConfig) *GetListenerAttributeResponseBody {
	s.XForwardedForConfig = v
	return s
}

type GetListenerAttributeResponseBodyAclConfig struct {
	// The ID of the server group to which requests are forwarded.
	AclRelations []*GetListenerAttributeResponseBodyAclConfigAclRelations `json:"AclRelations,omitempty" xml:"AclRelations,omitempty" type:"Repeated"`
	// Indicates whether HTTP/2 is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	//
	// >  Only HTTPS listeners support this parameter.
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
}

func (s GetListenerAttributeResponseBodyAclConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyAclConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyAclConfig) SetAclRelations(v []*GetListenerAttributeResponseBodyAclConfigAclRelations) *GetListenerAttributeResponseBodyAclConfig {
	s.AclRelations = v
	return s
}

func (s *GetListenerAttributeResponseBodyAclConfig) SetAclType(v string) *GetListenerAttributeResponseBodyAclConfig {
	s.AclType = &v
	return s
}

type GetListenerAttributeResponseBodyAclConfigAclRelations struct {
	// The action type.
	//
	// If **ForwardGroup** is returned, it indicates that requests are forwarded to multiple vServer groups.
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Indicates whether gzip compression is enabled to compress specific types of files. Valid values:
	//
	// - **true**: enabled
	// - **false**: disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerAttributeResponseBodyAclConfigAclRelations) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyAclConfigAclRelations) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyAclConfigAclRelations) SetAclId(v string) *GetListenerAttributeResponseBodyAclConfigAclRelations {
	s.AclId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyAclConfigAclRelations) SetStatus(v string) *GetListenerAttributeResponseBodyAclConfigAclRelations {
	s.Status = &v
	return s
}

type GetListenerAttributeResponseBodyCaCertificates struct {
	// The frontend port that is used by the ALB instance.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The status of the listener. Valid values:
	//
	// *   **Provisioning**: The listener is being created.
	// *   **Running**: The listener is running.
	// *   **Configuring**: The listener is being configured.
	// *   **Stopped**: The listener is stopped.
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The listener protocol. Valid values: **HTTP**, **HTTPS**, and **QUIC**.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerAttributeResponseBodyCaCertificates) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyCaCertificates) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyCaCertificates) SetCertificateId(v string) *GetListenerAttributeResponseBodyCaCertificates {
	s.CertificateId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyCaCertificates) SetIsDefault(v bool) *GetListenerAttributeResponseBodyCaCertificates {
	s.IsDefault = &v
	return s
}

func (s *GetListenerAttributeResponseBodyCaCertificates) SetStatus(v string) *GetListenerAttributeResponseBodyCaCertificates {
	s.Status = &v
	return s
}

type GetListenerAttributeResponseBodyCertificates struct {
	// The description of the listener.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s GetListenerAttributeResponseBodyCertificates) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyCertificates) SetCertificateId(v string) *GetListenerAttributeResponseBodyCertificates {
	s.CertificateId = &v
	return s
}

type GetListenerAttributeResponseBodyDefaultActions struct {
	// Indicates whether custom headers are recorded in the access log. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	ForwardGroupConfig *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The sampling rate of Xtrace. Valid values: 1 to 10000.
	//
	// >  This parameter takes effect when **TracingEnabled** is set to **true**.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetListenerAttributeResponseBodyDefaultActions) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyDefaultActions) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyDefaultActions) SetForwardGroupConfig(v *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) *GetListenerAttributeResponseBodyDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *GetListenerAttributeResponseBodyDefaultActions) SetType(v string) *GetListenerAttributeResponseBodyDefaultActions {
	s.Type = &v
	return s
}

type GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig struct {
	// The configuration of Xtrace. Xtrace is used to record the requests sent to ALB.
	ServerGroupTuples []*GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// Indicates whether Xtrace is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	//
	// >  **true** is returned only if the access log feature is enabled.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type GetListenerAttributeResponseBodyLogConfig struct {
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertIssuerDNEnabled** is set to **true**.
	//
	// The name is 1 to 40 characters in length, and can contain lowercase letters, hyphens (-), underscores (\_), and digits.
	//
	// >  Only HTTPS listeners support this parameter.
	AccessLogRecordCustomizedHeadersEnabled *bool `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-issuerdn` header is used to retrieve information about the authority that issues the client certificate. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	//
	// >  Only HTTPS listeners support this parameter.
	AccessLogTracingConfig *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig `json:"AccessLogTracingConfig,omitempty" xml:"AccessLogTracingConfig,omitempty" type:"Struct"`
}

func (s GetListenerAttributeResponseBodyLogConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyLogConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyLogConfig) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *GetListenerAttributeResponseBodyLogConfig {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyLogConfig) SetAccessLogTracingConfig(v *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) *GetListenerAttributeResponseBodyLogConfig {
	s.AccessLogTracingConfig = v
	return s
}

type GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig struct {
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertSubjectDNEnabled** is set to **true**.
	//
	// The name is 1 to 40 characters in length, and can contain lowercase letters, hyphens (-), underscores (\_), and digits.
	//
	// >  Only HTTPS listeners support this parameter.
	TracingEnabled *bool `json:"TracingEnabled,omitempty" xml:"TracingEnabled,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-subjectdn` header is used to retrieve information about the owner of the client certificate. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	//
	// >  Only HTTPS listeners support this parameter.
	TracingSample *int32 `json:"TracingSample,omitempty" xml:"TracingSample,omitempty"`
	// Indicates whether the `X-Forwarded-Client-Port` header is used to retrieve the client port. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	//
	// >  HTTP and HTTPS listeners support this parameter.
	TracingType *string `json:"TracingType,omitempty" xml:"TracingType,omitempty"`
}

func (s GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) SetTracingEnabled(v bool) *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig {
	s.TracingEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) SetTracingSample(v int32) *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig {
	s.TracingSample = &v
	return s
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) SetTracingType(v string) *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig {
	s.TracingType = &v
	return s
}

type GetListenerAttributeResponseBodyQuicConfig struct {
	// Indicates whether the `X-Forwarded-Proto` header is used to retrieve the listener protocol. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	//
	// >  HTTP, HTTPS, and QUIC listeners support this parameter.
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// Indicates whether the `SLB-ID` header is used to retrieve the ID of the ALB instance. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	//
	// >  HTTP, HTTPS, and QUIC listeners support this parameter.
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s GetListenerAttributeResponseBodyQuicConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyQuicConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyQuicConfig) SetQuicListenerId(v string) *GetListenerAttributeResponseBodyQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyQuicConfig) SetQuicUpgradeEnabled(v bool) *GetListenerAttributeResponseBodyQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

type GetListenerAttributeResponseBodyXForwardedForConfig struct {
	XForwardedForClientCertClientVerifyAlias   *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	XForwardedForClientCertClientVerifyEnabled *bool   `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	XForwardedForClientCertFingerprintAlias    *string `json:"XForwardedForClientCertFingerprintAlias,omitempty" xml:"XForwardedForClientCertFingerprintAlias,omitempty"`
	XForwardedForClientCertFingerprintEnabled  *bool   `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	XForwardedForClientCertIssuerDNAlias       *string `json:"XForwardedForClientCertIssuerDNAlias,omitempty" xml:"XForwardedForClientCertIssuerDNAlias,omitempty"`
	XForwardedForClientCertIssuerDNEnabled     *bool   `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	XForwardedForClientCertSubjectDNAlias      *string `json:"XForwardedForClientCertSubjectDNAlias,omitempty" xml:"XForwardedForClientCertSubjectDNAlias,omitempty"`
	XForwardedForClientCertSubjectDNEnabled    *bool   `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	XForwardedForClientSourceIpsEnabled        *bool   `json:"XForwardedForClientSourceIpsEnabled,omitempty" xml:"XForwardedForClientSourceIpsEnabled,omitempty"`
	XForwardedForClientSourceIpsTrusted        *string `json:"XForwardedForClientSourceIpsTrusted,omitempty" xml:"XForwardedForClientSourceIpsTrusted,omitempty"`
	XForwardedForClientSrcPortEnabled          *bool   `json:"XForwardedForClientSrcPortEnabled,omitempty" xml:"XForwardedForClientSrcPortEnabled,omitempty"`
	XForwardedForEnabled                       *bool   `json:"XForwardedForEnabled,omitempty" xml:"XForwardedForEnabled,omitempty"`
	XForwardedForProtoEnabled                  *bool   `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	XForwardedForSLBIdEnabled                  *bool   `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	XForwardedForSLBPortEnabled                *bool   `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s GetListenerAttributeResponseBodyXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientSourceIpsEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientSourceIpsEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientSourceIpsTrusted(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientSourceIpsTrusted = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

type GetListenerAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponse) SetHeaders(v map[string]*string) *GetListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetListenerAttributeResponse) SetStatusCode(v int32) *GetListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListenerAttributeResponse) SetBody(v *GetListenerAttributeResponseBody) *GetListenerAttributeResponse {
	s.Body = v
	return s
}

type GetListenerHealthStatusRequest struct {
	// The ID of the listener.
	IncludeRule *bool `json:"IncludeRule,omitempty" xml:"IncludeRule,omitempty"`
	// The health check status of the server groups that are associated with the listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listening protocol.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The listening port.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetListenerHealthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusRequest) SetIncludeRule(v bool) *GetListenerHealthStatusRequest {
	s.IncludeRule = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetListenerId(v string) *GetListenerHealthStatusRequest {
	s.ListenerId = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetMaxResults(v int64) *GetListenerHealthStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetNextToken(v string) *GetListenerHealthStatusRequest {
	s.NextToken = &v
	return s
}

type GetListenerHealthStatusResponseBody struct {
	// The information about server groups.
	ListenerHealthStatus []*GetListenerHealthStatusResponseBodyListenerHealthStatus `json:"ListenerHealthStatus,omitempty" xml:"ListenerHealthStatus,omitempty" type:"Repeated"`
	NextToken            *string                                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Indicates whether health checks are enabled. If **on** is returned, it indicates that health checks are enabled.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of unhealthy backend servers.
	RuleHealthStatus []*GetListenerHealthStatusResponseBodyRuleHealthStatus `json:"RuleHealthStatus,omitempty" xml:"RuleHealthStatus,omitempty" type:"Repeated"`
}

func (s GetListenerHealthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBody) SetListenerHealthStatus(v []*GetListenerHealthStatusResponseBodyListenerHealthStatus) *GetListenerHealthStatusResponseBody {
	s.ListenerHealthStatus = v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetNextToken(v string) *GetListenerHealthStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetRequestId(v string) *GetListenerHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetRuleHealthStatus(v []*GetListenerHealthStatusResponseBodyRuleHealthStatus) *GetListenerHealthStatusResponseBody {
	s.RuleHealthStatus = v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatus struct {
	// Indicates whether health checks are enabled. If **on** is returned, it indicates that health checks are enabled.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The list of unhealthy backend servers.
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The backend port.
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The cause of the health check failure.
	ServerGroupInfos []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos `json:"ServerGroupInfos,omitempty" xml:"ServerGroupInfos,omitempty" type:"Repeated"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatus) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatus) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetListenerId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ListenerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetListenerPort(v int32) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ListenerPort = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetListenerProtocol(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ListenerProtocol = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetServerGroupInfos(v []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ServerGroupInfos = v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos struct {
	// The list of server groups.
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The HTTP status code returned from the backend server. For example, **302**.
	//
	// >  This value is returned only if the return value of `ReasonCode` is **RESPONSE_MISMATCH**.
	HealthCheckEnabled *string `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The HTTP status code returned after backend servers pass health checks.
	//
	// Valid values: **HTTP\_2xx**, **HTTP\_3xx**, **HTTP\_4xx**, and **HTTP\_5xx**. Multiple status codes are separated by commas (,).
	//
	// >  This value is returned only if the return value of **ReasonCode** is **RESPONSE_MISMATCH**.
	NonNormalServers []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers `json:"NonNormalServers,omitempty" xml:"NonNormalServers,omitempty" type:"Repeated"`
	// The ID of the forwarding rule.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetActionType(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.ActionType = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetHealthCheckEnabled(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.HealthCheckEnabled = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetNonNormalServers(v []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.NonNormalServers = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetServerGroupId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.ServerGroupId = &v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers struct {
	// The reason why **Status** is Unhealthy. Only HTTP and HTTPS listeners support this parameter.
	//
	// *   **CONNECT_TIMEOUT**: Application Load Balancer (ALB) failed to connect to the backend server within the specified period of time.
	// *   **CONNECT_FAILED**: ALB failed to connect to the backend server.
	// *   **RECV_RESPONSE_FAILED**: ALB failed to receive a response from the backend server.
	// *   **RECV_RESPONSE_TIMEOUT**: ALB failed to receive a response from the backend server within the specified period of time.
	// *   **SEND_REQUEST_FAILED**: ALB failed to send a request to the backend server.
	// *   **SEND_REQUEST_TIMEOUT**: ALB failed to send a request to the backend server within the specified period of time.
	// *   **RESPONSE_FORMAT_ERROR**: The format of the response from the backend server is invalid.
	// *   **RESPONSE_FORMAT_ERROR**: The HTTP status code returned from the backend server is not the expected one.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server.
	Reason *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The action specified for the server group.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The ID of the request.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The health check status of the forwarding rules.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetPort(v int32) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.Port = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetReason(v *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.Reason = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetServerId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.ServerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetServerIp(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.ServerIp = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetStatus(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.Status = &v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason struct {
	// The IP address of the backend server.
	ActualResponse *string `json:"ActualResponse,omitempty" xml:"ActualResponse,omitempty"`
	// The health check status. Valid values:
	//
	// *   **Initial**: indicates that health checks are configured for the ALB instance, but no data was found.
	// *   **Unhealthy**: indicates that the backend server consecutively fails health checks.
	// *   **Unused**: indicates that the weight of the backend server is 0.
	// *   **Unavailable**: indicates that health checks are disabled.
	ExpectedResponse *string `json:"ExpectedResponse,omitempty" xml:"ExpectedResponse,omitempty"`
	// The ID of the server group that is associated with the listener.
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) SetActualResponse(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason {
	s.ActualResponse = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) SetExpectedResponse(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason {
	s.ExpectedResponse = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) SetReasonCode(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason {
	s.ReasonCode = &v
	return s
}

type GetListenerHealthStatusResponseBodyRuleHealthStatus struct {
	// The backend port.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The cause of the health check failure.
	ServerGroupInfos []*GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos `json:"ServerGroupInfos,omitempty" xml:"ServerGroupInfos,omitempty" type:"Repeated"`
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatus) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatus) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatus) SetRuleId(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatus {
	s.RuleId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatus) SetServerGroupInfos(v []*GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) *GetListenerHealthStatusResponseBodyRuleHealthStatus {
	s.ServerGroupInfos = v
	return s
}

type GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos struct {
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The HTTP status code returned from the backend server. For example, **302**.
	//
	// >  A value is returned only if the return value of **ReasonCode** is **RESPONSE_MISMATCH**.
	HealthCheckEnabled *string `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The HTTP status code returned after backend servers pass health checks.
	//
	// Valid values: **HTTP\_2xx**, **HTTP\_3xx**, **HTTP\_4xx**, and **HTTP\_5xx**. Multiple status codes are separated by commas (,).
	//
	// >  A value is returned only if the return value of **ReasonCode** is **RESPONSE_MISMATCH**.
	NonNormalServers []*GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers `json:"NonNormalServers,omitempty" xml:"NonNormalServers,omitempty" type:"Repeated"`
	ServerGroupId    *string                                                                                `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) SetActionType(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos {
	s.ActionType = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) SetHealthCheckEnabled(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos {
	s.HealthCheckEnabled = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) SetNonNormalServers(v []*GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos {
	s.NonNormalServers = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) SetServerGroupId(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos {
	s.ServerGroupId = &v
	return s
}

type GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers struct {
	// The reason why **Status** is Unhealthy. Only HTTP and HTTPS listeners support this parameter.
	//
	// *   **CONNECT_TIMEOUT**: ALB failed to connect to the backend server within the specified period of time.
	// *   **CONNECT_FAILED**: ALB failed to connect to the backend server.
	// *   **RECV_RESPONSE_FAILED**: ALB failed to receive a response from the backend server.
	// *   **RECV_RESPONSE_TIMEOUT**: ALB failed to receive a response from the backend server within the specified period of time.
	// *   **SEND_REQUEST_FAILED**: ALB failed to send a request to the backend server.
	// *   **SEND_REQUEST_TIMEOUT**: ALB failed to send a request to the backend server within the specified period of time.
	// *   **RESPONSE_FORMAT_ERROR**: The format of the response from the backend server is invalid.
	// *   **RESPONSE_FORMAT_ERROR**: The HTTP status code returned from the backend server is not the expected one.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server.
	Reason *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The action specified for the server group.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// - If **NextToken** is empty, it indicates that no next query is to be sent.
	// - If a value of **NextToken** is returned, the value is the token that is used for the next query.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) SetPort(v int32) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers {
	s.Port = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) SetReason(v *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers {
	s.Reason = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) SetServerId(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers {
	s.ServerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) SetServerIp(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers {
	s.ServerIp = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) SetStatus(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers {
	s.Status = &v
	return s
}

type GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason struct {
	// The IP address of the server group.
	ActualResponse *string `json:"ActualResponse,omitempty" xml:"ActualResponse,omitempty"`
	// The health check status. Valid values:
	//
	// *   **Initial**: indicates that health checks are configured for the ALB instance, but no data was found.
	// *   **Unhealthy**: indicates that the backend server consecutively fails health checks.
	// *   **Unused**: indicates that the weight of the backend server is 0.
	// *   **Unavailable**: indicates that health checks are disabled.
	ExpectedResponse *string `json:"ExpectedResponse,omitempty" xml:"ExpectedResponse,omitempty"`
	// The ID of the server group that is associated with the listener.
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) SetActualResponse(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason {
	s.ActualResponse = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) SetExpectedResponse(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason {
	s.ExpectedResponse = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) SetReasonCode(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason {
	s.ReasonCode = &v
	return s
}

type GetListenerHealthStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetListenerHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetListenerHealthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponse) SetHeaders(v map[string]*string) *GetListenerHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *GetListenerHealthStatusResponse) SetStatusCode(v int32) *GetListenerHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListenerHealthStatusResponse) SetBody(v *GetListenerHealthStatusResponseBody) *GetListenerHealthStatusResponse {
	s.Body = v
	return s
}

type GetLoadBalancerAttributeRequest struct {
	// The Logstore.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s GetLoadBalancerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *GetLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

type GetLoadBalancerAttributeResponseBody struct {
	// The mode used to assign IP addresses to zones of the ALB instance. Valid values:
	//
	// *   **Fixed:** assigns a static IP address to the ALB instance.
	// *   **Dynamic:** dynamically assigns an IP address to each zone of the ALB instance.
	AccessLogConfig *GetLoadBalancerAttributeResponseBodyAccessLogConfig `json:"AccessLogConfig,omitempty" xml:"AccessLogConfig,omitempty" type:"Struct"`
	// The time when the resource was created. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	AddressAllocatedMode *string `json:"AddressAllocatedMode,omitempty" xml:"AddressAllocatedMode,omitempty"`
	AddressIpVersion     *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The domain name of the ALB instance.
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The configuration of the deletion protection feature.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// Indicates whether the deletion protection feature is enabled. Valid values:
	//
	// *   **true:** The deletion protection feature is enabled.
	// *   **false:** The deletion protection feature is disabled.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the deletion protection feature was enabled. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	DNSName *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	// The configuration of the billing method of the ALB instance.
	DeletionProtectionConfig *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig `json:"DeletionProtectionConfig,omitempty" xml:"DeletionProtectionConfig,omitempty" type:"Struct"`
	Ipv6AddressType          *string                                                       `json:"Ipv6AddressType,omitempty" xml:"Ipv6AddressType,omitempty"`
	// The edition of the ALB instance. The features and billing rules vary based on the edition of the ALB instance. Valid values:
	//
	// *   **Basic:** basic.
	// *   **Standard:** standard.
	// *   **StandardWithWaf:** WAF-enabled.
	LoadBalancerBillingConfig *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	// The name of the ALB instance.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	LoadBalancerBussinessStatus *string `json:"LoadBalancerBussinessStatus,omitempty" xml:"LoadBalancerBussinessStatus,omitempty"`
	// The configuration of the operation lock.
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// The reason why the ALB instance was locked. This parameter is available only if you specify the **LoadBalancerBussinessStatus** parameter to **Abnormal**.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The type of the lock. Valid values:
	//
	// *   **SecurityLocked:** The ALB instance was locked due to security reasons.
	// *   **RelatedResourceLocked:** The ALB instance was locked due to association issues.
	// *   **FinancialLocked:** The ALB instance was locked due to overdue payments.
	// *   **ResidualLocked:** The ALB instance was locked because the associated EIP bandwidth plans were released.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The status of the ALB instance. Valid values:
	//
	// *   **Inactive:** The ALB instance is disabled. ALB instances in the Inactive state do not forward traffic.
	// *   **Active:** The ALB instance is running.
	// *   **Provisioning:** The ALB instance is being created.
	// *   **Configuring:** The ALB instance is being modified.
	// *   **CreateFailed:** The system failed to create the ALB instance. In this case, you are not charged for the ALB instance.
	LoadBalancerOperationLocks []*GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks `json:"LoadBalancerOperationLocks,omitempty" xml:"LoadBalancerOperationLocks,omitempty" type:"Repeated"`
	// The status of the configuration read-only mode. Valid values:
	//
	// *   **NonProtection:** The configuration read-only mode is disabled. In this case, you cannot specify the ModificationProtectionReason parameter. If you specify the ModificationProtectionReason parameter, the value of the parameter is cleared.
	// *   **ConsoleProtection:** The configuration read-only mode is enabled. In this case, you can specify the ModificationProtectionReason parameter.
	//
	// >  If you set this parameter to **ConsoleProtection**, you cannot modify the configurations of the instance in the ALB console. However, you can call API operations to modify the configurations of the instance.
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The ID of the region where the ALB instance was deployed.
	ModificationProtectionConfig *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// The tags that were added to the ALB instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The key of tag N.
	//
	// The key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. The key cannot contain `http://` or `https://`.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The value of tag N.
	//
	// The value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The value cannot contain `http://` or `https://`.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the VPC in which the ALB instance was deployed.
	Tags []*GetLoadBalancerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The IPv4 address that is used by the ALB instance.
	//
	// This parameter is available only if you set the **AddressIPVersion** parameter to **IPv4**. The AddressType parameter determines whether the IPv4 address is public or private.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The IPv6 address that is used by the ALB instance.
	//
	// This parameter is available only if you set the **AddressIPVersion** parameter to **DualStack**. The AddressType parameter determines whether the IPv6 address is public or private.
	ZoneMappings []*GetLoadBalancerAttributeResponseBodyZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s GetLoadBalancerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBody) SetAccessLogConfig(v *GetLoadBalancerAttributeResponseBodyAccessLogConfig) *GetLoadBalancerAttributeResponseBody {
	s.AccessLogConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressAllocatedMode(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressAllocatedMode = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressIpVersion(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressIpVersion = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressType(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetBandwidthPackageId(v string) *GetLoadBalancerAttributeResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetCreateTime(v string) *GetLoadBalancerAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetDNSName(v string) *GetLoadBalancerAttributeResponseBody {
	s.DNSName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetDeletionProtectionConfig(v *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) *GetLoadBalancerAttributeResponseBody {
	s.DeletionProtectionConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetIpv6AddressType(v string) *GetLoadBalancerAttributeResponseBody {
	s.Ipv6AddressType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerBillingConfig(v *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerBillingConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerBussinessStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerBussinessStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerEdition(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerEdition = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerId(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerName(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerOperationLocks(v []*GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerOperationLocks = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetModificationProtectionConfig(v *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) *GetLoadBalancerAttributeResponseBody {
	s.ModificationProtectionConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetRegionId(v string) *GetLoadBalancerAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetRequestId(v string) *GetLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetResourceGroupId(v string) *GetLoadBalancerAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetTags(v []*GetLoadBalancerAttributeResponseBodyTags) *GetLoadBalancerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetVpcId(v string) *GetLoadBalancerAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetZoneMappings(v []*GetLoadBalancerAttributeResponseBodyZoneMappings) *GetLoadBalancerAttributeResponseBody {
	s.ZoneMappings = v
	return s
}

type GetLoadBalancerAttributeResponseBodyAccessLogConfig struct {
	// The type of the IPv4 address that is used by the ALB instance. Valid values:
	//
	// *   **Internet:** The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. In this case, the ALB instance can be accessed over the Internet.
	// *   **Intranet:** The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. In this case, the ALB instance can be accessed over the virtual private cloud (VPC) where the ALB instance is deployed.
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The ID of the Elastic IP Address (EIP) bandwidth plan which is associated with the ALB instance if the ALB instance uses a public IP address.
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyAccessLogConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyAccessLogConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyAccessLogConfig) SetLogProject(v string) *GetLoadBalancerAttributeResponseBodyAccessLogConfig {
	s.LogProject = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyAccessLogConfig) SetLogStore(v string) *GetLoadBalancerAttributeResponseBodyAccessLogConfig {
	s.LogStore = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig struct {
	// The billing method.
	//
	// The value **PostPay** is returned. The value indicates the pay-as-you-go billing method.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The status of the ALB instance. Valid values:
	//
	// *   **Abnormal:** The ALB instance is not working as expected.
	// *   **Normal:** The ALB instance is working as expected.
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetEnabled(v bool) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.Enabled = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetEnabledTime(v string) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.EnabledTime = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig struct {
	// The ID of the ALB instance.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) SetPayType(v string) *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks struct {
	// The configuration of the configuration read-only mode.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The reason why the configuration read-only mode was enabled. The reason must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The reason must start with a letter.
	//
	// This parameter is valid only if you set the **ModificationProtectionStatus** parameter to **ConsoleProtection**.
	LockType *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) SetLockReason(v string) *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks {
	s.LockReason = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) SetLockType(v string) *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks {
	s.LockType = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyModificationProtectionConfig struct {
	// The ID of the request.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the resource group.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) SetReason(v string) *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) SetStatus(v string) *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	s.Status = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyTags struct {
	// The zones and the vSwitches. You must specify at least two zones.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The IP addresses that are used by the ALB instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetKey(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetValue(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.Value = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyZoneMappings struct {
	// The ID of the vSwitch in the zone. Each zone can contain only one vSwitch and one subnet.
	LoadBalancerAddresses []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses `json:"LoadBalancerAddresses,omitempty" xml:"LoadBalancerAddresses,omitempty" type:"Repeated"`
	// The type of IPv6 address that is used by the ALB instance. Valid values:
	//
	// *   **Internet:** The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. Therefore, the ALB instance can be accessed over the Internet.
	// *   **Intranet:** The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. Therefore, the ALB instance can be accessed over the VPC in which the ALB instance is deployed.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetLoadBalancerAddresses(v []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.LoadBalancerAddresses = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetVSwitchId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetZoneId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.ZoneId = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses struct {
	// The ID of the zone where the ALB instance was deployed.
	//
	// You can call the [DescribeZones](~~189196~~) operation to query the zones of the ALB instance.
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	EipType      *string `json:"EipType,omitempty" xml:"EipType,omitempty"`
	// The protocol version. Valid values:
	//
	// *   **IPv4:** IPv4.
	// *   **DualStack:** dual stack.
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetAddress(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.Address = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetAllocationId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.AllocationId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetEipType(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.EipType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetIpv6Address(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.Ipv6Address = &v
	return s
}

type GetLoadBalancerAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLoadBalancerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLoadBalancerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponse) SetHeaders(v map[string]*string) *GetLoadBalancerAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetLoadBalancerAttributeResponse) SetStatusCode(v int32) *GetLoadBalancerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoadBalancerAttributeResponse) SetBody(v *GetLoadBalancerAttributeResponseBody) *GetLoadBalancerAttributeResponse {
	s.Body = v
	return s
}

type ListAScriptsRequest struct {
	// The IDs of the AScript rules. You can specify at most 20 IDs at a time.
	AScriptIds []*string `json:"AScriptIds,omitempty" xml:"AScriptIds,omitempty" type:"Repeated"`
	// The names of the AScript rules. You can specify at most 10 names at a time.
	AScriptNames []*string `json:"AScriptNames,omitempty" xml:"AScriptNames,omitempty" type:"Repeated"`
	// The IDs of the listeners. You can specify at most 20 IDs in each request.
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return.
	//
	// Valid values: **1** to **100**.
	//
	// Default value: **20**. If you do not set this parameter, the default value is used.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no subsequent query is to be sent.
	// *   If a value of **NextToken** is returned, the value is the token used for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAScriptsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAScriptsRequest) GoString() string {
	return s.String()
}

func (s *ListAScriptsRequest) SetAScriptIds(v []*string) *ListAScriptsRequest {
	s.AScriptIds = v
	return s
}

func (s *ListAScriptsRequest) SetAScriptNames(v []*string) *ListAScriptsRequest {
	s.AScriptNames = v
	return s
}

func (s *ListAScriptsRequest) SetListenerIds(v []*string) *ListAScriptsRequest {
	s.ListenerIds = v
	return s
}

func (s *ListAScriptsRequest) SetMaxResults(v int32) *ListAScriptsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAScriptsRequest) SetNextToken(v string) *ListAScriptsRequest {
	s.NextToken = &v
	return s
}

type ListAScriptsResponseBody struct {
	// The list of the AScript rules.
	AScripts []*ListAScriptsResponseBodyAScripts `json:"AScripts,omitempty" xml:"AScripts,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no subsequent query is to be sent.
	// *   If a value of **NextToken** is returned, the value is the token used for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// >  This parameter is optional and is not returned by default.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAScriptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAScriptsResponseBody) SetAScripts(v []*ListAScriptsResponseBodyAScripts) *ListAScriptsResponseBody {
	s.AScripts = v
	return s
}

func (s *ListAScriptsResponseBody) SetMaxResults(v int32) *ListAScriptsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAScriptsResponseBody) SetNextToken(v string) *ListAScriptsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAScriptsResponseBody) SetRequestId(v string) *ListAScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAScriptsResponseBody) SetTotalCount(v int32) *ListAScriptsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAScriptsResponseBodyAScripts struct {
	// The ID of the AScript rule.
	AScriptId *string `json:"AScriptId,omitempty" xml:"AScriptId,omitempty"`
	// The name of the AScript rule.
	AScriptName *string `json:"AScriptName,omitempty" xml:"AScriptName,omitempty"`
	// The status of the AScript rule. Valid values:
	//
	// *   **Creating**: being created.
	// *   **Available**: created or updated.
	// *   **Configuring**: being updated.
	// *   **Deleting**: being deleted.
	AScriptStatus *string `json:"AScriptStatus,omitempty" xml:"AScriptStatus,omitempty"`
	// Indicates whether the AScript rule is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the Application Load Balancer (ALB) instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The content of the AScript rule.
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
}

func (s ListAScriptsResponseBodyAScripts) String() string {
	return tea.Prettify(s)
}

func (s ListAScriptsResponseBodyAScripts) GoString() string {
	return s.String()
}

func (s *ListAScriptsResponseBodyAScripts) SetAScriptId(v string) *ListAScriptsResponseBodyAScripts {
	s.AScriptId = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetAScriptName(v string) *ListAScriptsResponseBodyAScripts {
	s.AScriptName = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetAScriptStatus(v string) *ListAScriptsResponseBodyAScripts {
	s.AScriptStatus = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetEnabled(v bool) *ListAScriptsResponseBodyAScripts {
	s.Enabled = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetListenerId(v string) *ListAScriptsResponseBodyAScripts {
	s.ListenerId = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetLoadBalancerId(v string) *ListAScriptsResponseBodyAScripts {
	s.LoadBalancerId = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetScriptContent(v string) *ListAScriptsResponseBodyAScripts {
	s.ScriptContent = &v
	return s
}

type ListAScriptsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAScriptsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAScriptsResponse) GoString() string {
	return s.String()
}

func (s *ListAScriptsResponse) SetHeaders(v map[string]*string) *ListAScriptsResponse {
	s.Headers = v
	return s
}

func (s *ListAScriptsResponse) SetStatusCode(v int32) *ListAScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAScriptsResponse) SetBody(v *ListAScriptsResponseBody) *ListAScriptsResponse {
	s.Body = v
	return s
}

type ListAclEntriesRequest struct {
	// The ACL ID.
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The number of entries per page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// *   You do not need to specify this parameter for the first request.
	// *   You must specify the token that is obtained from the previous query as the value of **NextToken**.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAclEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAclEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListAclEntriesRequest) SetAclId(v string) *ListAclEntriesRequest {
	s.AclId = &v
	return s
}

func (s *ListAclEntriesRequest) SetMaxResults(v int32) *ListAclEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAclEntriesRequest) SetNextToken(v string) *ListAclEntriesRequest {
	s.NextToken = &v
	return s
}

type ListAclEntriesResponseBody struct {
	// The ACL entries.
	AclEntries []*ListAclEntriesResponseBodyAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The number of entries per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// *   If **NextToken** is empty, no next page exists.
	// *   If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAclEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAclEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclEntriesResponseBody) SetAclEntries(v []*ListAclEntriesResponseBodyAclEntries) *ListAclEntriesResponseBody {
	s.AclEntries = v
	return s
}

func (s *ListAclEntriesResponseBody) SetMaxResults(v int32) *ListAclEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAclEntriesResponseBody) SetNextToken(v string) *ListAclEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAclEntriesResponseBody) SetRequestId(v string) *ListAclEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAclEntriesResponseBody) SetTotalCount(v int32) *ListAclEntriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAclEntriesResponseBodyAclEntries struct {
	// The description of the ACL entry. The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (\_).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The CIDR block of the ACL entry.
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The status of the ACL entry. Valid values:
	//
	// *   **Adding**
	// *   **Available**
	// *   **Removing**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAclEntriesResponseBodyAclEntries) String() string {
	return tea.Prettify(s)
}

func (s ListAclEntriesResponseBodyAclEntries) GoString() string {
	return s.String()
}

func (s *ListAclEntriesResponseBodyAclEntries) SetDescription(v string) *ListAclEntriesResponseBodyAclEntries {
	s.Description = &v
	return s
}

func (s *ListAclEntriesResponseBodyAclEntries) SetEntry(v string) *ListAclEntriesResponseBodyAclEntries {
	s.Entry = &v
	return s
}

func (s *ListAclEntriesResponseBodyAclEntries) SetStatus(v string) *ListAclEntriesResponseBodyAclEntries {
	s.Status = &v
	return s
}

type ListAclEntriesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAclEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAclEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAclEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListAclEntriesResponse) SetHeaders(v map[string]*string) *ListAclEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListAclEntriesResponse) SetStatusCode(v int32) *ListAclEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAclEntriesResponse) SetBody(v *ListAclEntriesResponseBody) *ListAclEntriesResponse {
	s.Body = v
	return s
}

type ListAclRelationsRequest struct {
	// The listeners associated with the specified ACL.
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
}

func (s ListAclRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAclRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListAclRelationsRequest) SetAclIds(v []*string) *ListAclRelationsRequest {
	s.AclIds = v
	return s
}

type ListAclRelationsResponseBody struct {
	// The ID of the listener.
	AclRelations []*ListAclRelationsResponseBodyAclRelations `json:"AclRelations,omitempty" xml:"AclRelations,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAclRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAclRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponseBody) SetAclRelations(v []*ListAclRelationsResponseBodyAclRelations) *ListAclRelationsResponseBody {
	s.AclRelations = v
	return s
}

func (s *ListAclRelationsResponseBody) SetRequestId(v string) *ListAclRelationsResponseBody {
	s.RequestId = &v
	return s
}

type ListAclRelationsResponseBodyAclRelations struct {
	// The port of the listener.
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The protocol used by the listener.
	RelatedListeners []*ListAclRelationsResponseBodyAclRelationsRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Repeated"`
}

func (s ListAclRelationsResponseBodyAclRelations) String() string {
	return tea.Prettify(s)
}

func (s ListAclRelationsResponseBodyAclRelations) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponseBodyAclRelations) SetAclId(v string) *ListAclRelationsResponseBodyAclRelations {
	s.AclId = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelations) SetRelatedListeners(v []*ListAclRelationsResponseBodyAclRelationsRelatedListeners) *ListAclRelationsResponseBodyAclRelations {
	s.RelatedListeners = v
	return s
}

type ListAclRelationsResponseBodyAclRelationsRelatedListeners struct {
	// The ID of the SLB instance.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The status of association between the listener and the ACL.
	//
	// *   **Associating**: The listener is being associated with the ACL.
	// *   **Associated**: The listener is associated with the ACL.
	// *   **Dissociating**: The listener is being disassociated from the ACL.
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the request.
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId   *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAclRelationsResponseBodyAclRelationsRelatedListeners) String() string {
	return tea.Prettify(s)
}

func (s ListAclRelationsResponseBodyAclRelationsRelatedListeners) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetListenerId(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.ListenerId = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetListenerPort(v int32) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetListenerProtocol(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetLoadBalancerId(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetStatus(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.Status = &v
	return s
}

type ListAclRelationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAclRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAclRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAclRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponse) SetHeaders(v map[string]*string) *ListAclRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListAclRelationsResponse) SetStatusCode(v int32) *ListAclRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAclRelationsResponse) SetBody(v *ListAclRelationsResponseBody) *ListAclRelationsResponse {
	s.Body = v
	return s
}

type ListAclsRequest struct {
	// The names of the network ACLs. You can specify at most 10 network ACL names in each request.
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	// The ID of the resource group. You can filter the query results based on the specified ID.
	AclNames []*string `json:"AclNames,omitempty" xml:"AclNames,omitempty" type:"Repeated"`
	// The maximum number of entries to return. This parameter is optional. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The operation that you want to perform. Set the value to **ListAcls**.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IP version. **IPv4** is returned.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListAclsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAclsRequest) GoString() string {
	return s.String()
}

func (s *ListAclsRequest) SetAclIds(v []*string) *ListAclsRequest {
	s.AclIds = v
	return s
}

func (s *ListAclsRequest) SetAclNames(v []*string) *ListAclsRequest {
	s.AclNames = v
	return s
}

func (s *ListAclsRequest) SetMaxResults(v int32) *ListAclsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAclsRequest) SetNextToken(v string) *ListAclsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAclsRequest) SetResourceGroupId(v string) *ListAclsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListAclsResponseBody struct {
	// The status of the network ACL. Valid values:
	//
	// *   **Creating**: The network ACL is being created.
	// *   **Available**: The network ACL is available.
	// *   **Configuring**: The network ACL is being configured.
	Acls []*ListAclsResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	// The ID of the request.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the resource group.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The operation that you want to perform. Set the value to **ListAcls**.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the network ACL. You can specify at most 20 network ACL IDs in each request.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAclsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBody) SetAcls(v []*ListAclsResponseBodyAcls) *ListAclsResponseBody {
	s.Acls = v
	return s
}

func (s *ListAclsResponseBody) SetMaxResults(v int32) *ListAclsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAclsResponseBody) SetNextToken(v string) *ListAclsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAclsResponseBody) SetRequestId(v string) *ListAclsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAclsResponseBody) SetTotalCount(v int32) *ListAclsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAclsResponseBodyAcls struct {
	// The ID of the network ACL.
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the network ACL. You can specify at most 20 network ACL IDs in each request.
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The name of the network ACL.
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The status of configuration management. Valid values:
	//
	// *   **true**: configuration management is enabled.
	// *   **false**: configuration management is disabled.
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The network ACLs.
	ConfigManagedEnabled *bool   `json:"ConfigManagedEnabled,omitempty" xml:"ConfigManagedEnabled,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no next query is to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the value to the value of NextToken that is returned from the last call.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListAclsResponseBodyAcls) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBodyAcls) SetAclId(v string) *ListAclsResponseBodyAcls {
	s.AclId = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAclName(v string) *ListAclsResponseBodyAcls {
	s.AclName = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAclStatus(v string) *ListAclsResponseBodyAcls {
	s.AclStatus = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAddressIPVersion(v string) *ListAclsResponseBodyAcls {
	s.AddressIPVersion = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetConfigManagedEnabled(v bool) *ListAclsResponseBodyAcls {
	s.ConfigManagedEnabled = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetCreateTime(v string) *ListAclsResponseBodyAcls {
	s.CreateTime = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetResourceGroupId(v string) *ListAclsResponseBodyAcls {
	s.ResourceGroupId = &v
	return s
}

type ListAclsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAclsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAclsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponse) GoString() string {
	return s.String()
}

func (s *ListAclsResponse) SetHeaders(v map[string]*string) *ListAclsResponse {
	s.Headers = v
	return s
}

func (s *ListAclsResponse) SetStatusCode(v int32) *ListAclsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAclsResponse) SetBody(v *ListAclsResponseBody) *ListAclsResponse {
	s.Body = v
	return s
}

type ListAsynJobsRequest struct {
	// The operation that you want to perform.
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The time when the task started. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The time when the task ended. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the asynchronous task.
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no next query is to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the value to the value of **NextToken** that is returned from the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the associated resource. Valid values:
	//
	// *   **loadbalancer**: an Application Load Balancer (ALB) instance
	// *   **listener**: a listener
	// *   **rule**: a forwarding rule
	// *   **acl**: a network access control list (ACL)
	// *   **securitypolicy**: a security policy
	// *   **servergroup**: a server group
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListAsynJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsynJobsRequest) GoString() string {
	return s.String()
}

func (s *ListAsynJobsRequest) SetApiName(v string) *ListAsynJobsRequest {
	s.ApiName = &v
	return s
}

func (s *ListAsynJobsRequest) SetBeginTime(v int64) *ListAsynJobsRequest {
	s.BeginTime = &v
	return s
}

func (s *ListAsynJobsRequest) SetEndTime(v int64) *ListAsynJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAsynJobsRequest) SetJobIds(v []*string) *ListAsynJobsRequest {
	s.JobIds = v
	return s
}

func (s *ListAsynJobsRequest) SetMaxResults(v int64) *ListAsynJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAsynJobsRequest) SetNextToken(v string) *ListAsynJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAsynJobsRequest) SetResourceIds(v []*string) *ListAsynJobsRequest {
	s.ResourceIds = v
	return s
}

func (s *ListAsynJobsRequest) SetResourceType(v string) *ListAsynJobsRequest {
	s.ResourceType = &v
	return s
}

type ListAsynJobsResponseBody struct {
	// The list of tasks.
	Jobs []*ListAsynJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no next query is to be sent.
	// *   If a value of **NextToken** is returned, the value is the token that is used for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAsynJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAsynJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsynJobsResponseBody) SetJobs(v []*ListAsynJobsResponseBodyJobs) *ListAsynJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListAsynJobsResponseBody) SetMaxResults(v int64) *ListAsynJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAsynJobsResponseBody) SetNextToken(v string) *ListAsynJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAsynJobsResponseBody) SetRequestId(v string) *ListAsynJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsynJobsResponseBody) SetTotalCount(v int64) *ListAsynJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAsynJobsResponseBodyJobs struct {
	// The operation performed.
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The timestamp that indicates the start time of the task. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// If the **Status** parameter indicates Failed, an error code is returned.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// If the **Status** parameter indicates Failed, an error message is returned.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the task.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The timestamp that indicates the end time of the task. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The operation type. Valid values:
	//
	// *   **Create**: creation
	// *   **Update**: modification
	// *   **Delete**: deletion
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The ID of the associated resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the associated resource. Valid values:
	//
	// *   **loadbalancer**: an ALB instance
	// *   **listener**: a listener
	// *   **rule**: a forwarding rule
	// *   **acl**: a network ACL
	// *   **securitypolicy**: a security policy
	// *   **servergroup**: a server group
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the task. Valid values:
	//
	// *   **Succeeded**: The task is successful.
	// *   **Failed**: The task failed.
	// *   **Processing**: The task is being processed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAsynJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListAsynJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListAsynJobsResponseBodyJobs) SetApiName(v string) *ListAsynJobsResponseBodyJobs {
	s.ApiName = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetCreateTime(v int64) *ListAsynJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetErrorCode(v string) *ListAsynJobsResponseBodyJobs {
	s.ErrorCode = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetErrorMessage(v string) *ListAsynJobsResponseBodyJobs {
	s.ErrorMessage = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetId(v string) *ListAsynJobsResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetModifyTime(v int64) *ListAsynJobsResponseBodyJobs {
	s.ModifyTime = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetOperateType(v string) *ListAsynJobsResponseBodyJobs {
	s.OperateType = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetResourceId(v string) *ListAsynJobsResponseBodyJobs {
	s.ResourceId = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetResourceType(v string) *ListAsynJobsResponseBodyJobs {
	s.ResourceType = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetStatus(v string) *ListAsynJobsResponseBodyJobs {
	s.Status = &v
	return s
}

type ListAsynJobsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAsynJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAsynJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAsynJobsResponse) GoString() string {
	return s.String()
}

func (s *ListAsynJobsResponse) SetHeaders(v map[string]*string) *ListAsynJobsResponse {
	s.Headers = v
	return s
}

func (s *ListAsynJobsResponse) SetStatusCode(v int32) *ListAsynJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsynJobsResponse) SetBody(v *ListAsynJobsResponseBody) *ListAsynJobsResponse {
	s.Body = v
	return s
}

type ListHealthCheckTemplatesRequest struct {
	// The IDs of health check templates.
	HealthCheckTemplateIds []*string `json:"HealthCheckTemplateIds,omitempty" xml:"HealthCheckTemplateIds,omitempty" type:"Repeated"`
	// The health check templates.
	HealthCheckTemplateNames []*string `json:"HealthCheckTemplateNames,omitempty" xml:"HealthCheckTemplateNames,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// *   You do not need to specify this parameter for the first request.
	// *   You must specify the token that is obtained from the previous query as the value of **NextToken**.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListHealthCheckTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHealthCheckTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesRequest) SetHealthCheckTemplateIds(v []*string) *ListHealthCheckTemplatesRequest {
	s.HealthCheckTemplateIds = v
	return s
}

func (s *ListHealthCheckTemplatesRequest) SetHealthCheckTemplateNames(v []*string) *ListHealthCheckTemplatesRequest {
	s.HealthCheckTemplateNames = v
	return s
}

func (s *ListHealthCheckTemplatesRequest) SetMaxResults(v int32) *ListHealthCheckTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHealthCheckTemplatesRequest) SetNextToken(v string) *ListHealthCheckTemplatesRequest {
	s.NextToken = &v
	return s
}

type ListHealthCheckTemplatesResponseBody struct {
	// The health check templates.
	HealthCheckTemplates []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplates `json:"HealthCheckTemplates,omitempty" xml:"HealthCheckTemplates,omitempty" type:"Repeated"`
	// The number of entries returned per page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// *   If **NextToken** is empty, no next page exists.
	// *   If a value of **NextToken** was returned in the previous query, specify the value to obtain the next set of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHealthCheckTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHealthCheckTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesResponseBody) SetHealthCheckTemplates(v []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) *ListHealthCheckTemplatesResponseBody {
	s.HealthCheckTemplates = v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetMaxResults(v int32) *ListHealthCheckTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetNextToken(v string) *ListHealthCheckTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetRequestId(v string) *ListHealthCheckTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetTotalCount(v int32) *ListHealthCheckTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListHealthCheckTemplatesResponseBodyHealthCheckTemplates struct {
	// The status code.
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The port that is used for health checks.
	//
	// Valid values: \*\* 0 to 65535\*\*.
	//
	// Default value: **0**. If you set the value to 0, the port of a backend server is used for health checks.
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// **$SERVER_IP** (default): the private IP addresses of backend servers. If you do not set the HealthCheckHost parameter or set the parameter to $SERVER_IP, the Application Load Balancer (ALB) uses the private IP addresses of backend servers for health checks.
	//
	// **domain**: The domain name must be 1 to 80 characters in length and can contain letters, digits, periods (.), and hyphens (-).
	//
	// > This parameter is valid only if the `HealthCheckProtocol` parameter is set to **HTTP**.
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The HTTP version that is used for health checks.
	//
	// Valid values: **HTTP 1.0** and **HTTP 1.1**.
	//
	// Default value: **HTTP 1.1**.
	//
	// > This parameter is valid only if the `HealthCheckProtocol` parameter is set to **HTTP**.
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval at which health checks are performed. Unit: seconds. Valid values: **1 to 50**. Default value: **2**.
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The method that you want to use for the health check. Valid values:
	//
	// *   **HEAD**: By default, the ALB instance sends HEAD requests to a backend server to perform HTTP health checks.
	// *   **GET**: If the length of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	// *   **POST**: gRPC health checks automatically use the POST method.
	//
	// > This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length and can contain only letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&). The URL can also contain the following extended characters: `_ ; ~ ! ( ) * [ ] @ $ ^ : \" , +`. The URL must start with a forward slash (/).
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that you want to use for health checks. Valid values:
	//
	// *   **HTTP** (default): To perform HTTP health checks, ALB sends HEAD or GET requests to a backend server to check whether the backend server is healthy.
	// *   **TCP**: To perform TCP health checks, ALB sends SYN packets to a backend server to check whether the port of the backend server is available to receive requests.
	// *   **gRPC**: To perform gRPC health checks, ALB sends POST or GET requests to a backend server to check whether the backend server is healthy.
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The ID of the health check template.
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// The name of the health check template.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// The timeout period of a health check response. If a backend Elastic Compute Service (ECS) instance does not return a health check response within the specified timeout period, the server fails the health check.
	//
	// Valid values: **1 to 300**. Unit: seconds.
	//
	// Default value: **5**.
	//
	// > If the value of the **HealthCheckTimeout** parameter is smaller than that of the **HealthCheckInterval** parameter, the timeout period specified by the **HealthCheckTimeout** parameter is ignored and the value of the **HealthCheckInterval** parameter is used as the timeout period.
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail** to **success**.
	//
	// Valid values: **2 to 10**.
	//
	// Default value: **3**.
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success** to **fail**.
	//
	// Valid values: **2 to 10**.
	//
	// Default value: **3**.
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckCodes(v []*string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckCodes = v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckConnectPort(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckHost(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckHost = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckHttpVersion(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckInterval(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckInterval = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckMethod(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckMethod = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckPath(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckPath = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckProtocol(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckProtocol = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckTemplateId(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckTemplateName(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckTimeout(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckTimeout = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthyThreshold(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthyThreshold = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetUnhealthyThreshold(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.UnhealthyThreshold = &v
	return s
}

type ListHealthCheckTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHealthCheckTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHealthCheckTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHealthCheckTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesResponse) SetHeaders(v map[string]*string) *ListHealthCheckTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListHealthCheckTemplatesResponse) SetStatusCode(v int32) *ListHealthCheckTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHealthCheckTemplatesResponse) SetBody(v *ListHealthCheckTemplatesResponseBody) *ListHealthCheckTemplatesResponse {
	s.Body = v
	return s
}

type ListListenerCertificatesRequest struct {
	// The token that is used for the next query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no next query is to be sent.
	// *   If **NextToken** is not empty, the value indicates the token that is used for the next query.
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The maximum number of entries returned.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The type of the certificate. Valid values: **Ca** and **Server**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the listener. You must specify the ID of an HTTPS listener or a QUIC listener.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListListenerCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesRequest) SetCertificateType(v string) *ListListenerCertificatesRequest {
	s.CertificateType = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetListenerId(v string) *ListListenerCertificatesRequest {
	s.ListenerId = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetMaxResults(v int32) *ListListenerCertificatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetNextToken(v string) *ListListenerCertificatesRequest {
	s.NextToken = &v
	return s
}

type ListListenerCertificatesResponseBody struct {
	// Indicates whether the certificate is the default certificate of the listener. Valid values:
	//
	// *   **true**: the default certificate
	// *   **false**: an additional certificate
	Certificates []*ListListenerCertificatesResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The ID of the request.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The total number of entries returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of certificates.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the certificate. Only server certificates are supported.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenerCertificatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponseBody) SetCertificates(v []*ListListenerCertificatesResponseBodyCertificates) *ListListenerCertificatesResponseBody {
	s.Certificates = v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetMaxResults(v int32) *ListListenerCertificatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetNextToken(v string) *ListListenerCertificatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetRequestId(v string) *ListListenerCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetTotalCount(v int32) *ListListenerCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListListenerCertificatesResponseBodyCertificates struct {
	// Indicates whether the certificate is associated with the listener. Valid values:
	//
	// *   **Associating**: The certificate is being associated with the listener.
	// *   **Associated**: The certificate is associated with the listener.
	// *   **Diassociating**: The certificate is being disassociated from the listener.
	CertificateId   *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The type of the certificate.
	IsDefault *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListListenerCertificatesResponseBodyCertificates) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetCertificateId(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.CertificateId = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetCertificateType(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.CertificateType = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetIsDefault(v bool) *ListListenerCertificatesResponseBodyCertificates {
	s.IsDefault = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetStatus(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.Status = &v
	return s
}

type ListListenerCertificatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListListenerCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListListenerCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponse) SetHeaders(v map[string]*string) *ListListenerCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListListenerCertificatesResponse) SetStatusCode(v int32) *ListListenerCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListListenerCertificatesResponse) SetBody(v *ListListenerCertificatesResponseBody) *ListListenerCertificatesResponse {
	s.Body = v
	return s
}

type ListListenersRequest struct {
	// The listener IDs. You can specify up to 20 IDs.
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// The listener protocol. Valid values:
	//
	// *   **HTTP**
	// *   **HTTPS**
	// *   **QUIC**
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The Server Load Balancer (SLB) instance IDs. You can specify up to 20 IDs.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return. This parameter is optional. Valid values: **1 to 100**. If you do not specify this parameter, the default value **20** is used.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// *   You do not need to specify this parameter for the first request.
	// *   You must specify the token that is obtained from the previous query as the value of **NextToken**.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListListenersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenersRequest) GoString() string {
	return s.String()
}

func (s *ListListenersRequest) SetListenerIds(v []*string) *ListListenersRequest {
	s.ListenerIds = v
	return s
}

func (s *ListListenersRequest) SetListenerProtocol(v string) *ListListenersRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *ListListenersRequest) SetLoadBalancerIds(v []*string) *ListListenersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListListenersRequest) SetMaxResults(v int32) *ListListenersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenersRequest) SetNextToken(v string) *ListListenersRequest {
	s.NextToken = &v
	return s
}

type ListListenersResponseBody struct {
	// The listeners.
	Listeners []*ListListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position where the query stopped. If this parameter is not returned, all data is queried.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBody) SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersResponseBody) SetMaxResults(v int32) *ListListenersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListListenersResponseBody) SetNextToken(v string) *ListListenersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListListenersResponseBody) SetRequestId(v string) *ListListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersResponseBody) SetTotalCount(v int32) *ListListenersResponseBody {
	s.TotalCount = &v
	return s
}

type ListListenersResponseBodyListeners struct {
	// The actions of the default forwarding rule.
	DefaultActions []*ListListenersResponseBodyListenersDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
	// Indicates whether GZIP compression is enabled to compress specific types of files. Valid values:
	//
	// *   **true**
	// *   **false**
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// Indicates whether HTTP/2 is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The timeout period of an idle connection. Unit: seconds. Valid values: **1 to 60**.
	//
	// If no request is received within the specified timeout period, SLB closes the connection. SLB establishes the connection again when a new connection request is received.
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The description of the listener. The description must be 2 to 256 characters in length.
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The listener ID.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The frontend port that is used by the SLB instance. Valid values: **1 to 65535**.
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listener protocol. Valid values:
	//
	// *   **HTTP**
	// *   **HTTPS**
	// *   **QUIC**
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The status of the listener. Valid values:
	//
	// *   **Provisioning**
	// *   **Running**
	// *   **Configuring**
	// *   **Stopped**
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// The SLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The logging configuration.
	LogConfig *ListListenersResponseBodyListenersLogConfig `json:"LogConfig,omitempty" xml:"LogConfig,omitempty" type:"Struct"`
	// The configuration information when the listener is associated with a QUIC listener.
	QuicConfig *ListListenersResponseBodyListenersQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// The timeout period of a request. Unit: seconds. Valid values: **1 to 180**.
	//
	// If no responses are received from the backend server within the specified timeout period, SLB returns an `HTTP 504` error code to the client.
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The security policy.
	//
	// > This parameter is available only when you create an HTTPS listener.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The configuration of the `XForward` headers.
	XForwardedForConfig *ListListenersResponseBodyListenersXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s ListListenersResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListeners) SetDefaultActions(v []*ListListenersResponseBodyListenersDefaultActions) *ListListenersResponseBodyListeners {
	s.DefaultActions = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetGzipEnabled(v bool) *ListListenersResponseBodyListeners {
	s.GzipEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetHttp2Enabled(v bool) *ListListenersResponseBodyListeners {
	s.Http2Enabled = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetIdleTimeout(v int32) *ListListenersResponseBodyListeners {
	s.IdleTimeout = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerDescription(v string) *ListListenersResponseBodyListeners {
	s.ListenerDescription = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerId(v string) *ListListenersResponseBodyListeners {
	s.ListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerPort(v int32) *ListListenersResponseBodyListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerProtocol(v string) *ListListenersResponseBodyListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerStatus(v string) *ListListenersResponseBodyListeners {
	s.ListenerStatus = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetLoadBalancerId(v string) *ListListenersResponseBodyListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetLogConfig(v *ListListenersResponseBodyListenersLogConfig) *ListListenersResponseBodyListeners {
	s.LogConfig = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetQuicConfig(v *ListListenersResponseBodyListenersQuicConfig) *ListListenersResponseBodyListeners {
	s.QuicConfig = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetRequestTimeout(v int32) *ListListenersResponseBodyListeners {
	s.RequestTimeout = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetSecurityPolicyId(v string) *ListListenersResponseBodyListeners {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetXForwardedForConfig(v *ListListenersResponseBodyListenersXForwardedForConfig) *ListListenersResponseBodyListeners {
	s.XForwardedForConfig = v
	return s
}

type ListListenersResponseBodyListenersDefaultActions struct {
	// The configuration of the forwarding action. This parameter is returned and takes effect only if the action type is **FowardGroup**.
	ForwardGroupConfig *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The type of the action. If **ForwardGroup** is returned, requests are forwarded to multiple vServer groups.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListListenersResponseBodyListenersDefaultActions) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersDefaultActions) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersDefaultActions) SetForwardGroupConfig(v *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) *ListListenersResponseBodyListenersDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *ListListenersResponseBodyListenersDefaultActions) SetType(v string) *ListListenersResponseBodyListenersDefaultActions {
	s.Type = &v
	return s
}

type ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig struct {
	// The server groups to which requests are forwarded.
	ServerGroupTuples []*ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// The ID of the server group to which requests are forwarded.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type ListListenersResponseBodyListenersLogConfig struct {
	// Indicates whether custom headers are carried in the access log. Valid values:
	//
	// *   **true**
	// *   **false**
	AccessLogRecordCustomizedHeadersEnabled *bool `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// The configuration of Xtrace that is used to record the requests sent to SLB.
	AccessLogTracingConfig *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig `json:"AccessLogTracingConfig,omitempty" xml:"AccessLogTracingConfig,omitempty" type:"Struct"`
}

func (s ListListenersResponseBodyListenersLogConfig) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersLogConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersLogConfig) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *ListListenersResponseBodyListenersLogConfig {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersLogConfig) SetAccessLogTracingConfig(v *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) *ListListenersResponseBodyListenersLogConfig {
	s.AccessLogTracingConfig = v
	return s
}

type ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig struct {
	// Indicates whether Xtrace is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > You can set this parameter to **true** only if the access log feature is enabled by specifying **AccessLogEnabled**.
	TracingEnabled *bool `json:"TracingEnabled,omitempty" xml:"TracingEnabled,omitempty"`
	// The sampling rate of Xtrace. Valid values: **1 to 10000**.
	//
	// > This parameter takes effect only if **TracingEnabled** is set to **true**.
	TracingSample *int32 `json:"TracingSample,omitempty" xml:"TracingSample,omitempty"`
	// The Xtrace type. This parameter can be set to **Zipkin**.
	//
	// > This parameter takes effect only if **TracingEnabled** is set to **true**.
	TracingType *string `json:"TracingType,omitempty" xml:"TracingType,omitempty"`
}

func (s ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) SetTracingEnabled(v bool) *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig {
	s.TracingEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) SetTracingSample(v int32) *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig {
	s.TracingSample = &v
	return s
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) SetTracingType(v string) *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig {
	s.TracingType = &v
	return s
}

type ListListenersResponseBodyListenersQuicConfig struct {
	// The QUIC listener ID. This parameter is returned when **QuicUpgradeEnabled** is set to **true**. Only HTTPS listeners support this parameter.
	//
	// > You must add the HTTPS listener and the QUIC listener to the same ALB instance. In addition, make sure that the QUIC listener has never been associated with another listener.
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// Indicates whether QUIC upgrade is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > Only HTTPS listeners support this parameter.
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s ListListenersResponseBodyListenersQuicConfig) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersQuicConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersQuicConfig) SetQuicListenerId(v string) *ListListenersResponseBodyListenersQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListenersQuicConfig) SetQuicUpgradeEnabled(v bool) *ListListenersResponseBodyListenersQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

type ListListenersResponseBodyListenersXForwardedForConfig struct {
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertClientVerifyEnabled** is set to **true**.
	//
	// The name is 1 to 40 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertClientVerifyAlias *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-clientverify` header is used to obtain the verification result of the client certificate. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertClientVerifyEnabled *bool `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertFingerprintEnabled** is set to **true**.
	//
	// The name is 1 to 40 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertFingerprintAlias *string `json:"XForwardedForClientCertFingerprintAlias,omitempty" xml:"XForwardedForClientCertFingerprintAlias,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-fingerprint` header is used to retrieve the fingerprint of the client certificate. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertFingerprintEnabled *bool `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertIssuerDNEnabled** is set to **true**.
	//
	// The name is 1 to 40 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertIssuerDNAlias *string `json:"XForwardedForClientCertIssuerDNAlias,omitempty" xml:"XForwardedForClientCertIssuerDNAlias,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-issuerdn` header is used to retrieve information about the authority that issues the client certificate. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertIssuerDNEnabled *bool `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertSubjectDNEnabled** is set to **true**.
	//
	// The name is 1 to 40 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertSubjectDNAlias *string `json:"XForwardedForClientCertSubjectDNAlias,omitempty" xml:"XForwardedForClientCertSubjectDNAlias,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-subjectdn` header is used to retrieve information about the owner of the client certificate. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertSubjectDNEnabled *bool `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	// Indicates whether the `X-Forwarded-Client-Ip` header is used to retrieve the source port of the SLB instance. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP, HTTPS, and QUIC listeners support this parameter.
	XForwardedForClientSourceIpsEnabled *bool `json:"XForwardedForClientSourceIpsEnabled,omitempty" xml:"XForwardedForClientSourceIpsEnabled,omitempty"`
	// The trusted proxy IP address.
	//
	// ALB traverses `X-Forwarded-For` backward and selects the first IP address that is not in the trusted IP address list as the real IP address of the client. The IP address is used in source IP address throttling.
	XForwardedForClientSourceIpsTrusted *string `json:"XForwardedForClientSourceIpsTrusted,omitempty" xml:"XForwardedForClientSourceIpsTrusted,omitempty"`
	// Indicates whether the `X-Forwarded-Client-Port` header is used to retrieve the client port. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP and HTTPS listeners support this parameter.
	XForwardedForClientSrcPortEnabled *bool `json:"XForwardedForClientSrcPortEnabled,omitempty" xml:"XForwardedForClientSrcPortEnabled,omitempty"`
	// Indicates whether the `X-Forwarded-For` header is used to retrieve the client IP address. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP and HTTPS listeners support this parameter.
	XForwardedForEnabled *bool `json:"XForwardedForEnabled,omitempty" xml:"XForwardedForEnabled,omitempty"`
	// Indicates whether the `X-Forwarded-Proto` header is used to retrieve the listener protocol. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP, HTTPS, and QUIC listeners support this parameter.
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// Indicates whether the `SLB-ID` header is used to retrieve the ID of the SLB instance. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP, HTTPS, and QUIC listeners support this parameter.
	XForwardedForSLBIdEnabled *bool `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	// Indicates whether the `X-Forwarded-Port` header is used to retrieve the listener port of the SLB instance. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP, HTTPS, and QUIC listeners support this parameter.
	XForwardedForSLBPortEnabled *bool `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s ListListenersResponseBodyListenersXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientSourceIpsEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientSourceIpsEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientSourceIpsTrusted(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientSourceIpsTrusted = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

type ListListenersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListListenersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListListenersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponse) GoString() string {
	return s.String()
}

func (s *ListListenersResponse) SetHeaders(v map[string]*string) *ListListenersResponse {
	s.Headers = v
	return s
}

func (s *ListListenersResponse) SetStatusCode(v int32) *ListListenersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListListenersResponse) SetBody(v *ListListenersResponseBody) *ListListenersResponse {
	s.Body = v
	return s
}

type ListLoadBalancersRequest struct {
	// The network type. Valid values:
	//
	// *   **Internet**: The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. Therefore, the ALB instance can be accessed over the Internet.
	// *   **Intranet**: The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. In this case, the ALB instance can be accessed over the VPC where the ALB instance is deployed.
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The service status of the ALB instance. Valid values:
	//
	// *   **Abnormal**
	// *   **Normal**
	LoadBalancerBussinessStatus *string `json:"LoadBalancerBussinessStatus,omitempty" xml:"LoadBalancerBussinessStatus,omitempty"`
	// The instance IDs. You can specify at most 20 ALB instance IDs.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The names of the instances. You can specify at most 10 names.
	LoadBalancerNames []*string `json:"LoadBalancerNames,omitempty" xml:"LoadBalancerNames,omitempty" type:"Repeated"`
	// The status of the ALB instance. Valid values:
	//
	// *   **Inactive**: The ALB instance is disabled. The listeners do not forward traffic.
	// *   **Active**: The ALB instance is running.
	// *   **Provisioning**: The ALB instance is being created.
	// *   **Configuring**: The ALB instance is being modified.
	// *   **CreateFailed**: The system failed to create the ALB instance. In this case, you are not charged for the ALB instance. You can only delete the ALB instance. By default, the system deletes the ALB instances that are in the CreateFailed state within the last day.
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The number of entries to return on each page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// *   You do not need to specify this parameter for the first request.
	// *   You must specify the token that is obtained from the previous query as the value of **NextToken**.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The billing method of the ALB instance. Set the value to
	//
	// **PostPay**, which specifies the pay-as-you-go billing method. This is the default value.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags added to the ALB instance.
	Tag []*ListLoadBalancersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the ALB instance belongs. You can specify at most 10 IDs.
	VpcIds []*string `json:"VpcIds,omitempty" xml:"VpcIds,omitempty" type:"Repeated"`
	// The ID of the zone where the ALB instance is deployed.
	//
	// You can call the [DescribeZones](~~189196~~) operation to query zones.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListLoadBalancersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequest) SetAddressType(v string) *ListLoadBalancersRequest {
	s.AddressType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerBussinessStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerBussinessStatus = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerIds(v []*string) *ListLoadBalancersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerNames(v []*string) *ListLoadBalancersRequest {
	s.LoadBalancerNames = v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersRequest) SetMaxResults(v int32) *ListLoadBalancersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersRequest) SetNextToken(v string) *ListLoadBalancersRequest {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersRequest) SetPayType(v string) *ListLoadBalancersRequest {
	s.PayType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetResourceGroupId(v string) *ListLoadBalancersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersRequest) SetTag(v []*ListLoadBalancersRequestTag) *ListLoadBalancersRequest {
	s.Tag = v
	return s
}

func (s *ListLoadBalancersRequest) SetVpcIds(v []*string) *ListLoadBalancersRequest {
	s.VpcIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetZoneId(v string) *ListLoadBalancersRequest {
	s.ZoneId = &v
	return s
}

type ListLoadBalancersRequestTag struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLoadBalancersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersRequestTag) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequestTag) SetKey(v string) *ListLoadBalancersRequestTag {
	s.Key = &v
	return s
}

func (s *ListLoadBalancersRequestTag) SetValue(v string) *ListLoadBalancersRequestTag {
	s.Value = &v
	return s
}

type ListLoadBalancersResponseBody struct {
	// The list of ALB instances.
	LoadBalancers []*ListLoadBalancersResponseBodyLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// *   If **NextToken** is empty, no next page exists.
	// *   If a value is returned for **NextToken**, the value is used to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLoadBalancersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBody) SetLoadBalancers(v []*ListLoadBalancersResponseBodyLoadBalancers) *ListLoadBalancersResponseBody {
	s.LoadBalancers = v
	return s
}

func (s *ListLoadBalancersResponseBody) SetMaxResults(v int32) *ListLoadBalancersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetNextToken(v string) *ListLoadBalancersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetRequestId(v string) *ListLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetTotalCount(v int32) *ListLoadBalancersResponseBody {
	s.TotalCount = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancers struct {
	// The configuration of the access log.
	AccessLogConfig *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig `json:"AccessLogConfig,omitempty" xml:"AccessLogConfig,omitempty" type:"Struct"`
	// The mode in which IP addresses are allocated. Valid values:
	//
	// *   **Fixed**: allocates a static IP address to the ALB instance.
	// *   **Dynamic**: dynamically allocates an IP address to each zone of the ALB instance.
	AddressAllocatedMode *string `json:"AddressAllocatedMode,omitempty" xml:"AddressAllocatedMode,omitempty"`
	// The IP version. Valid values:
	//
	// *   **IPv4**
	// *   **DualStack**
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The type of IP address that the ALB instance uses to provide services. Valid values:
	//
	// *   **Internet**: The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. Therefore, the ALB instance can be accessed over the Internet.
	// *   **Intranet**: The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. In this case, the ALB instance can be accessed over the VPC where the ALB instance is deployed.
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the EIP bandwidth plan that is associated with the NLB instance if the NLB instance uses a public IP address.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The time when the resource was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The domain name of the ALB instance.
	DNSName *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	// The configuration of deletion protection.
	DeletionProtectionConfig *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig `json:"DeletionProtectionConfig,omitempty" xml:"DeletionProtectionConfig,omitempty" type:"Struct"`
	// The type of IPv6 address that is used by the ALB instance. Valid values:
	//
	// *   **Internet**: The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. Therefore, the ALB instance can be accessed over the Internet.
	// *   **Intranet**: The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. Therefore, the ALB instance can be accessed over the VPC in which the ALB instance is deployed.
	Ipv6AddressType *string `json:"Ipv6AddressType,omitempty" xml:"Ipv6AddressType,omitempty"`
	// The configuration of the billing method.
	LoadBalancerBillingConfig *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	// The business status of the ALB instance. Valid values:
	//
	// *   **Abnormal**
	// *   **Normal**
	LoadBalancerBussinessStatus *string `json:"LoadBalancerBussinessStatus,omitempty" xml:"LoadBalancerBussinessStatus,omitempty"`
	// The edition of the ALB instance. Different editions have different limits and support different billing methods. Valid values:
	//
	// *   **Basic**: basic
	// *   **Standard**: standard
	// *   **StandardWithWaf**: WAF-enabled
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// The ID of the ALB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the NLB instance.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The configuration of the operation lock.
	LoadBalancerOperationLocks []*ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks `json:"LoadBalancerOperationLocks,omitempty" xml:"LoadBalancerOperationLocks,omitempty" type:"Repeated"`
	// The status of the ALB instance. Valid values:
	//
	// *   **Inactive**: The ALB instance is disabled. The listeners do not forward traffic.
	// *   **Active**: The ALB instance is running.
	// *   **Provisioning**: The ALB instance is being created.
	// *   **Configuring**: The ALB instance is being modified.
	// *   **CreateFailed**: The system failed to create the ALB instance.
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The configuration read-only mode.
	ModificationProtectionConfig *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags that are added to the instance.
	Tags []*ListLoadBalancersResponseBodyLoadBalancersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC to which the ALB instance belongs.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancers) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancers) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAccessLogConfig(v *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AccessLogConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAddressAllocatedMode(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AddressAllocatedMode = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAddressIpVersion(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AddressIpVersion = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAddressType(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AddressType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetBandwidthPackageId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.BandwidthPackageId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetCreateTime(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.CreateTime = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetDNSName(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.DNSName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetDeletionProtectionConfig(v *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.DeletionProtectionConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetIpv6AddressType(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Ipv6AddressType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerBillingConfig(v *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerBillingConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerBussinessStatus(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerBussinessStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerEdition(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerEdition = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerName(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerOperationLocks(v []*ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerOperationLocks = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerStatus(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetModificationProtectionConfig(v *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.ModificationProtectionConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetResourceGroupId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetTags(v []*ListLoadBalancersResponseBodyLoadBalancersTags) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Tags = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetVpcId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.VpcId = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig struct {
	// The log project.
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The Logstore.
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) SetLogProject(v string) *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig {
	s.LogProject = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) SetLogStore(v string) *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig {
	s.LogStore = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig struct {
	// Indicates whether deletion protection is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time when deletion protection is enabled.
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) SetEnabled(v bool) *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig {
	s.Enabled = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) SetEnabledTime(v string) *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig {
	s.EnabledTime = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig struct {
	// The billing method. Valid values:
	//
	// Only **PostPay** may be returned, which indicates the pay-as-you-go billing method.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) SetPayType(v string) *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks struct {
	// The reason why the ALB instance is locked. This parameter is available only when **LoadBalancerBussinessStatus** is set to **Abnormal**.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The lock type. Valid values:
	//
	// *   **SecurityLocked**: The ALB instance is locked due to security reasons.
	// *   **RelatedResourceLocked**: The ALB instance is locked due to association issues.
	// *   **FinancialLocked**: The ALB instance is locked due to overdue payments.
	// *   **ResidualLocked**: The ALB instance is locked because the associated resources have overdue payments and the resources are released.
	LockType *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) SetLockReason(v string) *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks {
	s.LockReason = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) SetLockType(v string) *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks {
	s.LockType = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig struct {
	// The reason why deletion protection is enabled.
	//
	// It must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). It must start with a letter.
	//
	// This parameter takes effect only when **ModificationProtectionStatus** is set to **ConsoleProtection**.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Indicates whether the configuration read-only mode is enabled for the ALB instance. Valid values:
	//
	// *   **NonProtection**: The configuration read-only mode is disabled. In this case, you cannot specify ModificationProtectionReason. If you specify ModificationProtectionReason, the value of the parameter is cleared.
	// *   **ConsoleProtection**: The configuration read-only mode is enabled. In this case, you can specify ModificationProtectionReason.
	//
	// > If you set this parameter to **ConsoleProtection**, you cannot use the ALB console to modify instance configurations. However, you can call API operations to modify instance configurations.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) SetReason(v string) *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) SetStatus(v string) *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig {
	s.Status = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersTags struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersTags) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersTags) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) SetKey(v string) *ListLoadBalancersResponseBodyLoadBalancersTags {
	s.Key = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) SetValue(v string) *ListLoadBalancersResponseBodyLoadBalancersTags {
	s.Value = &v
	return s
}

type ListLoadBalancersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLoadBalancersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponse) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponse) SetHeaders(v map[string]*string) *ListLoadBalancersResponse {
	s.Headers = v
	return s
}

func (s *ListLoadBalancersResponse) SetStatusCode(v int32) *ListLoadBalancersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLoadBalancersResponse) SetBody(v *ListLoadBalancersResponseBody) *ListLoadBalancersResponse {
	s.Body = v
	return s
}

type ListRulesRequest struct {
	// The direction to which the forwarding rule is applied. Valid values:
	//
	// *   **Request** (default): The forwarding rule is applied to the client requests received by ALB.
	// *   **Response**: The forwarding rule is applied to the responses returned by backend servers.
	//
	// > You cannot set this parameter to Response if you use basic ALB instances.
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The listener IDs.
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// The Application Load Balancer (ALB) instance IDs.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return.
	//
	// Valid values: **1 to 100**.
	//
	// Default value: **20**. If you do not specify this parameter, the default value is used.
	//
	// > This parameter is optional.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The starting point of the current query. If you do not specify this parameter, the query starts from the beginning.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The forwarding rules.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s ListRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) SetDirection(v string) *ListRulesRequest {
	s.Direction = &v
	return s
}

func (s *ListRulesRequest) SetListenerIds(v []*string) *ListRulesRequest {
	s.ListenerIds = v
	return s
}

func (s *ListRulesRequest) SetLoadBalancerIds(v []*string) *ListRulesRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListRulesRequest) SetMaxResults(v int32) *ListRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRulesRequest) SetNextToken(v string) *ListRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRulesRequest) SetRuleIds(v []*string) *ListRulesRequest {
	s.RuleIds = v
	return s
}

type ListRulesResponseBody struct {
	// The maximum number of entries returned.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// *   If **NextToken** is empty, no next page exists.
	// *   If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The forwarding rules.
	Rules []*ListRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) SetMaxResults(v int32) *ListRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRulesResponseBody) SetNextToken(v string) *ListRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetRules(v []*ListRulesResponseBodyRules) *ListRulesResponseBody {
	s.Rules = v
	return s
}

func (s *ListRulesResponseBody) SetTotalCount(v int32) *ListRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListRulesResponseBodyRules struct {
	// 转发规则的方向。取值：
	//
	// * Request（默认值）：请求类型，对从客户端发送到ALB的报文进行条件匹配并进行相应的处理。
	//
	// * Response：响应类型，对从后端服务器组返回到ALB的报文进行条件匹配并进行相应的处理。
	//
	// > 基础版的ALB实例不支持Response类型。
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The ID of the listener to which the forwarding rule belongs.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the ALB instance to which the forwarding rule belongs.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The priority of the forwarding rule. Valid values: **1 to 10000**. A lower value indicates a higher priority.
	//
	// > The priority of each forwarding rule added to a listener must be unique.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The actions of the forwarding rule.
	RuleActions []*ListRulesResponseBodyRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// The match conditions of the forwarding rule.
	RuleConditions []*ListRulesResponseBodyRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The forwarding rule ID.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the forwarding rule. Valid values:
	//
	// *   **Provisioning**
	// *   **Configuring**
	// *   **Available**
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
}

func (s ListRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRules) SetDirection(v string) *ListRulesResponseBodyRules {
	s.Direction = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetListenerId(v string) *ListRulesResponseBodyRules {
	s.ListenerId = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetLoadBalancerId(v string) *ListRulesResponseBodyRules {
	s.LoadBalancerId = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetPriority(v int32) *ListRulesResponseBodyRules {
	s.Priority = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleActions(v []*ListRulesResponseBodyRulesRuleActions) *ListRulesResponseBodyRules {
	s.RuleActions = v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleConditions(v []*ListRulesResponseBodyRulesRuleConditions) *ListRulesResponseBodyRules {
	s.RuleConditions = v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleId(v string) *ListRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleName(v string) *ListRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleStatus(v string) *ListRulesResponseBodyRules {
	s.RuleStatus = &v
	return s
}

type ListRulesResponseBodyRulesRuleActions struct {
	// The CORS configuration.
	CorsConfig *ListRulesResponseBodyRulesRuleActionsCorsConfig `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	// The configuration of the custom response.
	FixedResponseConfig *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	// The configurations of the server groups.
	ForwardGroupConfig *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The configuration of the header to be inserted.
	InsertHeaderConfig *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// The priority of the action. Valid values: **1 to 50000**. A lower value indicates a higher priority. The actions of a forwarding rule are applied in descending order of priority. This parameter is not empty. The priority of each action within a forwarding rule is unique.
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The configuration of the redirect action.
	RedirectConfig *ListRulesResponseBodyRulesRuleActionsRedirectConfig `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	// 去除HTTP头部配置。
	RemoveHeaderConfig *ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig `json:"RemoveHeaderConfig,omitempty" xml:"RemoveHeaderConfig,omitempty" type:"Struct"`
	// The configuration of the rewrite action.
	RewriteConfig *ListRulesResponseBodyRulesRuleActionsRewriteConfig `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	// The configuration of the action to throttle traffic.
	TrafficLimitConfig *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig `json:"TrafficLimitConfig,omitempty" xml:"TrafficLimitConfig,omitempty" type:"Struct"`
	// The configuration of the action to mirror traffic.
	TrafficMirrorConfig *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig `json:"TrafficMirrorConfig,omitempty" xml:"TrafficMirrorConfig,omitempty" type:"Struct"`
	// The action type. Valid values:
	//
	// *   **ForwardGroup**: forwards a request to multiple vServer groups.
	// *   **Redirect**: redirects a request.
	// *   **FixedResponse**: returns a custom response.
	// *   **Rewrite**: rewrites a request.
	// *   **InsertHeader**: inserts a header.
	// *   **RemoveHeaderConfig**: deletes a header.
	// *   **TrafficLimitConfig**: throttles network traffic.
	// *   **TrafficMirrorConfig**: mirrors traffic.
	// *   **CorsConfig**: forwards requests based on CORS.
	//
	// The following action types are supported:
	//
	// *   **FinalType**: the last action to be performed in a forwarding rule. Each forwarding rule can contain only one FinalType action. You can specify a **ForwardGroup**, **Redirect**, or **FixedResponse** action as the FinalType action.
	// *   **ExtType**: the action or the actions to be performed before the **FinalType** action. A forwarding rule can contain one or more **ExtType** actions. To specify an ExtType action, you must specify a **FinalType** action. You can specify multiple **InsertHeader** actions or one **Rewrite** action.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActions) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActions) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActions) SetCorsConfig(v *ListRulesResponseBodyRulesRuleActionsCorsConfig) *ListRulesResponseBodyRulesRuleActions {
	s.CorsConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetFixedResponseConfig(v *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) *ListRulesResponseBodyRulesRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetForwardGroupConfig(v *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) *ListRulesResponseBodyRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetInsertHeaderConfig(v *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) *ListRulesResponseBodyRulesRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetOrder(v int32) *ListRulesResponseBodyRulesRuleActions {
	s.Order = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetRedirectConfig(v *ListRulesResponseBodyRulesRuleActionsRedirectConfig) *ListRulesResponseBodyRulesRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetRemoveHeaderConfig(v *ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig) *ListRulesResponseBodyRulesRuleActions {
	s.RemoveHeaderConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetRewriteConfig(v *ListRulesResponseBodyRulesRuleActionsRewriteConfig) *ListRulesResponseBodyRulesRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetTrafficLimitConfig(v *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) *ListRulesResponseBodyRulesRuleActions {
	s.TrafficLimitConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetTrafficMirrorConfig(v *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) *ListRulesResponseBodyRulesRuleActions {
	s.TrafficMirrorConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetType(v string) *ListRulesResponseBodyRulesRuleActions {
	s.Type = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsCorsConfig struct {
	// Indicates whether credentials can be carried in CORS requests. Valid values:
	//
	// *   **on**
	// *   **off**
	AllowCredentials *string `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	// The allowed headers for CORS requests.
	AllowHeaders []*string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty" type:"Repeated"`
	// The allowed HTTP methods for CORS requests.
	AllowMethods []*string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty" type:"Repeated"`
	// The origin sites that are allowed.
	AllowOrigin []*string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty" type:"Repeated"`
	// The headers that can be exposed.
	ExposeHeaders []*string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty" type:"Repeated"`
	// The maximum cache time of preflight requests in the browser. Unit: seconds.
	//
	// Valid values: **-1** to **172800**.
	MaxAge *int64 `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsCorsConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsCorsConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetAllowCredentials(v string) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetAllowHeaders(v []*string) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.AllowHeaders = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetAllowMethods(v []*string) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.AllowMethods = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetAllowOrigin(v []*string) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.AllowOrigin = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetExposeHeaders(v []*string) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.ExposeHeaders = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetMaxAge(v int64) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.MaxAge = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsFixedResponseConfig struct {
	// The content of the custom response. The content is up to 1 KB in size, and can contain only ASCII characters.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The format of the response.
	//
	// Valid values: **text/plain**, **text/css**, **text/html**, **application/javascript**, and **application/json**.
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The HTTP status code in the response. Valid values: **HTTP\_2xx**, **HTTP\_4xx**, and **HTTP\_5xx**. **x** must be a digit.
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) SetContent(v string) *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) SetContentType(v string) *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) SetHttpCode(v string) *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsForwardGroupConfig struct {
	// The server groups to which requests are forwarded.
	ServerGroupTuples []*ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	// The server group to which requests are forwarded.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The weight. Valid values: **0** to **100**.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) SetWeight(v int32) *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig struct {
	// The key of the header to be inserted. The key must be 1 to 40 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The key specified in `InsertHeader` must be unique.
	//
	// > : **Cookie** and **Host** are not supported.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the header to be inserted.
	//
	// *   If **ValueType** is set to **SystemDefined**, one of the following header values is supported:
	//
	//     *   **ClientSrcPort**: the client port.
	//     *   **ClientSrcIp**: the client IP address.
	//     *   **Protocol**: the request protocol (HTTP or HTTPS).
	//     *   **SLBId**: the ALB instance ID.
	//     *   **SLBPort**: the listener port.
	//
	// *   If **ValueType** is set to **UserDefined**, you can specify a custom header value. The header value must be 1 to 128 characters in length, and can contain printable characters whose ASCII values are `greater than or equal to 32 and lower than 127`. You can use asterisks (\*) and question marks (?) as wildcards. The value cannot start or end with a space character.
	//
	// *   If **ValueType** is set to **ReferenceHeader**, you can reference one of the request headers. The header value must be 1 to 128 characters in length, and can contain lowercase letters, digits, underscores (\_), and hyphens (-).
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The type of the header. Valid values:
	//
	// *   **UserDefined:** a user-defined header.
	// *   **ReferenceHeader:** a header that is referenced from a request header.
	// *   **SystemDefined:** a system-defined header.
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) SetKey(v string) *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) SetValue(v string) *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) SetValueType(v string) *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsRedirectConfig struct {
	// The hostname to which requests are redirected. Valid values:
	//
	// *   **${host}** (default): If you set the value to ${host}, you cannot append other characters.
	//
	// *   A custom value that meets the following requirements:
	//
	//     *   The hostname is 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), and periods (.). Asterisks (\*) and question marks (?) can be used as wildcards.
	//     *   The hostname contains at least one period (.) but does not start or end with a period (.).
	//     *   The rightmost domain label contains only letters and wildcard characters. It does not contain digits or hyphens (-).
	//     *   The domain labels do not start or end with hyphens (-).
	//     *   You can use asterisks (\*) and question marks (?) anywhere in a domain label as wildcard characters.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The redirect type. Valid values: **301**, **302**, **303**, **307**, and **308**.
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The path to which requests are redirected. Valid values:
	//
	// *   **${path}** (default): You can reference \*\*${host}**, **${protocol}**, and**${port}**. The path can consist of **${host}**,**${protocol}**, and **${port}\*\*. Each variable can be used only once. You can specify one or more of the preceding variables in each request. You can also combine them with a custom value.
	//
	// *   A custom value that meets the following requirements:
	//
	//     *   The value is 1 to 128 characters in length.
	//     *   It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ .+ / & ~ @ :`. It cannot contain the following special characters: `" % # ; ! ( ) [ ] ^ , "`. You can use asterisks (\*) and question marks (?) as wildcards.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port to which requests are redirected. Valid values:
	//
	// *   **${port}** (default): If you set the value to ${port}, you cannot append other characters.
	// *   You can also enter a port number. Valid values: **1 to 63335**.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The redirect protocol. Valid values:
	//
	// *   **${protocol}** (default): If you set the value to ${protocol}, you cannot append other characters.
	// *   **HTTP** or **HTTPS**.
	//
	// > HTTPS listeners support only HTTP to HTTPS redirection.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The query string to which requests are forwarded. The query string must be 1 to 128 characters in length, and can contain printable characters, excluding uppercase letters and the following special characters: `# [ ] { } \ | < > &`.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsRedirectConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetHost(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetHttpCode(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetPath(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetPort(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetProtocol(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetQuery(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig struct {
	// 去除的头字段名称，长度为1\~40个字符，支持大小写字母a~z、数字、下划线（_）和短划线（-）。头字段名称不能重复用于RemoveHeader中。
	//
	// * 请求方向（Direction取值为Request）：不允许将头名称设置为以下字段（不区分大小写）：`slb-id`、`slb-ip`、`x-forwarded-for`、`x-forwarded-proto`、`x-forwarded-eip`、`x-forwarded-port`、`x-forwarded-client-srcport`、`connection`、`upgrade`、`content-length`、`transfer-encoding`、`keep-alive`、`te`、`host`、`cookie`、`remoteip`、`authority`。
	// * 响应方向（Direction取值为Response）：响应方向不允许将头名称设置为以下字段（不区分大小写）：`connection`、`upgrade`、`content-length`、`transfer-encoding`。
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig) SetKey(v string) *ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig {
	s.Key = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsRewriteConfig struct {
	// The hostname to which requests are redirected. Valid values:
	//
	// *   **${host}** (default): If you set the value to ${host}, you cannot append other characters.
	//
	// *   A custom value that meets the following requirements:
	//
	//     *   The hostname is 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), and periods (.). Asterisks (\*) and question marks (?) can be used as wildcards.
	//     *   The hostname contains at least one period (.) but does not start or end with a period (.).
	//     *   The rightmost domain label contains only letters and wildcard characters. It does not contain digits or hyphens (-).
	//     *   The domain labels do not start or end with hyphens (-).
	//     *   You can use asterisks (\*) and question marks (?) anywhere in a domain label as wildcard characters.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The path to which requests are forwarded. The path is 1 to 128 characters in length and starts with a forward slash (/). The path can contain letters, digits, asterisks (\*), question marks (?), and the following special characters: `$ - _ . + / & ~ @ :`. The path does not contain the following special characters: `" % # ; ! ( ) [ ] ^ , "`.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The query string to which requests are forwarded. The query string must be 1 to 128 characters in length, and can contain printable characters, excluding uppercase letters and the following special characters: `# [ ] { } \ | < > &`.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsRewriteConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) SetHost(v string) *ListRulesResponseBodyRulesRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) SetPath(v string) *ListRulesResponseBodyRulesRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) SetQuery(v string) *ListRulesResponseBodyRulesRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig struct {
	// The QPS of each IP address. Valid values: **1 to 100000**.
	//
	// > If you specify this parameter and **QPS**, the value of **PerIpQps** must be smaller than the value of **QPS**.
	PerIpQps *int32 `json:"PerIpQps,omitempty" xml:"PerIpQps,omitempty"`
	// The number of queries per second (QPS). Valid values: **1** to **100000**.
	QPS *int32 `json:"QPS,omitempty" xml:"QPS,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) SetPerIpQps(v int32) *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig {
	s.PerIpQps = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) SetQPS(v int32) *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig {
	s.QPS = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig struct {
	// The server groups to which network traffic is mirrored.
	MirrorGroupConfig *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig `json:"MirrorGroupConfig,omitempty" xml:"MirrorGroupConfig,omitempty" type:"Struct"`
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) SetMirrorGroupConfig(v *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig {
	s.MirrorGroupConfig = v
	return s
}

type ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig struct {
	// The server group to which network traffic is mirrored.
	ServerGroupTuples []*ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) SetServerGroupTuples(v []*ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples struct {
	// The server group ID.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The weight. Valid values: **0** to **100**.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetServerGroupId(v string) *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetWeight(v int32) *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

type ListRulesResponseBodyRulesRuleConditions struct {
	// The configuration of the cookie.
	CookieConfig *ListRulesResponseBodyRulesRuleConditionsCookieConfig `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	// The configuration of the header.
	HeaderConfig *ListRulesResponseBodyRulesRuleConditionsHeaderConfig `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	// The configurations of the hosts.
	HostConfig *ListRulesResponseBodyRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// The configurations of the request methods.
	MethodConfig *ListRulesResponseBodyRulesRuleConditionsMethodConfig `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	// The configurations of the paths.
	PathConfig *ListRulesResponseBodyRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// The configurations of the query strings.
	QueryStringConfig *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	// 响应HTTP头部配置。
	ResponseHeaderConfig *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig `json:"ResponseHeaderConfig,omitempty" xml:"ResponseHeaderConfig,omitempty" type:"Struct"`
	// 响应状态码配置。
	ResponseStatusCodeConfig *ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig `json:"ResponseStatusCodeConfig,omitempty" xml:"ResponseStatusCodeConfig,omitempty" type:"Struct"`
	// The configuration of the source IP addresses based on which user traffic is matched.
	SourceIpConfig *ListRulesResponseBodyRulesRuleConditionsSourceIpConfig `json:"SourceIpConfig,omitempty" xml:"SourceIpConfig,omitempty" type:"Struct"`
	// The type of forwarding rule. Valid values:
	//
	// *   **Host**: Requests are distributed based on hosts.
	// *   **Path**: Requests are distributed based on paths.
	// *   **Header**: Requests are distributed based on HTTP headers.
	// *   **QueryString**: Requests are distributed based on query strings.
	// *   **Method**: Requests are distributed based on request methods.
	// *   **Cookie**: Requests are distributed based on cookies.
	// *   **SourceIp**: Requests are distributed based on source IP addresses.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetCookieConfig(v *ListRulesResponseBodyRulesRuleConditionsCookieConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetHeaderConfig(v *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetHostConfig(v *ListRulesResponseBodyRulesRuleConditionsHostConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetMethodConfig(v *ListRulesResponseBodyRulesRuleConditionsMethodConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetPathConfig(v *ListRulesResponseBodyRulesRuleConditionsPathConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetQueryStringConfig(v *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetResponseHeaderConfig(v *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.ResponseHeaderConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetResponseStatusCodeConfig(v *ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.ResponseStatusCodeConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetSourceIpConfig(v *ListRulesResponseBodyRulesRuleConditionsSourceIpConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.SourceIpConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetType(v string) *ListRulesResponseBodyRulesRuleConditions {
	s.Type = &v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsCookieConfig struct {
	// The values of the cookie.
	Values []*ListRulesResponseBodyRulesRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfig) SetValues(v []*ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) *ListRulesResponseBodyRulesRuleConditionsCookieConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsCookieConfigValues struct {
	// The key of the cookie. The key is 1 to 100 characters in length, and can contain printable characters such as lowercase letters, asterisks (\*), and question marks (?). The key cannot contain uppercase letters, space characters, or the following special characters: `# [ ] { } \ | < > &`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the cookie. The value is 1 to 128 characters in length, and can contain printable characters such as lowercase letters, asterisks (\*), and question marks (?). Uppercase letters, space characters, and the following special characters are not supported: `# [ ] { } \ | < > &`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) SetKey(v string) *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) SetValue(v string) *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsHeaderConfig struct {
	// The key of the header. The key must be 1 to 40 characters in length. It can contain letters, digits, hyphens (-), and underscores (\_). Cookie and Host are not supported.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of the header.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) SetKey(v string) *ListRulesResponseBodyRulesRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsHostConfig struct {
	// The hostnames.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsHostConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsMethodConfig struct {
	// The request methods.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsMethodConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsMethodConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsMethodConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsPathConfig struct {
	// The paths.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsPathConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsQueryStringConfig struct {
	// The query strings.
	Values []*ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) SetValues(v []*ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues struct {
	// They key of the query string. The key must be 1 to 100 characters in length, and can contain printable characters such as lowercase letters, asterisks (\*), and question marks (?). The key cannot contain uppercase letters, space characters, or the following special characters: `# [ ] { } \ | < > &`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the query string. The value must be 1 to 128 characters in length, and can contain printable characters such as lowercase letters, asterisks (\*), and question marks (?). However, uppercase letters, space characters, and the following special characters are not supported: `# [ ] { } \ | < > &`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) SetKey(v string) *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) SetValue(v string) *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig struct {
	// 响应HTTP头部键。长度为1\~40个字符。支持字母a~z、数字、短划线（-）和下划线（_）。不支持Cookie和Host。
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 响应HTTP头部值列表。
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) SetKey(v string) *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig struct {
	// 响应状态码列表。
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsSourceIpConfig struct {
	// The source IP addresses.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsSourceIpConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsSourceIpConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsSourceIpConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsSourceIpConfig {
	s.Values = v
	return s
}

type ListRulesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRulesResponse) SetHeaders(v map[string]*string) *ListRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRulesResponse) SetStatusCode(v int32) *ListRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRulesResponse) SetBody(v *ListRulesResponseBody) *ListRulesResponse {
	s.Body = v
	return s
}

type ListSecurityPoliciesRequest struct {
	// The number of entries per page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// *   You do not need to specify this parameter for the first request.
	// *   You must specify the token that is obtained from the previous query as the value of **NextToken**.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security policy IDs. You can specify up to 20 IDs.
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitempty" xml:"SecurityPolicyIds,omitempty" type:"Repeated"`
	// The names of the security policies. You can specify up to 10 names.
	SecurityPolicyNames []*string `json:"SecurityPolicyNames,omitempty" xml:"SecurityPolicyNames,omitempty" type:"Repeated"`
}

func (s ListSecurityPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesRequest) SetMaxResults(v int32) *ListSecurityPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecurityPoliciesRequest) SetNextToken(v string) *ListSecurityPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecurityPoliciesRequest) SetResourceGroupId(v string) *ListSecurityPoliciesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecurityPoliciesRequest) SetSecurityPolicyIds(v []*string) *ListSecurityPoliciesRequest {
	s.SecurityPolicyIds = v
	return s
}

func (s *ListSecurityPoliciesRequest) SetSecurityPolicyNames(v []*string) *ListSecurityPoliciesRequest {
	s.SecurityPolicyNames = v
	return s
}

type ListSecurityPoliciesResponseBody struct {
	// The number of entries per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// *   If **NextToken** is empty, no next page exists.
	// *   If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The supported security policies.
	SecurityPolicies []*ListSecurityPoliciesResponseBodySecurityPolicies `json:"SecurityPolicies,omitempty" xml:"SecurityPolicies,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecurityPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesResponseBody) SetMaxResults(v int32) *ListSecurityPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetNextToken(v string) *ListSecurityPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetRequestId(v string) *ListSecurityPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetSecurityPolicies(v []*ListSecurityPoliciesResponseBodySecurityPolicies) *ListSecurityPoliciesResponseBody {
	s.SecurityPolicies = v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetTotalCount(v int32) *ListSecurityPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type ListSecurityPoliciesResponseBodySecurityPolicies struct {
	// The supported cipher suites.
	Ciphers    []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	CreateTime *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The resource group ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security policy ID.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The name of the security policy.
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	// The status of the security policy. Valid values:
	//
	// *   **Configuring**
	// *   **Available**
	SecurityPolicyStatus *string `json:"SecurityPolicyStatus,omitempty" xml:"SecurityPolicyStatus,omitempty"`
	// The supported TLS protocol versions.
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s ListSecurityPoliciesResponseBodySecurityPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPoliciesResponseBodySecurityPolicies) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetCiphers(v []*string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.Ciphers = v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetCreateTime(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetResourceGroupId(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyId(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyName(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyName = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyStatus(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyStatus = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetTLSVersions(v []*string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.TLSVersions = v
	return s
}

type ListSecurityPoliciesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSecurityPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecurityPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesResponse) SetHeaders(v map[string]*string) *ListSecurityPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityPoliciesResponse) SetStatusCode(v int32) *ListSecurityPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityPoliciesResponse) SetBody(v *ListSecurityPoliciesResponseBody) *ListSecurityPoliciesResponse {
	s.Body = v
	return s
}

type ListSecurityPolicyRelationsRequest struct {
	// The listeners that associated with the security policy.
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitempty" xml:"SecurityPolicyIds,omitempty" type:"Repeated"`
}

func (s ListSecurityPolicyRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsRequest) SetSecurityPolicyIds(v []*string) *ListSecurityPolicyRelationsRequest {
	s.SecurityPolicyIds = v
	return s
}

type ListSecurityPolicyRelationsResponseBody struct {
	// The ID of the listener.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The port that is used by the listener.
	SecrityPolicyRelations []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations `json:"SecrityPolicyRelations,omitempty" xml:"SecrityPolicyRelations,omitempty" type:"Repeated"`
}

func (s ListSecurityPolicyRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponseBody) SetRequestId(v string) *ListSecurityPolicyRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBody) SetSecrityPolicyRelations(v []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) *ListSecurityPolicyRelationsResponseBody {
	s.SecrityPolicyRelations = v
	return s
}

type ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations struct {
	// The protocol that is used by the listener.
	RelatedListeners []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Repeated"`
	SecurityPolicyId *string                                                                          `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) SetRelatedListeners(v []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations {
	s.RelatedListeners = v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) SetSecurityPolicyId(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations {
	s.SecurityPolicyId = &v
	return s
}

type ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners struct {
	// The ID of the SLB instance.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the security policy.
	ListenerPort     *int64  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId   *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetListenerId(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.ListenerId = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetListenerPort(v int64) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetListenerProtocol(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetLoadBalancerId(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.LoadBalancerId = &v
	return s
}

type ListSecurityPolicyRelationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSecurityPolicyRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecurityPolicyRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponse) SetHeaders(v map[string]*string) *ListSecurityPolicyRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityPolicyRelationsResponse) SetStatusCode(v int32) *ListSecurityPolicyRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponse) SetBody(v *ListSecurityPolicyRelationsResponseBody) *ListSecurityPolicyRelationsResponse {
	s.Body = v
	return s
}

type ListServerGroupServersRequest struct {
	// The number of entries to return. Valid values: **1** to **100**. If you do not specify a value, the default value **20** is used.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no next queries are to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the value to the value of **NextToken** that is returned from the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the server group.
	ServerGroupId *string                             `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	ServerIds     []*string                           `json:"ServerIds,omitempty" xml:"ServerIds,omitempty" type:"Repeated"`
	Tag           []*ListServerGroupServersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServerGroupServersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersRequest) SetMaxResults(v int32) *ListServerGroupServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersRequest) SetNextToken(v string) *ListServerGroupServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerGroupId(v string) *ListServerGroupServersRequest {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerIds(v []*string) *ListServerGroupServersRequest {
	s.ServerIds = v
	return s
}

func (s *ListServerGroupServersRequest) SetTag(v []*ListServerGroupServersRequestTag) *ListServerGroupServersRequest {
	s.Tag = v
	return s
}

type ListServerGroupServersRequestTag struct {
	// The tag key. You can specify up to 10 tag keys.
	//
	// It can be at most 64 characters in length, and cannot contain `http://` or `https://`. It must not start with `aliyun` or `acs:`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 10 tag values.
	//
	// It can be up to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServerGroupServersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersRequestTag) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersRequestTag) SetKey(v string) *ListServerGroupServersRequestTag {
	s.Key = &v
	return s
}

func (s *ListServerGroupServersRequestTag) SetValue(v string) *ListServerGroupServersRequestTag {
	s.Value = &v
	return s
}

type ListServerGroupServersResponseBody struct {
	// The maximum number of entries returned.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no subsequent query is to be sent.
	// *   If a value is returned for **NextToken**, the value is the token that is used for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of backend servers.
	Servers []*ListServerGroupServersResponseBodyServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBody) SetMaxResults(v int32) *ListServerGroupServersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetNextToken(v string) *ListServerGroupServersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetRequestId(v string) *ListServerGroupServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetServers(v []*ListServerGroupServersResponseBodyServers) *ListServerGroupServersResponseBody {
	s.Servers = v
	return s
}

func (s *ListServerGroupServersResponseBody) SetTotalCount(v int32) *ListServerGroupServersResponseBody {
	s.TotalCount = &v
	return s
}

type ListServerGroupServersResponseBodyServers struct {
	// The description of the backend server.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port used by the backend server. Valid values: **1** to **65535**.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Indicates whether the remote IP address feature is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	RemoteIpEnabled *bool `json:"RemoteIpEnabled,omitempty" xml:"RemoteIpEnabled,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The ID of the backend server.
	//
	// >  If **ServerType** is set to **Fc**, **ServerId** is the ARN of a function.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address in inclusive ENI mode.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server.
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The status of the backend server. Valid values:
	//
	// *   **Adding**: The backend server is being added.
	// *   **Available**: The backend server is added.
	// *   **Configuring**: The backend server is being configured.
	// *   **Removing**: The backend server is being removed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The weight of the backend server. An ECS instance with a higher weight receives more requests.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListServerGroupServersResponseBodyServers) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponseBodyServers) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBodyServers) SetDescription(v string) *ListServerGroupServersResponseBodyServers {
	s.Description = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetPort(v int32) *ListServerGroupServersResponseBodyServers {
	s.Port = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetRemoteIpEnabled(v bool) *ListServerGroupServersResponseBodyServers {
	s.RemoteIpEnabled = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerGroupId(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerId(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerId = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerIp(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerIp = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerType(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerType = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetStatus(v string) *ListServerGroupServersResponseBodyServers {
	s.Status = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetWeight(v int32) *ListServerGroupServersResponseBodyServers {
	s.Weight = &v
	return s
}

type ListServerGroupServersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServerGroupServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServerGroupServersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponse) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponse) SetHeaders(v map[string]*string) *ListServerGroupServersResponse {
	s.Headers = v
	return s
}

func (s *ListServerGroupServersResponse) SetStatusCode(v int32) *ListServerGroupServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerGroupServersResponse) SetBody(v *ListServerGroupServersResponseBody) *ListServerGroupServersResponse {
	s.Body = v
	return s
}

type ListServerGroupsRequest struct {
	// The number of entries to return on each page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no next query is to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the value to the value of **NextToken** that is returned from the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group.
	ResourceGroupId  *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerGroupIds   []*string                     `json:"ServerGroupIds,omitempty" xml:"ServerGroupIds,omitempty" type:"Repeated"`
	ServerGroupNames []*string                     `json:"ServerGroupNames,omitempty" xml:"ServerGroupNames,omitempty" type:"Repeated"`
	Tag              []*ListServerGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupsRequest) SetMaxResults(v int32) *ListServerGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsRequest) SetNextToken(v string) *ListServerGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsRequest) SetResourceGroupId(v string) *ListServerGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupIds(v []*string) *ListServerGroupsRequest {
	s.ServerGroupIds = v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupNames(v []*string) *ListServerGroupsRequest {
	s.ServerGroupNames = v
	return s
}

func (s *ListServerGroupsRequest) SetTag(v []*ListServerGroupsRequestTag) *ListServerGroupsRequest {
	s.Tag = v
	return s
}

func (s *ListServerGroupsRequest) SetVpcId(v string) *ListServerGroupsRequest {
	s.VpcId = &v
	return s
}

type ListServerGroupsRequestTag struct {
	// The tag keys. You can specify up to 10 tags.
	//
	// It can be at most 64 characters in length, and cannot contain `http://` or `https://`. It must not start with `aliyun` or `acs:`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values. You can specify up to 10 tags.
	//
	// It can be at most 128 characters in length, and cannot contain `http://` or `https://`. It must not start with `aliyun` or `acs:`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServerGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *ListServerGroupsRequestTag) SetKey(v string) *ListServerGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *ListServerGroupsRequestTag) SetValue(v string) *ListServerGroupsRequestTag {
	s.Value = &v
	return s
}

type ListServerGroupsResponseBody struct {
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no next query is to be sent.
	// *   If a value of **NextToken** is returned, the value is the token that is used for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of server groups.
	ServerGroups []*ListServerGroupsResponseBodyServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// The number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBody) SetMaxResults(v int32) *ListServerGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetNextToken(v string) *ListServerGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetRequestId(v string) *ListServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetServerGroups(v []*ListServerGroupsResponseBodyServerGroups) *ListServerGroupsResponseBody {
	s.ServerGroups = v
	return s
}

func (s *ListServerGroupsResponseBody) SetTotalCount(v int32) *ListServerGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListServerGroupsResponseBodyServerGroups struct {
	// Indicates whether configuration management is enabled. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	ConfigManagedEnabled *bool   `json:"ConfigManagedEnabled,omitempty" xml:"ConfigManagedEnabled,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The health check configurations.
	HealthCheckConfig *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// Indicates whether IPv6 is supported. Valid values:
	//
	// *   **true**: supported
	// *   **false**: not supported
	Ipv6Enabled *bool `json:"Ipv6Enabled,omitempty" xml:"Ipv6Enabled,omitempty"`
	// The backend protocol. Valid values:
	//
	// *   **HTTP**: allows you to associate an HTTPS, HTTP, or QUIC listener with the server group.
	// *   **HTTPS**: allows you to associate HTTPS listeners with backend servers.
	// *   **GRPC**: If you select this option, you can associate the server group with HTTPS and QUIC listeners.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The routing algorithm. Valid values:
	//
	// *   **Wrr**: Backend servers with higher weights receive more requests than backend servers with lower weights.
	// *   **Wlc**: Requests are distributed based on the weight and load of each backend server. The load refers to the number of connections on a backend server. If multiple backend servers have the same weight, requests are forwarded to the backend server with the least connections.
	// *   **Sch**: enables consistent hashing. Requests from the same source IP address are distributed to the same backend server.
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of backend servers in the server group.
	ServerCount *int32 `json:"ServerCount,omitempty" xml:"ServerCount,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The name of the server group.
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The status of the server group. Valid values:
	//
	// *   **Creating**: The server group is being created.
	// *   **Available**: The server group is available
	// *   **Configuring**: The server group is being configured.
	ServerGroupStatus *string `json:"ServerGroupStatus,omitempty" xml:"ServerGroupStatus,omitempty"`
	// The type of the server group. Valid values:
	//
	// *   **Instance**: a server group of the Instance type.
	// *   **Ip**: a server group of the IP type.
	// *   **Fc**: a server group of the Function Compute type.
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// The name of the server group.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The configuration of session persistence.
	StickySessionConfig *ListServerGroupsResponseBodyServerGroupsStickySessionConfig `json:"StickySessionConfig,omitempty" xml:"StickySessionConfig,omitempty" type:"Struct"`
	// The tags that are added to the server group.
	Tags []*ListServerGroupsResponseBodyServerGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// url一致性hash参数配置
	UchConfig *ListServerGroupsResponseBodyServerGroupsUchConfig `json:"UchConfig,omitempty" xml:"UchConfig,omitempty" type:"Struct"`
	// Indicates whether persistent TCP connections are enabled. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	UpstreamKeepaliveEnabled *bool `json:"UpstreamKeepaliveEnabled,omitempty" xml:"UpstreamKeepaliveEnabled,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroups) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroups) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroups) SetConfigManagedEnabled(v bool) *ListServerGroupsResponseBodyServerGroups {
	s.ConfigManagedEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetCreateTime(v string) *ListServerGroupsResponseBodyServerGroups {
	s.CreateTime = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetHealthCheckConfig(v *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) *ListServerGroupsResponseBodyServerGroups {
	s.HealthCheckConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetIpv6Enabled(v bool) *ListServerGroupsResponseBodyServerGroups {
	s.Ipv6Enabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetProtocol(v string) *ListServerGroupsResponseBodyServerGroups {
	s.Protocol = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetResourceGroupId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetScheduler(v string) *ListServerGroupsResponseBodyServerGroups {
	s.Scheduler = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerCount(v int32) *ListServerGroupsResponseBodyServerGroups {
	s.ServerCount = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupName(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupName = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupStatus(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupStatus = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupType(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupType = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServiceName(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServiceName = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetStickySessionConfig(v *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) *ListServerGroupsResponseBodyServerGroups {
	s.StickySessionConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetTags(v []*ListServerGroupsResponseBodyServerGroupsTags) *ListServerGroupsResponseBodyServerGroups {
	s.Tags = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetUchConfig(v *ListServerGroupsResponseBodyServerGroupsUchConfig) *ListServerGroupsResponseBodyServerGroups {
	s.UchConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetUpstreamKeepaliveEnabled(v bool) *ListServerGroupsResponseBodyServerGroups {
	s.UpstreamKeepaliveEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetVpcId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.VpcId = &v
	return s
}

type ListServerGroupsResponseBodyServerGroupsHealthCheckConfig struct {
	// The HTTP status codes that indicate a successful health check.
	//
	// *   If **HealthCheckProtocol** is set to **HTTP**, **HealthCheckCodes** can be set to **http\_2xx**, **http\_3xx**, **http\_4xx**, and **http\_5xx**. Separate multiple HTTP status codes with commas (,).
	// *   If **HealthCheckProtocol** is set to **gRPC**, **HealthCheckCodes** can be set to **0 to 99**. Value ranges are supported. You can enter at most 20 value ranges and must separate them with commas (,).
	//
	// >  This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The backend port that is used for health checks. Valid values: **0** to **65535**.
	//
	// **0** indicates that the port on a backend server is used for health checks.
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The domain name that is used for health checks. The domain name meets the following requirements:
	//
	// *   The domain name is 1 to 80 characters in length.
	// *   The domain name can contain lowercase letters, digits, hyphens (-), and periods (.).
	// *   It contains at least one period (.) but does not start or end with a period (.).
	// *   The rightmost domain label of the domain name contains only letters, and does not contain digits or hyphens (-).
	// *   The domain name does not start or end with a hyphen (-).
	//
	// >  This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP**.
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The HTTP version that is used for health checks.
	//
	// Valid values: **HTTP1.0** and **HTTP1.1**.
	//
	// >  This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP**.
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds. Valid values: **1** to **50**.
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP method that is used for health checks. Valid values:
	//
	// *   **GET**: If the length of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	// *   **POST**: By default, gRPC health checks use the POST method.
	// *   **HEAD**: By default, HTTP health checks use the HEAD method.
	//
	// >  This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The path that is used for health checks.
	//
	// >  This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP**.
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// *   **HTTP**: To perform HTTP health checks, ALB sends HEAD or GET requests to a backend server to check whether the backend server is healthy.
	// *   **TCP**: To perform TCP health checks, ALB sends SYN packets to a backend server to check whether the port of the backend server is available to receive requests.
	// *   **gRPC**: To perform gRPC health checks, ALB sends POST or GET requests to a backend server to check whether the backend server is healthy.
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The timeout period of a health check. If a backend server does not respond within the specified timeout period, the backend server fails the health check. Unit: seconds.
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail** to **success**.
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success** to **fail**.
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckCodes(v []*string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckCodes = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckConnectPort(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckEnabled(v bool) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckHost(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckHost = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckHttpVersion(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckInterval(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckMethod(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckPath(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckPath = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckProtocol(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckProtocol = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckTimeout(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthyThreshold(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetUnhealthyThreshold(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

type ListServerGroupsResponseBodyServerGroupsStickySessionConfig struct {
	// The cookie that is configured on the backend server.
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of a cookie. Unit: seconds. Valid values: **1** to **86400**.
	//
	// >  This parameter takes effect when the **StickySessionEnabled** parameter is set to **true** and the **StickySessionType** parameter is set to **Insert**.
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// Indicates whether session persistence is enabled. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	StickySessionEnabled *bool `json:"StickySessionEnabled,omitempty" xml:"StickySessionEnabled,omitempty"`
	// The method that is used to handle a cookie. Valid values:
	//
	// *   **Insert**: inserts a cookie.
	//
	//     ALB inserts a cookie (SERVERID) into the first HTTP or HTTPS response packet that is sent to a client. The next request from the client contains this cookie and the listener forwards this request to the recorded backend server.
	//
	// *   **Server**: rewrites a cookie.
	//
	//     When ALB detects a user-defined cookie, it overwrites the original cookie with the user-defined cookie. The next request from the client carries the user-defined cookie, and the listener will distribute this request to the recorded backend server.
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsStickySessionConfig) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsStickySessionConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) SetCookie(v string) *ListServerGroupsResponseBodyServerGroupsStickySessionConfig {
	s.Cookie = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) SetCookieTimeout(v int32) *ListServerGroupsResponseBodyServerGroupsStickySessionConfig {
	s.CookieTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) SetStickySessionEnabled(v bool) *ListServerGroupsResponseBodyServerGroupsStickySessionConfig {
	s.StickySessionEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) SetStickySessionType(v string) *ListServerGroupsResponseBodyServerGroupsStickySessionConfig {
	s.StickySessionType = &v
	return s
}

type ListServerGroupsResponseBodyServerGroupsTags struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsTags) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsTags) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) SetKey(v string) *ListServerGroupsResponseBodyServerGroupsTags {
	s.Key = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) SetValue(v string) *ListServerGroupsResponseBodyServerGroupsTags {
	s.Value = &v
	return s
}

type ListServerGroupsResponseBodyServerGroupsUchConfig struct {
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 一致性hash参数值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsUchConfig) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsUchConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsUchConfig) SetType(v string) *ListServerGroupsResponseBodyServerGroupsUchConfig {
	s.Type = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsUchConfig) SetValue(v string) *ListServerGroupsResponseBodyServerGroupsUchConfig {
	s.Value = &v
	return s
}

type ListServerGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponse) SetHeaders(v map[string]*string) *ListServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListServerGroupsResponse) SetStatusCode(v int32) *ListServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerGroupsResponse) SetBody(v *ListServerGroupsResponseBody) *ListServerGroupsResponse {
	s.Body = v
	return s
}

type ListSystemSecurityPoliciesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security policies.
	SecurityPolicies []*ListSystemSecurityPoliciesResponseBodySecurityPolicies `json:"SecurityPolicies,omitempty" xml:"SecurityPolicies,omitempty" type:"Repeated"`
}

func (s ListSystemSecurityPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponseBody) SetRequestId(v string) *ListSystemSecurityPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) SetSecurityPolicies(v []*ListSystemSecurityPoliciesResponseBodySecurityPolicies) *ListSystemSecurityPoliciesResponseBody {
	s.SecurityPolicies = v
	return s
}

type ListSystemSecurityPoliciesResponseBodySecurityPolicies struct {
	// The supported cipher suites, which are determined by the **TLS protocol version**.
	//
	// The specified cipher suites must be supported by at least one **TLS protocol version** that you select. For example, if you set the TLSVersions.N parameter to TLSv1.3, you can specify only cipher suites that are supported by TLSv1.3.
	//
	// *   TLS 1.0 and TLS 1.1 support the following cipher suites:
	//
	//     *   ECDHE-ECDSA-AES128-SHA
	//     *   ECDHE-ECDSA-AES256-SHA
	//     *   ECDHE-RSA-AES128-SHA
	//     *   ECDHE-RSA-AES256-SHA
	//     *   AES128-SHA
	//     *   AES256-SHA
	//
	//     <!---->
	//
	//     *   DES-CBC3-SHA
	//
	// *   TLS 1.2 supports the following cipher suites:
	//
	//     *   ECDHE-ECDSA-AES128-SHA
	//     *   ECDHE-ECDSA-AES256-SHA
	//     *   ECDHE-RSA-AES128-SHA
	//     *   ECDHE-RSA-AES256-SHA
	//     *   AES128-SHA
	//     *   AES256-SHA
	//     *   DES-CBC3-SHA
	//     *   ECDHE-ECDSA-AES128-GCM-SHA256
	//     *   ECDHE-ECDSA-AES256-GCM-SHA384
	//     *   ECDHE-ECDSA-AES128-SHA256
	//     *   ECDHE-ECDSA-AES256-SHA384
	//     *   ECDHE-RSA-AES128-GCM-SHA256
	//     *   ECDHE-RSA-AES256-GCM-SHA384
	//     *   ECDHE-RSA-AES128-SHA256
	//     *   ECDHE-RSA-AES256-SHA384
	//     *   AES128-GCM-SHA256
	//     *   AES256-GCM-SHA384
	//     *   AES128-SHA256
	//     *   AES256-SHA256
	//
	// *   TLS 1.3 supports the following cipher suites:
	//
	//     *   TLS_AES\_128\_GCM_SHA256
	//     *   TLS_AES\_256\_GCM_SHA384
	//     *   TLS_CHACHA20\_POLY1305\_SHA256
	//     *   TLS_AES\_128\_CCM_SHA256
	//     *   TLS_AES\_128\_CCM\_8\_SHA256
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The ID of the security policy.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The TLS protocol versions that are supported. Valid values: **TLSv1.0**, **TLSv1.1**, **TLSv1.2**, and **TLSv1.3**.
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s ListSystemSecurityPoliciesResponseBodySecurityPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponseBodySecurityPolicies) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetCiphers(v []*string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.Ciphers = v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyId(v string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetTLSVersions(v []*string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.TLSVersions = v
	return s
}

type ListSystemSecurityPoliciesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSystemSecurityPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSystemSecurityPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponse) SetHeaders(v map[string]*string) *ListSystemSecurityPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListSystemSecurityPoliciesResponse) SetStatusCode(v int32) *ListSystemSecurityPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponse) SetBody(v *ListSystemSecurityPoliciesResponseBody) *ListSystemSecurityPoliciesResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// The type of the tag.
	//
	// Valid values: **Custom**, **System**, and **All**.
	//
	// Default value: **All**.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The tag key. It can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries to return on each page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no next query is to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the value to the value of **NextToken** that is returned from the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the resource. Valid values:
	//
	// *   **acl**: a network access control list (ACL)
	// *   **loadbalancer**: an ALB instance
	// *   **securitypolicy**: a security policy
	// *   **servergroup**: a server group
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetCategory(v string) *ListTagKeysRequest {
	s.Category = &v
	return s
}

func (s *ListTagKeysRequest) SetKeyword(v string) *ListTagKeysRequest {
	s.Keyword = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no next query is to be sent.
	// *   If a value of **NextToken** is returned, the value is the token that is used for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tag keys.
	TagKeys []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// The number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetMaxResults(v int32) *ListTagKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagKeysResponseBodyTagKeys struct {
	// The type of the tag.
	//
	// Valid values: **Custom**, **System**, and **All**.
	//
	// Default value: **All**.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The tags that match all of the filter conditions.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) SetCategory(v string) *ListTagKeysResponseBodyTagKeys {
	s.Category = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagKey(v string) *ListTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The ID of the resource.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The type of the resource. Valid values:
	//
	// *   **acl**: a network access control list (ACL)
	// *   **loadbalancer**: an Application Load Balancer (ALB) instance
	// *   **securitypolicy**: a security policy
	// *   **servergroup**: a server group
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tag key. It can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The tags.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag value. It can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The number of entries returned per page.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no next query is to be sent.
	// *   If a value of **NextToken** is returned, the value is the token that is used for the next query.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// The ID of the request.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The tags that match the specified keys and values.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the resource. Valid values:
	//
	// *   **acl**: a network ACL
	// *   **loadbalancer**: an ALB instance
	// *   **securitypolicy**: a security policy
	// *   **servergroup**: a server group
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetMaxResults(v int32) *ListTagResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The tag key.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The tag value.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	// The ID of the resource.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The type of the resource. Valid values:
	//
	// *   **loadbalancer**: an ALB instance
	// *   **acl**: a network access control list (ACL)
	// *   **securitypolicy**: a security policy
	// *   **servergroup**: a server group
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries returned per page.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The tag key. It can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no next query is to be sent.
	// *   If a value of **NextToken** is returned, the value is the token that is used for the next query.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceId(v string) *ListTagValuesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

type ListTagValuesResponseBody struct {
	// The ID of the request.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The tags that match all the filter conditions.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries returned.
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagValues  []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetMaxResults(v int32) *ListTagValuesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetTagValues(v []*string) *ListTagValuesResponseBody {
	s.TagValues = v
	return s
}

func (s *ListTagValuesResponseBody) SetTotalCount(v int32) *ListTagValuesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	// The ID of the resource group to which the cloud resource is to be moved.
	//
	// >  You can use resource groups to classify cloud resources that belong to your Apsara Stack tenant account to facilitate resource management and permission control. For more information, see [What is resource management?](~~94475~~).
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the cloud resource that you want to move.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// *   **loadbalancer**: an Application Load Balancer (ALB) instance.
	// *   **acl**: an access control list (ACL).
	// *   **securitypolicy**: a security policy.
	// *   **servergroup**: a server group.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type RemoveEntriesFromAclRequest struct {
	// The ID of the ACL.
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId** as **ClientToken**. The value of **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// A list of removed access control entries.
	Entries []*string `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
}

func (s RemoveEntriesFromAclRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclRequest) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclRequest) SetAclId(v string) *RemoveEntriesFromAclRequest {
	s.AclId = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetClientToken(v string) *RemoveEntriesFromAclRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetDryRun(v bool) *RemoveEntriesFromAclRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetEntries(v []*string) *RemoveEntriesFromAclRequest {
	s.Entries = v
	return s
}

type RemoveEntriesFromAclResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveEntriesFromAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclResponseBody) SetJobId(v string) *RemoveEntriesFromAclResponseBody {
	s.JobId = &v
	return s
}

func (s *RemoveEntriesFromAclResponseBody) SetRequestId(v string) *RemoveEntriesFromAclResponseBody {
	s.RequestId = &v
	return s
}

type RemoveEntriesFromAclResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveEntriesFromAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveEntriesFromAclResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclResponse) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclResponse) SetHeaders(v map[string]*string) *RemoveEntriesFromAclResponse {
	s.Headers = v
	return s
}

func (s *RemoveEntriesFromAclResponse) SetStatusCode(v int32) *RemoveEntriesFromAclResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveEntriesFromAclResponse) SetBody(v *RemoveEntriesFromAclResponseBody) *RemoveEntriesFromAclResponse {
	s.Body = v
	return s
}

type RemoveServersFromServerGroupRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** may be different for each API request.
	ServerGroupId *string                                       `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	Servers       []*RemoveServersFromServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s RemoveServersFromServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupRequest) SetClientToken(v string) *RemoveServersFromServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetDryRun(v bool) *RemoveServersFromServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetServerGroupId(v string) *RemoveServersFromServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetServers(v []*RemoveServersFromServerGroupRequestServers) *RemoveServersFromServerGroupRequest {
	s.Servers = v
	return s
}

type RemoveServersFromServerGroupRequestServers struct {
	// Specifies whether only to precheck this request. Valid values:
	//
	// *   **trues**: prechecks the request but does not remove the server from the server group. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the asynchronous task.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The ID of the request.
	ServerIp   *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
}

func (s RemoveServersFromServerGroupRequestServers) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupRequestServers) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupRequestServers) SetPort(v int32) *RemoveServersFromServerGroupRequestServers {
	s.Port = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerId(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerIp(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerIp = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerType(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerType = &v
	return s
}

type RemoveServersFromServerGroupResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveServersFromServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupResponseBody) SetJobId(v string) *RemoveServersFromServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetRequestId(v string) *RemoveServersFromServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveServersFromServerGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveServersFromServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveServersFromServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupResponse) SetHeaders(v map[string]*string) *RemoveServersFromServerGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveServersFromServerGroupResponse) SetStatusCode(v int32) *RemoveServersFromServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveServersFromServerGroupResponse) SetBody(v *RemoveServersFromServerGroupResponseBody) *RemoveServersFromServerGroupResponse {
	s.Body = v
	return s
}

type ReplaceServersInServerGroupRequest struct {
	// The backend servers that you want to add to the server group. You can specify up to 40 backend servers in each call.
	AddedServers []*ReplaceServersInServerGroupRequestAddedServers `json:"AddedServers,omitempty" xml:"AddedServers,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx` HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The backend servers that you want to remove.
	RemovedServers []*ReplaceServersInServerGroupRequestRemovedServers `json:"RemovedServers,omitempty" xml:"RemovedServers,omitempty" type:"Repeated"`
	// The ID of the server group.
	//
	// > You cannot perform this operation on a server group of the Function type.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s ReplaceServersInServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplaceServersInServerGroupRequest) GoString() string {
	return s.String()
}

func (s *ReplaceServersInServerGroupRequest) SetAddedServers(v []*ReplaceServersInServerGroupRequestAddedServers) *ReplaceServersInServerGroupRequest {
	s.AddedServers = v
	return s
}

func (s *ReplaceServersInServerGroupRequest) SetClientToken(v string) *ReplaceServersInServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ReplaceServersInServerGroupRequest) SetDryRun(v bool) *ReplaceServersInServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *ReplaceServersInServerGroupRequest) SetRemovedServers(v []*ReplaceServersInServerGroupRequestRemovedServers) *ReplaceServersInServerGroupRequest {
	s.RemovedServers = v
	return s
}

func (s *ReplaceServersInServerGroupRequest) SetServerGroupId(v string) *ReplaceServersInServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

type ReplaceServersInServerGroupRequestAddedServers struct {
	// The description of the backend server. The description must be 2 to 256 characters in length, and can contain letters, digits, periods (.), underscores (\_), hyphens (-), commas (,), semicolons (;), forward slashes (/), and at signs (@). You can specify at most 40 servers in each call.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port used by the server group. Valid values: **1** to **65535**. You can specify at most 40 servers in each call.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server. You can specify up to 40 server IDs in each call.
	//
	// *   If the server group type is **Instance**, set the ServerId parameter to the ID of an ECS instance, an ENI, or an elastic container instance. These backend servers are specified by **Ecs**, **Eni**, or **Eci**.
	// *   If the server group type is **Ip**, set the ServerId parameter to an IP address specified in the server group.
	//
	// > You cannot perform this operation on a server group of the Function type. You can call the [ListServerGroups](~~213627~~) operation to query information about the server group type so that you can set ServerId to a proper value.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address in inclusive ENI mode. You can specify at most 40 servers in each call.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server that you want to remove from the server group. You can specify up to 40 backend servers in each call. Valid values:
	//
	// *   **Ecs**
	// *   **Eni**
	// *   **Eci**
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The weight of the backend server that you want to add to the server group. You can specify up to 40 backend servers in each call.
	//
	// Valid values: **0** to **100**. Default value: **100**. If the weight of a backend server is set to **0**, no requests are forwarded to the backend server.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ReplaceServersInServerGroupRequestAddedServers) String() string {
	return tea.Prettify(s)
}

func (s ReplaceServersInServerGroupRequestAddedServers) GoString() string {
	return s.String()
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetDescription(v string) *ReplaceServersInServerGroupRequestAddedServers {
	s.Description = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetPort(v int32) *ReplaceServersInServerGroupRequestAddedServers {
	s.Port = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetServerId(v string) *ReplaceServersInServerGroupRequestAddedServers {
	s.ServerId = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetServerIp(v string) *ReplaceServersInServerGroupRequestAddedServers {
	s.ServerIp = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetServerType(v string) *ReplaceServersInServerGroupRequestAddedServers {
	s.ServerType = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetWeight(v int32) *ReplaceServersInServerGroupRequestAddedServers {
	s.Weight = &v
	return s
}

type ReplaceServersInServerGroupRequestRemovedServers struct {
	// The port that is used by the backend server. Valid values: **1** to **65535**. You can specify at most 40 servers in each call.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server. You can specify up to 40 server IDs in each call.
	//
	// *   If the server group type is **Instance**, set the ServerId parameter to the ID of an ECS instance, an ENI, or an elastic container instance. These backend servers are specified by **Ecs**, **Eni**, or **Eci**.
	// *   If the server group type is **Ip**, set the ServerId parameter to an IP address specified in the server group.
	//
	// > You cannot perform this operation on a server group of the Function type. You can call the [ListServerGroups](~~213627~~) operation to query information about the server group type so that you can set ServerId to a proper value.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address in inclusive ENI mode. You can specify at most 40 servers in each call.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server that you want to remove from the server group. You can specify up to 40 backend servers in each call. Valid values:
	//
	// *   **Ecs**
	// *   **Eni**
	// *   **Eci**
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
}

func (s ReplaceServersInServerGroupRequestRemovedServers) String() string {
	return tea.Prettify(s)
}

func (s ReplaceServersInServerGroupRequestRemovedServers) GoString() string {
	return s.String()
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) SetPort(v int32) *ReplaceServersInServerGroupRequestRemovedServers {
	s.Port = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) SetServerId(v string) *ReplaceServersInServerGroupRequestRemovedServers {
	s.ServerId = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) SetServerIp(v string) *ReplaceServersInServerGroupRequestRemovedServers {
	s.ServerIp = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) SetServerType(v string) *ReplaceServersInServerGroupRequestRemovedServers {
	s.ServerType = &v
	return s
}

type ReplaceServersInServerGroupResponseBody struct {
	// The asynchronous task ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplaceServersInServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplaceServersInServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceServersInServerGroupResponseBody) SetJobId(v string) *ReplaceServersInServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *ReplaceServersInServerGroupResponseBody) SetRequestId(v string) *ReplaceServersInServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type ReplaceServersInServerGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReplaceServersInServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplaceServersInServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplaceServersInServerGroupResponse) GoString() string {
	return s.String()
}

func (s *ReplaceServersInServerGroupResponse) SetHeaders(v map[string]*string) *ReplaceServersInServerGroupResponse {
	s.Headers = v
	return s
}

func (s *ReplaceServersInServerGroupResponse) SetStatusCode(v int32) *ReplaceServersInServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceServersInServerGroupResponse) SetBody(v *ReplaceServersInServerGroupResponseBody) *ReplaceServersInServerGroupResponse {
	s.Body = v
	return s
}

type StartListenerRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ListenerId  *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s StartListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s StartListenerRequest) GoString() string {
	return s.String()
}

func (s *StartListenerRequest) SetClientToken(v string) *StartListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *StartListenerRequest) SetDryRun(v bool) *StartListenerRequest {
	s.DryRun = &v
	return s
}

func (s *StartListenerRequest) SetListenerId(v string) *StartListenerRequest {
	s.ListenerId = &v
	return s
}

type StartListenerResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartListenerResponseBody) GoString() string {
	return s.String()
}

func (s *StartListenerResponseBody) SetJobId(v string) *StartListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *StartListenerResponseBody) SetRequestId(v string) *StartListenerResponseBody {
	s.RequestId = &v
	return s
}

type StartListenerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s StartListenerResponse) GoString() string {
	return s.String()
}

func (s *StartListenerResponse) SetHeaders(v map[string]*string) *StartListenerResponse {
	s.Headers = v
	return s
}

func (s *StartListenerResponse) SetStatusCode(v int32) *StartListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *StartListenerResponse) SetBody(v *StartListenerResponseBody) *StartListenerResponse {
	s.Body = v
	return s
}

type StopListenerRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ListenerId  *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s StopListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s StopListenerRequest) GoString() string {
	return s.String()
}

func (s *StopListenerRequest) SetClientToken(v string) *StopListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *StopListenerRequest) SetDryRun(v bool) *StopListenerRequest {
	s.DryRun = &v
	return s
}

func (s *StopListenerRequest) SetListenerId(v string) *StopListenerRequest {
	s.ListenerId = &v
	return s
}

type StopListenerResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopListenerResponseBody) GoString() string {
	return s.String()
}

func (s *StopListenerResponseBody) SetJobId(v string) *StopListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *StopListenerResponseBody) SetRequestId(v string) *StopListenerResponseBody {
	s.RequestId = &v
	return s
}

type StopListenerResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s StopListenerResponse) GoString() string {
	return s.String()
}

func (s *StopListenerResponse) SetHeaders(v map[string]*string) *StopListenerResponse {
	s.Headers = v
	return s
}

func (s *StopListenerResponse) SetStatusCode(v int32) *StopListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *StopListenerResponse) SetBody(v *StopListenerResponseBody) *StopListenerResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The tag key. It can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The tags.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag value. It can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The ID of the request.
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnTagResourcesRequest struct {
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The key of the tag that you want to remove. It can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The tags that you want to remove.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The value of the tag that you want to remove. It can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	Tag []*UnTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the request.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTag(v []*UnTagResourcesRequestTag) *UnTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

type UnTagResourcesRequestTag struct {
	// The tag keys. It can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Specifies whether to remove all tags from the specified resource. Valid values:
	//
	// *   **true**: removes all tags from the resource.
	// *   **false**: does not remove all tags from the specified resource. This is the default value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UnTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequestTag) SetKey(v string) *UnTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *UnTagResourcesRequestTag) SetValue(v string) *UnTagResourcesRequestTag {
	s.Value = &v
	return s
}

type UnTagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponseBody) SetRequestId(v string) *UnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UnTagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetStatusCode(v int32) *UnTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnTagResourcesResponse) SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse {
	s.Body = v
	return s
}

type UpdateAScriptsRequest struct {
	// The AScript rules.
	AScripts []*UpdateAScriptsRequestAScripts `json:"AScripts,omitempty" xml:"AScripts,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false**(default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s UpdateAScriptsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAScriptsRequest) GoString() string {
	return s.String()
}

func (s *UpdateAScriptsRequest) SetAScripts(v []*UpdateAScriptsRequestAScripts) *UpdateAScriptsRequest {
	s.AScripts = v
	return s
}

func (s *UpdateAScriptsRequest) SetClientToken(v string) *UpdateAScriptsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAScriptsRequest) SetDryRun(v bool) *UpdateAScriptsRequest {
	s.DryRun = &v
	return s
}

type UpdateAScriptsRequestAScripts struct {
	// The AScript rule ID.
	AScriptId *string `json:"AScriptId,omitempty" xml:"AScriptId,omitempty"`
	// The name of the AScript rule.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	AScriptName *string `json:"AScriptName,omitempty" xml:"AScriptName,omitempty"`
	// Specifies whether to enable the AScript rule. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The content of the AScript rule.
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
}

func (s UpdateAScriptsRequestAScripts) String() string {
	return tea.Prettify(s)
}

func (s UpdateAScriptsRequestAScripts) GoString() string {
	return s.String()
}

func (s *UpdateAScriptsRequestAScripts) SetAScriptId(v string) *UpdateAScriptsRequestAScripts {
	s.AScriptId = &v
	return s
}

func (s *UpdateAScriptsRequestAScripts) SetAScriptName(v string) *UpdateAScriptsRequestAScripts {
	s.AScriptName = &v
	return s
}

func (s *UpdateAScriptsRequestAScripts) SetEnabled(v bool) *UpdateAScriptsRequestAScripts {
	s.Enabled = &v
	return s
}

func (s *UpdateAScriptsRequestAScripts) SetScriptContent(v string) *UpdateAScriptsRequestAScripts {
	s.ScriptContent = &v
	return s
}

type UpdateAScriptsResponseBody struct {
	// The asynchronous task ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAScriptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAScriptsResponseBody) SetJobId(v string) *UpdateAScriptsResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateAScriptsResponseBody) SetRequestId(v string) *UpdateAScriptsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAScriptsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAScriptsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAScriptsResponse) GoString() string {
	return s.String()
}

func (s *UpdateAScriptsResponse) SetHeaders(v map[string]*string) *UpdateAScriptsResponse {
	s.Headers = v
	return s
}

func (s *UpdateAScriptsResponse) SetStatusCode(v int32) *UpdateAScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAScriptsResponse) SetBody(v *UpdateAScriptsResponseBody) *UpdateAScriptsResponse {
	s.Body = v
	return s
}

type UpdateAclAttributeRequest struct {
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: only prechecks the request and does not perform the requested operation. The system checks the required parameters, request format, and service limits. If the request fails the precheck, an error code is returned based on the cause of the failure. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false**: prechecks the request and performs the requested operation. After the request passes the precheck, an `HTTP 2xx` status code is returned and the system performs the operation. This is the default value.
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the request.
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The name of the ACL. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s UpdateAclAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeRequest) SetAclId(v string) *UpdateAclAttributeRequest {
	s.AclId = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetAclName(v string) *UpdateAclAttributeRequest {
	s.AclName = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetClientToken(v string) *UpdateAclAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetDryRun(v bool) *UpdateAclAttributeRequest {
	s.DryRun = &v
	return s
}

type UpdateAclAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAclAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeResponseBody) SetRequestId(v string) *UpdateAclAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAclAttributeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAclAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAclAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeResponse) SetHeaders(v map[string]*string) *UpdateAclAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAclAttributeResponse) SetStatusCode(v int32) *UpdateAclAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAclAttributeResponse) SetBody(v *UpdateAclAttributeResponseBody) *UpdateAclAttributeResponse {
	s.Body = v
	return s
}

type UpdateHealthCheckTemplateAttributeRequest struct {
	// The port that is used for health checks. Valid values: **0 to 65535**.
	//
	// Default value: **0**. If you set the value to 0, the port of the backend server is used for health checks.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request.
	DryRun           *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The HTTP status codes that are used to determine whether the backend server passes the health check.
	//
	// *   If **HealthCheckProtocol** is set to **HTTP**, **HealthCheckCodes** can be set to **http\_2xx** (default), **http\_3xx**, **http\_4xx**, and **http\_5xx**. Separate multiple HTTP status codes with a comma (,).
	// *   If **HealthCheckProtocol** is set to **gRPC**, **HealthCheckCodes** can be set to **0 to 99**. Default value: **0**. Value ranges are supported. You can enter 20 value ranges at most and must separate each range with a comma (,).
	//
	// >  This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds. Valid values: **1 to 50**. Default value: **2**.
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success** to **fail**.
	//
	// Valid values: **2 to 10**.
	//
	// Default value: **3**.
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The URL path that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length and can contain only letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&). The URL can also contain the following extended characters: `_ ; ~ ! ( ) * [ ] @ $ ^ : \" , +`.
	//
	// The URL path must start with a forward slash (/).
	//
	// >  This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP version that is used for health checks.
	//
	// Valid values: **HTTP1.0** and **HTTP1.1**.
	//
	// Default value: **HTTP1.1**.
	//
	// >  This parameter takes effect only when the `HealthCheckProtocol` parameter is set to **HTTP**.
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// *   **HTTP** (default): To perform HTTP health checks, ALB sends HEAD or GET requests to a backend server to check whether the backend server is healthy.
	// *   **TCP**: To perform TCP health checks, ALB sends SYN packets to a backend server to check whether the port of the backend server is available to receive requests.
	// *   **gRPC**: To perform gRPC health checks, ALB sends POST or GET requests to a backend server to check whether the backend server is healthy.
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail** to **success**.
	//
	// Valid values: **2 to 10**.
	//
	// Default value: **3**.
	HealthCheckProtocol   *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// *   **$SERVER_IP** (default): the private IP addresses of backend servers. If you do not set the HealthCheckHost parameter or set the parameter to $SERVER_IP, the Application Load Balancer (ALB) uses the private IP addresses of backend servers for health checks.
	// *   **domain**: The domain name must be 1 to 80 characters in length and can contain letters, digits, periods (.), and hyphens (-).
	//
	// >  This parameter takes effect only when the `HealthCheckProtocol` parameter is set to **HTTP**.
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// The HTTP method that is used for health checks. Valid values:
	//
	// *   **HEAD**: By default, HTTP health checks use the HEAD method.
	// *   **GET**: If the size of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	// *   **POST**: By default, gRPC health checks use the POST method.
	//
	// >  This parameter takes effect only when the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// Specifies whether to only precheck this request. Valid values:
	//
	// *   **true**: prechecks the request without modifying the attributes of the health check template. The system checks the required parameters, request syntax, and limits. If the request fails to pass the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, a 2xx HTTP status code is returned and the health check template is modified.
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The ID of the template.
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s UpdateHealthCheckTemplateAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHealthCheckTemplateAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetClientToken(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetDryRun(v bool) *UpdateHealthCheckTemplateAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckCodes(v []*string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckCodes = v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckConnectPort(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckHost(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckHost = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckHttpVersion(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckInterval(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckMethod(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckPath(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckPath = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckProtocol(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckProtocol = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckTemplateId(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckTemplateName(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckTimeout(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthyThreshold(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetUnhealthyThreshold(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

type UpdateHealthCheckTemplateAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHealthCheckTemplateAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHealthCheckTemplateAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckTemplateAttributeResponseBody) SetRequestId(v string) *UpdateHealthCheckTemplateAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateHealthCheckTemplateAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateHealthCheckTemplateAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateHealthCheckTemplateAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHealthCheckTemplateAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckTemplateAttributeResponse) SetHeaders(v map[string]*string) *UpdateHealthCheckTemplateAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeResponse) SetStatusCode(v int32) *UpdateHealthCheckTemplateAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeResponse) SetBody(v *UpdateHealthCheckTemplateAttributeResponseBody) *UpdateHealthCheckTemplateAttributeResponse {
	s.Body = v
	return s
}

type UpdateListenerAttributeRequest struct {
	// The certificate authority (CA) certificates.
	CaCertificates []*UpdateListenerAttributeRequestCaCertificates `json:"CaCertificates,omitempty" xml:"CaCertificates,omitempty" type:"Repeated"`
	// Specifies whether to enable mutual authentication. Valid values:
	//
	// *   **true**
	// *   **false**
	CaEnabled *bool `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// The certificates.
	Certificates []*UpdateListenerAttributeRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The actions of the default forwarding rule.
	DefaultActions []*UpdateListenerAttributeRequestDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable GZIP compression for specific types of files. Valid values:
	//
	// *   **true**
	// *   **false**
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// Specifies whether to enable HTTP/2. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The timeout period of an idle connection. Unit: seconds. Valid values: **1 to 60**.
	//
	// If no request is received within the specified timeout period, ALB closes the current connection. When another request is received, ALB establishes a new connection.
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The name of the listener.
	//
	// The name must be 2 to 256 characters in length, and can contain letters, digits, and the following special characters: , . ; / @ \_ -.
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The ID of the Application Load Balancer (ALB) listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The configuration information when the listener is associated with a QUIC listener.
	QuicConfig *UpdateListenerAttributeRequestQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// The timeout period of a request. Unit: seconds. Valid values: **1 to 180**.
	//
	// If no response is received from the backend server within the specified timeout period, ALB returns an `HTTP 504` error code to the client.
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The security policy ID. System security policies and custom security policies are supported.
	//
	// > This parameter is available only when you create an HTTPS listener.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The configuration of the XForwardFor headers.
	XForwardedForConfig *UpdateListenerAttributeRequestXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s UpdateListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequest) SetCaCertificates(v []*UpdateListenerAttributeRequestCaCertificates) *UpdateListenerAttributeRequest {
	s.CaCertificates = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCaEnabled(v bool) *UpdateListenerAttributeRequest {
	s.CaEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCertificates(v []*UpdateListenerAttributeRequestCertificates) *UpdateListenerAttributeRequest {
	s.Certificates = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetClientToken(v string) *UpdateListenerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetDefaultActions(v []*UpdateListenerAttributeRequestDefaultActions) *UpdateListenerAttributeRequest {
	s.DefaultActions = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetDryRun(v bool) *UpdateListenerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetGzipEnabled(v bool) *UpdateListenerAttributeRequest {
	s.GzipEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetHttp2Enabled(v bool) *UpdateListenerAttributeRequest {
	s.Http2Enabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetIdleTimeout(v int32) *UpdateListenerAttributeRequest {
	s.IdleTimeout = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerDescription(v string) *UpdateListenerAttributeRequest {
	s.ListenerDescription = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerId(v string) *UpdateListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetQuicConfig(v *UpdateListenerAttributeRequestQuicConfig) *UpdateListenerAttributeRequest {
	s.QuicConfig = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetRequestTimeout(v int32) *UpdateListenerAttributeRequest {
	s.RequestTimeout = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetSecurityPolicyId(v string) *UpdateListenerAttributeRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetXForwardedForConfig(v *UpdateListenerAttributeRequestXForwardedForConfig) *UpdateListenerAttributeRequest {
	s.XForwardedForConfig = v
	return s
}

type UpdateListenerAttributeRequestCaCertificates struct {
}

func (s UpdateListenerAttributeRequestCaCertificates) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestCaCertificates) GoString() string {
	return s.String()
}

type UpdateListenerAttributeRequestCertificates struct {
	// The certificate ID. Only server certificates are supported. You can specify up to 20 certificate IDs.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s UpdateListenerAttributeRequestCertificates) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestCertificates) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestCertificates) SetCertificateId(v string) *UpdateListenerAttributeRequestCertificates {
	s.CertificateId = &v
	return s
}

type UpdateListenerAttributeRequestDefaultActions struct {
	// The configuration of the forwarding action. This parameter is required and takes effect when **Type** is set to **FowardGroup**. You can specify configurations for up to 20 forwarding actions.
	ForwardGroupConfig *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The type of the action. You can specify only one action type.
	//
	// Set the value to **ForwardGroup** to forward requests to multiple vServer groups.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateListenerAttributeRequestDefaultActions) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestDefaultActions) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestDefaultActions) SetForwardGroupConfig(v *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) *UpdateListenerAttributeRequestDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *UpdateListenerAttributeRequestDefaultActions) SetType(v string) *UpdateListenerAttributeRequestDefaultActions {
	s.Type = &v
	return s
}

type UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig struct {
	// The server groups to which requests are forwarded.
	ServerGroupTuples []*UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// The server group to which requests are forwarded.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type UpdateListenerAttributeRequestQuicConfig struct {
	// The QUIC listener ID. This parameter is required if **QuicUpgradeEnabled** is set to **true**. Only HTTPS listeners support this parameter.
	//
	// > You must add the HTTPS listener and the QUIC listener to the same ALB instance. In addition, make sure that the QUIC listener has never been associated with another listener.
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// Specifies whether to enable QUIC upgrade. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > Only HTTPS listeners support this parameter.
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s UpdateListenerAttributeRequestQuicConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestQuicConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestQuicConfig) SetQuicListenerId(v string) *UpdateListenerAttributeRequestQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *UpdateListenerAttributeRequestQuicConfig) SetQuicUpgradeEnabled(v bool) *UpdateListenerAttributeRequestQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

type UpdateListenerAttributeRequestXForwardedForConfig struct {
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertClientVerifyEnabled** is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (\_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertClientVerifyAlias *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-clientverify` header to retrieve the verification result of the client certificate. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertClientVerifyEnabled *bool `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertFingerprintEnabled** is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (\_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertFingerprintAlias *string `json:"XForwardedForClientCertFingerprintAlias,omitempty" xml:"XForwardedForClientCertFingerprintAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-fingerprint` header to retrieve the fingerprint of the client certificate. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertFingerprintEnabled *bool `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertIssuerDNEnabled** is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (\_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertIssuerDNAlias *string `json:"XForwardedForClientCertIssuerDNAlias,omitempty" xml:"XForwardedForClientCertIssuerDNAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-issuerdn` header to retrieve information about the authority that issues the client certificate. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertIssuerDNEnabled *bool `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertSubjectDNEnabled** is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (\_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertSubjectDNAlias *string `json:"XForwardedForClientCertSubjectDNAlias,omitempty" xml:"XForwardedForClientCertSubjectDNAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-subjectdn` header to retrieve information about the owner of the client certificate. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	XForwardedForClientCertSubjectDNEnabled *bool `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-Client-Ip` header to retrieve the source IP address. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP, HTTPS, and QUIC listeners support this parameter. By default, the feature that corresponds to this parameter is unavailable. If you want to use this feature, contact your account manager.
	XForwardedForClientSourceIpsEnabled *bool `json:"XForwardedForClientSourceIpsEnabled,omitempty" xml:"XForwardedForClientSourceIpsEnabled,omitempty"`
	// The trusted proxy IP address.
	//
	// ALB traverses `X-Forwarded-For` backward and selects the first IP address that is not in the trusted IP address list as the real IP address of the client. The IP address is used in source IP address throttling.
	XForwardedForClientSourceIpsTrusted *string `json:"XForwardedForClientSourceIpsTrusted,omitempty" xml:"XForwardedForClientSourceIpsTrusted,omitempty"`
	// Specifies whether to use the `X-Forwarded-Client-Port` header to retrieve the client port. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP and HTTPS listeners support this parameter.
	XForwardedForClientSrcPortEnabled *bool `json:"XForwardedForClientSrcPortEnabled,omitempty" xml:"XForwardedForClientSrcPortEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-For` header to retrieve the client IP address. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP and HTTPS listeners support this parameter.
	XForwardedForEnabled *bool `json:"XForwardedForEnabled,omitempty" xml:"XForwardedForEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-Proto` header to retrieve the listener protocol. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP, HTTPS, and QUIC listeners support this parameter.
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// Specifies whether to use the `SLB-ID` header to retrieve the ID of the ALB instance. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP, HTTPS, and QUIC listeners support this parameter.
	XForwardedForSLBIdEnabled *bool `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-Port` header to retrieve the listener port. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > HTTP, HTTPS, and QUIC listeners support this parameter.
	XForwardedForSLBPortEnabled *bool `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s UpdateListenerAttributeRequestXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientSourceIpsEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientSourceIpsEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientSourceIpsTrusted(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientSourceIpsTrusted = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

type UpdateListenerAttributeResponseBody struct {
	// The asynchronous task ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeResponseBody) SetJobId(v string) *UpdateListenerAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateListenerAttributeResponseBody) SetRequestId(v string) *UpdateListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateListenerAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeResponse) SetHeaders(v map[string]*string) *UpdateListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateListenerAttributeResponse) SetStatusCode(v int32) *UpdateListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateListenerAttributeResponse) SetBody(v *UpdateListenerAttributeResponseBody) *UpdateListenerAttributeResponse {
	s.Body = v
	return s
}

type UpdateListenerLogConfigRequest struct {
	// Specifies whether to record custom headers in the access log. Valid values:
	//
	// *   **true**: yes
	// *   **false** (default): no
	//
	// >  This parameter can be set to **true** only when the access log feature of ALB is enabled by specifying **AccessLogEnabled**.
	AccessLogRecordCustomizedHeadersEnabled *bool `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// The configuration information about the Xtrace feature.
	AccessLogTracingConfig *UpdateListenerLogConfigRequestAccessLogTracingConfig `json:"AccessLogTracingConfig,omitempty" xml:"AccessLogTracingConfig,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId** as **ClientToken**. The value of **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false**: sends the request. If the request passes the check, the **HTTP\_2xx** status code is returned and the operation is performed. This is the default value.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener of the ALB instance.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s UpdateListenerLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerLogConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigRequest) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *UpdateListenerLogConfigRequest {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetAccessLogTracingConfig(v *UpdateListenerLogConfigRequestAccessLogTracingConfig) *UpdateListenerLogConfigRequest {
	s.AccessLogTracingConfig = v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetClientToken(v string) *UpdateListenerLogConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetDryRun(v bool) *UpdateListenerLogConfigRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetListenerId(v string) *UpdateListenerLogConfigRequest {
	s.ListenerId = &v
	return s
}

type UpdateListenerLogConfigRequestAccessLogTracingConfig struct {
	// Specifies whether to enable the Xtrace feature. Valid values:
	//
	// *   **true**: yes
	// *   **false** (default): no
	//
	// >  This parameter can be set to **true** only when the access log feature of ALB is enabled by specifying **AccessLogEnabled**.
	TracingEnabled *bool `json:"TracingEnabled,omitempty" xml:"TracingEnabled,omitempty"`
	// The sampling rate of the Xtrace feature.
	//
	// Valid values: **1 to 10000**.
	//
	// >  This parameter is valid only if the **TracingEnabled** parameter is set to **true**.
	TracingSample *int32 `json:"TracingSample,omitempty" xml:"TracingSample,omitempty"`
	// The type of Xtrace. Set the value to **Zipkin**.
	//
	// >  This parameter is valid only if the **TracingEnabled** parameter is set to **true**.
	TracingType *string `json:"TracingType,omitempty" xml:"TracingType,omitempty"`
}

func (s UpdateListenerLogConfigRequestAccessLogTracingConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerLogConfigRequestAccessLogTracingConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) SetTracingEnabled(v bool) *UpdateListenerLogConfigRequestAccessLogTracingConfig {
	s.TracingEnabled = &v
	return s
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) SetTracingSample(v int32) *UpdateListenerLogConfigRequestAccessLogTracingConfig {
	s.TracingSample = &v
	return s
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) SetTracingType(v string) *UpdateListenerLogConfigRequestAccessLogTracingConfig {
	s.TracingType = &v
	return s
}

type UpdateListenerLogConfigResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListenerLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigResponseBody) SetJobId(v string) *UpdateListenerLogConfigResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateListenerLogConfigResponseBody) SetRequestId(v string) *UpdateListenerLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateListenerLogConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateListenerLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateListenerLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerLogConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigResponse) SetHeaders(v map[string]*string) *UpdateListenerLogConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateListenerLogConfigResponse) SetStatusCode(v int32) *UpdateListenerLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateListenerLogConfigResponse) SetBody(v *UpdateListenerLogConfigResponseBody) *UpdateListenerLogConfigResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerAddressTypeConfigRequest struct {
	// The ID of the zone where the ALB instance is deployed. You can specify up to 10 zones for an ALB instance.
	//
	// You can call the [DescribeZones](~~189196~~) operation to query the most recent zone list.
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the ALB instance.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new network type of the ALB instance. Valid values:
	//
	// *   **Internet**: The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. Therefore, the ALB instance can be accessed over the Internet.
	// *   **Intranet**: The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. Therefore, the ALB instance can be accessed over the virtual private cloud (VPC) where the ALB instance is deployed.
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of an ALB instance. You can specify up to 10 zones for an ALB instance.
	LoadBalancerId *string                                                   `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	ZoneMappings   []*UpdateLoadBalancerAddressTypeConfigRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s UpdateLoadBalancerAddressTypeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetAddressType(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.AddressType = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetClientToken(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetDryRun(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetZoneMappings(v []*UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.ZoneMappings = v
	return s
}

type UpdateLoadBalancerAddressTypeConfigRequestZoneMappings struct {
	// The ID of the asynchronous task.
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The ID of the elastic IP address (EIP). You can specify up to 10 zones for an ALB instance.
	//
	// >  This parameter is required if you want to change the network type from internal-facing to Internet-facing.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the request.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetAllocationId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.AllocationId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetVSwitchId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetZoneId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type UpdateLoadBalancerAddressTypeConfigResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerAddressTypeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetJobId(v string) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetRequestId(v string) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerAddressTypeConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLoadBalancerAddressTypeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerAddressTypeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerAddressTypeConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponse) SetStatusCode(v int32) *UpdateLoadBalancerAddressTypeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponse) SetBody(v *UpdateLoadBalancerAddressTypeConfigResponseBody) *UpdateLoadBalancerAddressTypeConfigResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerAttributeRequest struct {
	// The status of the configuration read-only mode. Valid values:
	//
	// *   **NonProtection**: disables the configuration read-only mode. In this case, you cannot set the **ModificationProtectionReason** parameter. If you set the **ModificationProtectionReason** parameter, the value of the parameter is cleared.
	// *   **ConsoleProtection**: enables the configuration read-only mode. In this case, you can set the **ModificationProtectionReason** parameter.
	//
	// >  If you set this parameter to **ConsoleProtection**, you cannot modify instance configurations in the ALB console. However, you can modify instance configurations by calling API operations.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the synchronous task.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the request.
	LoadBalancerId               *string                                                         `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName             *string                                                         `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	ModificationProtectionConfig *UpdateLoadBalancerAttributeRequestModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
}

func (s UpdateLoadBalancerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeRequest) SetClientToken(v string) *UpdateLoadBalancerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetDryRun(v bool) *UpdateLoadBalancerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerName(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetModificationProtectionConfig(v *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) *UpdateLoadBalancerAttributeRequest {
	s.ModificationProtectionConfig = v
	return s
}

type UpdateLoadBalancerAttributeRequestModificationProtectionConfig struct {
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateLoadBalancerAttributeRequestModificationProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeRequestModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) SetReason(v string) *UpdateLoadBalancerAttributeRequestModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) SetStatus(v string) *UpdateLoadBalancerAttributeRequestModificationProtectionConfig {
	s.Status = &v
	return s
}

type UpdateLoadBalancerAttributeResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetJobId(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetRequestId(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLoadBalancerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerAttributeResponse) SetStatusCode(v int32) *UpdateLoadBalancerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponse) SetBody(v *UpdateLoadBalancerAttributeResponseBody) *UpdateLoadBalancerAttributeResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerEditionRequest struct {
	// The ID of the ALB instance.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The edition of the ALB instance. Different editions have different limits and support different billing methods.
	//
	// *   **Basic**: basic
	// *   **Standard**: standard
	// *   **StandardWithWaf**: WAF-enabled
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// The ID of the request.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s UpdateLoadBalancerEditionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerEditionRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerEditionRequest) SetClientToken(v string) *UpdateLoadBalancerEditionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerEditionRequest) SetDryRun(v bool) *UpdateLoadBalancerEditionRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerEditionRequest) SetLoadBalancerEdition(v string) *UpdateLoadBalancerEditionRequest {
	s.LoadBalancerEdition = &v
	return s
}

func (s *UpdateLoadBalancerEditionRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerEditionRequest {
	s.LoadBalancerId = &v
	return s
}

type UpdateLoadBalancerEditionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerEditionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerEditionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerEditionResponseBody) SetRequestId(v string) *UpdateLoadBalancerEditionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerEditionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLoadBalancerEditionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerEditionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerEditionResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerEditionResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerEditionResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerEditionResponse) SetStatusCode(v int32) *UpdateLoadBalancerEditionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerEditionResponse) SetBody(v *UpdateLoadBalancerEditionResponseBody) *UpdateLoadBalancerEditionResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerZonesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and sends the request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the ALB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The zones and the vSwitches. You must specify at least two zones. The specified zones overwrite the existing configurations.
	ZoneMappings []*UpdateLoadBalancerZonesRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s UpdateLoadBalancerZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesRequest) SetClientToken(v string) *UpdateLoadBalancerZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetDryRun(v bool) *UpdateLoadBalancerZonesRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerZonesRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetZoneMappings(v []*UpdateLoadBalancerZonesRequestZoneMappings) *UpdateLoadBalancerZonesRequest {
	s.ZoneMappings = v
	return s
}

type UpdateLoadBalancerZonesRequestZoneMappings struct {
	// The ID of the vSwitch in the zone. By default, you can specify only one vSwitch (subnet) for each zone of an ALB instance. You can specify up to 10 zone IDs.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the zone. You can call the [DescribeZones](~~189196~~) operation to query the zones. You can specify up to 10 zone IDs.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateLoadBalancerZonesRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetVSwitchId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetZoneId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type UpdateLoadBalancerZonesResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesResponseBody) SetJobId(v string) *UpdateLoadBalancerZonesResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponseBody) SetRequestId(v string) *UpdateLoadBalancerZonesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerZonesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLoadBalancerZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerZonesResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerZonesResponse) SetStatusCode(v int32) *UpdateLoadBalancerZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponse) SetBody(v *UpdateLoadBalancerZonesResponseBody) *UpdateLoadBalancerZonesResponse {
	s.Body = v
	return s
}

type UpdateRuleAttributeRequest struct {
	// The key of the header. The key must be 1 to 40 characters in length and can contain letters, digits, hyphens (-), and underscores (\_). You cannot set Cookie or Host.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The value of the header. The header values within a forwarding rule must be unique. The header must meet the following requirements:
	//
	// *   The value must be 1 to 128 characters in length.
	// *   It can contain printable characters whose ASCII values are `greater than or equal to 32 and lower than 127`. You can use asterisks (\*) and question marks (?) as wildcard characters.
	// *   The header value cannot start or end with a space character.
	DryRun         *bool                                       `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Priority       *int32                                      `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RuleActions    []*UpdateRuleAttributeRequestRuleActions    `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	RuleConditions []*UpdateRuleAttributeRequestRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The value of the cookie. The value must be 1 to 128 characters in length, and can contain printable characters such as lowercase letters, asterisks (\*), and question marks (?). However, uppercase letters, space characters, and the following special characters are not supported: `# [ ] { } \ | < > &`.
	RuleId   *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s UpdateRuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequest) SetClientToken(v string) *UpdateRuleAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetDryRun(v bool) *UpdateRuleAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetPriority(v int32) *UpdateRuleAttributeRequest {
	s.Priority = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetRuleActions(v []*UpdateRuleAttributeRequestRuleActions) *UpdateRuleAttributeRequest {
	s.RuleActions = v
	return s
}

func (s *UpdateRuleAttributeRequest) SetRuleConditions(v []*UpdateRuleAttributeRequestRuleConditions) *UpdateRuleAttributeRequest {
	s.RuleConditions = v
	return s
}

func (s *UpdateRuleAttributeRequest) SetRuleId(v string) *UpdateRuleAttributeRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetRuleName(v string) *UpdateRuleAttributeRequest {
	s.RuleName = &v
	return s
}

type UpdateRuleAttributeRequestRuleActions struct {
	CorsConfig          *UpdateRuleAttributeRequestRuleActionsCorsConfig          `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	FixedResponseConfig *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	ForwardGroupConfig  *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig  `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	InsertHeaderConfig  *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig  `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// The port to which requests are distributed. Valid values:
	//
	// *   **${port}** (default): If you set the value to ${port}, you cannot append other characters.
	// *   Other valid values: **1 to 63335**.
	Order               *int32                                                    `json:"Order,omitempty" xml:"Order,omitempty"`
	RedirectConfig      *UpdateRuleAttributeRequestRuleActionsRedirectConfig      `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	RewriteConfig       *UpdateRuleAttributeRequestRuleActionsRewriteConfig       `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	TrafficLimitConfig  *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig  `json:"TrafficLimitConfig,omitempty" xml:"TrafficLimitConfig,omitempty" type:"Struct"`
	TrafficMirrorConfig *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig `json:"TrafficMirrorConfig,omitempty" xml:"TrafficMirrorConfig,omitempty" type:"Struct"`
	// The ID of the vServer group.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActions) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActions) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActions) SetCorsConfig(v *UpdateRuleAttributeRequestRuleActionsCorsConfig) *UpdateRuleAttributeRequestRuleActions {
	s.CorsConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetFixedResponseConfig(v *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) *UpdateRuleAttributeRequestRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetForwardGroupConfig(v *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) *UpdateRuleAttributeRequestRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetInsertHeaderConfig(v *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) *UpdateRuleAttributeRequestRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetOrder(v int32) *UpdateRuleAttributeRequestRuleActions {
	s.Order = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetRedirectConfig(v *UpdateRuleAttributeRequestRuleActionsRedirectConfig) *UpdateRuleAttributeRequestRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetRewriteConfig(v *UpdateRuleAttributeRequestRuleActionsRewriteConfig) *UpdateRuleAttributeRequestRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetTrafficLimitConfig(v *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) *UpdateRuleAttributeRequestRuleActions {
	s.TrafficLimitConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetTrafficMirrorConfig(v *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) *UpdateRuleAttributeRequestRuleActions {
	s.TrafficMirrorConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetType(v string) *UpdateRuleAttributeRequestRuleActions {
	s.Type = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsCorsConfig struct {
	// The ID of the asynchronous task.
	AllowCredentials *string   `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	AllowHeaders     []*string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty" type:"Repeated"`
	AllowMethods     []*string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty" type:"Repeated"`
	AllowOrigin      []*string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty" type:"Repeated"`
	ExposeHeaders    []*string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty" type:"Repeated"`
	// The ID of the request.
	MaxAge *int64 `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsCorsConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsCorsConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetAllowCredentials(v string) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetAllowHeaders(v []*string) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.AllowHeaders = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetAllowMethods(v []*string) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.AllowMethods = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetAllowOrigin(v []*string) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.AllowOrigin = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetExposeHeaders(v []*string) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.ExposeHeaders = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetMaxAge(v int64) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.MaxAge = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsFixedResponseConfig struct {
	// The weight of the server group. A larger value specifies a higher weight. A server group with a higher weight receives more requests. Valid values: **1** to **100**.
	//
	// *   If only one destination server group exists, the weight is **100** by default.
	// *   If more than one destination server group exists, you must specify weights.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Specifies whether to enable session persistence. Valid values:
	//
	// *   **true**: yes
	// *   **false** (default): no
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The timeout period. Unit: seconds. Valid values: 1 to 86400.
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) SetContent(v string) *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) SetContentType(v string) *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) SetHttpCode(v string) *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsForwardGroupConfig struct {
	ServerGroupStickySession *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession `json:"ServerGroupStickySession,omitempty" xml:"ServerGroupStickySession,omitempty" type:"Struct"`
	ServerGroupTuples        []*UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples      `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) SetServerGroupStickySession(v *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig {
	s.ServerGroupStickySession = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession struct {
	// The type of the header. Valid values:
	//
	// *   **UserDefined**: a custom header
	// *   **ReferenceHeader**: a header that references the request headers
	// *   **SystemDefined**: a system-defined header
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The priority of the action. Valid values: **1 to 50000**. A lower value specifies a higher priority. The actions of a forwarding rule are applied in descending order of priority. This parameter cannot be left empty. The priority of each action within a forwarding rule must be unique. You can specify priorities for at most 20 actions.
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) SetEnabled(v bool) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Enabled = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) SetTimeout(v int32) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Timeout = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples struct {
	// The key of the header to be inserted. The key must be 1 to 40 characters in length, and can contain lowercase letters, digits, underscores (\_), and hyphens (-). The header key specified by **InsertHeaderConfig** must be unique.
	//
	// >  You cannot set one of the following header keys (case-insensitive): `slb-id`, `slb-ip`, `x-forwarded-for`, `x-forwarded-proto`, `x-forwarded-eip`, `x-forwarded-port`, `x-forwarded-client-srcport`, `connection`, `upgrade`, `content-length`, `transfer-encoding`, `keep-alive`, `te`, `host`, `cookie`, `remoteip`, and `authority`.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The value of the header to be inserted.
	//
	// *   If **ValueType** is set to **SystemDefined**, you can specify one of the following header values:
	//
	//     *   **ClientSrcPort**: the client port.
	//     *   **ClientSrcIp**: the client IP address.
	//     *   **Protocol**: the request protocol. You can set the protocol to HTTP or HTTPS.
	//     *   **SLBId**: the ID of the Application Load Balancer (ALB) instance.
	//     *   **SLBPort**: the listening port.
	//
	// *   If **ValueType** is set to **UserDefined**, you can specify a custom header. The header value must be 1 to 128 characters in length, and can contain printable characters whose ASCII values are `greater than or equal to 32 and lower than 127`. You can use asterisks (\*) and question marks (?) as wildcards. The header value cannot start or end with a space character.
	//
	// *   If **ValueType** is set to **ReferenceHeader**, you can reference the request headers. The header value must be 1 to 128 characters in length, and can contain lowercase letters, digits, underscores (\_), and hyphens (-).
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) SetWeight(v int32) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig struct {
	// The hostname to which requests are distributed. Valid values:
	//
	// *   **${host}** (default): If you set the value to ${host}, you cannot append other characters.
	//
	// *   If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     *   The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), and periods (.). You can use asterisks (\*) and question marks (?) as wildcard characters.
	//     *   The hostname must contain at least one period (.) but cannot start or end with a period (.).
	//     *   The rightmost domain label can contain only letters and wildcard characters. It cannot contain digits or hyphens (-).
	//     *   The domain labels cannot start or end with a hyphen (-).
	//     *   You can use an asterisk (\*) and question mark (?) anywhere in a domain label as wildcards.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The redirect type. Valid values: **301**, **302**, **303**, **307**, and **308**.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The path to which requests are forwarded. Valid values:
	//
	// *   Default value: **${path}**. **${host}**, **${protocol}**, and **${port}** are also supported. Each variable cannot be specified more than once. You can specify one or more of the preceding variables in each request. You can also combine them with the following characters.
	//
	// *   A custom value. You must make sure that the custom value meets the following requirements:
	//
	//     *   The value is 1 to 128 characters in length.
	//     *   It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ .+ / & ~ @ :`. It cannot contain the following special characters: `" % # ; ! ( ) [ ]^ , "`. You can use asterisks (\*) and question marks (?) as wildcards.
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) SetKey(v string) *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) SetValue(v string) *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) SetValueType(v string) *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsRedirectConfig struct {
	// The redirect protocol. Valid values:
	//
	// *   **${protocol}** (default): If you set the value to ${protocol}, you cannot append other characters.
	// *   Valid values: **HTTP** and **HTTPS**.
	//
	// >  HTTPS listeners do not support HTTPS to HTTP redirects.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The query string of the URL to which requests are distributed. Valid values:
	//
	// *   Default value: **${query}**. **${host}**, **${protocol}**, and **${port}** are also supported. Each variable cannot be specified more than once. You can specify one or more of the preceding variables in each request. You can also combine them with the following characters.
	//
	// *   A custom value. You must make sure that the custom value meets the following requirements:
	//
	//     *   The value is 1 to 128 characters in length.
	//     *   It can contain printable characters, except space characters, the special characters `# [ ] { } \ | < > &`, and uppercase letters.
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The hostname to which requests are forwarded. Valid values:
	//
	// *   **${host}** (default): If you set the value to ${host}, you cannot append other characters.
	//
	// *   If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     *   The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), and periods (.). You can use asterisks (\*) and question marks (?) as wildcard characters.
	//     *   The hostname must contain at least one period (.) but cannot start or end with a period (.).
	//     *   The rightmost domain label can contain only letters and wildcard characters. It cannot contain digits or hyphens (-).
	//     *   The domain labels cannot start or end with a hyphen (-). You can use an asterisk (\*) and question mark (?) anywhere in a domain label as wildcards.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The path to which requests are forwarded. Valid values:
	//
	// *   Default value: **${path}**. **${host}**, **${protocol}**, and **${port}** are also supported. Each variable cannot be specified more than once. You can specify one or more of the preceding variables in each request. You can also combine them with the following characters.
	//
	// *   A custom value. You must make sure that the custom value meets the following requirements:
	//
	//     *   The value is 1 to 128 characters in length.
	//     *   It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ .+ / & ~ @ :`. It cannot contain the following special characters: `" % # ; ! ( ) [ ]^ , "`. You can use asterisks (\*) and question marks (?) as wildcards.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The query string of the URL to which requests are distributed. Valid values:
	//
	// *   Default value: **${query}**. **${host}**, **${protocol}**, and **${port}** are also supported. Each variable cannot be specified more than once. You can specify one or more of the preceding variables in each request. You can also combine them with the following characters.
	//
	// *   A custom value. You must make sure that the custom value meets the following requirements:
	//
	//     *   The value is 1 to 128 characters in length.
	//     *   It can contain printable characters, except space characters, the special characters `# [ ] { } \ | < > &`, and uppercase letters.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The action type. You can specify at most 11 types of action. Valid values:
	//
	// *   **ForwardGroup**: forwards a request to multiple vServer groups.
	// *   **Redirect**: redirects a request.
	// *   **FixedResponse**: returns a custom response.
	// *   **Rewrite**: rewrites a request.
	// *   **InsertHeader**: inserts a header.
	// *   **RemoveHeaderConfig**: deletes a header.
	// *   **TrafficLimitConfig**: throttles network traffic.
	// *   **TrafficMirrorConfig**: mirrors traffic.
	// *   **CorsConfig**: forwards requests based on CORS.
	//
	// You can specify the final action and the actions that you want to perform before the final action:
	//
	// *   **FinalType**: the last action to be performed in a forwarding rule. Each forwarding rule can contain only one FinalType action. You can specify a **ForwardGroup**, **Redirect**, or **FixedResponse** action as the FinalType action.
	// *   **ExtType**: the action to be performed before the FinalType action. A forwarding rule can contain one or more ExtType actions. To specify this parameter, you must also specify FinalType. You can specify multiple **InsertHeader** actions or one **Rewrite** action.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsRedirectConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetHost(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetHttpCode(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetPath(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetPort(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetProtocol(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetQuery(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsRewriteConfig struct {
	// The queries per second (QPS). Valid values: **1 to 100000**.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The QPS per IP address. Valid values: **1 to 100000**.
	//
	// >  If both **QPS** and **PerIpQps** are set, make sure that the **QPS** value is smaller than the PerIpQps value.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The type of destination to which network traffic is mirrored. Valid values:
	//
	// *   **ForwardGroupMirror**: a server group
	// *   **SlsMirror**: Log Service
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsRewriteConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) SetHost(v string) *UpdateRuleAttributeRequestRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) SetPath(v string) *UpdateRuleAttributeRequestRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) SetQuery(v string) *UpdateRuleAttributeRequestRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig struct {
	// The allowed HTTP methods for CORS requests. Valid values:
	//
	// *   **GET**
	// *   **POST**
	// *   **PUT**
	// *   **DELETE**
	// *   **HEAD**
	// *   **OPTIONS**
	// *   **PATCH**
	PerIpQps *int32 `json:"PerIpQps,omitempty" xml:"PerIpQps,omitempty"`
	// The allowed origins of CORS requests. You can set this parameter to `*` or one or more values. A value cannot be `*`.
	//
	// *   A value must start with a `http://` or `https://`, followed by a valid domain name or a first-level wildcard domain name, such as `*.test.abc.example.com`.
	// *   You can specify a port in a value. Port range: **1** to **65535**.
	QPS *int32 `json:"QPS,omitempty" xml:"QPS,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) SetPerIpQps(v int32) *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig {
	s.PerIpQps = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) SetQPS(v int32) *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig {
	s.QPS = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig struct {
	MirrorGroupConfig *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig `json:"MirrorGroupConfig,omitempty" xml:"MirrorGroupConfig,omitempty" type:"Struct"`
	// The allowed headers for CORS requests. You can specify `*` or specify one or more values. Separate multiple values with commas (,). A value can contain only letters and digits. It cannot start or end with underscores (\_) or hyphens (-). It can be up to 32 characters in length.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) SetMirrorGroupConfig(v *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig {
	s.MirrorGroupConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) SetTargetType(v string) *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig {
	s.TargetType = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig struct {
	ServerGroupTuples []*UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) SetServerGroupTuples(v []*UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples struct {
	// The headers that can be exposed. You can specify `*` or specify one or more values. Separate multiple values with commas (,). A value can contain only letters and digits. It cannot start or end with underscores (\_) or hyphens (-). It can be up to 32 characters in length.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type UpdateRuleAttributeRequestRuleConditions struct {
	CookieConfig      *UpdateRuleAttributeRequestRuleConditionsCookieConfig      `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	HeaderConfig      *UpdateRuleAttributeRequestRuleConditionsHeaderConfig      `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	HostConfig        *UpdateRuleAttributeRequestRuleConditionsHostConfig        `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	MethodConfig      *UpdateRuleAttributeRequestRuleConditionsMethodConfig      `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	PathConfig        *UpdateRuleAttributeRequestRuleConditionsPathConfig        `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	QueryStringConfig *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	SourceIpConfig    *UpdateRuleAttributeRequestRuleConditionsSourceIpConfig    `json:"SourceIpConfig,omitempty" xml:"SourceIpConfig,omitempty" type:"Struct"`
	// The HTTP status code in the response. Valid values: **HTTP\_2xx**, **HTTP\_4xx**, and **HTTP\_5xx**. **x** must be a digit.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditions) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetCookieConfig(v *UpdateRuleAttributeRequestRuleConditionsCookieConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetHeaderConfig(v *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetHostConfig(v *UpdateRuleAttributeRequestRuleConditionsHostConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.HostConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetMethodConfig(v *UpdateRuleAttributeRequestRuleConditionsMethodConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetPathConfig(v *UpdateRuleAttributeRequestRuleConditionsPathConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.PathConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetQueryStringConfig(v *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetSourceIpConfig(v *UpdateRuleAttributeRequestRuleConditionsSourceIpConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.SourceIpConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetType(v string) *UpdateRuleAttributeRequestRuleConditions {
	s.Type = &v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsCookieConfig struct {
	Values []*UpdateRuleAttributeRequestRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfig) SetValues(v []*UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) *UpdateRuleAttributeRequestRuleConditionsCookieConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsCookieConfigValues struct {
	// The hostname. The hostname must meet the following requirements:
	//
	// *   The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), periods (.), asterisks (\*), and question marks (?).
	// *   The hostname must contain at least one period (.) but cannot start or end with a period (.).
	// *   The rightmost domain label can contain only letters, asterisks (\*), and question marks (?), and cannot contain digits or hyphens (-).
	// *   The domain labels cannot start or end with a hyphen (-). You can specify asterisks (∗) and question marks (?) anywhere in a domain label.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The request method.
	//
	// Valid values: **HEAD**, **GET**, **POST**, **OPTIONS**, **PUT**, **PATCH**, and **DELETE**.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) SetKey(v string) *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) SetValue(v string) *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsHeaderConfig struct {
	// The path. The path must meet the following requirements:
	//
	// *   It must be 1 to 128 characters in length.
	// *   It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ .+ / & ~ @ :`. It cannot contain the following special characters: `" % # ; ! ( ) [ ]^ , "`. You can use asterisks (\*) and question marks (?) as wildcards.
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) SetKey(v string) *UpdateRuleAttributeRequestRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsHostConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsHostConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsHostConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsMethodConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsMethodConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsMethodConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsMethodConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsPathConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsPathConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsPathConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsQueryStringConfig struct {
	Values []*UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) SetValues(v []*UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues struct {
	// The content of the custom response. The content can be up to 1 KB in size, and can contain only ASCII characters.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The content type.
	//
	// Valid values: **text/plain**, **text/css**, **text/html**, **application/javascript**, and **application/json**.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) SetKey(v string) *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) SetValue(v string) *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsSourceIpConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsSourceIpConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsSourceIpConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsSourceIpConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsSourceIpConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRuleAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeResponseBody) SetJobId(v string) *UpdateRuleAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateRuleAttributeResponseBody) SetRequestId(v string) *UpdateRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRuleAttributeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleAttributeResponse) SetStatusCode(v int32) *UpdateRuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleAttributeResponse) SetBody(v *UpdateRuleAttributeResponseBody) *UpdateRuleAttributeResponse {
	s.Body = v
	return s
}

type UpdateRulesAttributeRequest struct {
	// The content of the custom response. The content can be up to 1 KB in size and can contain only ASCII characters.
	ClientToken *string                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool                               `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Rules       []*UpdateRulesAttributeRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequest) SetClientToken(v string) *UpdateRulesAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRulesAttributeRequest) SetDryRun(v bool) *UpdateRulesAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateRulesAttributeRequest) SetRules(v []*UpdateRulesAttributeRequestRules) *UpdateRulesAttributeRequest {
	s.Rules = v
	return s
}

type UpdateRulesAttributeRequestRules struct {
	// The format of the response.
	//
	// Valid values: **text/plain**, **text/css**, **text/html**, **application/javascript**, and **application/json**.
	Priority       *int32                                            `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RuleActions    []*UpdateRulesAttributeRequestRulesRuleActions    `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	RuleConditions []*UpdateRulesAttributeRequestRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The ID of the asynchronous task.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The HTTP status code in the response. Valid values: **HTTP\_2xx**, **HTTP\_4xx**, and **HTTP\_5xx**. **x** must be a digit.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s UpdateRulesAttributeRequestRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRules) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRules) SetPriority(v int32) *UpdateRulesAttributeRequestRules {
	s.Priority = &v
	return s
}

func (s *UpdateRulesAttributeRequestRules) SetRuleActions(v []*UpdateRulesAttributeRequestRulesRuleActions) *UpdateRulesAttributeRequestRules {
	s.RuleActions = v
	return s
}

func (s *UpdateRulesAttributeRequestRules) SetRuleConditions(v []*UpdateRulesAttributeRequestRulesRuleConditions) *UpdateRulesAttributeRequestRules {
	s.RuleConditions = v
	return s
}

func (s *UpdateRulesAttributeRequestRules) SetRuleId(v string) *UpdateRulesAttributeRequestRules {
	s.RuleId = &v
	return s
}

func (s *UpdateRulesAttributeRequestRules) SetRuleName(v string) *UpdateRulesAttributeRequestRules {
	s.RuleName = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActions struct {
	CorsConfig          *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig          `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	FixedResponseConfig *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	ForwardGroupConfig  *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig  `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	InsertHeaderConfig  *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig  `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// The path to which requests are forwarded.
	//
	// *   Default value: **${path}**. **${host}**, **${protocol}**, and **${port}** are also supported. Each variable can be specified only once. You can specify one or more of the preceding variables in each request. You can also concatenate them with the following characters.
	//
	// *   A custom value. You must ensure that the custom value meets the following requirements:
	//
	//     *   The custom value must be 1 to 128 characters in length. You can use asterisks (\*) and question marks (?) as wildcards.
	//     *   The custom value can contain letters, digits, and the following special characters: `$ - _ . + / & ~ @ : \" * ?`. The custom value must start with a forward slash (/) and cannot contain the following characters: `" % # ; ! ( ) [ ] ^ , "`.
	Order               *int32                                                          `json:"Order,omitempty" xml:"Order,omitempty"`
	RedirectConfig      *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig      `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	RemoveHeaderConfig  *UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig  `json:"RemoveHeaderConfig,omitempty" xml:"RemoveHeaderConfig,omitempty" type:"Struct"`
	RewriteConfig       *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig       `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	TrafficLimitConfig  *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig  `json:"TrafficLimitConfig,omitempty" xml:"TrafficLimitConfig,omitempty" type:"Struct"`
	TrafficMirrorConfig *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig `json:"TrafficMirrorConfig,omitempty" xml:"TrafficMirrorConfig,omitempty" type:"Struct"`
	// The type of destination to which network traffic is mirrored. Valid values:
	//
	// *   **ForwardGroupMirror:** a server group.
	// *   **SlsMirror:** Log Service.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActions) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActions) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetCorsConfig(v *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.CorsConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetFixedResponseConfig(v *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetForwardGroupConfig(v *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetInsertHeaderConfig(v *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetOrder(v int32) *UpdateRulesAttributeRequestRulesRuleActions {
	s.Order = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetRedirectConfig(v *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetRemoveHeaderConfig(v *UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.RemoveHeaderConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetRewriteConfig(v *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetTrafficLimitConfig(v *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.TrafficLimitConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetTrafficMirrorConfig(v *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.TrafficMirrorConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetType(v string) *UpdateRulesAttributeRequestRulesRuleActions {
	s.Type = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsCorsConfig struct {
	// The value of the cookie. The value must be 1 to 128 characters in length, and can contain printable characters such as lowercase letters, asterisks (\*), and question marks (?). However, the value cannot contain uppercase letters, space characters, or the following special characters: `# [ ] { } \ | < > &`.
	AllowCredentials *string   `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	AllowHeaders     []*string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty" type:"Repeated"`
	AllowMethods     []*string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty" type:"Repeated"`
	AllowOrigin      []*string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty" type:"Repeated"`
	ExposeHeaders    []*string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty" type:"Repeated"`
	// The key of the header. The key must be 1 to 40 characters in length, and can contain lowercase letters, digits, hyphens (-), and underscores (\_). Cookie and Host are not supported.
	MaxAge *int64 `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetAllowCredentials(v string) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetAllowHeaders(v []*string) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.AllowHeaders = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetAllowMethods(v []*string) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.AllowMethods = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetAllowOrigin(v []*string) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.AllowOrigin = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetExposeHeaders(v []*string) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.ExposeHeaders = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetMaxAge(v int64) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.MaxAge = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig struct {
	// The server group to which requests are forwarded.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The weight of the server group. A larger value indicates a higher weight. A server group with a higher weight receives more requests. Valid values: **1** to **100**.
	//
	// *   When **N** is 1, the default value **100** is used.
	// *   When **N** is greater than 1, you must specify the **weight** for each server group.
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Specifies whether to enable session persistence. Valid values:
	//
	// *   **true**: enables session persistence.
	// *   **false** (default): disables session persistence.
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) SetContent(v string) *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) SetContentType(v string) *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) SetHttpCode(v string) *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig struct {
	ServerGroupStickySession *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession `json:"ServerGroupStickySession,omitempty" xml:"ServerGroupStickySession,omitempty" type:"Struct"`
	ServerGroupTuples        []*UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples      `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) SetServerGroupStickySession(v *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig {
	s.ServerGroupStickySession = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession struct {
	// The key of the header to be inserted. The key must be 1 to 40 characters in length, and can contain lowercase letters, digits, underscores (\_), and hyphens (-). The key specified in the `InsertHeader` parameter must be unique.
	//
	// >  You cannot set one of the following keys (case-insensitive): `slb-id`, `slb-ip`, `x-forwarded-for`, `x-forwarded-proto`, `x-forwarded-eip`, `x-forwarded-port`, `x-forwarded-client-srcport`, `connection`, `upgrade`, `content-length`, `transfer-encoding`, `keep-alive`, `te, host`, `cookie`, `remoteip`, and `authority`.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The value of the header to be inserted.
	//
	// *   If **ValueType** is set to **SystemDefined**, you can set one of the following values:
	//
	//     *   **ClientSrcPort:** the port of the client.
	//     *   **ClientSrcIp:** the IP address of the client.
	//     *   **Protocol:** the request protocol. You can set the protocol to HTTP or HTTPS.
	//     *   **SLBId:** the ID of the ALB instance.
	//     *   **SLBPort:** the listening port of the ALB instance.
	//
	// *   If **ValueType** is set to **UserDefined**, you can specify a custom header value. The header value must be 1 to 128 characters in length, and can contain printable characters whose ASCII values are `greater than or equal to 32 and lower than 127`. You can use asterisks (\*) and question marks (?) as wildcard characters. The value cannot start or end with a space character.
	//
	// *   If **ValueType** is set to **ReferenceHeader**, you can reference one of the request headers. The header value must be 1 to 128 characters in length, and can contain lowercase letters, digits, underscores (\_), and hyphens (-).
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) SetEnabled(v bool) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Enabled = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) SetTimeout(v int32) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Timeout = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	// The timeout period of sessions. Unit: seconds. Valid values: **1** to **86400**.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// Specifies whether to overwrite the request header. Valid values:
	//
	// *   **true** overwrites the request header.
	// *   **false** (default): does not overwrite the request header.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) SetWeight(v int32) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig struct {
	// The type of header. Valid values:
	//
	// *   **UserDefined:** a user-defined header.
	// *   **ReferenceHeader:** a header that references a field of a request header.
	// *   **SystemDefined:** a system-defined header.
	CoverEnabled *bool `json:"CoverEnabled,omitempty" xml:"CoverEnabled,omitempty"`
	// The priority of the action. Valid values: **1** to **50000**. A lower value indicates a higher priority. The actions of a forwarding rule are applied in descending order of priority. This parameter cannot be left empty. The priority of each action within a forwarding rule must be unique. You can specify priorities for at most 20 actions.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The hostname to which requests are forwarded.
	//
	// Take note of the following rules when you specify a hostname:
	//
	// *   The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, and the following special characters: - . \* = ~ \_ + \ ^ ! $ & | ( ) \[ ] ?.
	// *   The hostname must contain at least one period (.) but cannot start or end with a period (.).
	// *   The rightmost domain label can contain only letters, asterisks (\*), and question marks (?), and cannot contain digits or hyphens (-). The leftmost `domain label` can be an asterisk (\*).
	// *   The domain labels cannot start or end with a hyphen (-). You can specify asterisks (∗) and question marks (?) anywhere in a domain label.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The HTTP status code that indicates the redirect type. Valid values: **301**, **302**, **303**, **307**, and **308**.
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) SetCoverEnabled(v bool) *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig {
	s.CoverEnabled = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) SetValue(v string) *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) SetValueType(v string) *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig struct {
	// The port to which requests are forwarded.
	//
	// Valid values: **1** to **63335**.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The redirect protocol.
	//
	// Valid values: **HTTP** and **HTTPS**.
	//
	// >  HTTPS listeners do not support HTTPS to HTTP redirection.
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The query string to which requests are forwarded.
	//
	// The query string must be 1 to 128 characters in length, and can contain printable characters, excluding uppercase letters and the following special characters: `# [ ] { } \ | < > &`.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The hostname to which requests are forwarded.
	//
	// Take note of the following rules when you specify a hostname:
	//
	// *   The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, and the following special characters: - . \* = ~ \_ + \ ^ ! $ & | ( ) \[ ] ?.
	// *   The hostname must contain at least one period (.) but cannot start or end with a period (.).
	// *   The rightmost domain label can contain only letters, asterisks (\*), and question marks (?), and cannot contain digits or hyphens (-). The leftmost `domain label` can be an asterisk (\*).
	// *   The domain labels cannot start or end with a hyphen (-). You can specify asterisks (∗) and question marks (?) anywhere in a domain label.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The path to which requests are forwarded.
	//
	// *   Default value: **${path}**. **${host}**, **${protocol}**, and **${port}** are also supported. Each variable can be specified only once. You can specify one or more of the preceding variables in each request. You can also concatenate them with the following characters.
	//
	// *   A custom value. You must ensure that the custom value meets the following requirements:
	//
	//     *   The custom value must be 1 to 128 characters in length. You can use asterisks (\*) and question marks (?) as wildcards.
	//     *   The custom value can contain letters, digits, and the following special characters: `$ - _ . + / & ~ @ : \" * ?`. The custom value must start with a forward slash (/) and cannot contain the following characters: `" % # ; ! ( ) [ ] ^ , "`.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The query string to which requests are forwarded.
	//
	// The query string must be 1 to 128 characters in length, and can contain printable characters, excluding uppercase letters and the following special characters: `# [ ] { } \ | < > &`.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetHost(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetHttpCode(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetPath(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetPort(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetProtocol(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetQuery(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig struct {
	// The allowed headers for CORS requests. You can specify an asterisk (`*`) or one or more values. Separate multiple values with commas (,). A value can contain only letters and digits. It cannot start or end with underscores (\_) or hyphens (-). It can be up to 32 characters in length.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig {
	s.Key = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig struct {
	// The type of forwarding rule. You can specify at most seven types. Valid values:
	//
	// *   **Host**: Requests are forwarded based on hosts.
	// *   **Path**: Requests are forwarded based on paths.
	// *   **Header**: Requests are forwarded based on HTTP headers.
	// *   **QueryString**: Requests are forwarded based on query strings.
	// *   **Method**: Requests are forwarded based on request methods.
	// *   **Cookie**: Requests are forwarded based on cookies.
	// *   **SourceIp**: Requests are forwarded based on source IP addresses.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The number of queries per second (QPS). Valid values: **1** to **100000**.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The QPS per IP address. Valid values: **1** to **100000**.
	//
	// >  If both the QPS and PerIpQps properties are specified, make sure that the value of the QPS property is smaller than the value of the PerIpQps property.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) SetHost(v string) *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) SetPath(v string) *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) SetQuery(v string) *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig struct {
	// The key of the header.
	//
	// *   The key must be 1 to 40 characters in length,
	// *   and can contain lowercase letters, digits, hyphens (-), and underscores (\_).
	// *   Cookie and Host are not supported.
	PerIpQps *int32 `json:"PerIpQps,omitempty" xml:"PerIpQps,omitempty"`
	// The ID of the server group.
	QPS *int32 `json:"QPS,omitempty" xml:"QPS,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) SetPerIpQps(v int32) *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig {
	s.PerIpQps = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) SetQPS(v int32) *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig {
	s.QPS = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig struct {
	MirrorGroupConfig *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig `json:"MirrorGroupConfig,omitempty" xml:"MirrorGroupConfig,omitempty" type:"Struct"`
	// The allowed origins of CORS requests. You can set this parameter to an asterisk (`*`) or one or more values. The values cannot be asterisks (`*`).
	//
	// *   Each value must start with `http://` or `https://`, followed by a valid domain name or a first-level wildcard domain name, such as `*.test.abc.example.com`.
	// *   You can specify a port in a value. Port range: **1** to **65535**.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) SetMirrorGroupConfig(v *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig {
	s.MirrorGroupConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) SetTargetType(v string) *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig {
	s.TargetType = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig struct {
	ServerGroupTuples []*UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) SetServerGroupTuples(v []*UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples struct {
	// The allowed HTTP methods for CORS requests. Valid values:
	//
	// *   **GET**
	// *   **POST**
	// *   **PUT**
	// *   **DELETE**
	// *   **HEAD**
	// *   **OPTIONS**
	// *   **PATCH**
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditions struct {
	CookieConfig             *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig             `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	HeaderConfig             *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig             `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	HostConfig               *UpdateRulesAttributeRequestRulesRuleConditionsHostConfig               `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	MethodConfig             *UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig             `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	PathConfig               *UpdateRulesAttributeRequestRulesRuleConditionsPathConfig               `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	QueryStringConfig        *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig        `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	ResponseHeaderConfig     *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig     `json:"ResponseHeaderConfig,omitempty" xml:"ResponseHeaderConfig,omitempty" type:"Struct"`
	ResponseStatusCodeConfig *UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig `json:"ResponseStatusCodeConfig,omitempty" xml:"ResponseStatusCodeConfig,omitempty" type:"Struct"`
	SourceIpConfig           *UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig           `json:"SourceIpConfig,omitempty" xml:"SourceIpConfig,omitempty" type:"Struct"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// *   **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false**: sends the request. If the request passes the dry run, the system returns an `HTTP 2xx` status code and performs the operation. This is the default value.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetCookieConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetHeaderConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetHostConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsHostConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetMethodConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetPathConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsPathConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetQueryStringConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetResponseHeaderConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.ResponseHeaderConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetResponseStatusCodeConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.ResponseStatusCodeConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetSourceIpConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.SourceIpConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetType(v string) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.Type = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig struct {
	Values []*UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig) SetValues(v []*UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig {
	s.Values = v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues struct {
	// The value of the response header. The value must be 1 to 128 characters in length. It can contain printable characters whose ASCII values are `greater than or equal to 32 and smaller than 127`. For example, lowercase letters, asterisks (\*), and question marks (?). The value cannot start or end with a space character.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The hostname.
	//
	// Take note of the following rules when you specify a hostname:
	//
	// *   The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, and the following special characters: - . \* = ~ \_ + \ ^ ! $ & | ( ) \[ ] ?.
	// *   The hostname must contain at least one period (.) but cannot start or end with a period (.).
	// *   The rightmost domain label can contain only letters, asterisks (\*), and question marks (?), and cannot contain digits or hyphens (-). The leftmost `domain label` can be an asterisk (\*).
	// *   The domain labels cannot start or end with a hyphen (-). You can specify asterisks (∗) and question marks (?) anywhere in a domain label.
	// *   If you specify an exact hostname or a wildcard hostname, the hostname cannot start with a tilde (~).
	// *   If you specify a regular expression, the regular expression cannot start with an asterisk (\*). The regular expression is not case-sensitive.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) SetValue(v string) *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig struct {
	// The request methods.
	//
	// Valid values: **HEAD**, **GET**, **POST**, **OPTIONS**, **PUT**, **PATCH**, and **DELETE**.
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditionsHostConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsHostConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig {
	s.Values = v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditionsPathConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsPathConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig struct {
	Values []*UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig) SetValues(v []*UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues struct {
	// The value of the response header.
	//
	// *   The value must be 1 to 128 characters in length.
	// *   It can contain printable characters whose ASCII values are `greater than or equal to 32 and smaller than 127`. For example, lowercase letters, asterisks (\*), and question marks (?).
	// *   The value cannot start or end with a space character.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The status code in responses. Valid values: **100** to **599**.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) SetValue(v string) *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig struct {
	// The type of action specified in the forwarding rule. You can specify at most 11 types of actions. Valid values:
	//
	// *   **ForwardGroup:** forwards a request to multiple vServer groups.
	// *   **Redirect:** redirects a request.
	// *   **FixedResponse:** returns a custom response.
	// *   **Rewrite:** rewrites a request.
	// *   **InsertHeader:** adds a header to a request.
	// *   **RemoveHeaderConfig:** deletes the header of a request.
	// *   **TrafficLimitConfig:** throttles traffic.
	// *   **TrafficMirrorConfig:** mirrors traffic.
	// *   **CorsConfig:** forwards requests based on CORS.
	//
	// The preceding actions can be classified into two broad types:
	//
	// *   **FinalType:** the last action to be performed in a forwarding rule. Each forwarding rule can contain only one FinalType action. You can specify a **ForwardGroup**, **Redirect**, or **FixedResponse** action as the FinalType action.
	// *   **ExtType:** the actions to be performed before the FinalType action. A forwarding rule can contain one or more ExtType actions. To specify this parameter, you must also specify FinalType. You can specify multiple **InsertHeader** actions or one **Rewrite** action.
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig {
	s.Values = v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig {
	s.Values = v
	return s
}

type UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig {
	s.Values = v
	return s
}

type UpdateRulesAttributeResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRulesAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeResponseBody) SetJobId(v string) *UpdateRulesAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateRulesAttributeResponseBody) SetRequestId(v string) *UpdateRulesAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRulesAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRulesAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRulesAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRulesAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeResponse) SetHeaders(v map[string]*string) *UpdateRulesAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRulesAttributeResponse) SetStatusCode(v int32) *UpdateRulesAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRulesAttributeResponse) SetBody(v *UpdateRulesAttributeResponseBody) *UpdateRulesAttributeResponse {
	s.Body = v
	return s
}

type UpdateSecurityPolicyAttributeRequest struct {
	// The supported cipher suites.
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The security policy ID.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The name of the security policy.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	// The supported TLS protocol versions.
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s UpdateSecurityPolicyAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeRequest) SetCiphers(v []*string) *UpdateSecurityPolicyAttributeRequest {
	s.Ciphers = v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetClientToken(v string) *UpdateSecurityPolicyAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetDryRun(v bool) *UpdateSecurityPolicyAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetSecurityPolicyId(v string) *UpdateSecurityPolicyAttributeRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetSecurityPolicyName(v string) *UpdateSecurityPolicyAttributeRequest {
	s.SecurityPolicyName = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetTLSVersions(v []*string) *UpdateSecurityPolicyAttributeRequest {
	s.TLSVersions = v
	return s
}

type UpdateSecurityPolicyAttributeResponseBody struct {
	// The asynchronous task ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSecurityPolicyAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetJobId(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetRequestId(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSecurityPolicyAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSecurityPolicyAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSecurityPolicyAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeResponse) SetHeaders(v map[string]*string) *UpdateSecurityPolicyAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponse) SetStatusCode(v int32) *UpdateSecurityPolicyAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponse) SetBody(v *UpdateSecurityPolicyAttributeResponseBody) *UpdateSecurityPolicyAttributeResponse {
	s.Body = v
	return s
}

type UpdateServerGroupAttributeRequest struct {
	// The configuration of health checks.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The backend port that is used for health checks.
	//
	// Valid values: **0** to **65535**.
	//
	// If you set the value to **0**, the ports of backend servers are used for health checks.
	//
	// >  This parameter takes effect when the **HealthCheckEnabled** parameter is set to **true**.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// *   **true** (default): enables the feature.
	// *   **false**: disables the feature.
	HealthCheckConfig *UpdateServerGroupAttributeRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, an HTTP `2xx` status code is returned and the operation is performed.
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The ID of the asynchronous task.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses RequestId as ClientToken. The value of RequestId for each API request may be different.
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The ID of the request.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The timeout period of a cookie. Unit: seconds.
	//
	// Valid values: **1** to **86400**.
	//
	// >  This parameter takes effect when the **StickySessionEnabled** parameter is set to **true** and the **StickySessionType** parameter is set to **Insert**.
	StickySessionConfig *UpdateServerGroupAttributeRequestStickySessionConfig `json:"StickySessionConfig,omitempty" xml:"StickySessionConfig,omitempty" type:"Struct"`
	// url一致性hash参数配置
	UchConfig *UpdateServerGroupAttributeRequestUchConfig `json:"UchConfig,omitempty" xml:"UchConfig,omitempty" type:"Struct"`
}

func (s UpdateServerGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequest) SetClientToken(v string) *UpdateServerGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetDryRun(v bool) *UpdateServerGroupAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetHealthCheckConfig(v *UpdateServerGroupAttributeRequestHealthCheckConfig) *UpdateServerGroupAttributeRequest {
	s.HealthCheckConfig = v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetScheduler(v string) *UpdateServerGroupAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetServerGroupId(v string) *UpdateServerGroupAttributeRequest {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetServerGroupName(v string) *UpdateServerGroupAttributeRequest {
	s.ServerGroupName = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetServiceName(v string) *UpdateServerGroupAttributeRequest {
	s.ServiceName = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetStickySessionConfig(v *UpdateServerGroupAttributeRequestStickySessionConfig) *UpdateServerGroupAttributeRequest {
	s.StickySessionConfig = v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetUchConfig(v *UpdateServerGroupAttributeRequestUchConfig) *UpdateServerGroupAttributeRequest {
	s.UchConfig = v
	return s
}

type UpdateServerGroupAttributeRequestHealthCheckConfig struct {
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// Valid values: **1** to **50**.
	//
	// >  This parameter takes effect when the **HealthCheckEnabled** parameter is set to **true**.
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The domain name that is used for health checks. The domain name must meet the following requirements:
	//
	// *   The domain name must be 1 to 80 characters in length.
	// *   The domain name can contain lowercase letters, digits, hyphens (-), and periods (.).
	// *   It must contain at least one period (.) but cannot start or end with a period (.).
	// *   The rightmost field of the domain name can contain only letters and cannot contain digits or hyphens (-).
	// *   Other fields cannot start or end with a hyphen (-).
	//
	// >  This parameter takes effect when the **HealthCheckEnabled** parameter is set to true and the **HealthCheckProtocol** parameter is set to **HTTP**.
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The HTTP status codes that indicate a successful health check.
	//
	// *   If **HealthCheckProtocol** is set to **HTTP**, **HealthCheckCodes** can be set to **http\_2xx** (default), **http\_3xx**, **http\_4xx**, and **http\_5xx**. Separate multiple HTTP status codes with commas (,).
	// *   If **HealthCheckProtocol** is set to **gRPC**, **HealthCheckCodes** can be set to **0 to 99**. Default value: **0**. Value ranges are supported. You can enter up to 20 value ranges. Separate multiple value ranges with commas (,).
	//
	// >  This parameter takes effect when the **HealthCheckEnabled** parameter is set to **true** and the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The version of HTTP that is used for health checks. Valid values:
	//
	// *   **HTTP1.0**
	// *   **HTTP1.1**
	//
	// >  This parameter takes effect when the **HealthCheckEnabled** parameter is set to true and the **HealthCheckProtocol** parameter is set to **HTTP**.
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The HTTP method that is used for health checks. Valid values:
	//
	// *   **GET**: If the length of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	// *   **POST**: By default, gRPC health checks use the POST method.
	// *   **HEAD**: By default, HTTP health checks use the HEAD method.
	//
	// >  This parameter takes effect when the **HealthCheckEnabled** parameter is set to true and the **HealthCheckProtocol** parameter is set to **HTTP** or **gRPC**.
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The path that is used for health checks.
	//
	// The path must be 1 to 80 characters in length and can contain only letters, digits, and the following special characters: `- / . % ? # & =`. It can also contain the following extended characters: `_ ; ~ ! ( ) * [ ] @ $ ^ : \" , +`. The path must start with a forward slash (`/`).
	//
	// >  This parameter takes effect when the **HealthCheckEnabled** parameter is set to **true** and the **HealthCheckProtocol** parameter is set to **HTTP**.
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// *   **HTTP**: To perform HTTP health checks, Application Load Balancer (ALB) sends HEAD or GET requests to a backend server to check whether the backend server is healthy.
	// *   **TCP**: To perform TCP health checks, ALB sends SYN packets to the backend server to check whether the port of the backend server is available to receive requests.
	// *   **gRPC**: To perform gRPC health checks, ALB sends POST or GET requests to a backend server to check whether the backend server is healthy.
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The timeout period of a health check response. If a backend server, such as an Elastic Compute Service (ECS) instance, does not return a health check response within the specified timeout period, the server fails the health check. Unit: seconds.
	//
	// Valid values: **1** to **300**.
	//
	// >
	// *   If the value of the **HealthCheckTimeout** parameter is smaller than that of the **HealthCheckInterval** parameter, the timeout period specified by the **HealthCheckTimeout** parameter is ignored and the value of the **HealthCheckInterval** parameter is used as the timeout period.
	// *   This parameter takes effect when the **HealthCheckEnabled** parameter is set to **true**.
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail** to **success**.
	//
	// Valid values: **2** to **10**.
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success** to **fail**.
	//
	// Valid values: **2** to **10**.
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The configuration of session persistence.
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The cookie that is configured on the server.
	//
	// The cookie must be 1 to 200 characters in length and can contain only ASCII characters and digits. It cannot contain commas (,), semicolons (;), or space characters. It cannot start with a dollar sign ($).
	//
	// >  This parameter takes effect when the **StickySessionEnabled** parameter is set to **true** and the **StickySessionType** parameter is set to **Server**.
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckCodes(v []*string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckCodes = v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckConnectPort(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckEnabled(v bool) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckHost(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckHost = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckHttpVersion(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckMethod(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckPath(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckPath = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckProtocol(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckProtocol = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckTimeout(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckTimeout = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthyThreshold(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

type UpdateServerGroupAttributeRequestStickySessionConfig struct {
	// Specifies whether to enable session persistence. Valid values:
	//
	// *   **true**: enables session persistence.
	// *   **false** (default): disables session persistence.
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The method that is used to handle a cookie. Valid values:
	//
	// *   **Insert**: inserts a cookie.
	//
	//     ALB inserts a cookie (SERVERID) into the first HTTP or HTTPS response packet that is sent to a client. The next request from the client contains this cookie and the listener forwards this request to the recorded backend server.
	//
	// *   **Server**: rewrites a cookie.
	//
	//     When ALB detects a user-defined cookie, it overwrites the original cookie with the user-defined cookie. Subsequent requests to ALB carry this user-defined cookie, and ALB determines the destination servers of the requests based on the cookies.
	//
	// >  This parameter takes effect when the **StickySessionEnabled** parameter is set to **true**.
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The ID of the server group.
	StickySessionEnabled *bool `json:"StickySessionEnabled,omitempty" xml:"StickySessionEnabled,omitempty"`
	// This parameter is available only if the ALB Ingress controller is used. In this case, set this parameter to the name of the `Kubernetes Service` that is associated with the server group.
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
}

func (s UpdateServerGroupAttributeRequestStickySessionConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestStickySessionConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetCookie(v string) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.Cookie = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetCookieTimeout(v int32) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.CookieTimeout = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetStickySessionEnabled(v bool) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.StickySessionEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetStickySessionType(v string) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.StickySessionType = &v
	return s
}

type UpdateServerGroupAttributeRequestUchConfig struct {
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 一致性hash参数值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateServerGroupAttributeRequestUchConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestUchConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestUchConfig) SetType(v string) *UpdateServerGroupAttributeRequestUchConfig {
	s.Type = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestUchConfig) SetValue(v string) *UpdateServerGroupAttributeRequestUchConfig {
	s.Value = &v
	return s
}

type UpdateServerGroupAttributeResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServerGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeResponseBody) SetJobId(v string) *UpdateServerGroupAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetRequestId(v string) *UpdateServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServerGroupAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServerGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServerGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateServerGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServerGroupAttributeResponse) SetStatusCode(v int32) *UpdateServerGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServerGroupAttributeResponse) SetBody(v *UpdateServerGroupAttributeResponseBody) *UpdateServerGroupAttributeResponse {
	s.Body = v
	return s
}

type UpdateServerGroupServersAttributeRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The type of the backend server. You can specify at most 40 servers in each call. Valid values:
	//
	// *   **Ecs**: an ECS instance
	// *   **Eni**: an ENI
	// *   **Eci**: an elastic container instance
	// *   **Ip**: an IP address
	// *   **Fc**: a function
	ServerGroupId *string                                            `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	Servers       []*UpdateServerGroupServersAttributeRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s UpdateServerGroupServersAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeRequest) SetClientToken(v string) *UpdateServerGroupServersAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetDryRun(v bool) *UpdateServerGroupServersAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetServerGroupId(v string) *UpdateServerGroupServersAttributeRequest {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetServers(v []*UpdateServerGroupServersAttributeRequestServers) *UpdateServerGroupServersAttributeRequest {
	s.Servers = v
	return s
}

type UpdateServerGroupServersAttributeRequestServers struct {
	// The weight of the backend server. Valid values: **0** to **100**. Default value: **100**. If the weight of a backend server is set to **0**, no requests are forwarded to the backend server. You can specify at most 40 servers in each call.
	//
	// >  You do not need to set this parameter if **ServerType** is set to **Fc**.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId** as **ClientToken**. The value of **RequestId** may be different for each API request.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The ID of the asynchronous task.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The ID of the request.
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	Weight     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateServerGroupServersAttributeRequestServers) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeRequestServers) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetDescription(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.Description = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetPort(v int32) *UpdateServerGroupServersAttributeRequestServers {
	s.Port = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerId(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerIp(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerIp = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerType(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerType = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetWeight(v int32) *UpdateServerGroupServersAttributeRequestServers {
	s.Weight = &v
	return s
}

type UpdateServerGroupServersAttributeResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServerGroupServersAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetJobId(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetRequestId(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServerGroupServersAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServerGroupServersAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServerGroupServersAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeResponse) SetHeaders(v map[string]*string) *UpdateServerGroupServersAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServerGroupServersAttributeResponse) SetStatusCode(v int32) *UpdateServerGroupServersAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponse) SetBody(v *UpdateServerGroupServersAttributeResponseBody) *UpdateServerGroupServersAttributeResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("alb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Each ACL can contain IP addresses or CIDR blocks. Take note of the following limits on ACLs:
 *     *   The maximum number of IP entries that can be added to an ACL with each Alibaba Cloud account at a time: 20
 *     *   The maximum number of IP entries that each ACL can contain: 1,000
 * *   **AddEntriesToAcl** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAclEntries](~~213616~~) operation to query the status of the task.
 *     *   If the ACL is in the **Adding** state, the IP entries are being added.
 *     *   If the ACL is in the **Available** state, the IP entries are added.
 *
 * @param request AddEntriesToAclRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddEntriesToAclResponse
 */
func (client *Client) AddEntriesToAclWithOptions(request *AddEntriesToAclRequest, runtime *util.RuntimeOptions) (_result *AddEntriesToAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclEntries)) {
		query["AclEntries"] = request.AclEntries
	}

	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddEntriesToAcl"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddEntriesToAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Each ACL can contain IP addresses or CIDR blocks. Take note of the following limits on ACLs:
 *     *   The maximum number of IP entries that can be added to an ACL with each Alibaba Cloud account at a time: 20
 *     *   The maximum number of IP entries that each ACL can contain: 1,000
 * *   **AddEntriesToAcl** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAclEntries](~~213616~~) operation to query the status of the task.
 *     *   If the ACL is in the **Adding** state, the IP entries are being added.
 *     *   If the ACL is in the **Available** state, the IP entries are added.
 *
 * @param request AddEntriesToAclRequest
 * @return AddEntriesToAclResponse
 */
func (client *Client) AddEntriesToAcl(request *AddEntriesToAclRequest) (_result *AddEntriesToAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddEntriesToAclResponse{}
	_body, _err := client.AddEntriesToAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * **AddServersToServerGroup** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
 * 1.  You can call the [ListServerGroups](~~213627~~) operation to query the status of a server group.
 * *   If a server group is in the **Configuring** state, it indicates that the server group is being modified.
 * *   If a server group is in the **Available** state, it indicates that the server group is running.
 * 2.  You can call the [ListServerGroupServers](~~213628~~) operation to query the status of a backend server.
 * *   If a backend server is in the **Adding** state, it indicates that the backend server is being added to a server group.
 * *   If a backend server is in the **Available** state, it indicates that the server is running.
 *
 * @param request AddServersToServerGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddServersToServerGroupResponse
 */
func (client *Client) AddServersToServerGroupWithOptions(request *AddServersToServerGroupRequest, runtime *util.RuntimeOptions) (_result *AddServersToServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Servers)) {
		query["Servers"] = request.Servers
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddServersToServerGroup"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddServersToServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * **AddServersToServerGroup** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
 * 1.  You can call the [ListServerGroups](~~213627~~) operation to query the status of a server group.
 * *   If a server group is in the **Configuring** state, it indicates that the server group is being modified.
 * *   If a server group is in the **Available** state, it indicates that the server group is running.
 * 2.  You can call the [ListServerGroupServers](~~213628~~) operation to query the status of a backend server.
 * *   If a backend server is in the **Adding** state, it indicates that the backend server is being added to a server group.
 * *   If a backend server is in the **Available** state, it indicates that the server is running.
 *
 * @param request AddServersToServerGroupRequest
 * @return AddServersToServerGroupResponse
 */
func (client *Client) AddServersToServerGroup(request *AddServersToServerGroupRequest) (_result *AddServersToServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddServersToServerGroupResponse{}
	_body, _err := client.AddServersToServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyHealthCheckTemplateToServerGroupWithOptions(request *ApplyHealthCheckTemplateToServerGroupRequest, runtime *util.RuntimeOptions) (_result *ApplyHealthCheckTemplateToServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTemplateId)) {
		query["HealthCheckTemplateId"] = request.HealthCheckTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyHealthCheckTemplateToServerGroup"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyHealthCheckTemplateToServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyHealthCheckTemplateToServerGroup(request *ApplyHealthCheckTemplateToServerGroupRequest) (_result *ApplyHealthCheckTemplateToServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyHealthCheckTemplateToServerGroupResponse{}
	_body, _err := client.ApplyHealthCheckTemplateToServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Specifies whether only to precheck this request. Valid values:
 * *   **true**: sends the precheck request but does not associate the ACLs with the listener. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
 * *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
 *
 * @param request AssociateAclsWithListenerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AssociateAclsWithListenerResponse
 */
func (client *Client) AssociateAclsWithListenerWithOptions(request *AssociateAclsWithListenerRequest, runtime *util.RuntimeOptions) (_result *AssociateAclsWithListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclIds)) {
		query["AclIds"] = request.AclIds
	}

	if !tea.BoolValue(util.IsUnset(request.AclType)) {
		query["AclType"] = request.AclType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateAclsWithListener"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateAclsWithListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Specifies whether only to precheck this request. Valid values:
 * *   **true**: sends the precheck request but does not associate the ACLs with the listener. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
 * *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
 *
 * @param request AssociateAclsWithListenerRequest
 * @return AssociateAclsWithListenerResponse
 */
func (client *Client) AssociateAclsWithListener(request *AssociateAclsWithListenerRequest) (_result *AssociateAclsWithListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateAclsWithListenerResponse{}
	_body, _err := client.AssociateAclsWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The **AssociateAdditionalCertificatesWithListener** operation is asynchronous. After you send a request, the system returns a request ID, but the operation is still being performed in the background. You can call the [GetListenerAttribute](~~2254865~~) operation to query the status of an additional certificate:
 * *   If the HTTPS or QUIC listener is in the **Associating** state, the additional certificate is being added.
 * *   If the HTTPS or QUIC listener is in the **Associated** state, the additional certificate is added.
 *
 * @param request AssociateAdditionalCertificatesWithListenerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AssociateAdditionalCertificatesWithListenerResponse
 */
func (client *Client) AssociateAdditionalCertificatesWithListenerWithOptions(request *AssociateAdditionalCertificatesWithListenerRequest, runtime *util.RuntimeOptions) (_result *AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Certificates)) {
		query["Certificates"] = request.Certificates
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateAdditionalCertificatesWithListener"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The **AssociateAdditionalCertificatesWithListener** operation is asynchronous. After you send a request, the system returns a request ID, but the operation is still being performed in the background. You can call the [GetListenerAttribute](~~2254865~~) operation to query the status of an additional certificate:
 * *   If the HTTPS or QUIC listener is in the **Associating** state, the additional certificate is being added.
 * *   If the HTTPS or QUIC listener is in the **Associated** state, the additional certificate is added.
 *
 * @param request AssociateAdditionalCertificatesWithListenerRequest
 * @return AssociateAdditionalCertificatesWithListenerResponse
 */
func (client *Client) AssociateAdditionalCertificatesWithListener(request *AssociateAdditionalCertificatesWithListenerRequest) (_result *AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.AssociateAdditionalCertificatesWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Specifies whether only to precheck the request. Valid values:
 * *   **true**: sends the precheck request but does not associate the EIP bandwidth plan with the ALB instance. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
 * *   **false** (default): sends the API request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
 *
 * @param request AttachCommonBandwidthPackageToLoadBalancerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AttachCommonBandwidthPackageToLoadBalancerResponse
 */
func (client *Client) AttachCommonBandwidthPackageToLoadBalancerWithOptions(request *AttachCommonBandwidthPackageToLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *AttachCommonBandwidthPackageToLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachCommonBandwidthPackageToLoadBalancer"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachCommonBandwidthPackageToLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Specifies whether only to precheck the request. Valid values:
 * *   **true**: sends the precheck request but does not associate the EIP bandwidth plan with the ALB instance. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
 * *   **false** (default): sends the API request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
 *
 * @param request AttachCommonBandwidthPackageToLoadBalancerRequest
 * @return AttachCommonBandwidthPackageToLoadBalancerResponse
 */
func (client *Client) AttachCommonBandwidthPackageToLoadBalancer(request *AttachCommonBandwidthPackageToLoadBalancerRequest) (_result *AttachCommonBandwidthPackageToLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachCommonBandwidthPackageToLoadBalancerResponse{}
	_body, _err := client.AttachCommonBandwidthPackageToLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Specifies whether to enable the AScript rule. Valid values:
 * *   **true**: enables the AScript rule.
 * *   **false** (default): disables the AScript rule.
 *
 * @param request CreateAScriptsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateAScriptsResponse
 */
func (client *Client) CreateAScriptsWithOptions(request *CreateAScriptsRequest, runtime *util.RuntimeOptions) (_result *CreateAScriptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AScripts)) {
		query["AScripts"] = request.AScripts
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAScripts"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAScriptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Specifies whether to enable the AScript rule. Valid values:
 * *   **true**: enables the AScript rule.
 * *   **false** (default): disables the AScript rule.
 *
 * @param request CreateAScriptsRequest
 * @return CreateAScriptsResponse
 */
func (client *Client) CreateAScripts(request *CreateAScriptsRequest) (_result *CreateAScriptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAScriptsResponse{}
	_body, _err := client.CreateAScriptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Specifies whether to check the request without performing the operation. Valid values:
 * *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
 * *   **false** (default): sends the request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
 *
 * @param request CreateAclRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateAclResponse
 */
func (client *Client) CreateAclWithOptions(request *CreateAclRequest, runtime *util.RuntimeOptions) (_result *CreateAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclName)) {
		query["AclName"] = request.AclName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAcl"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Specifies whether to check the request without performing the operation. Valid values:
 * *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
 * *   **false** (default): sends the request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
 *
 * @param request CreateAclRequest
 * @return CreateAclResponse
 */
func (client *Client) CreateAcl(request *CreateAclRequest) (_result *CreateAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAclResponse{}
	_body, _err := client.CreateAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHealthCheckTemplateWithOptions(request *CreateHealthCheckTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateHealthCheckTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckCodes)) {
		query["HealthCheckCodes"] = request.HealthCheckCodes
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectPort)) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckHost)) {
		query["HealthCheckHost"] = request.HealthCheckHost
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckHttpVersion)) {
		query["HealthCheckHttpVersion"] = request.HealthCheckHttpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckInterval)) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckMethod)) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckPath)) {
		query["HealthCheckPath"] = request.HealthCheckPath
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckProtocol)) {
		query["HealthCheckProtocol"] = request.HealthCheckProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTemplateName)) {
		query["HealthCheckTemplateName"] = request.HealthCheckTemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTimeout)) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthyThreshold)) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.UnhealthyThreshold)) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHealthCheckTemplate"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHealthCheckTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHealthCheckTemplate(request *CreateHealthCheckTemplateRequest) (_result *CreateHealthCheckTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHealthCheckTemplateResponse{}
	_body, _err := client.CreateHealthCheckTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * **CreateListener** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetListenerAttribute](~~214353~~) operation to query the status of the HTTP, HTTPS, or QUIC listener.
 * *   If the HTTP, HTTPS, or QUIC listener is in the **Provisioning** state, it indicates that the listener is being created.
 * *   If the HTTP, HTTPS, or QUIC listener is in the **Running** state, it indicates that the listener has been created.
 *
 * @param request CreateListenerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateListenerResponse
 */
func (client *Client) CreateListenerWithOptions(request *CreateListenerRequest, runtime *util.RuntimeOptions) (_result *CreateListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaCertificates)) {
		query["CaCertificates"] = request.CaCertificates
	}

	if !tea.BoolValue(util.IsUnset(request.CaEnabled)) {
		query["CaEnabled"] = request.CaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Certificates)) {
		query["Certificates"] = request.Certificates
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultActions)) {
		query["DefaultActions"] = request.DefaultActions
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.GzipEnabled)) {
		query["GzipEnabled"] = request.GzipEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Http2Enabled)) {
		query["Http2Enabled"] = request.Http2Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerDescription)) {
		query["ListenerDescription"] = request.ListenerDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.QuicConfig)) {
		query["QuicConfig"] = request.QuicConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RequestTimeout)) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedForConfig)) {
		query["XForwardedForConfig"] = request.XForwardedForConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateListener"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * **CreateListener** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetListenerAttribute](~~214353~~) operation to query the status of the HTTP, HTTPS, or QUIC listener.
 * *   If the HTTP, HTTPS, or QUIC listener is in the **Provisioning** state, it indicates that the listener is being created.
 * *   If the HTTP, HTTPS, or QUIC listener is in the **Running** state, it indicates that the listener has been created.
 *
 * @param request CreateListenerRequest
 * @return CreateListenerResponse
 */
func (client *Client) CreateListener(request *CreateListenerRequest) (_result *CreateListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateListenerResponse{}
	_body, _err := client.CreateListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * **CreateLoadBalancer** is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](~~214362~~) operation to query the status of an ALB instance.
 * *   If an ALB instance is in the **Provisioning** state, it indicates that the ALB instance is being created.
 * *   If an ALB instance is in the **Active** state, it indicates that the ALB instance is created.
 *
 * @param request CreateLoadBalancerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateLoadBalancerResponse
 */
func (client *Client) CreateLoadBalancerWithOptions(request *CreateLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *CreateLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressAllocatedMode)) {
		query["AddressAllocatedMode"] = request.AddressAllocatedMode
	}

	if !tea.BoolValue(util.IsUnset(request.AddressIpVersion)) {
		query["AddressIpVersion"] = request.AddressIpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeletionProtectionEnabled)) {
		query["DeletionProtectionEnabled"] = request.DeletionProtectionEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerBillingConfig)) {
		query["LoadBalancerBillingConfig"] = request.LoadBalancerBillingConfig
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerEdition)) {
		query["LoadBalancerEdition"] = request.LoadBalancerEdition
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerName)) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	if !tea.BoolValue(util.IsUnset(request.ModificationProtectionConfig)) {
		query["ModificationProtectionConfig"] = request.ModificationProtectionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneMappings)) {
		query["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoadBalancer"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * **CreateLoadBalancer** is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](~~214362~~) operation to query the status of an ALB instance.
 * *   If an ALB instance is in the **Provisioning** state, it indicates that the ALB instance is being created.
 * *   If an ALB instance is in the **Active** state, it indicates that the ALB instance is created.
 *
 * @param request CreateLoadBalancerRequest
 * @return CreateLoadBalancerResponse
 */
func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (_result *CreateLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CreateLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Take note of the following limits:
 * *   When you configure the **Redirect** action, you can use the default value only for the **HttpCode** parameter. Do not use the default values for the other parameters.
 * *   If you specify the **Rewrite** action together with other actions in a forwarding rule, make sure that the **ForwardGroup** action is specified.
 * *   **CreateRule** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListRules](~~214379~~) operation to query the status of a forwarding rule.
 *     *   If a forwarding rule is in the **Provisioning** state, the forwarding rule is being created.
 *     *   If a forwarding rule is in the **Available** state, the forwarding rule is created.
 * *   You can set **RuleConditions** and **RuleActions** to add conditions and actions to a forwarding rule. The limits on conditions and actions are:
 *     *   Limits on conditions: You can specify at most 5 conditions for a basic Application Load Balancer (ALB) instance, at most 10 conditions for a standard ALB instance, and at most 10 conditions for a WAF-enabled ALB instance.
 *     *   Limits on actions: You can specify at most 3 actions for a basic ALB instance, at most 5 actions for a standard ALB instance, and at most 5 actions for a WAF-enabled ALB instance.
 *
 * @param request CreateRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateRuleResponse
 */
func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RuleActions)) {
		query["RuleActions"] = request.RuleActions
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditions)) {
		query["RuleConditions"] = request.RuleConditions
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRule"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Take note of the following limits:
 * *   When you configure the **Redirect** action, you can use the default value only for the **HttpCode** parameter. Do not use the default values for the other parameters.
 * *   If you specify the **Rewrite** action together with other actions in a forwarding rule, make sure that the **ForwardGroup** action is specified.
 * *   **CreateRule** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListRules](~~214379~~) operation to query the status of a forwarding rule.
 *     *   If a forwarding rule is in the **Provisioning** state, the forwarding rule is being created.
 *     *   If a forwarding rule is in the **Available** state, the forwarding rule is created.
 * *   You can set **RuleConditions** and **RuleActions** to add conditions and actions to a forwarding rule. The limits on conditions and actions are:
 *     *   Limits on conditions: You can specify at most 5 conditions for a basic Application Load Balancer (ALB) instance, at most 10 conditions for a standard ALB instance, and at most 10 conditions for a WAF-enabled ALB instance.
 *     *   Limits on actions: You can specify at most 3 actions for a basic ALB instance, at most 5 actions for a standard ALB instance, and at most 5 actions for a WAF-enabled ALB instance.
 *
 * @param request CreateRuleRequest
 * @return CreateRuleResponse
 */
func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Specifies whether to precheck the request without performing the operation. Valid values:
 * *   **true**: prechecks the request but does not create the forwarding rule. The system checks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
 * *   **false** (default): sends the request. If the request passes the precheck, the system returns an `HTTP 2xx` status code and creates the forwarding rule.
 *
 * @param request CreateRulesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateRulesResponse
 */
func (client *Client) CreateRulesWithOptions(request *CreateRulesRequest, runtime *util.RuntimeOptions) (_result *CreateRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.Rules)) {
		query["Rules"] = request.Rules
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRules"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Specifies whether to precheck the request without performing the operation. Valid values:
 * *   **true**: prechecks the request but does not create the forwarding rule. The system checks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
 * *   **false** (default): sends the request. If the request passes the precheck, the system returns an `HTTP 2xx` status code and creates the forwarding rule.
 *
 * @param request CreateRulesRequest
 * @return CreateRulesResponse
 */
func (client *Client) CreateRules(request *CreateRulesRequest) (_result *CreateRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRulesResponse{}
	_body, _err := client.CreateRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecurityPolicyWithOptions(request *CreateSecurityPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateSecurityPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ciphers)) {
		query["Ciphers"] = request.Ciphers
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyName)) {
		query["SecurityPolicyName"] = request.SecurityPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.TLSVersions)) {
		query["TLSVersions"] = request.TLSVersions
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSecurityPolicy"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSecurityPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecurityPolicy(request *CreateSecurityPolicyRequest) (_result *CreateSecurityPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecurityPolicyResponse{}
	_body, _err := client.CreateSecurityPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * **CreateServerGroup** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListServerGroups](~~213627~~) operation to query the status of the task.
 * *   If a server group is in the **Creating** state, it indicates that the server group is being created.
 * *   If a server group is in the **Available** state, it indicates that the server group is created.
 *
 * @param request CreateServerGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateServerGroupResponse
 */
func (client *Client) CreateServerGroupWithOptions(request *CreateServerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConfig)) {
		query["HealthCheckConfig"] = request.HealthCheckConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		query["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupName)) {
		query["ServerGroupName"] = request.ServerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupType)) {
		query["ServerGroupType"] = request.ServerGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StickySessionConfig)) {
		query["StickySessionConfig"] = request.StickySessionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.UchConfig)) {
		query["UchConfig"] = request.UchConfig
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServerGroup"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * **CreateServerGroup** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListServerGroups](~~213627~~) operation to query the status of the task.
 * *   If a server group is in the **Creating** state, it indicates that the server group is being created.
 * *   If a server group is in the **Available** state, it indicates that the server group is created.
 *
 * @param request CreateServerGroupRequest
 * @return CreateServerGroupResponse
 */
func (client *Client) CreateServerGroup(request *CreateServerGroupRequest) (_result *CreateServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServerGroupResponse{}
	_body, _err := client.CreateServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAScriptsWithOptions(request *DeleteAScriptsRequest, runtime *util.RuntimeOptions) (_result *DeleteAScriptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AScriptIds)) {
		query["AScriptIds"] = request.AScriptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAScripts"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAScriptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAScripts(request *DeleteAScriptsRequest) (_result *DeleteAScriptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAScriptsResponse{}
	_body, _err := client.DeleteAScriptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the asynchronous task.
 *
 * @param request DeleteAclRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteAclResponse
 */
func (client *Client) DeleteAclWithOptions(request *DeleteAclRequest, runtime *util.RuntimeOptions) (_result *DeleteAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAcl"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ID of the asynchronous task.
 *
 * @param request DeleteAclRequest
 * @return DeleteAclResponse
 */
func (client *Client) DeleteAcl(request *DeleteAclRequest) (_result *DeleteAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAclResponse{}
	_body, _err := client.DeleteAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHealthCheckTemplatesWithOptions(request *DeleteHealthCheckTemplatesRequest, runtime *util.RuntimeOptions) (_result *DeleteHealthCheckTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTemplateIds)) {
		query["HealthCheckTemplateIds"] = request.HealthCheckTemplateIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHealthCheckTemplates"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHealthCheckTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHealthCheckTemplates(request *DeleteHealthCheckTemplatesRequest) (_result *DeleteHealthCheckTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHealthCheckTemplatesResponse{}
	_body, _err := client.DeleteHealthCheckTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The **DeleteRouteEntry** operation is asynchronous. After you send a request, the system returns the request ID. However, the operation is still being performed in the system background. You can call [GetListenerAttribute](~~2254865~~) to query the status of a listener.
 * *   If the listener is in the **Deleting** state, the vSwitch is being deleted.
 * *   If the listener cannot be found, the listener is deleted.
 *
 * @param request DeleteListenerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteListenerResponse
 */
func (client *Client) DeleteListenerWithOptions(request *DeleteListenerRequest, runtime *util.RuntimeOptions) (_result *DeleteListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteListener"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The **DeleteRouteEntry** operation is asynchronous. After you send a request, the system returns the request ID. However, the operation is still being performed in the system background. You can call [GetListenerAttribute](~~2254865~~) to query the status of a listener.
 * *   If the listener is in the **Deleting** state, the vSwitch is being deleted.
 * *   If the listener cannot be found, the listener is deleted.
 *
 * @param request DeleteListenerRequest
 * @return DeleteListenerResponse
 */
func (client *Client) DeleteListener(request *DeleteListenerRequest) (_result *DeleteListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteListenerResponse{}
	_body, _err := client.DeleteListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the request.
 *
 * @param request DeleteLoadBalancerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteLoadBalancerResponse
 */
func (client *Client) DeleteLoadBalancerWithOptions(request *DeleteLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *DeleteLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLoadBalancer"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ID of the request.
 *
 * @param request DeleteLoadBalancerRequest
 * @return DeleteLoadBalancerResponse
 */
func (client *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (_result *DeleteLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.DeleteLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the asynchronous task.
 *
 * @param request DeleteRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteRuleResponse
 */
func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRule"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ID of the asynchronous task.
 *
 * @param request DeleteRuleRequest
 * @return DeleteRuleResponse
 */
func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The **DeleteRules** operation is asynchronous. After you send a request, the system returns a request ID. However, the operation is still being performed in the system background. You can call [ListRules](~~214379~~) to query the status of forwarding rules.
 * *   If the forwarding rules are in the **Deleting** state, the forwarding rules are being deleted.
 * *   If the forwarding rules cannot be found, the forwarding rules are deleted.
 *
 * @param request DeleteRulesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteRulesResponse
 */
func (client *Client) DeleteRulesWithOptions(request *DeleteRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RuleIds)) {
		query["RuleIds"] = request.RuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRules"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The **DeleteRules** operation is asynchronous. After you send a request, the system returns a request ID. However, the operation is still being performed in the system background. You can call [ListRules](~~214379~~) to query the status of forwarding rules.
 * *   If the forwarding rules are in the **Deleting** state, the forwarding rules are being deleted.
 * *   If the forwarding rules cannot be found, the forwarding rules are deleted.
 *
 * @param request DeleteRulesRequest
 * @return DeleteRulesResponse
 */
func (client *Client) DeleteRules(request *DeleteRulesRequest) (_result *DeleteRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRulesResponse{}
	_body, _err := client.DeleteRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecurityPolicyWithOptions(request *DeleteSecurityPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecurityPolicy"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSecurityPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecurityPolicy(request *DeleteSecurityPolicyRequest) (_result *DeleteSecurityPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecurityPolicyResponse{}
	_body, _err := client.DeleteSecurityPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the asynchronous task.
 *
 * @param request DeleteServerGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteServerGroupResponse
 */
func (client *Client) DeleteServerGroupWithOptions(request *DeleteServerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServerGroup"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ID of the asynchronous task.
 *
 * @param request DeleteServerGroupRequest
 * @return DeleteServerGroupResponse
 */
func (client *Client) DeleteServerGroup(request *DeleteServerGroupRequest) (_result *DeleteServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServerGroupResponse{}
	_body, _err := client.DeleteServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Specifies whether to check the request without performing the operation. Valid values:
 * *   **true**: checks the request without perform the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
 * *   **false** (default): sends the request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed.
 *
 * @param request DetachCommonBandwidthPackageFromLoadBalancerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DetachCommonBandwidthPackageFromLoadBalancerResponse
 */
func (client *Client) DetachCommonBandwidthPackageFromLoadBalancerWithOptions(request *DetachCommonBandwidthPackageFromLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *DetachCommonBandwidthPackageFromLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachCommonBandwidthPackageFromLoadBalancer"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachCommonBandwidthPackageFromLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Specifies whether to check the request without performing the operation. Valid values:
 * *   **true**: checks the request without perform the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
 * *   **false** (default): sends the request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed.
 *
 * @param request DetachCommonBandwidthPackageFromLoadBalancerRequest
 * @return DetachCommonBandwidthPackageFromLoadBalancerResponse
 */
func (client *Client) DetachCommonBandwidthPackageFromLoadBalancer(request *DetachCommonBandwidthPackageFromLoadBalancerRequest) (_result *DetachCommonBandwidthPackageFromLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachCommonBandwidthPackageFromLoadBalancerResponse{}
	_body, _err := client.DetachCommonBandwidthPackageFromLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableDeletionProtectionWithOptions(request *DisableDeletionProtectionRequest, runtime *util.RuntimeOptions) (_result *DisableDeletionProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableDeletionProtection"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableDeletionProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableDeletionProtection(request *DisableDeletionProtectionRequest) (_result *DisableDeletionProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableDeletionProtectionResponse{}
	_body, _err := client.DisableDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableLoadBalancerAccessLogWithOptions(request *DisableLoadBalancerAccessLogRequest, runtime *util.RuntimeOptions) (_result *DisableLoadBalancerAccessLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableLoadBalancerAccessLog"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableLoadBalancerAccessLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableLoadBalancerAccessLog(request *DisableLoadBalancerAccessLogRequest) (_result *DisableLoadBalancerAccessLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableLoadBalancerAccessLogResponse{}
	_body, _err := client.DisableLoadBalancerAccessLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * - After the DisableLoadBalancerIpv6Internet operation is called, the value of the **Ipv6AddressType** parameter is changed to **Intranet** and the type of the IPv6 address of the ALB instance is changed from public to private. If you upgrade the instance or the instance scales elastic network interface (ENIs) along with workloads, private IPv6 addresses are automatically enabled for the instance and the new ENIs. You can call the [GetLoadBalancerAttribute](/help/en/server-load-balancer/latest/getloadbalancerattribute) operation to query the value of the **Ipv6AddressType** parameter.
 * - **DisableLoadBalancerIpv6Internet** is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](/help/en/server-load-balancer/latest/getloadbalancerattribute) operation to query the network type of the IPv6 address of an ALB instance.
 *       - If the ALB instance is in the **Configuring** state, the network type of the IPv6 address that is used by the ALB instance is being changed.
 *   - If the ALB instance is in the **Active** state, the network type of the IPv6 address that is used by the ALB instance has been changed successfully.
 * ## Prerequisites
 * An ALB instance is created and IPv4/IPv6 dual stack is enabled for the instance. You can call the [CreateLoadBalancer](~~214358~~) operation and set the **AddressIpVersion** parameter to **DualStack** to create a dual-stack ALB instance.
 * When **AddressIpVersion** is set to **DualStack**:
 *    *   If you set the **AddressType** parameter to **Internet**, the ALB instance uses a public IPv4 IP address and a private IPv6 address.
 *    *   If you set the **AddressType** parameter to **Intranet**, the ALB instance uses a private IPv4 IP address and a private IPv6 address.
 *
 * @param request DisableLoadBalancerIpv6InternetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DisableLoadBalancerIpv6InternetResponse
 */
func (client *Client) DisableLoadBalancerIpv6InternetWithOptions(request *DisableLoadBalancerIpv6InternetRequest, runtime *util.RuntimeOptions) (_result *DisableLoadBalancerIpv6InternetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableLoadBalancerIpv6Internet"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableLoadBalancerIpv6InternetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * - After the DisableLoadBalancerIpv6Internet operation is called, the value of the **Ipv6AddressType** parameter is changed to **Intranet** and the type of the IPv6 address of the ALB instance is changed from public to private. If you upgrade the instance or the instance scales elastic network interface (ENIs) along with workloads, private IPv6 addresses are automatically enabled for the instance and the new ENIs. You can call the [GetLoadBalancerAttribute](/help/en/server-load-balancer/latest/getloadbalancerattribute) operation to query the value of the **Ipv6AddressType** parameter.
 * - **DisableLoadBalancerIpv6Internet** is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](/help/en/server-load-balancer/latest/getloadbalancerattribute) operation to query the network type of the IPv6 address of an ALB instance.
 *       - If the ALB instance is in the **Configuring** state, the network type of the IPv6 address that is used by the ALB instance is being changed.
 *   - If the ALB instance is in the **Active** state, the network type of the IPv6 address that is used by the ALB instance has been changed successfully.
 * ## Prerequisites
 * An ALB instance is created and IPv4/IPv6 dual stack is enabled for the instance. You can call the [CreateLoadBalancer](~~214358~~) operation and set the **AddressIpVersion** parameter to **DualStack** to create a dual-stack ALB instance.
 * When **AddressIpVersion** is set to **DualStack**:
 *    *   If you set the **AddressType** parameter to **Internet**, the ALB instance uses a public IPv4 IP address and a private IPv6 address.
 *    *   If you set the **AddressType** parameter to **Intranet**, the ALB instance uses a private IPv4 IP address and a private IPv6 address.
 *
 * @param request DisableLoadBalancerIpv6InternetRequest
 * @return DisableLoadBalancerIpv6InternetResponse
 */
func (client *Client) DisableLoadBalancerIpv6Internet(request *DisableLoadBalancerIpv6InternetRequest) (_result *DisableLoadBalancerIpv6InternetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableLoadBalancerIpv6InternetResponse{}
	_body, _err := client.DisableLoadBalancerIpv6InternetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The operation that you want to perform. Set the value to **DissociateAclsFromListener**.
 *
 * @param request DissociateAclsFromListenerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DissociateAclsFromListenerResponse
 */
func (client *Client) DissociateAclsFromListenerWithOptions(request *DissociateAclsFromListenerRequest, runtime *util.RuntimeOptions) (_result *DissociateAclsFromListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclIds)) {
		query["AclIds"] = request.AclIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DissociateAclsFromListener"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DissociateAclsFromListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The operation that you want to perform. Set the value to **DissociateAclsFromListener**.
 *
 * @param request DissociateAclsFromListenerRequest
 * @return DissociateAclsFromListenerResponse
 */
func (client *Client) DissociateAclsFromListener(request *DissociateAclsFromListenerRequest) (_result *DissociateAclsFromListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateAclsFromListenerResponse{}
	_body, _err := client.DissociateAclsFromListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the certificate. Only server certificates are supported. A maximum of 20 certificate IDs are supported.
 *
 * @param request DissociateAdditionalCertificatesFromListenerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DissociateAdditionalCertificatesFromListenerResponse
 */
func (client *Client) DissociateAdditionalCertificatesFromListenerWithOptions(request *DissociateAdditionalCertificatesFromListenerRequest, runtime *util.RuntimeOptions) (_result *DissociateAdditionalCertificatesFromListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Certificates)) {
		query["Certificates"] = request.Certificates
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DissociateAdditionalCertificatesFromListener"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DissociateAdditionalCertificatesFromListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ID of the certificate. Only server certificates are supported. A maximum of 20 certificate IDs are supported.
 *
 * @param request DissociateAdditionalCertificatesFromListenerRequest
 * @return DissociateAdditionalCertificatesFromListenerResponse
 */
func (client *Client) DissociateAdditionalCertificatesFromListener(request *DissociateAdditionalCertificatesFromListenerRequest) (_result *DissociateAdditionalCertificatesFromListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateAdditionalCertificatesFromListenerResponse{}
	_body, _err := client.DissociateAdditionalCertificatesFromListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableDeletionProtectionWithOptions(request *EnableDeletionProtectionRequest, runtime *util.RuntimeOptions) (_result *EnableDeletionProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableDeletionProtection"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableDeletionProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableDeletionProtection(request *EnableDeletionProtectionRequest) (_result *EnableDeletionProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableDeletionProtectionResponse{}
	_body, _err := client.EnableDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableLoadBalancerAccessLogWithOptions(request *EnableLoadBalancerAccessLogRequest, runtime *util.RuntimeOptions) (_result *EnableLoadBalancerAccessLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.LogProject)) {
		query["LogProject"] = request.LogProject
	}

	if !tea.BoolValue(util.IsUnset(request.LogStore)) {
		query["LogStore"] = request.LogStore
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableLoadBalancerAccessLog"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableLoadBalancerAccessLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableLoadBalancerAccessLog(request *EnableLoadBalancerAccessLogRequest) (_result *EnableLoadBalancerAccessLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableLoadBalancerAccessLogResponse{}
	_body, _err := client.EnableLoadBalancerAccessLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * - After the EnableLoadBalancerIpv6Internet operation is called, the value of the **Ipv6AddressType** parameter is changed to **Internet** and the type of the IPv6 address of the ALB instance is changed from private to public. If you upgrade the instance or the instance scales elastic network interface (ENIs) along with workloads, public IPv6 addresses are automatically enabled for the instance and the new ENIs. You can call the [GetLoadBalancerAttribute](/help/en/server-load-balancer/latest/getloadbalancerattribute) operation to query the value of the **Ipv6AddressType** parameter.
 * - **EnableLoadBalancerIpv6Internet** is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](/help/en/server-load-balancer/latest/getloadbalancerattribute) operation to query the network type of the IPv6 address of an ALB instance.
 *       - If the ALB instance is in the **Configuring** state, the network type of the IPv6 address that is used by the ALB instance is being changed.
 *   - If the ALB instance is in the **Active** state, the network type of the IPv6 address that is used by the ALB instance has been changed successfully.
 * ## Prerequisites
 * An ALB instance is created and IPv4/IPv6 dual stack is enabled for the instance. You can call the [CreateLoadBalancer](~~214358~~) operation and set the **AddressIpVersion** parameter to **DualStack** to create a dual-stack ALB instance.
 * When **AddressIpVersion** is set to **DualStack**:
 *   *   If you set the **AddressType** parameter to **Internet**, the ALB instance uses a public IPv4 IP address and a private IPv6 address.
 *   *   If you set the **AddressType** parameter to **Intranet**, the ALB instance uses a private IPv4 IP address and a private IPv6 address.
 *
 * @param request EnableLoadBalancerIpv6InternetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return EnableLoadBalancerIpv6InternetResponse
 */
func (client *Client) EnableLoadBalancerIpv6InternetWithOptions(request *EnableLoadBalancerIpv6InternetRequest, runtime *util.RuntimeOptions) (_result *EnableLoadBalancerIpv6InternetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableLoadBalancerIpv6Internet"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableLoadBalancerIpv6InternetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * - After the EnableLoadBalancerIpv6Internet operation is called, the value of the **Ipv6AddressType** parameter is changed to **Internet** and the type of the IPv6 address of the ALB instance is changed from private to public. If you upgrade the instance or the instance scales elastic network interface (ENIs) along with workloads, public IPv6 addresses are automatically enabled for the instance and the new ENIs. You can call the [GetLoadBalancerAttribute](/help/en/server-load-balancer/latest/getloadbalancerattribute) operation to query the value of the **Ipv6AddressType** parameter.
 * - **EnableLoadBalancerIpv6Internet** is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetLoadBalancerAttribute](/help/en/server-load-balancer/latest/getloadbalancerattribute) operation to query the network type of the IPv6 address of an ALB instance.
 *       - If the ALB instance is in the **Configuring** state, the network type of the IPv6 address that is used by the ALB instance is being changed.
 *   - If the ALB instance is in the **Active** state, the network type of the IPv6 address that is used by the ALB instance has been changed successfully.
 * ## Prerequisites
 * An ALB instance is created and IPv4/IPv6 dual stack is enabled for the instance. You can call the [CreateLoadBalancer](~~214358~~) operation and set the **AddressIpVersion** parameter to **DualStack** to create a dual-stack ALB instance.
 * When **AddressIpVersion** is set to **DualStack**:
 *   *   If you set the **AddressType** parameter to **Internet**, the ALB instance uses a public IPv4 IP address and a private IPv6 address.
 *   *   If you set the **AddressType** parameter to **Intranet**, the ALB instance uses a private IPv4 IP address and a private IPv6 address.
 *
 * @param request EnableLoadBalancerIpv6InternetRequest
 * @return EnableLoadBalancerIpv6InternetResponse
 */
func (client *Client) EnableLoadBalancerIpv6Internet(request *EnableLoadBalancerIpv6InternetRequest) (_result *EnableLoadBalancerIpv6InternetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableLoadBalancerIpv6InternetResponse{}
	_body, _err := client.EnableLoadBalancerIpv6InternetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHealthCheckTemplateAttributeWithOptions(request *GetHealthCheckTemplateAttributeRequest, runtime *util.RuntimeOptions) (_result *GetHealthCheckTemplateAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HealthCheckTemplateId)) {
		query["HealthCheckTemplateId"] = request.HealthCheckTemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHealthCheckTemplateAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHealthCheckTemplateAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHealthCheckTemplateAttribute(request *GetHealthCheckTemplateAttributeRequest) (_result *GetHealthCheckTemplateAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHealthCheckTemplateAttributeResponse{}
	_body, _err := client.GetHealthCheckTemplateAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetListenerAttributeWithOptions(request *GetListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *GetListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetListenerAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetListenerAttribute(request *GetListenerAttributeRequest) (_result *GetListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetListenerAttributeResponse{}
	_body, _err := client.GetListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetListenerHealthStatusWithOptions(request *GetListenerHealthStatusRequest, runtime *util.RuntimeOptions) (_result *GetListenerHealthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeRule)) {
		query["IncludeRule"] = request.IncludeRule
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetListenerHealthStatus"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetListenerHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetListenerHealthStatus(request *GetListenerHealthStatusRequest) (_result *GetListenerHealthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetListenerHealthStatusResponse{}
	_body, _err := client.GetListenerHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLoadBalancerAttributeWithOptions(request *GetLoadBalancerAttributeRequest, runtime *util.RuntimeOptions) (_result *GetLoadBalancerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLoadBalancerAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoadBalancerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLoadBalancerAttribute(request *GetLoadBalancerAttributeRequest) (_result *GetLoadBalancerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoadBalancerAttributeResponse{}
	_body, _err := client.GetLoadBalancerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAScriptsWithOptions(request *ListAScriptsRequest, runtime *util.RuntimeOptions) (_result *ListAScriptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AScriptIds)) {
		query["AScriptIds"] = request.AScriptIds
	}

	if !tea.BoolValue(util.IsUnset(request.AScriptNames)) {
		query["AScriptNames"] = request.AScriptNames
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerIds)) {
		query["ListenerIds"] = request.ListenerIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAScripts"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAScriptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAScripts(request *ListAScriptsRequest) (_result *ListAScriptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAScriptsResponse{}
	_body, _err := client.ListAScriptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAclEntriesWithOptions(request *ListAclEntriesRequest, runtime *util.RuntimeOptions) (_result *ListAclEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAclEntries"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAclEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAclEntries(request *ListAclEntriesRequest) (_result *ListAclEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAclEntriesResponse{}
	_body, _err := client.ListAclEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAclRelationsWithOptions(request *ListAclRelationsRequest, runtime *util.RuntimeOptions) (_result *ListAclRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclIds)) {
		query["AclIds"] = request.AclIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAclRelations"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAclRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAclRelations(request *ListAclRelationsRequest) (_result *ListAclRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAclRelationsResponse{}
	_body, _err := client.ListAclRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The maximum number of network ACLs returned. This parameter is optional. Valid values: **1** to **100**. If this parameter is not set, the default value **20** is returned.
 *
 * @param request ListAclsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAclsResponse
 */
func (client *Client) ListAclsWithOptions(request *ListAclsRequest, runtime *util.RuntimeOptions) (_result *ListAclsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclIds)) {
		query["AclIds"] = request.AclIds
	}

	if !tea.BoolValue(util.IsUnset(request.AclNames)) {
		query["AclNames"] = request.AclNames
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAcls"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAclsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The maximum number of network ACLs returned. This parameter is optional. Valid values: **1** to **100**. If this parameter is not set, the default value **20** is returned.
 *
 * @param request ListAclsRequest
 * @return ListAclsResponse
 */
func (client *Client) ListAcls(request *ListAclsRequest) (_result *ListAclsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAclsResponse{}
	_body, _err := client.ListAclsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAsynJobsWithOptions(request *ListAsynJobsRequest, runtime *util.RuntimeOptions) (_result *ListAsynJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.JobIds)) {
		query["JobIds"] = request.JobIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAsynJobs"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAsynJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAsynJobs(request *ListAsynJobsRequest) (_result *ListAsynJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAsynJobsResponse{}
	_body, _err := client.ListAsynJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHealthCheckTemplatesWithOptions(request *ListHealthCheckTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListHealthCheckTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HealthCheckTemplateIds)) {
		query["HealthCheckTemplateIds"] = request.HealthCheckTemplateIds
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTemplateNames)) {
		query["HealthCheckTemplateNames"] = request.HealthCheckTemplateNames
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHealthCheckTemplates"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHealthCheckTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHealthCheckTemplates(request *ListHealthCheckTemplatesRequest) (_result *ListHealthCheckTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHealthCheckTemplatesResponse{}
	_body, _err := client.ListHealthCheckTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListListenerCertificatesWithOptions(request *ListListenerCertificatesRequest, runtime *util.RuntimeOptions) (_result *ListListenerCertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateType)) {
		query["CertificateType"] = request.CertificateType
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListenerCertificates"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListListenerCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListListenerCertificates(request *ListListenerCertificatesRequest) (_result *ListListenerCertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListenerCertificatesResponse{}
	_body, _err := client.ListListenerCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListListenersWithOptions(request *ListListenersRequest, runtime *util.RuntimeOptions) (_result *ListListenersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerIds)) {
		query["ListenerIds"] = request.ListenerIds
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerIds)) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListeners"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListListenersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListListeners(request *ListListenersRequest) (_result *ListListenersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListenersResponse{}
	_body, _err := client.ListListenersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLoadBalancersWithOptions(request *ListLoadBalancersRequest, runtime *util.RuntimeOptions) (_result *ListLoadBalancersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerBussinessStatus)) {
		query["LoadBalancerBussinessStatus"] = request.LoadBalancerBussinessStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerIds)) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerNames)) {
		query["LoadBalancerNames"] = request.LoadBalancerNames
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerStatus)) {
		query["LoadBalancerStatus"] = request.LoadBalancerStatus
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpcIds)) {
		query["VpcIds"] = request.VpcIds
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLoadBalancers"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLoadBalancers(request *ListLoadBalancersRequest) (_result *ListLoadBalancersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.ListLoadBalancersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *util.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerIds)) {
		query["ListenerIds"] = request.ListenerIds
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerIds)) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RuleIds)) {
		query["RuleIds"] = request.RuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRules"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecurityPoliciesWithOptions(request *ListSecurityPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListSecurityPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyIds)) {
		query["SecurityPolicyIds"] = request.SecurityPolicyIds
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyNames)) {
		query["SecurityPolicyNames"] = request.SecurityPolicyNames
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecurityPolicies"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSecurityPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecurityPolicies(request *ListSecurityPoliciesRequest) (_result *ListSecurityPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecurityPoliciesResponse{}
	_body, _err := client.ListSecurityPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecurityPolicyRelationsWithOptions(request *ListSecurityPolicyRelationsRequest, runtime *util.RuntimeOptions) (_result *ListSecurityPolicyRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyIds)) {
		query["SecurityPolicyIds"] = request.SecurityPolicyIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecurityPolicyRelations"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSecurityPolicyRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecurityPolicyRelations(request *ListSecurityPolicyRelationsRequest) (_result *ListSecurityPolicyRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecurityPolicyRelationsResponse{}
	_body, _err := client.ListSecurityPolicyRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServerGroupServersWithOptions(request *ListServerGroupServersRequest, runtime *util.RuntimeOptions) (_result *ListServerGroupServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIds)) {
		query["ServerIds"] = request.ServerIds
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServerGroupServers"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServerGroupServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServerGroupServers(request *ListServerGroupServersRequest) (_result *ListServerGroupServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServerGroupServersResponse{}
	_body, _err := client.ListServerGroupServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServerGroupsWithOptions(request *ListServerGroupsRequest, runtime *util.RuntimeOptions) (_result *ListServerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupIds)) {
		query["ServerGroupIds"] = request.ServerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupNames)) {
		query["ServerGroupNames"] = request.ServerGroupNames
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServerGroups"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServerGroups(request *ListServerGroupsRequest) (_result *ListServerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServerGroupsResponse{}
	_body, _err := client.ListServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSystemSecurityPoliciesWithOptions(runtime *util.RuntimeOptions) (_result *ListSystemSecurityPoliciesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListSystemSecurityPolicies"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSystemSecurityPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSystemSecurityPolicies() (_result *ListSystemSecurityPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSystemSecurityPoliciesResponse{}
	_body, _err := client.ListSystemSecurityPoliciesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The **RemoveEntriesFromAcl** operation is asynchronous. After you send a request, the system returns the request ID, but the operation is still being performed in the system background. You can call the [ListAclEntries](~~213616~~) operation to query the status of an entry in an ACL:
 * *   If an ACL is in the **Removing** state, the entry is being deleted.
 * *   If an ACL cannot be found, the entry is deleted.
 *
 * @param request RemoveEntriesFromAclRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveEntriesFromAclResponse
 */
func (client *Client) RemoveEntriesFromAclWithOptions(request *RemoveEntriesFromAclRequest, runtime *util.RuntimeOptions) (_result *RemoveEntriesFromAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Entries)) {
		query["Entries"] = request.Entries
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveEntriesFromAcl"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveEntriesFromAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The **RemoveEntriesFromAcl** operation is asynchronous. After you send a request, the system returns the request ID, but the operation is still being performed in the system background. You can call the [ListAclEntries](~~213616~~) operation to query the status of an entry in an ACL:
 * *   If an ACL is in the **Removing** state, the entry is being deleted.
 * *   If an ACL cannot be found, the entry is deleted.
 *
 * @param request RemoveEntriesFromAclRequest
 * @return RemoveEntriesFromAclResponse
 */
func (client *Client) RemoveEntriesFromAcl(request *RemoveEntriesFromAclRequest) (_result *RemoveEntriesFromAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveEntriesFromAclResponse{}
	_body, _err := client.RemoveEntriesFromAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The IP address in inclusive ENI mode. You can specify at most 40 servers in each call.
 *
 * @param request RemoveServersFromServerGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveServersFromServerGroupResponse
 */
func (client *Client) RemoveServersFromServerGroupWithOptions(request *RemoveServersFromServerGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveServersFromServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Servers)) {
		query["Servers"] = request.Servers
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveServersFromServerGroup"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveServersFromServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The IP address in inclusive ENI mode. You can specify at most 40 servers in each call.
 *
 * @param request RemoveServersFromServerGroupRequest
 * @return RemoveServersFromServerGroupResponse
 */
func (client *Client) RemoveServersFromServerGroup(request *RemoveServersFromServerGroupRequest) (_result *RemoveServersFromServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveServersFromServerGroupResponse{}
	_body, _err := client.RemoveServersFromServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * **ReplaceServersInServerGroup** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
 * 1.  You can call the [ListServerGroups](~~213627~~) operation to query the status of a server group.
 *     *   If a server group is in the **Configuring** state, it indicates that the server group is being modified.
 *     *   If a server group is in the **Available** state, it indicates that the server group is running.
 * 2.  You can call the [ListServerGroupServers](~~213628~~) operation to query the status of a backend server.
 *     *   If a backend server is in the **Replacing** state, it indicates that the server is being removed from the server group and a new server is added to the server group.
 *     *   If a backend server is in the \\*\\*Available\\*\\* state, it indicates that the server is running.
 *
 * @param request ReplaceServersInServerGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ReplaceServersInServerGroupResponse
 */
func (client *Client) ReplaceServersInServerGroupWithOptions(request *ReplaceServersInServerGroupRequest, runtime *util.RuntimeOptions) (_result *ReplaceServersInServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddedServers)) {
		query["AddedServers"] = request.AddedServers
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RemovedServers)) {
		query["RemovedServers"] = request.RemovedServers
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReplaceServersInServerGroup"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplaceServersInServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * **ReplaceServersInServerGroup** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
 * 1.  You can call the [ListServerGroups](~~213627~~) operation to query the status of a server group.
 *     *   If a server group is in the **Configuring** state, it indicates that the server group is being modified.
 *     *   If a server group is in the **Available** state, it indicates that the server group is running.
 * 2.  You can call the [ListServerGroupServers](~~213628~~) operation to query the status of a backend server.
 *     *   If a backend server is in the **Replacing** state, it indicates that the server is being removed from the server group and a new server is added to the server group.
 *     *   If a backend server is in the \\*\\*Available\\*\\* state, it indicates that the server is running.
 *
 * @param request ReplaceServersInServerGroupRequest
 * @return ReplaceServersInServerGroupResponse
 */
func (client *Client) ReplaceServersInServerGroup(request *ReplaceServersInServerGroupRequest) (_result *ReplaceServersInServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplaceServersInServerGroupResponse{}
	_body, _err := client.ReplaceServersInServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the asynchronous task.
 *
 * @param request StartListenerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartListenerResponse
 */
func (client *Client) StartListenerWithOptions(request *StartListenerRequest, runtime *util.RuntimeOptions) (_result *StartListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartListener"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ID of the asynchronous task.
 *
 * @param request StartListenerRequest
 * @return StartListenerResponse
 */
func (client *Client) StartListener(request *StartListenerRequest) (_result *StartListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartListenerResponse{}
	_body, _err := client.StartListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the asynchronous task.
 *
 * @param request StopListenerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopListenerResponse
 */
func (client *Client) StopListenerWithOptions(request *StopListenerRequest, runtime *util.RuntimeOptions) (_result *StopListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopListener"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ID of the asynchronous task.
 *
 * @param request StopListenerRequest
 * @return StopListenerResponse
 */
func (client *Client) StopListener(request *StopListenerRequest) (_result *StopListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopListenerResponse{}
	_body, _err := client.StopListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnTagResources"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   **UpdateAScripts** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAScripts](~~472574~~) operation to query the status of the task.
 *     *   If an AScript rule is in the **Configuring** state, the AScript rule is being updated.
 *     *   If an AScript rule is in the **Available** state, the AScript rule is updated.
 * *   In the following table, the maximum value of **N** is **4**.
 *
 * @param request UpdateAScriptsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateAScriptsResponse
 */
func (client *Client) UpdateAScriptsWithOptions(request *UpdateAScriptsRequest, runtime *util.RuntimeOptions) (_result *UpdateAScriptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AScripts)) {
		query["AScripts"] = request.AScripts
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAScripts"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAScriptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   **UpdateAScripts** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListAScripts](~~472574~~) operation to query the status of the task.
 *     *   If an AScript rule is in the **Configuring** state, the AScript rule is being updated.
 *     *   If an AScript rule is in the **Available** state, the AScript rule is updated.
 * *   In the following table, the maximum value of **N** is **4**.
 *
 * @param request UpdateAScriptsRequest
 * @return UpdateAScriptsResponse
 */
func (client *Client) UpdateAScripts(request *UpdateAScriptsRequest) (_result *UpdateAScriptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAScriptsResponse{}
	_body, _err := client.UpdateAScriptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAclAttributeWithOptions(request *UpdateAclAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateAclAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.AclName)) {
		query["AclName"] = request.AclName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAclAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAclAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAclAttribute(request *UpdateAclAttributeRequest) (_result *UpdateAclAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAclAttributeResponse{}
	_body, _err := client.UpdateAclAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateHealthCheckTemplateAttributeWithOptions(request *UpdateHealthCheckTemplateAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateHealthCheckTemplateAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckCodes)) {
		query["HealthCheckCodes"] = request.HealthCheckCodes
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectPort)) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckHost)) {
		query["HealthCheckHost"] = request.HealthCheckHost
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckHttpVersion)) {
		query["HealthCheckHttpVersion"] = request.HealthCheckHttpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckInterval)) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckMethod)) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckPath)) {
		query["HealthCheckPath"] = request.HealthCheckPath
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckProtocol)) {
		query["HealthCheckProtocol"] = request.HealthCheckProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTemplateId)) {
		query["HealthCheckTemplateId"] = request.HealthCheckTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTemplateName)) {
		query["HealthCheckTemplateName"] = request.HealthCheckTemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTimeout)) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthyThreshold)) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.UnhealthyThreshold)) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHealthCheckTemplateAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHealthCheckTemplateAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateHealthCheckTemplateAttribute(request *UpdateHealthCheckTemplateAttributeRequest) (_result *UpdateHealthCheckTemplateAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateHealthCheckTemplateAttributeResponse{}
	_body, _err := client.UpdateHealthCheckTemplateAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * **UpdateListenerAttribute** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetListenerAttribute](~~214353~~) operation to query the status of the task.
 * *   If a listener is in the **Configuring** state, the listener is being updated.
 * *   If a listener is in the **Running** state, the listener is updated.
 *
 * @param request UpdateListenerAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateListenerAttributeResponse
 */
func (client *Client) UpdateListenerAttributeWithOptions(request *UpdateListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaCertificates)) {
		query["CaCertificates"] = request.CaCertificates
	}

	if !tea.BoolValue(util.IsUnset(request.CaEnabled)) {
		query["CaEnabled"] = request.CaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Certificates)) {
		query["Certificates"] = request.Certificates
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultActions)) {
		query["DefaultActions"] = request.DefaultActions
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.GzipEnabled)) {
		query["GzipEnabled"] = request.GzipEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Http2Enabled)) {
		query["Http2Enabled"] = request.Http2Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerDescription)) {
		query["ListenerDescription"] = request.ListenerDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.QuicConfig)) {
		query["QuicConfig"] = request.QuicConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RequestTimeout)) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedForConfig)) {
		query["XForwardedForConfig"] = request.XForwardedForConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateListenerAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * **UpdateListenerAttribute** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetListenerAttribute](~~214353~~) operation to query the status of the task.
 * *   If a listener is in the **Configuring** state, the listener is being updated.
 * *   If a listener is in the **Running** state, the listener is updated.
 *
 * @param request UpdateListenerAttributeRequest
 * @return UpdateListenerAttributeResponse
 */
func (client *Client) UpdateListenerAttribute(request *UpdateListenerAttributeRequest) (_result *UpdateListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateListenerAttributeResponse{}
	_body, _err := client.UpdateListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The **UpdateListenerLogConfig** operation is asynchronous. After you send a request, the system returns the request ID, but the operation is still being performed in the system background. You can call [GetListenerAttribute](~~2254865~~) to query the status of a listener:
 * *   If a listener is in the **Configuring** state, the log configuration of the listener is being modified.
 * *   If a listener is in the **Running** state, the log configuration of the listener is modified.
 * > You can call this operation only if you enable the access log feature for the ALB instance that you want to manage.
 *
 * @param request UpdateListenerLogConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateListenerLogConfigResponse
 */
func (client *Client) UpdateListenerLogConfigWithOptions(request *UpdateListenerLogConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateListenerLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessLogRecordCustomizedHeadersEnabled)) {
		query["AccessLogRecordCustomizedHeadersEnabled"] = request.AccessLogRecordCustomizedHeadersEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogTracingConfig)) {
		query["AccessLogTracingConfig"] = request.AccessLogTracingConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateListenerLogConfig"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateListenerLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The **UpdateListenerLogConfig** operation is asynchronous. After you send a request, the system returns the request ID, but the operation is still being performed in the system background. You can call [GetListenerAttribute](~~2254865~~) to query the status of a listener:
 * *   If a listener is in the **Configuring** state, the log configuration of the listener is being modified.
 * *   If a listener is in the **Running** state, the log configuration of the listener is modified.
 * > You can call this operation only if you enable the access log feature for the ALB instance that you want to manage.
 *
 * @param request UpdateListenerLogConfigRequest
 * @return UpdateListenerLogConfigResponse
 */
func (client *Client) UpdateListenerLogConfig(request *UpdateListenerLogConfigRequest) (_result *UpdateListenerLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateListenerLogConfigResponse{}
	_body, _err := client.UpdateListenerLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The client token that is used to ensure the idempotence of the request.
 * You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
 * >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** may be different for each API request.
 *
 * @param request UpdateLoadBalancerAddressTypeConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateLoadBalancerAddressTypeConfigResponse
 */
func (client *Client) UpdateLoadBalancerAddressTypeConfigWithOptions(request *UpdateLoadBalancerAddressTypeConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerAddressTypeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneMappings)) {
		query["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerAddressTypeConfig"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerAddressTypeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The client token that is used to ensure the idempotence of the request.
 * You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
 * >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** may be different for each API request.
 *
 * @param request UpdateLoadBalancerAddressTypeConfigRequest
 * @return UpdateLoadBalancerAddressTypeConfigResponse
 */
func (client *Client) UpdateLoadBalancerAddressTypeConfig(request *UpdateLoadBalancerAddressTypeConfigRequest) (_result *UpdateLoadBalancerAddressTypeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerAddressTypeConfigResponse{}
	_body, _err := client.UpdateLoadBalancerAddressTypeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The name of the ALB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\\_), and hyphens (-). The name must start with a letter.
 *
 * @param request UpdateLoadBalancerAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateLoadBalancerAttributeResponse
 */
func (client *Client) UpdateLoadBalancerAttributeWithOptions(request *UpdateLoadBalancerAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerName)) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	if !tea.BoolValue(util.IsUnset(request.ModificationProtectionConfig)) {
		query["ModificationProtectionConfig"] = request.ModificationProtectionConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The name of the ALB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\\_), and hyphens (-). The name must start with a letter.
 *
 * @param request UpdateLoadBalancerAttributeRequest
 * @return UpdateLoadBalancerAttributeResponse
 */
func (client *Client) UpdateLoadBalancerAttribute(request *UpdateLoadBalancerAttributeRequest) (_result *UpdateLoadBalancerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerAttributeResponse{}
	_body, _err := client.UpdateLoadBalancerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The client token that is used to ensure the idempotence of the request.
 * You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
 * >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** may be different for each API request.
 *
 * @param request UpdateLoadBalancerEditionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateLoadBalancerEditionResponse
 */
func (client *Client) UpdateLoadBalancerEditionWithOptions(request *UpdateLoadBalancerEditionRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerEditionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerEdition)) {
		query["LoadBalancerEdition"] = request.LoadBalancerEdition
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerEdition"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerEditionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The client token that is used to ensure the idempotence of the request.
 * You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
 * >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** may be different for each API request.
 *
 * @param request UpdateLoadBalancerEditionRequest
 * @return UpdateLoadBalancerEditionResponse
 */
func (client *Client) UpdateLoadBalancerEdition(request *UpdateLoadBalancerEditionRequest) (_result *UpdateLoadBalancerEditionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerEditionResponse{}
	_body, _err := client.UpdateLoadBalancerEditionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * **UpdateLoadBalancerZones** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [GetLoadBalancerAttribute](~~214362~~) to query the status of the task.
 * *   If an ALB instance is in the **Configuring** state, the zones are being modified.
 * *   If an ALB instance is in the **Active** state, the zones are modified.
 * > You may be charged after you call UpdateLoadBalancerZones.
 *
 * @param request UpdateLoadBalancerZonesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateLoadBalancerZonesResponse
 */
func (client *Client) UpdateLoadBalancerZonesWithOptions(request *UpdateLoadBalancerZonesRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneMappings)) {
		query["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerZones"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * **UpdateLoadBalancerZones** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [GetLoadBalancerAttribute](~~214362~~) to query the status of the task.
 * *   If an ALB instance is in the **Configuring** state, the zones are being modified.
 * *   If an ALB instance is in the **Active** state, the zones are modified.
 * > You may be charged after you call UpdateLoadBalancerZones.
 *
 * @param request UpdateLoadBalancerZonesRequest
 * @return UpdateLoadBalancerZonesResponse
 */
func (client *Client) UpdateLoadBalancerZones(request *UpdateLoadBalancerZonesRequest) (_result *UpdateLoadBalancerZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerZonesResponse{}
	_body, _err := client.UpdateLoadBalancerZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Specifies whether to check the request without performing the operation. Valid values:
 * *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error code is returned based on the cause of the failure. If the request passes the check, the `DryRunOperation` error code is returned.
 * *   **false**: checks the request. If the request passes the check, the system returns an `HTTP 2xx` status code and performs the operation. This is the default value.
 *
 * @param request UpdateRuleAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateRuleAttributeResponse
 */
func (client *Client) UpdateRuleAttributeWithOptions(request *UpdateRuleAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateRuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RuleActions)) {
		query["RuleActions"] = request.RuleActions
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditions)) {
		query["RuleConditions"] = request.RuleConditions
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRuleAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRuleAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Specifies whether to check the request without performing the operation. Valid values:
 * *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error code is returned based on the cause of the failure. If the request passes the check, the `DryRunOperation` error code is returned.
 * *   **false**: checks the request. If the request passes the check, the system returns an `HTTP 2xx` status code and performs the operation. This is the default value.
 *
 * @param request UpdateRuleAttributeRequest
 * @return UpdateRuleAttributeResponse
 */
func (client *Client) UpdateRuleAttribute(request *UpdateRuleAttributeRequest) (_result *UpdateRuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRuleAttributeResponse{}
	_body, _err := client.UpdateRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The priority of the forwarding rule. Valid values: **1 to 10000**. A lower value specifies a higher priority. You can specify at most 20 forwarding rules.
 * >  The priority of each forwarding rule within a listener is unique.
 *
 * @param request UpdateRulesAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateRulesAttributeResponse
 */
func (client *Client) UpdateRulesAttributeWithOptions(request *UpdateRulesAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateRulesAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Rules)) {
		query["Rules"] = request.Rules
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRulesAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRulesAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The priority of the forwarding rule. Valid values: **1 to 10000**. A lower value specifies a higher priority. You can specify at most 20 forwarding rules.
 * >  The priority of each forwarding rule within a listener is unique.
 *
 * @param request UpdateRulesAttributeRequest
 * @return UpdateRulesAttributeResponse
 */
func (client *Client) UpdateRulesAttribute(request *UpdateRulesAttributeRequest) (_result *UpdateRulesAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRulesAttributeResponse{}
	_body, _err := client.UpdateRulesAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ##
 * **UpdateSecurityPolicyAttribute** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [ListSecurityPolicies](~~213609~~) to query the status of the task.
 * *   If a security policy is in the **Configuring** state, the security policy is being updated.
 * *   If a security policy is in the **Available** state, the security policy is updated.
 *
 * @param request UpdateSecurityPolicyAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateSecurityPolicyAttributeResponse
 */
func (client *Client) UpdateSecurityPolicyAttributeWithOptions(request *UpdateSecurityPolicyAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateSecurityPolicyAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ciphers)) {
		query["Ciphers"] = request.Ciphers
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyName)) {
		query["SecurityPolicyName"] = request.SecurityPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.TLSVersions)) {
		query["TLSVersions"] = request.TLSVersions
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSecurityPolicyAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSecurityPolicyAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ##
 * **UpdateSecurityPolicyAttribute** is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call [ListSecurityPolicies](~~213609~~) to query the status of the task.
 * *   If a security policy is in the **Configuring** state, the security policy is being updated.
 * *   If a security policy is in the **Available** state, the security policy is updated.
 *
 * @param request UpdateSecurityPolicyAttributeRequest
 * @return UpdateSecurityPolicyAttributeResponse
 */
func (client *Client) UpdateSecurityPolicyAttribute(request *UpdateSecurityPolicyAttributeRequest) (_result *UpdateSecurityPolicyAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSecurityPolicyAttributeResponse{}
	_body, _err := client.UpdateSecurityPolicyAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The name of the server group.
 * The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (\\_), and hyphens (-). The name must start with a letter.
 *
 * @param request UpdateServerGroupAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateServerGroupAttributeResponse
 */
func (client *Client) UpdateServerGroupAttributeWithOptions(request *UpdateServerGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServerGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConfig)) {
		query["HealthCheckConfig"] = request.HealthCheckConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		query["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupName)) {
		query["ServerGroupName"] = request.ServerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StickySessionConfig)) {
		query["StickySessionConfig"] = request.StickySessionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.UchConfig)) {
		query["UchConfig"] = request.UchConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServerGroupAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServerGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The name of the server group.
 * The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (\\_), and hyphens (-). The name must start with a letter.
 *
 * @param request UpdateServerGroupAttributeRequest
 * @return UpdateServerGroupAttributeResponse
 */
func (client *Client) UpdateServerGroupAttribute(request *UpdateServerGroupAttributeRequest) (_result *UpdateServerGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServerGroupAttributeResponse{}
	_body, _err := client.UpdateServerGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the server. You can specify at most 40 servers in each call.
 * *   If **ServerType** is set to **Ecs**, **Eni**, or **Eci**, set the ServerId parameter to the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance.
 * *   If **ServerType** is set to **Ip**, set the ServerId parameter to an IP address.
 * *   If **ServerType** is set to **Fc**, set the ServerId parameter to the Alibaba Cloud Resource Name (ARN) of a function.
 *
 * @param request UpdateServerGroupServersAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateServerGroupServersAttributeResponse
 */
func (client *Client) UpdateServerGroupServersAttributeWithOptions(request *UpdateServerGroupServersAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServerGroupServersAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		query["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Servers)) {
		query["Servers"] = request.Servers
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServerGroupServersAttribute"),
		Version:     tea.String("2020-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServerGroupServersAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ID of the server. You can specify at most 40 servers in each call.
 * *   If **ServerType** is set to **Ecs**, **Eni**, or **Eci**, set the ServerId parameter to the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance.
 * *   If **ServerType** is set to **Ip**, set the ServerId parameter to an IP address.
 * *   If **ServerType** is set to **Fc**, set the ServerId parameter to the Alibaba Cloud Resource Name (ARN) of a function.
 *
 * @param request UpdateServerGroupServersAttributeRequest
 * @return UpdateServerGroupServersAttributeResponse
 */
func (client *Client) UpdateServerGroupServersAttribute(request *UpdateServerGroupServersAttributeRequest) (_result *UpdateServerGroupServersAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServerGroupServersAttributeResponse{}
	_body, _err := client.UpdateServerGroupServersAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
