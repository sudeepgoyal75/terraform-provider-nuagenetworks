package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.0.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceWirelessPortTemplate() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceWirelessPortTemplateRead,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "generic_config": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "physical_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "wifi_frequency_band": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "wifi_mode": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "port_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "country_code": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "frequency_channel": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceWirelessPortTemplateRead(d *schema.ResourceData, m interface{}) error {
    filteredWirelessPortTemplates := vspk.WirelessPortTemplatesList{}
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

    WirelessPortTemplate := &vspk.WirelessPortTemplate{}

    if len(filteredWirelessPortTemplates) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredWirelessPortTemplates) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    WirelessPortTemplate = filteredWirelessPortTemplates[0]

    d.Set("name", WirelessPortTemplate.Name)
    d.Set("generic_config", WirelessPortTemplate.GenericConfig)
    d.Set("description", WirelessPortTemplate.Description)
    d.Set("physical_name", WirelessPortTemplate.PhysicalName)
    d.Set("wifi_frequency_band", WirelessPortTemplate.WifiFrequencyBand)
    d.Set("wifi_mode", WirelessPortTemplate.WifiMode)
    d.Set("port_type", WirelessPortTemplate.PortType)
    d.Set("country_code", WirelessPortTemplate.CountryCode)
    d.Set("frequency_channel", WirelessPortTemplate.FrequencyChannel)
    
    d.Set("id", WirelessPortTemplate.Identifier())
    d.Set("parent_id", WirelessPortTemplate.ParentID)
    d.Set("parent_type", WirelessPortTemplate.ParentType)
    d.Set("owner", WirelessPortTemplate.Owner)

    d.SetId(WirelessPortTemplate.Identifier())
    
    return nil
}