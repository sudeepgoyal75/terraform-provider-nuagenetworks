package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.3"
)

func resourceFlow() *schema.Resource {
    return &schema.Resource{
        Create: resourceFlowCreate,
        Read:   resourceFlowRead,
        Update: resourceFlowUpdate,
        Delete: resourceFlowDelete,
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
            "destination_tier_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "metadata": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "origin_tier_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_app": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceFlowCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Flow object
    o := &vspk.Flow{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_tier_id"); ok {
        o.DestinationTierID = attr.(string)
    }
    if attr, ok := d.GetOk("metadata"); ok {
        o.Metadata = attr.(string)
    }
    if attr, ok := d.GetOk("origin_tier_id"); ok {
        o.OriginTierID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.App{ID: d.Get("parent_app").(string)}
    err := parent.CreateFlow(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceFlowRead(d, m)
}

func resourceFlowRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Flow{
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
    d.Set("destination_tier_id", o.DestinationTierID)
    d.Set("metadata", o.Metadata)
    d.Set("entity_scope", o.EntityScope)
    d.Set("origin_tier_id", o.OriginTierID)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceFlowUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Flow{
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
    if attr, ok := d.GetOk("destination_tier_id"); ok {
        o.DestinationTierID = attr.(string)
    }
    if attr, ok := d.GetOk("metadata"); ok {
        o.Metadata = attr.(string)
    }
    if attr, ok := d.GetOk("origin_tier_id"); ok {
        o.OriginTierID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceFlowDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Flow{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}