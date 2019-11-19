package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceForwardingPathList() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceForwardingPathListRead,
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
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceForwardingPathListRead(d *schema.ResourceData, m interface{}) error {
    filteredForwardingPathLists := vspk.ForwardingPathListsList{}
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
    parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
    filteredForwardingPathLists, err = parent.ForwardingPathLists(fetchFilter)
    if err != nil {
        return err
    }

    ForwardingPathList := &vspk.ForwardingPathList{}

    if len(filteredForwardingPathLists) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredForwardingPathLists) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    ForwardingPathList = filteredForwardingPathLists[0]

    d.Set("name", ForwardingPathList.Name)
    d.Set("last_updated_by", ForwardingPathList.LastUpdatedBy)
    d.Set("description", ForwardingPathList.Description)
    d.Set("entity_scope", ForwardingPathList.EntityScope)
    d.Set("external_id", ForwardingPathList.ExternalID)
    
    d.Set("id", ForwardingPathList.Identifier())
    d.Set("parent_id", ForwardingPathList.ParentID)
    d.Set("parent_type", ForwardingPathList.ParentType)
    d.Set("owner", ForwardingPathList.Owner)

    d.SetId(ForwardingPathList.Identifier())
    
    return nil
}