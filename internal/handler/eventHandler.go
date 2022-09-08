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

type EventMessage struct {
	Content string
	Value   string
}

func (handler *EventHandler) Init(conf *setting.AppSetting) {
	handler.Configuration = conf
}

func (handler *EventHandler) Recive(msg string) *EventMessage {

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

func (handler *EventHandler) captureEvent(msg string) *EventMessage {
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

func analyzeEvent(msg string, cset *setting.UDPEventSetting) *EventMessage {
	reg := regexp.MustCompile(cset.CaptureReg)
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
	emsg := EventMessage{}
	content, value := cset.Content, cset.Value

	for k, v := range dic {
		log.Debug(fmt.Sprintf("key:%v ----->  value:%v", k, v))
		content = strings.Replace(content, "$"+k, v, -1)
		value = strings.Replace(value, "$"+k, v, -1)
	}
	emsg.Content = content
	emsg.Value = value

	return &emsg
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
