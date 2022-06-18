package params

type DockerListResp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	State       string `json:"state"`
	Status      string `json:"status"`
	Created     int64  `json:"created"`
	PrivatePort uint16  `json:"private_port"`
	PublicPort  uint16  `json:"public_port"`
}
