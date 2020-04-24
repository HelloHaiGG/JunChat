package utils

import (
	"github.com/mojocn/base64Captcha"
	"math/rand"
	"time"
)

/**
生成验证码，使用motocn/base64captcha
1.该项目只支持单机部署，分布式需要自己实现 Store 接口
2.该项目支持几种验证方式（digit/数字,string/字符串,math/数学,Chinese/中文）,其他需要 自己实现 Driver 接口
*/
func init() {
	//设置种子
	rand.Seed(time.Now().Unix())
}

/**
生成验证码id和base64编码字符串
*/
func CreateImageCode() (string, string) {

	//长度随机（4-6）
	//随机背景颜色

	cfg := base64Captcha.ConfigCharacter{
		Height:                 60,
		Width:                  180,
		Mode:                   base64Captcha.CaptchaModeNumberAlphabet, //数字和字符模式
		IsUseSimpleFont:        true,
		ComplexOfNoiseText:     base64Captcha.CaptchaComplexHigh, //干扰文本数量
		ComplexOfNoiseDot:      base64Captcha.CaptchaComplexHigh, //噪点计数
		IsShowHollowLine:       true,                             //显示空心线
		IsShowNoiseDot:         true,                             //显示噪点
		IsShowNoiseText:        true,                             //显示干扰文本
		IsShowSlimeLine:        true,                             //史莱姆线
		IsShowSineLine:         true,                             //正弦线
		ChineseCharacterSource: "",
		SequencedCharacters:    nil,
		UseCJKFonts:            false,
		CaptchaLen:             randomLen(),
		BgHashColor:            randomColor(),
		BgColor:                nil,
	}
	id, cap := base64Captcha.GenerateCaptcha("", cfg)
	return id, base64Captcha.CaptchaWriteToBase64Encoding(cap)
}

//通过验证码Id和验证码验证
func CheckCode(id, code string) bool {
	return base64Captcha.VerifyCaptcha(id, code)
}

func randomLen() int {
	//(4-6)
	return rand.Intn(3) + 4
}

func randomColor() string {
	//#000 - #FFF (0-16)
	pool := []string{
		"#000", "#111", "#222", "#333",
		"#444", "#555", "#666", "#777",
		"#888", "#999", "#AAA", "#BBB",
		"#CCC", "#DDD", "#EEE", "#FFF",
	}
	i := rand.Intn(16)
	return pool[i]
}
