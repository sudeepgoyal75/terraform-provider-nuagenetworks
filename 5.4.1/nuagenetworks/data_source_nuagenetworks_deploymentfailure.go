package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceDeploymentFailure() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceDeploymentFailureRead,
        Schema: map[string]*schema.Schema{
            "filter": dataSourceFiltersSchema(),
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_failure_reason": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_known_error": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "affected_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "affected_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "diff_map": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "error_condition": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "assoc_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "assoc_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_network_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_network_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "number_of_occurences": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "event_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile", "parent_redundancy_group", "parent_gateway"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile", "parent_redundancy_group", "parent_gateway"},
            },
            "parent_static_route": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile", "parent_redundancy_group", "parent_gateway"},
            },
            "parent_bridge_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_static_route", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile", "parent_redundancy_group", "parent_gateway"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile", "parent_redundancy_group", "parent_gateway"},
            },
            "parent_egress_profile": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile", "parent_redundancy_group", "parent_gateway"},
            },
            "parent_bgp_neighbor": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_vport", "parent_ingress_profile", "parent_redundancy_group", "parent_gateway"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_ingress_profile", "parent_redundancy_group", "parent_gateway"},
            },
            "parent_ingress_profile": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_redundancy_group", "parent_gateway"},
            },
            "parent_redundancy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile", "parent_gateway"},
            },
            "parent_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_subnet", "parent_egress_profile", "parent_bgp_neighbor", "parent_vport", "parent_ingress_profile", "parent_redundancy_group"},
            },
        },
    }
}


func dataSourceDeploymentFailureRead(d *schema.ResourceData, m interface{}) error {
    filteredDeploymentFailures := vspk.DeploymentFailuresList{}
    err := &bambou.Error{}
    fetchFilter := &bambou.FetchingInfo{}
    
    filters, filtersOk := d.GetOk("filter")
    if filtersOk {
        fetchFilter = bambou.NewFetchingInfo()
        for _, v := range filters.(*schema.Set).List() {
            m := v.(map[string]interface{})
            if fetchFilter.Filter != "" {
                fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string),  m["operator"].(string),  m["value"].(string))
            } else {
                fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
            }
           
        }
    }
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_static_route"); ok {
        parent := &vspk.StaticRoute{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_bridge_interface"); ok {
        parent := &vspk.BridgeInterface{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_egress_profile"); ok {
        parent := &vspk.EgressProfile{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_bgp_neighbor"); ok {
        parent := &vspk.BGPNeighbor{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ingress_profile"); ok {
        parent := &vspk.IngressProfile{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_redundancy_group"); ok {
        parent := &vspk.RedundancyGroup{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    }

    DeploymentFailure := &vspk.DeploymentFailure{}

    if len(filteredDeploymentFailures) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredDeploymentFailures) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    DeploymentFailure = filteredDeploymentFailures[0]

    d.Set("last_failure_reason", DeploymentFailure.LastFailureReason)
    d.Set("last_known_error", DeploymentFailure.LastKnownError)
    d.Set("last_updated_by", DeploymentFailure.LastUpdatedBy)
    d.Set("affected_entity_id", DeploymentFailure.AffectedEntityID)
    d.Set("affected_entity_type", DeploymentFailure.AffectedEntityType)
    d.Set("diff_map", DeploymentFailure.DiffMap)
    d.Set("entity_scope", DeploymentFailure.EntityScope)
    d.Set("error_condition", DeploymentFailure.ErrorCondition)
    d.Set("assoc_entity_id", DeploymentFailure.AssocEntityId)
    d.Set("assoc_entity_type", DeploymentFailure.AssocEntityType)
    d.Set("associated_network_entity_id", DeploymentFailure.AssociatedNetworkEntityID)
    d.Set("associated_network_entity_type", DeploymentFailure.AssociatedNetworkEntityType)
    d.Set("number_of_occurences", DeploymentFailure.NumberOfOccurences)
    d.Set("event_type", DeploymentFailure.EventType)
    d.Set("external_id", DeploymentFailure.ExternalID)
    
    d.Set("id", DeploymentFailure.Identifier())
    d.Set("parent_id", DeploymentFailure.ParentID)
    d.Set("parent_type", DeploymentFailure.ParentType)
    d.Set("owner", DeploymentFailure.Owner)

    d.SetId(DeploymentFailure.Identifier())
    
    return nil
}