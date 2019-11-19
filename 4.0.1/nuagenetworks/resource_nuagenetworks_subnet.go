package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.1"
)

func resourceSubnet() *schema.Resource {
    return &schema.Resource{
        Create: resourceSubnetCreate,
        Read:   resourceSubnetRead,
        Update: resourceSubnetUpdate,
        Delete: resourceSubnetDelete,
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
            "pat_enabled": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ip_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "maintenance_mode": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_mac_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "service_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vn_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "encryption": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "underlay": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "underlay_enabled": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "policy_group_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "route_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "split_subnet": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "proxy_arp": &schema.Schema{
                Type:     schema.TypeBool,
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
            "associated_multicast_channel_map_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_shared_network_resource_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "public": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "multicast": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_zone": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceSubnetCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Subnet object
    o := &vspk.Subnet{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("pat_enabled"); ok {
        o.PATEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("maintenance_mode"); ok {
        o.MaintenanceMode = attr.(string)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_mac_address"); ok {
        o.GatewayMACAddress = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("service_id"); ok {
        o.ServiceID = attr.(int)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("vn_id"); ok {
        o.VnId = attr.(int)
    }
    if attr, ok := d.GetOk("encryption"); ok {
        o.Encryption = attr.(string)
    }
    if attr, ok := d.GetOk("underlay"); ok {
        o.Underlay = attr.(bool)
    }
    if attr, ok := d.GetOk("underlay_enabled"); ok {
        o.UnderlayEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("policy_group_id"); ok {
        o.PolicyGroupID = attr.(int)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("route_target"); ok {
        o.RouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("split_subnet"); ok {
        o.SplitSubnet = attr.(bool)
    }
    if attr, ok := d.GetOk("proxy_arp"); ok {
        o.ProxyARP = attr.(bool)
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
    if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
        o.AssociatedMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_shared_network_resource_id"); ok {
        o.AssociatedSharedNetworkResourceID = attr.(string)
    }
    if attr, ok := d.GetOk("public"); ok {
        o.Public = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast"); ok {
        o.Multicast = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Zone{ID: d.Get("parent_zone").(string)}
    err := parent.CreateSubnet(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    if attr, ok := d.GetOk("ikegatewayconnections"); ok {
        o.AssignIKEGatewayConnections(attr.(vspk.IKEGatewayConnectionsList))
    }
    return resourceSubnetRead(d, m)
}

func resourceSubnetRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Subnet{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("pat_enabled", o.PATEnabled)
    d.Set("ip_type", o.IPType)
    d.Set("maintenance_mode", o.MaintenanceMode)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway", o.Gateway)
    d.Set("gateway_mac_address", o.GatewayMACAddress)
    d.Set("address", o.Address)
    d.Set("template_id", o.TemplateID)
    d.Set("service_id", o.ServiceID)
    d.Set("description", o.Description)
    d.Set("netmask", o.Netmask)
    d.Set("vn_id", o.VnId)
    d.Set("encryption", o.Encryption)
    d.Set("underlay", o.Underlay)
    d.Set("underlay_enabled", o.UnderlayEnabled)
    d.Set("entity_scope", o.EntityScope)
    d.Set("policy_group_id", o.PolicyGroupID)
    d.Set("route_distinguisher", o.RouteDistinguisher)
    d.Set("route_target", o.RouteTarget)
    d.Set("split_subnet", o.SplitSubnet)
    d.Set("proxy_arp", o.ProxyARP)
    d.Set("associated_application_id", o.AssociatedApplicationID)
    d.Set("associated_application_object_id", o.AssociatedApplicationObjectID)
    d.Set("associated_application_object_type", o.AssociatedApplicationObjectType)
    d.Set("associated_multicast_channel_map_id", o.AssociatedMulticastChannelMapID)
    d.Set("associated_shared_network_resource_id", o.AssociatedSharedNetworkResourceID)
    d.Set("public", o.Public)
    d.Set("multicast", o.Multicast)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceSubnetUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Subnet{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("pat_enabled"); ok {
        o.PATEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("maintenance_mode"); ok {
        o.MaintenanceMode = attr.(string)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_mac_address"); ok {
        o.GatewayMACAddress = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("service_id"); ok {
        o.ServiceID = attr.(int)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("vn_id"); ok {
        o.VnId = attr.(int)
    }
    if attr, ok := d.GetOk("encryption"); ok {
        o.Encryption = attr.(string)
    }
    if attr, ok := d.GetOk("underlay"); ok {
        o.Underlay = attr.(bool)
    }
    if attr, ok := d.GetOk("underlay_enabled"); ok {
        o.UnderlayEnabled = attr.(string)
    }
    if attr, ok := d.GetOk("policy_group_id"); ok {
        o.PolicyGroupID = attr.(int)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("route_target"); ok {
        o.RouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("split_subnet"); ok {
        o.SplitSubnet = attr.(bool)
    }
    if attr, ok := d.GetOk("proxy_arp"); ok {
        o.ProxyARP = attr.(bool)
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
    if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
        o.AssociatedMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_shared_network_resource_id"); ok {
        o.AssociatedSharedNetworkResourceID = attr.(string)
    }
    if attr, ok := d.GetOk("public"); ok {
        o.Public = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast"); ok {
        o.Multicast = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceSubnetDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Subnet{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}