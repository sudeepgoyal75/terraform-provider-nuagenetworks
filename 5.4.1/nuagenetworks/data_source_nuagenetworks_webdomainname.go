package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceWebDomainName() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceWebDomainNameRead,
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
                ConflictsWith: []string{"parent_web_category"},
            },
            "parent_web_category": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
        },
    }
}


func dataSourceWebDomainNameRead(d *schema.ResourceData, m interface{}) error {
    filteredWebDomainNames := vspk.WebDomainNamesList{}
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
        filteredWebDomainNames, err = parent.WebDomainNames(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_web_category"); ok {
        parent := &vspk.WebCategory{ID: attr.(string)}
        filteredWebDomainNames, err = parent.WebDomainNames(fetchFilter)
        if err != nil {
            return err
        }
    }

    WebDomainName := &vspk.WebDomainName{}

    if len(filteredWebDomainNames) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredWebDomainNames) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    WebDomainName = filteredWebDomainNames[0]

    d.Set("name", WebDomainName.Name)
    d.Set("last_updated_by", WebDomainName.LastUpdatedBy)
    d.Set("entity_scope", WebDomainName.EntityScope)
    d.Set("external_id", WebDomainName.ExternalID)
    
    d.Set("id", WebDomainName.Identifier())
    d.Set("parent_id", WebDomainName.ParentID)
    d.Set("parent_type", WebDomainName.ParentType)
    d.Set("owner", WebDomainName.Owner)

    d.SetId(WebDomainName.Identifier())
    
    return nil
}