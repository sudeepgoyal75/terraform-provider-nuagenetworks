package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
)

func resourceShuntLink() *schema.Resource {
    return &schema.Resource{
        Create: resourceShuntLinkCreate,
        Read:   resourceShuntLinkRead,
        Update: resourceShuntLinkUpdate,
        Delete: resourceShuntLinkDelete,
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
            "vlan_peer1_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "vlan_peer2_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
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
            "gateway_peer1_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_peer2_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "null",
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
            "parent_ns_redundant_gateway_group": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceShuntLinkCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize ShuntLink object
    o := &vspk.ShuntLink{
        VLANPeer1ID: d.Get("vlan_peer1_id").(string),
        VLANPeer2ID: d.Get("vlan_peer2_id").(string),
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.NSRedundantGatewayGroup{ID: d.Get("parent_ns_redundant_gateway_group").(string)}
    err := parent.CreateShuntLink(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceShuntLinkRead(d, m)
}

func resourceShuntLinkRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ShuntLink{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("vlan_peer1_id", o.VLANPeer1ID)
    d.Set("vlan_peer2_id", o.VLANPeer2ID)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway_peer1_id", o.GatewayPeer1ID)
    d.Set("gateway_peer2_id", o.GatewayPeer2ID)
    d.Set("permitted_action", o.PermittedAction)
    d.Set("description", o.Description)
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceShuntLinkUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ShuntLink{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.VLANPeer1ID = d.Get("vlan_peer1_id").(string)
    o.VLANPeer2ID = d.Get("vlan_peer2_id").(string)
    
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceShuntLinkDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ShuntLink{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}