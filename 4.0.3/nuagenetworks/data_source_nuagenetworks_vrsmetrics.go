package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVrsMetrics() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVrsMetricsRead,
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
            "al_ubr0_status": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "cpu_utilization": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "vrs_process": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "vrsvsc_status": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "receiving_metrics": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "memory_utilization": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "jesxmon_process": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "agent_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "assoc_vcenter_hypervisor_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceVrsMetricsRead(d *schema.ResourceData, m interface{}) error {
    filteredVrsMetrics := vspk.VrsMetricsList{}
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

    VrsMetrics := &vspk.VrsMetrics{}

    if len(filteredVrsMetrics) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVrsMetrics) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VrsMetrics = filteredVrsMetrics[0]

    d.Set("al_ubr0_status", VrsMetrics.ALUbr0Status)
    d.Set("cpu_utilization", VrsMetrics.CPUUtilization)
    d.Set("vrs_process", VrsMetrics.VRSProcess)
    d.Set("vrsvsc_status", VrsMetrics.VRSVSCStatus)
    d.Set("receiving_metrics", VrsMetrics.ReceivingMetrics)
    d.Set("memory_utilization", VrsMetrics.MemoryUtilization)
    d.Set("jesxmon_process", VrsMetrics.JesxmonProcess)
    d.Set("agent_name", VrsMetrics.AgentName)
    d.Set("assoc_vcenter_hypervisor_id", VrsMetrics.AssocVCenterHypervisorID)
    
    d.Set("id", VrsMetrics.Identifier())
    d.Set("parent_id", VrsMetrics.ParentID)
    d.Set("parent_type", VrsMetrics.ParentType)
    d.Set("owner", VrsMetrics.Owner)

    d.SetId(VrsMetrics.Identifier())
    
    return nil
}