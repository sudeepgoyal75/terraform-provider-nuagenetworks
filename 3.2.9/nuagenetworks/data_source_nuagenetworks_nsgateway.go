package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.9"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNSGateway() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNSGatewayRead,
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
            "nat_traversal_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "datapath_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "redundancy_group_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "pending": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "location_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "configuration_reload_state": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "configuration_status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "bootstrap_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "bootstrap_status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_gateway_security_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_gateway_security_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "auto_disc_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "system_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_ns_redundant_gateway_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_ns_redundant_gateway_group"},
            },
        },
    }
}


func dataSourceNSGatewayRead(d *schema.ResourceData, m interface{}) error {
    filteredNSGateways := vspk.NSGatewaysList{}
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
    if attr, ok := d.GetOk("parent_ns_redundant_gateway_group"); ok {
        parent := &vspk.NSRedundantGatewayGroup{ID: attr.(string)}
        filteredNSGateways, err = parent.NSGateways(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredNSGateways, err = parent.NSGateways(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredNSGateways, err = parent.NSGateways(fetchFilter)
        if err != nil {
            return err
        }
    }

    NSGateway := &vspk.NSGateway{}

    if len(filteredNSGateways) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNSGateways) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NSGateway = filteredNSGateways[0]

    d.Set("nat_traversal_enabled", NSGateway.NATTraversalEnabled)
    d.Set("name", NSGateway.Name)
    d.Set("last_updated_by", NSGateway.LastUpdatedBy)
    d.Set("datapath_id", NSGateway.DatapathID)
    d.Set("redundancy_group_id", NSGateway.RedundancyGroupID)
    d.Set("template_id", NSGateway.TemplateID)
    d.Set("pending", NSGateway.Pending)
    d.Set("permitted_action", NSGateway.PermittedAction)
    d.Set("personality", NSGateway.Personality)
    d.Set("description", NSGateway.Description)
    d.Set("enterprise_id", NSGateway.EnterpriseID)
    d.Set("entity_scope", NSGateway.EntityScope)
    d.Set("location_id", NSGateway.LocationID)
    d.Set("configuration_reload_state", NSGateway.ConfigurationReloadState)
    d.Set("configuration_status", NSGateway.ConfigurationStatus)
    d.Set("bootstrap_id", NSGateway.BootstrapID)
    d.Set("bootstrap_status", NSGateway.BootstrapStatus)
    d.Set("associated_gateway_security_id", NSGateway.AssociatedGatewaySecurityID)
    d.Set("associated_gateway_security_profile_id", NSGateway.AssociatedGatewaySecurityProfileID)
    d.Set("auto_disc_gateway_id", NSGateway.AutoDiscGatewayID)
    d.Set("external_id", NSGateway.ExternalID)
    d.Set("system_id", NSGateway.SystemID)
    
    d.Set("id", NSGateway.Identifier())
    d.Set("parent_id", NSGateway.ParentID)
    d.Set("parent_type", NSGateway.ParentType)
    d.Set("owner", NSGateway.Owner)

    d.SetId(NSGateway.Identifier())
    
    return nil
}