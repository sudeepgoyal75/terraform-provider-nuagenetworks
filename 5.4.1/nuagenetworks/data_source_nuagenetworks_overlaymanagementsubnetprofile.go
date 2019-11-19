package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceOverlayManagementSubnetProfile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceOverlayManagementSubnetProfileRead,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_dna_subnet_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "syslog_destination_ids": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "parent_overlay_management_profile": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceOverlayManagementSubnetProfileRead(d *schema.ResourceData, m interface{}) error {
    filteredOverlayManagementSubnetProfiles := vspk.OverlayManagementSubnetProfilesList{}
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
    parent := &vspk.OverlayManagementProfile{ID: d.Get("parent_overlay_management_profile").(string)}
    filteredOverlayManagementSubnetProfiles, err = parent.OverlayManagementSubnetProfiles(fetchFilter)
    if err != nil {
        return err
    }

    OverlayManagementSubnetProfile := &vspk.OverlayManagementSubnetProfile{}

    if len(filteredOverlayManagementSubnetProfiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredOverlayManagementSubnetProfiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    OverlayManagementSubnetProfile = filteredOverlayManagementSubnetProfiles[0]

    d.Set("name", OverlayManagementSubnetProfile.Name)
    d.Set("description", OverlayManagementSubnetProfile.Description)
    d.Set("associated_dna_subnet_id", OverlayManagementSubnetProfile.AssociatedDNASubnetID)
    d.Set("syslog_destination_ids", OverlayManagementSubnetProfile.SyslogDestinationIDs)
    
    d.Set("id", OverlayManagementSubnetProfile.Identifier())
    d.Set("parent_id", OverlayManagementSubnetProfile.ParentID)
    d.Set("parent_type", OverlayManagementSubnetProfile.ParentType)
    d.Set("owner", OverlayManagementSubnetProfile.Owner)

    d.SetId(OverlayManagementSubnetProfile.Identifier())
    
    return nil
}