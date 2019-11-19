package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceSAPIngressQoSProfile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceSAPIngressQoSProfileRead,
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


func dataSourceSAPIngressQoSProfileRead(d *schema.ResourceData, m interface{}) error {
    filteredSAPIngressQoSProfiles := vspk.SAPIngressQoSProfilesList{}
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
        filteredSAPIngressQoSProfiles, err = parent.SAPIngressQoSProfiles(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        filteredSAPIngressQoSProfiles, err = parent.SAPIngressQoSProfiles(fetchFilter)
        if err != nil {
            return err
        }
    }

    SAPIngressQoSProfile := &vspk.SAPIngressQoSProfile{}

    if len(filteredSAPIngressQoSProfiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredSAPIngressQoSProfiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    SAPIngressQoSProfile = filteredSAPIngressQoSProfiles[0]

    d.Set("name", SAPIngressQoSProfile.Name)
    d.Set("description", SAPIngressQoSProfile.Description)
    
    d.Set("id", SAPIngressQoSProfile.Identifier())
    d.Set("parent_id", SAPIngressQoSProfile.ParentID)
    d.Set("parent_type", SAPIngressQoSProfile.ParentType)
    d.Set("owner", SAPIngressQoSProfile.Owner)

    d.SetId(SAPIngressQoSProfile.Identifier())
    
    return nil
}