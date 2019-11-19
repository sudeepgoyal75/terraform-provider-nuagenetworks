package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.9"
)

func resourceKeyServerMonitorEncryptedSEK() *schema.Resource {
    return &schema.Resource{
        Create: resourceKeyServerMonitorEncryptedSEKCreate,
        Read:   resourceKeyServerMonitorEncryptedSEKRead,
        Update: resourceKeyServerMonitorEncryptedSEKUpdate,
        Delete: resourceKeyServerMonitorEncryptedSEKDelete,
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
            "nsg_certificate_serial_number": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_secured_data_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "key_server_certificate_serial_number": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_key_server_monitor_sek_creation_time": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Computed: true,
            },
            "associated_key_server_monitor_sekid": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_key_server_monitor": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceKeyServerMonitorEncryptedSEKCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize KeyServerMonitorEncryptedSEK object
    o := &vspk.KeyServerMonitorEncryptedSEK{
    }
    if attr, ok := d.GetOk("nsg_certificate_serial_number"); ok {
        o.NSGCertificateSerialNumber = attr.(float64)
    }
    if attr, ok := d.GetOk("gateway_secured_data_id"); ok {
        o.GatewaySecuredDataID = attr.(string)
    }
    if attr, ok := d.GetOk("key_server_certificate_serial_number"); ok {
        o.KeyServerCertificateSerialNumber = attr.(float64)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_sek_creation_time"); ok {
        o.AssociatedKeyServerMonitorSEKCreationTime = attr.(float64)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_sekid"); ok {
        o.AssociatedKeyServerMonitorSEKID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.KeyServerMonitor{ID: d.Get("parent_key_server_monitor").(string)}
    err := parent.CreateKeyServerMonitorEncryptedSEK(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceKeyServerMonitorEncryptedSEKRead(d, m)
}

func resourceKeyServerMonitorEncryptedSEKRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.KeyServerMonitorEncryptedSEK{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("nsg_certificate_serial_number", o.NSGCertificateSerialNumber)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway_secured_data_id", o.GatewaySecuredDataID)
    d.Set("key_server_certificate_serial_number", o.KeyServerCertificateSerialNumber)
    d.Set("entity_scope", o.EntityScope)
    d.Set("associated_key_server_monitor_sek_creation_time", o.AssociatedKeyServerMonitorSEKCreationTime)
    d.Set("associated_key_server_monitor_sekid", o.AssociatedKeyServerMonitorSEKID)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceKeyServerMonitorEncryptedSEKUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.KeyServerMonitorEncryptedSEK{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("nsg_certificate_serial_number"); ok {
        o.NSGCertificateSerialNumber = attr.(float64)
    }
    if attr, ok := d.GetOk("gateway_secured_data_id"); ok {
        o.GatewaySecuredDataID = attr.(string)
    }
    if attr, ok := d.GetOk("key_server_certificate_serial_number"); ok {
        o.KeyServerCertificateSerialNumber = attr.(float64)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_sek_creation_time"); ok {
        o.AssociatedKeyServerMonitorSEKCreationTime = attr.(float64)
    }
    if attr, ok := d.GetOk("associated_key_server_monitor_sekid"); ok {
        o.AssociatedKeyServerMonitorSEKID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceKeyServerMonitorEncryptedSEKDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.KeyServerMonitorEncryptedSEK{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}