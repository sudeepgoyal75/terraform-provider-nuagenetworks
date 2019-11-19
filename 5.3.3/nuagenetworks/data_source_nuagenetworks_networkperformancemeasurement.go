package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNetworkPerformanceMeasurement() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNetworkPerformanceMeasurementRead,
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
            "npm_type": &schema.Schema{
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_performance_monitor_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}


func dataSourceNetworkPerformanceMeasurementRead(d *schema.ResourceData, m interface{}) error {
    filteredNetworkPerformanceMeasurements := vspk.NetworkPerformanceMeasurementsList{}
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
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredNetworkPerformanceMeasurements, err = parent.NetworkPerformanceMeasurements(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredNetworkPerformanceMeasurements, err = parent.NetworkPerformanceMeasurements(fetchFilter)
        if err != nil {
            return err
        }
    }

    NetworkPerformanceMeasurement := &vspk.NetworkPerformanceMeasurement{}

    if len(filteredNetworkPerformanceMeasurements) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNetworkPerformanceMeasurements) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NetworkPerformanceMeasurement = filteredNetworkPerformanceMeasurements[0]

    d.Set("npm_type", NetworkPerformanceMeasurement.NPMType)
    d.Set("name", NetworkPerformanceMeasurement.Name)
    d.Set("last_updated_by", NetworkPerformanceMeasurement.LastUpdatedBy)
    d.Set("read_only", NetworkPerformanceMeasurement.ReadOnly)
    d.Set("description", NetworkPerformanceMeasurement.Description)
    d.Set("entity_scope", NetworkPerformanceMeasurement.EntityScope)
    d.Set("associated_performance_monitor_id", NetworkPerformanceMeasurement.AssociatedPerformanceMonitorID)
    d.Set("external_id", NetworkPerformanceMeasurement.ExternalID)
    
    d.Set("id", NetworkPerformanceMeasurement.Identifier())
    d.Set("parent_id", NetworkPerformanceMeasurement.ParentID)
    d.Set("parent_type", NetworkPerformanceMeasurement.ParentType)
    d.Set("owner", NetworkPerformanceMeasurement.Owner)

    d.SetId(NetworkPerformanceMeasurement.Identifier())
    
    return nil
}