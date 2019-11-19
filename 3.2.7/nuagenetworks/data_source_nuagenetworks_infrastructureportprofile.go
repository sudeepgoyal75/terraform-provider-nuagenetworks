package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.7"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceInfrastructurePortProfile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceInfrastructurePortProfileRead,
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
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "speed": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "uplink_tag": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mtu": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "duplex": &schema.Schema{
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


func dataSourceInfrastructurePortProfileRead(d *schema.ResourceData, m interface{}) error {
    filteredInfrastructurePortProfiles := vspk.InfrastructurePortProfilesList{}
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
        filteredInfrastructurePortProfiles, err = parent.InfrastructurePortProfiles(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredInfrastructurePortProfiles, err = parent.InfrastructurePortProfiles(fetchFilter)
        if err != nil {
            return err
        }
    }

    InfrastructurePortProfile := &vspk.InfrastructurePortProfile{}

    if len(filteredInfrastructurePortProfiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredInfrastructurePortProfiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    InfrastructurePortProfile = filteredInfrastructurePortProfiles[0]

    d.Set("name", InfrastructurePortProfile.Name)
    d.Set("last_updated_by", InfrastructurePortProfile.LastUpdatedBy)
    d.Set("description", InfrastructurePortProfile.Description)
    d.Set("enterprise_id", InfrastructurePortProfile.EnterpriseID)
    d.Set("entity_scope", InfrastructurePortProfile.EntityScope)
    d.Set("speed", InfrastructurePortProfile.Speed)
    d.Set("uplink_tag", InfrastructurePortProfile.UplinkTag)
    d.Set("mtu", InfrastructurePortProfile.Mtu)
    d.Set("duplex", InfrastructurePortProfile.Duplex)
    d.Set("external_id", InfrastructurePortProfile.ExternalID)
    
    d.Set("id", InfrastructurePortProfile.Identifier())
    d.Set("parent_id", InfrastructurePortProfile.ParentID)
    d.Set("parent_type", InfrastructurePortProfile.ParentType)
    d.Set("owner", InfrastructurePortProfile.Owner)

    d.SetId(InfrastructurePortProfile.Identifier())
    
    return nil
}