package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceMonitorscope() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceMonitorscopeRead,
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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "read_only": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "destination_nsgs": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "allow_all_destination_nsgs": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "allow_all_source_nsgs": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_nsgs": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_application": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_network_performance_measurement"},
            },
            "parent_network_performance_measurement": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_application"},
            },
        },
    }
}


func dataSourceMonitorscopeRead(d *schema.ResourceData, m interface{}) error {
    filteredMonitorscopes := vspk.MonitorscopesList{}
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
    if attr, ok := d.GetOk("parent_application"); ok {
        parent := &vspk.Application{ID: attr.(string)}
        filteredMonitorscopes, err = parent.Monitorscopes(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_network_performance_measurement"); ok {
        parent := &vspk.NetworkPerformanceMeasurement{ID: attr.(string)}
        filteredMonitorscopes, err = parent.Monitorscopes(fetchFilter)
        if err != nil {
            return err
        }
    }

    Monitorscope := &vspk.Monitorscope{}

    if len(filteredMonitorscopes) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredMonitorscopes) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Monitorscope = filteredMonitorscopes[0]

    d.Set("name", Monitorscope.Name)
    d.Set("last_updated_by", Monitorscope.LastUpdatedBy)
    d.Set("read_only", Monitorscope.ReadOnly)
    d.Set("destination_nsgs", Monitorscope.DestinationNSGs)
    d.Set("allow_all_destination_nsgs", Monitorscope.AllowAllDestinationNSGs)
    d.Set("allow_all_source_nsgs", Monitorscope.AllowAllSourceNSGs)
    d.Set("entity_scope", Monitorscope.EntityScope)
    d.Set("source_nsgs", Monitorscope.SourceNSGs)
    d.Set("external_id", Monitorscope.ExternalID)
    
    d.Set("id", Monitorscope.Identifier())
    d.Set("parent_id", Monitorscope.ParentID)
    d.Set("parent_type", Monitorscope.ParentType)
    d.Set("owner", Monitorscope.Owner)

    d.SetId(Monitorscope.Identifier())
    
    return nil
}