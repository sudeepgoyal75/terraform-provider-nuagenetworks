package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.9"
)

func resourceExternalService() *schema.Resource {
    return &schema.Resource{
        Create: resourceExternalServiceCreate,
        Read:   resourceExternalServiceRead,
        Update: resourceExternalServiceUpdate,
        Delete: resourceExternalServiceDelete,
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
            "service_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "direction": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "stage": &schema.Schema{
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

func resourceExternalServiceCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize ExternalService object
    o := &vspk.ExternalService{
        Name: d.Get("name").(string),
        ServiceType: d.Get("service_type").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("direction"); ok {
        o.Direction = attr.(string)
    }
    if attr, ok := d.GetOk("stage"); ok {
        o.Stage = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateExternalService(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceExternalServiceRead(d, m)
}

func resourceExternalServiceRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ExternalService{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("service_type", o.ServiceType)
    d.Set("description", o.Description)
    d.Set("direction", o.Direction)
    d.Set("entity_scope", o.EntityScope)
    d.Set("stage", o.Stage)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceExternalServiceUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ExternalService{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.ServiceType = d.Get("service_type").(string)
    
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("direction"); ok {
        o.Direction = attr.(string)
    }
    if attr, ok := d.GetOk("stage"); ok {
        o.Stage = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceExternalServiceDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ExternalService{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}