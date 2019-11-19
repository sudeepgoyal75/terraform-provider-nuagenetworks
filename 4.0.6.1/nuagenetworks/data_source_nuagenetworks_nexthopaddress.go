package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.6.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNextHopAddress() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNextHopAddressRead,
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
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_link": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceNextHopAddressRead(d *schema.ResourceData, m interface{}) error {
    filteredNextHopAddress := vspk.NextHopAddressList{}
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
    parent := &vspk.Link{ID: d.Get("parent_link").(string)}
    filteredNextHopAddress, err = parent.NextHopAddress(fetchFilter)
    if err != nil {
        return err
    }

    NextHopAddress := &vspk.NextHopAddress{}

    if len(filteredNextHopAddress) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNextHopAddress) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NextHopAddress = filteredNextHopAddress[0]

    d.Set("address", NextHopAddress.Address)
    d.Set("route_distinguisher", NextHopAddress.RouteDistinguisher)
    d.Set("type", NextHopAddress.Type)
    
    d.Set("id", NextHopAddress.Identifier())
    d.Set("parent_id", NextHopAddress.ParentID)
    d.Set("parent_type", NextHopAddress.ParentType)
    d.Set("owner", NextHopAddress.Owner)

    d.SetId(NextHopAddress.Identifier())
    
    return nil
}