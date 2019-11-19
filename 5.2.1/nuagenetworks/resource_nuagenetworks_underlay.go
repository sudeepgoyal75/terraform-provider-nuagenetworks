package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.2.1"
)

func resourceUnderlay() *schema.Resource {
    return &schema.Resource{
        Create: resourceUnderlayCreate,
        Read:   resourceUnderlayRead,
        Update: resourceUnderlayUpdate,
        Delete: resourceUnderlayDelete,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "underlay_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
        },
    }
}

func resourceUnderlayCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Underlay object
    o := &vspk.Underlay{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("underlay_id"); ok {
        o.UnderlayID = attr.(int)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateUnderlay(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceUnderlayRead(d, m)
}

func resourceUnderlayRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Underlay{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("description", o.Description)
    d.Set("underlay_id", o.UnderlayID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceUnderlayUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Underlay{
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
    if attr, ok := d.GetOk("underlay_id"); ok {
        o.UnderlayID = attr.(int)
    }

    o.Save()

    return nil
}

func resourceUnderlayDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Underlay{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}