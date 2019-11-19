package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNSPortInfo() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNSPortInfoRead,
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
            "wireless_ports": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "monitoring_ports": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "ports": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceNSPortInfoRead(d *schema.ResourceData, m interface{}) error {
    filteredNSPortInfos := vspk.NSPortInfosList{}
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
    parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
    filteredNSPortInfos, err = parent.NSPortInfos(fetchFilter)
    if err != nil {
        return err
    }

    NSPortInfo := &vspk.NSPortInfo{}

    if len(filteredNSPortInfos) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNSPortInfos) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NSPortInfo = filteredNSPortInfos[0]

    d.Set("wireless_ports", NSPortInfo.WirelessPorts)
    d.Set("monitoring_ports", NSPortInfo.MonitoringPorts)
    d.Set("ports", NSPortInfo.Ports)
    
    d.Set("id", NSPortInfo.Identifier())
    d.Set("parent_id", NSPortInfo.ParentID)
    d.Set("parent_type", NSPortInfo.ParentType)
    d.Set("owner", NSPortInfo.Owner)

    d.SetId(NSPortInfo.Identifier())
    
    return nil
}