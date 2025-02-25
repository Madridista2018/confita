// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
)

// SendRemindMail sends mail with remind information.
func SendSignupMail(title string, content string, receiver string, domain string) error {
	sender := beego.AppConfig.String("appname")

	title = fmt.Sprintf("Re: [%s] %s", sender, title)

	return SendEmail(title, content, sender, receiver)
}

func SendEmail(title string, content string, sender string, receivers ...string) error {
	err := auth.SendEmail(title, content, sender, receivers...)
	return err
}
