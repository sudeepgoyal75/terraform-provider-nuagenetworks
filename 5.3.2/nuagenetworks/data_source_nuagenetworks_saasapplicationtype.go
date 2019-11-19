package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceSaaSApplicationType() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceSaaSApplicationTypeRead,
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
            "read_only": &schema.Schema{
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
                ConflictsWith: []string{"parent_saa_s_application_group"},
            },
            "parent_saa_s_application_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
        },
    }
}


func dataSourceSaaSApplicationTypeRead(d *schema.ResourceData, m interface{}) error {
    filteredSaaSApplicationTypes := vspk.SaaSApplicationTypesList{}
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
        filteredSaaSApplicationTypes, err = parent.SaaSApplicationTypes(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_saa_s_application_group"); ok {
        parent := &vspk.SaaSApplicationGroup{ID: attr.(string)}
        filteredSaaSApplicationTypes, err = parent.SaaSApplicationTypes(fetchFilter)
        if err != nil {
            return err
        }
    }

    SaaSApplicationType := &vspk.SaaSApplicationType{}

    if len(filteredSaaSApplicationTypes) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredSaaSApplicationTypes) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    SaaSApplicationType = filteredSaaSApplicationTypes[0]

    d.Set("name", SaaSApplicationType.Name)
    d.Set("last_updated_by", SaaSApplicationType.LastUpdatedBy)
    d.Set("read_only", SaaSApplicationType.ReadOnly)
    d.Set("description", SaaSApplicationType.Description)
    d.Set("entity_scope", SaaSApplicationType.EntityScope)
    d.Set("external_id", SaaSApplicationType.ExternalID)
    
    d.Set("id", SaaSApplicationType.Identifier())
    d.Set("parent_id", SaaSApplicationType.ParentID)
    d.Set("parent_type", SaaSApplicationType.ParentType)
    d.Set("owner", SaaSApplicationType.Owner)

    d.SetId(SaaSApplicationType.Identifier())
    
    return nil
}