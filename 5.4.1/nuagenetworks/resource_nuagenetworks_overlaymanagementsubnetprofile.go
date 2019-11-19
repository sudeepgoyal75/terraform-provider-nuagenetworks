package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceOverlayManagementSubnetProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceOverlayManagementSubnetProfileCreate,
        Read:   resourceOverlayManagementSubnetProfileRead,
        Update: resourceOverlayManagementSubnetProfileUpdate,
        Delete: resourceOverlayManagementSubnetProfileDelete,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_dna_subnet_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "syslog_destination_ids": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "parent_overlay_management_profile": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceOverlayManagementSubnetProfileCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize OverlayManagementSubnetProfile object
    o := &vspk.OverlayManagementSubnetProfile{
        Name: d.Get("name").(string),
        AssociatedDNASubnetID: d.Get("associated_dna_subnet_id").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("syslog_destination_ids"); ok {
        o.SyslogDestinationIDs = attr.([]interface{})
    }
    parent := &vspk.OverlayManagementProfile{ID: d.Get("parent_overlay_management_profile").(string)}
    err := parent.CreateOverlayManagementSubnetProfile(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceOverlayManagementSubnetProfileRead(d, m)
}

func resourceOverlayManagementSubnetProfileRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayManagementSubnetProfile{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("description", o.Description)
    d.Set("associated_dna_subnet_id", o.AssociatedDNASubnetID)
    d.Set("syslog_destination_ids", o.SyslogDestinationIDs)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceOverlayManagementSubnetProfileUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayManagementSubnetProfile{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.AssociatedDNASubnetID = d.Get("associated_dna_subnet_id").(string)
    
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("syslog_destination_ids"); ok {
        o.SyslogDestinationIDs = attr.([]interface{})
    }

    o.Save()

    return nil
}

func resourceOverlayManagementSubnetProfileDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayManagementSubnetProfile{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}