package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceSyslogDestination() *schema.Resource {
    return &schema.Resource{
        Create: resourceSyslogDestinationCreate,
        Read:   resourceSyslogDestinationRead,
        Update: resourceSyslogDestinationUpdate,
        Delete: resourceSyslogDestinationDelete,
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
            "ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "ip_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "IPV4",
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "port": &schema.Schema{
                Type:     schema.TypeInt,
                Required: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "UDP",
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceSyslogDestinationCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize SyslogDestination object
    o := &vspk.SyslogDestination{
        IPAddress: d.Get("ip_address").(string),
        Name: d.Get("name").(string),
        Port: d.Get("port").(int),
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateSyslogDestination(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceSyslogDestinationRead(d, m)
}

func resourceSyslogDestinationRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SyslogDestination{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("ip_address", o.IPAddress)
    d.Set("ip_type", o.IPType)
    d.Set("name", o.Name)
    d.Set("description", o.Description)
    d.Set("port", o.Port)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceSyslogDestinationUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SyslogDestination{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.IPAddress = d.Get("ip_address").(string)
    o.Name = d.Get("name").(string)
    o.Port = d.Get("port").(int)
    
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }

    o.Save()

    return nil
}

func resourceSyslogDestinationDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SyslogDestination{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}