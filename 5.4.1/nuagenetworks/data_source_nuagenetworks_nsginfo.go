package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNSGInfo() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNSGInfoRead,
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
            "mac_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "aar_application_release_date": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "aar_application_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "bios_release_date": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "bios_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "sku": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "tpm_status": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "tpm_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "cpu_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nsg_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "uuid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "family": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "patches_detail": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "serial_number": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "libraries": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "cmd_detailed_status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "cmd_detailed_status_code": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "cmd_download_progress": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "cmd_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "cmd_last_updated_date": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "cmd_status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "cmd_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enterprise_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "bootstrap_status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "product_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ns_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "system_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}


func dataSourceNSGInfoRead(d *schema.ResourceData, m interface{}) error {
    filteredNSGInfos := vspk.NSGInfosList{}
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
    if attr, ok := d.GetOk("parent_ns_gateway"); ok {
        parent := &vspk.NSGateway{ID: attr.(string)}
        filteredNSGInfos, err = parent.NSGInfos(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredNSGInfos, err = parent.NSGInfos(fetchFilter)
        if err != nil {
            return err
        }
    }

    NSGInfo := &vspk.NSGInfo{}

    if len(filteredNSGInfos) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNSGInfos) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NSGInfo = filteredNSGInfos[0]

    d.Set("mac_address", NSGInfo.MACAddress)
    d.Set("aar_application_release_date", NSGInfo.AARApplicationReleaseDate)
    d.Set("aar_application_version", NSGInfo.AARApplicationVersion)
    d.Set("bios_release_date", NSGInfo.BIOSReleaseDate)
    d.Set("bios_version", NSGInfo.BIOSVersion)
    d.Set("sku", NSGInfo.SKU)
    d.Set("tpm_status", NSGInfo.TPMStatus)
    d.Set("tpm_version", NSGInfo.TPMVersion)
    d.Set("cpu_type", NSGInfo.CPUType)
    d.Set("nsg_version", NSGInfo.NSGVersion)
    d.Set("uuid", NSGInfo.UUID)
    d.Set("name", NSGInfo.Name)
    d.Set("family", NSGInfo.Family)
    d.Set("patches_detail", NSGInfo.PatchesDetail)
    d.Set("serial_number", NSGInfo.SerialNumber)
    d.Set("personality", NSGInfo.Personality)
    d.Set("libraries", NSGInfo.Libraries)
    d.Set("cmd_detailed_status", NSGInfo.CmdDetailedStatus)
    d.Set("cmd_detailed_status_code", NSGInfo.CmdDetailedStatusCode)
    d.Set("cmd_download_progress", NSGInfo.CmdDownloadProgress)
    d.Set("cmd_id", NSGInfo.CmdID)
    d.Set("cmd_last_updated_date", NSGInfo.CmdLastUpdatedDate)
    d.Set("cmd_status", NSGInfo.CmdStatus)
    d.Set("cmd_type", NSGInfo.CmdType)
    d.Set("enterprise_id", NSGInfo.EnterpriseID)
    d.Set("enterprise_name", NSGInfo.EnterpriseName)
    d.Set("entity_scope", NSGInfo.EntityScope)
    d.Set("bootstrap_status", NSGInfo.BootstrapStatus)
    d.Set("product_name", NSGInfo.ProductName)
    d.Set("associated_entity_type", NSGInfo.AssociatedEntityType)
    d.Set("associated_ns_gateway_id", NSGInfo.AssociatedNSGatewayID)
    d.Set("external_id", NSGInfo.ExternalID)
    d.Set("system_id", NSGInfo.SystemID)
    
    d.Set("id", NSGInfo.Identifier())
    d.Set("parent_id", NSGInfo.ParentID)
    d.Set("parent_type", NSGInfo.ParentType)
    d.Set("owner", NSGInfo.Owner)

    d.SetId(NSGInfo.Identifier())
    
    return nil
}