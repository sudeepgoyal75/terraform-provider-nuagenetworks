package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVNF() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVNFRead,
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
            "vnf_descriptor_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vnf_descriptor_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "cpu_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "nsg_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nsg_system_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ns_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "task_state": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_known_error": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_user_action": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "memory_mb": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vendor": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "allowed_actions": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "is_attached_to_descriptor": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "associated_vnf_metadata_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_vnf_threshold_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "storage_gb": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_ns_gateway"},
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
        },
    }
}


func dataSourceVNFRead(d *schema.ResourceData, m interface{}) error {
    filteredVNFs := vspk.VNFsList{}
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
        filteredVNFs, err = parent.VNFs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
        parent := &vspk.NSGateway{ID: attr.(string)}
        filteredVNFs, err = parent.VNFs(fetchFilter)
        if err != nil {
            return err
        }
    }

    VNF := &vspk.VNF{}

    if len(filteredVNFs) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVNFs) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VNF = filteredVNFs[0]

    d.Set("vnf_descriptor_id", VNF.VNFDescriptorID)
    d.Set("vnf_descriptor_name", VNF.VNFDescriptorName)
    d.Set("cpu_count", VNF.CPUCount)
    d.Set("nsg_name", VNF.NSGName)
    d.Set("nsg_system_id", VNF.NSGSystemID)
    d.Set("ns_gateway_id", VNF.NSGatewayID)
    d.Set("name", VNF.Name)
    d.Set("task_state", VNF.TaskState)
    d.Set("last_known_error", VNF.LastKnownError)
    d.Set("last_updated_by", VNF.LastUpdatedBy)
    d.Set("last_user_action", VNF.LastUserAction)
    d.Set("memory_mb", VNF.MemoryMB)
    d.Set("vendor", VNF.Vendor)
    d.Set("description", VNF.Description)
    d.Set("allowed_actions", VNF.AllowedActions)
    d.Set("enterprise_id", VNF.EnterpriseID)
    d.Set("entity_scope", VNF.EntityScope)
    d.Set("is_attached_to_descriptor", VNF.IsAttachedToDescriptor)
    d.Set("associated_vnf_metadata_id", VNF.AssociatedVNFMetadataID)
    d.Set("associated_vnf_threshold_policy_id", VNF.AssociatedVNFThresholdPolicyID)
    d.Set("status", VNF.Status)
    d.Set("storage_gb", VNF.StorageGB)
    d.Set("external_id", VNF.ExternalID)
    d.Set("type", VNF.Type)
    
    d.Set("id", VNF.Identifier())
    d.Set("parent_id", VNF.ParentID)
    d.Set("parent_type", VNF.ParentType)
    d.Set("owner", VNF.Owner)

    d.SetId(VNF.Identifier())
    
    return nil
}