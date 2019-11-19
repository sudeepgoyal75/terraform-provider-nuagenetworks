package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
)

func resourceBRConnection() *schema.Resource {
    return &schema.Resource{
        Create: resourceBRConnectionCreate,
        Read:   resourceBRConnectionRead,
        Update: resourceBRConnectionUpdate,
        Delete: resourceBRConnectionDelete,
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
                Default: "OPERATIONAL_LINK",
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "inherited": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mode": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "uplink_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
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

func resourceBRConnectionCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize BRConnection object
    o := &vspk.BRConnection{
    }
    if attr, ok := d.GetOk("dns_address"); ok {
        o.DNSAddress = attr.(string)
    }
    if attr, ok := d.GetOk("dns_address_v6"); ok {
        o.DNSAddressV6 = attr.(string)
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
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("mode"); ok {
        o.Mode = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_vlan_template"); ok {
        parent := &vspk.VLANTemplate{ID: attr.(string)}
        err := parent.CreateBRConnection(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vlan"); ok {
        parent := &vspk.VLAN{ID: attr.(string)}
        err := parent.CreateBRConnection(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceBRConnectionRead(d, m)
}

func resourceBRConnectionRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.BRConnection{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("dns_address", o.DNSAddress)
    d.Set("dns_address_v6", o.DNSAddressV6)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway", o.Gateway)
    d.Set("gateway_v6", o.GatewayV6)
    d.Set("address", o.Address)
    d.Set("address_family", o.AddressFamily)
    d.Set("address_v6", o.AddressV6)
    d.Set("advertisement_criteria", o.AdvertisementCriteria)
    d.Set("netmask", o.Netmask)
    d.Set("inherited", o.Inherited)
    d.Set("entity_scope", o.EntityScope)
    d.Set("mode", o.Mode)
    d.Set("uplink_id", o.UplinkID)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceBRConnectionUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.BRConnection{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("dns_address"); ok {
        o.DNSAddress = attr.(string)
    }
    if attr, ok := d.GetOk("dns_address_v6"); ok {
        o.DNSAddressV6 = attr.(string)
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
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("mode"); ok {
        o.Mode = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceBRConnectionDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.BRConnection{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}