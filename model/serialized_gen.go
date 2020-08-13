// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Session) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 13 {
		err = msgp.ArrayError{Wanted: 13, Got: zb0001}
		return
	}
	z.Id, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	z.Token, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Token")
		return
	}
	z.CreateAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "CreateAt")
		return
	}
	z.ExpiresAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "ExpiresAt")
		return
	}
	z.LastActivityAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "LastActivityAt")
		return
	}
	z.UserId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "UserId")
		return
	}
	z.DeviceId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "DeviceId")
		return
	}
	z.Roles, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Roles")
		return
	}
	z.IsOAuth, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "IsOAuth")
		return
	}
	z.ExpiredNotify, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "ExpiredNotify")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "Props")
		return
	}
	if z.Props == nil {
		z.Props = make(StringMap, zb0002)
	} else if len(z.Props) > 0 {
		for key := range z.Props {
			delete(z.Props, key)
		}
	}
	for zb0002 > 0 {
		zb0002--
		var za0001 string
		var za0002 string
		za0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Props")
			return
		}
		za0002, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Props", za0001)
			return
		}
		z.Props[za0001] = za0002
	}
	var zb0003 uint32
	zb0003, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "TeamMembers")
		return
	}
	if cap(z.TeamMembers) >= int(zb0003) {
		z.TeamMembers = (z.TeamMembers)[:zb0003]
	} else {
		z.TeamMembers = make([]*TeamMember, zb0003)
	}
	for za0003 := range z.TeamMembers {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, "TeamMembers", za0003)
				return
			}
			z.TeamMembers[za0003] = nil
		} else {
			if z.TeamMembers[za0003] == nil {
				z.TeamMembers[za0003] = new(TeamMember)
			}
			err = z.TeamMembers[za0003].DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "TeamMembers", za0003)
				return
			}
		}
	}
	z.Local, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "Local")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Session) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 13
	err = en.Append(0x9d)
	if err != nil {
		return
	}
	err = en.WriteString(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	err = en.WriteString(z.Token)
	if err != nil {
		err = msgp.WrapError(err, "Token")
		return
	}
	err = en.WriteInt64(z.CreateAt)
	if err != nil {
		err = msgp.WrapError(err, "CreateAt")
		return
	}
	err = en.WriteInt64(z.ExpiresAt)
	if err != nil {
		err = msgp.WrapError(err, "ExpiresAt")
		return
	}
	err = en.WriteInt64(z.LastActivityAt)
	if err != nil {
		err = msgp.WrapError(err, "LastActivityAt")
		return
	}
	err = en.WriteString(z.UserId)
	if err != nil {
		err = msgp.WrapError(err, "UserId")
		return
	}
	err = en.WriteString(z.DeviceId)
	if err != nil {
		err = msgp.WrapError(err, "DeviceId")
		return
	}
	err = en.WriteString(z.Roles)
	if err != nil {
		err = msgp.WrapError(err, "Roles")
		return
	}
	err = en.WriteBool(z.IsOAuth)
	if err != nil {
		err = msgp.WrapError(err, "IsOAuth")
		return
	}
	err = en.WriteBool(z.ExpiredNotify)
	if err != nil {
		err = msgp.WrapError(err, "ExpiredNotify")
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Props)))
	if err != nil {
		err = msgp.WrapError(err, "Props")
		return
	}
	for za0001, za0002 := range z.Props {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Props")
			return
		}
		err = en.WriteString(za0002)
		if err != nil {
			err = msgp.WrapError(err, "Props", za0001)
			return
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.TeamMembers)))
	if err != nil {
		err = msgp.WrapError(err, "TeamMembers")
		return
	}
	for za0003 := range z.TeamMembers {
		if z.TeamMembers[za0003] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.TeamMembers[za0003].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "TeamMembers", za0003)
				return
			}
		}
	}
	err = en.WriteBool(z.Local)
	if err != nil {
		err = msgp.WrapError(err, "Local")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Session) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 13
	o = append(o, 0x9d)
	o = msgp.AppendString(o, z.Id)
	o = msgp.AppendString(o, z.Token)
	o = msgp.AppendInt64(o, z.CreateAt)
	o = msgp.AppendInt64(o, z.ExpiresAt)
	o = msgp.AppendInt64(o, z.LastActivityAt)
	o = msgp.AppendString(o, z.UserId)
	o = msgp.AppendString(o, z.DeviceId)
	o = msgp.AppendString(o, z.Roles)
	o = msgp.AppendBool(o, z.IsOAuth)
	o = msgp.AppendBool(o, z.ExpiredNotify)
	o = msgp.AppendMapHeader(o, uint32(len(z.Props)))
	for za0001, za0002 := range z.Props {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendString(o, za0002)
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.TeamMembers)))
	for za0003 := range z.TeamMembers {
		if z.TeamMembers[za0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.TeamMembers[za0003].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "TeamMembers", za0003)
				return
			}
		}
	}
	o = msgp.AppendBool(o, z.Local)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Session) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 13 {
		err = msgp.ArrayError{Wanted: 13, Got: zb0001}
		return
	}
	z.Id, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	z.Token, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Token")
		return
	}
	z.CreateAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "CreateAt")
		return
	}
	z.ExpiresAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "ExpiresAt")
		return
	}
	z.LastActivityAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LastActivityAt")
		return
	}
	z.UserId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "UserId")
		return
	}
	z.DeviceId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DeviceId")
		return
	}
	z.Roles, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Roles")
		return
	}
	z.IsOAuth, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "IsOAuth")
		return
	}
	z.ExpiredNotify, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "ExpiredNotify")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Props")
		return
	}
	if z.Props == nil {
		z.Props = make(StringMap, zb0002)
	} else if len(z.Props) > 0 {
		for key := range z.Props {
			delete(z.Props, key)
		}
	}
	for zb0002 > 0 {
		var za0001 string
		var za0002 string
		zb0002--
		za0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Props")
			return
		}
		za0002, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Props", za0001)
			return
		}
		z.Props[za0001] = za0002
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "TeamMembers")
		return
	}
	if cap(z.TeamMembers) >= int(zb0003) {
		z.TeamMembers = (z.TeamMembers)[:zb0003]
	} else {
		z.TeamMembers = make([]*TeamMember, zb0003)
	}
	for za0003 := range z.TeamMembers {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			z.TeamMembers[za0003] = nil
		} else {
			if z.TeamMembers[za0003] == nil {
				z.TeamMembers[za0003] = new(TeamMember)
			}
			bts, err = z.TeamMembers[za0003].UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "TeamMembers", za0003)
				return
			}
		}
	}
	z.Local, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Local")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Session) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.Id) + msgp.StringPrefixSize + len(z.Token) + msgp.Int64Size + msgp.Int64Size + msgp.Int64Size + msgp.StringPrefixSize + len(z.UserId) + msgp.StringPrefixSize + len(z.DeviceId) + msgp.StringPrefixSize + len(z.Roles) + msgp.BoolSize + msgp.BoolSize + msgp.MapHeaderSize
	if z.Props != nil {
		for za0001, za0002 := range z.Props {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += msgp.ArrayHeaderSize
	for za0003 := range z.TeamMembers {
		if z.TeamMembers[za0003] == nil {
			s += msgp.NilSize
		} else {
			s += z.TeamMembers[za0003].Msgsize()
		}
	}
	s += msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *StringMap) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0003 uint32
	zb0003, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if (*z) == nil {
		(*z) = make(StringMap, zb0003)
	} else if len((*z)) > 0 {
		for key := range *z {
			delete((*z), key)
		}
	}
	for zb0003 > 0 {
		zb0003--
		var zb0001 string
		var zb0002 string
		zb0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		zb0002, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		(*z)[zb0001] = zb0002
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z StringMap) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(uint32(len(z)))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0004, zb0005 := range z {
		err = en.WriteString(zb0004)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		err = en.WriteString(zb0005)
		if err != nil {
			err = msgp.WrapError(err, zb0004)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z StringMap) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, uint32(len(z)))
	for zb0004, zb0005 := range z {
		o = msgp.AppendString(o, zb0004)
		o = msgp.AppendString(o, zb0005)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StringMap) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if (*z) == nil {
		(*z) = make(StringMap, zb0003)
	} else if len((*z)) > 0 {
		for key := range *z {
			delete((*z), key)
		}
	}
	for zb0003 > 0 {
		var zb0001 string
		var zb0002 string
		zb0003--
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		zb0002, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		(*z)[zb0001] = zb0002
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z StringMap) Msgsize() (s int) {
	s = msgp.MapHeaderSize
	if z != nil {
		for zb0004, zb0005 := range z {
			_ = zb0005
			s += msgp.StringPrefixSize + len(zb0004) + msgp.StringPrefixSize + len(zb0005)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TeamMember) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "TeamId":
			z.TeamId, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "TeamId")
				return
			}
		case "UserId":
			z.UserId, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "UserId")
				return
			}
		case "Roles":
			z.Roles, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Roles")
				return
			}
		case "DeleteAt":
			z.DeleteAt, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "DeleteAt")
				return
			}
		case "SchemeGuest":
			z.SchemeGuest, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "SchemeGuest")
				return
			}
		case "SchemeUser":
			z.SchemeUser, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "SchemeUser")
				return
			}
		case "SchemeAdmin":
			z.SchemeAdmin, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "SchemeAdmin")
				return
			}
		case "ExplicitRoles":
			z.ExplicitRoles, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "ExplicitRoles")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TeamMember) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 8
	// write "TeamId"
	err = en.Append(0x88, 0xa6, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.TeamId)
	if err != nil {
		err = msgp.WrapError(err, "TeamId")
		return
	}
	// write "UserId"
	err = en.Append(0xa6, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.UserId)
	if err != nil {
		err = msgp.WrapError(err, "UserId")
		return
	}
	// write "Roles"
	err = en.Append(0xa5, 0x52, 0x6f, 0x6c, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.Roles)
	if err != nil {
		err = msgp.WrapError(err, "Roles")
		return
	}
	// write "DeleteAt"
	err = en.Append(0xa8, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.DeleteAt)
	if err != nil {
		err = msgp.WrapError(err, "DeleteAt")
		return
	}
	// write "SchemeGuest"
	err = en.Append(0xab, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x47, 0x75, 0x65, 0x73, 0x74)
	if err != nil {
		return
	}
	err = en.WriteBool(z.SchemeGuest)
	if err != nil {
		err = msgp.WrapError(err, "SchemeGuest")
		return
	}
	// write "SchemeUser"
	err = en.Append(0xaa, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteBool(z.SchemeUser)
	if err != nil {
		err = msgp.WrapError(err, "SchemeUser")
		return
	}
	// write "SchemeAdmin"
	err = en.Append(0xab, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteBool(z.SchemeAdmin)
	if err != nil {
		err = msgp.WrapError(err, "SchemeAdmin")
		return
	}
	// write "ExplicitRoles"
	err = en.Append(0xad, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.ExplicitRoles)
	if err != nil {
		err = msgp.WrapError(err, "ExplicitRoles")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TeamMember) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "TeamId"
	o = append(o, 0x88, 0xa6, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64)
	o = msgp.AppendString(o, z.TeamId)
	// string "UserId"
	o = append(o, 0xa6, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64)
	o = msgp.AppendString(o, z.UserId)
	// string "Roles"
	o = append(o, 0xa5, 0x52, 0x6f, 0x6c, 0x65, 0x73)
	o = msgp.AppendString(o, z.Roles)
	// string "DeleteAt"
	o = append(o, 0xa8, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74)
	o = msgp.AppendInt64(o, z.DeleteAt)
	// string "SchemeGuest"
	o = append(o, 0xab, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x47, 0x75, 0x65, 0x73, 0x74)
	o = msgp.AppendBool(o, z.SchemeGuest)
	// string "SchemeUser"
	o = append(o, 0xaa, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72)
	o = msgp.AppendBool(o, z.SchemeUser)
	// string "SchemeAdmin"
	o = append(o, 0xab, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e)
	o = msgp.AppendBool(o, z.SchemeAdmin)
	// string "ExplicitRoles"
	o = append(o, 0xad, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73)
	o = msgp.AppendString(o, z.ExplicitRoles)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TeamMember) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "TeamId":
			z.TeamId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TeamId")
				return
			}
		case "UserId":
			z.UserId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "UserId")
				return
			}
		case "Roles":
			z.Roles, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Roles")
				return
			}
		case "DeleteAt":
			z.DeleteAt, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DeleteAt")
				return
			}
		case "SchemeGuest":
			z.SchemeGuest, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SchemeGuest")
				return
			}
		case "SchemeUser":
			z.SchemeUser, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SchemeUser")
				return
			}
		case "SchemeAdmin":
			z.SchemeAdmin, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SchemeAdmin")
				return
			}
		case "ExplicitRoles":
			z.ExplicitRoles, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ExplicitRoles")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TeamMember) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.TeamId) + 7 + msgp.StringPrefixSize + len(z.UserId) + 6 + msgp.StringPrefixSize + len(z.Roles) + 9 + msgp.Int64Size + 12 + msgp.BoolSize + 11 + msgp.BoolSize + 12 + msgp.BoolSize + 14 + msgp.StringPrefixSize + len(z.ExplicitRoles)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *User) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 31 {
		err = msgp.ArrayError{Wanted: 31, Got: zb0001}
		return
	}
	z.Id, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	z.CreateAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "CreateAt")
		return
	}
	z.UpdateAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "UpdateAt")
		return
	}
	z.DeleteAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "DeleteAt")
		return
	}
	z.Username, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Username")
		return
	}
	z.Password, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Password")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "AuthData")
			return
		}
		z.AuthData = nil
	} else {
		if z.AuthData == nil {
			z.AuthData = new(string)
		}
		*z.AuthData, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "AuthData")
			return
		}
	}
	z.AuthService, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "AuthService")
		return
	}
	z.Email, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Email")
		return
	}
	z.EmailVerified, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "EmailVerified")
		return
	}
	z.Nickname, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Nickname")
		return
	}
	z.FirstName, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "FirstName")
		return
	}
	z.LastName, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "LastName")
		return
	}
	z.Position, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Position")
		return
	}
	z.Roles, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Roles")
		return
	}
	z.AllowMarketing, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "AllowMarketing")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "Props")
		return
	}
	if z.Props == nil {
		z.Props = make(StringMap, zb0002)
	} else if len(z.Props) > 0 {
		for key := range z.Props {
			delete(z.Props, key)
		}
	}
	for zb0002 > 0 {
		zb0002--
		var za0001 string
		var za0002 string
		za0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Props")
			return
		}
		za0002, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Props", za0001)
			return
		}
		z.Props[za0001] = za0002
	}
	var zb0003 uint32
	zb0003, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "NotifyProps")
		return
	}
	if z.NotifyProps == nil {
		z.NotifyProps = make(StringMap, zb0003)
	} else if len(z.NotifyProps) > 0 {
		for key := range z.NotifyProps {
			delete(z.NotifyProps, key)
		}
	}
	for zb0003 > 0 {
		zb0003--
		var za0003 string
		var za0004 string
		za0003, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "NotifyProps")
			return
		}
		za0004, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "NotifyProps", za0003)
			return
		}
		z.NotifyProps[za0003] = za0004
	}
	z.LastPasswordUpdate, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "LastPasswordUpdate")
		return
	}
	z.LastPictureUpdate, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "LastPictureUpdate")
		return
	}
	z.FailedAttempts, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "FailedAttempts")
		return
	}
	z.Locale, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Locale")
		return
	}
	var zb0004 uint32
	zb0004, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "Timezone")
		return
	}
	if z.Timezone == nil {
		z.Timezone = make(StringMap, zb0004)
	} else if len(z.Timezone) > 0 {
		for key := range z.Timezone {
			delete(z.Timezone, key)
		}
	}
	for zb0004 > 0 {
		zb0004--
		var za0005 string
		var za0006 string
		za0005, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Timezone")
			return
		}
		za0006, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Timezone", za0005)
			return
		}
		z.Timezone[za0005] = za0006
	}
	z.MfaActive, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "MfaActive")
		return
	}
	z.MfaSecret, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "MfaSecret")
		return
	}
	z.LastActivityAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "LastActivityAt")
		return
	}
	z.IsBot, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "IsBot")
		return
	}
	z.BotDescription, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "BotDescription")
		return
	}
	z.BotLastIconUpdate, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "BotLastIconUpdate")
		return
	}
	z.TermsOfServiceId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceId")
		return
	}
	z.TermsOfServiceCreateAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceCreateAt")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *User) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 31
	err = en.Append(0xdc, 0x0, 0x1f)
	if err != nil {
		return
	}
	err = en.WriteString(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	err = en.WriteInt64(z.CreateAt)
	if err != nil {
		err = msgp.WrapError(err, "CreateAt")
		return
	}
	err = en.WriteInt64(z.UpdateAt)
	if err != nil {
		err = msgp.WrapError(err, "UpdateAt")
		return
	}
	err = en.WriteInt64(z.DeleteAt)
	if err != nil {
		err = msgp.WrapError(err, "DeleteAt")
		return
	}
	err = en.WriteString(z.Username)
	if err != nil {
		err = msgp.WrapError(err, "Username")
		return
	}
	err = en.WriteString(z.Password)
	if err != nil {
		err = msgp.WrapError(err, "Password")
		return
	}
	if z.AuthData == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteString(*z.AuthData)
		if err != nil {
			err = msgp.WrapError(err, "AuthData")
			return
		}
	}
	err = en.WriteString(z.AuthService)
	if err != nil {
		err = msgp.WrapError(err, "AuthService")
		return
	}
	err = en.WriteString(z.Email)
	if err != nil {
		err = msgp.WrapError(err, "Email")
		return
	}
	err = en.WriteBool(z.EmailVerified)
	if err != nil {
		err = msgp.WrapError(err, "EmailVerified")
		return
	}
	err = en.WriteString(z.Nickname)
	if err != nil {
		err = msgp.WrapError(err, "Nickname")
		return
	}
	err = en.WriteString(z.FirstName)
	if err != nil {
		err = msgp.WrapError(err, "FirstName")
		return
	}
	err = en.WriteString(z.LastName)
	if err != nil {
		err = msgp.WrapError(err, "LastName")
		return
	}
	err = en.WriteString(z.Position)
	if err != nil {
		err = msgp.WrapError(err, "Position")
		return
	}
	err = en.WriteString(z.Roles)
	if err != nil {
		err = msgp.WrapError(err, "Roles")
		return
	}
	err = en.WriteBool(z.AllowMarketing)
	if err != nil {
		err = msgp.WrapError(err, "AllowMarketing")
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Props)))
	if err != nil {
		err = msgp.WrapError(err, "Props")
		return
	}
	for za0001, za0002 := range z.Props {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Props")
			return
		}
		err = en.WriteString(za0002)
		if err != nil {
			err = msgp.WrapError(err, "Props", za0001)
			return
		}
	}
	err = en.WriteMapHeader(uint32(len(z.NotifyProps)))
	if err != nil {
		err = msgp.WrapError(err, "NotifyProps")
		return
	}
	for za0003, za0004 := range z.NotifyProps {
		err = en.WriteString(za0003)
		if err != nil {
			err = msgp.WrapError(err, "NotifyProps")
			return
		}
		err = en.WriteString(za0004)
		if err != nil {
			err = msgp.WrapError(err, "NotifyProps", za0003)
			return
		}
	}
	err = en.WriteInt64(z.LastPasswordUpdate)
	if err != nil {
		err = msgp.WrapError(err, "LastPasswordUpdate")
		return
	}
	err = en.WriteInt64(z.LastPictureUpdate)
	if err != nil {
		err = msgp.WrapError(err, "LastPictureUpdate")
		return
	}
	err = en.WriteInt(z.FailedAttempts)
	if err != nil {
		err = msgp.WrapError(err, "FailedAttempts")
		return
	}
	err = en.WriteString(z.Locale)
	if err != nil {
		err = msgp.WrapError(err, "Locale")
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Timezone)))
	if err != nil {
		err = msgp.WrapError(err, "Timezone")
		return
	}
	for za0005, za0006 := range z.Timezone {
		err = en.WriteString(za0005)
		if err != nil {
			err = msgp.WrapError(err, "Timezone")
			return
		}
		err = en.WriteString(za0006)
		if err != nil {
			err = msgp.WrapError(err, "Timezone", za0005)
			return
		}
	}
	err = en.WriteBool(z.MfaActive)
	if err != nil {
		err = msgp.WrapError(err, "MfaActive")
		return
	}
	err = en.WriteString(z.MfaSecret)
	if err != nil {
		err = msgp.WrapError(err, "MfaSecret")
		return
	}
	err = en.WriteInt64(z.LastActivityAt)
	if err != nil {
		err = msgp.WrapError(err, "LastActivityAt")
		return
	}
	err = en.WriteBool(z.IsBot)
	if err != nil {
		err = msgp.WrapError(err, "IsBot")
		return
	}
	err = en.WriteString(z.BotDescription)
	if err != nil {
		err = msgp.WrapError(err, "BotDescription")
		return
	}
	err = en.WriteInt64(z.BotLastIconUpdate)
	if err != nil {
		err = msgp.WrapError(err, "BotLastIconUpdate")
		return
	}
	err = en.WriteString(z.TermsOfServiceId)
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceId")
		return
	}
	err = en.WriteInt64(z.TermsOfServiceCreateAt)
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceCreateAt")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *User) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 31
	o = append(o, 0xdc, 0x0, 0x1f)
	o = msgp.AppendString(o, z.Id)
	o = msgp.AppendInt64(o, z.CreateAt)
	o = msgp.AppendInt64(o, z.UpdateAt)
	o = msgp.AppendInt64(o, z.DeleteAt)
	o = msgp.AppendString(o, z.Username)
	o = msgp.AppendString(o, z.Password)
	if z.AuthData == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendString(o, *z.AuthData)
	}
	o = msgp.AppendString(o, z.AuthService)
	o = msgp.AppendString(o, z.Email)
	o = msgp.AppendBool(o, z.EmailVerified)
	o = msgp.AppendString(o, z.Nickname)
	o = msgp.AppendString(o, z.FirstName)
	o = msgp.AppendString(o, z.LastName)
	o = msgp.AppendString(o, z.Position)
	o = msgp.AppendString(o, z.Roles)
	o = msgp.AppendBool(o, z.AllowMarketing)
	o = msgp.AppendMapHeader(o, uint32(len(z.Props)))
	for za0001, za0002 := range z.Props {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendString(o, za0002)
	}
	o = msgp.AppendMapHeader(o, uint32(len(z.NotifyProps)))
	for za0003, za0004 := range z.NotifyProps {
		o = msgp.AppendString(o, za0003)
		o = msgp.AppendString(o, za0004)
	}
	o = msgp.AppendInt64(o, z.LastPasswordUpdate)
	o = msgp.AppendInt64(o, z.LastPictureUpdate)
	o = msgp.AppendInt(o, z.FailedAttempts)
	o = msgp.AppendString(o, z.Locale)
	o = msgp.AppendMapHeader(o, uint32(len(z.Timezone)))
	for za0005, za0006 := range z.Timezone {
		o = msgp.AppendString(o, za0005)
		o = msgp.AppendString(o, za0006)
	}
	o = msgp.AppendBool(o, z.MfaActive)
	o = msgp.AppendString(o, z.MfaSecret)
	o = msgp.AppendInt64(o, z.LastActivityAt)
	o = msgp.AppendBool(o, z.IsBot)
	o = msgp.AppendString(o, z.BotDescription)
	o = msgp.AppendInt64(o, z.BotLastIconUpdate)
	o = msgp.AppendString(o, z.TermsOfServiceId)
	o = msgp.AppendInt64(o, z.TermsOfServiceCreateAt)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *User) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 31 {
		err = msgp.ArrayError{Wanted: 31, Got: zb0001}
		return
	}
	z.Id, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	z.CreateAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "CreateAt")
		return
	}
	z.UpdateAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "UpdateAt")
		return
	}
	z.DeleteAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DeleteAt")
		return
	}
	z.Username, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Username")
		return
	}
	z.Password, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Password")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.AuthData = nil
	} else {
		if z.AuthData == nil {
			z.AuthData = new(string)
		}
		*z.AuthData, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "AuthData")
			return
		}
	}
	z.AuthService, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "AuthService")
		return
	}
	z.Email, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Email")
		return
	}
	z.EmailVerified, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "EmailVerified")
		return
	}
	z.Nickname, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Nickname")
		return
	}
	z.FirstName, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "FirstName")
		return
	}
	z.LastName, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LastName")
		return
	}
	z.Position, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Position")
		return
	}
	z.Roles, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Roles")
		return
	}
	z.AllowMarketing, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "AllowMarketing")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Props")
		return
	}
	if z.Props == nil {
		z.Props = make(StringMap, zb0002)
	} else if len(z.Props) > 0 {
		for key := range z.Props {
			delete(z.Props, key)
		}
	}
	for zb0002 > 0 {
		var za0001 string
		var za0002 string
		zb0002--
		za0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Props")
			return
		}
		za0002, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Props", za0001)
			return
		}
		z.Props[za0001] = za0002
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "NotifyProps")
		return
	}
	if z.NotifyProps == nil {
		z.NotifyProps = make(StringMap, zb0003)
	} else if len(z.NotifyProps) > 0 {
		for key := range z.NotifyProps {
			delete(z.NotifyProps, key)
		}
	}
	for zb0003 > 0 {
		var za0003 string
		var za0004 string
		zb0003--
		za0003, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "NotifyProps")
			return
		}
		za0004, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "NotifyProps", za0003)
			return
		}
		z.NotifyProps[za0003] = za0004
	}
	z.LastPasswordUpdate, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LastPasswordUpdate")
		return
	}
	z.LastPictureUpdate, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LastPictureUpdate")
		return
	}
	z.FailedAttempts, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "FailedAttempts")
		return
	}
	z.Locale, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Locale")
		return
	}
	var zb0004 uint32
	zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Timezone")
		return
	}
	if z.Timezone == nil {
		z.Timezone = make(StringMap, zb0004)
	} else if len(z.Timezone) > 0 {
		for key := range z.Timezone {
			delete(z.Timezone, key)
		}
	}
	for zb0004 > 0 {
		var za0005 string
		var za0006 string
		zb0004--
		za0005, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Timezone")
			return
		}
		za0006, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Timezone", za0005)
			return
		}
		z.Timezone[za0005] = za0006
	}
	z.MfaActive, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "MfaActive")
		return
	}
	z.MfaSecret, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "MfaSecret")
		return
	}
	z.LastActivityAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LastActivityAt")
		return
	}
	z.IsBot, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "IsBot")
		return
	}
	z.BotDescription, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "BotDescription")
		return
	}
	z.BotLastIconUpdate, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "BotLastIconUpdate")
		return
	}
	z.TermsOfServiceId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceId")
		return
	}
	z.TermsOfServiceCreateAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceCreateAt")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *User) Msgsize() (s int) {
	s = 3 + msgp.StringPrefixSize + len(z.Id) + msgp.Int64Size + msgp.Int64Size + msgp.Int64Size + msgp.StringPrefixSize + len(z.Username) + msgp.StringPrefixSize + len(z.Password)
	if z.AuthData == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.AuthData)
	}
	s += msgp.StringPrefixSize + len(z.AuthService) + msgp.StringPrefixSize + len(z.Email) + msgp.BoolSize + msgp.StringPrefixSize + len(z.Nickname) + msgp.StringPrefixSize + len(z.FirstName) + msgp.StringPrefixSize + len(z.LastName) + msgp.StringPrefixSize + len(z.Position) + msgp.StringPrefixSize + len(z.Roles) + msgp.BoolSize + msgp.MapHeaderSize
	if z.Props != nil {
		for za0001, za0002 := range z.Props {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += msgp.MapHeaderSize
	if z.NotifyProps != nil {
		for za0003, za0004 := range z.NotifyProps {
			_ = za0004
			s += msgp.StringPrefixSize + len(za0003) + msgp.StringPrefixSize + len(za0004)
		}
	}
	s += msgp.Int64Size + msgp.Int64Size + msgp.IntSize + msgp.StringPrefixSize + len(z.Locale) + msgp.MapHeaderSize
	if z.Timezone != nil {
		for za0005, za0006 := range z.Timezone {
			_ = za0006
			s += msgp.StringPrefixSize + len(za0005) + msgp.StringPrefixSize + len(za0006)
		}
	}
	s += msgp.BoolSize + msgp.StringPrefixSize + len(z.MfaSecret) + msgp.Int64Size + msgp.BoolSize + msgp.StringPrefixSize + len(z.BotDescription) + msgp.Int64Size + msgp.StringPrefixSize + len(z.TermsOfServiceId) + msgp.Int64Size
	return
}
