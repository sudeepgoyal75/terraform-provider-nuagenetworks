package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceEnterprise() *schema.Resource {
    return &schema.Resource{
        Create: resourceEnterpriseCreate,
        Read:   resourceEnterpriseRead,
        Update: resourceEnterpriseUpdate,
        Delete: resourceEnterpriseDelete,
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
            "ldap_authorization_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "ldap_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "bgp_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "dhcp_lease_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "vnf_management_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "web_filter_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "receive_multi_cast_list_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "send_multi_cast_list_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "shared_enterprise": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "dictionary_version": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default: 2,
            },
            "virtual_firewall_rules_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "allow_advanced_qos_configuration": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "allow_gateway_management": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "allow_trusted_forwarding_class": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "allowed_forwarding_classes": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "allowed_forwarding_mode": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "floating_ips_quota": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "floating_ips_used": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "flow_collection_enabled": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "DISABLED",
            },
            "enable_application_performance_management": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "encryption_management_mode": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "enterprise_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "local_as": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "use_global_mac": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "associated_enterprise_security_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_group_key_encryption_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_key_server_monitor_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "customer_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "avatar_data": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "avatar_type": &schema.Schema{
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

func resourceEnterpriseCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Enterprise object
    o := &vspk.Enterprise{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("dhcp_lease_interval"); ok {
        o.DHCPLeaseInterval = attr.(int)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("dictionary_version"); ok {
        o.DictionaryVersion = attr.(int)
    }
    if attr, ok := d.GetOk("allow_advanced_qos_configuration"); ok {
        o.AllowAdvancedQOSConfiguration = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_gateway_management"); ok {
        o.AllowGatewayManagement = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_trusted_forwarding_class"); ok {
        o.AllowTrustedForwardingClass = attr.(bool)
    }
    if attr, ok := d.GetOk("allowed_forwarding_classes"); ok {
        o.AllowedForwardingClasses = attr.([]interface{})
    }
    if attr, ok := d.GetOk("floating_ips_quota"); ok {
        o.FloatingIPsQuota = attr.(int)
    }
    if attr, ok := d.GetOk("floating_ips_used"); ok {
        o.FloatingIPsUsed = attr.(int)
    }
    if attr, ok := d.GetOk("flow_collection_enabled"); ok {
        o.FlowCollectionEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("enable_application_performance_management"); ok {
        o.EnableApplicationPerformanceManagement = attr.(bool)
    }
    if attr, ok := d.GetOk("enterprise_profile_id"); ok {
        o.EnterpriseProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("local_as"); ok {
        o.LocalAS = attr.(int)
    }
    if attr, ok := d.GetOk("use_global_mac"); ok {
        o.UseGlobalMAC = attr.(bool)
    }
    if attr, ok := d.GetOk("customer_id"); ok {
        o.CustomerID = attr.(int)
    }
    if attr, ok := d.GetOk("avatar_data"); ok {
        o.AvatarData = attr.(string)
    }
    if attr, ok := d.GetOk("avatar_type"); ok {
        o.AvatarType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateEnterprise(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceEnterpriseRead(d, m)
}

func resourceEnterpriseRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Enterprise{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("ldap_authorization_enabled", o.LDAPAuthorizationEnabled)
    d.Set("ldap_enabled", o.LDAPEnabled)
    d.Set("bgp_enabled", o.BGPEnabled)
    d.Set("dhcp_lease_interval", o.DHCPLeaseInterval)
    d.Set("vnf_management_enabled", o.VNFManagementEnabled)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("web_filter_enabled", o.WebFilterEnabled)
    d.Set("receive_multi_cast_list_id", o.ReceiveMultiCastListID)
    d.Set("send_multi_cast_list_id", o.SendMultiCastListID)
    d.Set("description", o.Description)
    d.Set("shared_enterprise", o.SharedEnterprise)
    d.Set("dictionary_version", o.DictionaryVersion)
    d.Set("virtual_firewall_rules_enabled", o.VirtualFirewallRulesEnabled)
    d.Set("allow_advanced_qos_configuration", o.AllowAdvancedQOSConfiguration)
    d.Set("allow_gateway_management", o.AllowGatewayManagement)
    d.Set("allow_trusted_forwarding_class", o.AllowTrustedForwardingClass)
    d.Set("allowed_forwarding_classes", o.AllowedForwardingClasses)
    d.Set("allowed_forwarding_mode", o.AllowedForwardingMode)
    d.Set("floating_ips_quota", o.FloatingIPsQuota)
    d.Set("floating_ips_used", o.FloatingIPsUsed)
    d.Set("flow_collection_enabled", o.FlowCollectionEnabled)
    d.Set("enable_application_performance_management", o.EnableApplicationPerformanceManagement)
    d.Set("encryption_management_mode", o.EncryptionManagementMode)
    d.Set("enterprise_profile_id", o.EnterpriseProfileID)
    d.Set("entity_scope", o.EntityScope)
    d.Set("local_as", o.LocalAS)
    d.Set("use_global_mac", o.UseGlobalMAC)
    d.Set("associated_enterprise_security_id", o.AssociatedEnterpriseSecurityID)
    d.Set("associated_group_key_encryption_profile_id", o.AssociatedGroupKeyEncryptionProfileID)
    d.Set("associated_key_server_monitor_id", o.AssociatedKeyServerMonitorID)
    d.Set("customer_id", o.CustomerID)
    d.Set("avatar_data", o.AvatarData)
    d.Set("avatar_type", o.AvatarType)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceEnterpriseUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Enterprise{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("dhcp_lease_interval"); ok {
        o.DHCPLeaseInterval = attr.(int)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("dictionary_version"); ok {
        o.DictionaryVersion = attr.(int)
    }
    if attr, ok := d.GetOk("allow_advanced_qos_configuration"); ok {
        o.AllowAdvancedQOSConfiguration = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_gateway_management"); ok {
        o.AllowGatewayManagement = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_trusted_forwarding_class"); ok {
        o.AllowTrustedForwardingClass = attr.(bool)
    }
    if attr, ok := d.GetOk("allowed_forwarding_classes"); ok {
        o.AllowedForwardingClasses = attr.([]interface{})
    }
    if attr, ok := d.GetOk("floating_ips_quota"); ok {
        o.FloatingIPsQuota = attr.(int)
    }
    if attr, ok := d.GetOk("floating_ips_used"); ok {
        o.FloatingIPsUsed = attr.(int)
    }
    if attr, ok := d.GetOk("flow_collection_enabled"); ok {
        o.FlowCollectionEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("enable_application_performance_management"); ok {
        o.EnableApplicationPerformanceManagement = attr.(bool)
    }
    if attr, ok := d.GetOk("enterprise_profile_id"); ok {
        o.EnterpriseProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("local_as"); ok {
        o.LocalAS = attr.(int)
    }
    if attr, ok := d.GetOk("use_global_mac"); ok {
        o.UseGlobalMAC = attr.(bool)
    }
    if attr, ok := d.GetOk("customer_id"); ok {
        o.CustomerID = attr.(int)
    }
    if attr, ok := d.GetOk("avatar_data"); ok {
        o.AvatarData = attr.(string)
    }
    if attr, ok := d.GetOk("avatar_type"); ok {
        o.AvatarType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceEnterpriseDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Enterprise{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}