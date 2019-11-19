package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceEndPoint() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceEndPointRead,
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
            "parent_external_service": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceEndPointRead(d *schema.ResourceData, m interface{}) error {
    filteredEndPoints := vspk.EndPointsList{}
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
    parent := &vspk.ExternalService{ID: d.Get("parent_external_service").(string)}
    filteredEndPoints, err = parent.EndPoints(fetchFilter)
    if err != nil {
        return err
    }

    EndPoint := &vspk.EndPoint{}

    if len(filteredEndPoints) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredEndPoints) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    EndPoint = filteredEndPoints[0]

    d.Set("name", EndPoint.Name)
    d.Set("last_updated_by", EndPoint.LastUpdatedBy)
    d.Set("description", EndPoint.Description)
    d.Set("entity_scope", EndPoint.EntityScope)
    d.Set("external_id", EndPoint.ExternalID)
    
    d.Set("id", EndPoint.Identifier())
    d.Set("parent_id", EndPoint.ParentID)
    d.Set("parent_type", EndPoint.ParentType)
    d.Set("owner", EndPoint.Owner)

    d.SetId(EndPoint.Identifier())
    
    return nil
}