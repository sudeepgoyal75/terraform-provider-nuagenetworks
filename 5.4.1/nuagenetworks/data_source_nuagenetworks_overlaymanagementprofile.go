package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceOverlayManagementProfile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceOverlayManagementProfileRead,
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
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceOverlayManagementProfileRead(d *schema.ResourceData, m interface{}) error {
    filteredOverlayManagementProfiles := vspk.OverlayManagementProfilesList{}
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
    filteredOverlayManagementProfiles, err = parent.OverlayManagementProfiles(fetchFilter)
    if err != nil {
        return err
    }

    OverlayManagementProfile := &vspk.OverlayManagementProfile{}

    if len(filteredOverlayManagementProfiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredOverlayManagementProfiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    OverlayManagementProfile = filteredOverlayManagementProfiles[0]

    d.Set("name", OverlayManagementProfile.Name)
    d.Set("description", OverlayManagementProfile.Description)
    
    d.Set("id", OverlayManagementProfile.Identifier())
    d.Set("parent_id", OverlayManagementProfile.ParentID)
    d.Set("parent_type", OverlayManagementProfile.ParentType)
    d.Set("owner", OverlayManagementProfile.Owner)

    d.SetId(OverlayManagementProfile.Identifier())
    
    return nil
}