package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNetconfManager() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNetconfManagerRead,
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
            "release": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_vsp": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceNetconfManagerRead(d *schema.ResourceData, m interface{}) error {
    filteredNetconfManagers := vspk.NetconfManagersList{}
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
    parent := &vspk.VSP{ID: d.Get("parent_vsp").(string)}
    filteredNetconfManagers, err = parent.NetconfManagers(fetchFilter)
    if err != nil {
        return err
    }

    NetconfManager := &vspk.NetconfManager{}

    if len(filteredNetconfManagers) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNetconfManagers) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NetconfManager = filteredNetconfManagers[0]

    d.Set("name", NetconfManager.Name)
    d.Set("release", NetconfManager.Release)
    d.Set("status", NetconfManager.Status)
    
    d.Set("id", NetconfManager.Identifier())
    d.Set("parent_id", NetconfManager.ParentID)
    d.Set("parent_type", NetconfManager.ParentType)
    d.Set("owner", NetconfManager.Owner)

    d.SetId(NetconfManager.Identifier())
    
    return nil
}