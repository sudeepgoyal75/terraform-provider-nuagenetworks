package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceAllGateway() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceAllGatewayRead,
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
            "mac_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zfb_match_attribute": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zfb_match_value": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "bios_release_date": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "bios_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "cpu_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "uuid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "family": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "management_id": &schema.Schema{
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
            "patches": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_connected": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "gateway_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "redundancy_group_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "peer": &schema.Schema{
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
            "serial_number": &schema.Schema{
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
            "libraries": &schema.Schema{
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
            "bootstrap_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "bootstrap_status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "product_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "use_gateway_vlanvnid": &schema.Schema{
                Type:     schema.TypeBool,
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
            "associated_nsg_info_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_netconf_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vtep": &schema.Schema{
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
        },
    }
}


func dataSourceAllGatewayRead(d *schema.ResourceData, m interface{}) error {
    filteredAllGateways := vspk.AllGatewaysList{}
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
    filteredAllGateways, err = parent.AllGateways(fetchFilter)
    if err != nil {
        return err
    }

    AllGateway := &vspk.AllGateway{}

    if len(filteredAllGateways) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredAllGateways) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    AllGateway = filteredAllGateways[0]

    d.Set("mac_address", AllGateway.MACAddress)
    d.Set("zfb_match_attribute", AllGateway.ZFBMatchAttribute)
    d.Set("zfb_match_value", AllGateway.ZFBMatchValue)
    d.Set("bios_release_date", AllGateway.BIOSReleaseDate)
    d.Set("bios_version", AllGateway.BIOSVersion)
    d.Set("cpu_type", AllGateway.CPUType)
    d.Set("uuid", AllGateway.UUID)
    d.Set("name", AllGateway.Name)
    d.Set("family", AllGateway.Family)
    d.Set("management_id", AllGateway.ManagementID)
    d.Set("last_updated_by", AllGateway.LastUpdatedBy)
    d.Set("datapath_id", AllGateway.DatapathID)
    d.Set("patches", AllGateway.Patches)
    d.Set("gateway_connected", AllGateway.GatewayConnected)
    d.Set("gateway_version", AllGateway.GatewayVersion)
    d.Set("redundancy_group_id", AllGateway.RedundancyGroupID)
    d.Set("peer", AllGateway.Peer)
    d.Set("template_id", AllGateway.TemplateID)
    d.Set("pending", AllGateway.Pending)
    d.Set("serial_number", AllGateway.SerialNumber)
    d.Set("permitted_action", AllGateway.PermittedAction)
    d.Set("personality", AllGateway.Personality)
    d.Set("description", AllGateway.Description)
    d.Set("libraries", AllGateway.Libraries)
    d.Set("enterprise_id", AllGateway.EnterpriseID)
    d.Set("entity_scope", AllGateway.EntityScope)
    d.Set("location_id", AllGateway.LocationID)
    d.Set("bootstrap_id", AllGateway.BootstrapID)
    d.Set("bootstrap_status", AllGateway.BootstrapStatus)
    d.Set("product_name", AllGateway.ProductName)
    d.Set("use_gateway_vlanvnid", AllGateway.UseGatewayVLANVNID)
    d.Set("associated_gateway_security_id", AllGateway.AssociatedGatewaySecurityID)
    d.Set("associated_gateway_security_profile_id", AllGateway.AssociatedGatewaySecurityProfileID)
    d.Set("associated_nsg_info_id", AllGateway.AssociatedNSGInfoID)
    d.Set("associated_netconf_profile_id", AllGateway.AssociatedNetconfProfileID)
    d.Set("vtep", AllGateway.Vtep)
    d.Set("auto_disc_gateway_id", AllGateway.AutoDiscGatewayID)
    d.Set("external_id", AllGateway.ExternalID)
    d.Set("system_id", AllGateway.SystemID)
    
    d.Set("id", AllGateway.Identifier())
    d.Set("parent_id", AllGateway.ParentID)
    d.Set("parent_type", AllGateway.ParentType)
    d.Set("owner", AllGateway.Owner)

    d.SetId(AllGateway.Identifier())
    
    return nil
}