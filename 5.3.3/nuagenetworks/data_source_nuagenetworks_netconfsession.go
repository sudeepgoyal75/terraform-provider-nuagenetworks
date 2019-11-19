package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNetconfSession() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNetconfSessionRead,
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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_gateway_name": &schema.Schema{
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
            "parent_netconf_manager": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceNetconfSessionRead(d *schema.ResourceData, m interface{}) error {
    filteredNetconfSessions := vspk.NetconfSessionsList{}
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
    parent := &vspk.NetconfManager{ID: d.Get("parent_netconf_manager").(string)}
    filteredNetconfSessions, err = parent.NetconfSessions(fetchFilter)
    if err != nil {
        return err
    }

    NetconfSession := &vspk.NetconfSession{}

    if len(filteredNetconfSessions) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNetconfSessions) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NetconfSession = filteredNetconfSessions[0]

    d.Set("last_updated_by", NetconfSession.LastUpdatedBy)
    d.Set("entity_scope", NetconfSession.EntityScope)
    d.Set("associated_gateway_id", NetconfSession.AssociatedGatewayID)
    d.Set("associated_gateway_name", NetconfSession.AssociatedGatewayName)
    d.Set("status", NetconfSession.Status)
    d.Set("external_id", NetconfSession.ExternalID)
    
    d.Set("id", NetconfSession.Identifier())
    d.Set("parent_id", NetconfSession.ParentID)
    d.Set("parent_type", NetconfSession.ParentType)
    d.Set("owner", NetconfSession.Owner)

    d.SetId(NetconfSession.Identifier())
    
    return nil
}