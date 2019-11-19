package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.9"
)

func resourcePATNATPool() *schema.Resource {
    return &schema.Resource{
        Create: resourcePATNATPoolCreate,
        Read:   resourcePATNATPoolRead,
        Update: resourcePATNATPoolUpdate,
        Delete: resourcePATNATPoolDelete,
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
            "address_range": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "default_patip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
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
            "associated_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_gateway_type": &schema.Schema{
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

func resourcePATNATPoolCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize PATNATPool object
    o := &vspk.PATNATPool{
        Name: d.Get("name").(string),
        AddressRange: d.Get("address_range").(string),
    }
    if attr, ok := d.GetOk("default_patip"); ok {
        o.DefaultPATIP = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_gateway_id"); ok {
        o.AssociatedGatewayId = attr.(string)
    }
    if attr, ok := d.GetOk("associated_gateway_type"); ok {
        o.AssociatedGatewayType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := m.(*vspk.Me)
    err := parent.CreatePATNATPool(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourcePATNATPoolRead(d, m)
}

func resourcePATNATPoolRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PATNATPool{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("address_range", o.AddressRange)
    d.Set("default_patip", o.DefaultPATIP)
    d.Set("permitted_action", o.PermittedAction)
    d.Set("description", o.Description)
    d.Set("entity_scope", o.EntityScope)
    d.Set("associated_gateway_id", o.AssociatedGatewayId)
    d.Set("associated_gateway_type", o.AssociatedGatewayType)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourcePATNATPoolUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PATNATPool{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.AddressRange = d.Get("address_range").(string)
    
    if attr, ok := d.GetOk("default_patip"); ok {
        o.DefaultPATIP = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_gateway_id"); ok {
        o.AssociatedGatewayId = attr.(string)
    }
    if attr, ok := d.GetOk("associated_gateway_type"); ok {
        o.AssociatedGatewayType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourcePATNATPoolDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PATNATPool{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}