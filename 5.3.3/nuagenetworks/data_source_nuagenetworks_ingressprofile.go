package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceIngressProfile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceIngressProfileRead,
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
            "assoc_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ip_filter_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ip_filter_profile_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ipv6_filter_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ipv6_filter_profile_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_mac_filter_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_mac_filter_profile_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_sap_ingress_qo_s_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_sap_ingress_qo_s_profile_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_redundancy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_gateway"},
            },
            "parent_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_redundancy_group"},
            },
        },
    }
}


func dataSourceIngressProfileRead(d *schema.ResourceData, m interface{}) error {
    filteredIngressProfiles := vspk.IngressProfilesList{}
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
    if attr, ok := d.GetOk("parent_redundancy_group"); ok {
        parent := &vspk.RedundancyGroup{ID: attr.(string)}
        filteredIngressProfiles, err = parent.IngressProfiles(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        filteredIngressProfiles, err = parent.IngressProfiles(fetchFilter)
        if err != nil {
            return err
        }
    }

    IngressProfile := &vspk.IngressProfile{}

    if len(filteredIngressProfiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredIngressProfiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    IngressProfile = filteredIngressProfiles[0]

    d.Set("name", IngressProfile.Name)
    d.Set("last_updated_by", IngressProfile.LastUpdatedBy)
    d.Set("description", IngressProfile.Description)
    d.Set("entity_scope", IngressProfile.EntityScope)
    d.Set("assoc_entity_type", IngressProfile.AssocEntityType)
    d.Set("associated_ip_filter_profile_id", IngressProfile.AssociatedIPFilterProfileID)
    d.Set("associated_ip_filter_profile_name", IngressProfile.AssociatedIPFilterProfileName)
    d.Set("associated_ipv6_filter_profile_id", IngressProfile.AssociatedIPv6FilterProfileID)
    d.Set("associated_ipv6_filter_profile_name", IngressProfile.AssociatedIPv6FilterProfileName)
    d.Set("associated_mac_filter_profile_id", IngressProfile.AssociatedMACFilterProfileID)
    d.Set("associated_mac_filter_profile_name", IngressProfile.AssociatedMACFilterProfileName)
    d.Set("associated_sap_ingress_qo_s_profile_id", IngressProfile.AssociatedSAPIngressQoSProfileID)
    d.Set("associated_sap_ingress_qo_s_profile_name", IngressProfile.AssociatedSAPIngressQoSProfileName)
    d.Set("external_id", IngressProfile.ExternalID)
    
    d.Set("id", IngressProfile.Identifier())
    d.Set("parent_id", IngressProfile.ParentID)
    d.Set("parent_type", IngressProfile.ParentType)
    d.Set("owner", IngressProfile.Owner)

    d.SetId(IngressProfile.Identifier())
    
    return nil
}