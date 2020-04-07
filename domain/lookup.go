package domain

type LookupResponse struct {
	VendorDetails     VendorDetails     `json:"vendorDetails"`
	BlockDetails      BlockDetails      `json:"blockDetails"`
	MacAddressDetails MacAddressDetails `json:"macAddressDetails"`
}

type VendorDetails struct {
	Oui            string `json:"oui"`
	IsPrivate      bool   `json:"isPrivate"`
	CompanyName    string `json:"companyName"`
	CompanyAddress string `json:"companyAddress"`
	CountryCode    string `json:"countryCode"`
}

type BlockDetails struct {
	BlockFound          bool   `json:"blockFound"`
	BorderLeft          string `json:"borderLeft"`
	BorderRight         string `json:"borderRight"`
	BlockSize           int    `json:"blockSize"`
	AssignmentBlockSize string `json:"assignmentBlockSize"`
	DateCreated         string `json:"dateCreated"`
	DateUpdated         string `json:"dateUpdated"`
}

type MacAddressDetails struct {
	SearchTerm         string   `json:"searchTerm"`
	IsValid            bool     `json:"isValid"`
	VirtualMachine     string   `json:"virtualMachine"`
	Applications       []string `json:"applications"`
	TransmissionType   string   `json:"transmissionType"`
	AdministrationType string   `json:"administrationType"`
	WiresharkNotes     string   `json:"wiresharkNotes"`
	Comment            string   `json:"comment"`
}
