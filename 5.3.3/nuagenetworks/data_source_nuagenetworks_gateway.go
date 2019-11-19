package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceGateway() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceGatewayRead,
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
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise", "parent_redundancy_group"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_redundancy_group"},
            },
            "parent_redundancy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_enterprise"},
            },
        },
    }
}


func dataSourceGatewayRead(d *schema.ResourceData, m interface{}) error {
    filteredGateways := vspk.GatewaysList{}
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
    if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredGateways, err = parent.Gateways(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredGateways, err = parent.Gateways(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_redundancy_group"); ok {
        parent := &vspk.RedundancyGroup{ID: attr.(string)}
        filteredGateways, err = parent.Gateways(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredGateways, err = parent.Gateways(fetchFilter)
        if err != nil {
            return err
        }
    }

    Gateway := &vspk.Gateway{}

    if len(filteredGateways) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredGateways) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Gateway = filteredGateways[0]

    d.Set("mac_address", Gateway.MACAddress)
    d.Set("zfb_match_attribute", Gateway.ZFBMatchAttribute)
    d.Set("zfb_match_value", Gateway.ZFBMatchValue)
    d.Set("bios_release_date", Gateway.BIOSReleaseDate)
    d.Set("bios_version", Gateway.BIOSVersion)
    d.Set("cpu_type", Gateway.CPUType)
    d.Set("uuid", Gateway.UUID)
    d.Set("name", Gateway.Name)
    d.Set("family", Gateway.Family)
    d.Set("management_id", Gateway.ManagementID)
    d.Set("last_updated_by", Gateway.LastUpdatedBy)
    d.Set("datapath_id", Gateway.DatapathID)
    d.Set("patches", Gateway.Patches)
    d.Set("gateway_connected", Gateway.GatewayConnected)
    d.Set("gateway_version", Gateway.GatewayVersion)
    d.Set("redundancy_group_id", Gateway.RedundancyGroupID)
    d.Set("peer", Gateway.Peer)
    d.Set("template_id", Gateway.TemplateID)
    d.Set("pending", Gateway.Pending)
    d.Set("serial_number", Gateway.SerialNumber)
    d.Set("permitted_action", Gateway.PermittedAction)
    d.Set("personality", Gateway.Personality)
    d.Set("description", Gateway.Description)
    d.Set("libraries", Gateway.Libraries)
    d.Set("enterprise_id", Gateway.EnterpriseID)
    d.Set("entity_scope", Gateway.EntityScope)
    d.Set("location_id", Gateway.LocationID)
    d.Set("bootstrap_id", Gateway.BootstrapID)
    d.Set("bootstrap_status", Gateway.BootstrapStatus)
    d.Set("product_name", Gateway.ProductName)
    d.Set("use_gateway_vlanvnid", Gateway.UseGatewayVLANVNID)
    d.Set("associated_gateway_security_id", Gateway.AssociatedGatewaySecurityID)
    d.Set("associated_gateway_security_profile_id", Gateway.AssociatedGatewaySecurityProfileID)
    d.Set("associated_nsg_info_id", Gateway.AssociatedNSGInfoID)
    d.Set("associated_netconf_profile_id", Gateway.AssociatedNetconfProfileID)
    d.Set("vtep", Gateway.Vtep)
    d.Set("auto_disc_gateway_id", Gateway.AutoDiscGatewayID)
    d.Set("external_id", Gateway.ExternalID)
    d.Set("system_id", Gateway.SystemID)
    
    d.Set("id", Gateway.Identifier())
    d.Set("parent_id", Gateway.ParentID)
    d.Set("parent_type", Gateway.ParentType)
    d.Set("owner", Gateway.Owner)

    d.SetId(Gateway.Identifier())
    
    return nil
}