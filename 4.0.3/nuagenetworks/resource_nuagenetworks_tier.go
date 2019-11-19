package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.3"
)

func resourceTier() *schema.Resource {
    return &schema.Resource{
        Create: resourceTierCreate,
        Read:   resourceTierRead,
        Update: resourceTierUpdate,
        Delete: resourceTierDelete,
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
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "metadata": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_application_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_floating_ip_pool_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_network_macro_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_network_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_network_object_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "parent_app": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceTierCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Tier object
    o := &vspk.Tier{
        Name: d.Get("name").(string),
        Type: d.Get("type").(string),
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("metadata"); ok {
        o.Metadata = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_id"); ok {
        o.AssociatedApplicationID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_floating_ip_pool_id"); ok {
        o.AssociatedFloatingIPPoolID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_network_macro_id"); ok {
        o.AssociatedNetworkMacroID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_network_object_id"); ok {
        o.AssociatedNetworkObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_network_object_type"); ok {
        o.AssociatedNetworkObjectType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.App{ID: d.Get("parent_app").(string)}
    err := parent.CreateTier(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceTierRead(d, m)
}

func resourceTierRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Tier{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway", o.Gateway)
    d.Set("address", o.Address)
    d.Set("description", o.Description)
    d.Set("metadata", o.Metadata)
    d.Set("netmask", o.Netmask)
    d.Set("entity_scope", o.EntityScope)
    d.Set("associated_application_id", o.AssociatedApplicationID)
    d.Set("associated_floating_ip_pool_id", o.AssociatedFloatingIPPoolID)
    d.Set("associated_network_macro_id", o.AssociatedNetworkMacroID)
    d.Set("associated_network_object_id", o.AssociatedNetworkObjectID)
    d.Set("associated_network_object_type", o.AssociatedNetworkObjectType)
    d.Set("external_id", o.ExternalID)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceTierUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Tier{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.Type = d.Get("type").(string)
    
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("metadata"); ok {
        o.Metadata = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_id"); ok {
        o.AssociatedApplicationID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_floating_ip_pool_id"); ok {
        o.AssociatedFloatingIPPoolID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_network_macro_id"); ok {
        o.AssociatedNetworkMacroID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_network_object_id"); ok {
        o.AssociatedNetworkObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_network_object_type"); ok {
        o.AssociatedNetworkObjectType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceTierDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Tier{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}