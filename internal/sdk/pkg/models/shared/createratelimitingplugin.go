// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/utils"
)

// LimitBy - The entity that is used when aggregating the limits.
type LimitBy string

const (
	LimitByConsumer      LimitBy = "consumer"
	LimitByCredential    LimitBy = "credential"
	LimitByIP            LimitBy = "ip"
	LimitByService       LimitBy = "service"
	LimitByHeader        LimitBy = "header"
	LimitByPath          LimitBy = "path"
	LimitByConsumerGroup LimitBy = "consumer-group"
)

func (e LimitBy) ToPointer() *LimitBy {
	return &e
}

func (e *LimitBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer":
		fallthrough
	case "credential":
		fallthrough
	case "ip":
		fallthrough
	case "service":
		fallthrough
	case "header":
		fallthrough
	case "path":
		fallthrough
	case "consumer-group":
		*e = LimitBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LimitBy: %v", v)
	}
}

// Policy - The rate-limiting policies to use for retrieving and incrementing the limits.
type Policy string

const (
	PolicyLocal   Policy = "local"
	PolicyCluster Policy = "cluster"
	PolicyRedis   Policy = "redis"
)

func (e Policy) ToPointer() *Policy {
	return &e
}

func (e *Policy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "local":
		fallthrough
	case "cluster":
		fallthrough
	case "redis":
		*e = Policy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Policy: %v", v)
	}
}

