package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourcePolicyObjectGroup() *schema.Resource {
    return &schema.Resource{
        Read: dataSourcePolicyObjectGroupRead,
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
            "type": &schema.Schema{
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


func dataSourcePolicyObjectGroupRead(d *schema.ResourceData, m interface{}) error {
    filteredPolicyObjectGroups := vspk.PolicyObjectGroupsList{}
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
        filteredPolicyObjectGroups, err = parent.PolicyObjectGroups(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredPolicyObjectGroups, err = parent.PolicyObjectGroups(fetchFilter)
        if err != nil {
            return err
        }
    }

    PolicyObjectGroup := &vspk.PolicyObjectGroup{}

    if len(filteredPolicyObjectGroups) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredPolicyObjectGroups) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    PolicyObjectGroup = filteredPolicyObjectGroups[0]

    d.Set("name", PolicyObjectGroup.Name)
    d.Set("last_updated_by", PolicyObjectGroup.LastUpdatedBy)
    d.Set("description", PolicyObjectGroup.Description)
    d.Set("entity_scope", PolicyObjectGroup.EntityScope)
    d.Set("external_id", PolicyObjectGroup.ExternalID)
    d.Set("type", PolicyObjectGroup.Type)
    
    d.Set("id", PolicyObjectGroup.Identifier())
    d.Set("parent_id", PolicyObjectGroup.ParentID)
    d.Set("parent_type", PolicyObjectGroup.ParentType)
    d.Set("owner", PolicyObjectGroup.Owner)

    d.SetId(PolicyObjectGroup.Identifier())
    
    return nil
}