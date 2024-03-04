package enums

type MaintenanceStatus string

const (
	Scheduled          MaintenanceStatus = "Scheduled"           // json:"maintenance_status"
	InProgress         MaintenanceStatus = "In Progress"         // json:"maintenance_status"
	Completed          MaintenanceStatus = "Completed"           // json:"maintenance_status"
	PendingApproval    MaintenanceStatus = "Pending Approval"    // json:"maintenance_status"
	Canceled           MaintenanceStatus = "Canceled"            // json:"maintenance_status"
	AwaitingParts      MaintenanceStatus = "Awaiting Parts"      // json:"maintenance_status"
	AwaitingPayment    MaintenanceStatus = "Awaiting Payment"    // json:"maintenance_status"
	Rescheduled        MaintenanceStatus = "Rescheduled"         // json:"maintenance_status"
	MaintenanceFailed  MaintenanceStatus = "Maintenance Failed"  // json:"maintenance_status"
	AwaitingInspection MaintenanceStatus = "Awaiting Inspection" // json:"maintenance_status"
)