type CreateRateLimitingPluginConfig struct {
	// The number of HTTP requests that can be made per day.
	Day *float64 `json:"day,omitempty"`
	// Set a custom error code to return when the rate limit is exceeded.
	ErrorCode *float64 `default:"429" json:"error_code"`
	// Set a custom error message to return when the rate limit is exceeded.
	ErrorMessage *string `default:"API rate limit exceeded" json:"error_message"`
	// A boolean value that determines if the requests should be proxied even if Kong has troubles connecting a third-party data store. If `true`, requests will be proxied anyway, effectively disabling the rate-limiting function until the data store is working again. If `false`, then the clients will see `500` errors.
	FaultTolerant *bool `default:"true" json:"fault_tolerant"`
	// A string representing an HTTP header name.
	HeaderName *string `json:"header_name,omitempty"`
	// Optionally hide informative response headers.
	HideClientHeaders *bool `default:"false" json:"hide_client_headers"`
	// The number of HTTP requests that can be made per hour.
	Hour *float64 `json:"hour,omitempty"`
	// The entity that is used when aggregating the limits.
	LimitBy *LimitBy `default:"consumer" json:"limit_by"`
	// The number of HTTP requests that can be made per minute.
	Minute *float64 `json:"minute,omitempty"`
	// The number of HTTP requests that can be made per month.
	Month *float64 `json:"month,omitempty"`
	// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
	Path *string `json:"path,omitempty"`
	// The rate-limiting policies to use for retrieving and incrementing the limits.
	Policy *Policy `default:"local" json:"policy"`
	// When using the `redis` policy, this property specifies the Redis database to use.
	RedisDatabase *int64 `default:"0" json:"redis_database"`
	// A string representing a host name, such as example.com.
	RedisHost *string `json:"redis_host,omitempty"`
	// When using the `redis` policy, this property specifies the password to connect to the Redis server.
	RedisPassword *string `json:"redis_password,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	RedisPort *int64 `default:"6379" json:"redis_port"`
	// A string representing an SNI (server name indication) value for TLS.
	RedisServerName *string `json:"redis_server_name,omitempty"`
	// When using the `redis` policy, this property specifies if SSL is used to connect to the Redis server.
	RedisSsl *bool `default:"false" json:"redis_ssl"`
	// When using the `redis` policy with `redis_ssl` set to `true`, this property specifies it server SSL certificate is validated. Note that you need to configure the lua_ssl_trusted_certificate to specify the CA (or server) certificate used by your Redis server. You may also need to configure lua_ssl_verify_depth accordingly.
	RedisSslVerify *bool `default:"false" json:"redis_ssl_verify"`
	// When using the `redis` policy, this property specifies the timeout in milliseconds of any command submitted to the Redis server.
	RedisTimeout *float64 `default:"2000" json:"redis_timeout"`
	// When using the `redis` policy, this property specifies the username to connect to the Redis server when ACL authentication is desired.
	RedisUsername *string `json:"redis_username,omitempty"`
	// The number of HTTP requests that can be made per second.
	Second *float64 `json:"second,omitempty"`
	// How often to sync counter data to the central data store. A value of -1 results in synchronous behavior.
	SyncRate *float64 `default:"-1" json:"sync_rate"`
	// The number of HTTP requests that can be made per year.
	Year *float64 `json:"year,omitempty"`
}

func (c CreateRateLimitingPluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateRateLimitingPluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateRateLimitingPluginConfig) GetDay() *float64 {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *CreateRateLimitingPluginConfig) GetErrorCode() *float64 {
	if o == nil {
		return nil
	}
	return o.ErrorCode
}

func (o *CreateRateLimitingPluginConfig) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateRateLimitingPluginConfig) GetFaultTolerant() *bool {
	if o == nil {
		return nil
	}
	return o.FaultTolerant
}

func (o *CreateRateLimitingPluginConfig) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *CreateRateLimitingPluginConfig) GetHideClientHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.HideClientHeaders
}

func (o *CreateRateLimitingPluginConfig) GetHour() *float64 {
	if o == nil {
		return nil
	}
	return o.Hour
}

func (o *CreateRateLimitingPluginConfig) GetLimitBy() *LimitBy {
	if o == nil {
		return nil
	}
	return o.LimitBy
}

func (o *CreateRateLimitingPluginConfig) GetMinute() *float64 {
	if o == nil {
		return nil
	}
	return o.Minute
}

func (o *CreateRateLimitingPluginConfig) GetMonth() *float64 {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *CreateRateLimitingPluginConfig) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *CreateRateLimitingPluginConfig) GetPolicy() *Policy {
	if o == nil {
		return nil
	}
	return o.Policy
}

func (o *CreateRateLimitingPluginConfig) GetRedisDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.RedisDatabase
}

func (o *CreateRateLimitingPluginConfig) GetRedisHost() *string {
	if o == nil {
		return nil
	}
	return o.RedisHost
}

func (o *CreateRateLimitingPluginConfig) GetRedisPassword() *string {
	if o == nil {
		return nil
	}
	return o.RedisPassword
}

func (o *CreateRateLimitingPluginConfig) GetRedisPort() *int64 {
	if o == nil {
		return nil
	}
	return o.RedisPort
}

func (o *CreateRateLimitingPluginConfig) GetRedisServerName() *string {
	if o == nil {
		return nil
	}
	return o.RedisServerName
}

func (o *CreateRateLimitingPluginConfig) GetRedisSsl() *bool {
	if o == nil {
		return nil
	}
	return o.RedisSsl
}

func (o *CreateRateLimitingPluginConfig) GetRedisSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.RedisSslVerify
}

func (o *CreateRateLimitingPluginConfig) GetRedisTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RedisTimeout
}

func (o *CreateRateLimitingPluginConfig) GetRedisUsername() *string {
	if o == nil {
		return nil
	}
	return o.RedisUsername
}

func (o *CreateRateLimitingPluginConfig) GetSecond() *float64 {
	if o == nil {
		return nil
	}
	return o.Second
}

func (o *CreateRateLimitingPluginConfig) GetSyncRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SyncRate
}

func (o *CreateRateLimitingPluginConfig) GetYear() *float64 {
	if o == nil {
		return nil
	}
	return o.Year
}

// CreateRateLimitingPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateRateLimitingPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRateLimitingPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateRateLimitingPluginProtocols string

const (
	CreateRateLimitingPluginProtocolsGrpc           CreateRateLimitingPluginProtocols = "grpc"
	CreateRateLimitingPluginProtocolsGrpcs          CreateRateLimitingPluginProtocols = "grpcs"
	CreateRateLimitingPluginProtocolsHTTP           CreateRateLimitingPluginProtocols = "http"
	CreateRateLimitingPluginProtocolsHTTPS          CreateRateLimitingPluginProtocols = "https"
	CreateRateLimitingPluginProtocolsTCP            CreateRateLimitingPluginProtocols = "tcp"
	CreateRateLimitingPluginProtocolsTLS            CreateRateLimitingPluginProtocols = "tls"
	CreateRateLimitingPluginProtocolsTLSPassthrough CreateRateLimitingPluginProtocols = "tls_passthrough"
	CreateRateLimitingPluginProtocolsUDP            CreateRateLimitingPluginProtocols = "udp"
	CreateRateLimitingPluginProtocolsWs             CreateRateLimitingPluginProtocols = "ws"
	CreateRateLimitingPluginProtocolsWss            CreateRateLimitingPluginProtocols = "wss"
)

func (e CreateRateLimitingPluginProtocols) ToPointer() *CreateRateLimitingPluginProtocols {
	return &e
}

func (e *CreateRateLimitingPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = CreateRateLimitingPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRateLimitingPluginProtocols: %v", v)
	}
}

// CreateRateLimitingPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateRateLimitingPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRateLimitingPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRateLimitingPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateRateLimitingPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRateLimitingPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRateLimitingPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CreateRateLimitingPlugin struct {
	Config CreateRateLimitingPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *CreateRateLimitingPluginConsumer `json:"consumer,omitempty"`
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"rate-limiting" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateRateLimitingPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateRateLimitingPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateRateLimitingPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (c CreateRateLimitingPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateRateLimitingPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateRateLimitingPlugin) GetConfig() CreateRateLimitingPluginConfig {
	if o == nil {
		return CreateRateLimitingPluginConfig{}
	}
	return o.Config
}

func (o *CreateRateLimitingPlugin) GetConsumer() *CreateRateLimitingPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateRateLimitingPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateRateLimitingPlugin) GetName() string {
	return "rate-limiting"
}

func (o *CreateRateLimitingPlugin) GetProtocols() []CreateRateLimitingPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateRateLimitingPlugin) GetRoute() *CreateRateLimitingPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateRateLimitingPlugin) GetService() *CreateRateLimitingPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *CreateRateLimitingPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
