package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceFlowSecurityPolicy() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceFlowSecurityPolicyRead,
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
            "action": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_address_overwrite": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "flow_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_address_overwrite": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "associated_application_service_id": &schema.Schema{
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
            "parent_flow": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceFlowSecurityPolicyRead(d *schema.ResourceData, m interface{}) error {
    filteredFlowSecurityPolicies := vspk.FlowSecurityPoliciesList{}
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
    parent := &vspk.Flow{ID: d.Get("parent_flow").(string)}
    filteredFlowSecurityPolicies, err = parent.FlowSecurityPolicies(fetchFilter)
    if err != nil {
        return err
    }

    FlowSecurityPolicy := &vspk.FlowSecurityPolicy{}

    if len(filteredFlowSecurityPolicies) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredFlowSecurityPolicies) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    FlowSecurityPolicy = filteredFlowSecurityPolicies[0]

    d.Set("action", FlowSecurityPolicy.Action)
    d.Set("destination_address_overwrite", FlowSecurityPolicy.DestinationAddressOverwrite)
    d.Set("flow_id", FlowSecurityPolicy.FlowID)
    d.Set("entity_scope", FlowSecurityPolicy.EntityScope)
    d.Set("source_address_overwrite", FlowSecurityPolicy.SourceAddressOverwrite)
    d.Set("priority", FlowSecurityPolicy.Priority)
    d.Set("associated_application_service_id", FlowSecurityPolicy.AssociatedApplicationServiceID)
    d.Set("associated_network_object_id", FlowSecurityPolicy.AssociatedNetworkObjectID)
    d.Set("associated_network_object_type", FlowSecurityPolicy.AssociatedNetworkObjectType)
    d.Set("external_id", FlowSecurityPolicy.ExternalID)
    
    d.Set("id", FlowSecurityPolicy.Identifier())
    d.Set("parent_id", FlowSecurityPolicy.ParentID)
    d.Set("parent_type", FlowSecurityPolicy.ParentType)
    d.Set("owner", FlowSecurityPolicy.Owner)

    d.SetId(FlowSecurityPolicy.Identifier())
    
    return nil
}