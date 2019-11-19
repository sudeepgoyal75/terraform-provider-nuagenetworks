package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.9"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceKeyServerMonitorEncryptedSEK() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceKeyServerMonitorEncryptedSEKRead,
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
            "nsg_certificate_serial_number": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_secured_data_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "key_server_certificate_serial_number": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_key_server_monitor_sek_creation_time": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "associated_key_server_monitor_sekid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_key_server_monitor_sek": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_key_server_monitor"},
            },
            "parent_key_server_monitor": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_key_server_monitor_sek"},
            },
        },
    }
}


func dataSourceKeyServerMonitorEncryptedSEKRead(d *schema.ResourceData, m interface{}) error {
    filteredKeyServerMonitorEncryptedSEKs := vspk.KeyServerMonitorEncryptedSEKsList{}
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
    if attr, ok := d.GetOk("parent_key_server_monitor_sek"); ok {
        parent := &vspk.KeyServerMonitorSEK{ID: attr.(string)}
        filteredKeyServerMonitorEncryptedSEKs, err = parent.KeyServerMonitorEncryptedSEKs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_key_server_monitor"); ok {
        parent := &vspk.KeyServerMonitor{ID: attr.(string)}
        filteredKeyServerMonitorEncryptedSEKs, err = parent.KeyServerMonitorEncryptedSEKs(fetchFilter)
        if err != nil {
            return err
        }
    }

    KeyServerMonitorEncryptedSEK := &vspk.KeyServerMonitorEncryptedSEK{}

    if len(filteredKeyServerMonitorEncryptedSEKs) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredKeyServerMonitorEncryptedSEKs) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    KeyServerMonitorEncryptedSEK = filteredKeyServerMonitorEncryptedSEKs[0]

    d.Set("nsg_certificate_serial_number", KeyServerMonitorEncryptedSEK.NSGCertificateSerialNumber)
    d.Set("last_updated_by", KeyServerMonitorEncryptedSEK.LastUpdatedBy)
    d.Set("gateway_secured_data_id", KeyServerMonitorEncryptedSEK.GatewaySecuredDataID)
    d.Set("key_server_certificate_serial_number", KeyServerMonitorEncryptedSEK.KeyServerCertificateSerialNumber)
    d.Set("entity_scope", KeyServerMonitorEncryptedSEK.EntityScope)
    d.Set("associated_key_server_monitor_sek_creation_time", KeyServerMonitorEncryptedSEK.AssociatedKeyServerMonitorSEKCreationTime)
    d.Set("associated_key_server_monitor_sekid", KeyServerMonitorEncryptedSEK.AssociatedKeyServerMonitorSEKID)
    d.Set("external_id", KeyServerMonitorEncryptedSEK.ExternalID)
    
    d.Set("id", KeyServerMonitorEncryptedSEK.Identifier())
    d.Set("parent_id", KeyServerMonitorEncryptedSEK.ParentID)
    d.Set("parent_type", KeyServerMonitorEncryptedSEK.ParentType)
    d.Set("owner", KeyServerMonitorEncryptedSEK.Owner)

    d.SetId(KeyServerMonitorEncryptedSEK.Identifier())
    
    return nil
}