package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceWirelessPort() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceWirelessPortRead,
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
            "generic_config": &schema.Schema{
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
            "wifi_frequency_band": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "wifi_mode": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "port_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "country_code": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "frequency_channel": &schema.Schema{
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
            "parent_auto_discovered_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_ns_gateway"},
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_auto_discovered_gateway"},
            },
        },
    }
}


func dataSourceWirelessPortRead(d *schema.ResourceData, m interface{}) error {
    filteredWirelessPorts := vspk.WirelessPortsList{}
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
    if attr, ok := d.GetOk("parent_auto_discovered_gateway"); ok {
        parent := &vspk.AutoDiscoveredGateway{ID: attr.(string)}
        filteredWirelessPorts, err = parent.WirelessPorts(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
        parent := &vspk.NSGateway{ID: attr.(string)}
        filteredWirelessPorts, err = parent.WirelessPorts(fetchFilter)
        if err != nil {
            return err
        }
    }

    WirelessPort := &vspk.WirelessPort{}

    if len(filteredWirelessPorts) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredWirelessPorts) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    WirelessPort = filteredWirelessPorts[0]

    d.Set("vlan_range", WirelessPort.VLANRange)
    d.Set("name", WirelessPort.Name)
    d.Set("last_updated_by", WirelessPort.LastUpdatedBy)
    d.Set("generic_config", WirelessPort.GenericConfig)
    d.Set("permitted_action", WirelessPort.PermittedAction)
    d.Set("description", WirelessPort.Description)
    d.Set("physical_name", WirelessPort.PhysicalName)
    d.Set("wifi_frequency_band", WirelessPort.WifiFrequencyBand)
    d.Set("wifi_mode", WirelessPort.WifiMode)
    d.Set("entity_scope", WirelessPort.EntityScope)
    d.Set("port_type", WirelessPort.PortType)
    d.Set("country_code", WirelessPort.CountryCode)
    d.Set("frequency_channel", WirelessPort.FrequencyChannel)
    d.Set("use_user_mnemonic", WirelessPort.UseUserMnemonic)
    d.Set("user_mnemonic", WirelessPort.UserMnemonic)
    d.Set("associated_egress_qos_policy_id", WirelessPort.AssociatedEgressQOSPolicyID)
    d.Set("status", WirelessPort.Status)
    d.Set("external_id", WirelessPort.ExternalID)
    
    d.Set("id", WirelessPort.Identifier())
    d.Set("parent_id", WirelessPort.ParentID)
    d.Set("parent_type", WirelessPort.ParentType)
    d.Set("owner", WirelessPort.Owner)

    d.SetId(WirelessPort.Identifier())
    
    return nil
}