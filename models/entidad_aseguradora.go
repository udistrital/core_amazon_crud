package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type EntidadAseguradora struct {
	Id        int     `orm:"column(id);pk;auto"`
	Codigo    int     `orm:"column(codigo)"`
	Nombre    string  `orm:"column(nombre)"`
	Nit       string  `orm:"column(nit)"`
	Direccion string  `orm:"column(direccion)"`
	Ciudad    int     `orm:"column(ciudad)"`
	Telefono  float64 `orm:"column(telefono)"`
	Estado    bool    `orm:"column(estado)"`
}

func (t *EntidadAseguradora) TableName() string {
	return "entidad_aseguradora"
}

func init() {
	orm.RegisterModel(new(EntidadAseguradora))
}

// AddEntidadAseguradora insert a new EntidadAseguradora into database and returns
// last inserted Id on success.
func AddEntidadAseguradora(m *EntidadAseguradora) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEntidadAseguradoraById retrieves EntidadAseguradora by Id. Returns error if
// Id doesn't exist
func GetEntidadAseguradoraById(id int) (v *EntidadAseguradora, err error) {
	o := orm.NewOrm()
	v = &EntidadAseguradora{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEntidad_aseguradora retrieves all EntidadAseguradora matches certain condition. Returns empty list if
// no records exist
func GetAllEntidad_aseguradora(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EntidadAseguradora))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []EntidadAseguradora
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateEntidad_aseguradora updates EntidadAseguradora by Id and returns error if
// the record to be updated doesn't exist
func UpdateEntidad_aseguradoraById(m *EntidadAseguradora) (err error) {
	o := orm.NewOrm()
	v := EntidadAseguradora{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEntidad_aseguradora deletes EntidadAseguradora by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEntidad_aseguradora(id int64) (err error) {
	o := orm.NewOrm()
	v := EntidadAseguradora{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EntidadAseguradora{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
