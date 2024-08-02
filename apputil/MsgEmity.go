package apputil

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// 消息结构体,用于方法间传递消息
type MsgEmity struct {
	Gsuccess bool        `json:"success"`
	Gmsg     string      `json:"msg"`
	Gdata    interface{} `json:"data"`
}

/**
 * 构造消息结构体,返回包含'错误信息'的结构体
 * @param data 码值
 * @param msg 描述信息
 * @return 返回新创建的结构体
 */
func (m MsgEmity) Err(data int, msg ...interface{}) *MsgEmity {
	if len(msg) == 0 {
		return &MsgEmity{false, "", data}
	}

	var build strings.Builder
	for _, v := range msg {
		build.WriteString(fmt.Sprintf("%v", v))
	}

	return &MsgEmity{false, build.String(), data}
}

/**
 * 返回包含'错误信息'的结构体Json字符串
 * @param data 码值
 * @param msg 描述信息
 * @return 返回Json字符串
 */
func (m MsgEmity) ErrString(data int, msg ...interface{}) string {
	var build strings.Builder
	for _, v := range msg {
		build.WriteString(fmt.Sprintf("%v", v))
	}

	return m.Err(data, build.String()).ToStr()
}

/**
 * 构造消息结构体,返回包含'正确信息'的结构体
 * @param data 码值|数据
 * @param msg 描述信息
 * @return 返回新创建的结构体
 */
func (m MsgEmity) Success(data interface{}, msg ...interface{}) *MsgEmity {
	if len(msg) == 0 {
		return &MsgEmity{true, "", data}
	}

	var build strings.Builder
	for _, v := range msg {
		build.WriteString(fmt.Sprintf("%v", v))
	}

	return &MsgEmity{true, build.String(), data}
}

/**
 * 返回包含'正确信息'的结构体Json字符串
 * @param data 码值|数据
 * @param msg 描述信息
 * @return 返回Json字符串
 */
func (m MsgEmity) SuccessString(data interface{}, msg ...interface{}) string {
	text := ""
	if len(msg) != 0 {
		var build strings.Builder
		for _, v := range msg {
			build.WriteString(fmt.Sprintf("%v", v))
		}

		text = build.String()
	}

	return m.Success(data, text).ToStr()
}

/**
 * 重设'返回数据'
 * data 码值|数据
 * @return 返回原结构体 m
 */
func (m *MsgEmity) SetData(data interface{}) *MsgEmity {
	m.Gdata = data

	return m
}

/**
 * 重设'描述信息'
 * msg 描述信息
 * @return 返回原结构体 m
 */
func (m *MsgEmity) SetMsg(msg ...interface{}) *MsgEmity {
	result := m
	if len(msg) == 0 {
		return result
	}

	var build strings.Builder
	for _, v := range msg {
		build.WriteString(fmt.Sprintf("%v", v))
	}

	result.Gmsg = build.String()

	return result
}

/**
 * 累加'码值'
 * iCode 被累加值
 * @return 返回原结构体 m
 */
func (m *MsgEmity) IncCode(iCode int) *MsgEmity {
	result := m

	iData, err := strconv.Atoi(fmt.Sprintf("%v", result.Gdata)) //防止返回的类型未float64
	if err != nil {
		iData = 0
	}

	result.Gdata = iData + iCode

	return result
}

/**
 * 累加'码值'
 * iCode 被累加值
 * @return 返回原结构体 m
 */
func (m *MsgEmity) IncData(iCode int) int {
	if iCode == 0 {
		return m.Gdata.(int)
	}

	iData, err := strconv.Atoi(fmt.Sprintf("%v", m.Gdata)) //防止返回的类型未float64
	if err != nil {
		iData = 0
	}

	return iData + iCode
}

/**
 * 在描述信息后面累加'描述信息'
 * msg 描述信息
 * @return 返回原结构体 m
 */
func (m *MsgEmity) AppendMsg(msg ...interface{}) *MsgEmity {
	result := m

	var build strings.Builder
	build.WriteString(result.Gmsg)

	for _, v := range msg {
		build.WriteString(fmt.Sprintf("%v", v))
	}

	result.Gmsg = build.String()

	return result
}

/**
 * 在描述信息前面插入'描述信息'
 * msg 描述信息
 * @return 返回原结构体 m
 */
func (m *MsgEmity) InsertMsg(msg ...interface{}) *MsgEmity {
	result := m
	if len(msg) == 0 {
		return result
	}

	var build strings.Builder

	for _, v := range msg {
		build.WriteString(fmt.Sprintf("%v", v))
	}

	build.WriteString(result.Gmsg)

	result.Gmsg = build.String()

	return result
}

/**
 * 将结构体转换成json字符串输出
 * @return 返回json结构字符串
 */
func (m *MsgEmity) ToStr() string {
	ret_json, _ := json.Marshal(m)

	return string(ret_json)
}
