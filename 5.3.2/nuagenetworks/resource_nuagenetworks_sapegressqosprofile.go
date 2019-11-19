package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
)

func resourceSAPEgressQoSProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceSAPEgressQoSProfileCreate,
        Read:   resourceSAPEgressQoSProfileRead,
        Update: resourceSAPEgressQoSProfileUpdate,
        Delete: resourceSAPEgressQoSProfileDelete,
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

func resourceSAPEgressQoSProfileCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize SAPEgressQoSProfile object
    o := &vspk.SAPEgressQoSProfile{
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    parent := &vspk.RedundancyGroup{ID: d.Get("parent_redundancy_group").(string)}
    err := parent.CreateSAPEgressQoSProfile(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceSAPEgressQoSProfileRead(d, m)
}

func resourceSAPEgressQoSProfileRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SAPEgressQoSProfile{
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

func resourceSAPEgressQoSProfileUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SAPEgressQoSProfile{
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

func resourceSAPEgressQoSProfileDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SAPEgressQoSProfile{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}