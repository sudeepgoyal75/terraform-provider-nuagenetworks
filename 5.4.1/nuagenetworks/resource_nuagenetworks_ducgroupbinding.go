package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceDUCGroupBinding() *schema.Resource {
    return &schema.Resource{
        Create: resourceDUCGroupBindingCreate,
        Read:   resourceDUCGroupBindingRead,
        Update: resourceDUCGroupBindingUpdate,
        Delete: resourceDUCGroupBindingDelete,
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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "one_way_delay": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default: 50,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "associated_duc_group_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_ubr_group_function": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_ubr_group_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_nsg_group": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceDUCGroupBindingCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize DUCGroupBinding object
    o := &vspk.DUCGroupBinding{
    }
    if attr, ok := d.GetOk("one_way_delay"); ok {
        o.OneWayDelay = attr.(int)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("associated_duc_group_id"); ok {
        o.AssociatedDUCGroupID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.NSGGroup{ID: d.Get("parent_nsg_group").(string)}
    err := parent.CreateDUCGroupBinding(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceDUCGroupBindingRead(d, m)
}

func resourceDUCGroupBindingRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DUCGroupBinding{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("one_way_delay", o.OneWayDelay)
    d.Set("entity_scope", o.EntityScope)
    d.Set("priority", o.Priority)
    d.Set("associated_duc_group_id", o.AssociatedDUCGroupID)
    d.Set("associated_ubr_group_function", o.AssociatedUBRGroupFunction)
    d.Set("associated_ubr_group_name", o.AssociatedUBRGroupName)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceDUCGroupBindingUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DUCGroupBinding{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("one_way_delay"); ok {
        o.OneWayDelay = attr.(int)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("associated_duc_group_id"); ok {
        o.AssociatedDUCGroupID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceDUCGroupBindingDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DUCGroupBinding{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}