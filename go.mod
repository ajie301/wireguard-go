module golang.zx2c4.com/wireguard/windows

require (
	github.com/go-ole/go-ole v1.2.4
	github.com/lxn/walk v0.0.0-20191128110447-55ccb3a9f5c1
	github.com/lxn/win v0.0.0-20191128105842-2da648fda5b4
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	golang.org/x/sys v0.0.0-20191210023423-ac6580df4449
	golang.org/x/text v0.3.2
	golang.zx2c4.com/wireguard v0.0.20191013-0.20191128101113-ddfad453cf22
)

replace (
	github.com/lxn/walk => golang.zx2c4.com/wireguard/windows v0.0.0-20191128151049-87f28cc339ec
	github.com/lxn/win => golang.zx2c4.com/wireguard/windows v0.0.0-20191128151145-b4e4933852d5
)

go 1.13
