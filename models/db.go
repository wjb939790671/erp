package models

import (
	"reflect"

	"github.com/astaxie/beego/orm"
)

type DataBase struct{}

//insert model
//@param model address of model
//@return error of insert mondel and
func (db *DataBase) Add(model interface{}) (int64, error) {
	orm := orm.NewOrm()
	return orm.Insert(model)
}

//updata data
//@param model address of model
//@param cols 要更新的字段
//@return 数据库数据影响行数 和异常
func (db *DataBase) Updata(model interface{}, cols ...string) (int64, error) {
	orm := orm.NewOrm()
	return orm.Update(model, cols...)
}

//假删除
//@param model 是要删除结构体的地址，结构的id字段必须赋值并且delflage字段必须为false
//返回受影响的行数和异常错误
func (db *DataBase) Delete(model interface{}) (int64, error) {
	orm := orm.NewOrm()
	v := reflect.ValueOf(model).Elem()
	v.FieldByName("Delflage").SetBool(true)
	return orm.Update(model, "Delflage")
}

//删除
func (db *DataBase) RealDelete(model interface{}) (int64, error) {
	orm := orm.NewOrm()
	return orm.Delete(model, "Id")
}

//单个查询
//@param model 是要查询结构体的地址
//@param cols 查询条件字段
//返回受影响的行数和异常错误
func (db *DataBase) QueryOne(model interface{}, cols ...string) error {
	orm := orm.NewOrm()
	return orm.Read(model, cols...)
}

//分布查询
//@param model 是要查询结构体的地址
//@param list 结果集的地址
//@param cols 查询字段
//@param where 查询条件的切片
// @param cols 要查询的字段
// @param pageSize 查询页的数据的条数
// @param pageIndex 查询页的页码数
// @param orderBy 按照什么字段排序
//返回 return 结果总数和异常错误
func (db *DataBase) QueryTable(model, list interface{}, where map[string]interface{}, relatedSel []interface{}, orderBy []string, pageSize, pageIndex int, cols ...string) (int64, error) {
	orm := orm.NewOrm()
	qs := orm.QueryTable(model)
	if len(relatedSel) > 0 {
		qs = qs.RelatedSel(relatedSel...)
	}

	//根据条件筛选数据
	for index, value := range where {
		qs = qs.Filter(index, value)
	}
	qs = qs.OrderBy(orderBy...)
	total, qerror := qs.Count()
	if qerror != nil {
		total = 0
	}
	_, qerror = qs.Limit(pageSize, (pageIndex-1)*pageSize).All(list, cols...)
	//分页
	return total, qerror
}

/**
*表查询
*@param model 要查询的表
*@param where 查询条件的切片
*@param cols 要查询的字段
*@return 查询结果的总数，异常
 */
func (db *DataBase) QueryAll(model, list interface{}, where map[string]interface{}, orderBy []string, cols ...string) (int64, error) {
	orm := orm.NewOrm()

	qs := orm.QueryTable(model)
	//根据条件筛选数据
	for index, value := range where {
		qs = qs.Filter(index, value)
	}

	return qs.OrderBy(orderBy...).All(list, cols...)
}

//多对多的插入
func (db *DataBase) M2MAdd(model, dataModel interface{}, relatedSel string) (int64, error) {
	orm := orm.NewOrm()
	m2m := orm.QueryM2M(model, relatedSel)
	m2m.Clear()
	return m2m.Add(dataModel)
}

//加载多对多数据
func (db *DataBase) M2MQuery(model interface{}, relatedSel string) (int64, error) {
	orm := orm.NewOrm()
	if error := orm.Read(model); error != nil {
		return 0, error
	}
	return orm.LoadRelated(model, relatedSel)
}

//存在与否
func (db *DataBase) IsExist(model interface{}, cols ...string) bool {
	orm := orm.NewOrm()
	t := reflect.TypeOf(model).Elem()
	v := reflect.New(t).Elem()
	for _, value := range cols {
		v.FieldByName(value).Set(reflect.ValueOf(model).Elem().FieldByName(value))
	}
	if error := orm.Read(model, cols...); error != nil {
		return false
	}
	return true

}

//sql 语句查询
func (db *DataBase) Raw(sql string, list interface{}, args ...interface{}) (int64, error) {
	orm := orm.NewOrm()
	return orm.Raw(sql, args...).QueryRows(list)
}
