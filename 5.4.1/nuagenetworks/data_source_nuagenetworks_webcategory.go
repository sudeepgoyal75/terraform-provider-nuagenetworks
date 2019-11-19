package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceWebCategory() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceWebCategoryRead,
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
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_web_domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_web_domain_name"},
            },
        },
    }
}


func dataSourceWebCategoryRead(d *schema.ResourceData, m interface{}) error {
    filteredWebCategories := vspk.WebCategoriesList{}
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
    if attr, ok := d.GetOk("parent_web_domain_name"); ok {
        parent := &vspk.WebDomainName{ID: attr.(string)}
        filteredWebCategories, err = parent.WebCategories(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredWebCategories, err = parent.WebCategories(fetchFilter)
        if err != nil {
            return err
        }
    }

    WebCategory := &vspk.WebCategory{}

    if len(filteredWebCategories) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredWebCategories) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    WebCategory = filteredWebCategories[0]

    d.Set("name", WebCategory.Name)
    d.Set("last_updated_by", WebCategory.LastUpdatedBy)
    d.Set("default_category", WebCategory.DefaultCategory)
    d.Set("description", WebCategory.Description)
    d.Set("entity_scope", WebCategory.EntityScope)
    d.Set("external_id", WebCategory.ExternalID)
    d.Set("type", WebCategory.Type)
    
    d.Set("id", WebCategory.Identifier())
    d.Set("parent_id", WebCategory.ParentID)
    d.Set("parent_type", WebCategory.ParentType)
    d.Set("owner", WebCategory.Owner)

    d.SetId(WebCategory.Identifier())
    
    return nil
}