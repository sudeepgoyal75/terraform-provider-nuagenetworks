package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.10.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceExternalAppService() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceExternalAppServiceRead,
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
            "destination_nat_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_nat_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "destination_nat_mask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "metadata": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "egress_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "virtual_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "virtual_ip_required": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "ingress_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_nat_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_nat_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "associated_service_egress_group_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_service_egress_redirect_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_service_ingress_group_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_service_ingress_redirect_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain"},
            },
        },
    }
}


func dataSourceExternalAppServiceRead(d *schema.ResourceData, m interface{}) error {
    filteredExternalAppServices := vspk.ExternalAppServicesList{}
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
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredExternalAppServices, err = parent.ExternalAppServices(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredExternalAppServices, err = parent.ExternalAppServices(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredExternalAppServices, err = parent.ExternalAppServices(fetchFilter)
        if err != nil {
            return err
        }
    }

    ExternalAppService := &vspk.ExternalAppService{}

    if len(filteredExternalAppServices) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredExternalAppServices) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    ExternalAppService = filteredExternalAppServices[0]

    d.Set("name", ExternalAppService.Name)
    d.Set("last_updated_by", ExternalAppService.LastUpdatedBy)
    d.Set("description", ExternalAppService.Description)
    d.Set("destination_nat_address", ExternalAppService.DestinationNATAddress)
    d.Set("destination_nat_enabled", ExternalAppService.DestinationNATEnabled)
    d.Set("destination_nat_mask", ExternalAppService.DestinationNATMask)
    d.Set("metadata", ExternalAppService.Metadata)
    d.Set("egress_type", ExternalAppService.EgressType)
    d.Set("virtual_ip", ExternalAppService.VirtualIP)
    d.Set("virtual_ip_required", ExternalAppService.VirtualIPRequired)
    d.Set("ingress_type", ExternalAppService.IngressType)
    d.Set("entity_scope", ExternalAppService.EntityScope)
    d.Set("source_nat_address", ExternalAppService.SourceNATAddress)
    d.Set("source_nat_enabled", ExternalAppService.SourceNATEnabled)
    d.Set("associated_service_egress_group_id", ExternalAppService.AssociatedServiceEgressGroupID)
    d.Set("associated_service_egress_redirect_id", ExternalAppService.AssociatedServiceEgressRedirectID)
    d.Set("associated_service_ingress_group_id", ExternalAppService.AssociatedServiceIngressGroupID)
    d.Set("associated_service_ingress_redirect_id", ExternalAppService.AssociatedServiceIngressRedirectID)
    d.Set("external_id", ExternalAppService.ExternalID)
    
    d.Set("id", ExternalAppService.Identifier())
    d.Set("parent_id", ExternalAppService.ParentID)
    d.Set("parent_type", ExternalAppService.ParentType)
    d.Set("owner", ExternalAppService.Owner)

    d.SetId(ExternalAppService.Identifier())
    
    return nil
}