package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceApplicationService() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceApplicationServiceRead,
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
            "dscp": &schema.Schema{
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
            "destination_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "direction": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "protocol": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ether_type": &schema.Schema{
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


func dataSourceApplicationServiceRead(d *schema.ResourceData, m interface{}) error {
    filteredApplicationServices := vspk.ApplicationServicesList{}
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
        filteredApplicationServices, err = parent.ApplicationServices(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredApplicationServices, err = parent.ApplicationServices(fetchFilter)
        if err != nil {
            return err
        }
    }

    ApplicationService := &vspk.ApplicationService{}

    if len(filteredApplicationServices) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredApplicationServices) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    ApplicationService = filteredApplicationServices[0]

    d.Set("dscp", ApplicationService.DSCP)
    d.Set("name", ApplicationService.Name)
    d.Set("last_updated_by", ApplicationService.LastUpdatedBy)
    d.Set("description", ApplicationService.Description)
    d.Set("destination_port", ApplicationService.DestinationPort)
    d.Set("direction", ApplicationService.Direction)
    d.Set("entity_scope", ApplicationService.EntityScope)
    d.Set("source_port", ApplicationService.SourcePort)
    d.Set("protocol", ApplicationService.Protocol)
    d.Set("ether_type", ApplicationService.EtherType)
    d.Set("external_id", ApplicationService.ExternalID)
    
    d.Set("id", ApplicationService.Identifier())
    d.Set("parent_id", ApplicationService.ParentID)
    d.Set("parent_type", ApplicationService.ParentType)
    d.Set("owner", ApplicationService.Owner)

    d.SetId(ApplicationService.Identifier())
    
    return nil
}