package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNSGatewaySummary() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNSGatewaySummaryRead,
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
            "major_alarms_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "latitude": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "time_zone_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "minor_alarms_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "info_alarms_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "locality": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "longitude": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "bootstrap_status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "country": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "critical_alarms_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "nsg_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "state": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "system_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_enterprise", "parent_ns_gateway"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_enterprise", "parent_ns_gateway"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_ns_gateway"},
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_enterprise"},
            },
        },
    }
}


func dataSourceNSGatewaySummaryRead(d *schema.ResourceData, m interface{}) error {
    filteredNSGatewaySummaries := vspk.NSGatewaySummariesList{}
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
        filteredNSGatewaySummaries, err = parent.NSGatewaySummaries(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredNSGatewaySummaries, err = parent.NSGatewaySummaries(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredNSGatewaySummaries, err = parent.NSGatewaySummaries(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
        parent := &vspk.NSGateway{ID: attr.(string)}
        filteredNSGatewaySummaries, err = parent.NSGatewaySummaries(fetchFilter)
        if err != nil {
            return err
        }
    }

    NSGatewaySummary := &vspk.NSGatewaySummary{}

    if len(filteredNSGatewaySummaries) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNSGatewaySummaries) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NSGatewaySummary = filteredNSGatewaySummaries[0]

    d.Set("major_alarms_count", NSGatewaySummary.MajorAlarmsCount)
    d.Set("gateway_id", NSGatewaySummary.GatewayID)
    d.Set("gateway_name", NSGatewaySummary.GatewayName)
    d.Set("latitude", NSGatewaySummary.Latitude)
    d.Set("address", NSGatewaySummary.Address)
    d.Set("time_zone_id", NSGatewaySummary.TimeZoneID)
    d.Set("minor_alarms_count", NSGatewaySummary.MinorAlarmsCount)
    d.Set("info_alarms_count", NSGatewaySummary.InfoAlarmsCount)
    d.Set("enterprise_id", NSGatewaySummary.EnterpriseID)
    d.Set("locality", NSGatewaySummary.Locality)
    d.Set("longitude", NSGatewaySummary.Longitude)
    d.Set("bootstrap_status", NSGatewaySummary.BootstrapStatus)
    d.Set("country", NSGatewaySummary.Country)
    d.Set("critical_alarms_count", NSGatewaySummary.CriticalAlarmsCount)
    d.Set("nsg_version", NSGatewaySummary.NsgVersion)
    d.Set("state", NSGatewaySummary.State)
    d.Set("system_id", NSGatewaySummary.SystemID)
    
    d.Set("id", NSGatewaySummary.Identifier())
    d.Set("parent_id", NSGatewaySummary.ParentID)
    d.Set("parent_type", NSGatewaySummary.ParentType)
    d.Set("owner", NSGatewaySummary.Owner)

    d.SetId(NSGatewaySummary.Identifier())
    
    return nil
}