package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.3"
)

func resourceIKEEncryptionprofile() *schema.Resource {
    return &schema.Resource{
        Create: resourceIKEEncryptionprofileCreate,
        Read:   resourceIKEEncryptionprofileRead,
        Update: resourceIKEEncryptionprofileUpdate,
        Delete: resourceIKEEncryptionprofileDelete,
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
            "dpd_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "dpd_mode": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "dpd_retry_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "dpd_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "ip_sec_authentication_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ip_sec_dont_fragment": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "ip_sec_enable_pfs": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "ip_sec_encryption_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ip_sec_pre_fragment": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "ip_sec_sa_lifetime": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "ip_sec_sa_replay_window_size": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "isakmp_authentication_mode": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "isakmp_diffie_helman_group_identifier": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "isakmp_encryption_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "isakmp_encryption_key_lifetime": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "isakmp_hash_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "sequence": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceIKEEncryptionprofileCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize IKEEncryptionprofile object
    o := &vspk.IKEEncryptionprofile{
    }
    if attr, ok := d.GetOk("dpd_interval"); ok {
        o.DPDInterval = attr.(int)
    }
    if attr, ok := d.GetOk("dpd_mode"); ok {
        o.DPDMode = attr.(string)
    }
    if attr, ok := d.GetOk("dpd_retry_interval"); ok {
        o.DPDRetryInterval = attr.(int)
    }
    if attr, ok := d.GetOk("dpd_timeout"); ok {
        o.DPDTimeout = attr.(int)
    }
    if attr, ok := d.GetOk("ip_sec_authentication_algorithm"); ok {
        o.IPSecAuthenticationAlgorithm = attr.(string)
    }
    if attr, ok := d.GetOk("ip_sec_dont_fragment"); ok {
        o.IPSecDontFragment = attr.(bool)
    }
    if attr, ok := d.GetOk("ip_sec_enable_pfs"); ok {
        o.IPSecEnablePFS = attr.(bool)
    }
    if attr, ok := d.GetOk("ip_sec_encryption_algorithm"); ok {
        o.IPSecEncryptionAlgorithm = attr.(string)
    }
    if attr, ok := d.GetOk("ip_sec_pre_fragment"); ok {
        o.IPSecPreFragment = attr.(bool)
    }
    if attr, ok := d.GetOk("ip_sec_sa_lifetime"); ok {
        o.IPSecSALifetime = attr.(int)
    }
    if attr, ok := d.GetOk("ip_sec_sa_replay_window_size"); ok {
        o.IPSecSAReplayWindowSize = attr.(string)
    }
    if attr, ok := d.GetOk("isakmp_authentication_mode"); ok {
        o.ISAKMPAuthenticationMode = attr.(string)
    }
    if attr, ok := d.GetOk("isakmp_diffie_helman_group_identifier"); ok {
        o.ISAKMPDiffieHelmanGroupIdentifier = attr.(string)
    }
    if attr, ok := d.GetOk("isakmp_encryption_algorithm"); ok {
        o.ISAKMPEncryptionAlgorithm = attr.(string)
    }
    if attr, ok := d.GetOk("isakmp_encryption_key_lifetime"); ok {
        o.ISAKMPEncryptionKeyLifetime = attr.(int)
    }
    if attr, ok := d.GetOk("isakmp_hash_algorithm"); ok {
        o.ISAKMPHashAlgorithm = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("sequence"); ok {
        o.Sequence = attr.(int)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_enterprise_id"); ok {
        o.AssociatedEnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateIKEEncryptionprofile(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceIKEEncryptionprofileRead(d, m)
}

func resourceIKEEncryptionprofileRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IKEEncryptionprofile{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("dpd_interval", o.DPDInterval)
    d.Set("dpd_mode", o.DPDMode)
    d.Set("dpd_retry_interval", o.DPDRetryInterval)
    d.Set("dpd_timeout", o.DPDTimeout)
    d.Set("ip_sec_authentication_algorithm", o.IPSecAuthenticationAlgorithm)
    d.Set("ip_sec_dont_fragment", o.IPSecDontFragment)
    d.Set("ip_sec_enable_pfs", o.IPSecEnablePFS)
    d.Set("ip_sec_encryption_algorithm", o.IPSecEncryptionAlgorithm)
    d.Set("ip_sec_pre_fragment", o.IPSecPreFragment)
    d.Set("ip_sec_sa_lifetime", o.IPSecSALifetime)
    d.Set("ip_sec_sa_replay_window_size", o.IPSecSAReplayWindowSize)
    d.Set("isakmp_authentication_mode", o.ISAKMPAuthenticationMode)
    d.Set("isakmp_diffie_helman_group_identifier", o.ISAKMPDiffieHelmanGroupIdentifier)
    d.Set("isakmp_encryption_algorithm", o.ISAKMPEncryptionAlgorithm)
    d.Set("isakmp_encryption_key_lifetime", o.ISAKMPEncryptionKeyLifetime)
    d.Set("isakmp_hash_algorithm", o.ISAKMPHashAlgorithm)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("sequence", o.Sequence)
    d.Set("description", o.Description)
    d.Set("entity_scope", o.EntityScope)
    d.Set("associated_enterprise_id", o.AssociatedEnterpriseID)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceIKEEncryptionprofileUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IKEEncryptionprofile{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("dpd_interval"); ok {
        o.DPDInterval = attr.(int)
    }
    if attr, ok := d.GetOk("dpd_mode"); ok {
        o.DPDMode = attr.(string)
    }
    if attr, ok := d.GetOk("dpd_retry_interval"); ok {
        o.DPDRetryInterval = attr.(int)
    }
    if attr, ok := d.GetOk("dpd_timeout"); ok {
        o.DPDTimeout = attr.(int)
    }
    if attr, ok := d.GetOk("ip_sec_authentication_algorithm"); ok {
        o.IPSecAuthenticationAlgorithm = attr.(string)
    }
    if attr, ok := d.GetOk("ip_sec_dont_fragment"); ok {
        o.IPSecDontFragment = attr.(bool)
    }
    if attr, ok := d.GetOk("ip_sec_enable_pfs"); ok {
        o.IPSecEnablePFS = attr.(bool)
    }
    if attr, ok := d.GetOk("ip_sec_encryption_algorithm"); ok {
        o.IPSecEncryptionAlgorithm = attr.(string)
    }
    if attr, ok := d.GetOk("ip_sec_pre_fragment"); ok {
        o.IPSecPreFragment = attr.(bool)
    }
    if attr, ok := d.GetOk("ip_sec_sa_lifetime"); ok {
        o.IPSecSALifetime = attr.(int)
    }
    if attr, ok := d.GetOk("ip_sec_sa_replay_window_size"); ok {
        o.IPSecSAReplayWindowSize = attr.(string)
    }
    if attr, ok := d.GetOk("isakmp_authentication_mode"); ok {
        o.ISAKMPAuthenticationMode = attr.(string)
    }
    if attr, ok := d.GetOk("isakmp_diffie_helman_group_identifier"); ok {
        o.ISAKMPDiffieHelmanGroupIdentifier = attr.(string)
    }
    if attr, ok := d.GetOk("isakmp_encryption_algorithm"); ok {
        o.ISAKMPEncryptionAlgorithm = attr.(string)
    }
    if attr, ok := d.GetOk("isakmp_encryption_key_lifetime"); ok {
        o.ISAKMPEncryptionKeyLifetime = attr.(int)
    }
    if attr, ok := d.GetOk("isakmp_hash_algorithm"); ok {
        o.ISAKMPHashAlgorithm = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("sequence"); ok {
        o.Sequence = attr.(int)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_enterprise_id"); ok {
        o.AssociatedEnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceIKEEncryptionprofileDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IKEEncryptionprofile{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}