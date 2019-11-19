package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
)

func resourceIPv6FilterProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceIPv6FilterProfileCreate,
        Read:   resourceIPv6FilterProfileRead,
        Update: resourceIPv6FilterProfileUpdate,
        Delete: resourceIPv6FilterProfileDelete,
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
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_redundancy_group": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceIPv6FilterProfileCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize IPv6FilterProfile object
    o := &vspk.IPv6FilterProfile{
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    parent := &vspk.RedundancyGroup{ID: d.Get("parent_redundancy_group").(string)}
    err := parent.CreateIPv6FilterProfile(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceIPv6FilterProfileRead(d, m)
}

func resourceIPv6FilterProfileRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IPv6FilterProfile{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("description", o.Description)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceIPv6FilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IPv6FilterProfile{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }

    o.Save()

    return nil
}

func resourceIPv6FilterProfileDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IPv6FilterProfile{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}