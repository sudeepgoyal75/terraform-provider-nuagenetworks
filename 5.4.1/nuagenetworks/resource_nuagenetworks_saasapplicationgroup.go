package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceSaaSApplicationGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceSaaSApplicationGroupCreate,
        Read:   resourceSaaSApplicationGroupRead,
        Update: resourceSaaSApplicationGroupUpdate,
        Delete: resourceSaaSApplicationGroupDelete,
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
            "entity_scope": &schema.Schema{
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
                Required: true,
            },
        },
    }
}

func resourceSaaSApplicationGroupCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize SaaSApplicationGroup object
    o := &vspk.SaaSApplicationGroup{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateSaaSApplicationGroup(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    if attr, ok := d.GetOk("saasapplicationtypes"); ok {
        o.AssignSaaSApplicationTypes(attr.(vspk.SaaSApplicationTypesList))
    }
    return resourceSaaSApplicationGroupRead(d, m)
}

func resourceSaaSApplicationGroupRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SaaSApplicationGroup{
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
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceSaaSApplicationGroupUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SaaSApplicationGroup{
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
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceSaaSApplicationGroupDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SaaSApplicationGroup{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}