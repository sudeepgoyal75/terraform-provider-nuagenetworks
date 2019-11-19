package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNSGatewayMonitor() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNSGatewayMonitorRead,
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
            "vrsinfo": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vscs": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "nsginfo": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nsgstate": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nsgsummary": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceNSGatewayMonitorRead(d *schema.ResourceData, m interface{}) error {
    filteredNSGatewayMonitors := vspk.NSGatewayMonitorsList{}
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
    filteredNSGatewayMonitors, err = parent.NSGatewayMonitors(fetchFilter)
    if err != nil {
        return err
    }

    NSGatewayMonitor := &vspk.NSGatewayMonitor{}

    if len(filteredNSGatewayMonitors) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNSGatewayMonitors) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NSGatewayMonitor = filteredNSGatewayMonitors[0]

    d.Set("vrsinfo", NSGatewayMonitor.Vrsinfo)
    d.Set("vscs", NSGatewayMonitor.Vscs)
    d.Set("nsginfo", NSGatewayMonitor.Nsginfo)
    d.Set("nsgstate", NSGatewayMonitor.Nsgstate)
    d.Set("nsgsummary", NSGatewayMonitor.Nsgsummary)
    
    d.Set("id", NSGatewayMonitor.Identifier())
    d.Set("parent_id", NSGatewayMonitor.ParentID)
    d.Set("parent_type", NSGatewayMonitor.ParentType)
    d.Set("owner", NSGatewayMonitor.Owner)

    d.SetId(NSGatewayMonitor.Identifier())
    
    return nil
}