package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.8"
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
            "dns_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "password": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "IPv4",
            },
            "advertisement_criteria": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "interface_connection_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "AUTOMATIC",
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
            "uplink_id": &schema.Schema{
                Type:     schema.TypeString,
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
            "associated_underlay_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_vsc_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "auxiliary_link": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
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
    if attr, ok := d.GetOk("dns_address"); ok {
        o.DNSAddress = attr.(string)
    }
    if attr, ok := d.GetOk("password"); ok {
        o.Password = attr.(string)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("advertisement_criteria"); ok {
        o.AdvertisementCriteria = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
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
    if attr, ok := d.GetOk("uplink_id"); ok {
        o.UplinkID = attr.(string)
    }
    if attr, ok := d.GetOk("username"); ok {
        o.Username = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_underlay_id"); ok {
        o.AssocUnderlayID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_underlay_name"); ok {
        o.AssociatedUnderlayName = attr.(string)
    }
    if attr, ok := d.GetOk("associated_vsc_profile_id"); ok {
        o.AssociatedVSCProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("auxiliary_link"); ok {
        o.AuxiliaryLink = attr.(bool)
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

    d.Set("dns_address", o.DNSAddress)
    d.Set("password", o.Password)
    d.Set("gateway", o.Gateway)
    d.Set("address", o.Address)
    d.Set("advertisement_criteria", o.AdvertisementCriteria)
    d.Set("netmask", o.Netmask)
    d.Set("interface_connection_type", o.InterfaceConnectionType)
    d.Set("mode", o.Mode)
    d.Set("role", o.Role)
    d.Set("uplink_id", o.UplinkID)
    d.Set("username", o.Username)
    d.Set("assoc_underlay_id", o.AssocUnderlayID)
    d.Set("associated_underlay_name", o.AssociatedUnderlayName)
    d.Set("associated_vsc_profile_id", o.AssociatedVSCProfileID)
    d.Set("auxiliary_link", o.AuxiliaryLink)
    
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
    
    
    if attr, ok := d.GetOk("dns_address"); ok {
        o.DNSAddress = attr.(string)
    }
    if attr, ok := d.GetOk("password"); ok {
        o.Password = attr.(string)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("advertisement_criteria"); ok {
        o.AdvertisementCriteria = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
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
    if attr, ok := d.GetOk("uplink_id"); ok {
        o.UplinkID = attr.(string)
    }
    if attr, ok := d.GetOk("username"); ok {
        o.Username = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_underlay_id"); ok {
        o.AssocUnderlayID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_underlay_name"); ok {
        o.AssociatedUnderlayName = attr.(string)
    }
    if attr, ok := d.GetOk("associated_vsc_profile_id"); ok {
        o.AssociatedVSCProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("auxiliary_link"); ok {
        o.AuxiliaryLink = attr.(bool)
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