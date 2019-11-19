package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourcePolicyGroupCategory() *schema.Resource {
    return &schema.Resource{
        Read: dataSourcePolicyGroupCategoryRead,
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
            "default_category": &schema.Schema{
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
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_policy_group"},
            },
            "parent_policy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
        },
    }
}


func dataSourcePolicyGroupCategoryRead(d *schema.ResourceData, m interface{}) error {
    filteredPolicyGroupCategories := vspk.PolicyGroupCategoriesList{}
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
        filteredPolicyGroupCategories, err = parent.PolicyGroupCategories(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_policy_group"); ok {
        parent := &vspk.PolicyGroup{ID: attr.(string)}
        filteredPolicyGroupCategories, err = parent.PolicyGroupCategories(fetchFilter)
        if err != nil {
            return err
        }
    }

    PolicyGroupCategory := &vspk.PolicyGroupCategory{}

    if len(filteredPolicyGroupCategories) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredPolicyGroupCategories) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    PolicyGroupCategory = filteredPolicyGroupCategories[0]

    d.Set("name", PolicyGroupCategory.Name)
    d.Set("last_updated_by", PolicyGroupCategory.LastUpdatedBy)
    d.Set("default_category", PolicyGroupCategory.DefaultCategory)
    d.Set("description", PolicyGroupCategory.Description)
    d.Set("entity_scope", PolicyGroupCategory.EntityScope)
    d.Set("external_id", PolicyGroupCategory.ExternalID)
    
    d.Set("id", PolicyGroupCategory.Identifier())
    d.Set("parent_id", PolicyGroupCategory.ParentID)
    d.Set("parent_type", PolicyGroupCategory.ParentType)
    d.Set("owner", PolicyGroupCategory.Owner)

    d.SetId(PolicyGroupCategory.Identifier())
    
    return nil
}