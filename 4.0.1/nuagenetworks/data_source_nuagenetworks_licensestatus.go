package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceLicensestatus() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceLicensestatusRead,
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
            "total_licensed_nics_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "total_licensed_nsgs_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "total_licensed_used_nics_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "total_licensed_used_nsgs_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "total_licensed_used_vms_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "total_licensed_used_vrsgs_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "total_licensed_used_vrss_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "total_licensed_vms_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "total_licensed_vrsgs_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "total_licensed_vrss_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceLicensestatusRead(d *schema.ResourceData, m interface{}) error {
    filteredLicensestatus := vspk.LicensestatusList{}
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

    Licensestatus := &vspk.Licensestatus{}

    if len(filteredLicensestatus) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredLicensestatus) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Licensestatus = filteredLicensestatus[0]

    d.Set("total_licensed_nics_count", Licensestatus.TotalLicensedNICsCount)
    d.Set("total_licensed_nsgs_count", Licensestatus.TotalLicensedNSGsCount)
    d.Set("total_licensed_used_nics_count", Licensestatus.TotalLicensedUsedNICsCount)
    d.Set("total_licensed_used_nsgs_count", Licensestatus.TotalLicensedUsedNSGsCount)
    d.Set("total_licensed_used_vms_count", Licensestatus.TotalLicensedUsedVMsCount)
    d.Set("total_licensed_used_vrsgs_count", Licensestatus.TotalLicensedUsedVRSGsCount)
    d.Set("total_licensed_used_vrss_count", Licensestatus.TotalLicensedUsedVRSsCount)
    d.Set("total_licensed_vms_count", Licensestatus.TotalLicensedVMsCount)
    d.Set("total_licensed_vrsgs_count", Licensestatus.TotalLicensedVRSGsCount)
    d.Set("total_licensed_vrss_count", Licensestatus.TotalLicensedVRSsCount)
    
    d.Set("id", Licensestatus.Identifier())
    d.Set("parent_id", Licensestatus.ParentID)
    d.Set("parent_type", Licensestatus.ParentType)
    d.Set("owner", Licensestatus.Owner)

    d.SetId(Licensestatus.Identifier())
    
    return nil
}