package dto

type Request struct {
	// TODO HGA: MISS GROUP PROPERTY HERE
	GroupPriority    int    `json:"group_priority"`
	GroupId          string `json:"group_id"`
	AgentVersion     string `json:"agent_version" binding:"required"`
	HostName         string `json:"hostname"`
	DeviceName       string `json:"device_name"`
	Hwid             string `json:"hwid"`
	Act              string `json:"act"`
	MachineType      string `json:"machinetype"`
	LicenseKey       string `json:"license_key"`
	IPAddress        string `json:"ipAddress"`
	DeviceType       string `json:"device_type"`
	SerialNumeber    string `json:"serial_number"`
	LinkedId         string `json:"linked_id"`
	InstallMode      string `json:"install_mode"`
	CustomID         string `json:"custom_id"`
	TSR              string `json:"tsr"`
	IsOpswatFreeTool bool   `json:"isOpswatFreeTool"`
	RegCode          string `json:"reg_code"`
	Reassessment     int    `json:"reassessment"`
}
