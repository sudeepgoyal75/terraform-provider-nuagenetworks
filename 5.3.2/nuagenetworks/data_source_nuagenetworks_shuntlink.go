package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceShuntLink() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceShuntLinkRead,
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
            "vlan_peer1_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vlan_peer2_id": &schema.Schema{
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
            "gateway_peer1_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_peer2_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "peer1_ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "peer1_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "peer2_ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "peer2_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_ns_redundant_gateway_group": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceShuntLinkRead(d *schema.ResourceData, m interface{}) error {
    filteredShuntLinks := vspk.ShuntLinksList{}
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
    parent := &vspk.NSRedundantGatewayGroup{ID: d.Get("parent_ns_redundant_gateway_group").(string)}
    filteredShuntLinks, err = parent.ShuntLinks(fetchFilter)
    if err != nil {
        return err
    }

    ShuntLink := &vspk.ShuntLink{}

    if len(filteredShuntLinks) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredShuntLinks) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    ShuntLink = filteredShuntLinks[0]

    d.Set("vlan_peer1_id", ShuntLink.VLANPeer1ID)
    d.Set("vlan_peer2_id", ShuntLink.VLANPeer2ID)
    d.Set("name", ShuntLink.Name)
    d.Set("last_updated_by", ShuntLink.LastUpdatedBy)
    d.Set("gateway_peer1_id", ShuntLink.GatewayPeer1ID)
    d.Set("gateway_peer2_id", ShuntLink.GatewayPeer2ID)
    d.Set("peer1_ip_address", ShuntLink.Peer1IPAddress)
    d.Set("peer1_subnet", ShuntLink.Peer1Subnet)
    d.Set("peer2_ip_address", ShuntLink.Peer2IPAddress)
    d.Set("peer2_subnet", ShuntLink.Peer2Subnet)
    d.Set("description", ShuntLink.Description)
    d.Set("entity_scope", ShuntLink.EntityScope)
    d.Set("external_id", ShuntLink.ExternalID)
    
    d.Set("id", ShuntLink.Identifier())
    d.Set("parent_id", ShuntLink.ParentID)
    d.Set("parent_type", ShuntLink.ParentType)
    d.Set("owner", ShuntLink.Owner)

    d.SetId(ShuntLink.Identifier())
    
    return nil
}