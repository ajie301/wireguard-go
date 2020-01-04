module golang.zx2c4.com/wireguard

go 1.12

require (
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	golang.org/x/sys v0.0.0-20191210023423-ac6580df4449
	golang.org/x/text v0.3.2
	golang.zx2c4.com/wireguard/windows v0.0.37
)

replace golang.zx2c4.com/wireguard/windows v0.0.37 => github.com/clouddeepcn/wireguard-windows v0.0.38-0.20200104061846-b0f41825e206
