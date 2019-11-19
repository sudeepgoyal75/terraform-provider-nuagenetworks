package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11"
)

func resourceVCenter() *schema.Resource {
    return &schema.Resource{
        Create: resourceVCenterCreate,
        Read:   resourceVCenterRead,
        Update: resourceVCenterUpdate,
        Delete: resourceVCenterDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
        Schema: map[string]*schema.Schema{
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vrs_configuration_time_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "v_require_nuage_metadata": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "password": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "data_dns1": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "data_dns2": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "data_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "data_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "datapath_sync_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "secondary_nuage_controller": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "generic_split_activation": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "separate_data_network": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "destination_mirror_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "no_mirror",
            },
            "metadata_server_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "metadata_server_listen_port": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "metadata_server_port": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "metadata_service_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "network_uplink_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "network_uplink_interface_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "network_uplink_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "network_uplink_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nfs_log_server": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nfs_mount_path": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mgmt_dns1": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mgmt_dns2": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mgmt_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mgmt_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "dhcp_relay_server": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mirror_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "disable_network_discovery": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "site_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "old_agency_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "allow_data_dhcp": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "allow_mgmt_dhcp": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "flow_eviction_threshold": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "vm_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "connection_status": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "portgroup_metadata": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "host_level_management": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "nova_client_version": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_service_auth_url": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_service_endpoint": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_service_password": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_service_tenant": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_service_username": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_shared_secret": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_region_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "upgrade_package_password": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "upgrade_package_url": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "upgrade_package_username": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "upgrade_script_time_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "primary_nuage_controller": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vrs_config_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vrs_password": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vrs_user_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "user_name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "static_route": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "static_route_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "static_route_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ntp_server1": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ntp_server2": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "http_port": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "https_port": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "mtu": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "multi_vmssupport": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "multicast_receive_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_receive_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_receive_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_receive_range": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_send_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_send_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_send_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_source_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "customized_script_url": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "auto_resolve_frequency": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "ovf_url": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
        },
    }
}

func resourceVCenterCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VCenter object
    o := &vspk.VCenter{
        Name: d.Get("name").(string),
        Password: d.Get("password").(string),
        IpAddress: d.Get("ip_address").(string),
        UserName: d.Get("user_name").(string),
    }
    if attr, ok := d.GetOk("vrs_configuration_time_limit"); ok {
        o.VRSConfigurationTimeLimit = attr.(int)
    }
    if attr, ok := d.GetOk("v_require_nuage_metadata"); ok {
        o.VRequireNuageMetadata = attr.(bool)
    }
    if attr, ok := d.GetOk("data_dns1"); ok {
        o.DataDNS1 = attr.(string)
    }
    if attr, ok := d.GetOk("data_dns2"); ok {
        o.DataDNS2 = attr.(string)
    }
    if attr, ok := d.GetOk("data_gateway"); ok {
        o.DataGateway = attr.(string)
    }
    if attr, ok := d.GetOk("data_network_portgroup"); ok {
        o.DataNetworkPortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("datapath_sync_timeout"); ok {
        o.DatapathSyncTimeout = attr.(int)
    }
    if attr, ok := d.GetOk("secondary_nuage_controller"); ok {
        o.SecondaryNuageController = attr.(string)
    }
    if attr, ok := d.GetOk("generic_split_activation"); ok {
        o.GenericSplitActivation = attr.(bool)
    }
    if attr, ok := d.GetOk("separate_data_network"); ok {
        o.SeparateDataNetwork = attr.(bool)
    }
    if attr, ok := d.GetOk("personality"); ok {
        o.Personality = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_mirror_port"); ok {
        o.DestinationMirrorPort = attr.(string)
    }
    if attr, ok := d.GetOk("metadata_server_ip"); ok {
        o.MetadataServerIP = attr.(string)
    }
    if attr, ok := d.GetOk("metadata_server_listen_port"); ok {
        o.MetadataServerListenPort = attr.(int)
    }
    if attr, ok := d.GetOk("metadata_server_port"); ok {
        o.MetadataServerPort = attr.(int)
    }
    if attr, ok := d.GetOk("metadata_service_enabled"); ok {
        o.MetadataServiceEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("network_uplink_interface"); ok {
        o.NetworkUplinkInterface = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_gateway"); ok {
        o.NetworkUplinkInterfaceGateway = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_ip"); ok {
        o.NetworkUplinkInterfaceIp = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_netmask"); ok {
        o.NetworkUplinkInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("nfs_log_server"); ok {
        o.NfsLogServer = attr.(string)
    }
    if attr, ok := d.GetOk("nfs_mount_path"); ok {
        o.NfsMountPath = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_dns1"); ok {
        o.MgmtDNS1 = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_dns2"); ok {
        o.MgmtDNS2 = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_gateway"); ok {
        o.MgmtGateway = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_network_portgroup"); ok {
        o.MgmtNetworkPortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("dhcp_relay_server"); ok {
        o.DhcpRelayServer = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_network_portgroup"); ok {
        o.MirrorNetworkPortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("disable_network_discovery"); ok {
        o.DisableNetworkDiscovery = attr.(bool)
    }
    if attr, ok := d.GetOk("site_id"); ok {
        o.SiteId = attr.(string)
    }
    if attr, ok := d.GetOk("old_agency_name"); ok {
        o.OldAgencyName = attr.(string)
    }
    if attr, ok := d.GetOk("allow_data_dhcp"); ok {
        o.AllowDataDHCP = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_mgmt_dhcp"); ok {
        o.AllowMgmtDHCP = attr.(bool)
    }
    if attr, ok := d.GetOk("flow_eviction_threshold"); ok {
        o.FlowEvictionThreshold = attr.(int)
    }
    if attr, ok := d.GetOk("vm_network_portgroup"); ok {
        o.VmNetworkPortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("connection_status"); ok {
        o.ConnectionStatus = attr.(bool)
    }
    if attr, ok := d.GetOk("portgroup_metadata"); ok {
        o.PortgroupMetadata = attr.(bool)
    }
    if attr, ok := d.GetOk("host_level_management"); ok {
        o.HostLevelManagement = attr.(bool)
    }
    if attr, ok := d.GetOk("nova_client_version"); ok {
        o.NovaClientVersion = attr.(int)
    }
    if attr, ok := d.GetOk("nova_metadata_service_auth_url"); ok {
        o.NovaMetadataServiceAuthUrl = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_endpoint"); ok {
        o.NovaMetadataServiceEndpoint = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_password"); ok {
        o.NovaMetadataServicePassword = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_tenant"); ok {
        o.NovaMetadataServiceTenant = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_username"); ok {
        o.NovaMetadataServiceUsername = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_shared_secret"); ok {
        o.NovaMetadataSharedSecret = attr.(string)
    }
    if attr, ok := d.GetOk("nova_region_name"); ok {
        o.NovaRegionName = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_password"); ok {
        o.UpgradePackagePassword = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_url"); ok {
        o.UpgradePackageURL = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_username"); ok {
        o.UpgradePackageUsername = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_script_time_limit"); ok {
        o.UpgradeScriptTimeLimit = attr.(int)
    }
    if attr, ok := d.GetOk("primary_nuage_controller"); ok {
        o.PrimaryNuageController = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_config_id"); ok {
        o.VrsConfigID = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_password"); ok {
        o.VrsPassword = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_user_name"); ok {
        o.VrsUserName = attr.(string)
    }
    if attr, ok := d.GetOk("static_route"); ok {
        o.StaticRoute = attr.(string)
    }
    if attr, ok := d.GetOk("static_route_gateway"); ok {
        o.StaticRouteGateway = attr.(string)
    }
    if attr, ok := d.GetOk("static_route_netmask"); ok {
        o.StaticRouteNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("ntp_server1"); ok {
        o.NtpServer1 = attr.(string)
    }
    if attr, ok := d.GetOk("ntp_server2"); ok {
        o.NtpServer2 = attr.(string)
    }
    if attr, ok := d.GetOk("http_port"); ok {
        o.HttpPort = attr.(int)
    }
    if attr, ok := d.GetOk("https_port"); ok {
        o.HttpsPort = attr.(int)
    }
    if attr, ok := d.GetOk("mtu"); ok {
        o.Mtu = attr.(int)
    }
    if attr, ok := d.GetOk("multi_vmssupport"); ok {
        o.MultiVMSsupport = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast_receive_interface"); ok {
        o.MulticastReceiveInterface = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_interface_ip"); ok {
        o.MulticastReceiveInterfaceIP = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_interface_netmask"); ok {
        o.MulticastReceiveInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_range"); ok {
        o.MulticastReceiveRange = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface"); ok {
        o.MulticastSendInterface = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface_ip"); ok {
        o.MulticastSendInterfaceIP = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface_netmask"); ok {
        o.MulticastSendInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_source_portgroup"); ok {
        o.MulticastSourcePortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("customized_script_url"); ok {
        o.CustomizedScriptURL = attr.(string)
    }
    if attr, ok := d.GetOk("auto_resolve_frequency"); ok {
        o.AutoResolveFrequency = attr.(int)
    }
    if attr, ok := d.GetOk("ovf_url"); ok {
        o.OvfURL = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateVCenter(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceVCenterRead(d, m)
}

func resourceVCenterRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VCenter{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("vrs_configuration_time_limit", o.VRSConfigurationTimeLimit)
    d.Set("v_require_nuage_metadata", o.VRequireNuageMetadata)
    d.Set("name", o.Name)
    d.Set("password", o.Password)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("data_dns1", o.DataDNS1)
    d.Set("data_dns2", o.DataDNS2)
    d.Set("data_gateway", o.DataGateway)
    d.Set("data_network_portgroup", o.DataNetworkPortgroup)
    d.Set("datapath_sync_timeout", o.DatapathSyncTimeout)
    d.Set("secondary_nuage_controller", o.SecondaryNuageController)
    d.Set("generic_split_activation", o.GenericSplitActivation)
    d.Set("separate_data_network", o.SeparateDataNetwork)
    d.Set("personality", o.Personality)
    d.Set("description", o.Description)
    d.Set("destination_mirror_port", o.DestinationMirrorPort)
    d.Set("metadata_server_ip", o.MetadataServerIP)
    d.Set("metadata_server_listen_port", o.MetadataServerListenPort)
    d.Set("metadata_server_port", o.MetadataServerPort)
    d.Set("metadata_service_enabled", o.MetadataServiceEnabled)
    d.Set("network_uplink_interface", o.NetworkUplinkInterface)
    d.Set("network_uplink_interface_gateway", o.NetworkUplinkInterfaceGateway)
    d.Set("network_uplink_interface_ip", o.NetworkUplinkInterfaceIp)
    d.Set("network_uplink_interface_netmask", o.NetworkUplinkInterfaceNetmask)
    d.Set("nfs_log_server", o.NfsLogServer)
    d.Set("nfs_mount_path", o.NfsMountPath)
    d.Set("mgmt_dns1", o.MgmtDNS1)
    d.Set("mgmt_dns2", o.MgmtDNS2)
    d.Set("mgmt_gateway", o.MgmtGateway)
    d.Set("mgmt_network_portgroup", o.MgmtNetworkPortgroup)
    d.Set("dhcp_relay_server", o.DhcpRelayServer)
    d.Set("mirror_network_portgroup", o.MirrorNetworkPortgroup)
    d.Set("disable_network_discovery", o.DisableNetworkDiscovery)
    d.Set("site_id", o.SiteId)
    d.Set("old_agency_name", o.OldAgencyName)
    d.Set("allow_data_dhcp", o.AllowDataDHCP)
    d.Set("allow_mgmt_dhcp", o.AllowMgmtDHCP)
    d.Set("flow_eviction_threshold", o.FlowEvictionThreshold)
    d.Set("vm_network_portgroup", o.VmNetworkPortgroup)
    d.Set("entity_scope", o.EntityScope)
    d.Set("connection_status", o.ConnectionStatus)
    d.Set("portgroup_metadata", o.PortgroupMetadata)
    d.Set("host_level_management", o.HostLevelManagement)
    d.Set("nova_client_version", o.NovaClientVersion)
    d.Set("nova_metadata_service_auth_url", o.NovaMetadataServiceAuthUrl)
    d.Set("nova_metadata_service_endpoint", o.NovaMetadataServiceEndpoint)
    d.Set("nova_metadata_service_password", o.NovaMetadataServicePassword)
    d.Set("nova_metadata_service_tenant", o.NovaMetadataServiceTenant)
    d.Set("nova_metadata_service_username", o.NovaMetadataServiceUsername)
    d.Set("nova_metadata_shared_secret", o.NovaMetadataSharedSecret)
    d.Set("nova_region_name", o.NovaRegionName)
    d.Set("ip_address", o.IpAddress)
    d.Set("upgrade_package_password", o.UpgradePackagePassword)
    d.Set("upgrade_package_url", o.UpgradePackageURL)
    d.Set("upgrade_package_username", o.UpgradePackageUsername)
    d.Set("upgrade_script_time_limit", o.UpgradeScriptTimeLimit)
    d.Set("primary_nuage_controller", o.PrimaryNuageController)
    d.Set("vrs_config_id", o.VrsConfigID)
    d.Set("vrs_password", o.VrsPassword)
    d.Set("vrs_user_name", o.VrsUserName)
    d.Set("user_name", o.UserName)
    d.Set("static_route", o.StaticRoute)
    d.Set("static_route_gateway", o.StaticRouteGateway)
    d.Set("static_route_netmask", o.StaticRouteNetmask)
    d.Set("ntp_server1", o.NtpServer1)
    d.Set("ntp_server2", o.NtpServer2)
    d.Set("http_port", o.HttpPort)
    d.Set("https_port", o.HttpsPort)
    d.Set("mtu", o.Mtu)
    d.Set("multi_vmssupport", o.MultiVMSsupport)
    d.Set("multicast_receive_interface", o.MulticastReceiveInterface)
    d.Set("multicast_receive_interface_ip", o.MulticastReceiveInterfaceIP)
    d.Set("multicast_receive_interface_netmask", o.MulticastReceiveInterfaceNetmask)
    d.Set("multicast_receive_range", o.MulticastReceiveRange)
    d.Set("multicast_send_interface", o.MulticastSendInterface)
    d.Set("multicast_send_interface_ip", o.MulticastSendInterfaceIP)
    d.Set("multicast_send_interface_netmask", o.MulticastSendInterfaceNetmask)
    d.Set("multicast_source_portgroup", o.MulticastSourcePortgroup)
    d.Set("customized_script_url", o.CustomizedScriptURL)
    d.Set("auto_resolve_frequency", o.AutoResolveFrequency)
    d.Set("ovf_url", o.OvfURL)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVCenterUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VCenter{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.Password = d.Get("password").(string)
    o.IpAddress = d.Get("ip_address").(string)
    o.UserName = d.Get("user_name").(string)
    
    if attr, ok := d.GetOk("vrs_configuration_time_limit"); ok {
        o.VRSConfigurationTimeLimit = attr.(int)
    }
    if attr, ok := d.GetOk("v_require_nuage_metadata"); ok {
        o.VRequireNuageMetadata = attr.(bool)
    }
    if attr, ok := d.GetOk("data_dns1"); ok {
        o.DataDNS1 = attr.(string)
    }
    if attr, ok := d.GetOk("data_dns2"); ok {
        o.DataDNS2 = attr.(string)
    }
    if attr, ok := d.GetOk("data_gateway"); ok {
        o.DataGateway = attr.(string)
    }
    if attr, ok := d.GetOk("data_network_portgroup"); ok {
        o.DataNetworkPortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("datapath_sync_timeout"); ok {
        o.DatapathSyncTimeout = attr.(int)
    }
    if attr, ok := d.GetOk("secondary_nuage_controller"); ok {
        o.SecondaryNuageController = attr.(string)
    }
    if attr, ok := d.GetOk("generic_split_activation"); ok {
        o.GenericSplitActivation = attr.(bool)
    }
    if attr, ok := d.GetOk("separate_data_network"); ok {
        o.SeparateDataNetwork = attr.(bool)
    }
    if attr, ok := d.GetOk("personality"); ok {
        o.Personality = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_mirror_port"); ok {
        o.DestinationMirrorPort = attr.(string)
    }
    if attr, ok := d.GetOk("metadata_server_ip"); ok {
        o.MetadataServerIP = attr.(string)
    }
    if attr, ok := d.GetOk("metadata_server_listen_port"); ok {
        o.MetadataServerListenPort = attr.(int)
    }
    if attr, ok := d.GetOk("metadata_server_port"); ok {
        o.MetadataServerPort = attr.(int)
    }
    if attr, ok := d.GetOk("metadata_service_enabled"); ok {
        o.MetadataServiceEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("network_uplink_interface"); ok {
        o.NetworkUplinkInterface = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_gateway"); ok {
        o.NetworkUplinkInterfaceGateway = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_ip"); ok {
        o.NetworkUplinkInterfaceIp = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_netmask"); ok {
        o.NetworkUplinkInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("nfs_log_server"); ok {
        o.NfsLogServer = attr.(string)
    }
    if attr, ok := d.GetOk("nfs_mount_path"); ok {
        o.NfsMountPath = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_dns1"); ok {
        o.MgmtDNS1 = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_dns2"); ok {
        o.MgmtDNS2 = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_gateway"); ok {
        o.MgmtGateway = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_network_portgroup"); ok {
        o.MgmtNetworkPortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("dhcp_relay_server"); ok {
        o.DhcpRelayServer = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_network_portgroup"); ok {
        o.MirrorNetworkPortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("disable_network_discovery"); ok {
        o.DisableNetworkDiscovery = attr.(bool)
    }
    if attr, ok := d.GetOk("site_id"); ok {
        o.SiteId = attr.(string)
    }
    if attr, ok := d.GetOk("old_agency_name"); ok {
        o.OldAgencyName = attr.(string)
    }
    if attr, ok := d.GetOk("allow_data_dhcp"); ok {
        o.AllowDataDHCP = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_mgmt_dhcp"); ok {
        o.AllowMgmtDHCP = attr.(bool)
    }
    if attr, ok := d.GetOk("flow_eviction_threshold"); ok {
        o.FlowEvictionThreshold = attr.(int)
    }
    if attr, ok := d.GetOk("vm_network_portgroup"); ok {
        o.VmNetworkPortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("connection_status"); ok {
        o.ConnectionStatus = attr.(bool)
    }
    if attr, ok := d.GetOk("portgroup_metadata"); ok {
        o.PortgroupMetadata = attr.(bool)
    }
    if attr, ok := d.GetOk("host_level_management"); ok {
        o.HostLevelManagement = attr.(bool)
    }
    if attr, ok := d.GetOk("nova_client_version"); ok {
        o.NovaClientVersion = attr.(int)
    }
    if attr, ok := d.GetOk("nova_metadata_service_auth_url"); ok {
        o.NovaMetadataServiceAuthUrl = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_endpoint"); ok {
        o.NovaMetadataServiceEndpoint = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_password"); ok {
        o.NovaMetadataServicePassword = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_tenant"); ok {
        o.NovaMetadataServiceTenant = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_username"); ok {
        o.NovaMetadataServiceUsername = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_shared_secret"); ok {
        o.NovaMetadataSharedSecret = attr.(string)
    }
    if attr, ok := d.GetOk("nova_region_name"); ok {
        o.NovaRegionName = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_password"); ok {
        o.UpgradePackagePassword = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_url"); ok {
        o.UpgradePackageURL = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_username"); ok {
        o.UpgradePackageUsername = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_script_time_limit"); ok {
        o.UpgradeScriptTimeLimit = attr.(int)
    }
    if attr, ok := d.GetOk("primary_nuage_controller"); ok {
        o.PrimaryNuageController = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_config_id"); ok {
        o.VrsConfigID = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_password"); ok {
        o.VrsPassword = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_user_name"); ok {
        o.VrsUserName = attr.(string)
    }
    if attr, ok := d.GetOk("static_route"); ok {
        o.StaticRoute = attr.(string)
    }
    if attr, ok := d.GetOk("static_route_gateway"); ok {
        o.StaticRouteGateway = attr.(string)
    }
    if attr, ok := d.GetOk("static_route_netmask"); ok {
        o.StaticRouteNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("ntp_server1"); ok {
        o.NtpServer1 = attr.(string)
    }
    if attr, ok := d.GetOk("ntp_server2"); ok {
        o.NtpServer2 = attr.(string)
    }
    if attr, ok := d.GetOk("http_port"); ok {
        o.HttpPort = attr.(int)
    }
    if attr, ok := d.GetOk("https_port"); ok {
        o.HttpsPort = attr.(int)
    }
    if attr, ok := d.GetOk("mtu"); ok {
        o.Mtu = attr.(int)
    }
    if attr, ok := d.GetOk("multi_vmssupport"); ok {
        o.MultiVMSsupport = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast_receive_interface"); ok {
        o.MulticastReceiveInterface = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_interface_ip"); ok {
        o.MulticastReceiveInterfaceIP = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_interface_netmask"); ok {
        o.MulticastReceiveInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_range"); ok {
        o.MulticastReceiveRange = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface"); ok {
        o.MulticastSendInterface = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface_ip"); ok {
        o.MulticastSendInterfaceIP = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface_netmask"); ok {
        o.MulticastSendInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_source_portgroup"); ok {
        o.MulticastSourcePortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("customized_script_url"); ok {
        o.CustomizedScriptURL = attr.(string)
    }
    if attr, ok := d.GetOk("auto_resolve_frequency"); ok {
        o.AutoResolveFrequency = attr.(int)
    }
    if attr, ok := d.GetOk("ovf_url"); ok {
        o.OvfURL = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVCenterDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VCenter{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}