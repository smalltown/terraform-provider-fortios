// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure custom log fields.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogCustomField() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogCustomFieldCreate,
		Read:   resourceLogCustomFieldRead,
		Update: resourceLogCustomFieldUpdate,
		Delete: resourceLogCustomFieldDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"value": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
		},
	}
}

func resourceLogCustomFieldCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogCustomField(d)
	if err != nil {
		return fmt.Errorf("Error creating LogCustomField resource while getting object: %v", err)
	}

	o, err := c.CreateLogCustomField(obj)

	if err != nil {
		return fmt.Errorf("Error creating LogCustomField resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogCustomField")
	}

	return resourceLogCustomFieldRead(d, m)
}

func resourceLogCustomFieldUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogCustomField(d)
	if err != nil {
		return fmt.Errorf("Error updating LogCustomField resource while getting object: %v", err)
	}

	o, err := c.UpdateLogCustomField(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogCustomField resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogCustomField")
	}

	return resourceLogCustomFieldRead(d, m)
}

func resourceLogCustomFieldDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteLogCustomField(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting LogCustomField resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogCustomFieldRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogCustomField(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogCustomField resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogCustomField(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogCustomField resource from API: %v", err)
	}
	return nil
}

func flattenLogCustomFieldId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogCustomFieldName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogCustomFieldValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogCustomField(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenLogCustomFieldId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenLogCustomFieldName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenLogCustomFieldValue(o["value"], d, "value")); err != nil {
		if !fortiAPIPatch(o["value"]) {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenLogCustomFieldFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogCustomFieldId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFieldName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFieldValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogCustomField(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandLogCustomFieldId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandLogCustomFieldName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok {
		t, err := expandLogCustomFieldValue(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
