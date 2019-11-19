package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceVirtualFirewallRule() *schema.Resource {
    return &schema.Resource{
        Create: resourceVirtualFirewallRuleCreate,
        Read:   resourceVirtualFirewallRuleRead,
        Update: resourceVirtualFirewallRuleUpdate,
        Delete: resourceVirtualFirewallRuleDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
        Schema: map[string]*schema.Schema{
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "acl_template_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "icmp_code": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "icmp_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ipv6_address_override": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "dscp": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "action": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "address_override": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "web_filter_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "web_filter_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "destination_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "network_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "network_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "ANY",
            },
            "mirror_destination_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "flow_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "enterprise_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "location_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "location_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "policy_state": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "protocol": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_egress_entry_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_ingress_entry_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_l7_application_signature_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_live_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_live_template_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_traffic_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_traffic_type_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "stateful": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "stats_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "stats_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "ether_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "overlay_mirror_destination_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_virtual_firewall_policy": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceVirtualFirewallRuleCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VirtualFirewallRule object
    o := &vspk.VirtualFirewallRule{
        Action: d.Get("action").(string),
        LocationType: d.Get("location_type").(string),
    }
    if attr, ok := d.GetOk("icmp_code"); ok {
        o.ICMPCode = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_type"); ok {
        o.ICMPType = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address_override"); ok {
        o.IPv6AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("dscp"); ok {
        o.DSCP = attr.(string)
    }
    if attr, ok := d.GetOk("address_override"); ok {
        o.AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("web_filter_id"); ok {
        o.WebFilterID = attr.(string)
    }
    if attr, ok := d.GetOk("web_filter_type"); ok {
        o.WebFilterType = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("network_id"); ok {
        o.NetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("network_type"); ok {
        o.NetworkType = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_destination_id"); ok {
        o.MirrorDestinationID = attr.(string)
    }
    if attr, ok := d.GetOk("flow_logging_enabled"); ok {
        o.FlowLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("location_id"); ok {
        o.LocationID = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("protocol"); ok {
        o.Protocol = attr.(string)
    }
    if attr, ok := d.GetOk("associated_l7_application_signature_id"); ok {
        o.AssociatedL7ApplicationSignatureID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_traffic_type"); ok {
        o.AssociatedTrafficType = attr.(string)
    }
    if attr, ok := d.GetOk("associated_traffic_type_id"); ok {
        o.AssociatedTrafficTypeID = attr.(string)
    }
    if attr, ok := d.GetOk("stateful"); ok {
        o.Stateful = attr.(bool)
    }
    if attr, ok := d.GetOk("stats_logging_enabled"); ok {
        o.StatsLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("ether_type"); ok {
        o.EtherType = attr.(string)
    }
    if attr, ok := d.GetOk("overlay_mirror_destination_id"); ok {
        o.OverlayMirrorDestinationID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.VirtualFirewallPolicy{ID: d.Get("parent_virtual_firewall_policy").(string)}
    err := parent.CreateVirtualFirewallRule(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceVirtualFirewallRuleRead(d, m)
}

func resourceVirtualFirewallRuleRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VirtualFirewallRule{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("acl_template_name", o.ACLTemplateName)
    d.Set("icmp_code", o.ICMPCode)
    d.Set("icmp_type", o.ICMPType)
    d.Set("ipv6_address_override", o.IPv6AddressOverride)
    d.Set("dscp", o.DSCP)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("action", o.Action)
    d.Set("address_override", o.AddressOverride)
    d.Set("web_filter_id", o.WebFilterID)
    d.Set("web_filter_type", o.WebFilterType)
    d.Set("description", o.Description)
    d.Set("destination_port", o.DestinationPort)
    d.Set("network_id", o.NetworkID)
    d.Set("network_type", o.NetworkType)
    d.Set("mirror_destination_id", o.MirrorDestinationID)
    d.Set("flow_logging_enabled", o.FlowLoggingEnabled)
    d.Set("enterprise_name", o.EnterpriseName)
    d.Set("entity_scope", o.EntityScope)
    d.Set("location_id", o.LocationID)
    d.Set("location_type", o.LocationType)
    d.Set("policy_state", o.PolicyState)
    d.Set("domain_name", o.DomainName)
    d.Set("source_port", o.SourcePort)
    d.Set("priority", o.Priority)
    d.Set("protocol", o.Protocol)
    d.Set("associated_egress_entry_id", o.AssociatedEgressEntryID)
    d.Set("associated_ingress_entry_id", o.AssociatedIngressEntryID)
    d.Set("associated_l7_application_signature_id", o.AssociatedL7ApplicationSignatureID)
    d.Set("associated_live_entity_id", o.AssociatedLiveEntityID)
    d.Set("associated_live_template_id", o.AssociatedLiveTemplateID)
    d.Set("associated_traffic_type", o.AssociatedTrafficType)
    d.Set("associated_traffic_type_id", o.AssociatedTrafficTypeID)
    d.Set("stateful", o.Stateful)
    d.Set("stats_id", o.StatsID)
    d.Set("stats_logging_enabled", o.StatsLoggingEnabled)
    d.Set("ether_type", o.EtherType)
    d.Set("overlay_mirror_destination_id", o.OverlayMirrorDestinationID)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVirtualFirewallRuleUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VirtualFirewallRule{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Action = d.Get("action").(string)
    o.LocationType = d.Get("location_type").(string)
    
    if attr, ok := d.GetOk("icmp_code"); ok {
        o.ICMPCode = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_type"); ok {
        o.ICMPType = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address_override"); ok {
        o.IPv6AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("dscp"); ok {
        o.DSCP = attr.(string)
    }
    if attr, ok := d.GetOk("address_override"); ok {
        o.AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("web_filter_id"); ok {
        o.WebFilterID = attr.(string)
    }
    if attr, ok := d.GetOk("web_filter_type"); ok {
        o.WebFilterType = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("network_id"); ok {
        o.NetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("network_type"); ok {
        o.NetworkType = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_destination_id"); ok {
        o.MirrorDestinationID = attr.(string)
    }
    if attr, ok := d.GetOk("flow_logging_enabled"); ok {
        o.FlowLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("location_id"); ok {
        o.LocationID = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("protocol"); ok {
        o.Protocol = attr.(string)
    }
    if attr, ok := d.GetOk("associated_l7_application_signature_id"); ok {
        o.AssociatedL7ApplicationSignatureID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_traffic_type"); ok {
        o.AssociatedTrafficType = attr.(string)
    }
    if attr, ok := d.GetOk("associated_traffic_type_id"); ok {
        o.AssociatedTrafficTypeID = attr.(string)
    }
    if attr, ok := d.GetOk("stateful"); ok {
        o.Stateful = attr.(bool)
    }
    if attr, ok := d.GetOk("stats_logging_enabled"); ok {
        o.StatsLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("ether_type"); ok {
        o.EtherType = attr.(string)
    }
    if attr, ok := d.GetOk("overlay_mirror_destination_id"); ok {
        o.OverlayMirrorDestinationID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVirtualFirewallRuleDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VirtualFirewallRule{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}