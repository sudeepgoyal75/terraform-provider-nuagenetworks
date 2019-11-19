package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
)

func resourceGateway() *schema.Resource {
    return &schema.Resource{
        Create: resourceGatewayCreate,
        Read:   resourceGatewayRead,
        Update: resourceGatewayUpdate,
        Delete: resourceGatewayDelete,
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
            "mac_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "zfb_match_attribute": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "NONE",
            },
            "zfb_match_value": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "bios_release_date": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "bios_version": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "cpu_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "uuid": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "family": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "management_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "datapath_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "patches": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_connected": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "gateway_version": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "redundancy_group_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "peer": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "pending": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "serial_number": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "libraries": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
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
            "bootstrap_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "bootstrap_status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "product_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "use_gateway_vlanvnid": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "associated_gateway_security_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_gateway_security_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_nsg_info_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_netconf_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vtep": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "auto_disc_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "system_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourceGatewayCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Gateway object
    o := &vspk.Gateway{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("zfb_match_attribute"); ok {
        o.ZFBMatchAttribute = attr.(string)
    }
    if attr, ok := d.GetOk("zfb_match_value"); ok {
        o.ZFBMatchValue = attr.(string)
    }
    if attr, ok := d.GetOk("management_id"); ok {
        o.ManagementID = attr.(string)
    }
    if attr, ok := d.GetOk("peer"); ok {
        o.Peer = attr.(string)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("pending"); ok {
        o.Pending = attr.(bool)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("personality"); ok {
        o.Personality = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("use_gateway_vlanvnid"); ok {
        o.UseGatewayVLANVNID = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_netconf_profile_id"); ok {
        o.AssociatedNetconfProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("vtep"); ok {
        o.Vtep = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_me"); ok {
        parent := &vspk.Me{ID: attr.(string)}
        err := parent.CreateGateway(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        err := parent.CreateGateway(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    if attr, ok := d.GetOk("patnatpools"); ok {
        o.AssignPATNATPools(attr.(vspk.PATNATPoolsList))
    }
    return resourceGatewayRead(d, m)
}

func resourceGatewayRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Gateway{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("mac_address", o.MACAddress)
    d.Set("zfb_match_attribute", o.ZFBMatchAttribute)
    d.Set("zfb_match_value", o.ZFBMatchValue)
    d.Set("bios_release_date", o.BIOSReleaseDate)
    d.Set("bios_version", o.BIOSVersion)
    d.Set("cpu_type", o.CPUType)
    d.Set("uuid", o.UUID)
    d.Set("name", o.Name)
    d.Set("family", o.Family)
    d.Set("management_id", o.ManagementID)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("datapath_id", o.DatapathID)
    d.Set("patches", o.Patches)
    d.Set("gateway_connected", o.GatewayConnected)
    d.Set("gateway_version", o.GatewayVersion)
    d.Set("redundancy_group_id", o.RedundancyGroupID)
    d.Set("peer", o.Peer)
    d.Set("template_id", o.TemplateID)
    d.Set("pending", o.Pending)
    d.Set("serial_number", o.SerialNumber)
    d.Set("permitted_action", o.PermittedAction)
    d.Set("personality", o.Personality)
    d.Set("description", o.Description)
    d.Set("libraries", o.Libraries)
    d.Set("enterprise_id", o.EnterpriseID)
    d.Set("entity_scope", o.EntityScope)
    d.Set("location_id", o.LocationID)
    d.Set("bootstrap_id", o.BootstrapID)
    d.Set("bootstrap_status", o.BootstrapStatus)
    d.Set("product_name", o.ProductName)
    d.Set("use_gateway_vlanvnid", o.UseGatewayVLANVNID)
    d.Set("associated_gateway_security_id", o.AssociatedGatewaySecurityID)
    d.Set("associated_gateway_security_profile_id", o.AssociatedGatewaySecurityProfileID)
    d.Set("associated_nsg_info_id", o.AssociatedNSGInfoID)
    d.Set("associated_netconf_profile_id", o.AssociatedNetconfProfileID)
    d.Set("vtep", o.Vtep)
    d.Set("auto_disc_gateway_id", o.AutoDiscGatewayID)
    d.Set("external_id", o.ExternalID)
    d.Set("system_id", o.SystemID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceGatewayUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Gateway{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("zfb_match_attribute"); ok {
        o.ZFBMatchAttribute = attr.(string)
    }
    if attr, ok := d.GetOk("zfb_match_value"); ok {
        o.ZFBMatchValue = attr.(string)
    }
    if attr, ok := d.GetOk("management_id"); ok {
        o.ManagementID = attr.(string)
    }
    if attr, ok := d.GetOk("peer"); ok {
        o.Peer = attr.(string)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("pending"); ok {
        o.Pending = attr.(bool)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("personality"); ok {
        o.Personality = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("use_gateway_vlanvnid"); ok {
        o.UseGatewayVLANVNID = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_netconf_profile_id"); ok {
        o.AssociatedNetconfProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("vtep"); ok {
        o.Vtep = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceGatewayDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Gateway{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}