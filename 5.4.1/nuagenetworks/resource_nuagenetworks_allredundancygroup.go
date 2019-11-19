package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceAllRedundancyGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceAllRedundancyGroupCreate,
        Read:   resourceAllRedundancyGroupRead,
        Update: resourceAllRedundancyGroupUpdate,
        Delete: resourceAllRedundancyGroupDelete,
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
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_peer1_autodiscovered_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_peer1_connected": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "gateway_peer1_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_peer1_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_peer2_autodiscovered_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_peer2_connected": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "gateway_peer2_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "redundant_gateway_status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vtep": &schema.Schema{
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

func resourceAllRedundancyGroupCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize AllRedundancyGroup object
    o := &vspk.AllRedundancyGroup{
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateAllRedundancyGroup(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceAllRedundancyGroupRead(d, m)
}

func resourceAllRedundancyGroupRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.AllRedundancyGroup{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway_peer1_autodiscovered_gateway_id", o.GatewayPeer1AutodiscoveredGatewayID)
    d.Set("gateway_peer1_connected", o.GatewayPeer1Connected)
    d.Set("gateway_peer1_id", o.GatewayPeer1ID)
    d.Set("gateway_peer1_name", o.GatewayPeer1Name)
    d.Set("gateway_peer2_autodiscovered_gateway_id", o.GatewayPeer2AutodiscoveredGatewayID)
    d.Set("gateway_peer2_connected", o.GatewayPeer2Connected)
    d.Set("gateway_peer2_name", o.GatewayPeer2Name)
    d.Set("redundant_gateway_status", o.RedundantGatewayStatus)
    d.Set("permitted_action", o.PermittedAction)
    d.Set("personality", o.Personality)
    d.Set("description", o.Description)
    d.Set("enterprise_id", o.EnterpriseID)
    d.Set("entity_scope", o.EntityScope)
    d.Set("vtep", o.Vtep)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceAllRedundancyGroupUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.AllRedundancyGroup{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceAllRedundancyGroupDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.AllRedundancyGroup{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}