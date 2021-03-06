// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure virtual Access Point (VAP) groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerVapGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerVapGroupCreate,
		Read:   resourceWirelessControllerVapGroupRead,
		Update: resourceWirelessControllerVapGroupUpdate,
		Delete: resourceWirelessControllerVapGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"vaps": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessControllerVapGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerVapGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerVapGroup resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerVapGroup(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerVapGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerVapGroup")
	}

	return resourceWirelessControllerVapGroupRead(d, m)
}

func resourceWirelessControllerVapGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerVapGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVapGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerVapGroup(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVapGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerVapGroup")
	}

	return resourceWirelessControllerVapGroupRead(d, m)
}

func resourceWirelessControllerVapGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerVapGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerVapGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerVapGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerVapGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVapGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerVapGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVapGroup resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerVapGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapGroupComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapGroupVaps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenWirelessControllerVapGroupVapsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerVapGroupVapsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerVapGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerVapGroupName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerVapGroupComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vaps", flattenWirelessControllerVapGroupVaps(o["vaps"], d, "vaps")); err != nil {
			if !fortiAPIPatch(o["vaps"]) {
				return fmt.Errorf("Error reading vaps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vaps"); ok {
			if err = d.Set("vaps", flattenWirelessControllerVapGroupVaps(o["vaps"], d, "vaps")); err != nil {
				if !fortiAPIPatch(o["vaps"]) {
					return fmt.Errorf("Error reading vaps: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerVapGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerVapGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGroupComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGroupVaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandWirelessControllerVapGroupVapsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapGroupVapsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerVapGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerVapGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerVapGroupComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("vaps"); ok {
		t, err := expandWirelessControllerVapGroupVaps(d, v, "vaps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vaps"] = t
		}
	}

	return &obj, nil
}
