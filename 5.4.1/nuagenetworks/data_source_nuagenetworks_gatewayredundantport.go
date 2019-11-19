package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceGatewayRedundantPort() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceGatewayRedundantPortRead,
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
            "vlan_range": &schema.Schema{
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
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "physical_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "port_peer1_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "port_peer2_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "port_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "use_user_mnemonic": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "user_mnemonic": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_egress_qos_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_redundancy_group": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceGatewayRedundantPortRead(d *schema.ResourceData, m interface{}) error {
    filteredGatewayRedundantPorts := vspk.GatewayRedundantPortsList{}
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
    parent := &vspk.RedundancyGroup{ID: d.Get("parent_redundancy_group").(string)}
    filteredGatewayRedundantPorts, err = parent.GatewayRedundantPorts(fetchFilter)
    if err != nil {
        return err
    }

    GatewayRedundantPort := &vspk.GatewayRedundantPort{}

    if len(filteredGatewayRedundantPorts) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredGatewayRedundantPorts) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    GatewayRedundantPort = filteredGatewayRedundantPorts[0]

    d.Set("vlan_range", GatewayRedundantPort.VLANRange)
    d.Set("name", GatewayRedundantPort.Name)
    d.Set("last_updated_by", GatewayRedundantPort.LastUpdatedBy)
    d.Set("permitted_action", GatewayRedundantPort.PermittedAction)
    d.Set("description", GatewayRedundantPort.Description)
    d.Set("physical_name", GatewayRedundantPort.PhysicalName)
    d.Set("entity_scope", GatewayRedundantPort.EntityScope)
    d.Set("port_peer1_id", GatewayRedundantPort.PortPeer1ID)
    d.Set("port_peer2_id", GatewayRedundantPort.PortPeer2ID)
    d.Set("port_type", GatewayRedundantPort.PortType)
    d.Set("use_user_mnemonic", GatewayRedundantPort.UseUserMnemonic)
    d.Set("user_mnemonic", GatewayRedundantPort.UserMnemonic)
    d.Set("associated_egress_qos_policy_id", GatewayRedundantPort.AssociatedEgressQOSPolicyID)
    d.Set("status", GatewayRedundantPort.Status)
    d.Set("external_id", GatewayRedundantPort.ExternalID)
    
    d.Set("id", GatewayRedundantPort.Identifier())
    d.Set("parent_id", GatewayRedundantPort.ParentID)
    d.Set("parent_type", GatewayRedundantPort.ParentType)
    d.Set("owner", GatewayRedundantPort.Owner)

    d.SetId(GatewayRedundantPort.Identifier())
    
    return nil
}