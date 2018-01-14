package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// ConnectionLog connection log
// swagger:model ConnectionLog
type ConnectionLog struct {

	// Placeholder for description of property adf of obj type ConnectionLog field type str  type boolean
	// Required: true
	Adf bool `json:"adf"`

	// . Units(MILLISECONDS).
	AverageTurntime int32 `json:"average_turntime,omitempty"`

	// Number of client_dest_port.
	// Required: true
	ClientDestPort int32 `json:"client_dest_port"`

	// Number of client_ip.
	// Required: true
	ClientIP int32 `json:"client_ip"`

	// client_location of ConnectionLog.
	ClientLocation string `json:"client_location,omitempty"`

	// . Units(MILLISECONDS).
	// Required: true
	ClientRtt int32 `json:"client_rtt"`

	// Number of client_src_port.
	// Required: true
	ClientSrcPort int32 `json:"client_src_port"`

	// Placeholder for description of property connection_ended of obj type ConnectionLog field type str  type boolean
	// Required: true
	ConnectionEnded bool `json:"connection_ended"`

	//  Enum options - DNS_ENTRY_PASS_THROUGH, DNS_ENTRY_GSLB, DNS_ENTRY_VIRTUALSERVICE, DNS_ENTRY_STATIC, DNS_ENTRY_POLICY, DNS_ENTRY_LOCAL.
	DNSEtype string `json:"dns_etype,omitempty"`

	// dns_fqdn of ConnectionLog.
	DNSFqdn string `json:"dns_fqdn,omitempty"`

	// Number of dns_ips.
	DNSIps []int64 `json:"dns_ips,omitempty,omitempty"`

	//  Enum options - DNS_RECORD_OTHER, DNS_RECORD_A, DNS_RECORD_NS, DNS_RECORD_CNAME, DNS_RECORD_SOA, DNS_RECORD_PTR, DNS_RECORD_HINFO, DNS_RECORD_MX, DNS_RECORD_TXT, DNS_RECORD_RP, DNS_RECORD_DNSKEY, DNS_RECORD_AAAA, DNS_RECORD_SRV, DNS_RECORD_OPT, DNS_RECORD_RRSIG, DNS_RECORD_AXFR, DNS_RECORD_ANY.
	DNSQtype string `json:"dns_qtype,omitempty"`

	//  Field introduced in 17.1.1.
	DNSRequest *DNSRequest `json:"dns_request,omitempty"`

	// Placeholder for description of property dns_response of obj type ConnectionLog field type str  type object
	DNSResponse *DNSResponse `json:"dns_response,omitempty"`

	// gslbpool_name of ConnectionLog.
	GslbpoolName string `json:"gslbpool_name,omitempty"`

	// gslbservice of ConnectionLog.
	Gslbservice string `json:"gslbservice,omitempty"`

	// gslbservice_name of ConnectionLog.
	GslbserviceName string `json:"gslbservice_name,omitempty"`

	// Number of log_id.
	// Required: true
	LogID int32 `json:"log_id"`

	// microservice of ConnectionLog.
	Microservice string `json:"microservice,omitempty"`

	// microservice_name of ConnectionLog.
	MicroserviceName string `json:"microservice_name,omitempty"`

	// . Units(BYTES).
	// Required: true
	Mss int32 `json:"mss"`

	// network_security_policy_rule_name of ConnectionLog.
	NetworkSecurityPolicyRuleName string `json:"network_security_policy_rule_name,omitempty"`

	// Number of num_syn_retransmit.
	NumSynRetransmit int32 `json:"num_syn_retransmit,omitempty"`

	// Number of num_transaction.
	NumTransaction int32 `json:"num_transaction,omitempty"`

	// Number of num_window_shrink.
	NumWindowShrink int32 `json:"num_window_shrink,omitempty"`

	// Number of out_of_orders.
	// Required: true
	OutOfOrders int32 `json:"out_of_orders"`

	// pool of ConnectionLog.
	Pool string `json:"pool,omitempty"`

	// pool_name of ConnectionLog.
	PoolName string `json:"pool_name,omitempty"`

	//  Enum options - PROTOCOL_ICMP, PROTOCOL_TCP, PROTOCOL_UDP.
	Protocol string `json:"protocol,omitempty"`

	// Version of proxy protocol used to convey client connection information to the back-end servers.  A value of 0 indicates that proxy protocol is not used.  A value of 1 or 2 indicates the version of proxy protocol used. Enum options - PROXY_PROTOCOL_VERSION_1, PROXY_PROTOCOL_VERSION_2.
	ProxyProtocol string `json:"proxy_protocol,omitempty"`

	// Number of report_timestamp.
	// Required: true
	ReportTimestamp int64 `json:"report_timestamp"`

	// Number of retransmits.
	// Required: true
	Retransmits int32 `json:"retransmits"`

	// . Units(BYTES).
	// Required: true
	RxBytes int64 `json:"rx_bytes"`

	// Number of rx_pkts.
	// Required: true
	RxPkts int64 `json:"rx_pkts"`

	// Number of server_conn_src_ip.
	// Required: true
	ServerConnSrcIP int32 `json:"server_conn_src_ip"`

	// Number of server_dest_port.
	// Required: true
	ServerDestPort int32 `json:"server_dest_port"`

	// Number of server_ip.
	// Required: true
	ServerIP int32 `json:"server_ip"`

	// server_name of ConnectionLog.
	ServerName string `json:"server_name,omitempty"`

	// Number of server_num_window_shrink.
	ServerNumWindowShrink int32 `json:"server_num_window_shrink,omitempty"`

	// Number of server_out_of_orders.
	// Required: true
	ServerOutOfOrders int32 `json:"server_out_of_orders"`

	// Number of server_retransmits.
	// Required: true
	ServerRetransmits int32 `json:"server_retransmits"`

	// . Units(MILLISECONDS).
	// Required: true
	ServerRtt int32 `json:"server_rtt"`

	// . Units(BYTES).
	// Required: true
	ServerRxBytes int64 `json:"server_rx_bytes"`

	// Number of server_rx_pkts.
	// Required: true
	ServerRxPkts int64 `json:"server_rx_pkts"`

	// Number of server_src_port.
	// Required: true
	ServerSrcPort int32 `json:"server_src_port"`

	// Number of server_timeouts.
	// Required: true
	ServerTimeouts int32 `json:"server_timeouts"`

	// . Units(BYTES).
	// Required: true
	ServerTotalBytes int64 `json:"server_total_bytes"`

	// Number of server_total_pkts.
	// Required: true
	ServerTotalPkts int64 `json:"server_total_pkts"`

	// . Units(BYTES).
	// Required: true
	ServerTxBytes int64 `json:"server_tx_bytes"`

	// Number of server_tx_pkts.
	// Required: true
	ServerTxPkts int64 `json:"server_tx_pkts"`

	// Number of server_zero_window_size_events.
	// Required: true
	ServerZeroWindowSizeEvents int32 `json:"server_zero_window_size_events"`

	// service_engine of ConnectionLog.
	ServiceEngine string `json:"service_engine,omitempty"`

	// significance of ConnectionLog.
	Significance string `json:"significance,omitempty"`

	// Number of significant.
	// Required: true
	Significant int64 `json:"significant"`

	// List of enums which indicate why a log is significant. Enum options - ADF_CLIENT_CONN_SETUP_REFUSED, ADF_SERVER_CONN_SETUP_REFUSED, ADF_CLIENT_CONN_SETUP_TIMEDOUT, ADF_SERVER_CONN_SETUP_TIMEDOUT, ADF_CLIENT_CONN_SETUP_FAILED_INTERNAL, ADF_SERVER_CONN_SETUP_FAILED_INTERNAL, ADF_CLIENT_CONN_SETUP_FAILED_BAD_PACKET, ADF_UDP_CONN_SETUP_FAILED_INTERNAL, ADF_UDP_SERVER_CONN_SETUP_FAILED_INTERNAL, ADF_CLIENT_SENT_RESET, ADF_SERVER_SENT_RESET, ADF_CLIENT_CONN_TIMEDOUT, ADF_SERVER_CONN_TIMEDOUT, ADF_USER_DELETE_OPERATION, ADF_CLIENT_REQUEST_TIMEOUT, ADF_CLIENT_CONN_ABORTED, ADF_CLIENT_SSL_HANDSHAKE_FAILURE, ADF_CLIENT_CONN_FAILED, ADF_SERVER_CERTIFICATE_VERIFICATION_FAILED, ADF_SERVER_SIDE_SSL_HANDSHAKE_FAILED, ADF_IDLE_TIMEDOUT, ADF_CLIENT_HIGH_TIMEOUT_RETRANSMITS, ADF_SERVER_HIGH_TIMEOUT_RETRANSMITS, ADF_CLIENT_HIGH_RX_ZERO_WINDOW_SIZE_EVENTS, ADF_SERVER_HIGH_RX_ZERO_WINDOW_SIZE_EVENTS, ADF_CLIENT_RTT_ABOVE_SEC, ADF_SERVER_RTT_ABOVE_500MS, ADF_CLIENT_HIGH_TOTAL_RETRANSMITS, ADF_SERVER_HIGH_TOTAL_RETRANSMITS, ADF_CLIENT_HIGH_OUT_OF_ORDERS, ADF_SERVER_HIGH_OUT_OF_ORDERS, ADF_CLIENT_HIGH_TX_ZERO_WINDOW_SIZE_EVENTS, ADF_SERVER_HIGH_TX_ZERO_WINDOW_SIZE_EVENTS, ADF_CLIENT_POSSIBLE_WINDOW_STUCK, ADF_SERVER_POSSIBLE_WINDOW_STUCK, ADF_SERVER_UNANSWERED_SYNS, ADF_CLIENT_CLOSE_CONNECTION_ON_VS_UPDATE, ADF_RESPONSE_CODE_4XX, ADF_RESPONSE_CODE_5XX, ADF_LOAD_BALANCING_FAILED, ADF_DATASCRIPT_EXECUTION_FAILED, ADF_REQUEST_NO_POOL, ADF_RATE_LIMIT_DROP_CLIENT_IP, ADF_RATE_LIMIT_DROP_URI, ADF_RATE_LIMIT_DROP_CLIENT_IP_URI, ADF_RATE_LIMIT_DROP_UNKNOWN_URI, ADF_RATE_LIMIT_DROP_BAD_URI, ADF_REQUEST_VIRTUAL_HOSTING_APP_SELECT_FAILED, ADF_RATE_LIMIT_DROP_UNKNOWN_CIP, ADF_RATE_LIMIT_DROP_BAD_CIP, ADF_RATE_LIMIT_DROP_CLIENT_IP_BAD, ADF_RATE_LIMIT_DROP_URI_BAD, ADF_RATE_LIMIT_DROP_CLIENT_IP_URI_BAD, ADF_RATE_LIMIT_DROP_REQ, ADF_RATE_LIMIT_DROP_CLIENT_IP_CONN, ADF_RATE_LIMIT_DROP_CONN, ADF_RATE_LIMIT_DROP_HEADER, ADF_HTTP_VERSION_LT_1_0, ADF_CLIENT_HIGH_RESPONSE_TIME, ADF_SERVER_HIGH_RESPONSE_TIME, ADF_PERSISTENT_SERVER_CHANGE, ADF_DOS_SERVER_BAD_GATEWAY, ADF_DOS_SERVER_GATEWAY_TIMEOUT, ADF_DOS_CLIENT_SENT_RESET, ADF_DOS_CLIENT_CONN_TIMEOUT, ADF_DOS_CLIENT_REQUEST_TIMEOUT, ADF_DOS_CLIENT_CONN_ABORTED, ADF_DOS_CLIENT_BAD_REQUEST, ADF_DOS_CLIENT_REQUEST_ENTITY_TOO_LARGE, ADF_DOS_CLIENT_REQUEST_URI_TOO_LARGE, ADF_DOS_CLIENT_REQUEST_HEADER_TOO_LARGE, ADF_DOS_CLIENT_CLOSED_REQUEST, ADF_DOS_SSL_ERROR, ADF_X509_CLIENT_CERTIFICATE_VERIFICATION_FAILED, ADF_X509_CLIENT_CERTIFICATE_NOT_YET_VALID, ADF_X509_CLIENT_CERTIFICATE_EXPIRED, ADF_X509_CLIENT_CERTIFICATE_REVOKED, ADF_X509_CLIENT_CERTIFICATE_INVALID_CA, ADF_X509_CLIENT_CERTIFICATE_CRL_NOT_PRESENT, ADF_X509_CLIENT_CERTIFICATE_CRL_NOT_YET_VALID, ADF_X509_CLIENT_CERTIFICATE_CRL_EXPIRED, ADF_X509_CLIENT_CERTIFICATE_CRL_ERROR, ADF_X509_CLIENT_CERTIFICATE_CHAINING_ERROR, ADF_X509_CLIENT_CERTIFICATE_INTERNAL_ERROR, ADF_X509_CLIENT_CERTIFICATE_FORMAT_ERROR, ADF_UDP_PORT_NOT_REACHABLE, ADF_UDP_CONN_TIMEOUT, ADF_X509_SERVER_CERTIFICATE_VERIFICATION_FAILED, ADF_X509_SERVER_CERTIFICATE_NOT_YET_VALID, ADF_X509_SERVER_CERTIFICATE_EXPIRED, ADF_X509_SERVER_CERTIFICATE_REVOKED, ADF_X509_SERVER_CERTIFICATE_INVALID_CA, ADF_X509_SERVER_CERTIFICATE_CRL_NOT_PRESENT, ADF_X509_SERVER_CERTIFICATE_CRL_NOT_YET_VALID, ADF_X509_SERVER_CERTIFICATE_CRL_EXPIRED, ADF_X509_SERVER_CERTIFICATE_CRL_ERROR, ADF_X509_SERVER_CERTIFICATE_CHAINING_ERROR, ADF_X509_SERVER_CERTIFICATE_INTERNAL_ERROR, ADF_X509_SERVER_CERTIFICATE_FORMAT_ERROR, ADF_X509_SERVER_CERTIFICATE_HOSTNAME_ERROR, ADF_SSL_R_BAD_CHANGE_CIPHER_SPEC, ADF_SSL_R_BLOCK_CIPHER_PAD_IS_WRONG, ADF_SSL_R_DIGEST_CHECK_FAILED, ADF_SSL_R_ERROR_IN_RECEIVED_CIPHER_LIST, ADF_SSL_R_EXCESSIVE_MESSAGE_SIZE, ADF_SSL_R_LENGTH_MISMATCH, ADF_SSL_R_NO_CIPHERS_PASSED, ADF_SSL_R_NO_CIPHERS_SPECIFIED, ADF_SSL_R_NO_COMPRESSION_SPECIFIED, ADF_SSL_R_NO_SHARED_CIPHER, ADF_SSL_R_RECORD_LENGTH_MISMATCH, ADF_SSL_R_PARSE_TLSEXT, ADF_SSL_R_UNEXPECTED_MESSAGE, ADF_SSL_R_UNEXPECTED_RECORD, ADF_SSL_R_UNKNOWN_ALERT_TYPE, ADF_SSL_R_UNKNOWN_PROTOCOL, ADF_SSL_R_WRONG_VERSION_NUMBER, ADF_SSL_R_DECRYPTION_FAILED_OR_BAD_RECORD_MAC, ADF_SSL_R_RENEGOTIATE_EXT_TOO_LONG, ADF_SSL_R_RENEGOTIATION_ENCODING_ERR, ADF_SSL_R_RENEGOTIATION_MISMATCH, ADF_SSL_R_UNSAFE_LEGACY_RENEGOTIATION_DISABLED, ADF_SSL_R_SCSV_RECEIVED_WHEN_RENEGOTIATING, ADF_SSL_R_INAPPROPRIATE_FALLBACK, ADF_SSL_R_SSLV3_ALERT_UNEXPECTED_MESSAGE, ADF_SSL_R_SSLV3_ALERT_BAD_RECORD_MAC, ADF_SSL_R_TLSV1_ALERT_DECRYPTION_FAILED, ADF_SSL_R_TLSV1_ALERT_RECORD_OVERFLOW, ADF_SSL_R_SSLV3_ALERT_DECOMPRESSION_FAILURE, ADF_SSL_R_SSLV3_ALERT_HANDSHAKE_FAILURE, ADF_SSL_R_SSLV3_ALERT_NO_CERTIFICATE, ADF_SSL_R_SSLV3_ALERT_BAD_CERTIFICATE, ADF_SSL_R_SSLV3_ALERT_UNSUPPORTED_CERTIFICATE, ADF_SSL_R_SSLV3_ALERT_CERTIFICATE_REVOKED, ADF_SSL_R_SSLV3_ALERT_CERTIFICATE_EXPIRED, ADF_SSL_R_SSLV3_ALERT_CERTIFICATE_UNKNOWN, ADF_SSL_R_SSLV3_ALERT_ILLEGAL_PARAMETER, ADF_SSL_R_TLSV1_ALERT_UNKNOWN_CA, ADF_SSL_R_TLSV1_ALERT_ACCESS_DENIED, ADF_SSL_R_TLSV1_ALERT_DECODE_ERROR, ADF_SSL_R_TLSV1_ALERT_DECRYPT_ERROR, ADF_SSL_R_TLSV1_ALERT_EXPORT_RESTRICTION, ADF_SSL_R_TLSV1_ALERT_PROTOCOL_VERSION, ADF_SSL_R_TLSV1_ALERT_INSUFFICIENT_SECURITY, ADF_SSL_R_TLSV1_ALERT_INTERNAL_ERROR, ADF_SSL_R_TLSV1_ALERT_USER_CANCELLED, ADF_SSL_R_TLSV1_ALERT_NO_RENEGOTIATION, ADF_CLIENT_AUTH_UNKNOWN_USER, ADF_CLIENT_AUTH_LOGIN_FAILED, ADF_CLIENT_AUTH_MISSING_CREDENTIALS, ADF_CLIENT_AUTH_SERVER_CONN_ERROR, ADF_CLIENT_AUTH_USER_NOT_AUTHORIZED, ADF_CLIENT_AUTH_TIMED_OUT, ADF_CLIENT_AUTH_UNKNOWN_ERROR, ADF_CLIENT_DNS_FAILED_INVALID_QUERY, ADF_CLIENT_DNS_FAILED_INVALID_DOMAIN, ADF_CLIENT_DNS_FAILED_NO_SERVICE, ADF_CLIENT_DNS_FAILED_GS_DOWN, ADF_CLIENT_DNS_FAILED_NO_VALID_GS_MEMBER, ADF_SERVER_DNS_ERROR_RESPONSE, ADF_CLIENT_DNS_FAILED_UNSUPPORTED_QUERY, ADF_MEMORY_EXHAUSTED, ADF_CLIENT_DNS_POLICY_DROP, ADF_WAF_MATCH, ADF_USER_DELETE_OPERATION_DATASCRIPT_RESET_CONN, ADF_USER_DELETE_OPERATION_HTTP_RULE_SECURITY_ACTION_CLOSE_CONN, ADF_USER_DELETE_OPERATION_HTTP_RULE_SECURITY_RATE_LIMIT_ACTION_CLOSE_CONN, ADF_USER_DELETE_OPERATION_HTTP_RULE_MISSING_TOKEN_ACTION_CLOSE_CONN, ADF_HTTP_BAD_REQUEST_INVALID_HOST_IN_REQUEST_LINE, ADF_HTTP_BAD_REQUEST_RECEIVED_VERSION_LESS_THAN_10, ADF_HTTP_NOT_ALLOWED_DATASCRIPT_RESPONSE_RETURNED_4XX, ADF_HTTP_NOT_ALLOWED_RUM_FLAGGED_INVALID_METHOD, ADF_HTTP_NOT_ALLOWED_UNSUPPORTED_TRACE_METHOD, ADF_HTTP_REQUEST_TIMEOUT_WAITING_FOR_CLIENT, ADF_HTTP_BAD_REQUEST_CLIENT_SENT_INVALID_CONTENT_LENGTH, ADF_HTTP_BAD_REQUEST_CLIENT_SENT_HTTP11_WITHOUT_HOST_HDR, ADF_HTTP_BAD_REQUEST_FAILED_TO_PARSE_URI, ADF_HTTP_BAD_REQUEST_INVALID_HEADER_LINE, ADF_HTTP_BAD_REQUEST_ERROR_WHILE_READING_CLIENT_HEADERS, ADF_HTTP_BAD_REQUEST_CLIENT_SENT_DUPLICATE_HEADER, ADF_HTTP_BAD_REQUEST_CLIENT_SENT_INVALID_HOST_HEADER, ADF_HTTP_NOT_IMPLEMENTED_CLIENT_SENT_UNKNOWN_TRANSFER_ENCODING, ADF_HTTP_BAD_REQUEST_REQUESTED_SERVER_NAME_DIFFERS, ADF_HTTP_BAD_REQUEST_CLIENT_SENT_INVALID_CHUNKED_BODY, ADF_HTTP_BAD_REQUEST_INVALID_HEADER_IN_SPDY, ADF_HTTP_BAD_REQUEST_INVALID_HEADER_BLOCK_IN_SPDY, ADF_HTTP_BAD_REQUEST_DATA_ERROR_IN_SPDY, ADF_HTTP_BAD_REQUEST_NO_METHOD_URI_OR_PROT_IN_REQ_CREATE_SPDY, ADF_HTTP_BAD_REQUEST_CLIENT_PREMATURELY_CLOSED_SPDY_STREAM, ADF_HTTP_BAD_REQUEST_DATA_ERROR_IN_SPDY_READ_REQ_BODY, ADF_HTTP_BAD_REQUEST_CERT_ERROR, ADF_HTTP_BAD_REQUEST_PLAIN_HTTP_REQUEST_SENT_ON_HTTPS_PORT, ADF_HTTP_BAD_REQUEST_NO_CERT_ERROR, ADF_HTTP_BAD_REQUEST_HEADER_TOO_LARGE, ADF_SERVER_HIGH_RESPONSE_TIME_L7, ADF_SERVER_HIGH_RESPONSE_TIME_L4, ADF_COOKIE_SIZE_GREATER_THAN_MAX, ADF_COOKIE_SIZE_LESS_THAN_MIN_COOKIE_LEN, ADF_PERSISTENCE_PROFILE_KEYS_NOT_CONFIGURED, ADF_PERSISTENCE_COOKIE_VERSION_MISMATCH, ADF_COOKIE_ABSENT_FROM_KEYS_IN_PERSISTENCE_PROFILE, ADF_GSLB_SITE_PERSISTENCE_REMOTE_SITE_DOWN.
	SignificantLog []string `json:"significant_log,omitempty"`

	//  Field introduced in 17.2.5.
	SniHostname string `json:"sni_hostname,omitempty"`

	// ssl_cipher of ConnectionLog.
	SslCipher string `json:"ssl_cipher,omitempty"`

	// ssl_session_id of ConnectionLog.
	SslSessionID string `json:"ssl_session_id,omitempty"`

	// ssl_version of ConnectionLog.
	SslVersion string `json:"ssl_version,omitempty"`

	// Number of start_timestamp.
	// Required: true
	StartTimestamp int64 `json:"start_timestamp"`

	// Number of timeouts.
	// Required: true
	Timeouts int32 `json:"timeouts"`

	// . Units(BYTES).
	TotalBytes int64 `json:"total_bytes,omitempty"`

	// Number of total_pkts.
	TotalPkts int64 `json:"total_pkts,omitempty"`

	// . Units(MILLISECONDS).
	TotalTime int64 `json:"total_time,omitempty"`

	// . Units(BYTES).
	// Required: true
	TxBytes int64 `json:"tx_bytes"`

	// Number of tx_pkts.
	// Required: true
	TxPkts int64 `json:"tx_pkts"`

	// Placeholder for description of property udf of obj type ConnectionLog field type str  type boolean
	// Required: true
	Udf bool `json:"udf"`

	// Number of vcpu_id.
	// Required: true
	VcpuID int32 `json:"vcpu_id"`

	// virtualservice of ConnectionLog.
	// Required: true
	Virtualservice string `json:"virtualservice"`

	//  Field introduced in 17.1.1.
	VsIP int32 `json:"vs_ip,omitempty"`

	// Number of zero_window_size_events.
	// Required: true
	ZeroWindowSizeEvents int32 `json:"zero_window_size_events"`
}
