package api

// -------------------------------------------------------------------------------- //
// From struct_data.go
//
// func (dm *DataMap) CopyDataEntries(dep DataEntries, endpoint string, pointId string, name string) *DataEntries {
// 	var ret *DataEntries
// 	for range only.Once {
// 		var des DataEntries
// 		des = dep.Copy()
// 		for i := range des.Entries {
// 			des.Entries[i].SetEndpoint(endpoint, pointId)
// 			des.Entries[i].SetPointName(name)
// 			dm.Add(des.Entries[i])
// 		}
//
// 		if len(des.Entries) == 0 {
// 			fmt.Printf("OOOPS\n")
// 		}
// 		epn := des.Entries[0].EndPoint
// 		ret = dm.Map[epn]
// 	}
// 	return ret
// }
//
// func (dm *DataMap) TableSort() []string {
// 	var sorted []string
//
// 	for range only.Once {
// 		for p := range dm.DataTables {
// 			sorted = append(sorted, p)
// 		}
// 		sort.Strings(sorted)
// 	}
// 	return sorted
// }
//
// func (dm *DataMap) AddAny(endpoint string, parentDeviceId string, pid valueTypes.PointId, name string, groupName string, date valueTypes.DateTime, value interface{}, unit string, Type string, timeSpan string) {
//
// 	for range only.Once {
// 		var point Point
// 		p := GetPoint(parentDeviceId + "." + pid.String())
// 		if p == nil {
// 			// No point found. Create one.
// 			point = CreatePoint(parentDeviceId, pid, name, groupName, unit, Type, timeSpan)
// 		} else {
// 			point = *p
// 		}
//
// 		uvs, isNil, ok := valueTypes.AnyToUnitValue(value, unit, Type, valueTypes.DateTimeLayout)
// 		if !ok {
// 			fmt.Printf("ERROR: AddAny(endpoint '%s', parentId '%s', pid '%s', name '%s', date '%s', value '%v')",
// 				endpoint, parentDeviceId, pid, name, date, value)
// 			break
// 		}
// 		if isNil {
// 			point.ValueType += "(NIL)"
// 		}
//
// 		for _, uv := range uvs {
// 			if uv.GetUnit() != point.Unit {
// 				fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, uv.ValueFloat(), uv.GetUnit())
// 				point.Unit = uv.GetUnit()
// 			}
//
// 			var parent ParentDevice
// 			parent.Set(parentDeviceId)
// 			point.Parents.Add(parent)
//
// 			de := CreatePointDataEntry(endpoint, parentDeviceId, point, date, uv)
// 			de.Point = &point
// 			dm.Add(de)
// 		}
// 	}
// }
//
// func (dm *DataMap) AddUnitValue(endpoint string, parentDeviceId string, pid valueTypes.PointId, name string, groupName string, date valueTypes.DateTime, uv valueTypes.UnitValue, timeSpan string) {
//
// 	for range only.Once {
// 		var point Point
// 		p := GetPoint(parentDeviceId + "." + pid.String())
// 		if p == nil {
// 			// No point found. Create one.
// 			point = CreatePoint(parentDeviceId, pid, name, groupName, uv.GetUnit(), uv.Type(), timeSpan)
// 		} else {
// 			point = *p
// 		}
//
// 		if uv.GetUnit() != point.Unit {
// 			fmt.Printf("OOOPS: Unit mismatch - %s %s != %f %s\n", uv.String(), point.Unit, uv.ValueFloat(), uv.GetUnit())
// 			point.Unit = uv.GetUnit()
// 		}
//
// 		var parent ParentDevice
// 		parent.Set(parentDeviceId)
// 		point.Parents.Add(parent)
//
// 		de := CreatePointDataEntry(endpoint, parentDeviceId, point, date, uv)
// 		de.Point = &point
// 		dm.Add(de)
// 	}
// }

