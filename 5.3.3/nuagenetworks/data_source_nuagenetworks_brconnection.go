package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceBRConnection() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceBRConnectionRead,
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
            "dns_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dns_address_v6": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_v6": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address_family": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address_v6": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "advertisement_criteria": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "inherited": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mode": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "uplink_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_vlan_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan"},
            },
            "parent_vlan": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan_template"},
            },
        },
    }
}


func dataSourceBRConnectionRead(d *schema.ResourceData, m interface{}) error {
    filteredBRConnections := vspk.BRConnectionsList{}
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
    if attr, ok := d.GetOk("parent_vlan_template"); ok {
        parent := &vspk.VLANTemplate{ID: attr.(string)}
        filteredBRConnections, err = parent.BRConnections(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vlan"); ok {
        parent := &vspk.VLAN{ID: attr.(string)}
        filteredBRConnections, err = parent.BRConnections(fetchFilter)
        if err != nil {
            return err
        }
    }

    BRConnection := &vspk.BRConnection{}

    if len(filteredBRConnections) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredBRConnections) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    BRConnection = filteredBRConnections[0]

    d.Set("dns_address", BRConnection.DNSAddress)
    d.Set("dns_address_v6", BRConnection.DNSAddressV6)
    d.Set("last_updated_by", BRConnection.LastUpdatedBy)
    d.Set("gateway", BRConnection.Gateway)
    d.Set("gateway_v6", BRConnection.GatewayV6)
    d.Set("address", BRConnection.Address)
    d.Set("address_family", BRConnection.AddressFamily)
    d.Set("address_v6", BRConnection.AddressV6)
    d.Set("advertisement_criteria", BRConnection.AdvertisementCriteria)
    d.Set("netmask", BRConnection.Netmask)
    d.Set("inherited", BRConnection.Inherited)
    d.Set("entity_scope", BRConnection.EntityScope)
    d.Set("mode", BRConnection.Mode)
    d.Set("uplink_id", BRConnection.UplinkID)
    d.Set("external_id", BRConnection.ExternalID)
    
    d.Set("id", BRConnection.Identifier())
    d.Set("parent_id", BRConnection.ParentID)
    d.Set("parent_type", BRConnection.ParentType)
    d.Set("owner", BRConnection.Owner)

    d.SetId(BRConnection.Identifier())
    
    return nil
}