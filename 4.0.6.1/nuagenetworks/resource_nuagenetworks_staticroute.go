package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.6.1"
)

func resourceStaticRoute() *schema.Resource {
    return &schema.Resource{
        Create: resourceStaticRouteCreate,
        Read:   resourceStaticRouteRead,
        Update: resourceStaticRouteUpdate,
        Delete: resourceStaticRouteDelete,
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
            "ip_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ipv6_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "next_hop_ip": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_shared_network_resource": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain"},
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_shared_network_resource", "parent_l2_domain"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_shared_network_resource", "parent_domain"},
            },
        },
    }
}

func resourceStaticRouteCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize StaticRoute object
    o := &vspk.StaticRoute{
        Address: d.Get("address").(string),
        Netmask: d.Get("netmask").(string),
        NextHopIp: d.Get("next_hop_ip").(string),
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address"); ok {
        o.IPv6Address = attr.(string)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }
    if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
        parent := &vspk.SharedNetworkResource{ID: attr.(string)}
        err := parent.CreateStaticRoute(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        err := parent.CreateStaticRoute(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        err := parent.CreateStaticRoute(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceStaticRouteRead(d, m)
}

func resourceStaticRouteRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.StaticRoute{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("ip_type", o.IPType)
    d.Set("ipv6_address", o.IPv6Address)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("address", o.Address)
    d.Set("netmask", o.Netmask)
    d.Set("next_hop_ip", o.NextHopIp)
    d.Set("entity_scope", o.EntityScope)
    d.Set("route_distinguisher", o.RouteDistinguisher)
    d.Set("external_id", o.ExternalID)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceStaticRouteUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.StaticRoute{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Address = d.Get("address").(string)
    o.Netmask = d.Get("netmask").(string)
    o.NextHopIp = d.Get("next_hop_ip").(string)
    
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address"); ok {
        o.IPv6Address = attr.(string)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }

    o.Save()

    return nil
}

func resourceStaticRouteDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.StaticRoute{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}