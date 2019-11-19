package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceDeploymentFailure() *schema.Resource {
    return &schema.Resource{
        Create: resourceDeploymentFailureCreate,
        Read:   resourceDeploymentFailureRead,
        Update: resourceDeploymentFailureUpdate,
        Delete: resourceDeploymentFailureDelete,
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
            "last_failure_reason": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_known_error": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "affected_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "affected_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "diff_map": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "error_condition": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "assoc_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "assoc_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_network_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_network_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "number_of_occurences": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "event_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile"},
            },
            "parent_static_route": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile"},
            },
            "parent_bridge_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_static_route", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile"},
            },
            "parent_egress_profile": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile"},
            },
            "parent_bgp_neighbor": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_vport", "parent_ingress_profile"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_ingress_profile"},
            },
            "parent_ingress_profile": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport"},
            },
        },
    }
}

func resourceDeploymentFailureCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize DeploymentFailure object
    o := &vspk.DeploymentFailure{
    }
    if attr, ok := d.GetOk("last_failure_reason"); ok {
        o.LastFailureReason = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        err := parent.CreateDeploymentFailure(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_static_route"); ok {
        parent := &vspk.StaticRoute{ID: attr.(string)}
        err := parent.CreateDeploymentFailure(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_bridge_interface"); ok {
        parent := &vspk.BridgeInterface{ID: attr.(string)}
        err := parent.CreateDeploymentFailure(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        err := parent.CreateDeploymentFailure(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_egress_profile"); ok {
        parent := &vspk.EgressProfile{ID: attr.(string)}
        err := parent.CreateDeploymentFailure(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_bgp_neighbor"); ok {
        parent := &vspk.BGPNeighbor{ID: attr.(string)}
        err := parent.CreateDeploymentFailure(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        err := parent.CreateDeploymentFailure(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_ingress_profile"); ok {
        parent := &vspk.IngressProfile{ID: attr.(string)}
        err := parent.CreateDeploymentFailure(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceDeploymentFailureRead(d, m)
}

func resourceDeploymentFailureRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DeploymentFailure{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("last_failure_reason", o.LastFailureReason)
    d.Set("last_known_error", o.LastKnownError)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("affected_entity_id", o.AffectedEntityID)
    d.Set("affected_entity_type", o.AffectedEntityType)
    d.Set("diff_map", o.DiffMap)
    d.Set("entity_scope", o.EntityScope)
    d.Set("error_condition", o.ErrorCondition)
    d.Set("assoc_entity_id", o.AssocEntityId)
    d.Set("assoc_entity_type", o.AssocEntityType)
    d.Set("associated_network_entity_id", o.AssociatedNetworkEntityID)
    d.Set("associated_network_entity_type", o.AssociatedNetworkEntityType)
    d.Set("number_of_occurences", o.NumberOfOccurences)
    d.Set("event_type", o.EventType)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceDeploymentFailureUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DeploymentFailure{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("last_failure_reason"); ok {
        o.LastFailureReason = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceDeploymentFailureDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DeploymentFailure{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}