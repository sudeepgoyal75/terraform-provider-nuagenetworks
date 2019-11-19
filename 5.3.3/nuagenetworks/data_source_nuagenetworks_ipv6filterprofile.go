package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceIPv6FilterProfile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceIPv6FilterProfileRead,
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
            "assoc_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_redundancy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_gateway"},
            },
            "parent_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_redundancy_group"},
            },
        },
    }
}


func dataSourceIPv6FilterProfileRead(d *schema.ResourceData, m interface{}) error {
    filteredIPv6FilterProfiles := vspk.IPv6FilterProfilesList{}
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
    if attr, ok := d.GetOk("parent_redundancy_group"); ok {
        parent := &vspk.RedundancyGroup{ID: attr.(string)}
        filteredIPv6FilterProfiles, err = parent.IPv6FilterProfiles(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        filteredIPv6FilterProfiles, err = parent.IPv6FilterProfiles(fetchFilter)
        if err != nil {
            return err
        }
    }

    IPv6FilterProfile := &vspk.IPv6FilterProfile{}

    if len(filteredIPv6FilterProfiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredIPv6FilterProfiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    IPv6FilterProfile = filteredIPv6FilterProfiles[0]

    d.Set("name", IPv6FilterProfile.Name)
    d.Set("last_updated_by", IPv6FilterProfile.LastUpdatedBy)
    d.Set("description", IPv6FilterProfile.Description)
    d.Set("entity_scope", IPv6FilterProfile.EntityScope)
    d.Set("assoc_entity_type", IPv6FilterProfile.AssocEntityType)
    d.Set("external_id", IPv6FilterProfile.ExternalID)
    
    d.Set("id", IPv6FilterProfile.Identifier())
    d.Set("parent_id", IPv6FilterProfile.ParentID)
    d.Set("parent_type", IPv6FilterProfile.ParentType)
    d.Set("owner", IPv6FilterProfile.Owner)

    d.SetId(IPv6FilterProfile.Identifier())
    
    return nil
}