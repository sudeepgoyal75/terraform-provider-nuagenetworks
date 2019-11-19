package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceL7applicationsignature() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceL7applicationsignatureRead,
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
            "category": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "readonly": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dictionary_version": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "guidstring": &schema.Schema{
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


func dataSourceL7applicationsignatureRead(d *schema.ResourceData, m interface{}) error {
    filteredL7applicationsignatures := vspk.L7applicationsignaturesList{}
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
    filteredL7applicationsignatures, err = parent.L7applicationsignatures(fetchFilter)
    if err != nil {
        return err
    }

    L7applicationsignature := &vspk.L7applicationsignature{}

    if len(filteredL7applicationsignatures) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredL7applicationsignatures) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    L7applicationsignature = filteredL7applicationsignatures[0]

    d.Set("name", L7applicationsignature.Name)
    d.Set("category", L7applicationsignature.Category)
    d.Set("readonly", L7applicationsignature.Readonly)
    d.Set("description", L7applicationsignature.Description)
    d.Set("dictionary_version", L7applicationsignature.DictionaryVersion)
    d.Set("guidstring", L7applicationsignature.Guidstring)
    
    d.Set("id", L7applicationsignature.Identifier())
    d.Set("parent_id", L7applicationsignature.ParentID)
    d.Set("parent_type", L7applicationsignature.ParentType)
    d.Set("owner", L7applicationsignature.Owner)

    d.SetId(L7applicationsignature.Identifier())
    
    return nil
}