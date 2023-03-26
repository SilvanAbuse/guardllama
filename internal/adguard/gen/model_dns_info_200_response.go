/*
AdGuard Home

AdGuard Home REST-ish API.  Our admin web interface is built on top of this REST-ish API. 

API version: 0.107
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package adguard

import (
	"encoding/json"
)

// DnsInfo200Response struct for DnsInfo200Response
type DnsInfo200Response struct {
	// Bootstrap servers, port is optional after colon.  Empty value will reset it to default values. 
	BootstrapDns []string `json:"bootstrap_dns,omitempty"`
	// Upstream servers, port is optional after colon.  Empty value will reset it to default values. 
	UpstreamDns []string `json:"upstream_dns,omitempty"`
	UpstreamDnsFile *string `json:"upstream_dns_file,omitempty"`
	ProtectionEnabled *bool `json:"protection_enabled,omitempty"`
	DhcpAvailable *bool `json:"dhcp_available,omitempty"`
	Ratelimit *int32 `json:"ratelimit,omitempty"`
	BlockingMode *string `json:"blocking_mode,omitempty"`
	BlockingIpv4 *string `json:"blocking_ipv4,omitempty"`
	BlockingIpv6 *string `json:"blocking_ipv6,omitempty"`
	EdnsCsEnabled *bool `json:"edns_cs_enabled,omitempty"`
	DisableIpv6 *bool `json:"disable_ipv6,omitempty"`
	DnssecEnabled *bool `json:"dnssec_enabled,omitempty"`
	CacheSize *int32 `json:"cache_size,omitempty"`
	CacheTtlMin *int32 `json:"cache_ttl_min,omitempty"`
	CacheTtlMax *int32 `json:"cache_ttl_max,omitempty"`
	CacheOptimistic *bool `json:"cache_optimistic,omitempty"`
	UpstreamMode *string `json:"upstream_mode,omitempty"`
	UsePrivatePtrResolvers *bool `json:"use_private_ptr_resolvers,omitempty"`
	ResolveClients *bool `json:"resolve_clients,omitempty"`
	// Upstream servers, port is optional after colon.  Empty value will reset it to default values. 
	LocalPtrUpstreams []string `json:"local_ptr_upstreams,omitempty"`
	DefaultLocalPtrUpstreams []string `json:"default_local_ptr_upstreams,omitempty"`
}

// NewDnsInfo200Response instantiates a new DnsInfo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsInfo200Response() *DnsInfo200Response {
	this := DnsInfo200Response{}
	return &this
}

// NewDnsInfo200ResponseWithDefaults instantiates a new DnsInfo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsInfo200ResponseWithDefaults() *DnsInfo200Response {
	this := DnsInfo200Response{}
	return &this
}

// GetBootstrapDns returns the BootstrapDns field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetBootstrapDns() []string {
	if o == nil || o.BootstrapDns == nil {
		var ret []string
		return ret
	}
	return o.BootstrapDns
}

// GetBootstrapDnsOk returns a tuple with the BootstrapDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetBootstrapDnsOk() ([]string, bool) {
	if o == nil || o.BootstrapDns == nil {
		return nil, false
	}
	return o.BootstrapDns, true
}

// HasBootstrapDns returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasBootstrapDns() bool {
	if o != nil && o.BootstrapDns != nil {
		return true
	}

	return false
}

// SetBootstrapDns gets a reference to the given []string and assigns it to the BootstrapDns field.
func (o *DnsInfo200Response) SetBootstrapDns(v []string) {
	o.BootstrapDns = v
}

// GetUpstreamDns returns the UpstreamDns field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetUpstreamDns() []string {
	if o == nil || o.UpstreamDns == nil {
		var ret []string
		return ret
	}
	return o.UpstreamDns
}

// GetUpstreamDnsOk returns a tuple with the UpstreamDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetUpstreamDnsOk() ([]string, bool) {
	if o == nil || o.UpstreamDns == nil {
		return nil, false
	}
	return o.UpstreamDns, true
}

// HasUpstreamDns returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasUpstreamDns() bool {
	if o != nil && o.UpstreamDns != nil {
		return true
	}

	return false
}

// SetUpstreamDns gets a reference to the given []string and assigns it to the UpstreamDns field.
func (o *DnsInfo200Response) SetUpstreamDns(v []string) {
	o.UpstreamDns = v
}

// GetUpstreamDnsFile returns the UpstreamDnsFile field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetUpstreamDnsFile() string {
	if o == nil || o.UpstreamDnsFile == nil {
		var ret string
		return ret
	}
	return *o.UpstreamDnsFile
}

// GetUpstreamDnsFileOk returns a tuple with the UpstreamDnsFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetUpstreamDnsFileOk() (*string, bool) {
	if o == nil || o.UpstreamDnsFile == nil {
		return nil, false
	}
	return o.UpstreamDnsFile, true
}

// HasUpstreamDnsFile returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasUpstreamDnsFile() bool {
	if o != nil && o.UpstreamDnsFile != nil {
		return true
	}

	return false
}

// SetUpstreamDnsFile gets a reference to the given string and assigns it to the UpstreamDnsFile field.
func (o *DnsInfo200Response) SetUpstreamDnsFile(v string) {
	o.UpstreamDnsFile = &v
}

// GetProtectionEnabled returns the ProtectionEnabled field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetProtectionEnabled() bool {
	if o == nil || o.ProtectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ProtectionEnabled
}

// GetProtectionEnabledOk returns a tuple with the ProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetProtectionEnabledOk() (*bool, bool) {
	if o == nil || o.ProtectionEnabled == nil {
		return nil, false
	}
	return o.ProtectionEnabled, true
}

// HasProtectionEnabled returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasProtectionEnabled() bool {
	if o != nil && o.ProtectionEnabled != nil {
		return true
	}

	return false
}

// SetProtectionEnabled gets a reference to the given bool and assigns it to the ProtectionEnabled field.
func (o *DnsInfo200Response) SetProtectionEnabled(v bool) {
	o.ProtectionEnabled = &v
}

// GetDhcpAvailable returns the DhcpAvailable field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetDhcpAvailable() bool {
	if o == nil || o.DhcpAvailable == nil {
		var ret bool
		return ret
	}
	return *o.DhcpAvailable
}

// GetDhcpAvailableOk returns a tuple with the DhcpAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetDhcpAvailableOk() (*bool, bool) {
	if o == nil || o.DhcpAvailable == nil {
		return nil, false
	}
	return o.DhcpAvailable, true
}

// HasDhcpAvailable returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasDhcpAvailable() bool {
	if o != nil && o.DhcpAvailable != nil {
		return true
	}

	return false
}

// SetDhcpAvailable gets a reference to the given bool and assigns it to the DhcpAvailable field.
func (o *DnsInfo200Response) SetDhcpAvailable(v bool) {
	o.DhcpAvailable = &v
}

// GetRatelimit returns the Ratelimit field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetRatelimit() int32 {
	if o == nil || o.Ratelimit == nil {
		var ret int32
		return ret
	}
	return *o.Ratelimit
}

// GetRatelimitOk returns a tuple with the Ratelimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetRatelimitOk() (*int32, bool) {
	if o == nil || o.Ratelimit == nil {
		return nil, false
	}
	return o.Ratelimit, true
}

// HasRatelimit returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasRatelimit() bool {
	if o != nil && o.Ratelimit != nil {
		return true
	}

	return false
}

// SetRatelimit gets a reference to the given int32 and assigns it to the Ratelimit field.
func (o *DnsInfo200Response) SetRatelimit(v int32) {
	o.Ratelimit = &v
}

// GetBlockingMode returns the BlockingMode field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetBlockingMode() string {
	if o == nil || o.BlockingMode == nil {
		var ret string
		return ret
	}
	return *o.BlockingMode
}

// GetBlockingModeOk returns a tuple with the BlockingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetBlockingModeOk() (*string, bool) {
	if o == nil || o.BlockingMode == nil {
		return nil, false
	}
	return o.BlockingMode, true
}

// HasBlockingMode returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasBlockingMode() bool {
	if o != nil && o.BlockingMode != nil {
		return true
	}

	return false
}

// SetBlockingMode gets a reference to the given string and assigns it to the BlockingMode field.
func (o *DnsInfo200Response) SetBlockingMode(v string) {
	o.BlockingMode = &v
}

// GetBlockingIpv4 returns the BlockingIpv4 field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetBlockingIpv4() string {
	if o == nil || o.BlockingIpv4 == nil {
		var ret string
		return ret
	}
	return *o.BlockingIpv4
}

// GetBlockingIpv4Ok returns a tuple with the BlockingIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetBlockingIpv4Ok() (*string, bool) {
	if o == nil || o.BlockingIpv4 == nil {
		return nil, false
	}
	return o.BlockingIpv4, true
}

// HasBlockingIpv4 returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasBlockingIpv4() bool {
	if o != nil && o.BlockingIpv4 != nil {
		return true
	}

	return false
}

// SetBlockingIpv4 gets a reference to the given string and assigns it to the BlockingIpv4 field.
func (o *DnsInfo200Response) SetBlockingIpv4(v string) {
	o.BlockingIpv4 = &v
}

// GetBlockingIpv6 returns the BlockingIpv6 field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetBlockingIpv6() string {
	if o == nil || o.BlockingIpv6 == nil {
		var ret string
		return ret
	}
	return *o.BlockingIpv6
}

// GetBlockingIpv6Ok returns a tuple with the BlockingIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetBlockingIpv6Ok() (*string, bool) {
	if o == nil || o.BlockingIpv6 == nil {
		return nil, false
	}
	return o.BlockingIpv6, true
}

// HasBlockingIpv6 returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasBlockingIpv6() bool {
	if o != nil && o.BlockingIpv6 != nil {
		return true
	}

	return false
}

// SetBlockingIpv6 gets a reference to the given string and assigns it to the BlockingIpv6 field.
func (o *DnsInfo200Response) SetBlockingIpv6(v string) {
	o.BlockingIpv6 = &v
}

// GetEdnsCsEnabled returns the EdnsCsEnabled field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetEdnsCsEnabled() bool {
	if o == nil || o.EdnsCsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EdnsCsEnabled
}

// GetEdnsCsEnabledOk returns a tuple with the EdnsCsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetEdnsCsEnabledOk() (*bool, bool) {
	if o == nil || o.EdnsCsEnabled == nil {
		return nil, false
	}
	return o.EdnsCsEnabled, true
}

// HasEdnsCsEnabled returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasEdnsCsEnabled() bool {
	if o != nil && o.EdnsCsEnabled != nil {
		return true
	}

	return false
}

// SetEdnsCsEnabled gets a reference to the given bool and assigns it to the EdnsCsEnabled field.
func (o *DnsInfo200Response) SetEdnsCsEnabled(v bool) {
	o.EdnsCsEnabled = &v
}

// GetDisableIpv6 returns the DisableIpv6 field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetDisableIpv6() bool {
	if o == nil || o.DisableIpv6 == nil {
		var ret bool
		return ret
	}
	return *o.DisableIpv6
}

// GetDisableIpv6Ok returns a tuple with the DisableIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetDisableIpv6Ok() (*bool, bool) {
	if o == nil || o.DisableIpv6 == nil {
		return nil, false
	}
	return o.DisableIpv6, true
}

// HasDisableIpv6 returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasDisableIpv6() bool {
	if o != nil && o.DisableIpv6 != nil {
		return true
	}

	return false
}

// SetDisableIpv6 gets a reference to the given bool and assigns it to the DisableIpv6 field.
func (o *DnsInfo200Response) SetDisableIpv6(v bool) {
	o.DisableIpv6 = &v
}

// GetDnssecEnabled returns the DnssecEnabled field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetDnssecEnabled() bool {
	if o == nil || o.DnssecEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DnssecEnabled
}

// GetDnssecEnabledOk returns a tuple with the DnssecEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetDnssecEnabledOk() (*bool, bool) {
	if o == nil || o.DnssecEnabled == nil {
		return nil, false
	}
	return o.DnssecEnabled, true
}

// HasDnssecEnabled returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasDnssecEnabled() bool {
	if o != nil && o.DnssecEnabled != nil {
		return true
	}

	return false
}

// SetDnssecEnabled gets a reference to the given bool and assigns it to the DnssecEnabled field.
func (o *DnsInfo200Response) SetDnssecEnabled(v bool) {
	o.DnssecEnabled = &v
}

// GetCacheSize returns the CacheSize field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetCacheSize() int32 {
	if o == nil || o.CacheSize == nil {
		var ret int32
		return ret
	}
	return *o.CacheSize
}

// GetCacheSizeOk returns a tuple with the CacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetCacheSizeOk() (*int32, bool) {
	if o == nil || o.CacheSize == nil {
		return nil, false
	}
	return o.CacheSize, true
}

// HasCacheSize returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasCacheSize() bool {
	if o != nil && o.CacheSize != nil {
		return true
	}

	return false
}

// SetCacheSize gets a reference to the given int32 and assigns it to the CacheSize field.
func (o *DnsInfo200Response) SetCacheSize(v int32) {
	o.CacheSize = &v
}

// GetCacheTtlMin returns the CacheTtlMin field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetCacheTtlMin() int32 {
	if o == nil || o.CacheTtlMin == nil {
		var ret int32
		return ret
	}
	return *o.CacheTtlMin
}

// GetCacheTtlMinOk returns a tuple with the CacheTtlMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetCacheTtlMinOk() (*int32, bool) {
	if o == nil || o.CacheTtlMin == nil {
		return nil, false
	}
	return o.CacheTtlMin, true
}

// HasCacheTtlMin returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasCacheTtlMin() bool {
	if o != nil && o.CacheTtlMin != nil {
		return true
	}

	return false
}

// SetCacheTtlMin gets a reference to the given int32 and assigns it to the CacheTtlMin field.
func (o *DnsInfo200Response) SetCacheTtlMin(v int32) {
	o.CacheTtlMin = &v
}

// GetCacheTtlMax returns the CacheTtlMax field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetCacheTtlMax() int32 {
	if o == nil || o.CacheTtlMax == nil {
		var ret int32
		return ret
	}
	return *o.CacheTtlMax
}

// GetCacheTtlMaxOk returns a tuple with the CacheTtlMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetCacheTtlMaxOk() (*int32, bool) {
	if o == nil || o.CacheTtlMax == nil {
		return nil, false
	}
	return o.CacheTtlMax, true
}

// HasCacheTtlMax returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasCacheTtlMax() bool {
	if o != nil && o.CacheTtlMax != nil {
		return true
	}

	return false
}

// SetCacheTtlMax gets a reference to the given int32 and assigns it to the CacheTtlMax field.
func (o *DnsInfo200Response) SetCacheTtlMax(v int32) {
	o.CacheTtlMax = &v
}

// GetCacheOptimistic returns the CacheOptimistic field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetCacheOptimistic() bool {
	if o == nil || o.CacheOptimistic == nil {
		var ret bool
		return ret
	}
	return *o.CacheOptimistic
}

// GetCacheOptimisticOk returns a tuple with the CacheOptimistic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetCacheOptimisticOk() (*bool, bool) {
	if o == nil || o.CacheOptimistic == nil {
		return nil, false
	}
	return o.CacheOptimistic, true
}

// HasCacheOptimistic returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasCacheOptimistic() bool {
	if o != nil && o.CacheOptimistic != nil {
		return true
	}

	return false
}

// SetCacheOptimistic gets a reference to the given bool and assigns it to the CacheOptimistic field.
func (o *DnsInfo200Response) SetCacheOptimistic(v bool) {
	o.CacheOptimistic = &v
}

// GetUpstreamMode returns the UpstreamMode field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetUpstreamMode() string {
	if o == nil || o.UpstreamMode == nil {
		var ret string
		return ret
	}
	return *o.UpstreamMode
}

// GetUpstreamModeOk returns a tuple with the UpstreamMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetUpstreamModeOk() (*string, bool) {
	if o == nil || o.UpstreamMode == nil {
		return nil, false
	}
	return o.UpstreamMode, true
}

// HasUpstreamMode returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasUpstreamMode() bool {
	if o != nil && o.UpstreamMode != nil {
		return true
	}

	return false
}

// SetUpstreamMode gets a reference to the given string and assigns it to the UpstreamMode field.
func (o *DnsInfo200Response) SetUpstreamMode(v string) {
	o.UpstreamMode = &v
}

// GetUsePrivatePtrResolvers returns the UsePrivatePtrResolvers field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetUsePrivatePtrResolvers() bool {
	if o == nil || o.UsePrivatePtrResolvers == nil {
		var ret bool
		return ret
	}
	return *o.UsePrivatePtrResolvers
}

// GetUsePrivatePtrResolversOk returns a tuple with the UsePrivatePtrResolvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetUsePrivatePtrResolversOk() (*bool, bool) {
	if o == nil || o.UsePrivatePtrResolvers == nil {
		return nil, false
	}
	return o.UsePrivatePtrResolvers, true
}

// HasUsePrivatePtrResolvers returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasUsePrivatePtrResolvers() bool {
	if o != nil && o.UsePrivatePtrResolvers != nil {
		return true
	}

	return false
}

// SetUsePrivatePtrResolvers gets a reference to the given bool and assigns it to the UsePrivatePtrResolvers field.
func (o *DnsInfo200Response) SetUsePrivatePtrResolvers(v bool) {
	o.UsePrivatePtrResolvers = &v
}

// GetResolveClients returns the ResolveClients field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetResolveClients() bool {
	if o == nil || o.ResolveClients == nil {
		var ret bool
		return ret
	}
	return *o.ResolveClients
}

// GetResolveClientsOk returns a tuple with the ResolveClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetResolveClientsOk() (*bool, bool) {
	if o == nil || o.ResolveClients == nil {
		return nil, false
	}
	return o.ResolveClients, true
}

// HasResolveClients returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasResolveClients() bool {
	if o != nil && o.ResolveClients != nil {
		return true
	}

	return false
}

// SetResolveClients gets a reference to the given bool and assigns it to the ResolveClients field.
func (o *DnsInfo200Response) SetResolveClients(v bool) {
	o.ResolveClients = &v
}

// GetLocalPtrUpstreams returns the LocalPtrUpstreams field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetLocalPtrUpstreams() []string {
	if o == nil || o.LocalPtrUpstreams == nil {
		var ret []string
		return ret
	}
	return o.LocalPtrUpstreams
}

// GetLocalPtrUpstreamsOk returns a tuple with the LocalPtrUpstreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetLocalPtrUpstreamsOk() ([]string, bool) {
	if o == nil || o.LocalPtrUpstreams == nil {
		return nil, false
	}
	return o.LocalPtrUpstreams, true
}

// HasLocalPtrUpstreams returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasLocalPtrUpstreams() bool {
	if o != nil && o.LocalPtrUpstreams != nil {
		return true
	}

	return false
}

// SetLocalPtrUpstreams gets a reference to the given []string and assigns it to the LocalPtrUpstreams field.
func (o *DnsInfo200Response) SetLocalPtrUpstreams(v []string) {
	o.LocalPtrUpstreams = v
}

// GetDefaultLocalPtrUpstreams returns the DefaultLocalPtrUpstreams field value if set, zero value otherwise.
func (o *DnsInfo200Response) GetDefaultLocalPtrUpstreams() []string {
	if o == nil || o.DefaultLocalPtrUpstreams == nil {
		var ret []string
		return ret
	}
	return o.DefaultLocalPtrUpstreams
}

// GetDefaultLocalPtrUpstreamsOk returns a tuple with the DefaultLocalPtrUpstreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200Response) GetDefaultLocalPtrUpstreamsOk() ([]string, bool) {
	if o == nil || o.DefaultLocalPtrUpstreams == nil {
		return nil, false
	}
	return o.DefaultLocalPtrUpstreams, true
}

// HasDefaultLocalPtrUpstreams returns a boolean if a field has been set.
func (o *DnsInfo200Response) HasDefaultLocalPtrUpstreams() bool {
	if o != nil && o.DefaultLocalPtrUpstreams != nil {
		return true
	}

	return false
}

// SetDefaultLocalPtrUpstreams gets a reference to the given []string and assigns it to the DefaultLocalPtrUpstreams field.
func (o *DnsInfo200Response) SetDefaultLocalPtrUpstreams(v []string) {
	o.DefaultLocalPtrUpstreams = v
}

func (o DnsInfo200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BootstrapDns != nil {
		toSerialize["bootstrap_dns"] = o.BootstrapDns
	}
	if o.UpstreamDns != nil {
		toSerialize["upstream_dns"] = o.UpstreamDns
	}
	if o.UpstreamDnsFile != nil {
		toSerialize["upstream_dns_file"] = o.UpstreamDnsFile
	}
	if o.ProtectionEnabled != nil {
		toSerialize["protection_enabled"] = o.ProtectionEnabled
	}
	if o.DhcpAvailable != nil {
		toSerialize["dhcp_available"] = o.DhcpAvailable
	}
	if o.Ratelimit != nil {
		toSerialize["ratelimit"] = o.Ratelimit
	}
	if o.BlockingMode != nil {
		toSerialize["blocking_mode"] = o.BlockingMode
	}
	if o.BlockingIpv4 != nil {
		toSerialize["blocking_ipv4"] = o.BlockingIpv4
	}
	if o.BlockingIpv6 != nil {
		toSerialize["blocking_ipv6"] = o.BlockingIpv6
	}
	if o.EdnsCsEnabled != nil {
		toSerialize["edns_cs_enabled"] = o.EdnsCsEnabled
	}
	if o.DisableIpv6 != nil {
		toSerialize["disable_ipv6"] = o.DisableIpv6
	}
	if o.DnssecEnabled != nil {
		toSerialize["dnssec_enabled"] = o.DnssecEnabled
	}
	if o.CacheSize != nil {
		toSerialize["cache_size"] = o.CacheSize
	}
	if o.CacheTtlMin != nil {
		toSerialize["cache_ttl_min"] = o.CacheTtlMin
	}
	if o.CacheTtlMax != nil {
		toSerialize["cache_ttl_max"] = o.CacheTtlMax
	}
	if o.CacheOptimistic != nil {
		toSerialize["cache_optimistic"] = o.CacheOptimistic
	}
	if o.UpstreamMode != nil {
		toSerialize["upstream_mode"] = o.UpstreamMode
	}
	if o.UsePrivatePtrResolvers != nil {
		toSerialize["use_private_ptr_resolvers"] = o.UsePrivatePtrResolvers
	}
	if o.ResolveClients != nil {
		toSerialize["resolve_clients"] = o.ResolveClients
	}
	if o.LocalPtrUpstreams != nil {
		toSerialize["local_ptr_upstreams"] = o.LocalPtrUpstreams
	}
	if o.DefaultLocalPtrUpstreams != nil {
		toSerialize["default_local_ptr_upstreams"] = o.DefaultLocalPtrUpstreams
	}
	return json.Marshal(toSerialize)
}

type NullableDnsInfo200Response struct {
	value *DnsInfo200Response
	isSet bool
}

func (v NullableDnsInfo200Response) Get() *DnsInfo200Response {
	return v.value
}

func (v *NullableDnsInfo200Response) Set(val *DnsInfo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsInfo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsInfo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsInfo200Response(val *DnsInfo200Response) *NullableDnsInfo200Response {
	return &NullableDnsInfo200Response{value: val, isSet: true}
}

func (v NullableDnsInfo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsInfo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


