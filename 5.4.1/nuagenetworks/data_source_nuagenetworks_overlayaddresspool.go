package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceOverlayAddressPool() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceOverlayAddressPoolRead,
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
            "ip_type": &schema.Schema{
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "end_address_range": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "start_address_range": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
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


func dataSourceOverlayAddressPoolRead(d *schema.ResourceData, m interface{}) error {
    filteredOverlayAddressPools := vspk.OverlayAddressPoolsList{}
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
    filteredOverlayAddressPools, err = parent.OverlayAddressPools(fetchFilter)
    if err != nil {
        return err
    }

    OverlayAddressPool := &vspk.OverlayAddressPool{}

    if len(filteredOverlayAddressPools) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredOverlayAddressPools) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    OverlayAddressPool = filteredOverlayAddressPools[0]

    d.Set("ip_type", OverlayAddressPool.IPType)
    d.Set("name", OverlayAddressPool.Name)
    d.Set("last_updated_by", OverlayAddressPool.LastUpdatedBy)
    d.Set("description", OverlayAddressPool.Description)
    d.Set("end_address_range", OverlayAddressPool.EndAddressRange)
    d.Set("entity_scope", OverlayAddressPool.EntityScope)
    d.Set("associated_domain_id", OverlayAddressPool.AssociatedDomainID)
    d.Set("start_address_range", OverlayAddressPool.StartAddressRange)
    d.Set("external_id", OverlayAddressPool.ExternalID)
    
    d.Set("id", OverlayAddressPool.Identifier())
    d.Set("parent_id", OverlayAddressPool.ParentID)
    d.Set("parent_type", OverlayAddressPool.ParentType)
    d.Set("owner", OverlayAddressPool.Owner)

    d.SetId(OverlayAddressPool.Identifier())
    
    return nil
}