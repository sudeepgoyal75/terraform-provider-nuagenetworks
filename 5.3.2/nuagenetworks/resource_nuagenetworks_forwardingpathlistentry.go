package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
)

func resourceForwardingPathListEntry() *schema.Resource {
    return &schema.Resource{
        Create: resourceForwardingPathListEntryCreate,
        Read:   resourceForwardingPathListEntryRead,
        Update: resourceForwardingPathListEntryUpdate,
        Delete: resourceForwardingPathListEntryDelete,
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
            "fc_override": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "forwarding_action": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "uplink_preference": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_forwarding_path_list": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceForwardingPathListEntryCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize ForwardingPathListEntry object
    o := &vspk.ForwardingPathListEntry{
        ForwardingAction: d.Get("forwarding_action").(string),
    }
    if attr, ok := d.GetOk("fc_override"); ok {
        o.FCOverride = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_preference"); ok {
        o.UplinkPreference = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.ForwardingPathList{ID: d.Get("parent_forwarding_path_list").(string)}
    err := parent.CreateForwardingPathListEntry(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceForwardingPathListEntryRead(d, m)
}

func resourceForwardingPathListEntryRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ForwardingPathListEntry{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("fc_override", o.FCOverride)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("entity_scope", o.EntityScope)
    d.Set("forwarding_action", o.ForwardingAction)
    d.Set("uplink_preference", o.UplinkPreference)
    d.Set("priority", o.Priority)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceForwardingPathListEntryUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ForwardingPathListEntry{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.ForwardingAction = d.Get("forwarding_action").(string)
    
    if attr, ok := d.GetOk("fc_override"); ok {
        o.FCOverride = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_preference"); ok {
        o.UplinkPreference = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceForwardingPathListEntryDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ForwardingPathListEntry{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}