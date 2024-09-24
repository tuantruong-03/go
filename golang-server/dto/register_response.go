package dto

type Response struct {
	AgentToken   string `json:"agent_token"`
	DeviceId     string `json:"device_id"`
	LicenseKey   string `json:"license_key"`
	DeviceStatus string `json:"device_status"`
}
