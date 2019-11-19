package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
)

func resourceOverlayAddressPool() *schema.Resource {
    return &schema.Resource{
        Create: resourceOverlayAddressPoolCreate,
        Read:   resourceOverlayAddressPoolRead,
        Update: resourceOverlayAddressPoolUpdate,
        Delete: resourceOverlayAddressPoolDelete,
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
            "ip_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
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
            "end_address_range": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "start_address_range": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_link": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceOverlayAddressPoolCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize OverlayAddressPool object
    o := &vspk.OverlayAddressPool{
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("end_address_range"); ok {
        o.EndAddressRange = attr.(string)
    }
    if attr, ok := d.GetOk("associated_domain_id"); ok {
        o.AssociatedDomainID = attr.(string)
    }
    if attr, ok := d.GetOk("start_address_range"); ok {
        o.StartAddressRange = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Link{ID: d.Get("parent_link").(string)}
    err := parent.CreateOverlayAddressPool(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceOverlayAddressPoolRead(d, m)
}

func resourceOverlayAddressPoolRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayAddressPool{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("ip_type", o.IPType)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("description", o.Description)
    d.Set("end_address_range", o.EndAddressRange)
    d.Set("entity_scope", o.EntityScope)
    d.Set("associated_domain_id", o.AssociatedDomainID)
    d.Set("start_address_range", o.StartAddressRange)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceOverlayAddressPoolUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayAddressPool{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("end_address_range"); ok {
        o.EndAddressRange = attr.(string)
    }
    if attr, ok := d.GetOk("associated_domain_id"); ok {
        o.AssociatedDomainID = attr.(string)
    }
    if attr, ok := d.GetOk("start_address_range"); ok {
        o.StartAddressRange = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceOverlayAddressPoolDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayAddressPool{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}