package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.9"
)

func resourceDomainFIPAclTemplateEntry() *schema.Resource {
    return &schema.Resource{
        Create: resourceDomainFIPAclTemplateEntryCreate,
        Read:   resourceDomainFIPAclTemplateEntryRead,
        Update: resourceDomainFIPAclTemplateEntryUpdate,
        Delete: resourceDomainFIPAclTemplateEntryDelete,
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
                Required: true,
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
                Optional: true,
                Computed: true,
            },
            "action_details": &schema.Schema{
                Type:     schema.TypeMap,
                Optional: true,
                Computed: true,
            },
            "action_details_raw": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
            "address_override": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "reflexive": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "dest_pg_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "dest_pg_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "destination_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "destination_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "destination_value": &schema.Schema{
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
                Computed: true,
            },
            "mirror_destination_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "flow_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "enterprise_name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
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
                Optional: true,
                Computed: true,
            },
            "policy_state": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "source_pg_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "source_pg_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "source_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "source_value": &schema.Schema{
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
            "associated_application_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_application_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_application_object_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_live_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "stateful": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "stats_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "stats_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "ether_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_domain_fip_acl_template": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceDomainFIPAclTemplateEntryCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize DomainFIPAclTemplateEntry object
    o := &vspk.DomainFIPAclTemplateEntry{
        ACLTemplateName: d.Get("acl_template_name").(string),
        EnterpriseName: d.Get("enterprise_name").(string),
        DomainName: d.Get("domain_name").(string),
    }
    if attr, ok := d.GetOk("icmp_code"); ok {
        o.ICMPCode = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_type"); ok {
        o.ICMPType = attr.(string)
    }
    if attr, ok := d.GetOk("dscp"); ok {
        o.DSCP = attr.(string)
    }
    if attr, ok := d.GetOk("action"); ok {
        o.Action = attr.(string)
    }
    if attr, ok := d.GetOk("action_details"); ok {
        o.ActionDetails = attr.(interface{})
    }
    if attr, ok := d.GetOk("address_override"); ok {
        o.AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("reflexive"); ok {
        o.Reflexive = attr.(bool)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("dest_pg_id"); ok {
        o.DestPgId = attr.(string)
    }
    if attr, ok := d.GetOk("dest_pg_type"); ok {
        o.DestPgType = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("destination_type"); ok {
        o.DestinationType = attr.(string)
    }
    if attr, ok := d.GetOk("destination_value"); ok {
        o.DestinationValue = attr.(string)
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
    if attr, ok := d.GetOk("location_type"); ok {
        o.LocationType = attr.(string)
    }
    if attr, ok := d.GetOk("policy_state"); ok {
        o.PolicyState = attr.(string)
    }
    if attr, ok := d.GetOk("source_pg_id"); ok {
        o.SourcePgId = attr.(string)
    }
    if attr, ok := d.GetOk("source_pg_type"); ok {
        o.SourcePgType = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("source_type"); ok {
        o.SourceType = attr.(string)
    }
    if attr, ok := d.GetOk("source_value"); ok {
        o.SourceValue = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("protocol"); ok {
        o.Protocol = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_id"); ok {
        o.AssociatedApplicationID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_object_id"); ok {
        o.AssociatedApplicationObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_object_type"); ok {
        o.AssociatedApplicationObjectType = attr.(string)
    }
    if attr, ok := d.GetOk("associated_live_entity_id"); ok {
        o.AssociatedLiveEntityID = attr.(string)
    }
    if attr, ok := d.GetOk("stateful"); ok {
        o.Stateful = attr.(bool)
    }
    if attr, ok := d.GetOk("stats_id"); ok {
        o.StatsID = attr.(string)
    }
    if attr, ok := d.GetOk("stats_logging_enabled"); ok {
        o.StatsLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("ether_type"); ok {
        o.EtherType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.DomainFIPAclTemplate{ID: d.Get("parent_domain_fip_acl_template").(string)}
    err := parent.CreateDomainFIPAclTemplateEntry(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceDomainFIPAclTemplateEntryRead(d, m)
}

func resourceDomainFIPAclTemplateEntryRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DomainFIPAclTemplateEntry{
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
    d.Set("dscp", o.DSCP)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("action", o.Action)
    if v, ok := o.ActionDetails.(string); ok {
		raw := make(map[string]string)
		raw["raw"] = v
		d.Set("action_details_raw", raw)
	} else {
		d.Set("action_details", o.ActionDetails)
	}
    d.Set("address_override", o.AddressOverride)
    d.Set("reflexive", o.Reflexive)
    d.Set("description", o.Description)
    d.Set("dest_pg_id", o.DestPgId)
    d.Set("dest_pg_type", o.DestPgType)
    d.Set("destination_port", o.DestinationPort)
    d.Set("destination_type", o.DestinationType)
    d.Set("destination_value", o.DestinationValue)
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
    d.Set("source_pg_id", o.SourcePgId)
    d.Set("source_pg_type", o.SourcePgType)
    d.Set("source_port", o.SourcePort)
    d.Set("source_type", o.SourceType)
    d.Set("source_value", o.SourceValue)
    d.Set("priority", o.Priority)
    d.Set("protocol", o.Protocol)
    d.Set("associated_application_id", o.AssociatedApplicationID)
    d.Set("associated_application_object_id", o.AssociatedApplicationObjectID)
    d.Set("associated_application_object_type", o.AssociatedApplicationObjectType)
    d.Set("associated_live_entity_id", o.AssociatedLiveEntityID)
    d.Set("stateful", o.Stateful)
    d.Set("stats_id", o.StatsID)
    d.Set("stats_logging_enabled", o.StatsLoggingEnabled)
    d.Set("ether_type", o.EtherType)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceDomainFIPAclTemplateEntryUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DomainFIPAclTemplateEntry{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.ACLTemplateName = d.Get("acl_template_name").(string)
    o.EnterpriseName = d.Get("enterprise_name").(string)
    o.DomainName = d.Get("domain_name").(string)
    
    if attr, ok := d.GetOk("icmp_code"); ok {
        o.ICMPCode = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_type"); ok {
        o.ICMPType = attr.(string)
    }
    if attr, ok := d.GetOk("dscp"); ok {
        o.DSCP = attr.(string)
    }
    if attr, ok := d.GetOk("action"); ok {
        o.Action = attr.(string)
    }
    if attr, ok := d.GetOk("action_details"); ok {
        o.ActionDetails = attr.(interface{})
    }
    if attr, ok := d.GetOk("address_override"); ok {
        o.AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("reflexive"); ok {
        o.Reflexive = attr.(bool)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("dest_pg_id"); ok {
        o.DestPgId = attr.(string)
    }
    if attr, ok := d.GetOk("dest_pg_type"); ok {
        o.DestPgType = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("destination_type"); ok {
        o.DestinationType = attr.(string)
    }
    if attr, ok := d.GetOk("destination_value"); ok {
        o.DestinationValue = attr.(string)
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
    if attr, ok := d.GetOk("location_type"); ok {
        o.LocationType = attr.(string)
    }
    if attr, ok := d.GetOk("policy_state"); ok {
        o.PolicyState = attr.(string)
    }
    if attr, ok := d.GetOk("source_pg_id"); ok {
        o.SourcePgId = attr.(string)
    }
    if attr, ok := d.GetOk("source_pg_type"); ok {
        o.SourcePgType = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("source_type"); ok {
        o.SourceType = attr.(string)
    }
    if attr, ok := d.GetOk("source_value"); ok {
        o.SourceValue = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("protocol"); ok {
        o.Protocol = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_id"); ok {
        o.AssociatedApplicationID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_object_id"); ok {
        o.AssociatedApplicationObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_object_type"); ok {
        o.AssociatedApplicationObjectType = attr.(string)
    }
    if attr, ok := d.GetOk("associated_live_entity_id"); ok {
        o.AssociatedLiveEntityID = attr.(string)
    }
    if attr, ok := d.GetOk("stateful"); ok {
        o.Stateful = attr.(bool)
    }
    if attr, ok := d.GetOk("stats_id"); ok {
        o.StatsID = attr.(string)
    }
    if attr, ok := d.GetOk("stats_logging_enabled"); ok {
        o.StatsLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("ether_type"); ok {
        o.EtherType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceDomainFIPAclTemplateEntryDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DomainFIPAclTemplateEntry{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}