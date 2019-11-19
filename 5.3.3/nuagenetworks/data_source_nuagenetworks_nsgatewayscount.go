package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNSGatewaysCount() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNSGatewaysCountRead,
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
            "active_nsg_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "alarmed_nsg_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "inactive_nsg_count": &schema.Schema{
                Type:     schema.TypeInt,
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
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceNSGatewaysCountRead(d *schema.ResourceData, m interface{}) error {
    filteredNSGatewaysCounts := vspk.NSGatewaysCountsList{}
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
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    filteredNSGatewaysCounts, err = parent.NSGatewaysCounts(fetchFilter)
    if err != nil {
        return err
    }

    NSGatewaysCount := &vspk.NSGatewaysCount{}

    if len(filteredNSGatewaysCounts) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNSGatewaysCounts) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NSGatewaysCount = filteredNSGatewaysCounts[0]

    d.Set("active_nsg_count", NSGatewaysCount.ActiveNSGCount)
    d.Set("alarmed_nsg_count", NSGatewaysCount.AlarmedNSGCount)
    d.Set("inactive_nsg_count", NSGatewaysCount.InactiveNSGCount)
    d.Set("entity_scope", NSGatewaysCount.EntityScope)
    d.Set("external_id", NSGatewaysCount.ExternalID)
    
    d.Set("id", NSGatewaysCount.Identifier())
    d.Set("parent_id", NSGatewaysCount.ParentID)
    d.Set("parent_type", NSGatewaysCount.ParentType)
    d.Set("owner", NSGatewaysCount.Owner)

    d.SetId(NSGatewaysCount.Identifier())
    
    return nil
}