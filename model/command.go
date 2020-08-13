// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	COMMAND_METHOD_POST = "P"
	COMMAND_METHOD_GET  = "G"
	MIN_TRIGGER_LENGTH  = 1
	MAX_TRIGGER_LENGTH  = 128
)

type Command struct {
	Id               string `json:"id"`
	Token            string `json:"token"`
	CreateAt         int64  `json:"create_at"`
	UpdateAt         int64  `json:"update_at"`
	DeleteAt         int64  `json:"delete_at"`
	CreatorId        string `json:"creator_id"`
	TeamId           string `json:"team_id"`
	Trigger          string `json:"trigger"`
	Method           string `json:"method"`
	Username         string `json:"username"`
	IconURL          string `json:"icon_url"`
	AutoComplete     bool   `json:"auto_complete"`
	AutoCompleteDesc string `json:"auto_complete_desc"`
	AutoCompleteHint string `json:"auto_complete_hint"`
	DisplayName      string `json:"display_name"`
	Description      string `json:"description"`
	URL              string `json:"url"`
	// PluginId records the id of the plugin that created this Command. If it is blank, the Command
	// was not created by a plugin.
	PluginId         string            `json:"plugin_id"`
	AutocompleteData *AutocompleteData `db:"-" json:"autocomplete_data,omitempty"`
	// AutocompleteIconData is a base64 encoded svg
	AutocompleteIconData string `db:"-" json:"autocomplete_icon_data,omitempty"`
}

func (o *Command) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func CommandFromJson(data io.Reader) *Command {
	var o *Command
	json.NewDecoder(data).Decode(&o)
	return o
}

func CommandListToJson(l []*Command) string {
	b, _ := json.Marshal(l)
	return string(b)
}

func CommandListFromJson(data io.Reader) []*Command {
	var o []*Command
	json.NewDecoder(data).Decode(&o)
	return o
}

func (o *Command) IsValid() *AppError {

	if !IsValidId(o.Id) {
		return NewAppError("Command.IsValid", "model.command.is_valid.id.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.Token) != 26 {
		return NewAppError("Command.IsValid", "model.command.is_valid.token.app_error", nil, "", http.StatusBadRequest)
	}

	if o.CreateAt == 0 {
		return NewAppError("Command.IsValid", "model.command.is_valid.create_at.app_error", nil, "", http.StatusBadRequest)
	}

	if o.UpdateAt == 0 {
		return NewAppError("Command.IsValid", "model.command.is_valid.update_at.app_error", nil, "", http.StatusBadRequest)
	}

	// If the CreatorId is blank, this should be a command created by a plugin.
	if o.CreatorId == "" && !IsValidPluginId(o.PluginId) {
		return NewAppError("Command.IsValid", "model.command.is_valid.plugin_id.app_error", nil, "", http.StatusBadRequest)
	}

	// If the PluginId is blank, this should be a command associated with a userId.
	if o.PluginId == "" && !IsValidId(o.CreatorId) {
		return NewAppError("Command.IsValid", "model.command.is_valid.user_id.app_error", nil, "", http.StatusBadRequest)
	}

	if o.CreatorId != "" && o.PluginId != "" {
		return NewAppError("Command.IsValid", "model.command.is_valid.plugin_id.app_error", nil, "command cannot have both a CreatorId and a PluginId", http.StatusBadRequest)
	}

	if !IsValidId(o.TeamId) {
		return NewAppError("Command.IsValid", "model.command.is_valid.team_id.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.Trigger) < MIN_TRIGGER_LENGTH || len(o.Trigger) > MAX_TRIGGER_LENGTH || strings.Index(o.Trigger, "/") == 0 || strings.Contains(o.Trigger, " ") {
		return NewAppError("Command.IsValid", "model.command.is_valid.trigger.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.URL) == 0 || len(o.URL) > 1024 {
		return NewAppError("Command.IsValid", "model.command.is_valid.url.app_error", nil, "", http.StatusBadRequest)
	}

	u, err := url.ParseRequestURI(o.URL)
	if err != nil || !IsValidHttpUrl(o.URL) || u.Scheme == "" || u.Host == "" {
		return NewAppError("Command.IsValid", "model.command.is_valid.url_http.app_error", nil, "", http.StatusBadRequest)
	}

	if !(o.Method == COMMAND_METHOD_GET || o.Method == COMMAND_METHOD_POST) {
		return NewAppError("Command.IsValid", "model.command.is_valid.method.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.DisplayName) > 64 {
		return NewAppError("Command.IsValid", "model.command.is_valid.display_name.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.Description) > 128 {
		return NewAppError("Command.IsValid", "model.command.is_valid.description.app_error", nil, "", http.StatusBadRequest)
	}

	if o.AutocompleteData != nil {
		if err := o.AutocompleteData.IsValid(); err != nil {
			return NewAppError("Command.IsValid", "model.command.is_valid.autocomplete_data.app_error", nil, err.Error(), http.StatusBadRequest)
		}
	}

	return nil
}

func (o *Command) PreSave() {
	if o.Id == "" {
		o.Id = NewId()
	}

	if o.Token == "" {
		o.Token = NewId()
	}

	o.CreateAt = GetMillis()
	o.UpdateAt = o.CreateAt
}

func (o *Command) PreUpdate() {
	o.UpdateAt = GetMillis()
}

func (o *Command) Sanitize() {
	o.Token = ""
	o.CreatorId = ""
	o.Method = ""
	o.URL = ""
	o.Username = ""
	o.IconURL = ""
}