// -------------------------------------------------------------------------------- //
// From struct_de.go
//
// func (de *DataEntry) FullId() string {
// 	return de.EndPoint	// + "." + de.Point.Id.String()
// }
//
// func (de *DataEntry) SetEndpoint(endpoint string, pointId string) {
// 	de.EndPoint = endpoint + "." + pointId
// 	de.Point.Id.SetString(pointId)
// }
//
// func (de *DataEntry) SetPointId(pointId string) {
// 	de.Point.Id.SetString(pointId)
// }
//
// func (de *DataEntry) SetPointName(name string) {
// 	if name != "" {
// 		de.Point.SetName(name)
// 	}
// }
//
// func (de *DataEntry) MakeState(state bool) DataEntry {
// 	var ret DataEntry
// 	for range only.Once {
// 		// uv := valueTypes.SetUnitValueBool(state)
// 		// de.Value = uv.String()
// 		// de.ValueFloat = uv.Value()
// 		de.Value = valueTypes.SetUnitValueBool(state)
// 		de.Point.Unit = ""
// 		de.Point.ValueType = "Bool"
// 		de.Point.Valid = true
// 		de.Valid = true
// 		// de.EndPoint += ".state"
// 		de.Hide = false
// 	}
//
// 	return ret
// }
//
// func (de *DataEntry) MakeFloat(value float64, unit string, Type string) {
// 	for range only.Once {
// 		if unit == "" {
// 			unit = de.Point.Unit
// 		}
// 		if Type == "" {
// 			Type = de.Point.ValueType
// 		}
// 		// uv := valueTypes.SetUnitValueFloat(value, unit, Type)
// 		// de.Value = uv.String()
// 		// de.ValueFloat = uv.Value()
// 		de.Value = valueTypes.SetUnitValueFloat(value, unit, Type)
// 		de.Valid = true
// 		de.Hide = false
// 	}
// }
//
// func (de *DataEntry) Copy() DataEntry {
// 	var ret DataEntry
// 	ret = *de
// 	var point Point
// 	point = *de.Point
// 	ret.Point = &point
// 	return ret
// }

