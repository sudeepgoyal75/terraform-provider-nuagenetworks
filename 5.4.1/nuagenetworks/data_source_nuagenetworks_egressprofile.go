package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceEgressProfile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceEgressProfileRead,
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
            "associated_sap_egress_qo_s_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_sap_egress_qo_s_profile_name": &schema.Schema{
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


func dataSourceEgressProfileRead(d *schema.ResourceData, m interface{}) error {
    filteredEgressProfiles := vspk.EgressProfilesList{}
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
        filteredEgressProfiles, err = parent.EgressProfiles(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        filteredEgressProfiles, err = parent.EgressProfiles(fetchFilter)
        if err != nil {
            return err
        }
    }

    EgressProfile := &vspk.EgressProfile{}

    if len(filteredEgressProfiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredEgressProfiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    EgressProfile = filteredEgressProfiles[0]

    d.Set("name", EgressProfile.Name)
    d.Set("last_updated_by", EgressProfile.LastUpdatedBy)
    d.Set("description", EgressProfile.Description)
    d.Set("entity_scope", EgressProfile.EntityScope)
    d.Set("assoc_entity_type", EgressProfile.AssocEntityType)
    d.Set("associated_ip_filter_profile_id", EgressProfile.AssociatedIPFilterProfileID)
    d.Set("associated_ip_filter_profile_name", EgressProfile.AssociatedIPFilterProfileName)
    d.Set("associated_ipv6_filter_profile_id", EgressProfile.AssociatedIPv6FilterProfileID)
    d.Set("associated_ipv6_filter_profile_name", EgressProfile.AssociatedIPv6FilterProfileName)
    d.Set("associated_mac_filter_profile_id", EgressProfile.AssociatedMACFilterProfileID)
    d.Set("associated_mac_filter_profile_name", EgressProfile.AssociatedMACFilterProfileName)
    d.Set("associated_sap_egress_qo_s_profile_id", EgressProfile.AssociatedSAPEgressQoSProfileID)
    d.Set("associated_sap_egress_qo_s_profile_name", EgressProfile.AssociatedSAPEgressQoSProfileName)
    d.Set("external_id", EgressProfile.ExternalID)
    
    d.Set("id", EgressProfile.Identifier())
    d.Set("parent_id", EgressProfile.ParentID)
    d.Set("parent_type", EgressProfile.ParentType)
    d.Set("owner", EgressProfile.Owner)

    d.SetId(EgressProfile.Identifier())
    
    return nil
}