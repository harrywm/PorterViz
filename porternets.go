package portnernets

// PorterNetwork is an abstracted version of the Docker Network data type containing only the fields we are concerned with
type PorterNetwork struct {
	Name       string
	ID         string
	Driver     string
	Containers struct {
		Name  string
		ID    string
		Ports struct {
			IP          string
			PrivatePort string
			PublicPort  string
			Type        string
		}
		State string
	}
}

func (n *PorterNetwork) getInfo() string {
	return n.Name
}

func (n *PorterNetwork) getID() string {
	return n.ID
}

func (n *PorterNetwork) getDriver() string {
	return n.Driver
}
