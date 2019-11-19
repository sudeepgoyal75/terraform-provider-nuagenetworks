package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVNFThresholdPolicy() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVNFThresholdPolicyRead,
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
            "cpu_threshold": &schema.Schema{
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
            "action": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "memory_threshold": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "min_occurrence": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "monit_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "assoc_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "storage_threshold": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vnf"},
            },
            "parent_vnf": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
        },
    }
}


func dataSourceVNFThresholdPolicyRead(d *schema.ResourceData, m interface{}) error {
    filteredVNFThresholdPolicies := vspk.VNFThresholdPoliciesList{}
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
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredVNFThresholdPolicies, err = parent.VNFThresholdPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vnf"); ok {
        parent := &vspk.VNF{ID: attr.(string)}
        filteredVNFThresholdPolicies, err = parent.VNFThresholdPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredVNFThresholdPolicies, err = parent.VNFThresholdPolicies(fetchFilter)
        if err != nil {
            return err
        }
    }

    VNFThresholdPolicy := &vspk.VNFThresholdPolicy{}

    if len(filteredVNFThresholdPolicies) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVNFThresholdPolicies) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VNFThresholdPolicy = filteredVNFThresholdPolicies[0]

    d.Set("cpu_threshold", VNFThresholdPolicy.CPUThreshold)
    d.Set("name", VNFThresholdPolicy.Name)
    d.Set("last_updated_by", VNFThresholdPolicy.LastUpdatedBy)
    d.Set("action", VNFThresholdPolicy.Action)
    d.Set("memory_threshold", VNFThresholdPolicy.MemoryThreshold)
    d.Set("description", VNFThresholdPolicy.Description)
    d.Set("min_occurrence", VNFThresholdPolicy.MinOccurrence)
    d.Set("entity_scope", VNFThresholdPolicy.EntityScope)
    d.Set("monit_interval", VNFThresholdPolicy.MonitInterval)
    d.Set("assoc_entity_type", VNFThresholdPolicy.AssocEntityType)
    d.Set("storage_threshold", VNFThresholdPolicy.StorageThreshold)
    d.Set("external_id", VNFThresholdPolicy.ExternalID)
    
    d.Set("id", VNFThresholdPolicy.Identifier())
    d.Set("parent_id", VNFThresholdPolicy.ParentID)
    d.Set("parent_type", VNFThresholdPolicy.ParentType)
    d.Set("owner", VNFThresholdPolicy.Owner)

    d.SetId(VNFThresholdPolicy.Identifier())
    
    return nil
}