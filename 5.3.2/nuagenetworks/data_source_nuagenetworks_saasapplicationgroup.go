package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceSaaSApplicationGroup() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceSaaSApplicationGroupRead,
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
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
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


func dataSourceSaaSApplicationGroupRead(d *schema.ResourceData, m interface{}) error {
    filteredSaaSApplicationGroups := vspk.SaaSApplicationGroupsList{}
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
    filteredSaaSApplicationGroups, err = parent.SaaSApplicationGroups(fetchFilter)
    if err != nil {
        return err
    }

    SaaSApplicationGroup := &vspk.SaaSApplicationGroup{}

    if len(filteredSaaSApplicationGroups) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredSaaSApplicationGroups) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    SaaSApplicationGroup = filteredSaaSApplicationGroups[0]

    d.Set("name", SaaSApplicationGroup.Name)
    d.Set("last_updated_by", SaaSApplicationGroup.LastUpdatedBy)
    d.Set("description", SaaSApplicationGroup.Description)
    d.Set("entity_scope", SaaSApplicationGroup.EntityScope)
    d.Set("external_id", SaaSApplicationGroup.ExternalID)
    
    d.Set("id", SaaSApplicationGroup.Identifier())
    d.Set("parent_id", SaaSApplicationGroup.ParentID)
    d.Set("parent_type", SaaSApplicationGroup.ParentType)
    d.Set("owner", SaaSApplicationGroup.Owner)

    d.SetId(SaaSApplicationGroup.Identifier())
    
    return nil
}