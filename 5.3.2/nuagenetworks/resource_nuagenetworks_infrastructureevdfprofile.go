package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
)

func resourceInfrastructureEVDFProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceInfrastructureEVDFProfileCreate,
        Read:   resourceInfrastructureEVDFProfileRead,
        Update: resourceInfrastructureEVDFProfileUpdate,
        Delete: resourceInfrastructureEVDFProfileDelete,
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
            "ntp_server_key": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ntp_server_key_id": &schema.Schema{
                Type:     schema.TypeInt,
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
            "active_controller": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "service_ipv4_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "0.0.0.0/8",
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "proxy_dns_name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "use_two_factor": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "standby_controller": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nuage_platform": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "KVM",
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
        },
    }
}

func resourceInfrastructureEVDFProfileCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize InfrastructureEVDFProfile object
    o := &vspk.InfrastructureEVDFProfile{
        Name: d.Get("name").(string),
        ActiveController: d.Get("active_controller").(string),
        ProxyDNSName: d.Get("proxy_dns_name").(string),
    }
    if attr, ok := d.GetOk("ntp_server_key"); ok {
        o.NTPServerKey = attr.(string)
    }
    if attr, ok := d.GetOk("ntp_server_key_id"); ok {
        o.NTPServerKeyID = attr.(int)
    }
    if attr, ok := d.GetOk("service_ipv4_subnet"); ok {
        o.ServiceIPv4Subnet = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("use_two_factor"); ok {
        o.UseTwoFactor = attr.(bool)
    }
    if attr, ok := d.GetOk("standby_controller"); ok {
        o.StandbyController = attr.(string)
    }
    if attr, ok := d.GetOk("nuage_platform"); ok {
        o.NuagePlatform = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateInfrastructureEVDFProfile(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceInfrastructureEVDFProfileRead(d, m)
}

func resourceInfrastructureEVDFProfileRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.InfrastructureEVDFProfile{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("ntp_server_key", o.NTPServerKey)
    d.Set("ntp_server_key_id", o.NTPServerKeyID)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("active_controller", o.ActiveController)
    d.Set("service_ipv4_subnet", o.ServiceIPv4Subnet)
    d.Set("description", o.Description)
    d.Set("entity_scope", o.EntityScope)
    d.Set("proxy_dns_name", o.ProxyDNSName)
    d.Set("use_two_factor", o.UseTwoFactor)
    d.Set("standby_controller", o.StandbyController)
    d.Set("nuage_platform", o.NuagePlatform)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceInfrastructureEVDFProfileUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.InfrastructureEVDFProfile{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.ActiveController = d.Get("active_controller").(string)
    o.ProxyDNSName = d.Get("proxy_dns_name").(string)
    
    if attr, ok := d.GetOk("ntp_server_key"); ok {
        o.NTPServerKey = attr.(string)
    }
    if attr, ok := d.GetOk("ntp_server_key_id"); ok {
        o.NTPServerKeyID = attr.(int)
    }
    if attr, ok := d.GetOk("service_ipv4_subnet"); ok {
        o.ServiceIPv4Subnet = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("use_two_factor"); ok {
        o.UseTwoFactor = attr.(bool)
    }
    if attr, ok := d.GetOk("standby_controller"); ok {
        o.StandbyController = attr.(string)
    }
    if attr, ok := d.GetOk("nuage_platform"); ok {
        o.NuagePlatform = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceInfrastructureEVDFProfileDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.InfrastructureEVDFProfile{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}