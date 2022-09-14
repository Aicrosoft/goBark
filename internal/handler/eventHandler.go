package handler

import (
	"fmt"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/Aicrosoft/goBark/internal/setting"
)

type EventHandler struct {
	Configuration *setting.AppSetting
}

func (handler *EventHandler) Init(conf *setting.AppSetting) {
	handler.Configuration = conf
}

func (handler *EventHandler) Recive(msg string) *setting.MessageSetting {

	if handler.isIgnoreThis(msg) {
		//log.Info("ignore this")
		return nil
	}

	evt := handler.captureEvent(msg)
	if evt != nil {
		log.Info(fmt.Sprintf("Capture Event:%v", evt))
	}
	return evt

}

func (handler *EventHandler) captureEvent(msg string) *setting.MessageSetting {
	evntSettings := handler.Configuration.EventMessages
	if (len(evntSettings)) == 0 {
		return nil
	}
	for _, set := range evntSettings {
		event := analyzeEvent(msg, &set)
		if event != nil {
			return event
		}
	}

	return nil
}

func analyzeEvent(msg string, ues *setting.UDPEventSetting) *setting.MessageSetting {
	reg := regexp.MustCompile(ues.CaptureReg)
	matchValues := reg.FindStringSubmatch(msg)
	groupKeys := reg.SubexpNames()
	if (len(matchValues) == 0) || (len(matchValues) != len(groupKeys)) {
		return nil
	}
	dic := make(map[string]string)

	for i, name := range groupKeys {
		if i != 0 && name != "" {
			dic[name] = matchValues[i]
		}
	}
	em := setting.MessageSetting{}
	title, content, value := ues.Title, ues.Content, ues.Value

	for k, v := range dic {
		log.Debug(fmt.Sprintf("event capture result (k,v) %v : %v", k, v))
		title = strings.Replace(title, "$"+k, v, -1)
		content = strings.Replace(content, "$"+k, v, -1)
		value = strings.Replace(value, "$"+k, v, -1)
	}
	em.Title = title
	em.Content = content
	em.Value = value

	return &em
}

func (handler *EventHandler) isIgnoreThis(msg string) bool {
	keys := handler.Configuration.EventMessageIgnoreKeys
	if len(keys) == 0 {
		return false
	}
	for _, key := range keys {
		if strings.Contains(msg, key) {
			return true
		}
	}
	return false
}
