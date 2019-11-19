package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceAllRedundancyGroup() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceAllRedundancyGroupRead,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_peer1_autodiscovered_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_peer1_connected": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "gateway_peer1_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_peer1_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_peer2_autodiscovered_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_peer2_connected": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "gateway_peer2_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "redundant_gateway_status": &schema.Schema{
                Type:     schema.TypeString,
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
            "vtep": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceAllRedundancyGroupRead(d *schema.ResourceData, m interface{}) error {
    filteredAllRedundancyGroups := vspk.AllRedundancyGroupsList{}
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
    parent := m.(*vspk.Me)
    filteredAllRedundancyGroups, err = parent.AllRedundancyGroups(fetchFilter)
    if err != nil {
        return err
    }

    AllRedundancyGroup := &vspk.AllRedundancyGroup{}

    if len(filteredAllRedundancyGroups) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredAllRedundancyGroups) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    AllRedundancyGroup = filteredAllRedundancyGroups[0]

    d.Set("name", AllRedundancyGroup.Name)
    d.Set("last_updated_by", AllRedundancyGroup.LastUpdatedBy)
    d.Set("gateway_peer1_autodiscovered_gateway_id", AllRedundancyGroup.GatewayPeer1AutodiscoveredGatewayID)
    d.Set("gateway_peer1_connected", AllRedundancyGroup.GatewayPeer1Connected)
    d.Set("gateway_peer1_id", AllRedundancyGroup.GatewayPeer1ID)
    d.Set("gateway_peer1_name", AllRedundancyGroup.GatewayPeer1Name)
    d.Set("gateway_peer2_autodiscovered_gateway_id", AllRedundancyGroup.GatewayPeer2AutodiscoveredGatewayID)
    d.Set("gateway_peer2_connected", AllRedundancyGroup.GatewayPeer2Connected)
    d.Set("gateway_peer2_name", AllRedundancyGroup.GatewayPeer2Name)
    d.Set("redundant_gateway_status", AllRedundancyGroup.RedundantGatewayStatus)
    d.Set("permitted_action", AllRedundancyGroup.PermittedAction)
    d.Set("personality", AllRedundancyGroup.Personality)
    d.Set("description", AllRedundancyGroup.Description)
    d.Set("enterprise_id", AllRedundancyGroup.EnterpriseID)
    d.Set("entity_scope", AllRedundancyGroup.EntityScope)
    d.Set("vtep", AllRedundancyGroup.Vtep)
    d.Set("external_id", AllRedundancyGroup.ExternalID)
    
    d.Set("id", AllRedundancyGroup.Identifier())
    d.Set("parent_id", AllRedundancyGroup.ParentID)
    d.Set("parent_type", AllRedundancyGroup.ParentType)
    d.Set("owner", AllRedundancyGroup.Owner)

    d.SetId(AllRedundancyGroup.Identifier())
    
    return nil
}