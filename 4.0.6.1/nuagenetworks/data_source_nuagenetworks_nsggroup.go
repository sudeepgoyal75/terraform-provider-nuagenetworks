package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.6.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNSGGroup() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNSGGroupRead,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_nsgs": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}


func dataSourceNSGGroupRead(d *schema.ResourceData, m interface{}) error {
    filteredNSGGroups := vspk.NSGGroupsList{}
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
        filteredNSGGroups, err = parent.NSGGroups(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredNSGGroups, err = parent.NSGGroups(fetchFilter)
        if err != nil {
            return err
        }
    }

    NSGGroup := &vspk.NSGGroup{}

    if len(filteredNSGGroups) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNSGGroups) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NSGGroup = filteredNSGGroups[0]

    d.Set("name", NSGGroup.Name)
    d.Set("description", NSGGroup.Description)
    d.Set("associated_nsgs", NSGGroup.AssociatedNSGs)
    
    d.Set("id", NSGGroup.Identifier())
    d.Set("parent_id", NSGGroup.ParentID)
    d.Set("parent_type", NSGGroup.ParentType)
    d.Set("owner", NSGGroup.Owner)

    d.SetId(NSGGroup.Identifier())
    
    return nil
}