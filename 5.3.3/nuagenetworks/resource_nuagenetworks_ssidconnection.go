package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
)

func resourceSSIDConnection() *schema.Resource {
    return &schema.Resource{
        Create: resourceSSIDConnectionCreate,
        Read:   resourceSSIDConnectionRead,
        Update: resourceSSIDConnectionUpdate,
        Delete: resourceSSIDConnectionDelete,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "passphrase": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "readonly": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "redirect_option": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "ORIGINAL_REQUEST",
            },
            "redirect_url": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "generic_config": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "restricted": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "white_list": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "black_list": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "vlan_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "interface_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "broadcast_ssid": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: true,
            },
            "associated_captive_portal_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_egress_qos_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "authentication_mode": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "WPA2",
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_wireless_port": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceSSIDConnectionCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize SSIDConnection object
    o := &vspk.SSIDConnection{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("passphrase"); ok {
        o.Passphrase = attr.(string)
    }
    if attr, ok := d.GetOk("readonly"); ok {
        o.Readonly = attr.(bool)
    }
    if attr, ok := d.GetOk("redirect_option"); ok {
        o.RedirectOption = attr.(string)
    }
    if attr, ok := d.GetOk("redirect_url"); ok {
        o.RedirectURL = attr.(string)
    }
    if attr, ok := d.GetOk("generic_config"); ok {
        o.GenericConfig = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("restricted"); ok {
        o.Restricted = attr.(bool)
    }
    if attr, ok := d.GetOk("white_list"); ok {
        o.WhiteList = attr.([]interface{})
    }
    if attr, ok := d.GetOk("black_list"); ok {
        o.BlackList = attr.([]interface{})
    }
    if attr, ok := d.GetOk("broadcast_ssid"); ok {
        o.BroadcastSSID = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_captive_portal_profile_id"); ok {
        o.AssociatedCaptivePortalProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
        o.AssociatedEgressQOSPolicyID = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("authentication_mode"); ok {
        o.AuthenticationMode = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.WirelessPort{ID: d.Get("parent_wireless_port").(string)}
    err := parent.CreateSSIDConnection(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceSSIDConnectionRead(d, m)
}

func resourceSSIDConnectionRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SSIDConnection{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("passphrase", o.Passphrase)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway_id", o.GatewayID)
    d.Set("readonly", o.Readonly)
    d.Set("redirect_option", o.RedirectOption)
    d.Set("redirect_url", o.RedirectURL)
    d.Set("generic_config", o.GenericConfig)
    d.Set("permitted_action", o.PermittedAction)
    d.Set("description", o.Description)
    d.Set("restricted", o.Restricted)
    d.Set("white_list", o.WhiteList)
    d.Set("black_list", o.BlackList)
    d.Set("vlan_id", o.VlanID)
    d.Set("interface_name", o.InterfaceName)
    d.Set("entity_scope", o.EntityScope)
    d.Set("vport_id", o.VportID)
    d.Set("broadcast_ssid", o.BroadcastSSID)
    d.Set("associated_captive_portal_profile_id", o.AssociatedCaptivePortalProfileID)
    d.Set("associated_egress_qos_policy_id", o.AssociatedEgressQOSPolicyID)
    d.Set("status", o.Status)
    d.Set("authentication_mode", o.AuthenticationMode)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceSSIDConnectionUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SSIDConnection{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("passphrase"); ok {
        o.Passphrase = attr.(string)
    }
    if attr, ok := d.GetOk("readonly"); ok {
        o.Readonly = attr.(bool)
    }
    if attr, ok := d.GetOk("redirect_option"); ok {
        o.RedirectOption = attr.(string)
    }
    if attr, ok := d.GetOk("redirect_url"); ok {
        o.RedirectURL = attr.(string)
    }
    if attr, ok := d.GetOk("generic_config"); ok {
        o.GenericConfig = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("restricted"); ok {
        o.Restricted = attr.(bool)
    }
    if attr, ok := d.GetOk("white_list"); ok {
        o.WhiteList = attr.([]interface{})
    }
    if attr, ok := d.GetOk("black_list"); ok {
        o.BlackList = attr.([]interface{})
    }
    if attr, ok := d.GetOk("broadcast_ssid"); ok {
        o.BroadcastSSID = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_captive_portal_profile_id"); ok {
        o.AssociatedCaptivePortalProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
        o.AssociatedEgressQOSPolicyID = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("authentication_mode"); ok {
        o.AuthenticationMode = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceSSIDConnectionDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.SSIDConnection{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}