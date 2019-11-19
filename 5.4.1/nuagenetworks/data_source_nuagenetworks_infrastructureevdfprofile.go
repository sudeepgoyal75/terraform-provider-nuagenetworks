package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceInfrastructureEVDFProfile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceInfrastructureEVDFProfileRead,
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
            "ntp_server_key": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ntp_server_key_id": &schema.Schema{
                Type:     schema.TypeInt,
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
            "active_controller": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "service_ipv4_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "proxy_dns_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "use_two_factor": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "standby_controller": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nuage_platform": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceInfrastructureEVDFProfileRead(d *schema.ResourceData, m interface{}) error {
    filteredInfrastructureEVDFProfiles := vspk.InfrastructureEVDFProfilesList{}
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
    parent := m.(*vspk.Me)
    filteredInfrastructureEVDFProfiles, err = parent.InfrastructureEVDFProfiles(fetchFilter)
    if err != nil {
        return err
    }

    InfrastructureEVDFProfile := &vspk.InfrastructureEVDFProfile{}

    if len(filteredInfrastructureEVDFProfiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredInfrastructureEVDFProfiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    InfrastructureEVDFProfile = filteredInfrastructureEVDFProfiles[0]

    d.Set("ntp_server_key", InfrastructureEVDFProfile.NTPServerKey)
    d.Set("ntp_server_key_id", InfrastructureEVDFProfile.NTPServerKeyID)
    d.Set("name", InfrastructureEVDFProfile.Name)
    d.Set("last_updated_by", InfrastructureEVDFProfile.LastUpdatedBy)
    d.Set("active_controller", InfrastructureEVDFProfile.ActiveController)
    d.Set("service_ipv4_subnet", InfrastructureEVDFProfile.ServiceIPv4Subnet)
    d.Set("description", InfrastructureEVDFProfile.Description)
    d.Set("enterprise_id", InfrastructureEVDFProfile.EnterpriseID)
    d.Set("entity_scope", InfrastructureEVDFProfile.EntityScope)
    d.Set("proxy_dns_name", InfrastructureEVDFProfile.ProxyDNSName)
    d.Set("use_two_factor", InfrastructureEVDFProfile.UseTwoFactor)
    d.Set("standby_controller", InfrastructureEVDFProfile.StandbyController)
    d.Set("nuage_platform", InfrastructureEVDFProfile.NuagePlatform)
    d.Set("external_id", InfrastructureEVDFProfile.ExternalID)
    
    d.Set("id", InfrastructureEVDFProfile.Identifier())
    d.Set("parent_id", InfrastructureEVDFProfile.ParentID)
    d.Set("parent_type", InfrastructureEVDFProfile.ParentType)
    d.Set("owner", InfrastructureEVDFProfile.Owner)

    d.SetId(InfrastructureEVDFProfile.Identifier())
    
    return nil
}