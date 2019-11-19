package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourcePSPATMap() *schema.Resource {
    return &schema.Resource{
        Create: resourcePSPATMapCreate,
        Read:   resourcePSPATMapRead,
        Update: resourcePSPATMapUpdate,
        Delete: resourcePSPATMapDelete,
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
            "reserved_spatips": &schema.Schema{
                Type:     schema.TypeList,
                Required: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_spat_sources_pool_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_psnat_pool": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourcePSPATMapCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize PSPATMap object
    o := &vspk.PSPATMap{
        Name: d.Get("name").(string),
        ReservedSPATIPs: d.Get("reserved_spatips").([]interface{}),
        AssociatedSPATSourcesPoolID: d.Get("associated_spat_sources_pool_id").(string),
    }
    if attr, ok := d.GetOk("family"); ok {
        o.Family = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.PSNATPool{ID: d.Get("parent_psnat_pool").(string)}
    err := parent.CreatePSPATMap(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourcePSPATMapRead(d, m)
}

func resourcePSPATMapRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PSPATMap{
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
    d.Set("reserved_spatips", o.ReservedSPATIPs)
    d.Set("entity_scope", o.EntityScope)
    d.Set("associated_spat_sources_pool_id", o.AssociatedSPATSourcesPoolID)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourcePSPATMapUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PSPATMap{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.ReservedSPATIPs = d.Get("reserved_spatips").([]interface{})
    o.AssociatedSPATSourcesPoolID = d.Get("associated_spat_sources_pool_id").(string)
    
    if attr, ok := d.GetOk("family"); ok {
        o.Family = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourcePSPATMapDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PSPATMap{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}