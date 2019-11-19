package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVPort() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVPortRead,
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
            "vlanid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "has_attached_interfaces": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "active": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "address_spoofing": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zone_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "operational_state": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_floating_ip_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_multicast_channel_map_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_send_multicast_channel_map_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multi_nic_vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast": &schema.Schema{
                Type:     schema.TypeString,
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
            "system_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_zone", "parent_vrs", "parent_tier", "parent_multi_nic_vport", "parent_floating_ip"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_zone", "parent_vrs", "parent_tier", "parent_multi_nic_vport", "parent_floating_ip"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_redirection_target", "parent_policy_group", "parent_zone", "parent_vrs", "parent_tier", "parent_multi_nic_vport", "parent_floating_ip"},
            },
            "parent_redirection_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_policy_group", "parent_zone", "parent_vrs", "parent_tier", "parent_multi_nic_vport", "parent_floating_ip"},
            },
            "parent_policy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_zone", "parent_vrs", "parent_tier", "parent_multi_nic_vport", "parent_floating_ip"},
            },
            "parent_zone": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_vrs", "parent_tier", "parent_multi_nic_vport", "parent_floating_ip"},
            },
            "parent_vrs": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_zone", "parent_tier", "parent_multi_nic_vport", "parent_floating_ip"},
            },
            "parent_tier": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_zone", "parent_vrs", "parent_multi_nic_vport", "parent_floating_ip"},
            },
            "parent_multi_nic_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_zone", "parent_vrs", "parent_tier", "parent_floating_ip"},
            },
            "parent_floating_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_zone", "parent_vrs", "parent_tier", "parent_multi_nic_vport"},
            },
        },
    }
}


func dataSourceVPortRead(d *schema.ResourceData, m interface{}) error {
    filteredVPorts := vspk.VPortsList{}
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
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_redirection_target"); ok {
        parent := &vspk.RedirectionTarget{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_policy_group"); ok {
        parent := &vspk.PolicyGroup{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_zone"); ok {
        parent := &vspk.Zone{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vrs"); ok {
        parent := &vspk.VRS{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_tier"); ok {
        parent := &vspk.Tier{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_multi_nic_vport"); ok {
        parent := &vspk.MultiNICVPort{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_floating_ip"); ok {
        parent := &vspk.FloatingIp{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return err
        }
    }

    VPort := &vspk.VPort{}

    if len(filteredVPorts) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVPorts) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VPort = filteredVPorts[0]

    d.Set("vlanid", VPort.VLANID)
    d.Set("name", VPort.Name)
    d.Set("has_attached_interfaces", VPort.HasAttachedInterfaces)
    d.Set("last_updated_by", VPort.LastUpdatedBy)
    d.Set("active", VPort.Active)
    d.Set("address_spoofing", VPort.AddressSpoofing)
    d.Set("description", VPort.Description)
    d.Set("entity_scope", VPort.EntityScope)
    d.Set("domain_id", VPort.DomainID)
    d.Set("zone_id", VPort.ZoneID)
    d.Set("operational_state", VPort.OperationalState)
    d.Set("associated_floating_ip_id", VPort.AssociatedFloatingIPID)
    d.Set("associated_multicast_channel_map_id", VPort.AssociatedMulticastChannelMapID)
    d.Set("associated_send_multicast_channel_map_id", VPort.AssociatedSendMulticastChannelMapID)
    d.Set("multi_nic_vport_id", VPort.MultiNICVPortID)
    d.Set("multicast", VPort.Multicast)
    d.Set("external_id", VPort.ExternalID)
    d.Set("type", VPort.Type)
    d.Set("system_type", VPort.SystemType)
    
    d.Set("id", VPort.Identifier())
    d.Set("parent_id", VPort.ParentID)
    d.Set("parent_type", VPort.ParentType)
    d.Set("owner", VPort.Owner)

    d.SetId(VPort.Identifier())
    
    return nil
}