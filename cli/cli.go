package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mr-doggy/doggy-sdk-go/doggy"
	"github.com/spf13/viper"
)

type Session struct {
	AccessToken string `json:"access_token" yaml:"access_token" mapstructure:"access_token"`
	ExpiresIn   int64  `json:"expires_in" yaml:"expires_in" mapstructure:"expires_in"`
	CreatedAt   int64  `json:"created_at" yaml:"created_at" mapstructure:"created_at"`
	TokenType   string `json:"token_type" yaml:"token_type" mapstructure:"token_type"`
	Scope       string `json:"scope" yaml:"scope" mapstructure:"scope"`
	PhoneNumber string `json:"phone_number" yaml:"phone_number" mapstructure:"phone_number"`
}

type Config struct {
	Host    string   `json:"host" yaml:"host" mapstructure:"host"`
	Session *Session `json:"session" yaml:"session" mapstructure:"session"`
}

func NewConfig() *Config {
	return &Config{
		Host:    "api.doggy.code2code.cn",
		Session: &Session{},
	}
}

func (session *Session) Valid() bool {
	if session.AccessToken == "" {
		return false
	}
	if session.ExpiresIn <= 0 {
		return false
	}
	if session.CreatedAt <= 0 {
		return false
	}
	return session.CreatedAt+session.ExpiresIn > time.Now().Unix()
}

type DoggyCli struct {
	api     *doggy.APIClient
	session *Session
	config  *Config

	AppOp     *AppOp
	ReleaseOp *ReleaseOp
}

func NewDoggyCli() *DoggyCli {
	cliCfg := NewConfig()
	err := viper.Unmarshal(&cliCfg)

	doggyCfg := doggy.NewConfiguration()
	doggyCfg.Scheme = "https"
	doggyCfg.Host = cliCfg.Host
	api := doggy.NewAPIClient(doggyCfg)

	if err != nil {
		panic(err)
	}
	if cliCfg.Session == nil {
		cliCfg.Session = &Session{}
	}
	if cliCfg.Session.Valid() {
		api.GetConfig().AddDefaultHeader("Authorization", cliCfg.Session.TokenType+" "+cliCfg.Session.AccessToken)
	}
	return &DoggyCli{
		api:       api,
		session:   cliCfg.Session,
		config:    cliCfg,
		AppOp:     NewAppOp(api),
		ReleaseOp: NewReleaseOp(api),
	}
}

func (cli *DoggyCli) Login() error {
	if cli.session.Valid() {
		return fmt.Errorf("??????????????? ??????????????????")
	}
	var phoneNumber string
	if err := survey.AskOne(&survey.Input{Message: "?????????????????????+86:"}, &phoneNumber); err != nil {
		return err
	}
	if !strings.HasPrefix("+86", phoneNumber) {
		phoneNumber = "+86" + phoneNumber
	}
	_, err := cli.api.SmsApi.ApiAppSmsSendLoginCodePost(context.Background()).
		SendSmsCodeDto(doggy.SendSmsCodeDto{
			PhoneNumber: *doggy.NewNullableString(&phoneNumber),
		}).Execute()
	if err != nil {
		return err
	}
	var smsCode string
	if err := survey.AskOne(&survey.Input{Message: "????????????????????????:"}, &smsCode); err != nil {
		return err
	}
	id, err := gonanoid.New()
	if err != nil {
		return err
	}
	deviceToken := id
	v := url.Values{}
	v.Set("grant_type", "sms")
	v.Set("scope", "Doggy")
	v.Set("client_id", "Doggy_Sms_GrantType")
	v.Set("client_secret", "1q2w3e*")
	v.Set("phone_number", phoneNumber)
	v.Set("sms_code", smsCode)
	v.Set("device_token", deviceToken)
	v.Set("device_name", "doggy-cli")
	v.Set("device_platform_type", "other")
	v.Set("device_brand", "doggy-cli")
	v.Set("device_system_version", "1.0.0")

	rsp, err := cli.api.GetConfig().HTTPClient.PostForm(cli.api.GetConfig().Scheme+"://"+cli.api.GetConfig().Host+"/connect/token", v)
	if err != nil {
		return err
	}
	if rsp.StatusCode > 300 {
		return fmt.Errorf("??????????????????????????????????????????????????????")
	}
	bts, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(bts))
	session := &Session{}
	err = json.Unmarshal(bts, &session)
	if err != nil {
		return err
	}
	cli.session = session
	session.PhoneNumber = phoneNumber
	session.CreatedAt = time.Now().Unix()
	viper.Set("session", session)
	return viper.WriteConfig()
}

func (cli *DoggyCli) Logout() error {
	viper.Set("session", nil)
	return viper.WriteConfig()
}
