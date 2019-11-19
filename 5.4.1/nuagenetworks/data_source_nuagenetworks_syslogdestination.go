package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceSyslogDestination() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceSyslogDestinationRead,
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
            "ip_address": &schema.Schema{
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "port": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "type": &schema.Schema{
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


func dataSourceSyslogDestinationRead(d *schema.ResourceData, m interface{}) error {
    filteredSyslogDestinations := vspk.SyslogDestinationsList{}
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
    filteredSyslogDestinations, err = parent.SyslogDestinations(fetchFilter)
    if err != nil {
        return err
    }

    SyslogDestination := &vspk.SyslogDestination{}

    if len(filteredSyslogDestinations) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredSyslogDestinations) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    SyslogDestination = filteredSyslogDestinations[0]

    d.Set("ip_address", SyslogDestination.IPAddress)
    d.Set("ip_type", SyslogDestination.IPType)
    d.Set("name", SyslogDestination.Name)
    d.Set("description", SyslogDestination.Description)
    d.Set("port", SyslogDestination.Port)
    d.Set("type", SyslogDestination.Type)
    
    d.Set("id", SyslogDestination.Identifier())
    d.Set("parent_id", SyslogDestination.ParentID)
    d.Set("parent_type", SyslogDestination.ParentType)
    d.Set("owner", SyslogDestination.Owner)

    d.SetId(SyslogDestination.Identifier())
    
    return nil
}