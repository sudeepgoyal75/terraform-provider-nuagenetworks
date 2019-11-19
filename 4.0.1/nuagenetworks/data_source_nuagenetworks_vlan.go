package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVLAN() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVLANRead,
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
            "value": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "readonly": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "restricted": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "use_user_mnemonic": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "user_mnemonic": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_bgp_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_egress_qos_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_ns_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vsg_redundant_port", "parent_redundant_port", "parent_port"},
            },
            "parent_vsg_redundant_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_ns_port", "parent_redundant_port", "parent_port"},
            },
            "parent_redundant_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_ns_port", "parent_vsg_redundant_port", "parent_port"},
            },
            "parent_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_ns_port", "parent_vsg_redundant_port", "parent_redundant_port"},
            },
        },
    }
}


func dataSourceVLANRead(d *schema.ResourceData, m interface{}) error {
    filteredVLANs := vspk.VLANsList{}
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
    if attr, ok := d.GetOk("parent_ns_port"); ok {
        parent := &vspk.NSPort{ID: attr.(string)}
        filteredVLANs, err = parent.VLANs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vsg_redundant_port"); ok {
        parent := &vspk.VsgRedundantPort{ID: attr.(string)}
        filteredVLANs, err = parent.VLANs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_redundant_port"); ok {
        parent := &vspk.RedundantPort{ID: attr.(string)}
        filteredVLANs, err = parent.VLANs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_port"); ok {
        parent := &vspk.Port{ID: attr.(string)}
        filteredVLANs, err = parent.VLANs(fetchFilter)
        if err != nil {
            return err
        }
    }

    VLAN := &vspk.VLAN{}

    if len(filteredVLANs) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVLANs) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VLAN = filteredVLANs[0]

    d.Set("value", VLAN.Value)
    d.Set("last_updated_by", VLAN.LastUpdatedBy)
    d.Set("gateway_id", VLAN.GatewayID)
    d.Set("readonly", VLAN.Readonly)
    d.Set("template_id", VLAN.TemplateID)
    d.Set("permitted_action", VLAN.PermittedAction)
    d.Set("description", VLAN.Description)
    d.Set("restricted", VLAN.Restricted)
    d.Set("entity_scope", VLAN.EntityScope)
    d.Set("vport_id", VLAN.VportID)
    d.Set("use_user_mnemonic", VLAN.UseUserMnemonic)
    d.Set("user_mnemonic", VLAN.UserMnemonic)
    d.Set("associated_bgp_profile_id", VLAN.AssociatedBGPProfileID)
    d.Set("associated_egress_qos_policy_id", VLAN.AssociatedEgressQOSPolicyID)
    d.Set("status", VLAN.Status)
    d.Set("external_id", VLAN.ExternalID)
    
    d.Set("id", VLAN.Identifier())
    d.Set("parent_id", VLAN.ParentID)
    d.Set("parent_type", VLAN.ParentType)
    d.Set("owner", VLAN.Owner)

    d.SetId(VLAN.Identifier())
    
    return nil
}