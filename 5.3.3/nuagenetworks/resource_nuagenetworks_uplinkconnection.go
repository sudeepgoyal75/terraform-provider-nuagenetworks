package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
)

func resourceUplinkConnection() *schema.Resource {
    return &schema.Resource{
        Create: resourceUplinkConnectionCreate,
        Read:   resourceUplinkConnectionRead,
        Update: resourceUplinkConnectionUpdate,
        Delete: resourceUplinkConnectionDelete,
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
            "pat_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: true,
            },
            "dns_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "dns_address_v6": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "password": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_v6": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "address_family": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "IPV4",
            },
            "address_v6": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "advertisement_criteria": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "secondary_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vlan": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "underlay_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: true,
            },
            "underlay_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "inherited": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "installer_managed": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "interface_connection_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "AUTOMATIC",
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mode": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "Dynamic",
            },
            "role": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "PRIMARY",
            },
            "role_order": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "port_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "download_rate_limit": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Default: 8.0,
            },
            "uplink_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "username": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "assoc_underlay_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_bgp_neighbor_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_underlay_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "auxiliary_link": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_vlan_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan"},
            },
            "parent_vlan": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan_template"},
            },
        },
    }
}

func resourceUplinkConnectionCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize UplinkConnection object
    o := &vspk.UplinkConnection{
    }
    if attr, ok := d.GetOk("pat_enabled"); ok {
        o.PATEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("dns_address"); ok {
        o.DNSAddress = attr.(string)
    }
    if attr, ok := d.GetOk("dns_address_v6"); ok {
        o.DNSAddressV6 = attr.(string)
    }
    if attr, ok := d.GetOk("password"); ok {
        o.Password = attr.(string)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_v6"); ok {
        o.GatewayV6 = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("address_family"); ok {
        o.AddressFamily = attr.(string)
    }
    if attr, ok := d.GetOk("address_v6"); ok {
        o.AddressV6 = attr.(string)
    }
    if attr, ok := d.GetOk("advertisement_criteria"); ok {
        o.AdvertisementCriteria = attr.(string)
    }
    if attr, ok := d.GetOk("secondary_address"); ok {
        o.SecondaryAddress = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("vlan"); ok {
        o.Vlan = attr.(int)
    }
    if attr, ok := d.GetOk("underlay_enabled"); ok {
        o.UnderlayEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("underlay_id"); ok {
        o.UnderlayID = attr.(int)
    }
    if attr, ok := d.GetOk("installer_managed"); ok {
        o.InstallerManaged = attr.(bool)
    }
    if attr, ok := d.GetOk("interface_connection_type"); ok {
        o.InterfaceConnectionType = attr.(string)
    }
    if attr, ok := d.GetOk("mode"); ok {
        o.Mode = attr.(string)
    }
    if attr, ok := d.GetOk("role"); ok {
        o.Role = attr.(string)
    }
    if attr, ok := d.GetOk("download_rate_limit"); ok {
        o.DownloadRateLimit = attr.(float64)
    }
    if attr, ok := d.GetOk("uplink_id"); ok {
        o.UplinkID = attr.(int)
    }
    if attr, ok := d.GetOk("username"); ok {
        o.Username = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_underlay_id"); ok {
        o.AssocUnderlayID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_bgp_neighbor_id"); ok {
        o.AssociatedBGPNeighborID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_underlay_name"); ok {
        o.AssociatedUnderlayName = attr.(string)
    }
    if attr, ok := d.GetOk("auxiliary_link"); ok {
        o.AuxiliaryLink = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_vlan_template"); ok {
        parent := &vspk.VLANTemplate{ID: attr.(string)}
        err := parent.CreateUplinkConnection(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vlan"); ok {
        parent := &vspk.VLAN{ID: attr.(string)}
        err := parent.CreateUplinkConnection(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceUplinkConnectionRead(d, m)
}

func resourceUplinkConnectionRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.UplinkConnection{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("pat_enabled", o.PATEnabled)
    d.Set("dns_address", o.DNSAddress)
    d.Set("dns_address_v6", o.DNSAddressV6)
    d.Set("password", o.Password)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway", o.Gateway)
    d.Set("gateway_v6", o.GatewayV6)
    d.Set("address", o.Address)
    d.Set("address_family", o.AddressFamily)
    d.Set("address_v6", o.AddressV6)
    d.Set("advertisement_criteria", o.AdvertisementCriteria)
    d.Set("secondary_address", o.SecondaryAddress)
    d.Set("netmask", o.Netmask)
    d.Set("vlan", o.Vlan)
    d.Set("underlay_enabled", o.UnderlayEnabled)
    d.Set("underlay_id", o.UnderlayID)
    d.Set("inherited", o.Inherited)
    d.Set("installer_managed", o.InstallerManaged)
    d.Set("interface_connection_type", o.InterfaceConnectionType)
    d.Set("entity_scope", o.EntityScope)
    d.Set("mode", o.Mode)
    d.Set("role", o.Role)
    d.Set("role_order", o.RoleOrder)
    d.Set("port_name", o.PortName)
    d.Set("download_rate_limit", o.DownloadRateLimit)
    d.Set("uplink_id", o.UplinkID)
    d.Set("username", o.Username)
    d.Set("assoc_underlay_id", o.AssocUnderlayID)
    d.Set("associated_bgp_neighbor_id", o.AssociatedBGPNeighborID)
    d.Set("associated_underlay_name", o.AssociatedUnderlayName)
    d.Set("auxiliary_link", o.AuxiliaryLink)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceUplinkConnectionUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.UplinkConnection{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("pat_enabled"); ok {
        o.PATEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("dns_address"); ok {
        o.DNSAddress = attr.(string)
    }
    if attr, ok := d.GetOk("dns_address_v6"); ok {
        o.DNSAddressV6 = attr.(string)
    }
    if attr, ok := d.GetOk("password"); ok {
        o.Password = attr.(string)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_v6"); ok {
        o.GatewayV6 = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("address_family"); ok {
        o.AddressFamily = attr.(string)
    }
    if attr, ok := d.GetOk("address_v6"); ok {
        o.AddressV6 = attr.(string)
    }
    if attr, ok := d.GetOk("advertisement_criteria"); ok {
        o.AdvertisementCriteria = attr.(string)
    }
    if attr, ok := d.GetOk("secondary_address"); ok {
        o.SecondaryAddress = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("vlan"); ok {
        o.Vlan = attr.(int)
    }
    if attr, ok := d.GetOk("underlay_enabled"); ok {
        o.UnderlayEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("underlay_id"); ok {
        o.UnderlayID = attr.(int)
    }
    if attr, ok := d.GetOk("installer_managed"); ok {
        o.InstallerManaged = attr.(bool)
    }
    if attr, ok := d.GetOk("interface_connection_type"); ok {
        o.InterfaceConnectionType = attr.(string)
    }
    if attr, ok := d.GetOk("mode"); ok {
        o.Mode = attr.(string)
    }
    if attr, ok := d.GetOk("role"); ok {
        o.Role = attr.(string)
    }
    if attr, ok := d.GetOk("download_rate_limit"); ok {
        o.DownloadRateLimit = attr.(float64)
    }
    if attr, ok := d.GetOk("uplink_id"); ok {
        o.UplinkID = attr.(int)
    }
    if attr, ok := d.GetOk("username"); ok {
        o.Username = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_underlay_id"); ok {
        o.AssocUnderlayID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_bgp_neighbor_id"); ok {
        o.AssociatedBGPNeighborID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_underlay_name"); ok {
        o.AssociatedUnderlayName = attr.(string)
    }
    if attr, ok := d.GetOk("auxiliary_link"); ok {
        o.AuxiliaryLink = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceUplinkConnectionDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.UplinkConnection{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}