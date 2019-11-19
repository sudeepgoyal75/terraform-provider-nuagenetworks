package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.9"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNSPortStaticConfiguration() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNSPortStaticConfigurationRead,
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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enabled": &schema.Schema{
                Type:     schema.TypeBool,
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
            "parent_ns_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}


func dataSourceNSPortStaticConfigurationRead(d *schema.ResourceData, m interface{}) error {
    filteredNSPortStaticConfigurations := vspk.NSPortStaticConfigurationsList{}
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
    if attr, ok := d.GetOk("parent_ns_port"); ok {
        parent := &vspk.NSPort{ID: attr.(string)}
        filteredNSPortStaticConfigurations, err = parent.NSPortStaticConfigurations(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredNSPortStaticConfigurations, err = parent.NSPortStaticConfigurations(fetchFilter)
        if err != nil {
            return err
        }
    }

    NSPortStaticConfiguration := &vspk.NSPortStaticConfiguration{}

    if len(filteredNSPortStaticConfigurations) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNSPortStaticConfigurations) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NSPortStaticConfiguration = filteredNSPortStaticConfigurations[0]

    d.Set("dns_address", NSPortStaticConfiguration.DNSAddress)
    d.Set("last_updated_by", NSPortStaticConfiguration.LastUpdatedBy)
    d.Set("gateway", NSPortStaticConfiguration.Gateway)
    d.Set("address", NSPortStaticConfiguration.Address)
    d.Set("netmask", NSPortStaticConfiguration.Netmask)
    d.Set("enabled", NSPortStaticConfiguration.Enabled)
    d.Set("entity_scope", NSPortStaticConfiguration.EntityScope)
    d.Set("external_id", NSPortStaticConfiguration.ExternalID)
    
    d.Set("id", NSPortStaticConfiguration.Identifier())
    d.Set("parent_id", NSPortStaticConfiguration.ParentID)
    d.Set("parent_type", NSPortStaticConfiguration.ParentType)
    d.Set("owner", NSPortStaticConfiguration.Owner)

    d.SetId(NSPortStaticConfiguration.Identifier())
    
    return nil
}