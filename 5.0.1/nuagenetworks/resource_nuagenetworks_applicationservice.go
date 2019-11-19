package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.0.1"
)

func resourceApplicationService() *schema.Resource {
    return &schema.Resource{
        Create: resourceApplicationServiceCreate,
        Read:   resourceApplicationServiceRead,
        Update: resourceApplicationServiceUpdate,
        Delete: resourceApplicationServiceDelete,
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
            "dscp": &schema.Schema{
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
            "destination_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "direction": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "protocol": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "ether_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
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

func resourceApplicationServiceCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize ApplicationService object
    o := &vspk.ApplicationService{
        Name: d.Get("name").(string),
        Direction: d.Get("direction").(string),
        Protocol: d.Get("protocol").(string),
        EtherType: d.Get("ether_type").(string),
    }
    if attr, ok := d.GetOk("dscp"); ok {
        o.DSCP = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_me"); ok {
        parent := &vspk.Me{ID: attr.(string)}
        err := parent.CreateApplicationService(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        err := parent.CreateApplicationService(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceApplicationServiceRead(d, m)
}

func resourceApplicationServiceRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ApplicationService{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("dscp", o.DSCP)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("description", o.Description)
    d.Set("destination_port", o.DestinationPort)
    d.Set("direction", o.Direction)
    d.Set("entity_scope", o.EntityScope)
    d.Set("source_port", o.SourcePort)
    d.Set("protocol", o.Protocol)
    d.Set("ether_type", o.EtherType)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceApplicationServiceUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ApplicationService{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.Direction = d.Get("direction").(string)
    o.Protocol = d.Get("protocol").(string)
    o.EtherType = d.Get("ether_type").(string)
    
    if attr, ok := d.GetOk("dscp"); ok {
        o.DSCP = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceApplicationServiceDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ApplicationService{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}