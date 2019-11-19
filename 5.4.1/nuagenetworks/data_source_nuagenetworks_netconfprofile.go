package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNetconfProfile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNetconfProfileRead,
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
            "password": &schema.Schema{
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
            "port": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "user_name": &schema.Schema{
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
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}


func dataSourceNetconfProfileRead(d *schema.ResourceData, m interface{}) error {
    filteredNetconfProfiles := vspk.NetconfProfilesList{}
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
        filteredNetconfProfiles, err = parent.NetconfProfiles(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredNetconfProfiles, err = parent.NetconfProfiles(fetchFilter)
        if err != nil {
            return err
        }
    }

    NetconfProfile := &vspk.NetconfProfile{}

    if len(filteredNetconfProfiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNetconfProfiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NetconfProfile = filteredNetconfProfiles[0]

    d.Set("name", NetconfProfile.Name)
    d.Set("password", NetconfProfile.Password)
    d.Set("last_updated_by", NetconfProfile.LastUpdatedBy)
    d.Set("description", NetconfProfile.Description)
    d.Set("entity_scope", NetconfProfile.EntityScope)
    d.Set("port", NetconfProfile.Port)
    d.Set("user_name", NetconfProfile.UserName)
    d.Set("assoc_entity_type", NetconfProfile.AssocEntityType)
    d.Set("external_id", NetconfProfile.ExternalID)
    
    d.Set("id", NetconfProfile.Identifier())
    d.Set("parent_id", NetconfProfile.ParentID)
    d.Set("parent_type", NetconfProfile.ParentType)
    d.Set("owner", NetconfProfile.Owner)

    d.SetId(NetconfProfile.Identifier())
    
    return nil
}