package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
)

func resourceSPATSourcesPool() *schema.Resource {
    return &schema.Resource{
        Create: resourceSPATSourcesPoolCreate,
        Read:   resourceSPATSourcesPoolRead,
        Update: resourceSPATSourcesPoolUpdate,
        Delete: resourceSPATSourcesPoolDelete,
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
                Default: "IPV4",
            },
            "family": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "address_list": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
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
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceSPATSourcesPoolCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize SPATSourcesPool object
    o := &vspk.SPATSourcesPool{
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("family"); ok {
        o.Family = attr.(string)
    }
    if attr, ok := d.GetOk("address_list"); ok {
        o.AddressList = attr.([]interface{})
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
    err := parent.CreateSPATSourcesPool(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceSPATSourcesPoolRead(d, m)
}

func resourceSPATSourcesPoolRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SPATSourcesPool{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("family", o.Family)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("address_list", o.AddressList)
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceSPATSourcesPoolUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SPATSourcesPool{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("family"); ok {
        o.Family = attr.(string)
    }
    if attr, ok := d.GetOk("address_list"); ok {
        o.AddressList = attr.([]interface{})
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceSPATSourcesPoolDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SPATSourcesPool{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}