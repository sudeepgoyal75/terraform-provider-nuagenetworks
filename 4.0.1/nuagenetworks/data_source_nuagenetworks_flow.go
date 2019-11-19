package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceFlow() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceFlowRead,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_tier_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "metadata": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "origin_tier_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_app": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceFlowRead(d *schema.ResourceData, m interface{}) error {
    filteredFlows := vspk.FlowsList{}
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
    parent := &vspk.App{ID: d.Get("parent_app").(string)}
    filteredFlows, err = parent.Flows(fetchFilter)
    if err != nil {
        return err
    }

    Flow := &vspk.Flow{}

    if len(filteredFlows) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredFlows) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Flow = filteredFlows[0]

    d.Set("name", Flow.Name)
    d.Set("last_updated_by", Flow.LastUpdatedBy)
    d.Set("description", Flow.Description)
    d.Set("destination_tier_id", Flow.DestinationTierID)
    d.Set("metadata", Flow.Metadata)
    d.Set("entity_scope", Flow.EntityScope)
    d.Set("origin_tier_id", Flow.OriginTierID)
    d.Set("external_id", Flow.ExternalID)
    
    d.Set("id", Flow.Identifier())
    d.Set("parent_id", Flow.ParentID)
    d.Set("parent_type", Flow.ParentType)
    d.Set("owner", Flow.Owner)

    d.SetId(Flow.Identifier())
    
    return nil
}