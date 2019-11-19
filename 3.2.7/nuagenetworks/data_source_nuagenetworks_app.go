package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.7"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceApp() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceAppRead,
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
            "assoc_egress_acl_template_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "assoc_ingress_acl_template_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_domain_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_network_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_network_object_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
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


func dataSourceAppRead(d *schema.ResourceData, m interface{}) error {
    filteredApps := vspk.AppsList{}
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
    filteredApps, err = parent.Apps(fetchFilter)
    if err != nil {
        return err
    }

    App := &vspk.App{}

    if len(filteredApps) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredApps) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    App = filteredApps[0]

    d.Set("name", App.Name)
    d.Set("last_updated_by", App.LastUpdatedBy)
    d.Set("description", App.Description)
    d.Set("entity_scope", App.EntityScope)
    d.Set("assoc_egress_acl_template_id", App.AssocEgressACLTemplateId)
    d.Set("assoc_ingress_acl_template_id", App.AssocIngressACLTemplateId)
    d.Set("associated_domain_id", App.AssociatedDomainID)
    d.Set("associated_domain_type", App.AssociatedDomainType)
    d.Set("associated_network_object_id", App.AssociatedNetworkObjectID)
    d.Set("associated_network_object_type", App.AssociatedNetworkObjectType)
    d.Set("external_id", App.ExternalID)
    
    d.Set("id", App.Identifier())
    d.Set("parent_id", App.ParentID)
    d.Set("parent_type", App.ParentType)
    d.Set("owner", App.Owner)

    d.SetId(App.Identifier())
    
    return nil
}