// -------------------------------------------------------------------------------- //
// From struct_dpe.go
//
// func NewDataPointEntries() DataEntries {
// 	return DataEntries{
// 		Entries: []DataEntry{},
// 		// Map: &GoStruct.StructMap{},
// 	}
// }
//
// func (de *DataEntries) Hide() {
// 	for range only.Once {
// 		for i := range de.Entries {
// 			de.Entries[i].Hide = true
// 		}
// 	}
// }
//
// func (de *DataEntries) GetEntryValue(index int) valueTypes.UnitValue {
// 	var ret valueTypes.UnitValue
// 	for range only.Once {
// 		ref := de.GetEntry(index)
// 		if ref == nil {
// 			break
// 		}
// 		ret = ref.Value
// 	}
// 	return ret
// }
//
// func (de *DataEntries) GetFloat() float64 {
// 	var ret float64
// 	for range only.Once {
// 		ref := de.GetEntry(0)
// 		if ref == nil {
// 			break
// 		}
// 		ret = ref.Value.Value()
// 	}
// 	return ret
// }
//
// func (de *DataEntries) MatchPointId(pointId string) bool {
// 	var yes bool
// 	for range only.Once {
// 		for _, v := range de.Entries {
// 			if v.Point.Id.String() == pointId {
// 				yes = true
// 				break
// 			}
// 		}
// 	}
// 	return yes
// }
//
// func (de *DataEntries) GetUnits() string {
// 	var unit string
// 	for range only.Once {
// 		for _, v := range de.Entries {
// 			unit = v.Point.Unit
// 			break
// 		}
// 	}
// 	return unit
// }
//
// func (de *DataEntries) SetUnits(units string) *DataEntries {
// 	for range only.Once {
// 		for i := range de.Entries {
// 			de.Entries[i].Point.Unit = units
// 		}
// 	}
// 	return de
// }
//
// func (de *DataEntries) SetGroupName(groupName string) *DataEntries {
// 	for range only.Once {
// 		for i := range de.Entries {
// 			de.Entries[i].Point.GroupName = groupName
// 		}
// 	}
// 	return de
// }
//
// func (de *DataEntries) SetTimestamp(timeStamp valueTypes.DateTime) *DataEntries {
// 	for range only.Once {
// 		for i := range de.Entries {
// 			// dt := valueTypes.SetDateTimeString(timeStamp)
// 			de.Entries[i].Date = timeStamp
// 		}
// 	}
// 	return de
// }
//
// func (de *DataEntries) Copy() DataEntries {
// 	var ret DataEntries
// 	for _, d := range de.Entries {
// 		// var point Point
// 		// point = *d.Point
// 		// d.Point = &point
// 		ret.Entries = append(ret.Entries, d.Copy())
// 	}
// 	return ret
// }
//
// func (de *DataEntries) MakeState(state bool) *DataEntries {
// 	for i := range de.Entries {
// 		de.Entries[i].MakeState(state)
// 	}
// 	return de
// }
//
// func (de *DataEntries) SetFloat(value float64, unit string, Type string) *DataEntries {
// 	for i := range de.Entries {
// 		de.Entries[i].MakeFloat(value, unit, Type)
// 	}
// 	return de
// }
//
// func (de *DataEntries) FloatToState(value float64) *DataEntries {
// 	for i := range de.Entries {
// 		if value == 0 {
// 			de.Entries[i].MakeState(false)
// 			break
// 		}
// 		de.Entries[i].MakeState(true)
// 	}
// 	return de
// }
//
// func (dm *DataMap) CreateDataTables() Tables {
// 	tables := make(Tables, 0)
//
// 	for range only.Once {
// 		for name := range dm.StructMap.TableMap {
// 			// values = make(GoStruct.StructValuesMap)
//
// 			td := dm.StructMap.GetTableData(name)
// 			if !td.IsValid {
// 				continue
// 			}
//
// 			values := td.GetValues()
// 			if (values == nil) || (len(values) == 0) {
// 				fmt.Printf("No data table results for '%s'\n", name)
// 				break
// 			}
//
// 			headers := td.GetHeaders()
// 			table := output.NewTable(headers...)
// 			for row := range values {
// 				var items []interface{}
// 				for _, col := range td.Columns {
// 					items = append(items, values.GetCell(row, col))
// 				}
// 				dm.Error = table.AddRow(items...)
// 				if dm.Error != nil {
// 					break
// 				}
// 			}
// 			if dm.Error != nil {
// 				break
// 			}
//
// 			title := td.Current.DataStructure.DataTableTitle
// 			if title == "" {
// 				title = td.Current.DataStructure.DataTableName
// 			}
// 			if title == "" {
// 				title = valueTypes.PointToName(td.Current.DataStructure.DataTableId)
// 			}
// 			// if title == "" {
// 			// 	title = valueTypes.PointToName(td.Current.DataStructure.PointId)
// 			// }
// 			// dm.EndPoint.GetRequestArgNames()
//
// 			table.SetName(name)
// 			if title == "" {
// 				table.SetTitle("DataTable %s.%s", dm.EndPoint.GetArea(), td.Name)
// 				table.SetFilePrefix("%s.%s", dm.EndPoint.GetArea(), td.Name)
// 			} else {
// 				table.SetTitle("DataTable %s.%s (%s)", dm.EndPoint.GetArea(), td.Name, title)
// 				table.SetFilePrefix("%s.%s-%s", dm.EndPoint.GetArea(), td.Name, td.Current.DataStructure.DataTableId)
// 			}
//
// 			// table.Sort(td.SortOn)
// 			table.SetJson(nil)
// 			table.SetRaw(nil)
//
// 			table.SetGraphFilter("")	// @TODO - Consider setting graph options here instead of iSolarCloud/data.go:487
//
// 			// if sgd.Options.GraphRequest.TimeColumn == nil {
// 			// 	for _, col := range table.GetHeaders() {
// 			// 		val := value.GetCell(0, col)
// 			// 		if val.Type() == "DateTime" {
// 			// 			sgd.Options.GraphRequest.TimeColumn = &col
// 			// 			break
// 			// 		}
// 			// 	}
// 			// }
// 			//
// 			// if sgd.Options.GraphRequest.DataColumn == nil {
// 			// 	for _, col := range table.GetHeaders() {
// 			// 		val := value.GetCell(0, col)
// 			// 		if val.IsNumber() {
// 			// 			sgd.Options.GraphRequest.DataColumn = &col
// 			// 			break
// 			// 		}
// 			// 	}
// 			// }
// 			//
// 			// if sgd.Options.GraphRequest.ValueColumn == nil {
// 			// 	for _, col := range table.GetHeaders() {
// 			// 		val := value.GetCell(0, col)
// 			// 		if val.IsNumber() {
// 			// 			sgd.Options.GraphRequest.ValueColumn = &col
// 			// 			break
// 			// 		}
// 			// 	}
// 			// }
//
// 			tables[name] = Table {
// 				Values: values,
// 				Table:  table,
// 			}
// 			// values[name] = vals
// 			// tables[name] = table
// 		}
// 	}
//
// 	return tables
// }

