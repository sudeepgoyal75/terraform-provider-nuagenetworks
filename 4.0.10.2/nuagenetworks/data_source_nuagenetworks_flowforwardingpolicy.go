package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.10.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceFlowForwardingPolicy() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceFlowForwardingPolicyRead,
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
            "redirect_target_id": &schema.Schema{
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
            "type": &schema.Schema{
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


func dataSourceFlowForwardingPolicyRead(d *schema.ResourceData, m interface{}) error {
    filteredFlowForwardingPolicies := vspk.FlowForwardingPoliciesList{}
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
    filteredFlowForwardingPolicies, err = parent.FlowForwardingPolicies(fetchFilter)
    if err != nil {
        return err
    }

    FlowForwardingPolicy := &vspk.FlowForwardingPolicy{}

    if len(filteredFlowForwardingPolicies) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredFlowForwardingPolicies) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    FlowForwardingPolicy = filteredFlowForwardingPolicies[0]

    d.Set("redirect_target_id", FlowForwardingPolicy.RedirectTargetID)
    d.Set("destination_address_overwrite", FlowForwardingPolicy.DestinationAddressOverwrite)
    d.Set("flow_id", FlowForwardingPolicy.FlowID)
    d.Set("entity_scope", FlowForwardingPolicy.EntityScope)
    d.Set("source_address_overwrite", FlowForwardingPolicy.SourceAddressOverwrite)
    d.Set("associated_application_service_id", FlowForwardingPolicy.AssociatedApplicationServiceID)
    d.Set("associated_network_object_id", FlowForwardingPolicy.AssociatedNetworkObjectID)
    d.Set("associated_network_object_type", FlowForwardingPolicy.AssociatedNetworkObjectType)
    d.Set("external_id", FlowForwardingPolicy.ExternalID)
    d.Set("type", FlowForwardingPolicy.Type)
    
    d.Set("id", FlowForwardingPolicy.Identifier())
    d.Set("parent_id", FlowForwardingPolicy.ParentID)
    d.Set("parent_type", FlowForwardingPolicy.ParentType)
    d.Set("owner", FlowForwardingPolicy.Owner)

    d.SetId(FlowForwardingPolicy.Identifier())
    
    return nil
}