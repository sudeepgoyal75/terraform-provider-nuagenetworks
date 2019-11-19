package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
)

func resourceNetconfManager() *schema.Resource {
    return &schema.Resource{
        Create: resourceNetconfManagerCreate,
        Read:   resourceNetconfManagerRead,
        Update: resourceNetconfManagerUpdate,
        Delete: resourceNetconfManagerDelete,
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
            "release": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_vsp": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceNetconfManagerCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize NetconfManager object
    o := &vspk.NetconfManager{
    }
    parent := &vspk.VSP{ID: d.Get("parent_vsp").(string)}
    err := parent.CreateNetconfManager(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceNetconfManagerRead(d, m)
}

func resourceNetconfManagerRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetconfManager{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("release", o.Release)
    d.Set("status", o.Status)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceNetconfManagerUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetconfManager{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    

    o.Save()

    return nil
}

func resourceNetconfManagerDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetconfManager{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}