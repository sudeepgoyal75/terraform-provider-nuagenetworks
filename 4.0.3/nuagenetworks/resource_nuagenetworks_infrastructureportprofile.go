package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.3"
)

func resourceInfrastructurePortProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceInfrastructurePortProfileCreate,
        Read:   resourceInfrastructurePortProfileRead,
        Update: resourceInfrastructurePortProfileUpdate,
        Delete: resourceInfrastructurePortProfileDelete,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "speed": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "uplink_tag": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mtu": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "duplex": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourceInfrastructurePortProfileCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize InfrastructurePortProfile object
    o := &vspk.InfrastructurePortProfile{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("enterprise_id"); ok {
        o.EnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("speed"); ok {
        o.Speed = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_tag"); ok {
        o.UplinkTag = attr.(string)
    }
    if attr, ok := d.GetOk("mtu"); ok {
        o.Mtu = attr.(int)
    }
    if attr, ok := d.GetOk("duplex"); ok {
        o.Duplex = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_me"); ok {
        parent := &vspk.Me{ID: attr.(string)}
        err := parent.CreateInfrastructurePortProfile(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        err := parent.CreateInfrastructurePortProfile(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceInfrastructurePortProfileRead(d, m)
}

func resourceInfrastructurePortProfileRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.InfrastructurePortProfile{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("description", o.Description)
    d.Set("enterprise_id", o.EnterpriseID)
    d.Set("entity_scope", o.EntityScope)
    d.Set("speed", o.Speed)
    d.Set("uplink_tag", o.UplinkTag)
    d.Set("mtu", o.Mtu)
    d.Set("duplex", o.Duplex)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceInfrastructurePortProfileUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.InfrastructurePortProfile{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("enterprise_id"); ok {
        o.EnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("speed"); ok {
        o.Speed = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_tag"); ok {
        o.UplinkTag = attr.(string)
    }
    if attr, ok := d.GetOk("mtu"); ok {
        o.Mtu = attr.(int)
    }
    if attr, ok := d.GetOk("duplex"); ok {
        o.Duplex = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceInfrastructurePortProfileDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.InfrastructurePortProfile{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}