// -------------------------------------------------------------------------------- //
// From struct_data.go
//
// func (dm *DataMap) GetEntry(entry string, index int) *DataEntry {
// 	var ret *DataEntry
// 	for range only.Once {
// 		pe := dm.Map[entry]
// 		if pe.Entries != nil {
// 			ret = pe.GetEntry(index)
// 			break
// 		}
//
// 		for k, v := range dm.Map {
// 			if strings.HasSuffix(k, "." + entry) {
// 				ret = v.GetEntry(index)
// 				break
// 			}
// 		}
// 	}
// 	return ret
// }
//
// func (dm *DataMap) GetFloatValue(entry string, index int) float64 {
// 	var ret float64
// 	for range only.Once {
// 		pe := dm.GetEntry(entry, index)
// 		if pe.IsNotValid() {
// 			fmt.Printf("ERROR: GetFloatValue('%s', '%d')\n", entry, index)
// 			break
// 		}
// 		ret = pe.Value.ValueFloat()
// 	}
// 	return ret
// }
//
// func (dm *DataMap) GetValue(entry string, index int) float64 {
// 	var ret float64
// 	for range only.Once {
// 		v := dm.GetEntry(entry, index)
// 		if v.IsNotValid() {
// 			fmt.Printf("ERROR: GetValue('%s', %d)\n", entry, index)
// 			break
// 		}
//
// 		ret = v.Value.ValueFloat()
// 	}
// 	return ret
// }
//
// func (dm *DataMap) GetEntryFromPointId(pointId string) *DataEntries {
// 	var ret *DataEntries
// 	for range only.Once {
// 		for i, v := range dm.Map {
// 			if v.MatchPointId(pointId) {
// 				ret = dm.Map[i]
// 				break
// 			}
// 		}
// 	}
// 	return ret
// }
//
// func (dm *DataMap) SetEntryUnits(pointId string, unit string) {
// 	for range only.Once {
// 		for i, v := range dm.Map {
// 			if v.MatchPointId(pointId) {
// 				// e := dm.Map[i]
// 				// dm.Map[i] = e.SetUnits(unit)
// 				dm.Map[i].SetUnits(unit)
// 				break
// 			}
// 		}
// 	}
// }
//
// func (dm *DataMap) SetEntryGroupName(pointId string, groupName string) {
// 	for range only.Once {
// 		for i, v := range dm.Map {
// 			if v.MatchPointId(pointId) {
// 				// e := dm.Map[i]
// 				// dm.Map[i] = e.SetGroupName(groupName)
// 				dm.Map[i].SetGroupName(groupName)
// 				break
// 			}
// 		}
// 	}
// }
//
// func (dm *DataMap) SetEntryTimestamp(pointId string, timeStamp valueTypes.DateTime) {
// 	for range only.Once {
// 		for i, v := range dm.Map {
// 			if v.MatchPointId(pointId) {
// 				// e := dm.Map[i]
// 				// dm.Map[i] = e.SetTimestamp(timeStamp)
// 				dm.Map[i].SetTimestamp(timeStamp)
// 				break
// 			}
// 		}
// 	}
// }
//
// func (dm *DataMap) FromRefAddAlias(ref string, parentId string, pid string, name string) {
// 	for range only.Once {
// 		pe := dm.GetEntry(ref, 0)
// 		if pe.IsNotValid() {
// 			fmt.Printf("ERROR: FromRefAddAlias('%s', '%s', '%s', '%s')\n", ref, parentId, pid, name)
// 			break
// 		}
//
// 		de := CopyDataEntry(*pe, pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name, pe.Point.GroupName, pe.Point.Unit, pe.Point.ValueType)
// 		dm.Add(de)
// 	}
// }
//
// func (dm *DataMap) FromRefAddState(ref string, parentId string, pid string, name string) {
// 	for range only.Once {
// 		pe := dm.GetEntry(ref, 0)
// 		if pe.IsNotValid() {
// 			fmt.Printf("ERROR: FromRefAddState('%s', '%s', '%s', '%s')\n", ref, parentId, pid, name)
// 			break
// 		}
//
// 		de := CopyDataEntry(*pe, pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name, pe.Point.GroupName, pe.Point.Unit, pe.Point.ValueType)
// 		de.MakeState(pe.Value.ValueBool())
// 		// de := pe.CreateState(pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name)
// 		dm.Add(de)
// 	}
// }
//
// func (dm *DataMap) FromRefAddFloat(ref string, parentId string, pid string, name string, value float64) {
// 	for range only.Once {
// 		pe := dm.GetEntry(ref, 0)
// 		if pe.IsNotValid() {
// 			fmt.Printf("ERROR: FromRefAddFloat('%s', '%s', '%s', '%s')\n", ref, parentId, pid, name)
// 			break
// 		}
//
// 		de := CopyDataEntry(*pe, pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name, pe.Point.GroupName, pe.Point.Unit, pe.Point.ValueType)
// 		de.MakeFloat(value, "", "")
// 		// de := pe.CreateFloat(pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name, value)
// 		dm.Add(de)
// 	}
// }
//
// func CopyDataEntry(ref DataEntry, endpoint string, parentId string, pid valueTypes.PointId, name string, groupName string, unit string, Type string) DataEntry {
// 	var ret DataEntry
// 	for range only.Once {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
//
// 		point := CopyPoint(*ref.Point, parentId, pid, name, groupName, unit, Type)
// 		// point = &Point {
// 		// 	Parents:   de.Point.Parents,
// 		// 	Id:        pid,
// 		// 	GroupName: "alias",
// 		// 	Name:      name,
// 		// 	Unit:      de.Point.Unit,
// 		// 	UpdateFreq:  de.Point.UpdateFreq,
// 		// 	ValueType: de.Point.ValueType,
// 		// 	Valid:     true,
// 		// 	States:    de.Point.States,
// 		// }
// 		// var parent ParentDevice
// 		// parent.Set(parentId)
// 		// point.Parents.Add(parent)
// 		// point.Unit = "binary"
// 		// if point.Unit == "" {
// 		// 	point.Unit = ref.Unit()
// 		// }
// 		// point.Name = name
// 		// if point.Name == "" {
// 		// 	point.Name = pid.PointToName()
// 		// }
// 		// // if de2.Point.GroupName == "" {
// 		// // 	de2.Point.GroupName = groupName
// 		// // }
// 		// point.FixUnitType()
// 		// point.Valid = true
//
// 		ret.Point = point
// 		ret.EndPoint = endpoint
// 		ret.Parent.Set(parentId)
// 		ret.Valid = true
// 		ret.Hide = false
// 	}
//
// 	return ret
// }
//
// func CopyPoint(ref Point, parentId string, pid valueTypes.PointId, name string, groupName string, unit string, Type string) *Point {
// 	for range only.Once {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
//
// 		var parent ParentDevice
// 		parent.Set(parentId)
// 		ref.Parents.Add(parent)
// 		ref.Id = pid
// 		ref.Unit = unit
// 		ref.Name = name
// 		ref.UpdateFreq = ""
// 		ref.GroupName = groupName
// 		ref.ValueType = Type
// 		ref.Valid = true
// 		ref.States = nil
//
// 		ref.FixUnitType()
// 	}
//
// 	return &ref
// }
//
// func (dm *DataMap) HideEntry(pointId valueTypes.PointId) {
// 	for range only.Once {
// 		de := dm.GetEntryFromPointId(pointId)
// 		de.Hide()
// 	}
// }
//
// func (dm *DataMap) AddEntry(endpoint string, parentId string, point Point, date valueTypes.DateTime, value string) {
// 	for range only.Once {
// 		unit := point.Unit	// Save unit.
// 		vType := point.ValueType	// Save type.
//
// 		// Match to a previously defined point.
// 		p := GetPoint(point.Id.String())
// 		if p != nil {
// 			// No point found. Create one.
// 			p = CreatePoint(parentId, pid, name, groupName, unit, Type)
// 		}
// 		point = *p
//
// 		// var parents ParentDevices
// 		// parents.Add(ParentDevice{Key: device})
// 		var parent ParentDevice
// 		parent.Set(parentId)
// 		point.Parents.Add(parent)
//
// 		if point.Name == "" {
// 			point.Name = point.Id.PointToName()
// 		}
// 		// fid := JoinDevicePoint(parent.Key, point.Id)
// 		ref := valueTypes.SetUnitValueString(value, unit, vType)
// 		point.Unit = ref.Unit()
// 		point.Valid = true
//
// 		if _, ok := dm.DataPoints[point.Id.String()]; ok {
// 			fmt.Printf("BARF: %s\n", point.Id)
// 		}
//
// 		// dm.Add(JoinDevicePoint(endpoint, point.Id), DataEntry {
// 		dm.Add(DataEntry {
// 			EndPoint:   endpoint,
// 			// FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
// 			// FullId:     JoinDevicePoint(parent.Key, point.Id),
// 			Parent:     parent,
//
// 			Point:      &point,
// 			Date:       date,
// 			Value:      ref.String(),
// 			ValueFloat: ref.Value(),
// 			ValueBool:  ref.ValueBool(),
// 			Index:      0,
// 			Valid:      true,
// 			Hide:       false,
// 		})
// 	}
// }
//
// func (dm *DataMap) AddUnitValue(endpoint string, parentId string, pid valueTypes.PointId, name string, groupName string, date valueTypes.DateTime, ref valueTypes.UnitValue) {
// 	for range only.Once {
// 		if endpoint == "" {
// 			endpoint = GoStruct.GetCallerPackage(2)
// 		}
//
// 		ref = ref.UnitValueFix()
//
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
//
// 		point := GetPoint(pid.String())
// 		if point == nil {
// 			// No point found. Create one.
// 			point = CreatePoint(parentId, pid, name, groupName, ref.Unit(), ref.Type())
// 			// de := CreateDataEntry(endpoint, parentId, pid, name, groupName, date, ref)
// 			// dm.Add(de)
// 			// break
// 		}
//
// 		var parent ParentDevice
// 		parent.Set(parentId)
// 		point.Parents.Add(parent)
// 		if point.Unit == "" {
// 			point.Unit = ref.Unit()
// 		}
// 		if point.Name == "" {
// 			point.Name = name
// 		}
// 		if point.Name == "" {
// 			point.Name = pid.PointToName()
// 		}
// 		if point.GroupName == "" {
// 			point.GroupName = groupName
// 		}
// 		point.FixUnitType()
// 		point.Valid = true
//
// 		dm.Add(DataEntry {
// 			EndPoint:   endpoint,
// 			// FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
// 			// FullId:     JoinDevicePoint(parent.Key, point.Id),
// 			Parent:     parent,
//
// 			Point:      point,
// 			Date:       date,
// 			Value:      ref.String(),
// 			ValueFloat: ref.Value(),
// 			ValueBool:  ref.ValueBool(),
// 			Index:      0,
// 			Valid:      true,
// 			Hide:       false,
// 		})
// 	}
// }
//
// func (dm *DataMap) AddFloat(endpoint string, parentId string, pid PointId, name string, date valueTypes.DateTime, value float64) {
// 	for range only.Once {
// 		// fvs := Float64ToString(value)
// 		point := GetPoint(parentId, pid)
// 		if point == nil {
// 			// No UV found. Create one.
// 			dm.Add(pid, CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)))
// 			break
// 		}
//
// 		ref := valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)
// 		if ref.Unit() != point.Unit {
// 			fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, ref.ValueFloat(), ref.Unit())
// 			point.Unit = ref.Unit()
// 		}
//
// 		var parent ParentDevice
// 		parent.Set(parentId)
// 		point.Parents.Add(parent)
//
// 		dm.Add(pid, DataEntry {
// 			EndPoint:   endpoint,
// 			FullId:     JoinDevicePoint(endpoint, point.Id),
// 			// FullId:     JoinDevicePoint(parent.Key, point.Id),
// 			Parent:     parent,
//
// 			Date:       date,
// 			Point:      point,
// 			Value:      ref.String(),
// 			ValueFloat: ref.Value(),
// 		})
// 	}
//
// 	uv := valueTypes.SetUnitValueFloat(value, "", "float")
// 	de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 	// de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, UnitValue {
// 	// 	Unit:       "float",
// 	// 	Value:      fmt.Sprintf("%f", value),
// 	// 	ValueFloat: 0,
// 	// })
// 	dm.Add(pid, de)
// }
//
// func (dm *DataMap) AddString(endpoint string, parentId string, pid PointId, name string, date valueTypes.DateTime, value string) {
// 	dm.Add(pid, CreateDataEntryString(date, endpoint, parentId, pid, name, value))
// }
//
// func (dm *DataMap) AddInt(endpoint string, parentId string, pid PointId, name string, date valueTypes.DateTime, value int64) {
//
// 	for range only.Once {
// 		uvs, ok := valueTypes.AnyToUnitValue(value, "", "")
// 		if !ok {
// 			fmt.Printf("ERROR: AddInt(endpoint '%s', parentId '%s', pid '%s', name '%s', date '%s', value %d)",
// 				endpoint, parentId, pid, name, date, value)
// 			break
// 		}
// 		for _, uv := range uvs {
// 			de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 			dm.Add(pid, de)
// 		}
//
// 		// uv := valueTypes.SetUnitValueInteger(value, "", "int")
// 		// de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 		// // de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, UnitValue {
// 		// // 	Unit:       "int",
// 		// // 	Value:      fmt.Sprintf("%d", value),
// 		// // 	ValueFloat: float64(value),
// 		// // })
// 		// dm.Add(pid, de)
// 	}
// }
//
// func (dm *DataMap) AddAny(endpoint string, parentId string, pid valueTypes.PointId, name string, date valueTypes.DateTime, value interface{}) {
//
// 	for range only.Once {
// 		uvs, isNil, ok := valueTypes.AnyToUnitValue(value, "", "")
// 		if !ok {
// 			fmt.Printf("ERROR: AddAny(endpoint '%s', parentId '%s', pid '%s', name '%s', date '%s', value '%v')",
// 				endpoint, parentId, pid, name, date, value)
// 			break
// 		}
//
// 		point := GetPoint(parentId + "." + pid.String())
// 		if point == nil {
// 			// No UV found. Create one.
// 			for _, uv := range uvs {
// 				de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 				if isNil {
// 					de.Point.ValueType += "(NIL)"
// 				}
// 				dm.Add(de)
// 			}
// 			// dm.Add(pid, CreateDataEntryUnitValue(date, endpoint, parentId, pid, name,
// 			// 	valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)))
// 			break
// 		}
//
// 		// ref := valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)
// 		// if ref.Unit() != point.Unit {
// 		// 	fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, ref.ValueFloat(), ref.Unit())
// 		// 	point.Unit = ref.Unit()
// 		// }
//
// 		if isNil {
// 			point.ValueType += "(NIL)"
// 		}
//
// 		for _, uv := range uvs {
// 			if uv.Unit() != point.Unit {
// 				fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, uv.ValueFloat(), uv.Unit())
// 				point.Unit = uv.Unit()
// 			}
//
// 			var parent ParentDevice
// 			parent.Set(parentId)
// 			point.Parents.Add(parent)
//
// 			// CreateDataEntry
// 			de := DataEntry {
// 				EndPoint: endpoint,
// 				// FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
// 				Parent:   parent,
//
// 				Date:       date,
// 				Point:      point,
// 				Value:      uv.String(),
// 				ValueFloat: uv.Value(),
// 				ValueBool:  uv.ValueBool(),
// 				Index:      0,
// 				Valid:      true,
// 				Hide:       false,
// 			}
// 			dm.Add(de)
// 		}
//
// 		for _, uv := range uvs {
// 			de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 			dm.Add(de)
// 		}
// 	}
// }
//
// func (de *DataEntry) CreateFloat(endpoint string, parentId string, pid valueTypes.PointId, name string, groupName string, unit string, Type string, value float64) DataEntry {
// 	var ret DataEntry
// 	for range only.Once {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
//
// 		ret = de.CreateDataEntry(endpoint, parentId, pid, name, groupName, unit, Type)
// 		uv := valueTypes.SetUnitValueFloat(value, ret.Point.Unit, ret.Point.ValueType)
// 		ret.Value = uv.String()
// 		ret.ValueFloat = uv.Value()
// 		ret.Valid = true
// 		ret.Hide = false
// 	}
// 	return ret
// }
//
// func (de *DataEntry) CreateState(endpoint string, parentId string, pid valueTypes.PointId, name string) DataEntry {
// 	var ret DataEntry
// 	for range only.Once {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
//
// 		de2 := de.CreateDataEntry(endpoint, parentId, pid, name)
// 		if de2.ValueFloat == 0 {
// 			de2.Value = "false"
// 			de2.ValueBool = false
// 			de2.ValueFloat = 0
// 		} else {
// 			de2.Value = "true"
// 			de2.ValueBool = true
// 			de2.ValueFloat = 1
// 		}
// 		de2.Valid = true
// 		de2.Hide = false
//
// 		var parent ParentDevice
// 		parent.Set(parentId)
// 		de2.Point.Parents.Add(parent)
// 		de2.Point.Unit = "binary"
// 		if de2.Point.Unit == "" {
// 			de2.Point.Unit = ref.Unit()
// 		}
// 		de2.Point.Name = name
// 		if de2.Point.Name == "" {
// 			de2.Point.Name = pid.PointToName()
// 		}
// 		// if de2.Point.GroupName == "" {
// 		// 	de2.Point.GroupName = groupName
// 		// }
// 		de2.Point.FixUnitType()
// 		de2.Point.Valid = true
// 	}
//
// 	return ret
// }
//
// func CreateDataEntryActive(date valueTypes.DateTime, endpoint string, parentId string, pid valueTypes.PointId, name string, value float64) DataEntry {
// 	point := GetPoint(parentId, pid)
// 	if point == nil {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
// 		point = CreatePoint(parentId, pid, name, "state")
// 	}
//
// 	var parent ParentDevice
// 	parent.Set(parentId)
// 	point.Parents.Add(parent)
//
// 	return DataEntry {
// 		EndPoint:   endpoint,
// 		FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
// 		// FullId:     JoinDevicePoint(parent.Key, point.Id),
// 		Parent:     parent,
//
// 		Point:      point,
// 		Date:       date,
// 		Value:      fmt.Sprintf("%v", IsActive(value)),
// 		ValueFloat: 0,
// 		Index:      0,
// 	}
// }
//
// func CreateDataEntryString(date valueTypes.DateTime, endpoint string, parentId string, pid valueTypes.PointId, name string, value string) DataEntry {
// 	point := GetPoint(parentId, pid)
// 	if point == nil {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
// 		point = CreatePoint(parentId, pid, name, "string")
// 	}
//
// 	var parent ParentDevice
// 	parent.Set(parentId)
// 	point.Parents.Add(parent)
//
// 	return DataEntry {
// 		EndPoint:   endpoint,
// 		FullId:     valueTypes.JoinDataPoint(endpoint, pid.String()),
// 		// FullId:     JoinDevicePoint(parent.Key, pid),
// 		Parent:     parent,
//
// 		Point:      point,
// 		Date:       date,
// 		Value:      value,
// 		ValueFloat: 0,
// 		Index:      0,
// 	}
// }
//
// func CreateDataEntryUnitValue(date valueTypes.DateTime, endpoint string, parentId string, pid valueTypes.PointId, name string, value valueTypes.UnitValue) DataEntry {
// 	value = value.UnitValueFix()
//
// 	point := GetPoint(parentId + "." + pid.String())
// 	if point == nil {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
// 		point = CreatePoint(parentId, pid, name, value.Unit())
// 	}
//
// 	var parent ParentDevice
// 	parent.Set(parentId)
// 	point.Parents.Add(parent)
// 	point.Valid = true
//
// 	return DataEntry {
// 		EndPoint:   endpoint,
// 		// FullId:     valueTypes.JoinDataPoint(endpoint, pid.String()),
// 		// FullId:     JoinDevicePoint(parent.Key, pid),
// 		Parent:     parent,
//
// 		Point:      point,
// 		Date:       date,
// 		Value:      value.String(),
// 		ValueFloat: value.Value(),
// 		ValueBool:  value.ValueBool(),
// 		Index:      0,
// 		Valid:      true,
// 		Hide:       false,
// 	}
// }
//
// func CreatePoint(parentId string, pid valueTypes.PointId, name string, unit string) *Point {
// 	if name == "" {
// 		name = pid.PointToName()
// 	}
//
// 	var parents ParentDevices
// 	parents.Add(ParentDevice{Key: parentId})
//
// 	ret := &Point {
// 		Parents:   parents,
// 		Id:        pid,
// 		GroupName: parentId,
// 		Name:      name,
// 		Unit:      unit,
// 		UpdateFreq:  "",
// 		ValueType: "",
// 		Valid:     true,
// 		States:    nil,
// 	}
// 	ret.FixUnitType()
//
// 	return ret
// }
//
// func IsActive(value float64) bool {
// 	if (value > 0.01) || (value < -0.01) {
// 		return true
// 	}
// 	return false
// }
//
// func JoinDevicePoint(endpoint string, pid valueTypes.PointId) valueTypes.PointId {
// 	var ret valueTypes.PointId
// 	for range only.Once {
// 		if endpoint == "" {
// 			endpoint = "virtual"
// 		}
// 		ret = valueTypes.PointId(JoinWithDots(0, "", endpoint, pid))
// 	}
// 	return ret
// }
//
// func JoinStringsWithDots(args ...string) string {
// 	return strings.Join(args, ".")
// }
