package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceAutoDiscoverHypervisorFromDatacenter() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceAutoDiscoverHypervisorFromDatacenterRead,
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
            "network_list": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "assoc_vcenter_data_center_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "hypervisor_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceAutoDiscoverHypervisorFromDatacenterRead(d *schema.ResourceData, m interface{}) error {
    filteredAutoDiscoverHypervisorFromDatacenters := vspk.AutoDiscoverHypervisorFromDatacentersList{}
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

    AutoDiscoverHypervisorFromDatacenter := &vspk.AutoDiscoverHypervisorFromDatacenter{}

    if len(filteredAutoDiscoverHypervisorFromDatacenters) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredAutoDiscoverHypervisorFromDatacenters) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    AutoDiscoverHypervisorFromDatacenter = filteredAutoDiscoverHypervisorFromDatacenters[0]

    d.Set("network_list", AutoDiscoverHypervisorFromDatacenter.NetworkList)
    d.Set("assoc_vcenter_data_center_id", AutoDiscoverHypervisorFromDatacenter.AssocVCenterDataCenterId)
    d.Set("hypervisor_ip", AutoDiscoverHypervisorFromDatacenter.HypervisorIP)
    
    d.Set("id", AutoDiscoverHypervisorFromDatacenter.Identifier())
    d.Set("parent_id", AutoDiscoverHypervisorFromDatacenter.ParentID)
    d.Set("parent_type", AutoDiscoverHypervisorFromDatacenter.ParentType)
    d.Set("owner", AutoDiscoverHypervisorFromDatacenter.Owner)

    d.SetId(AutoDiscoverHypervisorFromDatacenter.Identifier())
    
    return nil
}